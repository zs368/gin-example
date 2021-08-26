// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetHttp() *HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 string               `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Timeout              *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Debug                bool                 `protobuf:"varint,4,opt,name=debug,proto3" json:"debug,omitempty"`
	PageSize             int32                `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	MaxPageSize          int32                `protobuf:"varint,6,opt,name=maxPageSize,proto3" json:"maxPageSize,omitempty"`
	UploadSavePath       string               `protobuf:"bytes,7,opt,name=uploadSavePath,proto3" json:"uploadSavePath,omitempty"`
	UploadServerUrl      string               `protobuf:"bytes,8,opt,name=uploadServerUrl,proto3" json:"uploadServerUrl,omitempty"`
	UploadImageMaxSize   int32                `protobuf:"varint,9,opt,name=uploadImageMaxSize,proto3" json:"uploadImageMaxSize,omitempty"`
	UploadImageAllowExts []string             `protobuf:"bytes,10,rep,name=uploadImageAllowExts,proto3" json:"uploadImageAllowExts,omitempty"`
	TraceName            string               `protobuf:"bytes,11,opt,name=traceName,proto3" json:"traceName,omitempty"`
	TracePort            string               `protobuf:"bytes,12,opt,name=tracePort,proto3" json:"tracePort,omitempty"`
	WsWriteWait          *durationpb.Duration `protobuf:"bytes,13,opt,name=wsWriteWait,proto3" json:"wsWriteWait,omitempty"`
	WsPongWait           *durationpb.Duration `protobuf:"bytes,14,opt,name=wsPongWait,proto3" json:"wsPongWait,omitempty"`
	WsMaxMessageSize     int64                `protobuf:"varint,15,opt,name=wsMaxMessageSize,proto3" json:"wsMaxMessageSize,omitempty"`
	WsMessageQueue       int32                `protobuf:"varint,16,opt,name=wsMessageQueue,proto3" json:"wsMessageQueue,omitempty"`
	WsOfflineNum         int32                `protobuf:"varint,17,opt,name=wsOfflineNum,proto3" json:"wsOfflineNum,omitempty"`
	JwtSecret            string               `protobuf:"bytes,18,opt,name=jwtSecret,proto3" json:"jwtSecret,omitempty"`
	JwtIssuer            string               `protobuf:"bytes,19,opt,name=jwtIssuer,proto3" json:"jwtIssuer,omitempty"`
	JwtExpire            *durationpb.Duration `protobuf:"bytes,20,opt,name=jwtExpire,proto3" json:"jwtExpire,omitempty"`
}

func (x *HTTP) Reset() {
	*x = HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTP) ProtoMessage() {}

func (x *HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTP.ProtoReflect.Descriptor instead.
func (*HTTP) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *HTTP) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HTTP) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *HTTP) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *HTTP) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *HTTP) GetMaxPageSize() int32 {
	if x != nil {
		return x.MaxPageSize
	}
	return 0
}

func (x *HTTP) GetUploadSavePath() string {
	if x != nil {
		return x.UploadSavePath
	}
	return ""
}

func (x *HTTP) GetUploadServerUrl() string {
	if x != nil {
		return x.UploadServerUrl
	}
	return ""
}

func (x *HTTP) GetUploadImageMaxSize() int32 {
	if x != nil {
		return x.UploadImageMaxSize
	}
	return 0
}

func (x *HTTP) GetUploadImageAllowExts() []string {
	if x != nil {
		return x.UploadImageAllowExts
	}
	return nil
}

func (x *HTTP) GetTraceName() string {
	if x != nil {
		return x.TraceName
	}
	return ""
}

func (x *HTTP) GetTracePort() string {
	if x != nil {
		return x.TracePort
	}
	return ""
}

func (x *HTTP) GetWsWriteWait() *durationpb.Duration {
	if x != nil {
		return x.WsWriteWait
	}
	return nil
}

func (x *HTTP) GetWsPongWait() *durationpb.Duration {
	if x != nil {
		return x.WsPongWait
	}
	return nil
}

func (x *HTTP) GetWsMaxMessageSize() int64 {
	if x != nil {
		return x.WsMaxMessageSize
	}
	return 0
}

func (x *HTTP) GetWsMessageQueue() int32 {
	if x != nil {
		return x.WsMessageQueue
	}
	return 0
}

func (x *HTTP) GetWsOfflineNum() int32 {
	if x != nil {
		return x.WsOfflineNum
	}
	return 0
}

func (x *HTTP) GetJwtSecret() string {
	if x != nil {
		return x.JwtSecret
	}
	return ""
}

func (x *HTTP) GetJwtIssuer() string {
	if x != nil {
		return x.JwtIssuer
	}
	return ""
}

func (x *HTTP) GetJwtExpire() *durationpb.Duration {
	if x != nil {
		return x.JwtExpire
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db           int32                `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

var File_internal_conf_conf_proto protoreflect.FileDescriptor

var file_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2f,
	0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12,
	0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x8e, 0x06, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d,
	0x61, 0x78, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x12,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x14,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x77, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x57, 0x61, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x77, 0x73,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x57, 0x61, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x77, 0x73, 0x50,
	0x6f, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x77, 0x73, 0x50, 0x6f, 0x6e, 0x67,
	0x57, 0x61, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x77, 0x73, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x77, 0x73, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x77, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x77, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x73, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x77, 0x73, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x77,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a,
	0x77, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x6a, 0x77, 0x74, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x22, 0xdd, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x1a, 0x9d, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63,
	0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_conf_conf_proto_rawDescOnce sync.Once
	file_internal_conf_conf_proto_rawDescData = file_internal_conf_conf_proto_rawDesc
)

func file_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_conf_proto_rawDescData)
	})
	return file_internal_conf_conf_proto_rawDescData
}

var file_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: example.internal.conf.Bootstrap
	(*HTTP)(nil),                // 1: example.internal.conf.HTTP
	(*Data)(nil),                // 2: example.internal.conf.Data
	(*Data_Database)(nil),       // 3: example.internal.conf.Data.Database
	(*Data_Redis)(nil),          // 4: example.internal.conf.Data.Redis
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
}
var file_internal_conf_conf_proto_depIdxs = []int32{
	1,  // 0: example.internal.conf.Bootstrap.http:type_name -> example.internal.conf.HTTP
	2,  // 1: example.internal.conf.Bootstrap.data:type_name -> example.internal.conf.Data
	5,  // 2: example.internal.conf.HTTP.timeout:type_name -> google.protobuf.Duration
	5,  // 3: example.internal.conf.HTTP.wsWriteWait:type_name -> google.protobuf.Duration
	5,  // 4: example.internal.conf.HTTP.wsPongWait:type_name -> google.protobuf.Duration
	5,  // 5: example.internal.conf.HTTP.jwtExpire:type_name -> google.protobuf.Duration
	3,  // 6: example.internal.conf.Data.database:type_name -> example.internal.conf.Data.Database
	4,  // 7: example.internal.conf.Data.redis:type_name -> example.internal.conf.Data.Redis
	5,  // 8: example.internal.conf.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	5,  // 9: example.internal.conf.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	5,  // 10: example.internal.conf.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_internal_conf_conf_proto_init() }
func file_internal_conf_conf_proto_init() {
	if File_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTP); i {
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
		file_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
			RawDescriptor: file_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_internal_conf_conf_proto = out.File
	file_internal_conf_conf_proto_rawDesc = nil
	file_internal_conf_conf_proto_goTypes = nil
	file_internal_conf_conf_proto_depIdxs = nil
}
