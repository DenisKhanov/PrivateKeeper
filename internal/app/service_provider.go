package app

import (
	grpcdata "github.com/DenisKhanov/PrivateKeeper/internal/api/grpc/data"
	grpcuser "github.com/DenisKhanov/PrivateKeeper/internal/api/grpc/user"
	"github.com/DenisKhanov/PrivateKeeper/internal/repositories"
	repodata "github.com/DenisKhanov/PrivateKeeper/internal/repositories/database/data"
	repouser "github.com/DenisKhanov/PrivateKeeper/internal/repositories/database/user"
	"github.com/DenisKhanov/PrivateKeeper/internal/repositories/s3"
	"github.com/DenisKhanov/PrivateKeeper/internal/services"
	srevicedata "github.com/DenisKhanov/PrivateKeeper/internal/services/data"
	sreviceuser "github.com/DenisKhanov/PrivateKeeper/internal/services/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

// serviceProvider manages the dependency injection for http_shortener-related components.
type serviceProvider struct {
	repositoryUser   repositories.UserRepository // Repository interface for grpc_keeper-related data
	repositoryData   repositories.DataRepository // Repository interface for grpc_keeper-related data
	repositoryS3Data repositories.S3Repository   // Repository interface for grpc_keeper-related data
	serviceUser      services.UserService        // Service interface for grpc_keeper-related operations
	serviceData      services.DataService        // Service interface for grpc_keeper-related operations
	grpcUser         *grpcuser.GRPCUser          //GRPC for grpc_keeper-related operations
	grpcData         *grpcdata.GRPCData          //GRPC for grpc_keeper-related operations
}

// newServiceProvider creates a new instance of the service provider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// KeeperUserRepository returns the repository for user-related data.
// If dbPool is nil, it initializes an in-memory repository, otherwise initializes a database repository.
func (s *serviceProvider) KeeperUserRepository(dbPool *pgxpool.Pool, storagePath string) repositories.UserRepository {
	logrus.Info("Creating new Keeper User Repository")
	var err error
	if s.repositoryUser == nil {
		logrus.Info("Initializing repository user")
		if dbPool == nil {
			logrus.Infof("DB pool is nil")
			//TODO продумать хранение на жестком
			//s.repositoryUser = url2.NewURLInMemoryRepo(storagePath)
		} else {
			if s.repositoryUser, err = repouser.NewPostgresUser(dbPool); err != nil {
				//TODO лучше вернуть ошибку из метода и обработать ее выше
				logrus.Fatal(err)
			}
		}
	}
	return s.repositoryUser
}

// KeeperDataRepository returns the repository for data-related data.
// If dbPool is nil, it initializes an in-memory repository, otherwise initializes a database repository.
func (s *serviceProvider) KeeperDataRepository(dbPool *pgxpool.Pool, storagePath string) repositories.DataRepository {
	logrus.Info("Creating new Keeper Data Repository.")
	if s.repositoryData == nil {
		logrus.Info("Initializing repository data.")
		if dbPool == nil {
			logrus.Infof("DB pool is nil.")
			//TODO продумать хранение на жестком
			//s.repositoryUser = url2.NewURLInMemoryRepo(storagePath)
		} else {
			s.repositoryData = repodata.NewPostgresData(dbPool)
		}
	}
	return s.repositoryData
}

// KeeperS3Repository returns the repository for S3-related data.
func (s *serviceProvider) KeeperS3Repository(client *minio.Client, bucket string) repositories.S3Repository {
	var err error
	s.repositoryS3Data, err = s3.NewMinio(client, bucket)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Creating new Keeper S3 Repository.")

	return s.repositoryS3Data
}

// KeeperUserService returns the service for user-related operations.
func (s *serviceProvider) KeeperUserService(dbPool *pgxpool.Pool, storagePath string) services.UserService {
	if s.serviceUser == nil {
		s.serviceUser = sreviceuser.NewServiceUser(
			s.KeeperUserRepository(dbPool, storagePath),
		)
	}
	return s.serviceUser
}

// KeeperDataService returns the service for data-related operations.
func (s *serviceProvider) KeeperDataService(dbPool *pgxpool.Pool, client *minio.Client, storagePath, bucket string) services.DataService {
	if s.serviceData == nil {
		s.serviceData = srevicedata.NewServiceData(
			s.KeeperDataRepository(dbPool, storagePath),
			s.KeeperS3Repository(client, bucket),
			dbPool,
		)
	}
	return s.serviceData
}

// KeeperUserGRPC returns the handler for user-related HTTP endpoints.
func (s *serviceProvider) KeeperUserGRPC(dbPool *pgxpool.Pool, storagePath string) *grpcuser.GRPCUser {
	logrus.Info("Creating Keeper GRPC.")
	if s.grpcUser == nil {
		logrus.Info("Initializing grpc user.")
		keeperGRPC := grpcuser.NewGRPCUser(s.KeeperUserService(dbPool, storagePath))
		s.grpcUser = keeperGRPC
	}
	return s.grpcUser
}

// KeeperDataGRPC returns the handler for data-related HTTP endpoints.
func (s *serviceProvider) KeeperDataGRPC(dbPool *pgxpool.Pool, client *minio.Client, storagePath, bucket string) *grpcdata.GRPCData {
	logrus.Info("Creating Keeper GRPC.")
	if s.grpcData == nil {
		logrus.Info("Initializing grpc data.")
		keeperGRPC := grpcdata.NewGRPCData(s.KeeperDataService(dbPool, client, storagePath, bucket))
		s.grpcData = keeperGRPC
	}
	return s.grpcData
}
