// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: resources/proto/deposit.proto

package resources

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

// DepositServiceClient is the client API for DepositService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepositServiceClient interface {
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	GetDeposit(ctx context.Context, in *GetDepositRequest, opts ...grpc.CallOption) (*GetDepositResponse, error)
}

type depositServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepositServiceClient(cc grpc.ClientConnInterface) DepositServiceClient {
	return &depositServiceClient{cc}
}

func (c *depositServiceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/resources.DepositService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depositServiceClient) GetDeposit(ctx context.Context, in *GetDepositRequest, opts ...grpc.CallOption) (*GetDepositResponse, error) {
	out := new(GetDepositResponse)
	err := c.cc.Invoke(ctx, "/resources.DepositService/GetDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepositServiceServer is the server API for DepositService service.
// All implementations must embed UnimplementedDepositServiceServer
// for forward compatibility
type DepositServiceServer interface {
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	GetDeposit(context.Context, *GetDepositRequest) (*GetDepositResponse, error)
	mustEmbedUnimplementedDepositServiceServer()
}

// UnimplementedDepositServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepositServiceServer struct {
}

func (UnimplementedDepositServiceServer) Deposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedDepositServiceServer) GetDeposit(context.Context, *GetDepositRequest) (*GetDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeposit not implemented")
}
func (UnimplementedDepositServiceServer) mustEmbedUnimplementedDepositServiceServer() {}

// UnsafeDepositServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepositServiceServer will
// result in compilation errors.
type UnsafeDepositServiceServer interface {
	mustEmbedUnimplementedDepositServiceServer()
}

func RegisterDepositServiceServer(s grpc.ServiceRegistrar, srv DepositServiceServer) {
	s.RegisterService(&DepositService_ServiceDesc, srv)
}

func _DepositService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.DepositService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositServiceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepositService_GetDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositServiceServer).GetDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.DepositService/GetDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositServiceServer).GetDeposit(ctx, req.(*GetDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepositService_ServiceDesc is the grpc.ServiceDesc for DepositService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepositService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resources.DepositService",
	HandlerType: (*DepositServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _DepositService_Deposit_Handler,
		},
		{
			MethodName: "GetDeposit",
			Handler:    _DepositService_GetDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources/proto/deposit.proto",
}