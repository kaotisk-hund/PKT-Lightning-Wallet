// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: restrpc/rest.proto

package restrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RestCommandCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description []string                        `protobuf:"bytes,1,rep,name=description,proto3" json:"description,omitempty"`
	Endpoints   map[string]string               `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subcategory map[string]*RestCommandCategory `protobuf:"bytes,3,rep,name=subcategory,proto3" json:"subcategory,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RestCommandCategory) Reset() {
	*x = RestCommandCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestCommandCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestCommandCategory) ProtoMessage() {}

func (x *RestCommandCategory) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestCommandCategory.ProtoReflect.Descriptor instead.
func (*RestCommandCategory) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{0}
}

func (x *RestCommandCategory) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *RestCommandCategory) GetEndpoints() map[string]string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *RestCommandCategory) GetSubcategory() map[string]*RestCommandCategory {
	if x != nil {
		return x.Subcategory
	}
	return nil
}

type RestMasterHelpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description []string                        `protobuf:"bytes,2,rep,name=description,proto3" json:"description,omitempty"`
	Category    map[string]*RestCommandCategory `protobuf:"bytes,3,rep,name=category,proto3" json:"category,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RestMasterHelpResponse) Reset() {
	*x = RestMasterHelpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestMasterHelpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestMasterHelpResponse) ProtoMessage() {}

func (x *RestMasterHelpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestMasterHelpResponse.ProtoReflect.Descriptor instead.
func (*RestMasterHelpResponse) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{1}
}

func (x *RestMasterHelpResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestMasterHelpResponse) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *RestMasterHelpResponse) GetCategory() map[string]*RestCommandCategory {
	if x != nil {
		return x.Category
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description []string `protobuf:"bytes,2,rep,name=description,proto3" json:"description,omitempty"`
	Repeated    bool     `protobuf:"varint,3,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Type        *Type    `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Field) GetRepeated() bool {
	if x != nil {
		return x.Repeated
	}
	return false
}

func (x *Field) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description []string `protobuf:"bytes,2,rep,name=description,proto3" json:"description,omitempty"`
	Fields      []*Field `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{3}
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Type) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type RestHelpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Service     string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Description []string `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty"`
	Request     *Type    `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	Response    *Type    `protobuf:"bytes,5,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RestHelpResponse) Reset() {
	*x = RestHelpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestHelpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestHelpResponse) ProtoMessage() {}

func (x *RestHelpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestHelpResponse.ProtoReflect.Descriptor instead.
func (*RestHelpResponse) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{4}
}

func (x *RestHelpResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestHelpResponse) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RestHelpResponse) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *RestHelpResponse) GetRequest() *Type {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RestHelpResponse) GetResponse() *Type {
	if x != nil {
		return x.Response
	}
	return nil
}

//  WebSocket request and response JSon messages
type WebSocketError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpCode uint32 `protobuf:"varint,1,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WebSocketError) Reset() {
	*x = WebSocketError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketError) ProtoMessage() {}

func (x *WebSocketError) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketError.ProtoReflect.Descriptor instead.
func (*WebSocketError) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{5}
}

func (x *WebSocketError) GetHttpCode() uint32 {
	if x != nil {
		return x.HttpCode
	}
	return 0
}

func (x *WebSocketError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//  WebSocket request and response protobuf messages
type WebSocketProtobufRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rest endpoint to send the request to
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// An arbitrary string which will be reflected back in the response
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	HasMore   bool   `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	// The data to post to the REST endpoint, if any.
	// Make sure this is the correct data structure based on the endpoint you are posting to.
	Payload *anypb.Any `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *WebSocketProtobufRequest) Reset() {
	*x = WebSocketProtobufRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketProtobufRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketProtobufRequest) ProtoMessage() {}

func (x *WebSocketProtobufRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketProtobufRequest.ProtoReflect.Descriptor instead.
func (*WebSocketProtobufRequest) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{6}
}

func (x *WebSocketProtobufRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *WebSocketProtobufRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *WebSocketProtobufRequest) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *WebSocketProtobufRequest) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type WebSocketJSonProtobufResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	HasMore   bool   `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	// Types that are assignable to Payload:
	//	*WebSocketJSonProtobufResponse_Ok
	//	*WebSocketJSonProtobufResponse_Error
	Payload isWebSocketJSonProtobufResponse_Payload `protobuf_oneof:"payload"`
}

func (x *WebSocketJSonProtobufResponse) Reset() {
	*x = WebSocketJSonProtobufResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restrpc_rest_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketJSonProtobufResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketJSonProtobufResponse) ProtoMessage() {}

func (x *WebSocketJSonProtobufResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restrpc_rest_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketJSonProtobufResponse.ProtoReflect.Descriptor instead.
func (*WebSocketJSonProtobufResponse) Descriptor() ([]byte, []int) {
	return file_restrpc_rest_proto_rawDescGZIP(), []int{7}
}

func (x *WebSocketJSonProtobufResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *WebSocketJSonProtobufResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (m *WebSocketJSonProtobufResponse) GetPayload() isWebSocketJSonProtobufResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *WebSocketJSonProtobufResponse) GetOk() *anypb.Any {
	if x, ok := x.GetPayload().(*WebSocketJSonProtobufResponse_Ok); ok {
		return x.Ok
	}
	return nil
}

func (x *WebSocketJSonProtobufResponse) GetError() *WebSocketError {
	if x, ok := x.GetPayload().(*WebSocketJSonProtobufResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isWebSocketJSonProtobufResponse_Payload interface {
	isWebSocketJSonProtobufResponse_Payload()
}

type WebSocketJSonProtobufResponse_Ok struct {
	Ok *anypb.Any `protobuf:"bytes,3,opt,name=ok,proto3,oneof"`
}

type WebSocketJSonProtobufResponse_Error struct {
	Error *WebSocketError `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*WebSocketJSonProtobufResponse_Ok) isWebSocketJSonProtobufResponse_Payload() {}

func (*WebSocketJSonProtobufResponse_Error) isWebSocketJSonProtobufResponse_Payload() {}

var File_restrpc_rest_proto protoreflect.FileDescriptor

var file_restrpc_rest_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a,
	0x0b, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x3c,
	0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf4, 0x01, 0x0a, 0x16, 0x52,
	0x65, 0x73, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x59, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x7c, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x64, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x6c, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47,
	0x0a, 0x0e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x18, 0x57, 0x65, 0x62, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x1d, 0x57,
	0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4a, 0x53, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x2f,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x74, 0x2d, 0x63, 0x61, 0x73,
	0x68, 0x2f, 0x70, 0x6b, 0x74, 0x64, 0x2f, 0x6c, 0x6e, 0x64, 0x2f, 0x6c, 0x6e, 0x72, 0x70, 0x63,
	0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restrpc_rest_proto_rawDescOnce sync.Once
	file_restrpc_rest_proto_rawDescData = file_restrpc_rest_proto_rawDesc
)

func file_restrpc_rest_proto_rawDescGZIP() []byte {
	file_restrpc_rest_proto_rawDescOnce.Do(func() {
		file_restrpc_rest_proto_rawDescData = protoimpl.X.CompressGZIP(file_restrpc_rest_proto_rawDescData)
	})
	return file_restrpc_rest_proto_rawDescData
}

var file_restrpc_rest_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_restrpc_rest_proto_goTypes = []interface{}{
	(*RestCommandCategory)(nil),           // 0: restrpc.RestCommandCategory
	(*RestMasterHelpResponse)(nil),        // 1: restrpc.RestMasterHelpResponse
	(*Field)(nil),                         // 2: restrpc.Field
	(*Type)(nil),                          // 3: restrpc.Type
	(*RestHelpResponse)(nil),              // 4: restrpc.RestHelpResponse
	(*WebSocketError)(nil),                // 5: restrpc.WebSocketError
	(*WebSocketProtobufRequest)(nil),      // 6: restrpc.WebSocketProtobufRequest
	(*WebSocketJSonProtobufResponse)(nil), // 7: restrpc.WebSocketJSonProtobufResponse
	nil,                                   // 8: restrpc.RestCommandCategory.EndpointsEntry
	nil,                                   // 9: restrpc.RestCommandCategory.SubcategoryEntry
	nil,                                   // 10: restrpc.RestMasterHelpResponse.CategoryEntry
	(*anypb.Any)(nil),                     // 11: google.protobuf.Any
}
var file_restrpc_rest_proto_depIdxs = []int32{
	8,  // 0: restrpc.RestCommandCategory.endpoints:type_name -> restrpc.RestCommandCategory.EndpointsEntry
	9,  // 1: restrpc.RestCommandCategory.subcategory:type_name -> restrpc.RestCommandCategory.SubcategoryEntry
	10, // 2: restrpc.RestMasterHelpResponse.category:type_name -> restrpc.RestMasterHelpResponse.CategoryEntry
	3,  // 3: restrpc.Field.type:type_name -> restrpc.Type
	2,  // 4: restrpc.Type.fields:type_name -> restrpc.Field
	3,  // 5: restrpc.RestHelpResponse.request:type_name -> restrpc.Type
	3,  // 6: restrpc.RestHelpResponse.response:type_name -> restrpc.Type
	11, // 7: restrpc.WebSocketProtobufRequest.payload:type_name -> google.protobuf.Any
	11, // 8: restrpc.WebSocketJSonProtobufResponse.ok:type_name -> google.protobuf.Any
	5,  // 9: restrpc.WebSocketJSonProtobufResponse.error:type_name -> restrpc.WebSocketError
	0,  // 10: restrpc.RestCommandCategory.SubcategoryEntry.value:type_name -> restrpc.RestCommandCategory
	0,  // 11: restrpc.RestMasterHelpResponse.CategoryEntry.value:type_name -> restrpc.RestCommandCategory
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_restrpc_rest_proto_init() }
func file_restrpc_rest_proto_init() {
	if File_restrpc_rest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restrpc_rest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestCommandCategory); i {
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
		file_restrpc_rest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestMasterHelpResponse); i {
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
		file_restrpc_rest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_restrpc_rest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_restrpc_rest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestHelpResponse); i {
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
		file_restrpc_rest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketError); i {
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
		file_restrpc_rest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketProtobufRequest); i {
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
		file_restrpc_rest_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketJSonProtobufResponse); i {
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
	file_restrpc_rest_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*WebSocketJSonProtobufResponse_Ok)(nil),
		(*WebSocketJSonProtobufResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_restrpc_rest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restrpc_rest_proto_goTypes,
		DependencyIndexes: file_restrpc_rest_proto_depIdxs,
		MessageInfos:      file_restrpc_rest_proto_msgTypes,
	}.Build()
	File_restrpc_rest_proto = out.File
	file_restrpc_rest_proto_rawDesc = nil
	file_restrpc_rest_proto_goTypes = nil
	file_restrpc_rest_proto_depIdxs = nil
}
