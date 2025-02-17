// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/callback/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryParamsRequest is the request for Query.Params.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{0}
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

// QueryParamsResponse is the response for Query.Params.
type QueryParamsResponse struct {
	// params defines all the module parameters.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{1}
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

// QueryEstimateCallbackFeesRequest is the request for Query.EstimateCallbackFees.
type QueryEstimateCallbackFeesRequest struct {
	// block_height is the height at which to estimate the callback fees
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QueryEstimateCallbackFeesRequest) Reset()         { *m = QueryEstimateCallbackFeesRequest{} }
func (m *QueryEstimateCallbackFeesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEstimateCallbackFeesRequest) ProtoMessage()    {}
func (*QueryEstimateCallbackFeesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{2}
}
func (m *QueryEstimateCallbackFeesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEstimateCallbackFeesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEstimateCallbackFeesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEstimateCallbackFeesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEstimateCallbackFeesRequest.Merge(m, src)
}
func (m *QueryEstimateCallbackFeesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEstimateCallbackFeesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEstimateCallbackFeesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEstimateCallbackFeesRequest proto.InternalMessageInfo

func (m *QueryEstimateCallbackFeesRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// QueryEstimateCallbackFeesResponse is the response for Query.EstimateCallbackFees.
type QueryEstimateCallbackFeesResponse struct {
	// total_fees is the total fees that needs to be paid by the contract to reserve a callback
	TotalFees *types.Coin `protobuf:"bytes,1,opt,name=total_fees,json=totalFees,proto3" json:"total_fees,omitempty"`
	// fee_split is the breakdown of the total_fees
	FeeSplit *CallbackFeesFeeSplit `protobuf:"bytes,2,opt,name=fee_split,json=feeSplit,proto3" json:"fee_split,omitempty"`
}

func (m *QueryEstimateCallbackFeesResponse) Reset()         { *m = QueryEstimateCallbackFeesResponse{} }
func (m *QueryEstimateCallbackFeesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEstimateCallbackFeesResponse) ProtoMessage()    {}
func (*QueryEstimateCallbackFeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{3}
}
func (m *QueryEstimateCallbackFeesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEstimateCallbackFeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEstimateCallbackFeesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEstimateCallbackFeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEstimateCallbackFeesResponse.Merge(m, src)
}
func (m *QueryEstimateCallbackFeesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEstimateCallbackFeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEstimateCallbackFeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEstimateCallbackFeesResponse proto.InternalMessageInfo

func (m *QueryEstimateCallbackFeesResponse) GetTotalFees() *types.Coin {
	if m != nil {
		return m.TotalFees
	}
	return nil
}

func (m *QueryEstimateCallbackFeesResponse) GetFeeSplit() *CallbackFeesFeeSplit {
	if m != nil {
		return m.FeeSplit
	}
	return nil
}

// QueryCallbacksRequest is the request for Query.Callbacks.
type QueryCallbacksRequest struct {
	// block_height is the height at which to query the callbacks
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QueryCallbacksRequest) Reset()         { *m = QueryCallbacksRequest{} }
func (m *QueryCallbacksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCallbacksRequest) ProtoMessage()    {}
func (*QueryCallbacksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{4}
}
func (m *QueryCallbacksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCallbacksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCallbacksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCallbacksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCallbacksRequest.Merge(m, src)
}
func (m *QueryCallbacksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCallbacksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCallbacksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCallbacksRequest proto.InternalMessageInfo

func (m *QueryCallbacksRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// QueryCallbacksResponse is the response for Query.Callbacks.
type QueryCallbacksResponse struct {
	// callbacks is the list of callbacks registered at the given height
	Callbacks []*Callback `protobuf:"bytes,1,rep,name=callbacks,proto3" json:"callbacks,omitempty"`
}

func (m *QueryCallbacksResponse) Reset()         { *m = QueryCallbacksResponse{} }
func (m *QueryCallbacksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCallbacksResponse) ProtoMessage()    {}
func (*QueryCallbacksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c34fd4ae1f0e6aa, []int{5}
}
func (m *QueryCallbacksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCallbacksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCallbacksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCallbacksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCallbacksResponse.Merge(m, src)
}
func (m *QueryCallbacksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCallbacksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCallbacksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCallbacksResponse proto.InternalMessageInfo

func (m *QueryCallbacksResponse) GetCallbacks() []*Callback {
	if m != nil {
		return m.Callbacks
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "archway.callback.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "archway.callback.v1.QueryParamsResponse")
	proto.RegisterType((*QueryEstimateCallbackFeesRequest)(nil), "archway.callback.v1.QueryEstimateCallbackFeesRequest")
	proto.RegisterType((*QueryEstimateCallbackFeesResponse)(nil), "archway.callback.v1.QueryEstimateCallbackFeesResponse")
	proto.RegisterType((*QueryCallbacksRequest)(nil), "archway.callback.v1.QueryCallbacksRequest")
	proto.RegisterType((*QueryCallbacksResponse)(nil), "archway.callback.v1.QueryCallbacksResponse")
}

func init() { proto.RegisterFile("archway/callback/v1/query.proto", fileDescriptor_0c34fd4ae1f0e6aa) }

var fileDescriptor_0c34fd4ae1f0e6aa = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0xe6, 0xf7, 0x8b, 0xc8, 0x86, 0xd3, 0x36, 0xa0, 0x92, 0x52, 0x37, 0x35, 0x12,
	0x04, 0xaa, 0xee, 0x2a, 0xa9, 0x8a, 0xf8, 0x73, 0x6b, 0xd5, 0x88, 0x1b, 0xc5, 0x88, 0x0b, 0x97,
	0x68, 0x6d, 0x26, 0x8e, 0x55, 0xdb, 0xeb, 0x7a, 0x37, 0x69, 0xcd, 0x09, 0x71, 0xe6, 0x80, 0xc4,
	0x33, 0xf0, 0x12, 0x3c, 0x41, 0x8f, 0x95, 0x90, 0x10, 0x27, 0x84, 0x12, 0x1e, 0x04, 0x79, 0xbd,
	0x6e, 0xa1, 0xb8, 0x81, 0xde, 0x56, 0xb3, 0xdf, 0xef, 0xcc, 0x67, 0x76, 0xc6, 0x46, 0xab, 0x2c,
	0x71, 0x47, 0x87, 0x2c, 0xa5, 0x2e, 0x0b, 0x02, 0x87, 0xb9, 0xfb, 0x74, 0xd2, 0xa5, 0x07, 0x63,
	0x48, 0x52, 0x12, 0x27, 0x5c, 0x72, 0xbc, 0xa8, 0x05, 0xa4, 0x10, 0x90, 0x49, 0xb7, 0xd5, 0xf4,
	0xb8, 0xc7, 0xd5, 0x3d, 0xcd, 0x4e, 0xb9, 0xb4, 0x75, 0xd3, 0xe3, 0xdc, 0x0b, 0x80, 0xb2, 0xd8,
	0xa7, 0x2c, 0x8a, 0xb8, 0x64, 0xd2, 0xe7, 0x91, 0xd0, 0xb7, 0xa6, 0xcb, 0x45, 0xc8, 0x05, 0x75,
	0x98, 0x00, 0x3a, 0xe9, 0x3a, 0x20, 0x59, 0x97, 0xba, 0xdc, 0x8f, 0xf4, 0xbd, 0x55, 0x46, 0x72,
	0x5a, 0x54, 0x69, 0xac, 0x26, 0xc2, 0xcf, 0x32, 0xb6, 0x3d, 0x96, 0xb0, 0x50, 0xd8, 0x70, 0x30,
	0x06, 0x21, 0xad, 0x3d, 0xb4, 0xf8, 0x5b, 0x54, 0xc4, 0x3c, 0x12, 0x80, 0x1f, 0xa2, 0x5a, 0xac,
	0x22, 0x4b, 0x46, 0xdb, 0xe8, 0x34, 0x7a, 0xcb, 0xa4, 0xa4, 0x15, 0x92, 0x9b, 0xb6, 0xff, 0x3b,
	0xfe, 0xb6, 0x5a, 0xb1, 0xb5, 0xc1, 0xda, 0x45, 0x6d, 0x95, 0x71, 0x57, 0x48, 0x3f, 0x64, 0x12,
	0x76, 0xb4, 0xa1, 0x0f, 0x50, 0x54, 0xc5, 0x6b, 0xe8, 0xaa, 0x13, 0x70, 0x77, 0x7f, 0x30, 0x02,
	0xdf, 0x1b, 0x49, 0x55, 0xa4, 0x6a, 0x37, 0x54, 0xec, 0x89, 0x0a, 0x59, 0x1f, 0x0d, 0xb4, 0x36,
	0x27, 0x8f, 0xe6, 0x7c, 0x80, 0x90, 0xe4, 0x92, 0x05, 0x83, 0x21, 0x40, 0xc1, 0x7a, 0x83, 0xe4,
	0xaf, 0x45, 0xb2, 0xd7, 0x22, 0xfa, 0xb5, 0xc8, 0x0e, 0xf7, 0x23, 0xbb, 0xae, 0xc4, 0x59, 0x06,
	0xdc, 0x47, 0xf5, 0x21, 0xc0, 0x40, 0xc4, 0x81, 0x2f, 0x97, 0x16, 0x94, 0xf1, 0x6e, 0x69, 0x93,
	0xbf, 0xd6, 0xed, 0x03, 0x3c, 0xcf, 0x0c, 0xf6, 0x95, 0xa1, 0x3e, 0x59, 0x8f, 0xd0, 0x35, 0x85,
	0x59, 0xc8, 0x2e, 0xd3, 0xe3, 0x0b, 0x74, 0xfd, 0xbc, 0x57, 0xf7, 0xf5, 0x18, 0xd5, 0x0b, 0x86,
	0xac, 0xad, 0x6a, 0xa7, 0xd1, 0x5b, 0x99, 0x4b, 0x67, 0x9f, 0xe9, 0x7b, 0x5f, 0xaa, 0xe8, 0x7f,
	0x95, 0x17, 0xbf, 0x31, 0x50, 0x2d, 0x1f, 0x12, 0xbe, 0x53, 0x6a, 0xff, 0x73, 0x23, 0x5a, 0x9d,
	0xbf, 0x0b, 0x73, 0x48, 0xeb, 0xd6, 0xdb, 0xcf, 0x3f, 0x3e, 0x2c, 0xac, 0xe0, 0x65, 0x5a, 0xb6,
	0x7e, 0xf9, 0x3a, 0xe0, 0x4f, 0x06, 0x6a, 0x96, 0x8d, 0x10, 0x6f, 0x5d, 0x5c, 0x67, 0xce, 0xea,
	0xb4, 0xee, 0x5f, 0xd6, 0xa6, 0x61, 0x37, 0x15, 0xec, 0x06, 0x5e, 0x2f, 0x85, 0x05, 0x6d, 0x1d,
	0x14, 0x41, 0xb5, 0x50, 0xf8, 0x9d, 0x81, 0xea, 0xa7, 0xc3, 0xc1, 0xf7, 0x2e, 0x2e, 0x7d, 0x7e,
	0xfa, 0xad, 0xf5, 0x7f, 0xd2, 0x6a, 0xb6, 0xdb, 0x8a, 0xad, 0x8d, 0x4d, 0x3a, 0xef, 0x3b, 0x16,
	0xdb, 0x4f, 0x8f, 0xa7, 0xa6, 0x71, 0x32, 0x35, 0x8d, 0xef, 0x53, 0xd3, 0x78, 0x3f, 0x33, 0x2b,
	0x27, 0x33, 0xb3, 0xf2, 0x75, 0x66, 0x56, 0x5e, 0x6e, 0x79, 0xbe, 0x1c, 0x8d, 0x1d, 0xe2, 0xf2,
	0x90, 0xbe, 0x4a, 0x43, 0x88, 0x84, 0xcf, 0xa3, 0xa3, 0xf4, 0x35, 0x4d, 0x78, 0x10, 0xb0, 0x38,
	0xde, 0x38, 0x64, 0x22, 0xa4, 0x47, 0x67, 0x79, 0x65, 0x1a, 0x83, 0x70, 0x6a, 0xea, 0xd7, 0xb0,
	0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x15, 0xe9, 0xba, 0xca, 0x04, 0x00, 0x00,
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
	// Params returns module parameters
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EstimateCallbackFees returns the total amount of callback fees a contract needs to pay to register the callback
	EstimateCallbackFees(ctx context.Context, in *QueryEstimateCallbackFeesRequest, opts ...grpc.CallOption) (*QueryEstimateCallbackFeesResponse, error)
	// Callbacks returns all the callbacks registered at a given height
	Callbacks(ctx context.Context, in *QueryCallbacksRequest, opts ...grpc.CallOption) (*QueryCallbacksResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/archway.callback.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EstimateCallbackFees(ctx context.Context, in *QueryEstimateCallbackFeesRequest, opts ...grpc.CallOption) (*QueryEstimateCallbackFeesResponse, error) {
	out := new(QueryEstimateCallbackFeesResponse)
	err := c.cc.Invoke(ctx, "/archway.callback.v1.Query/EstimateCallbackFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Callbacks(ctx context.Context, in *QueryCallbacksRequest, opts ...grpc.CallOption) (*QueryCallbacksResponse, error) {
	out := new(QueryCallbacksResponse)
	err := c.cc.Invoke(ctx, "/archway.callback.v1.Query/Callbacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns module parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EstimateCallbackFees returns the total amount of callback fees a contract needs to pay to register the callback
	EstimateCallbackFees(context.Context, *QueryEstimateCallbackFeesRequest) (*QueryEstimateCallbackFeesResponse, error)
	// Callbacks returns all the callbacks registered at a given height
	Callbacks(context.Context, *QueryCallbacksRequest) (*QueryCallbacksResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) EstimateCallbackFees(ctx context.Context, req *QueryEstimateCallbackFeesRequest) (*QueryEstimateCallbackFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateCallbackFees not implemented")
}
func (*UnimplementedQueryServer) Callbacks(ctx context.Context, req *QueryCallbacksRequest) (*QueryCallbacksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callbacks not implemented")
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
		FullMethod: "/archway.callback.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EstimateCallbackFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEstimateCallbackFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EstimateCallbackFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.callback.v1.Query/EstimateCallbackFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EstimateCallbackFees(ctx, req.(*QueryEstimateCallbackFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Callbacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallbacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Callbacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.callback.v1.Query/Callbacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Callbacks(ctx, req.(*QueryCallbacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "archway.callback.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EstimateCallbackFees",
			Handler:    _Query_EstimateCallbackFees_Handler,
		},
		{
			MethodName: "Callbacks",
			Handler:    _Query_Callbacks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archway/callback/v1/query.proto",
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

func (m *QueryEstimateCallbackFeesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEstimateCallbackFeesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEstimateCallbackFeesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEstimateCallbackFeesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEstimateCallbackFeesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEstimateCallbackFeesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeSplit != nil {
		{
			size, err := m.FeeSplit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TotalFees != nil {
		{
			size, err := m.TotalFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCallbacksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCallbacksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCallbacksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCallbacksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCallbacksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCallbacksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Callbacks) > 0 {
		for iNdEx := len(m.Callbacks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Callbacks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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

func (m *QueryEstimateCallbackFeesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	return n
}

func (m *QueryEstimateCallbackFeesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalFees != nil {
		l = m.TotalFees.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.FeeSplit != nil {
		l = m.FeeSplit.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCallbacksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	return n
}

func (m *QueryCallbacksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Callbacks) > 0 {
		for _, e := range m.Callbacks {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
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
func (m *QueryEstimateCallbackFeesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEstimateCallbackFeesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEstimateCallbackFeesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryEstimateCallbackFeesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEstimateCallbackFeesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEstimateCallbackFeesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFees", wireType)
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
			if m.TotalFees == nil {
				m.TotalFees = &types.Coin{}
			}
			if err := m.TotalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeSplit", wireType)
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
			if m.FeeSplit == nil {
				m.FeeSplit = &CallbackFeesFeeSplit{}
			}
			if err := m.FeeSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCallbacksRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCallbacksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCallbacksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryCallbacksResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCallbacksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCallbacksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Callbacks", wireType)
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
			m.Callbacks = append(m.Callbacks, &Callback{})
			if err := m.Callbacks[len(m.Callbacks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
