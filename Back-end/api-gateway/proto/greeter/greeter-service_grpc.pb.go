// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc3
// source: greeter/greeter-service.proto

package greeter

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
	GreeterService_Greet_FullMethodName     = "/GreeterService/Greet"
	GreeterService_GreetTest_FullMethodName = "/GreeterService/GreetTest"
)

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GreetTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GreeterService_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) GreetTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GreeterService_GreetTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	Greet(context.Context, *Request) (*Response, error)
	GreetTest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) Greet(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreeterServiceServer) GreetTest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetTest not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Greet(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_GreetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).GreetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_GreetTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).GreetTest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreeterService_Greet_Handler,
		},
		{
			MethodName: "GreetTest",
			Handler:    _GreeterService_GreetTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}

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
	Metadata: "greeter/greeter-service.proto",
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
	Metadata: "greeter/greeter-service.proto",
}

const (
	StakeholderService_RegistrationRpc_FullMethodName  = "/StakeholderService/RegistrationRpc"
	StakeholderService_GetProfileRpc_FullMethodName    = "/StakeholderService/GetProfileRpc"
	StakeholderService_UpdateProfileRpc_FullMethodName = "/StakeholderService/UpdateProfileRpc"
)

// StakeholderServiceClient is the client API for StakeholderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StakeholderServiceClient interface {
	RegistrationRpc(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	GetProfileRpc(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfileRpc(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
}

type stakeholderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStakeholderServiceClient(cc grpc.ClientConnInterface) StakeholderServiceClient {
	return &stakeholderServiceClient{cc}
}

func (c *stakeholderServiceClient) RegistrationRpc(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, StakeholderService_RegistrationRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakeholderServiceClient) GetProfileRpc(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, StakeholderService_GetProfileRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakeholderServiceClient) UpdateProfileRpc(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, StakeholderService_UpdateProfileRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakeholderServiceServer is the server API for StakeholderService service.
// All implementations must embed UnimplementedStakeholderServiceServer
// for forward compatibility
type StakeholderServiceServer interface {
	RegistrationRpc(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	GetProfileRpc(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfileRpc(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	mustEmbedUnimplementedStakeholderServiceServer()
}

// UnimplementedStakeholderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStakeholderServiceServer struct {
}

func (UnimplementedStakeholderServiceServer) RegistrationRpc(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationRpc not implemented")
}
func (UnimplementedStakeholderServiceServer) GetProfileRpc(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileRpc not implemented")
}
func (UnimplementedStakeholderServiceServer) UpdateProfileRpc(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfileRpc not implemented")
}
func (UnimplementedStakeholderServiceServer) mustEmbedUnimplementedStakeholderServiceServer() {}

// UnsafeStakeholderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StakeholderServiceServer will
// result in compilation errors.
type UnsafeStakeholderServiceServer interface {
	mustEmbedUnimplementedStakeholderServiceServer()
}

func RegisterStakeholderServiceServer(s grpc.ServiceRegistrar, srv StakeholderServiceServer) {
	s.RegisterService(&StakeholderService_ServiceDesc, srv)
}

func _StakeholderService_RegistrationRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).RegistrationRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_RegistrationRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).RegistrationRpc(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakeholderService_GetProfileRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).GetProfileRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_GetProfileRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).GetProfileRpc(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakeholderService_UpdateProfileRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).UpdateProfileRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_UpdateProfileRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).UpdateProfileRpc(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StakeholderService_ServiceDesc is the grpc.ServiceDesc for StakeholderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StakeholderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StakeholderService",
	HandlerType: (*StakeholderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrationRpc",
			Handler:    _StakeholderService_RegistrationRpc_Handler,
		},
		{
			MethodName: "GetProfileRpc",
			Handler:    _StakeholderService_GetProfileRpc_Handler,
		},
		{
			MethodName: "UpdateProfileRpc",
			Handler:    _StakeholderService_UpdateProfileRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}

const (
	AuthService_LoginRpc_FullMethodName = "/AuthService/LoginRpc"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	LoginRpc(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) LoginRpc(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_LoginRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	LoginRpc(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) LoginRpc(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginRpc not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_LoginRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_LoginRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginRpc(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginRpc",
			Handler:    _AuthService_LoginRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}
