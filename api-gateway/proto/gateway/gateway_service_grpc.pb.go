// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: gateway_service.proto

package gateway

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
	AuthService_SignIn_FullMethodName   = "/AuthService/SignIn"
	AuthService_Register_FullMethodName = "/AuthService/Register"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, AuthService_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, AuthService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	Register(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
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

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegistrationRequest))
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
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_service.proto",
}

const (
	AccomodationService_Create_FullMethodName  = "/AccomodationService/Create"
	AccomodationService_FindAll_FullMethodName = "/AccomodationService/FindAll"
)

// AccomodationServiceClient is the client API for AccomodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccomodationServiceClient interface {
	Create(ctx context.Context, in *CreateAccomodationRequest, opts ...grpc.CallOption) (*CreateAccomodationResponse, error)
	FindAll(ctx context.Context, in *FindAllAccomodationRequest, opts ...grpc.CallOption) (*FindAllAccomodationResponse, error)
}

type accomodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccomodationServiceClient(cc grpc.ClientConnInterface) AccomodationServiceClient {
	return &accomodationServiceClient{cc}
}

func (c *accomodationServiceClient) Create(ctx context.Context, in *CreateAccomodationRequest, opts ...grpc.CallOption) (*CreateAccomodationResponse, error) {
	out := new(CreateAccomodationResponse)
	err := c.cc.Invoke(ctx, AccomodationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) FindAll(ctx context.Context, in *FindAllAccomodationRequest, opts ...grpc.CallOption) (*FindAllAccomodationResponse, error) {
	out := new(FindAllAccomodationResponse)
	err := c.cc.Invoke(ctx, AccomodationService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccomodationServiceServer is the server API for AccomodationService service.
// All implementations must embed UnimplementedAccomodationServiceServer
// for forward compatibility
type AccomodationServiceServer interface {
	Create(context.Context, *CreateAccomodationRequest) (*CreateAccomodationResponse, error)
	FindAll(context.Context, *FindAllAccomodationRequest) (*FindAllAccomodationResponse, error)
	mustEmbedUnimplementedAccomodationServiceServer()
}

// UnimplementedAccomodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccomodationServiceServer struct {
}

func (UnimplementedAccomodationServiceServer) Create(context.Context, *CreateAccomodationRequest) (*CreateAccomodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccomodationServiceServer) FindAll(context.Context, *FindAllAccomodationRequest) (*FindAllAccomodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedAccomodationServiceServer) mustEmbedUnimplementedAccomodationServiceServer() {}

// UnsafeAccomodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccomodationServiceServer will
// result in compilation errors.
type UnsafeAccomodationServiceServer interface {
	mustEmbedUnimplementedAccomodationServiceServer()
}

func RegisterAccomodationServiceServer(s grpc.ServiceRegistrar, srv AccomodationServiceServer) {
	s.RegisterService(&AccomodationService_ServiceDesc, srv)
}

func _AccomodationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccomodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).Create(ctx, req.(*CreateAccomodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAccomodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).FindAll(ctx, req.(*FindAllAccomodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccomodationService_ServiceDesc is the grpc.ServiceDesc for AccomodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccomodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccomodationService",
	HandlerType: (*AccomodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccomodationService_Create_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _AccomodationService_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_service.proto",
}

const (
	ReservationService_Create_FullMethodName = "/ReservationService/Create"
	ReservationService_Delete_FullMethodName = "/ReservationService/Delete"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Create(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*ReservationId, error)
	Delete(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Create(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*ReservationId, error) {
	out := new(ReservationId)
	err := c.cc.Invoke(ctx, ReservationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Delete(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Create(context.Context, *CreateReservationRequest) (*ReservationId, error)
	Delete(context.Context, *ReservationId) (*DeleteReservationResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Create(context.Context, *CreateReservationRequest) (*ReservationId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReservationServiceServer) Delete(context.Context, *ReservationId) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Create(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Delete(ctx, req.(*ReservationId))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReservationService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_service.proto",
}
