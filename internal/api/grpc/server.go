package grpc

import (
	"net"

	pbOTP "go-aut-registration-user-otp/proto"

	"go-aut-registration-user-otp/internal/database"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	logger      *logrus.Logger
	GRPCAddress string
	db          database.Storage
}

func NewGRPCService(
	logger *logrus.Logger,
	GRPCAddress string,
	db database.Storage,
) *Server {
	return &Server{
		logger,
		GRPCAddress,
		db,
	}
}
func (s Server) Start() error {
	conn, err := net.Listen("tcp", s.GRPCAddress)
	if err != nil {
		s.logger.WithError(err).Errorln("no connection to grpc server")
		return err
	}
	grpcServer := grpc.NewServer()
	pbOTP.RegisterGoAuthOTPServer(grpcServer, &AuthOTPService{
		Logger: s.logger,
		store:  s.db,
	})

	if err := grpcServer.Serve(conn); err != nil {
		s.logger.WithError(err).Errorln("connection to grpc error")
		return err
	}
	return nil
}
