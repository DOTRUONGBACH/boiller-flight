// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: flight.proto

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

// FlightServiceRPCClient is the client API for FlightServiceRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceRPCClient interface {
	GetFlightById(ctx context.Context, in *GetFlightByIdRequest, opts ...grpc.CallOption) (*Flight, error)
	GetFlightByFlightCode(ctx context.Context, in *GetFlightByFlightCodeRequest, opts ...grpc.CallOption) (*Flight, error)
	CreateFlight(ctx context.Context, in *CreateFlightRequest, opts ...grpc.CallOption) (*Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, in *UpdateFlightByFlightCodeRequest, opts ...grpc.CallOption) (*Flight, error)
	UpdateAvailableSeat(ctx context.Context, in *UpdateFlightAvailableSeatRequest, opts ...grpc.CallOption) (*Flight, error)
	UpdateFlightStatus(ctx context.Context, in *UpdateFlightStatusRequest, opts ...grpc.CallOption) (*Flight, error)
	DeleteFlight(ctx context.Context, in *DeleteFlightRequest, opts ...grpc.CallOption) (*DeleteFlightResponse, error)
	FindFlightByDateRange(ctx context.Context, in *FindFlightByDateRangeRequest, opts ...grpc.CallOption) (*Flights, error)
}

type flightServiceRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceRPCClient(cc grpc.ClientConnInterface) FlightServiceRPCClient {
	return &flightServiceRPCClient{cc}
}

