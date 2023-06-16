// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: accomodation_service.proto

package accomodationGrpc

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
	AccomodationService_Create_FullMethodName                           = "/AccomodationService/Create"
	AccomodationService_FindAll_FullMethodName                          = "/AccomodationService/FindAll"
	AccomodationService_FindAllAccommodationIdsByOwnerId_FullMethodName = "/AccomodationService/FindAllAccommodationIdsByOwnerId"
	AccomodationService_DeleteByOwnerId_FullMethodName                  = "/AccomodationService/DeleteByOwnerId"
	AccomodationService_FindById_FullMethodName                         = "/AccomodationService/FindById"
	AccomodationService_GetBookingPrice_FullMethodName                  = "/AccomodationService/GetBookingPrice"
	AccomodationService_UpdatePricing_FullMethodName                    = "/AccomodationService/UpdatePricing"
	AccomodationService_SearchAndFilter_FullMethodName                  = "/AccomodationService/SearchAndFilter"
	AccomodationService_PopulateRecommended_FullMethodName              = "/AccomodationService/PopulateRecommended"
)

// AccomodationServiceClient is the client API for AccomodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccomodationServiceClient interface {
	Create(ctx context.Context, in *CreateAccomodationRequest, opts ...grpc.CallOption) (*CreateAccomodationResponse, error)
	FindAll(ctx context.Context, in *FindAllAccomodationRequest, opts ...grpc.CallOption) (*FindAllAccomodationResponse, error)
	FindAllAccommodationIdsByOwnerId(ctx context.Context, in *FindAllAccommodationIdsByOwnerIdRequest, opts ...grpc.CallOption) (*FindAllAccommodationIdsByOwnerIdResponse, error)
	DeleteByOwnerId(ctx context.Context, in *DeleteByOwnerIdRequest, opts ...grpc.CallOption) (*DeleteByOwnerIdResponse, error)
	FindById(ctx context.Context, in *FindAccommodationByIdRequest, opts ...grpc.CallOption) (*AccomodationResponse, error)
	GetBookingPrice(ctx context.Context, in *GetBookingPriceRequest, opts ...grpc.CallOption) (*GetBookingPriceResponse, error)
	UpdatePricing(ctx context.Context, in *UpdatePricingRequest, opts ...grpc.CallOption) (*UpdatePricingResponse, error)
	SearchAndFilter(ctx context.Context, in *SearchAndFilterRequest, opts ...grpc.CallOption) (*SearchAndFilterResponse, error)
	PopulateRecommended(ctx context.Context, in *PopulateRecommendedRequest, opts ...grpc.CallOption) (*PopulateRecommendedResponse, error)
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

