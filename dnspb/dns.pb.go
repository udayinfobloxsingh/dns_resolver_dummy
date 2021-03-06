// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: Proto/dns.proto

package dnspb

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

type DomainRes_Action int32

const (
	DomainRes_DROP  DomainRes_Action = 0
	DomainRes_BLOCK DomainRes_Action = 1
	DomainRes_ALLOW DomainRes_Action = 2
)

// Enum value maps for DomainRes_Action.
var (
	DomainRes_Action_name = map[int32]string{
		0: "DROP",
		1: "BLOCK",
		2: "ALLOW",
	}
	DomainRes_Action_value = map[string]int32{
		"DROP":  0,
		"BLOCK": 1,
		"ALLOW": 2,
	}
)

func (x DomainRes_Action) Enum() *DomainRes_Action {
	p := new(DomainRes_Action)
	*p = x
	return p
}

func (x DomainRes_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainRes_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_Proto_dns_proto_enumTypes[0].Descriptor()
}

func (DomainRes_Action) Type() protoreflect.EnumType {
	return &file_Proto_dns_proto_enumTypes[0]
}

func (x DomainRes_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainRes_Action.Descriptor instead.
func (DomainRes_Action) EnumDescriptor() ([]byte, []int) {
	return file_Proto_dns_proto_rawDescGZIP(), []int{1, 0}
}

type DomainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domainName,proto3" json:"domainName,omitempty"`
}

func (x *DomainReq) Reset() {
	*x = DomainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainReq) ProtoMessage() {}

func (x *DomainReq) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainReq.ProtoReflect.Descriptor instead.
func (*DomainReq) Descriptor() ([]byte, []int) {
	return file_Proto_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DomainReq) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

type DomainRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action DomainRes_Action `protobuf:"varint,1,opt,name=action,proto3,enum=dns.DomainRes_Action" json:"action,omitempty"`
}

func (x *DomainRes) Reset() {
	*x = DomainRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainRes) ProtoMessage() {}

func (x *DomainRes) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainRes.ProtoReflect.Descriptor instead.
func (*DomainRes) Descriptor() ([]byte, []int) {
	return file_Proto_dns_proto_rawDescGZIP(), []int{1}
}

func (x *DomainRes) GetAction() DomainRes_Action {
	if x != nil {
		return x.Action
	}
	return DomainRes_DROP
}

var File_Proto_dns_proto protoreflect.FileDescriptor

var file_Proto_dns_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x64, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x09, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x09, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x2d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x28, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f,
	0x50, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x32, 0x3f, 0x0a, 0x0a, 0x44, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0f, 0x44, 0x4e, 0x53, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x64, 0x6e, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_dns_proto_rawDescOnce sync.Once
	file_Proto_dns_proto_rawDescData = file_Proto_dns_proto_rawDesc
)

func file_Proto_dns_proto_rawDescGZIP() []byte {
	file_Proto_dns_proto_rawDescOnce.Do(func() {
		file_Proto_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_dns_proto_rawDescData)
	})
	return file_Proto_dns_proto_rawDescData
}

var file_Proto_dns_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Proto_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Proto_dns_proto_goTypes = []interface{}{
	(DomainRes_Action)(0), // 0: dns.DomainRes.Action
	(*DomainReq)(nil),     // 1: dns.DomainReq
	(*DomainRes)(nil),     // 2: dns.DomainRes
}
var file_Proto_dns_proto_depIdxs = []int32{
	0, // 0: dns.DomainRes.action:type_name -> dns.DomainRes.Action
	1, // 1: dns.DnsService.DNSauthenticate:input_type -> dns.DomainReq
	2, // 2: dns.DnsService.DNSauthenticate:output_type -> dns.DomainRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Proto_dns_proto_init() }
func file_Proto_dns_proto_init() {
	if File_Proto_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainReq); i {
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
		file_Proto_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainRes); i {
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
			RawDescriptor: file_Proto_dns_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Proto_dns_proto_goTypes,
		DependencyIndexes: file_Proto_dns_proto_depIdxs,
		EnumInfos:         file_Proto_dns_proto_enumTypes,
		MessageInfos:      file_Proto_dns_proto_msgTypes,
	}.Build()
	File_Proto_dns_proto = out.File
	file_Proto_dns_proto_rawDesc = nil
	file_Proto_dns_proto_goTypes = nil
	file_Proto_dns_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DnsServiceClient is the client API for DnsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DnsServiceClient interface {
	DNSauthenticate(ctx context.Context, in *DomainReq, opts ...grpc.CallOption) (*DomainRes, error)
}

type dnsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsServiceClient(cc grpc.ClientConnInterface) DnsServiceClient {
	return &dnsServiceClient{cc}
}

func (c *dnsServiceClient) DNSauthenticate(ctx context.Context, in *DomainReq, opts ...grpc.CallOption) (*DomainRes, error) {
	out := new(DomainRes)
	err := c.cc.Invoke(ctx, "/dns.DnsService/DNSauthenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsServiceServer is the server API for DnsService service.
type DnsServiceServer interface {
	DNSauthenticate(context.Context, *DomainReq) (*DomainRes, error)
}

// UnimplementedDnsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDnsServiceServer struct {
}

func (*UnimplementedDnsServiceServer) DNSauthenticate(context.Context, *DomainReq) (*DomainRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DNSauthenticate not implemented")
}

func RegisterDnsServiceServer(s *grpc.Server, srv DnsServiceServer) {
	s.RegisterService(&_DnsService_serviceDesc, srv)
}

func _DnsService_DNSauthenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).DNSauthenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dns.DnsService/DNSauthenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).DNSauthenticate(ctx, req.(*DomainReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DnsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dns.DnsService",
	HandlerType: (*DnsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DNSauthenticate",
			Handler:    _DnsService_DNSauthenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/dns.proto",
}
