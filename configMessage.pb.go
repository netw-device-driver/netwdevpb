// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.3
// source: configMessage.proto

package netwdevpb

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

type ConfigMessage_MessageKind int32

const (
	ConfigMessage_Request  ConfigMessage_MessageKind = 0
	ConfigMessage_Response ConfigMessage_MessageKind = 1
)

// Enum value maps for ConfigMessage_MessageKind.
var (
	ConfigMessage_MessageKind_name = map[int32]string{
		0: "Request",
		1: "Response",
	}
	ConfigMessage_MessageKind_value = map[string]int32{
		"Request":  0,
		"Response": 1,
	}
)

func (x ConfigMessage_MessageKind) Enum() *ConfigMessage_MessageKind {
	p := new(ConfigMessage_MessageKind)
	*p = x
	return p
}

func (x ConfigMessage_MessageKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigMessage_MessageKind) Descriptor() protoreflect.EnumDescriptor {
	return file_configMessage_proto_enumTypes[0].Descriptor()
}

func (ConfigMessage_MessageKind) Type() protoreflect.EnumType {
	return &file_configMessage_proto_enumTypes[0]
}

func (x ConfigMessage_MessageKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigMessage_MessageKind.Descriptor instead.
func (ConfigMessage_MessageKind) EnumDescriptor() ([]byte, []int) {
	return file_configMessage_proto_rawDescGZIP(), []int{2, 0}
}

type ConfigMessage_ActionType int32

const (
	ConfigMessage_Replace ConfigMessage_ActionType = 0
	ConfigMessage_Update  ConfigMessage_ActionType = 1
	ConfigMessage_Delete  ConfigMessage_ActionType = 2
)

// Enum value maps for ConfigMessage_ActionType.
var (
	ConfigMessage_ActionType_name = map[int32]string{
		0: "Replace",
		1: "Update",
		2: "Delete",
	}
	ConfigMessage_ActionType_value = map[string]int32{
		"Replace": 0,
		"Update":  1,
		"Delete":  2,
	}
)

func (x ConfigMessage_ActionType) Enum() *ConfigMessage_ActionType {
	p := new(ConfigMessage_ActionType)
	*p = x
	return p
}

func (x ConfigMessage_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigMessage_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_configMessage_proto_enumTypes[1].Descriptor()
}

func (ConfigMessage_ActionType) Type() protoreflect.EnumType {
	return &file_configMessage_proto_enumTypes[1]
}

func (x ConfigMessage_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigMessage_ActionType.Descriptor instead.
func (ConfigMessage_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_configMessage_proto_rawDescGZIP(), []int{2, 1}
}

type CacheStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheObject string `protobuf:"bytes,1,opt,name=CacheObject,proto3" json:"CacheObject,omitempty"`
	Level       int32  `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *CacheStatusRequest) Reset() {
	*x = CacheStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheStatusRequest) ProtoMessage() {}

func (x *CacheStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheStatusRequest.ProtoReflect.Descriptor instead.
func (*CacheStatusRequest) Descriptor() ([]byte, []int) {
	return file_configMessage_proto_rawDescGZIP(), []int{0}
}

func (x *CacheStatusRequest) GetCacheObject() string {
	if x != nil {
		return x.CacheObject
	}
	return ""
}

func (x *CacheStatusRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type CacheStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists  string         `protobuf:"bytes,1,opt,name=Exists,proto3" json:"Exists,omitempty"`
	Message *ConfigMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CacheStatusReply) Reset() {
	*x = CacheStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheStatusReply) ProtoMessage() {}

func (x *CacheStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_configMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheStatusReply.ProtoReflect.Descriptor instead.
func (*CacheStatusReply) Descriptor() ([]byte, []int) {
	return file_configMessage_proto_rawDescGZIP(), []int{1}
}

func (x *CacheStatusReply) GetExists() string {
	if x != nil {
		return x.Exists
	}
	return ""
}

func (x *CacheStatusReply) GetMessage() *ConfigMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type ConfigMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind                        ConfigMessage_MessageKind `protobuf:"varint,1,opt,name=Kind,proto3,enum=netwdevpb.ConfigMessage_MessageKind" json:"Kind,omitempty"`
	ApiGroup                    string                    `protobuf:"bytes,2,opt,name=ApiGroup,proto3" json:"ApiGroup,omitempty"`
	Resource                    string                    `protobuf:"bytes,3,opt,name=Resource,proto3" json:"Resource,omitempty"`
	ResourceName                string                    `protobuf:"bytes,4,opt,name=ResourceName,proto3" json:"ResourceName,omitempty"`
	Namespace                   string                    `protobuf:"bytes,5,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Level                       int32                     `protobuf:"varint,6,opt,name=Level,proto3" json:"Level,omitempty"`
	Action                      ConfigMessage_ActionType  `protobuf:"varint,7,opt,name=Action,proto3,enum=netwdevpb.ConfigMessage_ActionType" json:"Action,omitempty"`
	AggregateActionPath         string                    `protobuf:"bytes,8,opt,name=AggregateActionPath,proto3" json:"AggregateActionPath,omitempty"`
	AggregateActionPathSuccess  bool                      `protobuf:"varint,9,opt,name=AggregateActionPathSuccess,proto3" json:"AggregateActionPathSuccess,omitempty"`
	IndividualActionPath        []string                  `protobuf:"bytes,10,rep,name=IndividualActionPath,proto3" json:"IndividualActionPath,omitempty"`
	IndividualActionPathSuccess []bool                    `protobuf:"varint,11,rep,packed,name=IndividualActionPathSuccess,proto3" json:"IndividualActionPathSuccess,omitempty"`
	ConfigData                  [][]byte                  `protobuf:"bytes,12,rep,name=ConfigData,proto3" json:"ConfigData,omitempty"`
	StatusData                  [][]byte                  `protobuf:"bytes,13,rep,name=StatusData,proto3" json:"StatusData,omitempty"`
	Dependencies                []string                  `protobuf:"bytes,14,rep,name=Dependencies,proto3" json:"Dependencies,omitempty"`
}

func (x *ConfigMessage) Reset() {
	*x = ConfigMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMessage) ProtoMessage() {}

func (x *ConfigMessage) ProtoReflect() protoreflect.Message {
	mi := &file_configMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMessage.ProtoReflect.Descriptor instead.
func (*ConfigMessage) Descriptor() ([]byte, []int) {
	return file_configMessage_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigMessage) GetKind() ConfigMessage_MessageKind {
	if x != nil {
		return x.Kind
	}
	return ConfigMessage_Request
}

func (x *ConfigMessage) GetApiGroup() string {
	if x != nil {
		return x.ApiGroup
	}
	return ""
}

func (x *ConfigMessage) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ConfigMessage) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ConfigMessage) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ConfigMessage) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ConfigMessage) GetAction() ConfigMessage_ActionType {
	if x != nil {
		return x.Action
	}
	return ConfigMessage_Replace
}

