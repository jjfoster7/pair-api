// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: proto/pairapi.proto

package pairapi

import (
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

// Create Account //
type AccountCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Pic   string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`
}

func (x *AccountCreateReq) Reset() {
	*x = AccountCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateReq) ProtoMessage() {}

func (x *AccountCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateReq.ProtoReflect.Descriptor instead.
func (*AccountCreateReq) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{0}
}

func (x *AccountCreateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountCreateReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountCreateReq) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

type AccountCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Pic      string `protobuf:"bytes,5,opt,name=pic,proto3" json:"pic,omitempty"`
	Points   int32  `protobuf:"varint,6,opt,name=points,proto3" json:"points,omitempty"`
	Peers    int32  `protobuf:"varint,7,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *AccountCreateResp) Reset() {
	*x = AccountCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateResp) ProtoMessage() {}

func (x *AccountCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateResp.ProtoReflect.Descriptor instead.
func (*AccountCreateResp) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{1}
}

func (x *AccountCreateResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountCreateResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountCreateResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountCreateResp) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AccountCreateResp) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *AccountCreateResp) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *AccountCreateResp) GetPeers() int32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

// Get Account //
type AccountGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountGetReq) Reset() {
	*x = AccountGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGetReq) ProtoMessage() {}

func (x *AccountGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGetReq.ProtoReflect.Descriptor instead.
func (*AccountGetReq) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{2}
}

func (x *AccountGetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AccountGetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Pic      string `protobuf:"bytes,5,opt,name=pic,proto3" json:"pic,omitempty"`
	Points   int32  `protobuf:"varint,6,opt,name=points,proto3" json:"points,omitempty"`
	Peers    int32  `protobuf:"varint,7,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *AccountGetResp) Reset() {
	*x = AccountGetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGetResp) ProtoMessage() {}

func (x *AccountGetResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGetResp.ProtoReflect.Descriptor instead.
func (*AccountGetResp) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{3}
}

func (x *AccountGetResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountGetResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountGetResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountGetResp) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AccountGetResp) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *AccountGetResp) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *AccountGetResp) GetPeers() int32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

// Update Account //
type AccountUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Column string `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AccountUpdateReq) Reset() {
	*x = AccountUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdateReq) ProtoMessage() {}

func (x *AccountUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdateReq.ProtoReflect.Descriptor instead.
func (*AccountUpdateReq) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{4}
}

func (x *AccountUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountUpdateReq) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *AccountUpdateReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AccountUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Column string `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AccountUpdateResp) Reset() {
	*x = AccountUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdateResp) ProtoMessage() {}

func (x *AccountUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdateResp.ProtoReflect.Descriptor instead.
func (*AccountUpdateResp) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{5}
}

func (x *AccountUpdateResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountUpdateResp) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *AccountUpdateResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Start Pair //
type PairStartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (x *PairStartReq) Reset() {
	*x = PairStartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairStartReq) ProtoMessage() {}

func (x *PairStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairStartReq.ProtoReflect.Descriptor instead.
func (*PairStartReq) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{6}
}

func (x *PairStartReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PairStartReq) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

type PairStartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *PairStartResp) Reset() {
	*x = PairStartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairStartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairStartResp) ProtoMessage() {}

func (x *PairStartResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairStartResp.ProtoReflect.Descriptor instead.
func (*PairStartResp) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{7}
}

func (x *PairStartResp) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *PairStartResp) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// Finish Pair //
type PairFinishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *PairFinishReq) Reset() {
	*x = PairFinishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairFinishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairFinishReq) ProtoMessage() {}

func (x *PairFinishReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairFinishReq.ProtoReflect.Descriptor instead.
func (*PairFinishReq) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{8}
}

func (x *PairFinishReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PairFinishReq) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *PairFinishReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type PairFinishResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created int64 `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *PairFinishResp) Reset() {
	*x = PairFinishResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pairapi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairFinishResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairFinishResp) ProtoMessage() {}

