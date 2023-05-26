// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: customer.proto

package pb

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

// CustomerServiceRPCClient is the client API for CustomerServiceRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceRPCClient interface {
	GetCustomerByCitizenId(ctx context.Context, in *GetCustomerByCitizenIdRequest, opts ...grpc.CallOption) (*Customer, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error)
}

type customerServiceRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceRPCClient(cc grpc.ClientConnInterface) CustomerServiceRPCClient {
	return &customerServiceRPCClient{cc}
}

func (c *customerServiceRPCClient) GetCustomerByCitizenId(ctx context.Context, in *GetCustomerByCitizenIdRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/training.CustomerServiceRPC/GetCustomerByCitizenId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceRPCClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/training.CustomerServiceRPC/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceRPCClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/training.CustomerServiceRPC/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceRPCClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error) {
	out := new(DeleteCustomerResponse)
	err := c.cc.Invoke(ctx, "/training.CustomerServiceRPC/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceRPCServer is the server API for CustomerServiceRPC service.
// All implementations must embed UnimplementedCustomerServiceRPCServer
// for forward compatibility
type CustomerServiceRPCServer interface {
	GetCustomerByCitizenId(context.Context, *GetCustomerByCitizenIdRequest) (*Customer, error)
	CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*Customer, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error)
	mustEmbedUnimplementedCustomerServiceRPCServer()
}

// UnimplementedCustomerServiceRPCServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceRPCServer struct {
}

func (UnimplementedCustomerServiceRPCServer) GetCustomerByCitizenId(context.Context, *GetCustomerByCitizenIdRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByCitizenId not implemented")
}
func (UnimplementedCustomerServiceRPCServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceRPCServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServiceRPCServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServiceRPCServer) mustEmbedUnimplementedCustomerServiceRPCServer() {}

// UnsafeCustomerServiceRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceRPCServer will
// result in compilation errors.
type UnsafeCustomerServiceRPCServer interface {
	mustEmbedUnimplementedCustomerServiceRPCServer()
}

func RegisterCustomerServiceRPCServer(s grpc.ServiceRegistrar, srv CustomerServiceRPCServer) {
	s.RegisterService(&CustomerServiceRPC_ServiceDesc, srv)
}

func _CustomerServiceRPC_GetCustomerByCitizenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerByCitizenIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceRPCServer).GetCustomerByCitizenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.CustomerServiceRPC/GetCustomerByCitizenId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceRPCServer).GetCustomerByCitizenId(ctx, req.(*GetCustomerByCitizenIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServiceRPC_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceRPCServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.CustomerServiceRPC/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceRPCServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServiceRPC_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceRPCServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.CustomerServiceRPC/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceRPCServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerServiceRPC_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceRPCServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.CustomerServiceRPC/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceRPCServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerServiceRPC_ServiceDesc is the grpc.ServiceDesc for CustomerServiceRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerServiceRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.CustomerServiceRPC",
	HandlerType: (*CustomerServiceRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerByCitizenId",
			Handler:    _CustomerServiceRPC_GetCustomerByCitizenId_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerServiceRPC_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerServiceRPC_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerServiceRPC_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
