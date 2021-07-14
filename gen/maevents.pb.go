// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.13.0
// source: maevents.proto

package gen

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MAUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId  string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ServiceId string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ExtraId   string `protobuf:"bytes,3,opt,name=extra_id,json=extraId,proto3" json:"extra_id,omitempty"`
}

func (x *MAUser) Reset() {
	*x = MAUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maevents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAUser) ProtoMessage() {}

func (x *MAUser) ProtoReflect() protoreflect.Message {
	mi := &file_maevents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAUser.ProtoReflect.Descriptor instead.
func (*MAUser) Descriptor() ([]byte, []int) {
	return file_maevents_proto_rawDescGZIP(), []int{0}
}

func (x *MAUser) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *MAUser) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *MAUser) GetExtraId() string {
	if x != nil {
		return x.ExtraId
	}
	return ""
}

type MARegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *MAUser              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Os          string               `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	AppId       string               `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppVersion  string               `protobuf:"bytes,4,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Time        *timestamp.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Screen      string               `protobuf:"bytes,6,opt,name=screen,proto3" json:"screen,omitempty"`
	Ip          string               `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	NetworkType string               `protobuf:"bytes,8,opt,name=network_type,json=networkType,proto3" json:"network_type,omitempty"`
	Carrier     string               `protobuf:"bytes,9,opt,name=carrier,proto3" json:"carrier,omitempty"`
	Channel     string               `protobuf:"bytes,10,opt,name=channel,proto3" json:"channel,omitempty"`
	Extra       string               `protobuf:"bytes,11,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *MARegister) Reset() {
	*x = MARegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maevents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MARegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MARegister) ProtoMessage() {}

func (x *MARegister) ProtoReflect() protoreflect.Message {
	mi := &file_maevents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MARegister.ProtoReflect.Descriptor instead.
func (*MARegister) Descriptor() ([]byte, []int) {
	return file_maevents_proto_rawDescGZIP(), []int{1}
}

func (x *MARegister) GetUser() *MAUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MARegister) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *MARegister) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *MARegister) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *MARegister) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MARegister) GetScreen() string {
	if x != nil {
		return x.Screen
	}
	return ""
}

func (x *MARegister) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *MARegister) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *MARegister) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *MARegister) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *MARegister) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type MAEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User             *MAUser              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AppId            string               `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	SessionId        string               `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Time             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Ip               string               `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	NetworkType      string               `protobuf:"bytes,6,opt,name=network_type,json=networkType,proto3" json:"network_type,omitempty"`
	Carrier          string               `protobuf:"bytes,7,opt,name=carrier,proto3" json:"carrier,omitempty"`
	PageTitle        string               `protobuf:"bytes,8,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	PageUrl          string               `protobuf:"bytes,9,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	SourceUrl        string               `protobuf:"bytes,10,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	SourceModule     string               `protobuf:"bytes,11,opt,name=source_module,json=sourceModule,proto3" json:"source_module,omitempty"`
	EventName        string               `protobuf:"bytes,12,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	EventAttr1Type   string               `protobuf:"bytes,21,opt,name=event_attr1_type,json=eventAttr1Type,proto3" json:"event_attr1_type,omitempty"`
	EventAttr1Value  string               `protobuf:"bytes,22,opt,name=event_attr1_value,json=eventAttr1Value,proto3" json:"event_attr1_value,omitempty"`
	EventAttr2Type   string               `protobuf:"bytes,23,opt,name=event_attr2_type,json=eventAttr2Type,proto3" json:"event_attr2_type,omitempty"`
	EventAttr2Value  string               `protobuf:"bytes,24,opt,name=event_attr2_value,json=eventAttr2Value,proto3" json:"event_attr2_value,omitempty"`
	EventAttr3Type   string               `protobuf:"bytes,25,opt,name=event_attr3_type,json=eventAttr3Type,proto3" json:"event_attr3_type,omitempty"`
	EventAttr3Value  string               `protobuf:"bytes,26,opt,name=event_attr3_value,json=eventAttr3Value,proto3" json:"event_attr3_value,omitempty"`
	EventAttr4Type   string               `protobuf:"bytes,27,opt,name=event_attr4_type,json=eventAttr4Type,proto3" json:"event_attr4_type,omitempty"`
	EventAttr4Value  string               `protobuf:"bytes,28,opt,name=event_attr4_value,json=eventAttr4Value,proto3" json:"event_attr4_value,omitempty"`
	EventAttr5Type   string               `protobuf:"bytes,29,opt,name=event_attr5_type,json=eventAttr5Type,proto3" json:"event_attr5_type,omitempty"`
	EventAttr5Value  string               `protobuf:"bytes,30,opt,name=event_attr5_value,json=eventAttr5Value,proto3" json:"event_attr5_value,omitempty"`
	EventAttr6Type   string               `protobuf:"bytes,31,opt,name=event_attr6_type,json=eventAttr6Type,proto3" json:"event_attr6_type,omitempty"`
	EventAttr6Value  string               `protobuf:"bytes,32,opt,name=event_attr6_value,json=eventAttr6Value,proto3" json:"event_attr6_value,omitempty"`
	EventAttr7Type   string               `protobuf:"bytes,33,opt,name=event_attr7_type,json=eventAttr7Type,proto3" json:"event_attr7_type,omitempty"`
	EventAttr7Value  string               `protobuf:"bytes,34,opt,name=event_attr7_value,json=eventAttr7Value,proto3" json:"event_attr7_value,omitempty"`
	EventAttr8Type   string               `protobuf:"bytes,35,opt,name=event_attr8_type,json=eventAttr8Type,proto3" json:"event_attr8_type,omitempty"`
	EventAttr8Value  string               `protobuf:"bytes,36,opt,name=event_attr8_value,json=eventAttr8Value,proto3" json:"event_attr8_value,omitempty"`
	EventAttr9Type   string               `protobuf:"bytes,37,opt,name=event_attr9_type,json=eventAttr9Type,proto3" json:"event_attr9_type,omitempty"`
	EventAttr9Value  string               `protobuf:"bytes,38,opt,name=event_attr9_value,json=eventAttr9Value,proto3" json:"event_attr9_value,omitempty"`
	EventAttr10Type  string               `protobuf:"bytes,39,opt,name=event_attr10_type,json=eventAttr10Type,proto3" json:"event_attr10_type,omitempty"`
	EventAttr10Value string               `protobuf:"bytes,40,opt,name=event_attr10_value,json=eventAttr10Value,proto3" json:"event_attr10_value,omitempty"`
}

func (x *MAEvent) Reset() {
	*x = MAEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maevents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAEvent) ProtoMessage() {}

func (x *MAEvent) ProtoReflect() protoreflect.Message {
	mi := &file_maevents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAEvent.ProtoReflect.Descriptor instead.
func (*MAEvent) Descriptor() ([]byte, []int) {
	return file_maevents_proto_rawDescGZIP(), []int{2}
}

func (x *MAEvent) GetUser() *MAUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MAEvent) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *MAEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *MAEvent) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MAEvent) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *MAEvent) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *MAEvent) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *MAEvent) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *MAEvent) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *MAEvent) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *MAEvent) GetSourceModule() string {
	if x != nil {
		return x.SourceModule
	}
	return ""
}

func (x *MAEvent) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *MAEvent) GetEventAttr1Type() string {
	if x != nil {
		return x.EventAttr1Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr1Value() string {
	if x != nil {
		return x.EventAttr1Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr2Type() string {
	if x != nil {
		return x.EventAttr2Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr2Value() string {
	if x != nil {
		return x.EventAttr2Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr3Type() string {
	if x != nil {
		return x.EventAttr3Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr3Value() string {
	if x != nil {
		return x.EventAttr3Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr4Type() string {
	if x != nil {
		return x.EventAttr4Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr4Value() string {
	if x != nil {
		return x.EventAttr4Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr5Type() string {
	if x != nil {
		return x.EventAttr5Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr5Value() string {
	if x != nil {
		return x.EventAttr5Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr6Type() string {
	if x != nil {
		return x.EventAttr6Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr6Value() string {
	if x != nil {
		return x.EventAttr6Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr7Type() string {
	if x != nil {
		return x.EventAttr7Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr7Value() string {
	if x != nil {
		return x.EventAttr7Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr8Type() string {
	if x != nil {
		return x.EventAttr8Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr8Value() string {
	if x != nil {
		return x.EventAttr8Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr9Type() string {
	if x != nil {
		return x.EventAttr9Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr9Value() string {
	if x != nil {
		return x.EventAttr9Value
	}
	return ""
}

func (x *MAEvent) GetEventAttr10Type() string {
	if x != nil {
		return x.EventAttr10Type
	}
	return ""
}

func (x *MAEvent) GetEventAttr10Value() string {
	if x != nil {
		return x.EventAttr10Value
	}
	return ""
}

type MACrash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *MAUser              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Step      string               `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Time      *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	PageTitle string               `protobuf:"bytes,4,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	PageUrl   string               `protobuf:"bytes,5,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	Log       string               `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *MACrash) Reset() {
	*x = MACrash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maevents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MACrash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MACrash) ProtoMessage() {}

