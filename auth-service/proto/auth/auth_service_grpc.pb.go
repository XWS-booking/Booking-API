// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: auth_service.proto

package authGrpc

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
	AuthService_SignIn_FullMethodName                        = "/AuthService/SignIn"
	AuthService_Register_FullMethodName                      = "/AuthService/Register"
	AuthService_UpdatePersonalInfo_FullMethodName            = "/AuthService/UpdatePersonalInfo"
	AuthService_DeleteProfile_FullMethodName                 = "/AuthService/DeleteProfile"
	AuthService_GetUser_FullMethodName                       = "/AuthService/GetUser"
	AuthService_FindById_FullMethodName                      = "/AuthService/FindById"
	AuthService_ChangePassword_FullMethodName                = "/AuthService/ChangePassword"
	AuthService_GetHostRatingWithGuestInfo_FullMethodName    = "/AuthService/GetHostRatingWithGuestInfo"
	AuthService_ProfileDeletion_FullMethodName               = "/AuthService/ProfileDeletion"
	AuthService_ChangeHostDistinguishedStatus_FullMethodName = "/AuthService/ChangeHostDistinguishedStatus"
	AuthService_GetFeaturedHosts_FullMethodName              = "/AuthService/GetFeaturedHosts"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error)
	DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	FindById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	GetHostRatingWithGuestInfo(ctx context.Context, in *GetHostRatingWithGuestInfoRequest, opts ...grpc.CallOption) (*GetHostRatingWithGuestInfoResponse, error)
	ProfileDeletion(ctx context.Context, in *ProfileDeletionRequest, opts ...grpc.CallOption) (*ProfileDeletionResponse, error)
	ChangeHostDistinguishedStatus(ctx context.Context, in *ChangeHostDistinguishedStatusRequest, opts ...grpc.CallOption) (*ChangeHostDistinguishedStatusResponse, error)
	GetFeaturedHosts(ctx context.Context, in *GetFeaturedHostsRequest, opts ...grpc.CallOption) (*GetFeaturedHostsResponse, error)
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

func (c *authServiceClient) UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error) {
	out := new(UpdatePersonalInfoResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdatePersonalInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error) {
	out := new(DeleteProfileResponse)
	err := c.cc.Invoke(ctx, AuthService_DeleteProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, AuthService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) FindById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error) {
	out := new(FindUserByIdResponse)
	err := c.cc.Invoke(ctx, AuthService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetHostRatingWithGuestInfo(ctx context.Context, in *GetHostRatingWithGuestInfoRequest, opts ...grpc.CallOption) (*GetHostRatingWithGuestInfoResponse, error) {
	out := new(GetHostRatingWithGuestInfoResponse)
	err := c.cc.Invoke(ctx, AuthService_GetHostRatingWithGuestInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ProfileDeletion(ctx context.Context, in *ProfileDeletionRequest, opts ...grpc.CallOption) (*ProfileDeletionResponse, error) {
	out := new(ProfileDeletionResponse)
	err := c.cc.Invoke(ctx, AuthService_ProfileDeletion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangeHostDistinguishedStatus(ctx context.Context, in *ChangeHostDistinguishedStatusRequest, opts ...grpc.CallOption) (*ChangeHostDistinguishedStatusResponse, error) {
	out := new(ChangeHostDistinguishedStatusResponse)
	err := c.cc.Invoke(ctx, AuthService_ChangeHostDistinguishedStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFeaturedHosts(ctx context.Context, in *GetFeaturedHostsRequest, opts ...grpc.CallOption) (*GetFeaturedHostsResponse, error) {
	out := new(GetFeaturedHostsResponse)
	err := c.cc.Invoke(ctx, AuthService_GetFeaturedHosts_FullMethodName, in, out, opts...)
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
	UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error)
	DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	FindById(context.Context, *FindUserByIdRequest) (*FindUserByIdResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	GetHostRatingWithGuestInfo(context.Context, *GetHostRatingWithGuestInfoRequest) (*GetHostRatingWithGuestInfoResponse, error)
	ProfileDeletion(context.Context, *ProfileDeletionRequest) (*ProfileDeletionResponse, error)
	ChangeHostDistinguishedStatus(context.Context, *ChangeHostDistinguishedStatusRequest) (*ChangeHostDistinguishedStatusResponse, error)
	GetFeaturedHosts(context.Context, *GetFeaturedHostsRequest) (*GetFeaturedHostsResponse, error)
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
func (UnimplementedAuthServiceServer) UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalInfo not implemented")
}
func (UnimplementedAuthServiceServer) DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedAuthServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServiceServer) FindById(context.Context, *FindUserByIdRequest) (*FindUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedAuthServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServiceServer) GetHostRatingWithGuestInfo(context.Context, *GetHostRatingWithGuestInfoRequest) (*GetHostRatingWithGuestInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostRatingWithGuestInfo not implemented")
}
func (UnimplementedAuthServiceServer) ProfileDeletion(context.Context, *ProfileDeletionRequest) (*ProfileDeletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileDeletion not implemented")
}
func (UnimplementedAuthServiceServer) ChangeHostDistinguishedStatus(context.Context, *ChangeHostDistinguishedStatusRequest) (*ChangeHostDistinguishedStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeHostDistinguishedStatus not implemented")
}
func (UnimplementedAuthServiceServer) GetFeaturedHosts(context.Context, *GetFeaturedHostsRequest) (*GetFeaturedHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeaturedHosts not implemented")
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

func _AuthService_UpdatePersonalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePersonalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdatePersonalInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePersonalInfo(ctx, req.(*UpdatePersonalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteProfile(ctx, req.(*DeleteProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).FindById(ctx, req.(*FindUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetHostRatingWithGuestInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostRatingWithGuestInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetHostRatingWithGuestInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetHostRatingWithGuestInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetHostRatingWithGuestInfo(ctx, req.(*GetHostRatingWithGuestInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ProfileDeletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ProfileDeletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ProfileDeletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ProfileDeletion(ctx, req.(*ProfileDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangeHostDistinguishedStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeHostDistinguishedStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangeHostDistinguishedStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangeHostDistinguishedStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangeHostDistinguishedStatus(ctx, req.(*ChangeHostDistinguishedStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFeaturedHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeaturedHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFeaturedHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetFeaturedHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFeaturedHosts(ctx, req.(*GetFeaturedHostsRequest))
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
		{
			MethodName: "UpdatePersonalInfo",
			Handler:    _AuthService_UpdatePersonalInfo_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _AuthService_DeleteProfile_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _AuthService_FindById_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthService_ChangePassword_Handler,
		},
		{
			MethodName: "GetHostRatingWithGuestInfo",
			Handler:    _AuthService_GetHostRatingWithGuestInfo_Handler,
		},
		{
			MethodName: "ProfileDeletion",
			Handler:    _AuthService_ProfileDeletion_Handler,
		},
		{
			MethodName: "ChangeHostDistinguishedStatus",
			Handler:    _AuthService_ChangeHostDistinguishedStatus_Handler,
		},
		{
			MethodName: "GetFeaturedHosts",
			Handler:    _AuthService_GetFeaturedHosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
