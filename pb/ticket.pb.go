// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: ticket.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket_Class int32

const (
	Ticket_Economy  Ticket_Class = 0
	Ticket_Business Ticket_Class = 1
)

// Enum value maps for Ticket_Class.
var (
	Ticket_Class_name = map[int32]string{
		0: "Economy",
		1: "Business",
	}
	Ticket_Class_value = map[string]int32{
		"Economy":  0,
		"Business": 1,
	}
)

func (x Ticket_Class) Enum() *Ticket_Class {
	p := new(Ticket_Class)
	*p = x
	return p
}

func (x Ticket_Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Class) Descriptor() protoreflect.EnumDescriptor {
	return file_ticket_proto_enumTypes[0].Descriptor()
}

func (Ticket_Class) Type() protoreflect.EnumType {
	return &file_ticket_proto_enumTypes[0]
}

func (x Ticket_Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Class.Descriptor instead.
func (Ticket_Class) EnumDescriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0, 0}
}

type Ticket_Status int32

const (
	Ticket_Paid     Ticket_Status = 0
	Ticket_Canceled Ticket_Status = 1
)

// Enum value maps for Ticket_Status.
var (
	Ticket_Status_name = map[int32]string{
		0: "Paid",
		1: "Canceled",
	}
	Ticket_Status_value = map[string]int32{
		"Paid":     0,
		"Canceled": 1,
	}
)

func (x Ticket_Status) Enum() *Ticket_Status {
	p := new(Ticket_Status)
	*p = x
	return p
}

func (x Ticket_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_ticket_proto_enumTypes[1].Descriptor()
}

func (Ticket_Status) Type() protoreflect.EnumType {
	return &file_ticket_proto_enumTypes[1]
}

func (x Ticket_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Status.Descriptor instead.
func (Ticket_Status) EnumDescriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0, 1}
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Class       Ticket_Class         `protobuf:"varint,2,opt,name=class,proto3,enum=training.Ticket_Class" json:"class,omitempty"`
	Status      Ticket_Status        `protobuf:"varint,3,opt,name=status,proto3,enum=training.Ticket_Status" json:"status,omitempty"`
	Flight      *Flight              `protobuf:"bytes,4,opt,name=flight,proto3" json:"flight,omitempty"`
	Booking     *Booking             `protobuf:"bytes,5,opt,name=booking,proto3" json:"booking,omitempty"`
	TicketOwner *TicketOwner         `protobuf:"bytes,6,opt,name=ticketOwner,proto3" json:"ticketOwner,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetClass() Ticket_Class {
	if x != nil {
		return x.Class
	}
	return Ticket_Economy
}

func (x *Ticket) GetStatus() Ticket_Status {
	if x != nil {
		return x.Status
	}
	return Ticket_Paid
}

func (x *Ticket) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

func (x *Ticket) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

func (x *Ticket) GetTicketOwner() *TicketOwner {
	if x != nil {
		return x.TicketOwner
	}
	return nil
}

func (x *Ticket) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetTicketByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketByIdRequest) Reset() {
	*x = GetTicketByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketByIdRequest) ProtoMessage() {}

func (x *GetTicketByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTicketByIdRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *GetTicketByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDeleted bool `protobuf:"varint,1,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTicketResponse) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03, 0x0a, 0x06, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x22, 0x0a, 0x05, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x10, 0x01, 0x22, 0x20, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x64, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x22,
	0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xad, 0x03, 0x0a, 0x10, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x50, 0x43, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3f, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x47,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x7d, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ticket_proto_goTypes = []interface{}{
	(Ticket_Class)(0),                            // 0: training.Ticket.Class
	(Ticket_Status)(0),                           // 1: training.Ticket.Status
	(*Ticket)(nil),                               // 2: training.Ticket
	(*GetTicketByIdRequest)(nil),                 // 3: training.GetTicketByIdRequest
	(*DeleteTicketRequest)(nil),                  // 4: training.DeleteTicketRequest
	(*DeleteTicketResponse)(nil),                 // 5: training.DeleteTicketResponse
	(*Flight)(nil),                               // 6: training.Flight
	(*Booking)(nil),                              // 7: training.Booking
	(*TicketOwner)(nil),                          // 8: training.TicketOwner
	(*timestamp.Timestamp)(nil),                  // 9: google.protobuf.Timestamp
	(*CreateTicketRequest)(nil),                  // 10: training.CreateTicketRequest
	(*UpdateTicketByIdRequest)(nil),              // 11: training.UpdateTicketByIdRequest
	(*UpdateTicketStatusByFlightIdRequest)(nil),  // 12: training.UpdateTicketStatusByFlightIdRequest
	(*UpdateTicketStatusByFlightIdResponse)(nil), // 13: training.UpdateTicketStatusByFlightIdResponse
}
var file_ticket_proto_depIdxs = []int32{
	0,  // 0: training.Ticket.class:type_name -> training.Ticket.Class
	1,  // 1: training.Ticket.status:type_name -> training.Ticket.Status
	6,  // 2: training.Ticket.flight:type_name -> training.Flight
	7,  // 3: training.Ticket.booking:type_name -> training.Booking
	8,  // 4: training.Ticket.ticketOwner:type_name -> training.TicketOwner
	9,  // 5: training.Ticket.created_at:type_name -> google.protobuf.Timestamp
	9,  // 6: training.Ticket.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 7: training.TicketServiceRPC.GetTicketById:input_type -> training.GetTicketByIdRequest
	10, // 8: training.TicketServiceRPC.CreateTicket:input_type -> training.CreateTicketRequest
	11, // 9: training.TicketServiceRPC.UpdateTicketById:input_type -> training.UpdateTicketByIdRequest
	12, // 10: training.TicketServiceRPC.UpdateTicketStatusByFlightId:input_type -> training.UpdateTicketStatusByFlightIdRequest
	4,  // 11: training.TicketServiceRPC.DeleteTicket:input_type -> training.DeleteTicketRequest
	2,  // 12: training.TicketServiceRPC.GetTicketById:output_type -> training.Ticket
	2,  // 13: training.TicketServiceRPC.CreateTicket:output_type -> training.Ticket
	2,  // 14: training.TicketServiceRPC.UpdateTicketById:output_type -> training.Ticket
	13, // 15: training.TicketServiceRPC.UpdateTicketStatusByFlightId:output_type -> training.UpdateTicketStatusByFlightIdResponse
	5,  // 16: training.TicketServiceRPC.DeleteTicket:output_type -> training.DeleteTicketResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	file_flight_proto_init()
	file_ticketOwner_proto_init()
	file_booking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketByIdRequest); i {
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
		file_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketRequest); i {
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
		file_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketResponse); i {
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
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		EnumInfos:         file_ticket_proto_enumTypes,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