func (x *MACrash) ProtoReflect() protoreflect.Message {
	mi := &file_maevents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MACrash.ProtoReflect.Descriptor instead.
func (*MACrash) Descriptor() ([]byte, []int) {
	return file_maevents_proto_rawDescGZIP(), []int{3}
}

func (x *MACrash) GetUser() *MAUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MACrash) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

func (x *MACrash) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MACrash) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *MACrash) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *MACrash) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_maevents_proto protoreflect.FileDescriptor

var file_maevents_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x06, 0x4d, 0x41, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x49, 0x64, 0x22, 0xb6, 0x02, 0x0a, 0x0a, 0x4d, 0x41, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4d, 0x41, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0xd6, 0x09, 0x0a, 0x07,
	0x4d, 0x41, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x41, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x31, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x31, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x31, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x31, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x32, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x32,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x33, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x33, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x33, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x33, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x34, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x34, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x34, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x35, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x35, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x35, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x35, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x36, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x36, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x36, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x36, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x37, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x37, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x37, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x37, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x38, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x38, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x38, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x24, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x38, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x39, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x39, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x39, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x39, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x31, 0x30, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x31, 0x30, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x31, 0x30, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x31, 0x30, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x07, 0x4d, 0x41, 0x43, 0x72, 0x61, 0x73, 0x68,
	0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x4d, 0x41, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x30, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maevents_proto_rawDescOnce sync.Once
	file_maevents_proto_rawDescData = file_maevents_proto_rawDesc
)

