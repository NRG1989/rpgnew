

genmock:
	mockgen -source=./internal/database/interface.go -destination=./internal/api/grpc/mocks/mock_interface.go

otp: ##To update files to work with OTP RPG-New API
    protoc --go_out=./proto ./proto/otp.proto
    protoc --go-grpc_out=./proto ./proto/otp.proto