func (c *accomodationServiceClient) FindAllAccommodationIdsByOwnerId(ctx context.Context, in *FindAllAccommodationIdsByOwnerIdRequest, opts ...grpc.CallOption) (*FindAllAccommodationIdsByOwnerIdResponse, error) {
	out := new(FindAllAccommodationIdsByOwnerIdResponse)
	err := c.cc.Invoke(ctx, AccomodationService_FindAllAccommodationIdsByOwnerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) DeleteByOwnerId(ctx context.Context, in *DeleteByOwnerIdRequest, opts ...grpc.CallOption) (*DeleteByOwnerIdResponse, error) {
	out := new(DeleteByOwnerIdResponse)
	err := c.cc.Invoke(ctx, AccomodationService_DeleteByOwnerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) FindById(ctx context.Context, in *FindAccommodationByIdRequest, opts ...grpc.CallOption) (*AccomodationResponse, error) {
	out := new(AccomodationResponse)
	err := c.cc.Invoke(ctx, AccomodationService_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) GetBookingPrice(ctx context.Context, in *GetBookingPriceRequest, opts ...grpc.CallOption) (*GetBookingPriceResponse, error) {
	out := new(GetBookingPriceResponse)
	err := c.cc.Invoke(ctx, AccomodationService_GetBookingPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) UpdatePricing(ctx context.Context, in *UpdatePricingRequest, opts ...grpc.CallOption) (*UpdatePricingResponse, error) {
	out := new(UpdatePricingResponse)
	err := c.cc.Invoke(ctx, AccomodationService_UpdatePricing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) SearchAndFilter(ctx context.Context, in *SearchAndFilterRequest, opts ...grpc.CallOption) (*SearchAndFilterResponse, error) {
	out := new(SearchAndFilterResponse)
	err := c.cc.Invoke(ctx, AccomodationService_SearchAndFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationServiceClient) PopulateRecommended(ctx context.Context, in *PopulateRecommendedRequest, opts ...grpc.CallOption) (*PopulateRecommendedResponse, error) {
	out := new(PopulateRecommendedResponse)
	err := c.cc.Invoke(ctx, AccomodationService_PopulateRecommended_FullMethodName, in, out, opts...)
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
	FindAllAccommodationIdsByOwnerId(context.Context, *FindAllAccommodationIdsByOwnerIdRequest) (*FindAllAccommodationIdsByOwnerIdResponse, error)
	DeleteByOwnerId(context.Context, *DeleteByOwnerIdRequest) (*DeleteByOwnerIdResponse, error)
	FindById(context.Context, *FindAccommodationByIdRequest) (*AccomodationResponse, error)
	GetBookingPrice(context.Context, *GetBookingPriceRequest) (*GetBookingPriceResponse, error)
	UpdatePricing(context.Context, *UpdatePricingRequest) (*UpdatePricingResponse, error)
	SearchAndFilter(context.Context, *SearchAndFilterRequest) (*SearchAndFilterResponse, error)
	PopulateRecommended(context.Context, *PopulateRecommendedRequest) (*PopulateRecommendedResponse, error)
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
func (UnimplementedAccomodationServiceServer) FindAllAccommodationIdsByOwnerId(context.Context, *FindAllAccommodationIdsByOwnerIdRequest) (*FindAllAccommodationIdsByOwnerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAccommodationIdsByOwnerId not implemented")
}
func (UnimplementedAccomodationServiceServer) DeleteByOwnerId(context.Context, *DeleteByOwnerIdRequest) (*DeleteByOwnerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByOwnerId not implemented")
}
func (UnimplementedAccomodationServiceServer) FindById(context.Context, *FindAccommodationByIdRequest) (*AccomodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedAccomodationServiceServer) GetBookingPrice(context.Context, *GetBookingPriceRequest) (*GetBookingPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingPrice not implemented")
}
func (UnimplementedAccomodationServiceServer) UpdatePricing(context.Context, *UpdatePricingRequest) (*UpdatePricingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePricing not implemented")
}
func (UnimplementedAccomodationServiceServer) SearchAndFilter(context.Context, *SearchAndFilterRequest) (*SearchAndFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAndFilter not implemented")
}
func (UnimplementedAccomodationServiceServer) PopulateRecommended(context.Context, *PopulateRecommendedRequest) (*PopulateRecommendedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopulateRecommended not implemented")
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

func _AccomodationService_FindAllAccommodationIdsByOwnerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAccommodationIdsByOwnerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).FindAllAccommodationIdsByOwnerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_FindAllAccommodationIdsByOwnerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).FindAllAccommodationIdsByOwnerId(ctx, req.(*FindAllAccommodationIdsByOwnerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_DeleteByOwnerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByOwnerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).DeleteByOwnerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_DeleteByOwnerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).DeleteByOwnerId(ctx, req.(*DeleteByOwnerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccommodationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).FindById(ctx, req.(*FindAccommodationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_GetBookingPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).GetBookingPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_GetBookingPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).GetBookingPrice(ctx, req.(*GetBookingPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_UpdatePricing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePricingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).UpdatePricing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_UpdatePricing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).UpdatePricing(ctx, req.(*UpdatePricingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_SearchAndFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAndFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).SearchAndFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_SearchAndFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).SearchAndFilter(ctx, req.(*SearchAndFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationService_PopulateRecommended_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopulateRecommendedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).PopulateRecommended(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_PopulateRecommended_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).PopulateRecommended(ctx, req.(*PopulateRecommendedRequest))
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
		{
			MethodName: "FindAllAccommodationIdsByOwnerId",
			Handler:    _AccomodationService_FindAllAccommodationIdsByOwnerId_Handler,
		},
		{
			MethodName: "DeleteByOwnerId",
			Handler:    _AccomodationService_DeleteByOwnerId_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _AccomodationService_FindById_Handler,
		},
		{
			MethodName: "GetBookingPrice",
			Handler:    _AccomodationService_GetBookingPrice_Handler,
		},
		{
			MethodName: "UpdatePricing",
			Handler:    _AccomodationService_UpdatePricing_Handler,
		},
		{
			MethodName: "SearchAndFilter",
			Handler:    _AccomodationService_SearchAndFilter_Handler,
		},
		{
			MethodName: "PopulateRecommended",
			Handler:    _AccomodationService_PopulateRecommended_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accomodation_service.proto",
}
