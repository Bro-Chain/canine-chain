// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jklmint/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryInflationRequest is the request type for the Query/Inflation RPC method.
type QueryInflationRequest struct {
}

func (m *QueryInflationRequest) Reset()         { *m = QueryInflationRequest{} }
func (m *QueryInflationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInflationRequest) ProtoMessage()    {}
func (*QueryInflationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{2}
}
func (m *QueryInflationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationRequest.Merge(m, src)
}
func (m *QueryInflationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationRequest proto.InternalMessageInfo

// QueryInflationResponse is the response type for the Query/Inflation RPC
// method.
type QueryInflationResponse struct {
	// inflation is the current minting inflation value.
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
}

func (m *QueryInflationResponse) Reset()         { *m = QueryInflationResponse{} }
func (m *QueryInflationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInflationResponse) ProtoMessage()    {}
func (*QueryInflationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{3}
}
func (m *QueryInflationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationResponse.Merge(m, src)
}
func (m *QueryInflationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationResponse proto.InternalMessageInfo

type QueryGetInflationRequest struct {
}

func (m *QueryGetInflationRequest) Reset()         { *m = QueryGetInflationRequest{} }
func (m *QueryGetInflationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetInflationRequest) ProtoMessage()    {}
func (*QueryGetInflationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{4}
}
func (m *QueryGetInflationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetInflationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetInflationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetInflationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetInflationRequest.Merge(m, src)
}
func (m *QueryGetInflationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetInflationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetInflationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetInflationRequest proto.InternalMessageInfo

type QueryGetInflationResponse struct {
	Inflation string `protobuf:"bytes,1,opt,name=inflation,proto3" json:"inflation,omitempty"`
}

func (m *QueryGetInflationResponse) Reset()         { *m = QueryGetInflationResponse{} }
func (m *QueryGetInflationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetInflationResponse) ProtoMessage()    {}
func (*QueryGetInflationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34daa9d33a6ad5, []int{5}
}
func (m *QueryGetInflationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetInflationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetInflationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetInflationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetInflationResponse.Merge(m, src)
}
func (m *QueryGetInflationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetInflationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetInflationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetInflationResponse proto.InternalMessageInfo

func (m *QueryGetInflationResponse) GetInflation() string {
	if m != nil {
		return m.Inflation
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "jackaldao.canine.jklmint.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "jackaldao.canine.jklmint.QueryParamsResponse")
	proto.RegisterType((*QueryInflationRequest)(nil), "jackaldao.canine.jklmint.QueryInflationRequest")
	proto.RegisterType((*QueryInflationResponse)(nil), "jackaldao.canine.jklmint.QueryInflationResponse")
	proto.RegisterType((*QueryGetInflationRequest)(nil), "jackaldao.canine.jklmint.QueryGetInflationRequest")
	proto.RegisterType((*QueryGetInflationResponse)(nil), "jackaldao.canine.jklmint.QueryGetInflationResponse")
}

func init() { proto.RegisterFile("jklmint/query.proto", fileDescriptor_9f34daa9d33a6ad5) }

var fileDescriptor_9f34daa9d33a6ad5 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0xfe, 0x54, 0xaa, 0xb9, 0xc9, 0x57, 0xa0, 0x84, 0x23, 0x57, 0x05, 0xe9, 0x74,
	0x3a, 0x5d, 0xed, 0xbb, 0xde, 0xc4, 0xc2, 0x50, 0x21, 0x10, 0x12, 0x03, 0x54, 0x62, 0x61, 0x41,
	0x4e, 0xea, 0x0b, 0xb9, 0x26, 0x76, 0x5a, 0xbb, 0x88, 0xae, 0x8c, 0x4c, 0x48, 0x48, 0x7c, 0x07,
	0xbe, 0x49, 0xc7, 0x4a, 0x2c, 0xc0, 0x50, 0xa1, 0x96, 0x0f, 0x82, 0x6a, 0x3b, 0x69, 0xe9, 0x1f,
	0xb5, 0x4c, 0x89, 0xec, 0xf7, 0x79, 0x9e, 0xdf, 0xfb, 0xda, 0x86, 0xfb, 0x57, 0x9d, 0x24, 0x8d,
	0xb9, 0x22, 0xdd, 0x3e, 0xeb, 0x0d, 0x70, 0xd6, 0x13, 0x4a, 0xa0, 0xea, 0x15, 0x0d, 0x3b, 0x34,
	0x69, 0x53, 0x81, 0x43, 0xca, 0x63, 0xce, 0xb0, 0xad, 0x72, 0x2b, 0x91, 0x88, 0x84, 0x2e, 0x22,
	0xb3, 0x3f, 0x53, 0xef, 0x1e, 0x44, 0x42, 0x44, 0x09, 0x23, 0x34, 0x8b, 0x09, 0xe5, 0x5c, 0x28,
	0xaa, 0x62, 0xc1, 0xa5, 0xdd, 0x3d, 0x09, 0x85, 0x4c, 0x85, 0x24, 0x01, 0x95, 0xcc, 0xc4, 0x90,
	0xf7, 0xe7, 0x01, 0x53, 0xf4, 0x9c, 0x64, 0x34, 0x8a, 0xb9, 0x2e, 0xb6, 0xb5, 0x95, 0x1c, 0x27,
	0xa3, 0x3d, 0x9a, 0x5a, 0x07, 0xbf, 0x02, 0xd1, 0xab, 0x99, 0xee, 0xa5, 0x5e, 0x6c, 0xb1, 0x6e,
	0x9f, 0x49, 0xe5, 0xbf, 0x86, 0xfb, 0xff, 0xac, 0xca, 0x4c, 0x70, 0xc9, 0xd0, 0x63, 0x58, 0x32,
	0xe2, 0x2a, 0xa8, 0x81, 0xe3, 0x5b, 0x8d, 0x1a, 0xde, 0xd4, 0x0d, 0x36, 0xca, 0xe6, 0x8d, 0xe1,
	0xf8, 0xd0, 0x69, 0x59, 0x95, 0x7f, 0x17, 0xde, 0xd6, 0xb6, 0xcf, 0xf9, 0x65, 0xa2, 0xd1, 0xf2,
	0xbc, 0x4b, 0x78, 0x67, 0x79, 0xc3, 0x46, 0xbe, 0x80, 0xe5, 0x38, 0x5f, 0xd4, 0xa9, 0x7b, 0x4d,
	0x3c, 0xf3, 0xfc, 0x35, 0x3e, 0x3c, 0x8a, 0x62, 0xf5, 0xae, 0x1f, 0xe0, 0x50, 0xa4, 0xc4, 0xce,
	0xc1, 0x7c, 0xea, 0xb2, 0xdd, 0x21, 0x6a, 0x90, 0x31, 0x89, 0x9f, 0xb0, 0xb0, 0x35, 0x37, 0xf0,
	0x5d, 0x58, 0xd5, 0x39, 0xcf, 0x98, 0x5a, 0x61, 0x78, 0x04, 0xef, 0xad, 0xd9, 0xb3, 0x18, 0x07,
	0xcb, 0x18, 0xe5, 0x05, 0xdb, 0xc6, 0xcf, 0xeb, 0xf0, 0xa6, 0xd6, 0xa2, 0x4f, 0x00, 0x96, 0x4c,
	0xeb, 0xe8, 0x74, 0xf3, 0x70, 0x56, 0x27, 0xee, 0xd6, 0x77, 0xac, 0x36, 0x3c, 0xfe, 0xc3, 0x8f,
	0xdf, 0xff, 0x7c, 0xb9, 0xf6, 0x00, 0xdd, 0xcf, 0x3b, 0xd7, 0x27, 0x3b, 0x3f, 0x7b, 0x4d, 0xf0,
	0x15, 0xc0, 0x72, 0xd1, 0x0a, 0x22, 0x5b, 0x12, 0x96, 0x07, 0xe2, 0x9e, 0xed, 0x2e, 0xb0, 0x54,
	0x47, 0x9a, 0xaa, 0x86, 0xbc, 0xb5, 0x54, 0xc5, 0xbc, 0xd0, 0x37, 0x00, 0xf7, 0x16, 0xc7, 0x8c,
	0x1a, 0x5b, 0xa2, 0xd6, 0x9c, 0x97, 0x7b, 0xf1, 0x5f, 0x1a, 0x4b, 0x78, 0xa6, 0x09, 0x4f, 0xd0,
	0x31, 0x31, 0xe2, 0x7a, 0x9b, 0x0a, 0x62, 0xd4, 0x24, 0x7f, 0x1f, 0x11, 0x53, 0x6f, 0x0b, 0xd6,
	0xe6, 0xd3, 0xe1, 0xc4, 0x03, 0xa3, 0x89, 0x07, 0x7e, 0x4f, 0x3c, 0xf0, 0x79, 0xea, 0x39, 0xa3,
	0xa9, 0xe7, 0xfc, 0x98, 0x7a, 0xce, 0x9b, 0xd3, 0x85, 0xfb, 0xb7, 0xea, 0xf6, 0xa1, 0xf0, 0xd3,
	0x37, 0x31, 0x28, 0xe9, 0xf7, 0x76, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x04, 0x11, 0xa8, 0xae,
	0x16, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Inflation returns the current minting inflation value.
	Inflation(ctx context.Context, in *QueryInflationRequest, opts ...grpc.CallOption) (*QueryInflationResponse, error)
	// Queries a list of GetInflation items.
	GetInflation(ctx context.Context, in *QueryGetInflationRequest, opts ...grpc.CallOption) (*QueryGetInflationResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/jackaldao.canine.jklmint.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Inflation(ctx context.Context, in *QueryInflationRequest, opts ...grpc.CallOption) (*QueryInflationResponse, error) {
	out := new(QueryInflationResponse)
	err := c.cc.Invoke(ctx, "/jackaldao.canine.jklmint.Query/Inflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetInflation(ctx context.Context, in *QueryGetInflationRequest, opts ...grpc.CallOption) (*QueryGetInflationResponse, error) {
	out := new(QueryGetInflationResponse)
	err := c.cc.Invoke(ctx, "/jackaldao.canine.jklmint.Query/GetInflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Inflation returns the current minting inflation value.
	Inflation(context.Context, *QueryInflationRequest) (*QueryInflationResponse, error)
	// Queries a list of GetInflation items.
	GetInflation(context.Context, *QueryGetInflationRequest) (*QueryGetInflationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Inflation(ctx context.Context, req *QueryInflationRequest) (*QueryInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inflation not implemented")
}
func (*UnimplementedQueryServer) GetInflation(ctx context.Context, req *QueryGetInflationRequest) (*QueryGetInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInflation not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jackaldao.canine.jklmint.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Inflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Inflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jackaldao.canine.jklmint.Query/Inflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Inflation(ctx, req.(*QueryInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jackaldao.canine.jklmint.Query/GetInflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetInflation(ctx, req.(*QueryGetInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jackaldao.canine.jklmint.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Inflation",
			Handler:    _Query_Inflation_Handler,
		},
		{
			MethodName: "GetInflation",
			Handler:    _Query_GetInflation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jklmint/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryInflationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryInflationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetInflationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetInflationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetInflationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetInflationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetInflationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetInflationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inflation) > 0 {
		i -= len(m.Inflation)
		copy(dAtA[i:], m.Inflation)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Inflation)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryInflationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryInflationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetInflationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetInflationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Inflation)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryInflationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryInflationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetInflationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetInflationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetInflationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetInflationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetInflationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetInflationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inflation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
