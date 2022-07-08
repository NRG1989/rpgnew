package http

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"go-aut-registration-user-otp/internal/config"
)

type otpServer struct {
	*echo.Echo
	ctx    context.Context
	cfg    *config.Config
	logger *logrus.Logger
}

func NewOtpServer(ctx context.Context, cfg *config.Config, logger *logrus.Logger) *otpServer {
	e := echo.New()
	srv := &otpServer{
		e,
		ctx,
		cfg,
		logger,
	}

	logger.Info("try to run api")

	return srv
}

func (srv *otpServer) StartServer() {
	if err := srv.Echo.Start(srv.cfg.API.Address); err != nil {
		srv.logger.Fatal(err)
	}
}
