// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: business/business.proto

package business

import (
	finance "github.com/oriastanjung/grpc_server/proto/finance"
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

type CreateBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateBusinessRequest) Reset() {
	*x = CreateBusinessRequest{}
	mi := &file_business_business_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessRequest) ProtoMessage() {}

func (x *CreateBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessRequest.ProtoReflect.Descriptor instead.
func (*CreateBusinessRequest) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBusinessRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBusinessRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateBusinessResponse) Reset() {
	*x = CreateBusinessResponse{}
	mi := &file_business_business_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessResponse) ProtoMessage() {}

func (x *CreateBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessResponse.ProtoReflect.Descriptor instead.
func (*CreateBusinessResponse) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBusinessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateBusinessRequest) Reset() {
	*x = UpdateBusinessRequest{}
	mi := &file_business_business_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBusinessRequest) ProtoMessage() {}

func (x *UpdateBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBusinessRequest.ProtoReflect.Descriptor instead.
func (*UpdateBusinessRequest) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBusinessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBusinessRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBusinessRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateBusinessResponse) Reset() {
	*x = UpdateBusinessResponse{}
	mi := &file_business_business_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBusinessResponse) ProtoMessage() {}

func (x *UpdateBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBusinessResponse.ProtoReflect.Descriptor instead.
func (*UpdateBusinessResponse) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBusinessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBusinessRequest) Reset() {
	*x = DeleteBusinessRequest{}
	mi := &file_business_business_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBusinessRequest) ProtoMessage() {}

func (x *DeleteBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBusinessRequest.ProtoReflect.Descriptor instead.
func (*DeleteBusinessRequest) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBusinessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBusinessResponse) Reset() {
	*x = DeleteBusinessResponse{}
	mi := &file_business_business_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBusinessResponse) ProtoMessage() {}

