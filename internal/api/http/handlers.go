package http

import (
	"go-aut-registration-user-otp/internal/database"

	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Logger *logrus.Logger
	store  database.Storage
}

func NewHandler(store database.Storage, logger *logrus.Logger) *Handlers {
	return &Handlers{
		store:  store,
		Logger: logger,
	}
}
