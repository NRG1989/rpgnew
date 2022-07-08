package grpc

import (
	"context"

	"go-aut-registration-user-otp/internal/database"

	pbOTP "go-aut-registration-user-otp/proto"

	"github.com/sirupsen/logrus"
)

type AuthOTPService struct {
	Logger *logrus.Logger
	store  database.Storage
	pbOTP.UnimplementedGoAuthOTPServer
}

func NewAuthOTPService(store database.Storage, logger *logrus.Logger) *AuthOTPService {
	return &AuthOTPService{
		store:  store,
		Logger: logger,
	}
}

func (h *AuthOTPService) SendOTPCode(ctx context.Context, request *pbOTP.SendOTPCodeRequest) (*pbOTP.SendOTPCodeResponse, error) {

	otp_keys := generateOTPsForUser()

	externalID, err := h.store.GetExternalID(ctx, h.Logger, otp_keys)
	if err != nil {

		h.Logger.Printf("Client was not add: %s", err)
		return &pbOTP.SendOTPCodeResponse{}, err
	}

	if err = h.store.AddOTPClient(ctx, h.Logger, externalID, request.Phone); err != nil {
		if err != nil {
			return nil, err
		}

		h.Logger.Printf("ID was not add: %s", err)
		return &pbOTP.SendOTPCodeResponse{}, err
	}
	h.Logger.Info("ID added to DB")

	return &pbOTP.SendOTPCodeResponse{}, nil
}

func (h *AuthOTPService) PassRecoceryByOTP(ctx context.Context, req *pbOTP.PassRecoceryByOTPRequest) (*pbOTP.PassRecoceryByOTPResponse, error) {

	extID, err := h.store.FindIDByPhone(ctx, h.Logger, req.Phone)
	if err != nil {
		h.Logger.Error(err)
		return &pbOTP.PassRecoceryByOTPResponse{}, err
	}

	if err := h.store.UpdateOTP(ctx, h.Logger, req.Otp, extID); err != nil {
		h.Logger.Error(err)
		return &pbOTP.PassRecoceryByOTPResponse{}, err
	}

	h.Logger.Info("code was accepted successfully")

	return &pbOTP.PassRecoceryByOTPResponse{}, nil
}
