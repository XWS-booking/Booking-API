// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: reservation_service.proto

package reservationGrpc

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
	ReservationService_Create_FullMethodName                                     = "/ReservationService/Create"
	ReservationService_Delete_FullMethodName                                     = "/ReservationService/Delete"
	ReservationService_Confirm_FullMethodName                                    = "/ReservationService/Confirm"
	ReservationService_Reject_FullMethodName                                     = "/ReservationService/Reject"
	ReservationService_FindAllReservedAccommodations_FullMethodName              = "/ReservationService/FindAllReservedAccommodations"
	ReservationService_CheckActiveReservationsForGuest_FullMethodName            = "/ReservationService/CheckActiveReservationsForGuest"
	ReservationService_CheckActiveReservationsForAccommodations_FullMethodName   = "/ReservationService/CheckActiveReservationsForAccommodations"
	ReservationService_CancelReservation_FullMethodName                          = "/ReservationService/CancelReservation"
	ReservationService_IsAccommodationAvailable_FullMethodName                   = "/ReservationService/IsAccommodationAvailable"
	ReservationService_FindAllByBuyerId_FullMethodName                           = "/ReservationService/FindAllByBuyerId"
	ReservationService_FindAllByAccommodationId_FullMethodName                   = "/ReservationService/FindAllByAccommodationId"
	ReservationService_FindNumberOfBuyersCancellations_FullMethodName            = "/ReservationService/FindNumberOfBuyersCancellations"
	ReservationService_UpdateReservationRating_FullMethodName                    = "/ReservationService/UpdateReservationRating"
	ReservationService_CheckIfGuestHasReservationInAccommodations_FullMethodName = "/ReservationService/CheckIfGuestHasReservationInAccommodations"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Create(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*ReservationId, error)
	Delete(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	Confirm(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*ReservationResponse, error)
	Reject(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*ReservationResponse, error)
	FindAllReservedAccommodations(ctx context.Context, in *FindAllReservedAccommodationsRequest, opts ...grpc.CallOption) (*FindAllReservedAccommodationsResponse, error)
	CheckActiveReservationsForGuest(ctx context.Context, in *CheckActiveReservationsForGuestRequest, opts ...grpc.CallOption) (*CheckActiveReservationsForGuestResponse, error)
	CheckActiveReservationsForAccommodations(ctx context.Context, in *CheckActiveReservationsForAccommodationsRequest, opts ...grpc.CallOption) (*CheckActiveReservationsForAccommodationsResponse, error)
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	IsAccommodationAvailable(ctx context.Context, in *IsAccommodationAvailableRequest, opts ...grpc.CallOption) (*IsAccommodationAvailableResponse, error)
	FindAllByBuyerId(ctx context.Context, in *FindAllReservationsByBuyerIdRequest, opts ...grpc.CallOption) (*FindAllReservationsByBuyerIdResponse, error)
	FindAllByAccommodationId(ctx context.Context, in *FindAllReservationsByAccommodationIdRequest, opts ...grpc.CallOption) (*FindAllReservationsByAccommodationIdResponse, error)
	FindNumberOfBuyersCancellations(ctx context.Context, in *NumberOfCancellationRequest, opts ...grpc.CallOption) (*NumberOfCancellationResponse, error)
	UpdateReservationRating(ctx context.Context, in *UpdateReservationRatingRequest, opts ...grpc.CallOption) (*UpdateReservationRatingResponse, error)
	CheckIfGuestHasReservationInAccommodations(ctx context.Context, in *CheckIfGuestHasReservationInAccommodationsRequest, opts ...grpc.CallOption) (*CheckIfGuestHasReservationInAccommodationsResponse, error)
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

func (c *reservationServiceClient) Confirm(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_Confirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Reject(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_Reject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) FindAllReservedAccommodations(ctx context.Context, in *FindAllReservedAccommodationsRequest, opts ...grpc.CallOption) (*FindAllReservedAccommodationsResponse, error) {
	out := new(FindAllReservedAccommodationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_FindAllReservedAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckActiveReservationsForGuest(ctx context.Context, in *CheckActiveReservationsForGuestRequest, opts ...grpc.CallOption) (*CheckActiveReservationsForGuestResponse, error) {
	out := new(CheckActiveReservationsForGuestResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckActiveReservationsForGuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckActiveReservationsForAccommodations(ctx context.Context, in *CheckActiveReservationsForAccommodationsRequest, opts ...grpc.CallOption) (*CheckActiveReservationsForAccommodationsResponse, error) {
	out := new(CheckActiveReservationsForAccommodationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckActiveReservationsForAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CancelReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) IsAccommodationAvailable(ctx context.Context, in *IsAccommodationAvailableRequest, opts ...grpc.CallOption) (*IsAccommodationAvailableResponse, error) {
	out := new(IsAccommodationAvailableResponse)
	err := c.cc.Invoke(ctx, ReservationService_IsAccommodationAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) FindAllByBuyerId(ctx context.Context, in *FindAllReservationsByBuyerIdRequest, opts ...grpc.CallOption) (*FindAllReservationsByBuyerIdResponse, error) {
	out := new(FindAllReservationsByBuyerIdResponse)
	err := c.cc.Invoke(ctx, ReservationService_FindAllByBuyerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) FindAllByAccommodationId(ctx context.Context, in *FindAllReservationsByAccommodationIdRequest, opts ...grpc.CallOption) (*FindAllReservationsByAccommodationIdResponse, error) {
	out := new(FindAllReservationsByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, ReservationService_FindAllByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) FindNumberOfBuyersCancellations(ctx context.Context, in *NumberOfCancellationRequest, opts ...grpc.CallOption) (*NumberOfCancellationResponse, error) {
	out := new(NumberOfCancellationResponse)
	err := c.cc.Invoke(ctx, ReservationService_FindNumberOfBuyersCancellations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateReservationRating(ctx context.Context, in *UpdateReservationRatingRequest, opts ...grpc.CallOption) (*UpdateReservationRatingResponse, error) {
	out := new(UpdateReservationRatingResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateReservationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckIfGuestHasReservationInAccommodations(ctx context.Context, in *CheckIfGuestHasReservationInAccommodationsRequest, opts ...grpc.CallOption) (*CheckIfGuestHasReservationInAccommodationsResponse, error) {
	out := new(CheckIfGuestHasReservationInAccommodationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckIfGuestHasReservationInAccommodations_FullMethodName, in, out, opts...)
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
	Confirm(context.Context, *ReservationId) (*ReservationResponse, error)
	Reject(context.Context, *ReservationId) (*ReservationResponse, error)
	FindAllReservedAccommodations(context.Context, *FindAllReservedAccommodationsRequest) (*FindAllReservedAccommodationsResponse, error)
	CheckActiveReservationsForGuest(context.Context, *CheckActiveReservationsForGuestRequest) (*CheckActiveReservationsForGuestResponse, error)
	CheckActiveReservationsForAccommodations(context.Context, *CheckActiveReservationsForAccommodationsRequest) (*CheckActiveReservationsForAccommodationsResponse, error)
	CancelReservation(context.Context, *CancelReservationRequest) (*ReservationResponse, error)
	IsAccommodationAvailable(context.Context, *IsAccommodationAvailableRequest) (*IsAccommodationAvailableResponse, error)
	FindAllByBuyerId(context.Context, *FindAllReservationsByBuyerIdRequest) (*FindAllReservationsByBuyerIdResponse, error)
	FindAllByAccommodationId(context.Context, *FindAllReservationsByAccommodationIdRequest) (*FindAllReservationsByAccommodationIdResponse, error)
	FindNumberOfBuyersCancellations(context.Context, *NumberOfCancellationRequest) (*NumberOfCancellationResponse, error)
	UpdateReservationRating(context.Context, *UpdateReservationRatingRequest) (*UpdateReservationRatingResponse, error)
	CheckIfGuestHasReservationInAccommodations(context.Context, *CheckIfGuestHasReservationInAccommodationsRequest) (*CheckIfGuestHasReservationInAccommodationsResponse, error)
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
func (UnimplementedReservationServiceServer) Confirm(context.Context, *ReservationId) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedReservationServiceServer) Reject(context.Context, *ReservationId) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reject not implemented")
}
func (UnimplementedReservationServiceServer) FindAllReservedAccommodations(context.Context, *FindAllReservedAccommodationsRequest) (*FindAllReservedAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllReservedAccommodations not implemented")
}
func (UnimplementedReservationServiceServer) CheckActiveReservationsForGuest(context.Context, *CheckActiveReservationsForGuestRequest) (*CheckActiveReservationsForGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckActiveReservationsForGuest not implemented")
}
func (UnimplementedReservationServiceServer) CheckActiveReservationsForAccommodations(context.Context, *CheckActiveReservationsForAccommodationsRequest) (*CheckActiveReservationsForAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckActiveReservationsForAccommodations not implemented")
}
func (UnimplementedReservationServiceServer) CancelReservation(context.Context, *CancelReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedReservationServiceServer) IsAccommodationAvailable(context.Context, *IsAccommodationAvailableRequest) (*IsAccommodationAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAccommodationAvailable not implemented")
}
func (UnimplementedReservationServiceServer) FindAllByBuyerId(context.Context, *FindAllReservationsByBuyerIdRequest) (*FindAllReservationsByBuyerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllByBuyerId not implemented")
}
func (UnimplementedReservationServiceServer) FindAllByAccommodationId(context.Context, *FindAllReservationsByAccommodationIdRequest) (*FindAllReservationsByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllByAccommodationId not implemented")
}
func (UnimplementedReservationServiceServer) FindNumberOfBuyersCancellations(context.Context, *NumberOfCancellationRequest) (*NumberOfCancellationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNumberOfBuyersCancellations not implemented")
}
func (UnimplementedReservationServiceServer) UpdateReservationRating(context.Context, *UpdateReservationRatingRequest) (*UpdateReservationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservationRating not implemented")
}
func (UnimplementedReservationServiceServer) CheckIfGuestHasReservationInAccommodations(context.Context, *CheckIfGuestHasReservationInAccommodationsRequest) (*CheckIfGuestHasReservationInAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfGuestHasReservationInAccommodations not implemented")
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

func _ReservationService_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Confirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Confirm(ctx, req.(*ReservationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Reject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Reject(ctx, req.(*ReservationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_FindAllReservedAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReservedAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).FindAllReservedAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_FindAllReservedAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).FindAllReservedAccommodations(ctx, req.(*FindAllReservedAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckActiveReservationsForGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckActiveReservationsForGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckActiveReservationsForGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckActiveReservationsForGuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckActiveReservationsForGuest(ctx, req.(*CheckActiveReservationsForGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckActiveReservationsForAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckActiveReservationsForAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckActiveReservationsForAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckActiveReservationsForAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckActiveReservationsForAccommodations(ctx, req.(*CheckActiveReservationsForAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CancelReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CancelReservation(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_IsAccommodationAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAccommodationAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).IsAccommodationAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_IsAccommodationAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).IsAccommodationAvailable(ctx, req.(*IsAccommodationAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_FindAllByBuyerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReservationsByBuyerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).FindAllByBuyerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_FindAllByBuyerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).FindAllByBuyerId(ctx, req.(*FindAllReservationsByBuyerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_FindAllByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReservationsByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).FindAllByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_FindAllByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).FindAllByAccommodationId(ctx, req.(*FindAllReservationsByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_FindNumberOfBuyersCancellations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberOfCancellationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).FindNumberOfBuyersCancellations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_FindNumberOfBuyersCancellations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).FindNumberOfBuyersCancellations(ctx, req.(*NumberOfCancellationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateReservationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateReservationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateReservationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateReservationRating(ctx, req.(*UpdateReservationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckIfGuestHasReservationInAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfGuestHasReservationInAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckIfGuestHasReservationInAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckIfGuestHasReservationInAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckIfGuestHasReservationInAccommodations(ctx, req.(*CheckIfGuestHasReservationInAccommodationsRequest))
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
		{
			MethodName: "Confirm",
			Handler:    _ReservationService_Confirm_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _ReservationService_Reject_Handler,
		},
		{
			MethodName: "FindAllReservedAccommodations",
			Handler:    _ReservationService_FindAllReservedAccommodations_Handler,
		},
		{
			MethodName: "CheckActiveReservationsForGuest",
			Handler:    _ReservationService_CheckActiveReservationsForGuest_Handler,
		},
		{
			MethodName: "CheckActiveReservationsForAccommodations",
			Handler:    _ReservationService_CheckActiveReservationsForAccommodations_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _ReservationService_CancelReservation_Handler,
		},
		{
			MethodName: "IsAccommodationAvailable",
			Handler:    _ReservationService_IsAccommodationAvailable_Handler,
		},
		{
			MethodName: "FindAllByBuyerId",
			Handler:    _ReservationService_FindAllByBuyerId_Handler,
		},
		{
			MethodName: "FindAllByAccommodationId",
			Handler:    _ReservationService_FindAllByAccommodationId_Handler,
		},
		{
			MethodName: "FindNumberOfBuyersCancellations",
			Handler:    _ReservationService_FindNumberOfBuyersCancellations_Handler,
		},
		{
			MethodName: "UpdateReservationRating",
			Handler:    _ReservationService_UpdateReservationRating_Handler,
		},
		{
			MethodName: "CheckIfGuestHasReservationInAccommodations",
			Handler:    _ReservationService_CheckIfGuestHasReservationInAccommodations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service.proto",
}
