// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: booking.proto

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

// BookingServiceRPCClient is the client API for BookingServiceRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceRPCClient interface {
	GetBookingByID(ctx context.Context, in *GetBookingByIdRequest, opts ...grpc.CallOption) (*Booking, error)
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	UpdateBookingStatus(ctx context.Context, in *UpdateBookingStatusRequest, opts ...grpc.CallOption) (*UpdateBookingStatusResponse, error)
	DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*DeleteBookingResponse, error)
	CustomerBookingHistory(ctx context.Context, in *CustomerBookingHistoryRequest, opts ...grpc.CallOption) (*Bookings, error)
}

type bookingServiceRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceRPCClient(cc grpc.ClientConnInterface) BookingServiceRPCClient {
	return &bookingServiceRPCClient{cc}
}

func (c *bookingServiceRPCClient) GetBookingByID(ctx context.Context, in *GetBookingByIdRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.BookingServiceRPC/GetBookingByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceRPCClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.BookingServiceRPC/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceRPCClient) UpdateBookingStatus(ctx context.Context, in *UpdateBookingStatusRequest, opts ...grpc.CallOption) (*UpdateBookingStatusResponse, error) {
	out := new(UpdateBookingStatusResponse)
	err := c.cc.Invoke(ctx, "/training.BookingServiceRPC/UpdateBookingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceRPCClient) DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*DeleteBookingResponse, error) {
	out := new(DeleteBookingResponse)
	err := c.cc.Invoke(ctx, "/training.BookingServiceRPC/DeleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceRPCClient) CustomerBookingHistory(ctx context.Context, in *CustomerBookingHistoryRequest, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/training.BookingServiceRPC/CustomerBookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceRPCServer is the server API for BookingServiceRPC service.
// All implementations must embed UnimplementedBookingServiceRPCServer
// for forward compatibility
type BookingServiceRPCServer interface {
	GetBookingByID(context.Context, *GetBookingByIdRequest) (*Booking, error)
	CreateBooking(context.Context, *CreateBookingRequest) (*Booking, error)
	UpdateBookingStatus(context.Context, *UpdateBookingStatusRequest) (*UpdateBookingStatusResponse, error)
	DeleteBooking(context.Context, *DeleteBookingRequest) (*DeleteBookingResponse, error)
	CustomerBookingHistory(context.Context, *CustomerBookingHistoryRequest) (*Bookings, error)
	mustEmbedUnimplementedBookingServiceRPCServer()
}

// UnimplementedBookingServiceRPCServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceRPCServer struct {
}

func (UnimplementedBookingServiceRPCServer) GetBookingByID(context.Context, *GetBookingByIdRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingByID not implemented")
}
func (UnimplementedBookingServiceRPCServer) CreateBooking(context.Context, *CreateBookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceRPCServer) UpdateBookingStatus(context.Context, *UpdateBookingStatusRequest) (*UpdateBookingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookingStatus not implemented")
}
func (UnimplementedBookingServiceRPCServer) DeleteBooking(context.Context, *DeleteBookingRequest) (*DeleteBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooking not implemented")
}
func (UnimplementedBookingServiceRPCServer) CustomerBookingHistory(context.Context, *CustomerBookingHistoryRequest) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerBookingHistory not implemented")
}
func (UnimplementedBookingServiceRPCServer) mustEmbedUnimplementedBookingServiceRPCServer() {}

// UnsafeBookingServiceRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceRPCServer will
// result in compilation errors.
type UnsafeBookingServiceRPCServer interface {
	mustEmbedUnimplementedBookingServiceRPCServer()
}

func RegisterBookingServiceRPCServer(s grpc.ServiceRegistrar, srv BookingServiceRPCServer) {
	s.RegisterService(&BookingServiceRPC_ServiceDesc, srv)
}

func _BookingServiceRPC_GetBookingByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceRPCServer).GetBookingByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.BookingServiceRPC/GetBookingByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceRPCServer).GetBookingByID(ctx, req.(*GetBookingByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingServiceRPC_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceRPCServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.BookingServiceRPC/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceRPCServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingServiceRPC_UpdateBookingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceRPCServer).UpdateBookingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.BookingServiceRPC/UpdateBookingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceRPCServer).UpdateBookingStatus(ctx, req.(*UpdateBookingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingServiceRPC_DeleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceRPCServer).DeleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.BookingServiceRPC/DeleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceRPCServer).DeleteBooking(ctx, req.(*DeleteBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingServiceRPC_CustomerBookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerBookingHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceRPCServer).CustomerBookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.BookingServiceRPC/CustomerBookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceRPCServer).CustomerBookingHistory(ctx, req.(*CustomerBookingHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingServiceRPC_ServiceDesc is the grpc.ServiceDesc for BookingServiceRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingServiceRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.BookingServiceRPC",
	HandlerType: (*BookingServiceRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookingByID",
			Handler:    _BookingServiceRPC_GetBookingByID_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _BookingServiceRPC_CreateBooking_Handler,
		},
		{
			MethodName: "UpdateBookingStatus",
			Handler:    _BookingServiceRPC_UpdateBookingStatus_Handler,
		},
		{
			MethodName: "DeleteBooking",
			Handler:    _BookingServiceRPC_DeleteBooking_Handler,
		},
		{
			MethodName: "CustomerBookingHistory",
			Handler:    _BookingServiceRPC_CustomerBookingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
