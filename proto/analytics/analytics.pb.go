// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/analytics/analytics.proto

package analytics

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PriceABC struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ABC                  string   `protobuf:"bytes,2,opt,name=ABC,proto3" json:"ABC,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceABC) Reset()         { *m = PriceABC{} }
func (m *PriceABC) String() string { return proto.CompactTextString(m) }
func (*PriceABC) ProtoMessage()    {}
func (*PriceABC) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb2a98f65dc1200, []int{0}
}

func (m *PriceABC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceABC.Unmarshal(m, b)
}
func (m *PriceABC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceABC.Marshal(b, m, deterministic)
}
func (m *PriceABC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceABC.Merge(m, src)
}
func (m *PriceABC) XXX_Size() int {
	return xxx_messageInfo_PriceABC.Size(m)
}
func (m *PriceABC) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceABC.DiscardUnknown(m)
}

var xxx_messageInfo_PriceABC proto.InternalMessageInfo

func (m *PriceABC) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PriceABC) GetABC() string {
	if m != nil {
		return m.ABC
	}
	return ""
}

type GetABCByPricesRequest struct {
	Year                 int32    `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetABCByPricesRequest) Reset()         { *m = GetABCByPricesRequest{} }
func (m *GetABCByPricesRequest) String() string { return proto.CompactTextString(m) }
func (*GetABCByPricesRequest) ProtoMessage()    {}
func (*GetABCByPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb2a98f65dc1200, []int{1}
}

func (m *GetABCByPricesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetABCByPricesRequest.Unmarshal(m, b)
}
func (m *GetABCByPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetABCByPricesRequest.Marshal(b, m, deterministic)
}
func (m *GetABCByPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetABCByPricesRequest.Merge(m, src)
}
func (m *GetABCByPricesRequest) XXX_Size() int {
	return xxx_messageInfo_GetABCByPricesRequest.Size(m)
}
func (m *GetABCByPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetABCByPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetABCByPricesRequest proto.InternalMessageInfo

func (m *GetABCByPricesRequest) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

type GetABCByPricesResponse struct {
	PriceABC             []*PriceABC `protobuf:"bytes,1,rep,name=priceABC,proto3" json:"priceABC,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetABCByPricesResponse) Reset()         { *m = GetABCByPricesResponse{} }
func (m *GetABCByPricesResponse) String() string { return proto.CompactTextString(m) }
func (*GetABCByPricesResponse) ProtoMessage()    {}
func (*GetABCByPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb2a98f65dc1200, []int{2}
}

func (m *GetABCByPricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetABCByPricesResponse.Unmarshal(m, b)
}
func (m *GetABCByPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetABCByPricesResponse.Marshal(b, m, deterministic)
}
func (m *GetABCByPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetABCByPricesResponse.Merge(m, src)
}
func (m *GetABCByPricesResponse) XXX_Size() int {
	return xxx_messageInfo_GetABCByPricesResponse.Size(m)
}
func (m *GetABCByPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetABCByPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetABCByPricesResponse proto.InternalMessageInfo

func (m *GetABCByPricesResponse) GetPriceABC() []*PriceABC {
	if m != nil {
		return m.PriceABC
	}
	return nil
}

func init() {
	proto.RegisterType((*PriceABC)(nil), "analytics.PriceABC")
	proto.RegisterType((*GetABCByPricesRequest)(nil), "analytics.GetABCByPricesRequest")
	proto.RegisterType((*GetABCByPricesResponse)(nil), "analytics.GetABCByPricesResponse")
}

func init() { proto.RegisterFile("proto/analytics/analytics.proto", fileDescriptor_7fb2a98f65dc1200) }

var fileDescriptor_7fb2a98f65dc1200 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcc, 0x4b, 0xcc, 0xa9, 0x2c, 0xc9, 0x4c, 0x2e, 0x46, 0xb0, 0xf4, 0xc0, 0x32,
	0x42, 0x9c, 0x70, 0x01, 0x25, 0x1d, 0x2e, 0x8e, 0x80, 0xa2, 0xcc, 0xe4, 0x54, 0x47, 0x27, 0x67,
	0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0xa6, 0xcc, 0x14,
	0x21, 0x01, 0x2e, 0x66, 0x47, 0x27, 0x67, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53,
	0x49, 0x9b, 0x4b, 0xd4, 0x3d, 0xb5, 0xc4, 0xd1, 0xc9, 0xd9, 0xa9, 0x12, 0xac, 0xab, 0x38, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0xa5, 0x32, 0x35, 0xb1, 0x08, 0xac, 0x99,
	0x35, 0x08, 0xcc, 0x56, 0xf2, 0xe4, 0x12, 0x43, 0x57, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a,
	0xa4, 0xcf, 0xc5, 0x51, 0x00, 0xb5, 0x54, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x58, 0x0f,
	0xe1, 0x46, 0x98, 0x7b, 0x82, 0xe0, 0x8a, 0x8c, 0xb2, 0xb9, 0x04, 0x1c, 0x61, 0xf2, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xe1, 0x5c, 0x7c, 0xa8, 0xc6, 0x0b, 0x29, 0x20, 0x19, 0x82,
	0xd5, 0x99, 0x52, 0x8a, 0x78, 0x54, 0x40, 0xdc, 0xa6, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x24, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x8b, 0x7a, 0x69, 0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	GetABCByPrices(ctx context.Context, in *GetABCByPricesRequest, opts ...grpc.CallOption) (*GetABCByPricesResponse, error)
}

type analyticsServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnalyticsServiceClient(cc *grpc.ClientConn) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) GetABCByPrices(ctx context.Context, in *GetABCByPricesRequest, opts ...grpc.CallOption) (*GetABCByPricesResponse, error) {
	out := new(GetABCByPricesResponse)
	err := c.cc.Invoke(ctx, "/analytics.AnalyticsService/GetABCByPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
type AnalyticsServiceServer interface {
	GetABCByPrices(context.Context, *GetABCByPricesRequest) (*GetABCByPricesResponse, error)
}

// UnimplementedAnalyticsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (*UnimplementedAnalyticsServiceServer) GetABCByPrices(ctx context.Context, req *GetABCByPricesRequest) (*GetABCByPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetABCByPrices not implemented")
}

func RegisterAnalyticsServiceServer(s *grpc.Server, srv AnalyticsServiceServer) {
	s.RegisterService(&_AnalyticsService_serviceDesc, srv)
}

func _AnalyticsService_GetABCByPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetABCByPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetABCByPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.AnalyticsService/GetABCByPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetABCByPrices(ctx, req.(*GetABCByPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnalyticsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetABCByPrices",
			Handler:    _AnalyticsService_GetABCByPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/analytics/analytics.proto",
}
