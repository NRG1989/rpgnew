// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __otp_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoAuthOTPClient is the client API for GoAuthOTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoAuthOTPClient interface {
	SendOTPCode(ctx context.Context, in *SendOTPCodeRequest, opts ...grpc.CallOption) (*SendOTPCodeResponse, error)
	PassRecoceryByOTP(ctx context.Context, in *PassRecoceryByOTPRequest, opts ...grpc.CallOption) (*PassRecoceryByOTPResponse, error)
}

type goAuthOTPClient struct {
	cc grpc.ClientConnInterface
}

func NewGoAuthOTPClient(cc grpc.ClientConnInterface) GoAuthOTPClient {
	return &goAuthOTPClient{cc}
}

func (c *goAuthOTPClient) SendOTPCode(ctx context.Context, in *SendOTPCodeRequest, opts ...grpc.CallOption) (*SendOTPCodeResponse, error) {
	out := new(SendOTPCodeResponse)
	err := c.cc.Invoke(ctx, "/otp_server.GoAuthOTP/SendOTPCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAuthOTPClient) PassRecoceryByOTP(ctx context.Context, in *PassRecoceryByOTPRequest, opts ...grpc.CallOption) (*PassRecoceryByOTPResponse, error) {
	out := new(PassRecoceryByOTPResponse)
	err := c.cc.Invoke(ctx, "/otp_server.GoAuthOTP/PassRecoceryByOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoAuthOTPServer is the server API for GoAuthOTP service.
// All implementations must embed UnimplementedGoAuthOTPServer
// for forward compatibility
type GoAuthOTPServer interface {
	SendOTPCode(context.Context, *SendOTPCodeRequest) (*SendOTPCodeResponse, error)
	PassRecoceryByOTP(context.Context, *PassRecoceryByOTPRequest) (*PassRecoceryByOTPResponse, error)
	mustEmbedUnimplementedGoAuthOTPServer()
}

// UnimplementedGoAuthOTPServer must be embedded to have forward compatible implementations.
type UnimplementedGoAuthOTPServer struct {
}

func (UnimplementedGoAuthOTPServer) SendOTPCode(context.Context, *SendOTPCodeRequest) (*SendOTPCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTPCode not implemented")
}
func (UnimplementedGoAuthOTPServer) PassRecoceryByOTP(context.Context, *PassRecoceryByOTPRequest) (*PassRecoceryByOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassRecoceryByOTP not implemented")
}
func (UnimplementedGoAuthOTPServer) mustEmbedUnimplementedGoAuthOTPServer() {}

// UnsafeGoAuthOTPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoAuthOTPServer will
// result in compilation errors.
type UnsafeGoAuthOTPServer interface {
	mustEmbedUnimplementedGoAuthOTPServer()
}

func RegisterGoAuthOTPServer(s grpc.ServiceRegistrar, srv GoAuthOTPServer) {
	s.RegisterService(&GoAuthOTP_ServiceDesc, srv)
}

func _GoAuthOTP_SendOTPCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOTPCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAuthOTPServer).SendOTPCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otp_server.GoAuthOTP/SendOTPCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAuthOTPServer).SendOTPCode(ctx, req.(*SendOTPCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAuthOTP_PassRecoceryByOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassRecoceryByOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAuthOTPServer).PassRecoceryByOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otp_server.GoAuthOTP/PassRecoceryByOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAuthOTPServer).PassRecoceryByOTP(ctx, req.(*PassRecoceryByOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoAuthOTP_ServiceDesc is the grpc.ServiceDesc for GoAuthOTP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoAuthOTP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otp_server.GoAuthOTP",
	HandlerType: (*GoAuthOTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTPCode",
			Handler:    _GoAuthOTP_SendOTPCode_Handler,
		},
		{
			MethodName: "PassRecoceryByOTP",
			Handler:    _GoAuthOTP_PassRecoceryByOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/otp.proto",
}
