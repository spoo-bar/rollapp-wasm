// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/cwerrors/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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
	return fileDescriptor_a1be36abcb817ffd, []int{0}
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
	return fileDescriptor_a1be36abcb817ffd, []int{1}
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

// QueryErrorsRequest is the request for Query.Errors.
type QueryErrorsRequest struct {
	// contract_address is the address of the contract whose errors to query for
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryErrorsRequest) Reset()         { *m = QueryErrorsRequest{} }
func (m *QueryErrorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryErrorsRequest) ProtoMessage()    {}
func (*QueryErrorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1be36abcb817ffd, []int{2}
}
func (m *QueryErrorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryErrorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryErrorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryErrorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryErrorsRequest.Merge(m, src)
}
func (m *QueryErrorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryErrorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryErrorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryErrorsRequest proto.InternalMessageInfo

func (m *QueryErrorsRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// QueryErrorsResponse is the response for Query.Errors.
type QueryErrorsResponse struct {
	// errors defines all the contract errors which will be returned
	Errors []SudoError `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors"`
}

func (m *QueryErrorsResponse) Reset()         { *m = QueryErrorsResponse{} }
func (m *QueryErrorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryErrorsResponse) ProtoMessage()    {}
func (*QueryErrorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1be36abcb817ffd, []int{3}
}
func (m *QueryErrorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryErrorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryErrorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryErrorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryErrorsResponse.Merge(m, src)
}
func (m *QueryErrorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryErrorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryErrorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryErrorsResponse proto.InternalMessageInfo

func (m *QueryErrorsResponse) GetErrors() []SudoError {
	if m != nil {
		return m.Errors
	}
	return nil
}

// QueryIsSubscribedRequest is the request for Query.IsSubscribed.
type QueryIsSubscribedRequest struct {
	// contract_address is the address of the contract to query if subscribed
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryIsSubscribedRequest) Reset()         { *m = QueryIsSubscribedRequest{} }
func (m *QueryIsSubscribedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIsSubscribedRequest) ProtoMessage()    {}
func (*QueryIsSubscribedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1be36abcb817ffd, []int{4}
}
func (m *QueryIsSubscribedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsSubscribedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsSubscribedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsSubscribedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsSubscribedRequest.Merge(m, src)
}
func (m *QueryIsSubscribedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsSubscribedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsSubscribedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsSubscribedRequest proto.InternalMessageInfo

func (m *QueryIsSubscribedRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// QueryIsSubscribedResponse is the response for Query.IsSubscribed.
type QueryIsSubscribedResponse struct {
	// subscribed defines if the contract is subscribed to sudo error callbacks
	Subscribed bool `protobuf:"varint,1,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	// subscription_valid_till defines the block height till which the
	// subscription is valid
	SubscriptionValidTill int64 `protobuf:"varint,2,opt,name=subscription_valid_till,json=subscriptionValidTill,proto3" json:"subscription_valid_till,omitempty"`
}

func (m *QueryIsSubscribedResponse) Reset()         { *m = QueryIsSubscribedResponse{} }
func (m *QueryIsSubscribedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIsSubscribedResponse) ProtoMessage()    {}
func (*QueryIsSubscribedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1be36abcb817ffd, []int{5}
}
func (m *QueryIsSubscribedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsSubscribedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsSubscribedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsSubscribedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsSubscribedResponse.Merge(m, src)
}
func (m *QueryIsSubscribedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsSubscribedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsSubscribedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsSubscribedResponse proto.InternalMessageInfo

func (m *QueryIsSubscribedResponse) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

func (m *QueryIsSubscribedResponse) GetSubscriptionValidTill() int64 {
	if m != nil {
		return m.SubscriptionValidTill
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "archway.cwerrors.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "archway.cwerrors.v1.QueryParamsResponse")
	proto.RegisterType((*QueryErrorsRequest)(nil), "archway.cwerrors.v1.QueryErrorsRequest")
	proto.RegisterType((*QueryErrorsResponse)(nil), "archway.cwerrors.v1.QueryErrorsResponse")
	proto.RegisterType((*QueryIsSubscribedRequest)(nil), "archway.cwerrors.v1.QueryIsSubscribedRequest")
	proto.RegisterType((*QueryIsSubscribedResponse)(nil), "archway.cwerrors.v1.QueryIsSubscribedResponse")
}

func init() { proto.RegisterFile("archway/cwerrors/v1/query.proto", fileDescriptor_a1be36abcb817ffd) }

var fileDescriptor_a1be36abcb817ffd = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xad, 0x2e, 0x3a, 0x15, 0x94, 0x69, 0xc5, 0x75, 0xab, 0xe9, 0x12, 0x05, 0x57,
	0xa1, 0x19, 0x76, 0x45, 0x41, 0x10, 0xc4, 0x42, 0x0f, 0x9e, 0xac, 0x59, 0xf1, 0xe0, 0x65, 0x99,
	0x24, 0x43, 0x3a, 0x90, 0xe4, 0xa5, 0x33, 0x93, 0xdd, 0xc6, 0x93, 0xf8, 0x09, 0x04, 0x4f, 0x7e,
	0x04, 0xbf, 0x49, 0x8f, 0x05, 0x2f, 0x9e, 0x44, 0x76, 0xfd, 0x20, 0xb2, 0x33, 0xb3, 0xed, 0x2e,
	0xa6, 0xad, 0x78, 0x9b, 0xbc, 0xf7, 0x7f, 0xff, 0xf7, 0x9b, 0xfc, 0x13, 0xb4, 0x45, 0x45, 0xb4,
	0x3f, 0xa6, 0x15, 0x89, 0xc6, 0x4c, 0x08, 0x10, 0x92, 0x8c, 0x7a, 0xe4, 0xa0, 0x64, 0xa2, 0xf2,
	0x0b, 0x01, 0x0a, 0xf0, 0xba, 0x15, 0xf8, 0x73, 0x81, 0x3f, 0xea, 0xb5, 0x37, 0x12, 0x48, 0x40,
	0xf7, 0xc9, 0xec, 0x64, 0xa4, 0xed, 0x3b, 0x09, 0x40, 0x92, 0x32, 0x42, 0x0b, 0x4e, 0x68, 0x9e,
	0x83, 0xa2, 0x8a, 0x43, 0x2e, 0x6d, 0xd7, 0x8d, 0x40, 0x66, 0x20, 0x49, 0x48, 0x25, 0x23, 0xa3,
	0x5e, 0xc8, 0x14, 0xed, 0x91, 0x08, 0x78, 0x6e, 0xfb, 0x5e, 0x1d, 0xc9, 0xc9, 0x52, 0xa3, 0xe9,
	0xd4, 0x69, 0x0a, 0x2a, 0x68, 0x66, 0x15, 0xde, 0x06, 0xc2, 0x6f, 0x66, 0xf4, 0x7b, 0xba, 0x18,
	0xb0, 0x83, 0x92, 0x49, 0xe5, 0xed, 0xa1, 0xf5, 0xa5, 0xaa, 0x2c, 0x20, 0x97, 0x0c, 0x3f, 0x43,
	0x4d, 0x33, 0xdc, 0x72, 0x3a, 0x4e, 0x77, 0xad, 0xbf, 0xe9, 0xd7, 0x5c, 0xd6, 0x37, 0x43, 0x3b,
	0x97, 0x8e, 0x7e, 0x6e, 0x35, 0x02, 0x3b, 0xe0, 0xbd, 0xb0, 0x7b, 0x76, 0xb5, 0xcc, 0xee, 0xc1,
	0x0f, 0xd1, 0x8d, 0x08, 0x72, 0x25, 0x68, 0xa4, 0x86, 0x34, 0x8e, 0x05, 0x93, 0xc6, 0xfa, 0x6a,
	0x70, 0x7d, 0x5e, 0x7f, 0x69, 0xca, 0xde, 0xc0, 0x22, 0xcd, 0x0d, 0x2c, 0xd2, 0x73, 0xd4, 0x34,
	0x9b, 0x5b, 0x4e, 0x67, 0xb5, 0xbb, 0xd6, 0x77, 0x6b, 0x91, 0x06, 0x65, 0x0c, 0x7a, 0x70, 0x4e,
	0x65, 0x5a, 0xde, 0x2e, 0x6a, 0x69, 0xd3, 0x57, 0x72, 0x50, 0x86, 0x32, 0x12, 0x3c, 0x64, 0xf1,
	0x7f, 0xb0, 0x49, 0x74, 0xbb, 0xc6, 0xc6, 0x12, 0xba, 0x08, 0xc9, 0x93, 0xaa, 0x76, 0xb8, 0x12,
	0x2c, 0x54, 0xf0, 0x53, 0x74, 0xcb, 0x3e, 0x15, 0xb3, 0xf8, 0x87, 0x23, 0x9a, 0xf2, 0x78, 0xa8,
	0x78, 0x9a, 0xb6, 0x56, 0x3a, 0x4e, 0x77, 0x35, 0xb8, 0xb9, 0xd8, 0x7e, 0x37, 0xeb, 0xbe, 0xe5,
	0x69, 0xda, 0xff, 0xb6, 0x8a, 0x2e, 0xeb, 0xad, 0xf8, 0xa3, 0x83, 0x9a, 0xe6, 0xa5, 0xe3, 0x07,
	0xb5, 0xd7, 0xff, 0x3b, 0xe1, 0x76, 0xf7, 0x62, 0xa1, 0xe1, 0xf7, 0xee, 0x7d, 0xfa, 0xfe, 0xfb,
	0xcb, 0xca, 0x5d, 0xbc, 0x49, 0xce, 0xfe, 0x98, 0x34, 0x82, 0x49, 0xe6, 0x3c, 0x84, 0xa5, 0xf0,
	0xcf, 0x43, 0x58, 0x0e, 0xf9, 0x02, 0x04, 0x73, 0xc2, 0x5f, 0x1d, 0x74, 0x6d, 0x31, 0x00, 0xbc,
	0x7d, 0xb6, 0x7f, 0x4d, 0xde, 0x6d, 0xff, 0x5f, 0xe5, 0x16, 0xea, 0x91, 0x86, 0xba, 0x8f, 0xbd,
	0x5a, 0x28, 0x2e, 0x87, 0xa7, 0x19, 0xef, 0xbc, 0x3e, 0x9a, 0xb8, 0xce, 0xf1, 0xc4, 0x75, 0x7e,
	0x4d, 0x5c, 0xe7, 0xf3, 0xd4, 0x6d, 0x1c, 0x4f, 0xdd, 0xc6, 0x8f, 0xa9, 0xdb, 0x78, 0xff, 0x24,
	0xe1, 0x6a, 0xbf, 0x0c, 0xfd, 0x08, 0x32, 0x12, 0x57, 0x19, 0xcb, 0x25, 0x87, 0xfc, 0xb0, 0xfa,
	0x40, 0x04, 0xa4, 0x29, 0x2d, 0x8a, 0xed, 0x31, 0x95, 0x19, 0x39, 0x3c, 0xf5, 0x56, 0x55, 0xc1,
	0x64, 0xd8, 0xd4, 0x7f, 0xef, 0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x83, 0x4b, 0xe3,
	0x8f, 0x04, 0x00, 0x00,
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
	// Params queries all the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Errors queries all the errors for a given contract.
	Errors(ctx context.Context, in *QueryErrorsRequest, opts ...grpc.CallOption) (*QueryErrorsResponse, error)
	// IsSubscribed queries if a contract is subscribed to sudo error callbacks.
	IsSubscribed(ctx context.Context, in *QueryIsSubscribedRequest, opts ...grpc.CallOption) (*QueryIsSubscribedResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/archway.cwerrors.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Errors(ctx context.Context, in *QueryErrorsRequest, opts ...grpc.CallOption) (*QueryErrorsResponse, error) {
	out := new(QueryErrorsResponse)
	err := c.cc.Invoke(ctx, "/archway.cwerrors.v1.Query/Errors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsSubscribed(ctx context.Context, in *QueryIsSubscribedRequest, opts ...grpc.CallOption) (*QueryIsSubscribedResponse, error) {
	out := new(QueryIsSubscribedResponse)
	err := c.cc.Invoke(ctx, "/archway.cwerrors.v1.Query/IsSubscribed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries all the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Errors queries all the errors for a given contract.
	Errors(context.Context, *QueryErrorsRequest) (*QueryErrorsResponse, error)
	// IsSubscribed queries if a contract is subscribed to sudo error callbacks.
	IsSubscribed(context.Context, *QueryIsSubscribedRequest) (*QueryIsSubscribedResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Errors(ctx context.Context, req *QueryErrorsRequest) (*QueryErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Errors not implemented")
}
func (*UnimplementedQueryServer) IsSubscribed(ctx context.Context, req *QueryIsSubscribedRequest) (*QueryIsSubscribedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSubscribed not implemented")
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
		FullMethod: "/archway.cwerrors.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Errors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Errors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.cwerrors.v1.Query/Errors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Errors(ctx, req.(*QueryErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IsSubscribed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIsSubscribedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IsSubscribed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.cwerrors.v1.Query/IsSubscribed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IsSubscribed(ctx, req.(*QueryIsSubscribedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "archway.cwerrors.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Errors",
			Handler:    _Query_Errors_Handler,
		},
		{
			MethodName: "IsSubscribed",
			Handler:    _Query_IsSubscribed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archway/cwerrors/v1/query.proto",
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

func (m *QueryErrorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryErrorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryErrorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryErrorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryErrorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryErrorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Errors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryIsSubscribedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsSubscribedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsSubscribedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsSubscribedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsSubscribedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsSubscribedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubscriptionValidTill != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SubscriptionValidTill))
		i--
		dAtA[i] = 0x10
	}
	if m.Subscribed {
		i--
		if m.Subscribed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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

func (m *QueryErrorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryErrorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryIsSubscribedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIsSubscribedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribed {
		n += 2
	}
	if m.SubscriptionValidTill != 0 {
		n += 1 + sovQuery(uint64(m.SubscriptionValidTill))
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
func (m *QueryErrorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryErrorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryErrorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryErrorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryErrorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryErrorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
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
			m.Errors = append(m.Errors, SudoError{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryIsSubscribedRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIsSubscribedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsSubscribedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIsSubscribedResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIsSubscribedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsSubscribedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribed = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionValidTill", wireType)
			}
			m.SubscriptionValidTill = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionValidTill |= int64(b&0x7F) << shift
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
