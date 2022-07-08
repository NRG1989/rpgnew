package grpc

import (
	"context"

	"testing"

	mock "go-aut-registration-user-otp/internal/api/grpc/mocks"

	pbOTP "go-aut-registration-user-otp/proto"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestSendOTPKeys(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock.NewMockStorage(ctrl)

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	handler := NewAuthOTPService(store, logger)

	store.EXPECT().GetExternalID(gomock.Any(), gomock.Any(), gomock.Any())
	store.EXPECT().AddOTPClient(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

	req := &pbOTP.SendOTPCodeRequest{Phone: "+700000000000"}

	response, err := handler.SendOTPCode(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, response)

}

func TestPassRecoceryByOTP(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock.NewMockStorage(ctrl)
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	handler := NewAuthOTPService(store, logger)

	store.EXPECT().FindIDByPhone(gomock.Any(), gomock.Any(), gomock.Any())
	store.EXPECT().UpdateOTP(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	req := &pbOTP.PassRecoceryByOTPRequest{Otp: "123031",
		Phone: "+71000333000"}

	response, err := handler.PassRecoceryByOTP(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, response)

}