func (c *flightServiceRPCClient) GetFlightById(ctx context.Context, in *GetFlightByIdRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/GetFlightById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) GetFlightByFlightCode(ctx context.Context, in *GetFlightByFlightCodeRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/GetFlightByFlightCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) CreateFlight(ctx context.Context, in *CreateFlightRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/CreateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) UpdateFlightByFlightCode(ctx context.Context, in *UpdateFlightByFlightCodeRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/UpdateFlightByFlightCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) UpdateAvailableSeat(ctx context.Context, in *UpdateFlightAvailableSeatRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/UpdateAvailableSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) UpdateFlightStatus(ctx context.Context, in *UpdateFlightStatusRequest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/UpdateFlightStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) DeleteFlight(ctx context.Context, in *DeleteFlightRequest, opts ...grpc.CallOption) (*DeleteFlightResponse, error) {
	out := new(DeleteFlightResponse)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/DeleteFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceRPCClient) FindFlightByDateRange(ctx context.Context, in *FindFlightByDateRangeRequest, opts ...grpc.CallOption) (*Flights, error) {
	out := new(Flights)
	err := c.cc.Invoke(ctx, "/training.FlightServiceRPC/FindFlightByDateRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceRPCServer is the server API for FlightServiceRPC service.
// All implementations must embed UnimplementedFlightServiceRPCServer
// for forward compatibility
type FlightServiceRPCServer interface {
	GetFlightById(context.Context, *GetFlightByIdRequest) (*Flight, error)
	GetFlightByFlightCode(context.Context, *GetFlightByFlightCodeRequest) (*Flight, error)
	CreateFlight(context.Context, *CreateFlightRequest) (*Flight, error)
	UpdateFlightByFlightCode(context.Context, *UpdateFlightByFlightCodeRequest) (*Flight, error)
	UpdateAvailableSeat(context.Context, *UpdateFlightAvailableSeatRequest) (*Flight, error)
	UpdateFlightStatus(context.Context, *UpdateFlightStatusRequest) (*Flight, error)
	DeleteFlight(context.Context, *DeleteFlightRequest) (*DeleteFlightResponse, error)
	FindFlightByDateRange(context.Context, *FindFlightByDateRangeRequest) (*Flights, error)
	mustEmbedUnimplementedFlightServiceRPCServer()
}

// UnimplementedFlightServiceRPCServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceRPCServer struct {
}

func (UnimplementedFlightServiceRPCServer) GetFlightById(context.Context, *GetFlightByIdRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightById not implemented")
}
func (UnimplementedFlightServiceRPCServer) GetFlightByFlightCode(context.Context, *GetFlightByFlightCodeRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightByFlightCode not implemented")
}
func (UnimplementedFlightServiceRPCServer) CreateFlight(context.Context, *CreateFlightRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServiceRPCServer) UpdateFlightByFlightCode(context.Context, *UpdateFlightByFlightCodeRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlightByFlightCode not implemented")
}
func (UnimplementedFlightServiceRPCServer) UpdateAvailableSeat(context.Context, *UpdateFlightAvailableSeatRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailableSeat not implemented")
}
func (UnimplementedFlightServiceRPCServer) UpdateFlightStatus(context.Context, *UpdateFlightStatusRequest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlightStatus not implemented")
}
func (UnimplementedFlightServiceRPCServer) DeleteFlight(context.Context, *DeleteFlightRequest) (*DeleteFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlight not implemented")
}
func (UnimplementedFlightServiceRPCServer) FindFlightByDateRange(context.Context, *FindFlightByDateRangeRequest) (*Flights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFlightByDateRange not implemented")
}
func (UnimplementedFlightServiceRPCServer) mustEmbedUnimplementedFlightServiceRPCServer() {}

// UnsafeFlightServiceRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceRPCServer will
// result in compilation errors.
type UnsafeFlightServiceRPCServer interface {
	mustEmbedUnimplementedFlightServiceRPCServer()
}

func RegisterFlightServiceRPCServer(s grpc.ServiceRegistrar, srv FlightServiceRPCServer) {
	s.RegisterService(&FlightServiceRPC_ServiceDesc, srv)
}

func _FlightServiceRPC_GetFlightById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlightByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).GetFlightById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/GetFlightById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).GetFlightById(ctx, req.(*GetFlightByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_GetFlightByFlightCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlightByFlightCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).GetFlightByFlightCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/GetFlightByFlightCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).GetFlightByFlightCode(ctx, req.(*GetFlightByFlightCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/CreateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).CreateFlight(ctx, req.(*CreateFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_UpdateFlightByFlightCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlightByFlightCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).UpdateFlightByFlightCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/UpdateFlightByFlightCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).UpdateFlightByFlightCode(ctx, req.(*UpdateFlightByFlightCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_UpdateAvailableSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlightAvailableSeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).UpdateAvailableSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/UpdateAvailableSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).UpdateAvailableSeat(ctx, req.(*UpdateFlightAvailableSeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_UpdateFlightStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlightStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).UpdateFlightStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/UpdateFlightStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).UpdateFlightStatus(ctx, req.(*UpdateFlightStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_DeleteFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).DeleteFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/DeleteFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).DeleteFlight(ctx, req.(*DeleteFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightServiceRPC_FindFlightByDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFlightByDateRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceRPCServer).FindFlightByDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FlightServiceRPC/FindFlightByDateRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceRPCServer).FindFlightByDateRange(ctx, req.(*FindFlightByDateRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightServiceRPC_ServiceDesc is the grpc.ServiceDesc for FlightServiceRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightServiceRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.FlightServiceRPC",
	HandlerType: (*FlightServiceRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlightById",
			Handler:    _FlightServiceRPC_GetFlightById_Handler,
		},
		{
			MethodName: "GetFlightByFlightCode",
			Handler:    _FlightServiceRPC_GetFlightByFlightCode_Handler,
		},
		{
			MethodName: "CreateFlight",
			Handler:    _FlightServiceRPC_CreateFlight_Handler,
		},
		{
			MethodName: "UpdateFlightByFlightCode",
			Handler:    _FlightServiceRPC_UpdateFlightByFlightCode_Handler,
		},
		{
			MethodName: "UpdateAvailableSeat",
			Handler:    _FlightServiceRPC_UpdateAvailableSeat_Handler,
		},
		{
			MethodName: "UpdateFlightStatus",
			Handler:    _FlightServiceRPC_UpdateFlightStatus_Handler,
		},
		{
			MethodName: "DeleteFlight",
			Handler:    _FlightServiceRPC_DeleteFlight_Handler,
		},
		{
			MethodName: "FindFlightByDateRange",
			Handler:    _FlightServiceRPC_FindFlightByDateRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flight.proto",
}
