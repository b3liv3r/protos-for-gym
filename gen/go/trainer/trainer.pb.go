// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: trainer/trainer.proto

package trainerv1

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

type ListForGymRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GymId int64 `protobuf:"varint,1,opt,name=gym_id,json=gymId,proto3" json:"gym_id,omitempty"`
}

func (x *ListForGymRequest) Reset() {
	*x = ListForGymRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForGymRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForGymRequest) ProtoMessage() {}

func (x *ListForGymRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForGymRequest.ProtoReflect.Descriptor instead.
func (*ListForGymRequest) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *ListForGymRequest) GetGymId() int64 {
	if x != nil {
		return x.GymId
	}
	return 0
}

type ListForGymResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trainers []*Trainers `protobuf:"bytes,1,rep,name=trainers,proto3" json:"trainers,omitempty"`
}

func (x *ListForGymResponse) Reset() {
	*x = ListForGymResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForGymResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForGymResponse) ProtoMessage() {}

func (x *ListForGymResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForGymResponse.ProtoReflect.Descriptor instead.
func (*ListForGymResponse) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *ListForGymResponse) GetTrainers() []*Trainers {
	if x != nil {
		return x.Trainers
	}
	return nil
}

type AvailableBookingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainerId int64 `protobuf:"varint,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
}

func (x *AvailableBookingListRequest) Reset() {
	*x = AvailableBookingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableBookingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableBookingListRequest) ProtoMessage() {}

func (x *AvailableBookingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableBookingListRequest.ProtoReflect.Descriptor instead.
func (*AvailableBookingListRequest) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *AvailableBookingListRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

type AvailableBookingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*Bookings `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *AvailableBookingListResponse) Reset() {
	*x = AvailableBookingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableBookingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableBookingListResponse) ProtoMessage() {}

func (x *AvailableBookingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableBookingListResponse.ProtoReflect.Descriptor instead.
func (*AvailableBookingListResponse) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{3}
}

func (x *AvailableBookingListResponse) GetBookings() []*Bookings {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type CurrentBookingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CurrentBookingListRequest) Reset() {
	*x = CurrentBookingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBookingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBookingListRequest) ProtoMessage() {}

func (x *CurrentBookingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBookingListRequest.ProtoReflect.Descriptor instead.
func (*CurrentBookingListRequest) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{4}
}

func (x *CurrentBookingListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CurrentBookingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*Bookings `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *CurrentBookingListResponse) Reset() {
	*x = CurrentBookingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBookingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBookingListResponse) ProtoMessage() {}

func (x *CurrentBookingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBookingListResponse.ProtoReflect.Descriptor instead.
func (*CurrentBookingListResponse) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{5}
}

func (x *CurrentBookingListResponse) GetBookings() []*Bookings {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableBookingId int64 `protobuf:"varint,1,opt,name=available_booking_id,json=availableBookingId,proto3" json:"available_booking_id,omitempty"`
	UserId             int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{6}
}

func (x *BookingRequest) GetAvailableBookingId() int64 {
	if x != nil {
		return x.AvailableBookingId
	}
	return 0
}

func (x *BookingRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{7}
}

func (x *BookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UnBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentBookingId int64 `protobuf:"varint,1,opt,name=current_booking_id,json=currentBookingId,proto3" json:"current_booking_id,omitempty"`
}

func (x *UnBookingRequest) Reset() {
	*x = UnBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnBookingRequest) ProtoMessage() {}

func (x *UnBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnBookingRequest.ProtoReflect.Descriptor instead.
func (*UnBookingRequest) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{8}
}

func (x *UnBookingRequest) GetCurrentBookingId() int64 {
	if x != nil {
		return x.CurrentBookingId
	}
	return 0
}

type UnBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UnBookingResponse) Reset() {
	*x = UnBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnBookingResponse) ProtoMessage() {}

func (x *UnBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnBookingResponse.ProtoReflect.Descriptor instead.
func (*UnBookingResponse) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{9}
}

func (x *UnBookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Trainers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainerId  int64  `protobuf:"varint,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	GymId      int64  `protobuf:"varint,2,opt,name=gym_id,json=gymId,proto3" json:"gym_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Speciality string `protobuf:"bytes,4,opt,name=speciality,proto3" json:"speciality,omitempty"`
}

func (x *Trainers) Reset() {
	*x = Trainers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trainers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trainers) ProtoMessage() {}

func (x *Trainers) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trainers.ProtoReflect.Descriptor instead.
func (*Trainers) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{10}
}

func (x *Trainers) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *Trainers) GetGymId() int64 {
	if x != nil {
		return x.GymId
	}
	return 0
}

func (x *Trainers) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trainers) GetSpeciality() string {
	if x != nil {
		return x.Speciality
	}
	return ""
}

