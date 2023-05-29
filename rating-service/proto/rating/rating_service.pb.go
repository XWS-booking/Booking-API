// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: rating_service.proto

package ratingGrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RateAccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,1,opt,name=AccommodationId,proto3" json:"AccommodationId,omitempty"`
	Rating          int32  `protobuf:"varint,2,opt,name=Rating,proto3" json:"Rating,omitempty"`
	GuestId         string `protobuf:"bytes,3,opt,name=GuestId,proto3" json:"GuestId,omitempty"`
}

func (x *RateAccommodationRequest) Reset() {
	*x = RateAccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationRequest) ProtoMessage() {}

func (x *RateAccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationRequest.ProtoReflect.Descriptor instead.
func (*RateAccommodationRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{0}
}

func (x *RateAccommodationRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *RateAccommodationRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RateAccommodationRequest) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

type RateAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RateAccommodationResponse) Reset() {
	*x = RateAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationResponse) ProtoMessage() {}

func (x *RateAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationResponse.ProtoReflect.Descriptor instead.
func (*RateAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{1}
}

func (x *RateAccommodationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccommodationRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteAccommodationRatingRequest) Reset() {
	*x = DeleteAccommodationRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccommodationRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccommodationRatingRequest) ProtoMessage() {}

func (x *DeleteAccommodationRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccommodationRatingRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccommodationRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAccommodationRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccommodationRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccommodationRatingResponse) Reset() {
	*x = DeleteAccommodationRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccommodationRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccommodationRatingResponse) ProtoMessage() {}

func (x *DeleteAccommodationRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccommodationRatingResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccommodationRatingResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{3}
}

type UpdateAccommodationRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Rating int32  `protobuf:"varint,2,opt,name=Rating,proto3" json:"Rating,omitempty"`
}

func (x *UpdateAccommodationRatingRequest) Reset() {
	*x = UpdateAccommodationRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccommodationRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccommodationRatingRequest) ProtoMessage() {}

func (x *UpdateAccommodationRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccommodationRatingRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccommodationRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAccommodationRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAccommodationRatingRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type UpdateAccommodationRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccommodationRatingResponse) Reset() {
	*x = UpdateAccommodationRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccommodationRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccommodationRatingResponse) ProtoMessage() {}

func (x *UpdateAccommodationRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccommodationRatingResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccommodationRatingResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{5}
}

type GetAllAccommodationRatingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,1,opt,name=AccommodationId,proto3" json:"AccommodationId,omitempty"`
}

func (x *GetAllAccommodationRatingsRequest) Reset() {
	*x = GetAllAccommodationRatingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAccommodationRatingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAccommodationRatingsRequest) ProtoMessage() {}

func (x *GetAllAccommodationRatingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAccommodationRatingsRequest.ProtoReflect.Descriptor instead.
func (*GetAllAccommodationRatingsRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllAccommodationRatingsRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

type GetAllAccommodationRatingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings []*AccommodationRatingItem `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *GetAllAccommodationRatingsResponse) Reset() {
	*x = GetAllAccommodationRatingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAccommodationRatingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAccommodationRatingsResponse) ProtoMessage() {}

func (x *GetAllAccommodationRatingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAccommodationRatingsResponse.ProtoReflect.Descriptor instead.
func (*GetAllAccommodationRatingsResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllAccommodationRatingsResponse) GetRatings() []*AccommodationRatingItem {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type AccommodationRatingItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccommodationId string                 `protobuf:"bytes,2,opt,name=AccommodationId,proto3" json:"AccommodationId,omitempty"`
	Rating          int32                  `protobuf:"varint,3,opt,name=Rating,proto3" json:"Rating,omitempty"`
	GuestId         string                 `protobuf:"bytes,4,opt,name=GuestId,proto3" json:"GuestId,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *AccommodationRatingItem) Reset() {
	*x = AccommodationRatingItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRatingItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRatingItem) ProtoMessage() {}

func (x *AccommodationRatingItem) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRatingItem.ProtoReflect.Descriptor instead.
func (*AccommodationRatingItem) Descriptor() ([]byte, []int) {
	return file_rating_service_proto_rawDescGZIP(), []int{8}
}

func (x *AccommodationRatingItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccommodationRatingItem) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *AccommodationRatingItem) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AccommodationRatingItem) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *AccommodationRatingItem) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_rating_service_proto protoreflect.FileDescriptor

