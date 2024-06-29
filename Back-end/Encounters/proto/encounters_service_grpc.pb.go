// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc3
// source: encounters_service.proto

package encounters

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

const (
	EncounterService_CreateEncounterRpc_FullMethodName  = "/EncounterService/CreateEncounterRpc"
	EncounterService_GetAllEncountersRpc_FullMethodName = "/EncounterService/GetAllEncountersRpc"
	EncounterService_GetEncounterByIDRpc_FullMethodName = "/EncounterService/GetEncounterByIDRpc"
)

// EncounterServiceClient is the client API for EncounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterServiceClient interface {
	CreateEncounterRpc(ctx context.Context, in *CreateEncounterRequest, opts ...grpc.CallOption) (*CreateEncounterResponse, error)
	GetAllEncountersRpc(ctx context.Context, in *GetAllEncountersRequest, opts ...grpc.CallOption) (*GetAllEncountersResponse, error)
	GetEncounterByIDRpc(ctx context.Context, in *GetEncounterByIDRequest, opts ...grpc.CallOption) (*GetEncounterByIDResponse, error)
}

type encounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterServiceClient(cc grpc.ClientConnInterface) EncounterServiceClient {
	return &encounterServiceClient{cc}
}

func (c *encounterServiceClient) CreateEncounterRpc(ctx context.Context, in *CreateEncounterRequest, opts ...grpc.CallOption) (*CreateEncounterResponse, error) {
	out := new(CreateEncounterResponse)
	err := c.cc.Invoke(ctx, EncounterService_CreateEncounterRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) GetAllEncountersRpc(ctx context.Context, in *GetAllEncountersRequest, opts ...grpc.CallOption) (*GetAllEncountersResponse, error) {
	out := new(GetAllEncountersResponse)
	err := c.cc.Invoke(ctx, EncounterService_GetAllEncountersRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) GetEncounterByIDRpc(ctx context.Context, in *GetEncounterByIDRequest, opts ...grpc.CallOption) (*GetEncounterByIDResponse, error) {
	out := new(GetEncounterByIDResponse)
	err := c.cc.Invoke(ctx, EncounterService_GetEncounterByIDRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterServiceServer is the server API for EncounterService service.
// All implementations must embed UnimplementedEncounterServiceServer
// for forward compatibility
type EncounterServiceServer interface {
	CreateEncounterRpc(context.Context, *CreateEncounterRequest) (*CreateEncounterResponse, error)
	GetAllEncountersRpc(context.Context, *GetAllEncountersRequest) (*GetAllEncountersResponse, error)
	GetEncounterByIDRpc(context.Context, *GetEncounterByIDRequest) (*GetEncounterByIDResponse, error)
	mustEmbedUnimplementedEncounterServiceServer()
}

// UnimplementedEncounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (UnimplementedEncounterServiceServer) CreateEncounterRpc(context.Context, *CreateEncounterRequest) (*CreateEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEncounterRpc not implemented")
}
func (UnimplementedEncounterServiceServer) GetAllEncountersRpc(context.Context, *GetAllEncountersRequest) (*GetAllEncountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEncountersRpc not implemented")
}
func (UnimplementedEncounterServiceServer) GetEncounterByIDRpc(context.Context, *GetEncounterByIDRequest) (*GetEncounterByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncounterByIDRpc not implemented")
}
func (UnimplementedEncounterServiceServer) mustEmbedUnimplementedEncounterServiceServer() {}

// UnsafeEncounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServiceServer will
// result in compilation errors.
type UnsafeEncounterServiceServer interface {
	mustEmbedUnimplementedEncounterServiceServer()
}

func RegisterEncounterServiceServer(s grpc.ServiceRegistrar, srv EncounterServiceServer) {
	s.RegisterService(&EncounterService_ServiceDesc, srv)
}

func _EncounterService_CreateEncounterRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).CreateEncounterRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_CreateEncounterRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).CreateEncounterRpc(ctx, req.(*CreateEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_GetAllEncountersRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEncountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).GetAllEncountersRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_GetAllEncountersRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).GetAllEncountersRpc(ctx, req.(*GetAllEncountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_GetEncounterByIDRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEncounterByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).GetEncounterByIDRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterService_GetEncounterByIDRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).GetEncounterByIDRpc(ctx, req.(*GetEncounterByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncounterService_ServiceDesc is the grpc.ServiceDesc for EncounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EncounterService",
	HandlerType: (*EncounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEncounterRpc",
			Handler:    _EncounterService_CreateEncounterRpc_Handler,
		},
		{
			MethodName: "GetAllEncountersRpc",
			Handler:    _EncounterService_GetAllEncountersRpc_Handler,
		},
		{
			MethodName: "GetEncounterByIDRpc",
			Handler:    _EncounterService_GetEncounterByIDRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encounters_service.proto",
}

const (
	EncounterExecutionService_CreateExecutionRpc_FullMethodName      = "/EncounterExecutionService/CreateExecutionRpc"
	EncounterExecutionService_GetAllExecutionsRpc_FullMethodName     = "/EncounterExecutionService/GetAllExecutionsRpc"
	EncounterExecutionService_GetExecutionByUserIDRpc_FullMethodName = "/EncounterExecutionService/GetExecutionByUserIDRpc"
	EncounterExecutionService_UpdateExecutionRpc_FullMethodName      = "/EncounterExecutionService/UpdateExecutionRpc"
)

// EncounterExecutionServiceClient is the client API for EncounterExecutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterExecutionServiceClient interface {
	CreateExecutionRpc(ctx context.Context, in *CreateExecutionRequest, opts ...grpc.CallOption) (*CreateExecutionResponse, error)
	GetAllExecutionsRpc(ctx context.Context, in *GetAllExecutionsRequest, opts ...grpc.CallOption) (*GetAllExecutionsResponse, error)
	GetExecutionByUserIDRpc(ctx context.Context, in *GetExecutionByUserIdRequest, opts ...grpc.CallOption) (*GetExecutionByUserIdResponse, error)
	UpdateExecutionRpc(ctx context.Context, in *UpdateExecutionRequest, opts ...grpc.CallOption) (*UpdateExecutionResponse, error)
}

type encounterExecutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterExecutionServiceClient(cc grpc.ClientConnInterface) EncounterExecutionServiceClient {
	return &encounterExecutionServiceClient{cc}
}

func (c *encounterExecutionServiceClient) CreateExecutionRpc(ctx context.Context, in *CreateExecutionRequest, opts ...grpc.CallOption) (*CreateExecutionResponse, error) {
	out := new(CreateExecutionResponse)
	err := c.cc.Invoke(ctx, EncounterExecutionService_CreateExecutionRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterExecutionServiceClient) GetAllExecutionsRpc(ctx context.Context, in *GetAllExecutionsRequest, opts ...grpc.CallOption) (*GetAllExecutionsResponse, error) {
	out := new(GetAllExecutionsResponse)
	err := c.cc.Invoke(ctx, EncounterExecutionService_GetAllExecutionsRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterExecutionServiceClient) GetExecutionByUserIDRpc(ctx context.Context, in *GetExecutionByUserIdRequest, opts ...grpc.CallOption) (*GetExecutionByUserIdResponse, error) {
	out := new(GetExecutionByUserIdResponse)
	err := c.cc.Invoke(ctx, EncounterExecutionService_GetExecutionByUserIDRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterExecutionServiceClient) UpdateExecutionRpc(ctx context.Context, in *UpdateExecutionRequest, opts ...grpc.CallOption) (*UpdateExecutionResponse, error) {
	out := new(UpdateExecutionResponse)
	err := c.cc.Invoke(ctx, EncounterExecutionService_UpdateExecutionRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterExecutionServiceServer is the server API for EncounterExecutionService service.
// All implementations must embed UnimplementedEncounterExecutionServiceServer
// for forward compatibility
type EncounterExecutionServiceServer interface {
	CreateExecutionRpc(context.Context, *CreateExecutionRequest) (*CreateExecutionResponse, error)
	GetAllExecutionsRpc(context.Context, *GetAllExecutionsRequest) (*GetAllExecutionsResponse, error)
	GetExecutionByUserIDRpc(context.Context, *GetExecutionByUserIdRequest) (*GetExecutionByUserIdResponse, error)
	UpdateExecutionRpc(context.Context, *UpdateExecutionRequest) (*UpdateExecutionResponse, error)
	mustEmbedUnimplementedEncounterExecutionServiceServer()
}

// UnimplementedEncounterExecutionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterExecutionServiceServer struct {
}

func (UnimplementedEncounterExecutionServiceServer) CreateExecutionRpc(context.Context, *CreateExecutionRequest) (*CreateExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExecutionRpc not implemented")
}
func (UnimplementedEncounterExecutionServiceServer) GetAllExecutionsRpc(context.Context, *GetAllExecutionsRequest) (*GetAllExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllExecutionsRpc not implemented")
}
func (UnimplementedEncounterExecutionServiceServer) GetExecutionByUserIDRpc(context.Context, *GetExecutionByUserIdRequest) (*GetExecutionByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecutionByUserIDRpc not implemented")
}
func (UnimplementedEncounterExecutionServiceServer) UpdateExecutionRpc(context.Context, *UpdateExecutionRequest) (*UpdateExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExecutionRpc not implemented")
}
func (UnimplementedEncounterExecutionServiceServer) mustEmbedUnimplementedEncounterExecutionServiceServer() {
}

// UnsafeEncounterExecutionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterExecutionServiceServer will
// result in compilation errors.
type UnsafeEncounterExecutionServiceServer interface {
	mustEmbedUnimplementedEncounterExecutionServiceServer()
}

func RegisterEncounterExecutionServiceServer(s grpc.ServiceRegistrar, srv EncounterExecutionServiceServer) {
	s.RegisterService(&EncounterExecutionService_ServiceDesc, srv)
}

func _EncounterExecutionService_CreateExecutionRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterExecutionServiceServer).CreateExecutionRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterExecutionService_CreateExecutionRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterExecutionServiceServer).CreateExecutionRpc(ctx, req.(*CreateExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterExecutionService_GetAllExecutionsRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterExecutionServiceServer).GetAllExecutionsRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterExecutionService_GetAllExecutionsRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterExecutionServiceServer).GetAllExecutionsRpc(ctx, req.(*GetAllExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterExecutionService_GetExecutionByUserIDRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExecutionByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterExecutionServiceServer).GetExecutionByUserIDRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterExecutionService_GetExecutionByUserIDRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterExecutionServiceServer).GetExecutionByUserIDRpc(ctx, req.(*GetExecutionByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterExecutionService_UpdateExecutionRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterExecutionServiceServer).UpdateExecutionRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncounterExecutionService_UpdateExecutionRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterExecutionServiceServer).UpdateExecutionRpc(ctx, req.(*UpdateExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncounterExecutionService_ServiceDesc is the grpc.ServiceDesc for EncounterExecutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncounterExecutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EncounterExecutionService",
	HandlerType: (*EncounterExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExecutionRpc",
			Handler:    _EncounterExecutionService_CreateExecutionRpc_Handler,
		},
		{
			MethodName: "GetAllExecutionsRpc",
			Handler:    _EncounterExecutionService_GetAllExecutionsRpc_Handler,
		},
		{
			MethodName: "GetExecutionByUserIDRpc",
			Handler:    _EncounterExecutionService_GetExecutionByUserIDRpc_Handler,
		},
		{
			MethodName: "UpdateExecutionRpc",
			Handler:    _EncounterExecutionService_UpdateExecutionRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encounters_service.proto",
}