func (x *ConfigMessage) GetAggregateActionPath() string {
	if x != nil {
		return x.AggregateActionPath
	}
	return ""
}

func (x *ConfigMessage) GetAggregateActionPathSuccess() bool {
	if x != nil {
		return x.AggregateActionPathSuccess
	}
	return false
}

func (x *ConfigMessage) GetIndividualActionPath() []string {
	if x != nil {
		return x.IndividualActionPath
	}
	return nil
}

func (x *ConfigMessage) GetIndividualActionPathSuccess() []bool {
	if x != nil {
		return x.IndividualActionPathSuccess
	}
	return nil
}

func (x *ConfigMessage) GetConfigData() [][]byte {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *ConfigMessage) GetStatusData() [][]byte {
	if x != nil {
		return x.StatusData
	}
	return nil
}

func (x *ConfigMessage) GetDependencies() []string {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

var File_configMessage_proto protoreflect.FileDescriptor

var file_configMessage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x64, 0x65, 0x76, 0x70, 0x62,
	0x22, 0x4c, 0x0a, 0x12, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x5e,
	0x0a, 0x10, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x64, 0x65, 0x76, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbf,
	0x05, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x64, 0x65, 0x76, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x70,
	0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x70,
	0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x64, 0x65, 0x76, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x1a, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x49, 0x6e, 0x64,
	0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64,
	0x75, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a,
	0x1b, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x08, 0x52, 0x1b, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x22, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x01, 0x22, 0x31, 0x0a,
	0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x02,
	0x32, 0x56, 0x0a, 0x0b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x47, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x64, 0x65, 0x76, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x64, 0x65, 0x76, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x2d, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x64, 0x65,
	0x76, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configMessage_proto_rawDescOnce sync.Once
	file_configMessage_proto_rawDescData = file_configMessage_proto_rawDesc
)

func file_configMessage_proto_rawDescGZIP() []byte {
	file_configMessage_proto_rawDescOnce.Do(func() {
		file_configMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_configMessage_proto_rawDescData)
	})
	return file_configMessage_proto_rawDescData
}

var file_configMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_configMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_configMessage_proto_goTypes = []interface{}{
	(ConfigMessage_MessageKind)(0), // 0: netwdevpb.ConfigMessage.MessageKind
	(ConfigMessage_ActionType)(0),  // 1: netwdevpb.ConfigMessage.ActionType
	(*CacheStatusRequest)(nil),     // 2: netwdevpb.CacheStatusRequest
	(*CacheStatusReply)(nil),       // 3: netwdevpb.CacheStatusReply
	(*ConfigMessage)(nil),          // 4: netwdevpb.ConfigMessage
}
var file_configMessage_proto_depIdxs = []int32{
	4, // 0: netwdevpb.CacheStatusReply.message:type_name -> netwdevpb.ConfigMessage
	0, // 1: netwdevpb.ConfigMessage.Kind:type_name -> netwdevpb.ConfigMessage.MessageKind
	1, // 2: netwdevpb.ConfigMessage.Action:type_name -> netwdevpb.ConfigMessage.ActionType
	2, // 3: netwdevpb.CacheStatus.Request:input_type -> netwdevpb.CacheStatusRequest
	3, // 4: netwdevpb.CacheStatus.Request:output_type -> netwdevpb.CacheStatusReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_configMessage_proto_init() }
func file_configMessage_proto_init() {
	if File_configMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheStatusRequest); i {
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
		file_configMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheStatusReply); i {
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
		file_configMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMessage); i {
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
			RawDescriptor: file_configMessage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configMessage_proto_goTypes,
		DependencyIndexes: file_configMessage_proto_depIdxs,
		EnumInfos:         file_configMessage_proto_enumTypes,
		MessageInfos:      file_configMessage_proto_msgTypes,
	}.Build()
	File_configMessage_proto = out.File
	file_configMessage_proto_rawDesc = nil
	file_configMessage_proto_goTypes = nil
	file_configMessage_proto_depIdxs = nil
}