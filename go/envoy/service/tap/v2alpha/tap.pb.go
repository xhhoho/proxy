// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tap.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2alpha "github.com/cilium/proxy/go/envoy/data/tap/v2alpha"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// [#not-implemented-hide:] Stream message for the Tap API. Envoy will open a stream to the server
// and stream taps without ever expecting a response.
type StreamTapsRequest struct {
	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamTapsRequest_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// The trace id. this can be used to merge together a streaming trace. Note that the trace_id
	// is not guaranteed to be spatially or temporally unique.
	TraceId uint64 `protobuf:"varint,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The trace data.
	Trace                *v2alpha.TraceWrapper `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StreamTapsRequest) Reset()         { *m = StreamTapsRequest{} }
func (m *StreamTapsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamTapsRequest) ProtoMessage()    {}
func (*StreamTapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e8ad5aa2b63d8d, []int{0}
}

func (m *StreamTapsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsRequest.Unmarshal(m, b)
}
func (m *StreamTapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsRequest.Marshal(b, m, deterministic)
}
func (m *StreamTapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsRequest.Merge(m, src)
}
func (m *StreamTapsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamTapsRequest.Size(m)
}
func (m *StreamTapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsRequest proto.InternalMessageInfo

func (m *StreamTapsRequest) GetIdentifier() *StreamTapsRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTapsRequest) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *StreamTapsRequest) GetTrace() *v2alpha.TraceWrapper {
	if m != nil {
		return m.Trace
	}
	return nil
}

