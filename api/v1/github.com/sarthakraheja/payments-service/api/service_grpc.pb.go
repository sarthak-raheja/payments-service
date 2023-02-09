// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/v1/service.proto

package api

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

// PaymentsServiceClient is the client API for PaymentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsServiceClient interface {
	// Sends a greeting
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
}

type paymentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsServiceClient(cc grpc.ClientConnInterface) PaymentsServiceClient {
	return &paymentsServiceClient{cc}
}

func (c *paymentsServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/api.PaymentsService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsServiceClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, "/api.PaymentsService/GetPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServiceServer is the server API for PaymentsService service.
// All implementations must embed UnimplementedPaymentsServiceServer
// for forward compatibility
type PaymentsServiceServer interface {
	// Sends a greeting
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	mustEmbedUnimplementedPaymentsServiceServer()
}

// UnimplementedPaymentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsServiceServer struct {
}

func (UnimplementedPaymentsServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentsServiceServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedPaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {}

// UnsafePaymentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsServiceServer will
// result in compilation errors.
type UnsafePaymentsServiceServer interface {
	mustEmbedUnimplementedPaymentsServiceServer()
}

func RegisterPaymentsServiceServer(s grpc.ServiceRegistrar, srv PaymentsServiceServer) {
	s.RegisterService(&PaymentsService_ServiceDesc, srv)
}

func _PaymentsService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PaymentsService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentsService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PaymentsService/GetPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentsService_ServiceDesc is the grpc.ServiceDesc for PaymentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PaymentsService",
	HandlerType: (*PaymentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentsService_CreatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _PaymentsService_GetPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/service.proto",
}
