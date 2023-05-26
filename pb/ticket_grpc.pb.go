// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: ticket.proto

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

// TicketServiceRPCClient is the client API for TicketServiceRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceRPCClient interface {
	GetTicketById(ctx context.Context, in *GetTicketByIdRequest, opts ...grpc.CallOption) (*Ticket, error)
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	UpdateTicketById(ctx context.Context, in *UpdateTicketByIdRequest, opts ...grpc.CallOption) (*Ticket, error)
	UpdateTicketStatusByFlightId(ctx context.Context, in *UpdateTicketStatusByFlightIdRequest, opts ...grpc.CallOption) (*UpdateTicketStatusByFlightIdResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
}

type ticketServiceRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceRPCClient(cc grpc.ClientConnInterface) TicketServiceRPCClient {
	return &ticketServiceRPCClient{cc}
}

func (c *ticketServiceRPCClient) GetTicketById(ctx context.Context, in *GetTicketByIdRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/training.TicketServiceRPC/GetTicketById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceRPCClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/training.TicketServiceRPC/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceRPCClient) UpdateTicketById(ctx context.Context, in *UpdateTicketByIdRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/training.TicketServiceRPC/UpdateTicketById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceRPCClient) UpdateTicketStatusByFlightId(ctx context.Context, in *UpdateTicketStatusByFlightIdRequest, opts ...grpc.CallOption) (*UpdateTicketStatusByFlightIdResponse, error) {
	out := new(UpdateTicketStatusByFlightIdResponse)
	err := c.cc.Invoke(ctx, "/training.TicketServiceRPC/UpdateTicketStatusByFlightId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceRPCClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, "/training.TicketServiceRPC/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceRPCServer is the server API for TicketServiceRPC service.
// All implementations must embed UnimplementedTicketServiceRPCServer
// for forward compatibility
type TicketServiceRPCServer interface {
	GetTicketById(context.Context, *GetTicketByIdRequest) (*Ticket, error)
	CreateTicket(context.Context, *CreateTicketRequest) (*Ticket, error)
	UpdateTicketById(context.Context, *UpdateTicketByIdRequest) (*Ticket, error)
	UpdateTicketStatusByFlightId(context.Context, *UpdateTicketStatusByFlightIdRequest) (*UpdateTicketStatusByFlightIdResponse, error)
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	mustEmbedUnimplementedTicketServiceRPCServer()
}

// UnimplementedTicketServiceRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceRPCServer struct {
}

func (UnimplementedTicketServiceRPCServer) GetTicketById(context.Context, *GetTicketByIdRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicketById not implemented")
}
func (UnimplementedTicketServiceRPCServer) CreateTicket(context.Context, *CreateTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketServiceRPCServer) UpdateTicketById(context.Context, *UpdateTicketByIdRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicketById not implemented")
}
func (UnimplementedTicketServiceRPCServer) UpdateTicketStatusByFlightId(context.Context, *UpdateTicketStatusByFlightIdRequest) (*UpdateTicketStatusByFlightIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicketStatusByFlightId not implemented")
}
func (UnimplementedTicketServiceRPCServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTicketServiceRPCServer) mustEmbedUnimplementedTicketServiceRPCServer() {}

// UnsafeTicketServiceRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceRPCServer will
// result in compilation errors.
type UnsafeTicketServiceRPCServer interface {
	mustEmbedUnimplementedTicketServiceRPCServer()
}

func RegisterTicketServiceRPCServer(s grpc.ServiceRegistrar, srv TicketServiceRPCServer) {
	s.RegisterService(&TicketServiceRPC_ServiceDesc, srv)
}

func _TicketServiceRPC_GetTicketById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceRPCServer).GetTicketById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.TicketServiceRPC/GetTicketById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceRPCServer).GetTicketById(ctx, req.(*GetTicketByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketServiceRPC_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceRPCServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.TicketServiceRPC/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceRPCServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketServiceRPC_UpdateTicketById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceRPCServer).UpdateTicketById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.TicketServiceRPC/UpdateTicketById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceRPCServer).UpdateTicketById(ctx, req.(*UpdateTicketByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketServiceRPC_UpdateTicketStatusByFlightId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketStatusByFlightIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceRPCServer).UpdateTicketStatusByFlightId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.TicketServiceRPC/UpdateTicketStatusByFlightId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceRPCServer).UpdateTicketStatusByFlightId(ctx, req.(*UpdateTicketStatusByFlightIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketServiceRPC_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceRPCServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.TicketServiceRPC/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceRPCServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketServiceRPC_ServiceDesc is the grpc.ServiceDesc for TicketServiceRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketServiceRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.TicketServiceRPC",
	HandlerType: (*TicketServiceRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTicketById",
			Handler:    _TicketServiceRPC_GetTicketById_Handler,
		},
		{
			MethodName: "CreateTicket",
			Handler:    _TicketServiceRPC_CreateTicket_Handler,
		},
		{
			MethodName: "UpdateTicketById",
			Handler:    _TicketServiceRPC_UpdateTicketById_Handler,
		},
		{
			MethodName: "UpdateTicketStatusByFlightId",
			Handler:    _TicketServiceRPC_UpdateTicketStatusByFlightId_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TicketServiceRPC_DeleteTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}
