// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/chronicle.proto

package debtor_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Chronicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Debtors []*Debtor              `protobuf:"bytes,2,rep,name=debtors,proto3" json:"debtors,omitempty"`
}

func (x *Chronicle) Reset() {
	*x = Chronicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chronicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chronicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chronicle) ProtoMessage() {}

func (x *Chronicle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chronicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chronicle.ProtoReflect.Descriptor instead.
func (*Chronicle) Descriptor() ([]byte, []int) {
	return file_proto_chronicle_proto_rawDescGZIP(), []int{0}
}

func (x *Chronicle) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Chronicle) GetDebtors() []*Debtor {
	if x != nil {
		return x.Debtors
	}
	return nil
}

type Debtor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint32 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (x *Debtor) Reset() {
	*x = Debtor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chronicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Debtor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Debtor) ProtoMessage() {}

func (x *Debtor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chronicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Debtor.ProtoReflect.Descriptor instead.
func (*Debtor) Descriptor() ([]byte, []int) {
	return file_proto_chronicle_proto_rawDescGZIP(), []int{1}
}

func (x *Debtor) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Debtor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Debtor) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_proto_chronicle_proto protoreflect.FileDescriptor

var file_proto_chronicle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63,
	0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x09, 0x43, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x62, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x48, 0x0a,
	0x06, 0x44, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x7f, 0x0a, 0x10, 0x43, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x41,
	0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65,
	0x1a, 0x14, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6b, 0x6b, 0x69, 0x6e, 0x7a, 0x69, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chronicle_proto_rawDescOnce sync.Once
	file_proto_chronicle_proto_rawDescData = file_proto_chronicle_proto_rawDesc
)

func file_proto_chronicle_proto_rawDescGZIP() []byte {
	file_proto_chronicle_proto_rawDescOnce.Do(func() {
		file_proto_chronicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chronicle_proto_rawDescData)
	})
	return file_proto_chronicle_proto_rawDescData
}

var file_proto_chronicle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chronicle_proto_goTypes = []interface{}{
	(*Chronicle)(nil),             // 0: chronicle.Chronicle
	(*Debtor)(nil),                // 1: chronicle.Debtor
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_chronicle_proto_depIdxs = []int32{
	2, // 0: chronicle.Chronicle.date:type_name -> google.protobuf.Timestamp
	1, // 1: chronicle.Chronicle.debtors:type_name -> chronicle.Debtor
	0, // 2: chronicle.ChronicleService.All:input_type -> chronicle.Chronicle
	0, // 3: chronicle.ChronicleService.Create:input_type -> chronicle.Chronicle
	0, // 4: chronicle.ChronicleService.All:output_type -> chronicle.Chronicle
	0, // 5: chronicle.ChronicleService.Create:output_type -> chronicle.Chronicle
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_chronicle_proto_init() }
func file_proto_chronicle_proto_init() {
	if File_proto_chronicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chronicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chronicle); i {
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
		file_proto_chronicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Debtor); i {
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
			RawDescriptor: file_proto_chronicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chronicle_proto_goTypes,
		DependencyIndexes: file_proto_chronicle_proto_depIdxs,
		MessageInfos:      file_proto_chronicle_proto_msgTypes,
	}.Build()
	File_proto_chronicle_proto = out.File
	file_proto_chronicle_proto_rawDesc = nil
	file_proto_chronicle_proto_goTypes = nil
	file_proto_chronicle_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChronicleServiceClient is the client API for ChronicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChronicleServiceClient interface {
	All(ctx context.Context, in *Chronicle, opts ...grpc.CallOption) (*Chronicle, error)
	Create(ctx context.Context, in *Chronicle, opts ...grpc.CallOption) (*Chronicle, error)
}

type chronicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChronicleServiceClient(cc grpc.ClientConnInterface) ChronicleServiceClient {
	return &chronicleServiceClient{cc}
}

func (c *chronicleServiceClient) All(ctx context.Context, in *Chronicle, opts ...grpc.CallOption) (*Chronicle, error) {
	out := new(Chronicle)
	err := c.cc.Invoke(ctx, "/chronicle.ChronicleService/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chronicleServiceClient) Create(ctx context.Context, in *Chronicle, opts ...grpc.CallOption) (*Chronicle, error) {
	out := new(Chronicle)
	err := c.cc.Invoke(ctx, "/chronicle.ChronicleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChronicleServiceServer is the server API for ChronicleService service.
type ChronicleServiceServer interface {
	All(context.Context, *Chronicle) (*Chronicle, error)
	Create(context.Context, *Chronicle) (*Chronicle, error)
}

// UnimplementedChronicleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChronicleServiceServer struct {
}

func (*UnimplementedChronicleServiceServer) All(context.Context, *Chronicle) (*Chronicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (*UnimplementedChronicleServiceServer) Create(context.Context, *Chronicle) (*Chronicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterChronicleServiceServer(s *grpc.Server, srv ChronicleServiceServer) {
	s.RegisterService(&_ChronicleService_serviceDesc, srv)
}

func _ChronicleService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chronicle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChronicleServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chronicle.ChronicleService/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChronicleServiceServer).All(ctx, req.(*Chronicle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChronicleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chronicle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChronicleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chronicle.ChronicleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChronicleServiceServer).Create(ctx, req.(*Chronicle))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChronicleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chronicle.ChronicleService",
	HandlerType: (*ChronicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _ChronicleService_All_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ChronicleService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chronicle.proto",
}
