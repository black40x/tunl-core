// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: tunl.proto

package commands

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

type ClientConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ClientConnect) Reset() {
	*x = ClientConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnect) ProtoMessage() {}

func (x *ClientConnect) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnect.ProtoReflect.Descriptor instead.
func (*ClientConnect) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConnect) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ClientConnect) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ServerHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Private bool   `protobuf:"varint,2,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *ServerHeader) Reset() {
	*x = ServerHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerHeader) ProtoMessage() {}

func (x *ServerHeader) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerHeader.ProtoReflect.Descriptor instead.
func (*ServerHeader) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{1}
}

func (x *ServerHeader) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerHeader) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type ServerConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix    string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PublicUrl string `protobuf:"bytes,2,opt,name=public_url,json=publicUrl,proto3" json:"public_url,omitempty"`
	Expire    int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *ServerConnect) Reset() {
	*x = ServerConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConnect) ProtoMessage() {}

func (x *ServerConnect) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConnect.ProtoReflect.Descriptor instead.
func (*ServerConnect) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{2}
}

func (x *ServerConnect) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ServerConnect) GetPublicUrl() string {
	if x != nil {
		return x.PublicUrl
	}
	return ""
}

func (x *ServerConnect) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{4}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type BodyChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Eof  bool   `protobuf:"varint,3,opt,name=eof,proto3" json:"eof,omitempty"`
}

func (x *BodyChunk) Reset() {
	*x = BodyChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyChunk) ProtoMessage() {}

func (x *BodyChunk) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyChunk.ProtoReflect.Descriptor instead.
func (*BodyChunk) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{5}
}

func (x *BodyChunk) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BodyChunk) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BodyChunk) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Domain   string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Expires  int64  `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	MaxAge   int32  `protobuf:"varint,6,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	Secure   bool   `protobuf:"varint,7,opt,name=secure,proto3" json:"secure,omitempty"`
	HttpOnly bool   `protobuf:"varint,8,opt,name=http_only,json=httpOnly,proto3" json:"http_only,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{6}
}

func (x *Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookie) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Cookie) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Cookie) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *Cookie) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Cookie) GetSecure() bool {
	if x != nil {
		return x.Secure
	}
	return false
}

func (x *Cookie) GetHttpOnly() bool {
	if x != nil {
		return x.HttpOnly
	}
	return false
}

type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Proto         string    `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Method        string    `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Uri           string    `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	ContentLength int64     `protobuf:"varint,5,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	Cookies       []*Cookie `protobuf:"bytes,6,rep,name=cookies,proto3" json:"cookies,omitempty"`
	Header        []*Header `protobuf:"bytes,7,rep,name=header,proto3" json:"header,omitempty"`
	RemoteAddr    string    `protobuf:"bytes,8,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	ErrorCode     int64     `protobuf:"varint,9,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{7}
}

func (x *HttpRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *HttpRequest) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *HttpRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HttpRequest) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *HttpRequest) GetCookies() []*Cookie {
	if x != nil {
		return x.Cookies
	}
	return nil
}

func (x *HttpRequest) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HttpRequest) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *HttpRequest) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type HttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Proto         string    `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Status        int32     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	ContentLength int64     `protobuf:"varint,4,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	Header        []*Header `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty"`
	ErrorCode     int64     `protobuf:"varint,6,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *HttpResponse) Reset() {
	*x = HttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpResponse) ProtoMessage() {}

func (x *HttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpResponse.ProtoReflect.Descriptor instead.
func (*HttpResponse) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{8}
}

func (x *HttpResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *HttpResponse) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *HttpResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HttpResponse) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *HttpResponse) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HttpResponse) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*Transfer_ServerHeader
	//	*Transfer_ClientConnect
	//	*Transfer_ServerConnect
	//	*Transfer_HttpRequest
	//	*Transfer_HttpResponse
	//	*Transfer_BodyChunk
	//	*Transfer_Error
	Command isTransfer_Command `protobuf_oneof:"Command"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunl_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_tunl_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_tunl_proto_rawDescGZIP(), []int{9}
}

func (m *Transfer) GetCommand() isTransfer_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Transfer) GetServerHeader() *ServerHeader {
	if x, ok := x.GetCommand().(*Transfer_ServerHeader); ok {
		return x.ServerHeader
	}
	return nil
}

func (x *Transfer) GetClientConnect() *ClientConnect {
	if x, ok := x.GetCommand().(*Transfer_ClientConnect); ok {
		return x.ClientConnect
	}
	return nil
}

func (x *Transfer) GetServerConnect() *ServerConnect {
	if x, ok := x.GetCommand().(*Transfer_ServerConnect); ok {
		return x.ServerConnect
	}
	return nil
}

func (x *Transfer) GetHttpRequest() *HttpRequest {
	if x, ok := x.GetCommand().(*Transfer_HttpRequest); ok {
		return x.HttpRequest
	}
	return nil
}

func (x *Transfer) GetHttpResponse() *HttpResponse {
	if x, ok := x.GetCommand().(*Transfer_HttpResponse); ok {
		return x.HttpResponse
	}
	return nil
}

func (x *Transfer) GetBodyChunk() *BodyChunk {
	if x, ok := x.GetCommand().(*Transfer_BodyChunk); ok {
		return x.BodyChunk
	}
	return nil
}

func (x *Transfer) GetError() *Error {
	if x, ok := x.GetCommand().(*Transfer_Error); ok {
		return x.Error
	}
	return nil
}

type isTransfer_Command interface {
	isTransfer_Command()
}

type Transfer_ServerHeader struct {
	ServerHeader *ServerHeader `protobuf:"bytes,1,opt,name=server_header,json=serverHeader,proto3,oneof"`
}

type Transfer_ClientConnect struct {
	ClientConnect *ClientConnect `protobuf:"bytes,2,opt,name=client_connect,json=clientConnect,proto3,oneof"`
}

type Transfer_ServerConnect struct {
	ServerConnect *ServerConnect `protobuf:"bytes,3,opt,name=server_connect,json=serverConnect,proto3,oneof"`
}

type Transfer_HttpRequest struct {
	HttpRequest *HttpRequest `protobuf:"bytes,4,opt,name=http_request,json=httpRequest,proto3,oneof"`
}

type Transfer_HttpResponse struct {
	HttpResponse *HttpResponse `protobuf:"bytes,5,opt,name=http_response,json=httpResponse,proto3,oneof"`
}

type Transfer_BodyChunk struct {
	BodyChunk *BodyChunk `protobuf:"bytes,6,opt,name=body_chunk,json=bodyChunk,proto3,oneof"`
}

type Transfer_Error struct {
	Error *Error `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*Transfer_ServerHeader) isTransfer_Command() {}

