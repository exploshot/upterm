// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Identifier_Type int32

const (
	Identifier_HOST   Identifier_Type = 0
	Identifier_CLIENT Identifier_Type = 1
)

// Enum value maps for Identifier_Type.
var (
	Identifier_Type_name = map[int32]string{
		0: "HOST",
		1: "CLIENT",
	}
	Identifier_Type_value = map[string]int32{
		"HOST":   0,
		"CLIENT": 1,
	}
)

func (x Identifier_Type) Enum() *Identifier_Type {
	p := new(Identifier_Type)
	*p = x
	return p
}

func (x Identifier_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Identifier_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Identifier_Type) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Identifier_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Identifier_Type.Descriptor instead.
func (Identifier_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3, 0}
}

type GetSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSessionRequest) Reset() {
	*x = GetSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionRequest) ProtoMessage() {}

func (x *GetSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionRequest.ProtoReflect.Descriptor instead.
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type GetSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId        string    `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Command          []string  `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	ForceCommand     []string  `protobuf:"bytes,3,rep,name=force_command,json=forceCommand,proto3" json:"force_command,omitempty"`
	Host             string    `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	NodeAddr         string    `protobuf:"bytes,5,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
	ConnectedClients []*Client `protobuf:"bytes,6,rep,name=connected_clients,json=connectedClients,proto3" json:"connected_clients,omitempty"`
}

func (x *GetSessionResponse) Reset() {
	*x = GetSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionResponse) ProtoMessage() {}

func (x *GetSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionResponse.ProtoReflect.Descriptor instead.
func (*GetSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetSessionResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *GetSessionResponse) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *GetSessionResponse) GetForceCommand() []string {
	if x != nil {
		return x.ForceCommand
	}
	return nil
}

func (x *GetSessionResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GetSessionResponse) GetNodeAddr() string {
	if x != nil {
		return x.NodeAddr
	}
	return ""
}

func (x *GetSessionResponse) GetConnectedClients() []*Client {
	if x != nil {
		return x.ConnectedClients
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Addr                 string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	PublicKeyFingerprint string `protobuf:"bytes,4,opt,name=public_key_fingerprint,json=publicKeyFingerprint,proto3" json:"public_key_fingerprint,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Client) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Client) GetPublicKeyFingerprint() string {
	if x != nil {
		return x.PublicKeyFingerprint
	}
	return ""
}

type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     Identifier_Type `protobuf:"varint,2,opt,name=type,proto3,enum=api.Identifier_Type" json:"type,omitempty"`
	NodeAddr string          `protobuf:"bytes,3,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Identifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Identifier) GetType() Identifier_Type {
	if x != nil {
		return x.Type
	}
	return Identifier_HOST
}

func (x *Identifier) GetNodeAddr() string {
	if x != nil {
		return x.NodeAddr
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x38, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7c, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x34, 0x0a,
	0x16, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x32, 0x4f, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_goTypes = []interface{}{
	(Identifier_Type)(0),       // 0: api.Identifier.Type
	(*GetSessionRequest)(nil),  // 1: api.GetSessionRequest
	(*GetSessionResponse)(nil), // 2: api.GetSessionResponse
	(*Client)(nil),             // 3: api.Client
	(*Identifier)(nil),         // 4: api.Identifier
}
var file_api_proto_depIdxs = []int32{
	3, // 0: api.GetSessionResponse.connected_clients:type_name -> api.Client
	0, // 1: api.Identifier.type:type_name -> api.Identifier.Type
	1, // 2: api.AdminService.GetSession:input_type -> api.GetSessionRequest
	2, // 3: api.AdminService.GetSession:output_type -> api.GetSessionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _AdminService_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