func (x *PairFinishResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pairapi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairFinishResp.ProtoReflect.Descriptor instead.
func (*PairFinishResp) Descriptor() ([]byte, []int) {
	return file_proto_pairapi_proto_rawDescGZIP(), []int{9}
}

func (x *PairFinishResp) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

var File_proto_pairapi_proto protoreflect.FileDescriptor

var file_proto_pairapi_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x22, 0x4a,
	0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x0d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x01,
	0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x10,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x11,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a,
	0x0c, 0x50, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x0d, 0x50, 0x61, 0x69, 0x72, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x50, 0x0a, 0x0d, 0x50, 0x61, 0x69, 0x72,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x50, 0x61,
	0x69, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0xdd, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x41,
	0x50, 0x49, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x69,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x69, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x61,
	0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61,
	0x69, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x61,
	0x69, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x61, 0x69, 0x72,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pairapi_proto_rawDescOnce sync.Once
	file_proto_pairapi_proto_rawDescData = file_proto_pairapi_proto_rawDesc
)

func file_proto_pairapi_proto_rawDescGZIP() []byte {
	file_proto_pairapi_proto_rawDescOnce.Do(func() {
		file_proto_pairapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pairapi_proto_rawDescData)
	})
	return file_proto_pairapi_proto_rawDescData
}

var file_proto_pairapi_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_pairapi_proto_goTypes = []interface{}{
	(*AccountCreateReq)(nil),  // 0: pairapi.AccountCreateReq
	(*AccountCreateResp)(nil), // 1: pairapi.AccountCreateResp
	(*AccountGetReq)(nil),     // 2: pairapi.AccountGetReq
	(*AccountGetResp)(nil),    // 3: pairapi.AccountGetResp
	(*AccountUpdateReq)(nil),  // 4: pairapi.AccountUpdateReq
	(*AccountUpdateResp)(nil), // 5: pairapi.AccountUpdateResp
	(*PairStartReq)(nil),      // 6: pairapi.PairStartReq
	(*PairStartResp)(nil),     // 7: pairapi.PairStartResp
	(*PairFinishReq)(nil),     // 8: pairapi.PairFinishReq
	(*PairFinishResp)(nil),    // 9: pairapi.PairFinishResp
}
var file_proto_pairapi_proto_depIdxs = []int32{
	0, // 0: pairapi.PairAPI.CreateAccount:input_type -> pairapi.AccountCreateReq
	2, // 1: pairapi.PairAPI.GetAccount:input_type -> pairapi.AccountGetReq
	4, // 2: pairapi.PairAPI.UpdateAccount:input_type -> pairapi.AccountUpdateReq
	6, // 3: pairapi.PairAPI.StartPair:input_type -> pairapi.PairStartReq
	8, // 4: pairapi.PairAPI.FinishPair:input_type -> pairapi.PairFinishReq
	1, // 5: pairapi.PairAPI.CreateAccount:output_type -> pairapi.AccountCreateResp
	3, // 6: pairapi.PairAPI.GetAccount:output_type -> pairapi.AccountGetResp
	5, // 7: pairapi.PairAPI.UpdateAccount:output_type -> pairapi.AccountUpdateResp
	7, // 8: pairapi.PairAPI.StartPair:output_type -> pairapi.PairStartResp
	9, // 9: pairapi.PairAPI.FinishPair:output_type -> pairapi.PairFinishResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pairapi_proto_init() }
func file_proto_pairapi_proto_init() {
	if File_proto_pairapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pairapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreateReq); i {
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
		file_proto_pairapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreateResp); i {
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
		file_proto_pairapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGetReq); i {
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
		file_proto_pairapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGetResp); i {
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
		file_proto_pairapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpdateReq); i {
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
		file_proto_pairapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpdateResp); i {
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
		file_proto_pairapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairStartReq); i {
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
		file_proto_pairapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairStartResp); i {
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
		file_proto_pairapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairFinishReq); i {
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
		file_proto_pairapi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairFinishResp); i {
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
			RawDescriptor: file_proto_pairapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pairapi_proto_goTypes,
		DependencyIndexes: file_proto_pairapi_proto_depIdxs,
		MessageInfos:      file_proto_pairapi_proto_msgTypes,
	}.Build()
	File_proto_pairapi_proto = out.File
	file_proto_pairapi_proto_rawDesc = nil
	file_proto_pairapi_proto_goTypes = nil
	file_proto_pairapi_proto_depIdxs = nil
}