func file_maevents_proto_rawDescGZIP() []byte {
	file_maevents_proto_rawDescOnce.Do(func() {
		file_maevents_proto_rawDescData = protoimpl.X.CompressGZIP(file_maevents_proto_rawDescData)
	})
	return file_maevents_proto_rawDescData
}

var file_maevents_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_maevents_proto_goTypes = []interface{}{
	(*MAUser)(nil),              // 0: MAUser
	(*MARegister)(nil),          // 1: MARegister
	(*MAEvent)(nil),             // 2: MAEvent
	(*MACrash)(nil),             // 3: MACrash
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_maevents_proto_depIdxs = []int32{
	0, // 0: MARegister.user:type_name -> MAUser
	4, // 1: MARegister.time:type_name -> google.protobuf.Timestamp
	0, // 2: MAEvent.user:type_name -> MAUser
	4, // 3: MAEvent.time:type_name -> google.protobuf.Timestamp
	0, // 4: MACrash.user:type_name -> MAUser
	4, // 5: MACrash.time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_maevents_proto_init() }
func file_maevents_proto_init() {
	if File_maevents_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maevents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MAUser); i {
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
		file_maevents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MARegister); i {
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
		file_maevents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MAEvent); i {
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
		file_maevents_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MACrash); i {
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
			RawDescriptor: file_maevents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_maevents_proto_goTypes,
		DependencyIndexes: file_maevents_proto_depIdxs,
		MessageInfos:      file_maevents_proto_msgTypes,
	}.Build()
	File_maevents_proto = out.File
	file_maevents_proto_rawDesc = nil
	file_maevents_proto_goTypes = nil
	file_maevents_proto_depIdxs = nil
}