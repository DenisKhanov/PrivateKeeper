package client

import (
	"bufio"
	"fmt"
	"github.com/DenisKhanov/PrivateKeeper/pkg/logcfg"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	binarypb "github.com/DenisKhanov/PrivateKeeper/internal/client/binary_data/pbclient"
	binaryservice "github.com/DenisKhanov/PrivateKeeper/internal/client/binary_data/service"
	"github.com/DenisKhanov/PrivateKeeper/internal/client/config"
	creditcardpb "github.com/DenisKhanov/PrivateKeeper/internal/client/credit_card/pbclient"
	creditcardservice "github.com/DenisKhanov/PrivateKeeper/internal/client/credit_card/service"
	datapb "github.com/DenisKhanov/PrivateKeeper/internal/client/data/pbclient"
	credentialspb "github.com/DenisKhanov/PrivateKeeper/internal/client/login_password/pbclient"
	credentialsservice "github.com/DenisKhanov/PrivateKeeper/internal/client/login_password/service"
	"github.com/DenisKhanov/PrivateKeeper/internal/client/state"
	textdatapb "github.com/DenisKhanov/PrivateKeeper/internal/client/text_data/pbclient"
	textdataservice "github.com/DenisKhanov/PrivateKeeper/internal/client/text_data/service"
	userpb "github.com/DenisKhanov/PrivateKeeper/internal/client/user/pbclient"

	dataservice "github.com/DenisKhanov/PrivateKeeper/internal/client/data/service"
	userservice "github.com/DenisKhanov/PrivateKeeper/internal/client/user/service"

	"github.com/DenisKhanov/PrivateKeeper/internal/proto/binary_data"
	"github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	"github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/user"
	"github.com/DenisKhanov/PrivateKeeper/pkg/tlsconfig"
)

func Run() {

	cfg, err := config.New()
	if err != nil {
		log.Println("Failed to initialize config", err.Error())
		os.Exit(1)
	}

	logFileName := "keeperClient.log"
	logcfg.RunLoggerConfig(cfg.EnvLogLevel, logFileName)

	clientTLS, err := tlsconfig.NewClientTLS(cfg.ClientCert, cfg.ClientKey, cfg.ClientCa)
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize clientTLS")
		os.Exit(1)
	}

	grpcClient, err := grpc.NewClient(cfg.GRPCServer, grpc.WithTransportCredentials(credentials.NewTLS(clientTLS)))
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize grpcClient")
		os.Exit(1)
	}

	clientState := state.NewClientState()

	userClient := userpb.NewUserPBClient(KeeperUserV1.NewKeeperUserV1Client(grpcClient))
	userService := userservice.NewUserService(userClient, clientState)

	dataClient := datapb.NewDataPBClient(KeeperDataV1.NewKeeperDataV1Client(grpcClient))
	dataService := dataservice.NewDataService(dataClient, clientState)

	creditCardClient := creditcardpb.NewCreditCardPBClient(KeeperDataV1.NewKeeperDataV1Client(grpcClient))
	creditCardService := creditcardservice.NewUserService(creditCardClient, clientState)

	textDataClient := textdatapb.NewCreditCardPBClient(KeeperDataV1.NewKeeperDataV1Client(grpcClient))
	textDataService := textdataservice.NewTextDataService(textDataClient, clientState)

	loginPasswordClient := credentialspb.NewCredentialsPBClient(KeeperDataV1.NewKeeperDataV1Client(grpcClient))
	loginPasswordService := credentialsservice.NewCredentialsService(loginPasswordClient, clientState)

	binaryClient := binarypb.NewBinaryDataPBClient(binary_data.NewBinaryDataServiceClient(grpcClient))
	binaryService := binaryservice.NewBinaryDataService(binaryClient, clientState)

	scanner := bufio.NewScanner(os.Stdin)

	blue := color.New(color.FgBlue).SprintFunc()

	for {
		if clientState.IsAuthorized() {
			fmt.Printf("You are authorized as %s\n", blue(clientState.GetLogin()))
		} else {
			fmt.Printf("You are not authorized, please login or register\n")
		}

		if clientState.GetDirPath() == "" {
			fmt.Printf("Working directory is not set \n")
		} else {
			fmt.Printf("Working directory is set to %s\n", blue(clientState.GetDirPath()))
		}

		fmt.Println("Input command number to proceed")
		fmt.Println("[0] - quit")
		fmt.Println("[1] - login")
		fmt.Println("[2] - register")
		fmt.Println("[3] - look all data info")
		fmt.Println("[4] - save credit card")
		fmt.Println("[5] - load credit cards")
		fmt.Println("[6] - save text data")
		fmt.Println("[7] - load text data")
		fmt.Println("[8] - save login/password")
		fmt.Println("[9] - load login/password")
		fmt.Println("[10] - save binary file")
		fmt.Println("[11] - load binary files")
		fmt.Println("[12] - set working directory")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "0":
			fmt.Println("Application shutdown.")
			return
		case "1":
			userService.LoginUser()
		case "2":
			userService.RegisterUser()
		case "3":
			dataService.LoadAllDataInfo()
		case "4":
			creditCardService.Save()
		case "5":
			creditCardService.Load()
		case "6":
			textDataService.Save()
		case "7":
			textDataService.Load()
		case "8":
			loginPasswordService.Save()
		case "9":
			loginPasswordService.Load()
		case "10":
			binaryService.Save()
		case "11":
			binaryService.Load()
		case "12":
			clientState.SetWorkingDirectory()
		default:
			fmt.Println("Unknown command, please try again")
		}
	}
}
