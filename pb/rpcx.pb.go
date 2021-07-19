// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.3
// source: rpcx.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requester string            `protobuf:"bytes,1,opt,name=requester,proto3" json:"requester,omitempty"`                                                                                        // 调用者对自身的描述，必须，由调用者填写。如:ui#192.168.10.2:8211
	TraceId   string            `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`                                                                             // 请求的唯一ID，必须，由调用者填写
	Timestamp int64             `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                                                                       // 事件发生的时间，必须，由调用者填写
	Version   string            `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`                                                                                            // 传输协议版本号，必须，由调用者填写
	Operator  string            `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`                                                                                          // 操作者ID，如web操作的操作用户，由调用者填写
	ReplyCode int32             `protobuf:"varint,6,opt,name=reply_code,json=replyCode,proto3" json:"reply_code,omitempty"`                                                                      // 服务端应答码，由各业务自己定义枚举
	ReplyMsg  string            `protobuf:"bytes,7,opt,name=reply_msg,json=replyMsg,proto3" json:"reply_msg,omitempty"`                                                                          // 服务端应错误，由各业务自己定义
	Metadata  map[string]string `protobuf:"bytes,10,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 交互数据
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_rpcx_proto_msgTypes[0]
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
	return file_rpcx_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetRequester() string {
	if x != nil {
		return x.Requester
	}
	return ""
}

func (x *Header) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Header) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Header) GetReplyCode() int32 {
	if x != nil {
		return x.ReplyCode
	}
	return 0
}

func (x *Header) GetReplyMsg() string {
	if x != nil {
		return x.ReplyMsg
	}
	return ""
}

func (x *Header) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_rpcx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_rpcx_proto_rawDescGZIP(), []int{1}
}

type Ping_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *Ping_Req) Reset() {
	*x = Ping_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping_Req) ProtoMessage() {}

func (x *Ping_Req) ProtoReflect() protoreflect.Message {
	mi := &file_rpcx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping_Req.ProtoReflect.Descriptor instead.
func (*Ping_Req) Descriptor() ([]byte, []int) {
	return file_rpcx_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Ping_Req) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

type Ping_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *Ping_Resp) Reset() {
	*x = Ping_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping_Resp) ProtoMessage() {}

func (x *Ping_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_rpcx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping_Resp.ProtoReflect.Descriptor instead.
func (*Ping_Resp) Descriptor() ([]byte, []int) {
	return file_rpcx_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Ping_Resp) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_rpcx_proto protoreflect.FileDescriptor

var file_rpcx_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x70, 0x63, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xc4, 0x02, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x1a,
	0x29, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x2a, 0x0a, 0x04, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcx_proto_rawDescOnce sync.Once
	file_rpcx_proto_rawDescData = file_rpcx_proto_rawDesc
)

func file_rpcx_proto_rawDescGZIP() []byte {
	file_rpcx_proto_rawDescOnce.Do(func() {
		file_rpcx_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcx_proto_rawDescData)
	})
	return file_rpcx_proto_rawDescData
}

var file_rpcx_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpcx_proto_goTypes = []interface{}{
	(*Header)(nil),    // 0: pb.Header
	(*Ping)(nil),      // 1: pb.Ping
	nil,               // 2: pb.Header.MetadataEntry
	(*Ping_Req)(nil),  // 3: pb.Ping.Req
	(*Ping_Resp)(nil), // 4: pb.Ping.Resp
}
var file_rpcx_proto_depIdxs = []int32{
	2, // 0: pb.Header.metadata:type_name -> pb.Header.MetadataEntry
	0, // 1: pb.Ping.Req.header:type_name -> pb.Header
	0, // 2: pb.Ping.Resp.header:type_name -> pb.Header
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpcx_proto_init() }
func file_rpcx_proto_init() {
	if File_rpcx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpcx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_rpcx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping_Req); i {
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
		file_rpcx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping_Resp); i {
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
			RawDescriptor: file_rpcx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpcx_proto_goTypes,
		DependencyIndexes: file_rpcx_proto_depIdxs,
		MessageInfos:      file_rpcx_proto_msgTypes,
	}.Build()
	File_rpcx_proto = out.File
	file_rpcx_proto_rawDesc = nil
	file_rpcx_proto_goTypes = nil
	file_rpcx_proto_depIdxs = nil
}