type Bookings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int64                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	Activity  string                 `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Bookings) Reset() {
	*x = Bookings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_trainer_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookings) ProtoMessage() {}

func (x *Bookings) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_trainer_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookings.ProtoReflect.Descriptor instead.
func (*Bookings) Descriptor() ([]byte, []int) {
	return file_trainer_trainer_proto_rawDescGZIP(), []int{11}
}

func (x *Bookings) GetBookingId() int64 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *Bookings) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *Bookings) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Bookings) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_trainer_trainer_proto protoreflect.FileDescriptor

var file_trainer_trainer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x79, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x79, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x79, 0x6d, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x79, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x22, 0x3c, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4d, 0x0a, 0x1c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0x34, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x1a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x5b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x10,
	0x55, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2d,
	0x0a, 0x11, 0x55, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a,
	0x08, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x79, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x79, 0x6d, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x96, 0x03,
	0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x47, 0x79, 0x6d, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x79, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x79, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x14, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x55, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x79, 0x6d, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trainer_trainer_proto_rawDescOnce sync.Once
	file_trainer_trainer_proto_rawDescData = file_trainer_trainer_proto_rawDesc
)

func file_trainer_trainer_proto_rawDescGZIP() []byte {
	file_trainer_trainer_proto_rawDescOnce.Do(func() {
		file_trainer_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_trainer_trainer_proto_rawDescData)
	})
	return file_trainer_trainer_proto_rawDescData
}

var file_trainer_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_trainer_trainer_proto_goTypes = []interface{}{
	(*ListForGymRequest)(nil),            // 0: trainer.ListForGymRequest
	(*ListForGymResponse)(nil),           // 1: trainer.ListForGymResponse
	(*AvailableBookingListRequest)(nil),  // 2: trainer.AvailableBookingListRequest
	(*AvailableBookingListResponse)(nil), // 3: trainer.AvailableBookingListResponse
	(*CurrentBookingListRequest)(nil),    // 4: trainer.CurrentBookingListRequest
	(*CurrentBookingListResponse)(nil),   // 5: trainer.CurrentBookingListResponse
	(*BookingRequest)(nil),               // 6: trainer.BookingRequest
	(*BookingResponse)(nil),              // 7: trainer.BookingResponse
	(*UnBookingRequest)(nil),             // 8: trainer.UnBookingRequest
	(*UnBookingResponse)(nil),            // 9: trainer.UnBookingResponse
	(*Trainers)(nil),                     // 10: trainer.Trainers
	(*Bookings)(nil),                     // 11: trainer.Bookings
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
}
var file_trainer_trainer_proto_depIdxs = []int32{
	10, // 0: trainer.ListForGymResponse.trainers:type_name -> trainer.Trainers
	11, // 1: trainer.AvailableBookingListResponse.bookings:type_name -> trainer.Bookings
	11, // 2: trainer.CurrentBookingListResponse.bookings:type_name -> trainer.Bookings
	12, // 3: trainer.Bookings.start_time:type_name -> google.protobuf.Timestamp
	12, // 4: trainer.Bookings.end_time:type_name -> google.protobuf.Timestamp
	0,  // 5: trainer.Trainer.ListForGym:input_type -> trainer.ListForGymRequest
	2,  // 6: trainer.Trainer.AvailableBookingList:input_type -> trainer.AvailableBookingListRequest
	4,  // 7: trainer.Trainer.CurrentBookingList:input_type -> trainer.CurrentBookingListRequest
	6,  // 8: trainer.Trainer.Booking:input_type -> trainer.BookingRequest
	8,  // 9: trainer.Trainer.UnBooking:input_type -> trainer.UnBookingRequest
	1,  // 10: trainer.Trainer.ListForGym:output_type -> trainer.ListForGymResponse
	3,  // 11: trainer.Trainer.AvailableBookingList:output_type -> trainer.AvailableBookingListResponse
	5,  // 12: trainer.Trainer.CurrentBookingList:output_type -> trainer.CurrentBookingListResponse
	7,  // 13: trainer.Trainer.Booking:output_type -> trainer.BookingResponse
	9,  // 14: trainer.Trainer.UnBooking:output_type -> trainer.UnBookingResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_trainer_trainer_proto_init() }
func file_trainer_trainer_proto_init() {
	if File_trainer_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trainer_trainer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForGymRequest); i {
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
		file_trainer_trainer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForGymResponse); i {
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
		file_trainer_trainer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableBookingListRequest); i {
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
		file_trainer_trainer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableBookingListResponse); i {
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
		file_trainer_trainer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentBookingListRequest); i {
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
		file_trainer_trainer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentBookingListResponse); i {
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
		file_trainer_trainer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_trainer_trainer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
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
		file_trainer_trainer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnBookingRequest); i {
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
		file_trainer_trainer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnBookingResponse); i {
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
		file_trainer_trainer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trainers); i {
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
		file_trainer_trainer_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookings); i {
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
			RawDescriptor: file_trainer_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trainer_trainer_proto_goTypes,
		DependencyIndexes: file_trainer_trainer_proto_depIdxs,
		MessageInfos:      file_trainer_trainer_proto_msgTypes,
	}.Build()
	File_trainer_trainer_proto = out.File
	file_trainer_trainer_proto_rawDesc = nil
	file_trainer_trainer_proto_goTypes = nil
	file_trainer_trainer_proto_depIdxs = nil
}
