package app

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/config"
	myGRPC "github.com/DenisKhanov/PrivateKeeper/internal/api/grpc/interceptors"
	protodata "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	protouser "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/user"
	"github.com/DenisKhanov/PrivateKeeper/pkg/logcfg"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// App represents the application structure responsible for initializing dependencies
// and running the serverGRPC.
type App struct {
	serviceProvider *serviceProvider  // The service provider for dependency injection
	config          *config.ENVConfig // The configuration object for the application
	dbPool          *pgxpool.Pool     // The connection pool to the database
	minioClient     *minio.Client     // The client s3 minio storage
	trustedSubnets  []*net.IPNet      // The collection trusted subnet
	serverGRPC      *grpc.Server      // The serverGRPC instance
}

// NewApp creates a new instance of the application.
func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Run starts the application and runs the grpc_keeper serverGRPC.
func (a *App) Run() {
	a.runKeeperServer()
}

// initDeps initializes all dependencies required by the application.
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initTrustedSubnets,
		a.initDBConnection,
		a.initS3Client,
		a.initServiceProvider,
		a.initKeeperGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// initConfig initializes the application configuration.
func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	a.config = cfg
	config.PrintProjectInfo()
	return nil
}

// parseSubnets parses a string containing a list of CIDR subnets and returns them as a []*net.IPNet objects.
func (a *App) initTrustedSubnets(_ context.Context) error {
	var subnets []*net.IPNet
	if a.config.EnvSubnet != "" {
		subStr := strings.Split(a.config.EnvSubnet, ",")
		for _, subnetStr := range subStr {
			_, subnetIPNet, err := net.ParseCIDR(subnetStr)
			if err != nil {
				logrus.WithError(err).Error("error parsing string CIDR")
				return err
			}
			subnets = append(subnets, subnetIPNet)
		}
	}
	a.trustedSubnets = subnets
	return nil
}

// initDBConnection initializes the connection to the database.
func (a *App) initDBConnection(ctx context.Context) error {
	logrus.Infof("DB config: %+v", a.config.EnvDataBase)
	if a.config.EnvDataBase != "" {
		confPool, err := pgxpool.ParseConfig(a.config.EnvDataBase)
		if err != nil {
			logrus.WithError(err).Error("Error parsing config")
			return err
		}
		confPool.MaxConns = 50
		confPool.MinConns = 10
		a.dbPool, err = pgxpool.NewWithConfig(ctx, confPool)
		if err != nil {
			logrus.WithError(err).Error("Don't connect to DB")
			return err
		}
	} else {
		logrus.Infof("config EnvDataBase is empty")
	}
	return nil
}

// initS3Client initializes the s3 client
func (a *App) initS3Client(_ context.Context) error {
	minioClient, err := minio.New(a.config.EnvS3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(a.config.EnvS3AccessKey, a.config.EnvS3SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.WithError(err).Errorf("Create minio client failed on: %v", a.config.EnvS3Endpoint)
		return err
	}
	a.minioClient = minioClient
	return nil
}

// initServiceProvider initializes the service provider for dependency injection.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initKeeperGRPCServer initializes the  serverGRPC with interceptors.
func (a *App) initKeeperGRPCServer(_ context.Context) error {
	keeperUserGRPC := a.serviceProvider.KeeperUserGRPC(a.dbPool, a.config.EnvStoragePath)
	keeperDataGRPC := a.serviceProvider.KeeperDataGRPC(a.dbPool, a.minioClient, a.config.EnvStoragePath, a.config.EnvS3Bucket)

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(myGRPC.UnaryLoggerInterceptor,
		myGRPC.UnaryTrustedSubnetsInterceptor(a.trustedSubnets),
		myGRPC.UnaryPrivateAuthInterceptor),
	)
	reflection.Register(server)
	a.serverGRPC = server

	// registration service
	protodata.RegisterKeeperDataV1Server(server, keeperDataGRPC)
	protouser.RegisterKeeperUserV1Server(server, keeperUserGRPC)
	return nil

}

// runKeeperServer starts the gRPC server with graceful shutdown.
func (a *App) runKeeperServer() {
	logcfg.RunLoggerConfig(a.config.EnvLogLevel)

	//run gRPC server
	go func() {
		listen, err := net.Listen("tcp", a.config.EnvGRPC)
		if err != nil {
			logrus.Error(err)
		}

		logrus.Infof("Starting server gRPC on: %s", a.config.EnvGRPC)
		if err = a.serverGRPC.Serve(listen); err != nil {
			logrus.WithError(err).Error("The server gRPC  failed to start")
		}
	}()

	// Shutdown signal with grace period of 5 seconds
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-signalChan
	logrus.Infof("Shutting down HTTP & gRPC servers with signal : %v...", sig)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		a.serverGRPC.GracefulStop()
		logrus.Infof("gRPC servers closed")
		wg.Done()
	}()

	//TODO избавиться от приведения типов

	//If the input shutdown signal, batch URLs saving to file
	if a.dbPool == nil {
		//err := a.serviceProvider.repositoryUser.(url.InMemoryRepository).SaveBatchToFile()
		//if err != nil {
		//	logrus.WithError(err).Error("Error save memory in file")
		//}
	} else {
		a.dbPool.Close()
	}
	wg.Wait()
	logrus.Info("Server exited")
}