func (x *DeleteBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBusinessResponse.ProtoReflect.Descriptor instead.
func (*DeleteBusinessResponse) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBusinessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BusinessModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	UserId      string                  `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt   string                  `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string                  `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Finance     []*finance.FinanceModel `protobuf:"bytes,7,rep,name=finance,proto3" json:"finance,omitempty"`
}

func (x *BusinessModel) Reset() {
	*x = BusinessModel{}
	mi := &file_business_business_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BusinessModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessModel) ProtoMessage() {}

func (x *BusinessModel) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessModel.ProtoReflect.Descriptor instead.
func (*BusinessModel) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{6}
}

func (x *BusinessModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BusinessModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BusinessModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BusinessModel) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BusinessModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BusinessModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BusinessModel) GetFinance() []*finance.FinanceModel {
	if x != nil {
		return x.Finance
	}
	return nil
}

type GetBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBusinessRequest) Reset() {
	*x = GetBusinessRequest{}
	mi := &file_business_business_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessRequest) ProtoMessage() {}

func (x *GetBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessRequest) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{7}
}

func (x *GetBusinessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Business *BusinessModel `protobuf:"bytes,1,opt,name=business,proto3" json:"business,omitempty"`
}

func (x *GetBusinessResponse) Reset() {
	*x = GetBusinessResponse{}
	mi := &file_business_business_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessResponse) ProtoMessage() {}

func (x *GetBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessResponse) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{8}
}

func (x *GetBusinessResponse) GetBusiness() *BusinessModel {
	if x != nil {
		return x.Business
	}
	return nil
}

type ListBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ListBusinessRequest) Reset() {
	*x = ListBusinessRequest{}
	mi := &file_business_business_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBusinessRequest) ProtoMessage() {}

func (x *ListBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBusinessRequest.ProtoReflect.Descriptor instead.
func (*ListBusinessRequest) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{9}
}

func (x *ListBusinessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Business []*BusinessModel `protobuf:"bytes,1,rep,name=business,proto3" json:"business,omitempty"`
}

func (x *ListBusinessResponse) Reset() {
	*x = ListBusinessResponse{}
	mi := &file_business_business_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBusinessResponse) ProtoMessage() {}

func (x *ListBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_business_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBusinessResponse.ProtoReflect.Descriptor instead.
func (*ListBusinessResponse) Descriptor() ([]byte, []int) {
	return file_business_business_proto_rawDescGZIP(), []int{10}
}

func (x *ListBusinessResponse) GetBusiness() []*BusinessModel {
	if x != nil {
		return x.Business
	}
	return nil
}

var File_business_business_proto protoreflect.FileDescriptor

var file_business_business_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x1a, 0x15, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xda, 0x01,
	0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x32, 0xb1, 0x03, 0x0a, 0x15, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x61, 0x73,
	0x74, 0x61, 0x6e, 0x6a, 0x75, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_business_business_proto_rawDescOnce sync.Once
	file_business_business_proto_rawDescData = file_business_business_proto_rawDesc
)

func file_business_business_proto_rawDescGZIP() []byte {
	file_business_business_proto_rawDescOnce.Do(func() {
		file_business_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_business_business_proto_rawDescData)
	})
	return file_business_business_proto_rawDescData
}

var file_business_business_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_business_business_proto_goTypes = []any{
	(*CreateBusinessRequest)(nil),  // 0: business.CreateBusinessRequest
	(*CreateBusinessResponse)(nil), // 1: business.CreateBusinessResponse
	(*UpdateBusinessRequest)(nil),  // 2: business.UpdateBusinessRequest
	(*UpdateBusinessResponse)(nil), // 3: business.UpdateBusinessResponse
	(*DeleteBusinessRequest)(nil),  // 4: business.DeleteBusinessRequest
	(*DeleteBusinessResponse)(nil), // 5: business.DeleteBusinessResponse
	(*BusinessModel)(nil),          // 6: business.BusinessModel
	(*GetBusinessRequest)(nil),     // 7: business.GetBusinessRequest
	(*GetBusinessResponse)(nil),    // 8: business.GetBusinessResponse
	(*ListBusinessRequest)(nil),    // 9: business.ListBusinessRequest
	(*ListBusinessResponse)(nil),   // 10: business.ListBusinessResponse
	(*finance.FinanceModel)(nil),   // 11: finance.FinanceModel
}
var file_business_business_proto_depIdxs = []int32{
	11, // 0: business.BusinessModel.finance:type_name -> finance.FinanceModel
	6,  // 1: business.GetBusinessResponse.business:type_name -> business.BusinessModel
	6,  // 2: business.ListBusinessResponse.business:type_name -> business.BusinessModel
	0,  // 3: business.BusinessRoutesService.CreateBusiness:input_type -> business.CreateBusinessRequest
	2,  // 4: business.BusinessRoutesService.UpdateBusiness:input_type -> business.UpdateBusinessRequest
	4,  // 5: business.BusinessRoutesService.DeleteBusiness:input_type -> business.DeleteBusinessRequest
	7,  // 6: business.BusinessRoutesService.GetBusiness:input_type -> business.GetBusinessRequest
	9,  // 7: business.BusinessRoutesService.ListBusiness:input_type -> business.ListBusinessRequest
	1,  // 8: business.BusinessRoutesService.CreateBusiness:output_type -> business.CreateBusinessResponse
	3,  // 9: business.BusinessRoutesService.UpdateBusiness:output_type -> business.UpdateBusinessResponse
	5,  // 10: business.BusinessRoutesService.DeleteBusiness:output_type -> business.DeleteBusinessResponse
	8,  // 11: business.BusinessRoutesService.GetBusiness:output_type -> business.GetBusinessResponse
	10, // 12: business.BusinessRoutesService.ListBusiness:output_type -> business.ListBusinessResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_business_business_proto_init() }
func file_business_business_proto_init() {
	if File_business_business_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_business_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_business_business_proto_goTypes,
		DependencyIndexes: file_business_business_proto_depIdxs,
		MessageInfos:      file_business_business_proto_msgTypes,
	}.Build()
	File_business_business_proto = out.File
	file_business_business_proto_rawDesc = nil
	file_business_business_proto_goTypes = nil
	file_business_business_proto_depIdxs = nil
}