type StreamTapsRequest_Identifier struct {
	// The node sending taps over the stream.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The opaque identifier that was set in the :ref:`output config
	// <envoy_api_field_service.tap.v2alpha.StreamingGrpcSink.tap_id>`.
	TapId                string   `protobuf:"bytes,2,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTapsRequest_Identifier) Reset()         { *m = StreamTapsRequest_Identifier{} }
func (m *StreamTapsRequest_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTapsRequest_Identifier) ProtoMessage()    {}
func (*StreamTapsRequest_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e8ad5aa2b63d8d, []int{0, 0}
}

func (m *StreamTapsRequest_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Unmarshal(m, b)
}
func (m *StreamTapsRequest_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTapsRequest_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsRequest_Identifier.Merge(m, src)
}
func (m *StreamTapsRequest_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Size(m)
}
func (m *StreamTapsRequest_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsRequest_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsRequest_Identifier proto.InternalMessageInfo

func (m *StreamTapsRequest_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamTapsRequest_Identifier) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

// [#not-implemented-hide:]
type StreamTapsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTapsResponse) Reset()         { *m = StreamTapsResponse{} }
func (m *StreamTapsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTapsResponse) ProtoMessage()    {}
func (*StreamTapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e8ad5aa2b63d8d, []int{1}
}

func (m *StreamTapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsResponse.Unmarshal(m, b)
}
func (m *StreamTapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsResponse.Marshal(b, m, deterministic)
}
func (m *StreamTapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsResponse.Merge(m, src)
}
func (m *StreamTapsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTapsResponse.Size(m)
}
func (m *StreamTapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamTapsRequest)(nil), "envoy.service.tap.v2alpha.StreamTapsRequest")
	proto.RegisterType((*StreamTapsRequest_Identifier)(nil), "envoy.service.tap.v2alpha.StreamTapsRequest.Identifier")
	proto.RegisterType((*StreamTapsResponse)(nil), "envoy.service.tap.v2alpha.StreamTapsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/tap.proto", fileDescriptor_e1e8ad5aa2b63d8d)
}

var fileDescriptor_e1e8ad5aa2b63d8d = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xaa, 0x13, 0x31,
	0x14, 0x86, 0xcd, 0x78, 0x5b, 0x6b, 0x04, 0xb9, 0x06, 0xe5, 0xf6, 0x0e, 0x0a, 0xa5, 0x16, 0xec,
	0x42, 0x13, 0x18, 0x11, 0xc1, 0x95, 0xcc, 0xae, 0x1b, 0x29, 0xd3, 0x81, 0x82, 0x1b, 0x39, 0x6d,
	0x8e, 0x18, 0x6c, 0x93, 0x98, 0xa4, 0xa3, 0x5d, 0xe9, 0x52, 0x7c, 0x09, 0xdf, 0xc3, 0x27, 0x70,
	0xeb, 0xeb, 0xb8, 0x92, 0x99, 0x4c, 0x6d, 0x8b, 0x14, 0xbc, 0xbb, 0x33, 0x9c, 0xef, 0xff, 0xff,
	0x73, 0xce, 0x84, 0x3e, 0x44, 0x5d, 0x99, 0xad, 0xf0, 0xe8, 0x2a, 0xb5, 0x44, 0x11, 0xc0, 0x8a,
	0x2a, 0x83, 0x95, 0x7d, 0x07, 0x75, 0xcd, 0xad, 0x33, 0xc1, 0xb0, 0xcb, 0x06, 0xe2, 0x2d, 0xc4,
	0xeb, 0x46, 0x0b, 0xa5, 0xf7, 0xa3, 0x1e, 0xac, 0x12, 0x55, 0x26, 0x96, 0xc6, 0xa1, 0x58, 0x80,
	0xc7, 0x28, 0x4c, 0x47, 0xb1, 0x2b, 0x21, 0xc0, 0x91, 0xf5, 0x47, 0x07, 0xd6, 0xa2, 0x6b, 0xa9,
	0x07, 0x1b, 0x69, 0x41, 0x80, 0xd6, 0x26, 0x40, 0x50, 0x46, 0x7b, 0xe1, 0x03, 0x84, 0x8d, 0x6f,
	0xdb, 0x17, 0x15, 0xac, 0x94, 0x84, 0x80, 0x62, 0x57, 0xc4, 0xc6, 0xf0, 0x7b, 0x42, 0xef, 0xcc,
	0x82, 0x43, 0x58, 0x97, 0x60, 0x7d, 0x81, 0x1f, 0x36, 0xe8, 0x03, 0x9b, 0x53, 0xaa, 0x24, 0xea,
	0xa0, 0xde, 0x2a, 0x74, 0x7d, 0x32, 0x20, 0xe3, 0x5b, 0xd9, 0x73, 0x7e, 0x72, 0x03, 0xfe, 0x8f,
	0x03, 0x9f, 0xfc, 0x95, 0x17, 0x07, 0x56, 0xec, 0x92, 0xf6, 0x82, 0x83, 0x25, 0xbe, 0x51, 0xb2,
	0x9f, 0x0c, 0xc8, 0xf8, 0xac, 0xb8, 0xd1, 0x7c, 0x4f, 0x24, 0x7b, 0x41, 0x3b, 0x4d, 0xd9, 0xbf,
	0xde, 0xc4, 0x8d, 0xda, 0xb8, 0x7a, 0xef, 0xa3, 0xac, 0xb2, 0x86, 0xe6, 0x71, 0xf9, 0x22, 0x4a,
	0xd2, 0xd7, 0x94, 0xee, 0x03, 0xd9, 0x33, 0x7a, 0xa6, 0x8d, 0xc4, 0x76, 0xee, 0x8b, 0xd6, 0x08,
	0xac, 0xe2, 0x55, 0xc6, 0xeb, 0xf3, 0xf2, 0x57, 0x46, 0x62, 0xde, 0xfb, 0x9d, 0x77, 0xbe, 0x91,
	0xe4, 0x9c, 0x14, 0x0d, 0xce, 0xee, 0xd1, 0x6e, 0x00, 0xbb, 0x9b, 0xec, 0x66, 0xd1, 0x09, 0x60,
	0x27, 0x72, 0x78, 0x97, 0xb2, 0xc3, 0xf5, 0xbc, 0x35, 0xda, 0x63, 0xf6, 0x99, 0xde, 0x2e, 0xc1,
	0xce, 0x94, 0x7e, 0x3f, 0x8b, 0xf7, 0x60, 0x6b, 0x4a, 0xf7, 0x1c, 0x7b, 0x7c, 0x95, 0x6b, 0xa5,
	0x4f, 0xfe, 0x93, 0x8e, 0xe1, 0xc3, 0x6b, 0x63, 0x92, 0xbf, 0xfc, 0xf1, 0xe5, 0xe7, 0xaf, 0x6e,
	0x72, 0x4e, 0xe8, 0x23, 0x65, 0xa2, 0xdc, 0x3a, 0xf3, 0x69, 0x7b, 0xda, 0x29, 0xef, 0x95, 0x60,
	0xa7, 0xf5, 0x5f, 0x9f, 0x92, 0xaf, 0x84, 0x2c, 0xba, 0xcd, 0x0b, 0x78, 0xfa, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xca, 0xa8, 0xbd, 0xbf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapSinkServiceClient is the client API for TapSinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapSinkServiceClient interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error)
}

type tapSinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapSinkServiceClient(cc *grpc.ClientConn) TapSinkServiceClient {
	return &tapSinkServiceClient{cc}
}

func (c *tapSinkServiceClient) StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapSinkService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapSinkService/StreamTaps", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapSinkServiceStreamTapsClient{stream}
	return x, nil
}

type TapSinkService_StreamTapsClient interface {
	Send(*StreamTapsRequest) error
	CloseAndRecv() (*StreamTapsResponse, error)
	grpc.ClientStream
}

type tapSinkServiceStreamTapsClient struct {
	grpc.ClientStream
}

func (x *tapSinkServiceStreamTapsClient) Send(m *StreamTapsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsClient) CloseAndRecv() (*StreamTapsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTapsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TapSinkServiceServer is the server API for TapSinkService service.
type TapSinkServiceServer interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(TapSinkService_StreamTapsServer) error
}

// UnimplementedTapSinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapSinkServiceServer struct {
}

func (*UnimplementedTapSinkServiceServer) StreamTaps(srv TapSinkService_StreamTapsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTaps not implemented")
}

func RegisterTapSinkServiceServer(s *grpc.Server, srv TapSinkServiceServer) {
	s.RegisterService(&_TapSinkService_serviceDesc, srv)
}

func _TapSinkService_StreamTaps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapSinkServiceServer).StreamTaps(&tapSinkServiceStreamTapsServer{stream})
}

type TapSinkService_StreamTapsServer interface {
	SendAndClose(*StreamTapsResponse) error
	Recv() (*StreamTapsRequest, error)
	grpc.ServerStream
}

type tapSinkServiceStreamTapsServer struct {
	grpc.ServerStream
}

func (x *tapSinkServiceStreamTapsServer) SendAndClose(m *StreamTapsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsServer) Recv() (*StreamTapsRequest, error) {
	m := new(StreamTapsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TapSinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapSinkService",
	HandlerType: (*TapSinkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTaps",
			Handler:       _TapSinkService_StreamTaps_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tap.proto",
}
