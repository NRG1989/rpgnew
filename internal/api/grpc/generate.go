package grpc

import (
	"fmt"

	otp "github.com/hgfischer/go-otp"
)

const (
	numOfKeys = 10
	lenOfKey  = 6
)

func generateOTPsForUser() string {

	var userOtpKeys [numOfKeys]string
	for i := 0; i < numOfKeys; i++ {
		totp := &otp.TOTP{
			Length: uint8(lenOfKey),
		}
		otp := totp.Get()
		userOtpKeys[i] = otp
	}
	return fmt.Sprintf("{%v}", userOtpKeys)
}
