// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: shorten.proto

// shorten v1

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorten_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *ShortenRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinifyURL string `protobuf:"bytes,1,opt,name=minifyURL,proto3" json:"minifyURL,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorten_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetMinifyURL() string {
	if x != nil {
		return x.MinifyURL
	}
	return ""
}

type FetchURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinifyURL string `protobuf:"bytes,1,opt,name=minifyURL,proto3" json:"minifyURL,omitempty"`
}

func (x *FetchURLRequest) Reset() {
	*x = FetchURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorten_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchURLRequest) ProtoMessage() {}

func (x *FetchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchURLRequest.ProtoReflect.Descriptor instead.
func (*FetchURLRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{2}
}

func (x *FetchURLRequest) GetMinifyURL() string {
	if x != nil {
		return x.MinifyURL
	}
	return ""
}

type FetchURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
}

func (x *FetchURLResponse) Reset() {
	*x = FetchURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorten_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchURLResponse) ProtoMessage() {}

func (x *FetchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchURLResponse.ProtoReflect.Descriptor instead.
func (*FetchURLResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{3}
}

func (x *FetchURLResponse) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

var File_shorten_proto protoreflect.FileDescriptor

var file_shorten_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0x34, 0x0a,
	0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52,
	0x4c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x69, 0x66, 0x79,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x69, 0x66,
	0x79, 0x55, 0x52, 0x4c, 0x22, 0x2f, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x69, 0x66,
	0x79, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x69,
	0x66, 0x79, 0x55, 0x52, 0x4c, 0x22, 0x34, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x32, 0xa3, 0x01, 0x0a, 0x0e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x75, 0x72, 0x6c, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x52, 0x4c, 0x12, 0x1d, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shorten_proto_rawDescOnce sync.Once
	file_shorten_proto_rawDescData = file_shorten_proto_rawDesc
)

func file_shorten_proto_rawDescGZIP() []byte {
	file_shorten_proto_rawDescOnce.Do(func() {
		file_shorten_proto_rawDescData = protoimpl.X.CompressGZIP(file_shorten_proto_rawDescData)
	})
	return file_shorten_proto_rawDescData
}

var file_shorten_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shorten_proto_goTypes = []interface{}{
	(*ShortenRequest)(nil),   // 0: urlshortener.ShortenRequest
	(*ShortenResponse)(nil),  // 1: urlshortener.ShortenResponse
	(*FetchURLRequest)(nil),  // 2: urlshortener.FetchURLRequest
	(*FetchURLResponse)(nil), // 3: urlshortener.FetchURLResponse
}
var file_shorten_proto_depIdxs = []int32{
	0, // 0: urlshortener.ShortenService.Shorten:input_type -> urlshortener.ShortenRequest
	2, // 1: urlshortener.ShortenService.FetchURL:input_type -> urlshortener.FetchURLRequest
	1, // 2: urlshortener.ShortenService.Shorten:output_type -> urlshortener.ShortenResponse
	3, // 3: urlshortener.ShortenService.FetchURL:output_type -> urlshortener.FetchURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shorten_proto_init() }
func file_shorten_proto_init() {
	if File_shorten_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shorten_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_shorten_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenResponse); i {
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
		file_shorten_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchURLRequest); i {
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
		file_shorten_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchURLResponse); i {
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
			RawDescriptor: file_shorten_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorten_proto_goTypes,
		DependencyIndexes: file_shorten_proto_depIdxs,
		MessageInfos:      file_shorten_proto_msgTypes,
	}.Build()
	File_shorten_proto = out.File
	file_shorten_proto_rawDesc = nil
	file_shorten_proto_goTypes = nil
	file_shorten_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShortenServiceClient is the client API for ShortenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortenServiceClient interface {
	Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
	FetchURL(ctx context.Context, in *FetchURLRequest, opts ...grpc.CallOption) (*FetchURLResponse, error)
}

type shortenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenServiceClient(cc grpc.ClientConnInterface) ShortenServiceClient {
	return &shortenServiceClient{cc}
}

func (c *shortenServiceClient) Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	out := new(ShortenResponse)
	err := c.cc.Invoke(ctx, "/urlshortener.ShortenService/Shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenServiceClient) FetchURL(ctx context.Context, in *FetchURLRequest, opts ...grpc.CallOption) (*FetchURLResponse, error) {
	out := new(FetchURLResponse)
	err := c.cc.Invoke(ctx, "/urlshortener.ShortenService/FetchURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenServiceServer is the server API for ShortenService service.
type ShortenServiceServer interface {
	Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error)
	FetchURL(context.Context, *FetchURLRequest) (*FetchURLResponse, error)
}

// UnimplementedShortenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShortenServiceServer struct {
}

func (*UnimplementedShortenServiceServer) Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (*UnimplementedShortenServiceServer) FetchURL(context.Context, *FetchURLRequest) (*FetchURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchURL not implemented")
}

func RegisterShortenServiceServer(s *grpc.Server, srv ShortenServiceServer) {
	s.RegisterService(&_ShortenService_serviceDesc, srv)
}

func _ShortenService_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenServiceServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlshortener.ShortenService/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenServiceServer).Shorten(ctx, req.(*ShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenService_FetchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenServiceServer).FetchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlshortener.ShortenService/FetchURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenServiceServer).FetchURL(ctx, req.(*FetchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShortenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "urlshortener.ShortenService",
	HandlerType: (*ShortenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shorten",
			Handler:    _ShortenService_Shorten_Handler,
		},
		{
			MethodName: "FetchURL",
			Handler:    _ShortenService_FetchURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shorten.proto",
}