func (*Transfer_ClientConnect) isTransfer_Command() {}

func (*Transfer_ServerConnect) isTransfer_Command() {}

func (*Transfer_HttpRequest) isTransfer_Command() {}

func (*Transfer_HttpResponse) isTransfer_Command() {}

func (*Transfer_BodyChunk) isTransfer_Command() {}

func (*Transfer_Error) isTransfer_Command() {}

var File_tunl_proto protoreflect.FileDescriptor

var file_tunl_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x75, 0x6e, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x5e,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x35,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x09, 0x42, 0x6f, 0x64, 0x79, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x22, 0xc6,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78,
	0x5f, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68,
	0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x98, 0x02, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x9d, 0x03, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x3a, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tunl_proto_rawDescOnce sync.Once
	file_tunl_proto_rawDescData = file_tunl_proto_rawDesc
)

func file_tunl_proto_rawDescGZIP() []byte {
	file_tunl_proto_rawDescOnce.Do(func() {
		file_tunl_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunl_proto_rawDescData)
	})
	return file_tunl_proto_rawDescData
}

var file_tunl_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tunl_proto_goTypes = []interface{}{
	(*ClientConnect)(nil), // 0: proto.ClientConnect
	(*ServerHeader)(nil),  // 1: proto.ServerHeader
	(*ServerConnect)(nil), // 2: proto.ServerConnect
	(*Error)(nil),         // 3: proto.Error
	(*Header)(nil),        // 4: proto.Header
	(*BodyChunk)(nil),     // 5: proto.BodyChunk
	(*Cookie)(nil),        // 6: proto.Cookie
	(*HttpRequest)(nil),   // 7: proto.HttpRequest
	(*HttpResponse)(nil),  // 8: proto.HttpResponse
	(*Transfer)(nil),      // 9: proto.Transfer
}
var file_tunl_proto_depIdxs = []int32{
	6,  // 0: proto.HttpRequest.cookies:type_name -> proto.Cookie
	4,  // 1: proto.HttpRequest.header:type_name -> proto.Header
	4,  // 2: proto.HttpResponse.header:type_name -> proto.Header
	1,  // 3: proto.Transfer.server_header:type_name -> proto.ServerHeader
	0,  // 4: proto.Transfer.client_connect:type_name -> proto.ClientConnect
	2,  // 5: proto.Transfer.server_connect:type_name -> proto.ServerConnect
	7,  // 6: proto.Transfer.http_request:type_name -> proto.HttpRequest
	8,  // 7: proto.Transfer.http_response:type_name -> proto.HttpResponse
	5,  // 8: proto.Transfer.body_chunk:type_name -> proto.BodyChunk
	3,  // 9: proto.Transfer.error:type_name -> proto.Error
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_tunl_proto_init() }
func file_tunl_proto_init() {
	if File_tunl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tunl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConnect); i {
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
		file_tunl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerHeader); i {
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
		file_tunl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConnect); i {
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
		file_tunl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_tunl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_tunl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyChunk); i {
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
		file_tunl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
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
		file_tunl_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
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
		file_tunl_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpResponse); i {
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
		file_tunl_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
	file_tunl_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Transfer_ServerHeader)(nil),
		(*Transfer_ClientConnect)(nil),
		(*Transfer_ServerConnect)(nil),
		(*Transfer_HttpRequest)(nil),
		(*Transfer_HttpResponse)(nil),
		(*Transfer_BodyChunk)(nil),
		(*Transfer_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tunl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tunl_proto_goTypes,
		DependencyIndexes: file_tunl_proto_depIdxs,
		MessageInfos:      file_tunl_proto_msgTypes,
	}.Build()
	File_tunl_proto = out.File
	file_tunl_proto_rawDesc = nil
	file_tunl_proto_goTypes = nil
	file_tunl_proto_depIdxs = nil
}
