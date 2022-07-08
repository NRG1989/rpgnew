package main

import (
	"context"
	"flag"

	"go-aut-registration-user-otp/internal/api/grpc"
	api "go-aut-registration-user-otp/internal/api/http"
	"go-aut-registration-user-otp/internal/config"
	"go-aut-registration-user-otp/internal/database"

	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	logger := logrus.New()

	configFilePath := flag.String("config", "./config.json", "path to configuration file")
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		logger.WithError(err).Fatal("reading config file error")
	}

	db, err := database.NewClient(cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("constructing database error")
	}
	storage := database.NewStorage(db)
	logger.Info("successfully connected to DB!!")

	otpServer := api.NewOtpServer(ctx, cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("constructing bot error")
	}

	go func() {
		logger.Info("starting 5014 gRPC server")
		if err := grpc.NewGRPCService(logger, cfg.GrpcCfg.AddressOTPSrv, storage).Start(); err != nil {
			logger.WithError(err).Fatal("connection to grpc error")
		}
	}()

	otpServer.StartServer()
}
