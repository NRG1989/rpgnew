package database

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Storage interface {
	AddOTPClient(ctx context.Context, logger *logrus.Logger, extID string, phone string) error
	GetExternalID(ctx context.Context, logger *logrus.Logger, otpKeys string) (string, error)
	FindIDByPhone(ctx context.Context, logger *logrus.Logger, phone string) (string, error)
	UpdateOTP(ctx context.Context, logger *logrus.Logger, otpKey, extID string) error
}