var file_rating_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x18, 0x52, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x19, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x20,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x23, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x23, 0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0xb5, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x92, 0x03, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x52, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rating_service_proto_rawDescOnce sync.Once
	file_rating_service_proto_rawDescData = file_rating_service_proto_rawDesc
)

func file_rating_service_proto_rawDescGZIP() []byte {
	file_rating_service_proto_rawDescOnce.Do(func() {
		file_rating_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rating_service_proto_rawDescData)
	})
	return file_rating_service_proto_rawDescData
}

var file_rating_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rating_service_proto_goTypes = []interface{}{
	(*RateAccommodationRequest)(nil),           // 0: RateAccommodationRequest
	(*RateAccommodationResponse)(nil),          // 1: RateAccommodationResponse
	(*DeleteAccommodationRatingRequest)(nil),   // 2: DeleteAccommodationRatingRequest
	(*DeleteAccommodationRatingResponse)(nil),  // 3: DeleteAccommodationRatingResponse
	(*UpdateAccommodationRatingRequest)(nil),   // 4: UpdateAccommodationRatingRequest
	(*UpdateAccommodationRatingResponse)(nil),  // 5: UpdateAccommodationRatingResponse
	(*GetAllAccommodationRatingsRequest)(nil),  // 6: GetAllAccommodationRatingsRequest
	(*GetAllAccommodationRatingsResponse)(nil), // 7: GetAllAccommodationRatingsResponse
	(*AccommodationRatingItem)(nil),            // 8: AccommodationRatingItem
	(*timestamppb.Timestamp)(nil),              // 9: google.protobuf.Timestamp
}
var file_rating_service_proto_depIdxs = []int32{
	8, // 0: GetAllAccommodationRatingsResponse.ratings:type_name -> AccommodationRatingItem
	9, // 1: AccommodationRatingItem.Time:type_name -> google.protobuf.Timestamp
	0, // 2: RatingService.RateAccommodation:input_type -> RateAccommodationRequest
	2, // 3: RatingService.DeleteAccommodationRating:input_type -> DeleteAccommodationRatingRequest
	4, // 4: RatingService.UpdateAccommodationRating:input_type -> UpdateAccommodationRatingRequest
	6, // 5: RatingService.GetAllAccommodationRatings:input_type -> GetAllAccommodationRatingsRequest
	1, // 6: RatingService.RateAccommodation:output_type -> RateAccommodationResponse
	3, // 7: RatingService.DeleteAccommodationRating:output_type -> DeleteAccommodationRatingResponse
	5, // 8: RatingService.UpdateAccommodationRating:output_type -> UpdateAccommodationRatingResponse
	7, // 9: RatingService.GetAllAccommodationRatings:output_type -> GetAllAccommodationRatingsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rating_service_proto_init() }
func file_rating_service_proto_init() {
	if File_rating_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rating_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccommodationRatingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccommodationRatingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccommodationRatingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccommodationRatingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAccommodationRatingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAccommodationRatingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rating_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRatingItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rating_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rating_service_proto_goTypes,
		DependencyIndexes: file_rating_service_proto_depIdxs,
		MessageInfos:      file_rating_service_proto_msgTypes,
	}.Build()
	File_rating_service_proto = out.File
	file_rating_service_proto_rawDesc = nil
	file_rating_service_proto_goTypes = nil
	file_rating_service_proto_depIdxs = nil
}
