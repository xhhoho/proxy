// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/router/v3/router.proto

package envoy_extensions_filters_http_router_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/accesslog/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// [#next-free-field: 7]
type Router struct {
	// Whether the router generates dynamic cluster statistics. Defaults to
	// true. Can be disabled in high performance scenarios.
	DynamicStats *wrappers.BoolValue `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats,proto3" json:"dynamic_stats,omitempty"`
	// Whether to start a child span for egress routed calls. This can be
	// useful in scenarios where other filters (auth, ratelimit, etc.) make
	// outbound calls and have child spans rooted at the same ingress
	// parent. Defaults to false.
	StartChildSpan bool `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan,proto3" json:"start_child_span,omitempty"`
	// Configuration for HTTP upstream logs emitted by the router. Upstream logs
	// are configured in the same way as access logs, but each log entry represents
	// an upstream request. Presuming retries are configured, multiple upstream
	// requests may be made for each downstream (inbound) request.
	UpstreamLog []*v3.AccessLog `protobuf:"bytes,3,rep,name=upstream_log,json=upstreamLog,proto3" json:"upstream_log,omitempty"`
	// Do not add any additional *x-envoy-* headers to requests or responses. This
	// only affects the :ref:`router filter generated *x-envoy-* headers
	// <config_http_filters_router_headers_set>`, other Envoy filters and the HTTP
	// connection manager may continue to set *x-envoy-* headers.
	SuppressEnvoyHeaders bool `protobuf:"varint,4,opt,name=suppress_envoy_headers,json=suppressEnvoyHeaders,proto3" json:"suppress_envoy_headers,omitempty"`
	// Specifies a list of HTTP headers to strictly validate. Envoy will reject a
	// request and respond with HTTP status 400 if the request contains an invalid
	// value for any of the headers listed in this field. Strict header checking
	// is only supported for the following headers:
	//
	// Value must be a ','-delimited list (i.e. no spaces) of supported retry
	// policy values:
	//
	// * :ref:`config_http_filters_router_x-envoy-retry-grpc-on`
	// * :ref:`config_http_filters_router_x-envoy-retry-on`
	//
	// Value must be an integer:
	//
	// * :ref:`config_http_filters_router_x-envoy-max-retries`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-timeout-ms`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-per-try-timeout-ms`
	StrictCheckHeaders []string `protobuf:"bytes,5,rep,name=strict_check_headers,json=strictCheckHeaders,proto3" json:"strict_check_headers,omitempty"`
	// If not set, ingress Envoy will ignore
	// :ref:`config_http_filters_router_x-envoy-expected-rq-timeout-ms` header, populated by egress
	// Envoy, when deriving timeout for upstream cluster.
	RespectExpectedRqTimeout bool     `protobuf:"varint,6,opt,name=respect_expected_rq_timeout,json=respectExpectedRqTimeout,proto3" json:"respect_expected_rq_timeout,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc46a0c2ce4e43c, []int{0}
}

func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetDynamicStats() *wrappers.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func (m *Router) GetUpstreamLog() []*v3.AccessLog {
	if m != nil {
		return m.UpstreamLog
	}
	return nil
}

func (m *Router) GetSuppressEnvoyHeaders() bool {
	if m != nil {
		return m.SuppressEnvoyHeaders
	}
	return false
}

func (m *Router) GetStrictCheckHeaders() []string {
	if m != nil {
		return m.StrictCheckHeaders
	}
	return nil
}

func (m *Router) GetRespectExpectedRqTimeout() bool {
	if m != nil {
		return m.RespectExpectedRqTimeout
	}
	return false
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.extensions.filters.http.router.v3.Router")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/router/v3/router.proto", fileDescriptor_4dc46a0c2ce4e43c)
}

var fileDescriptor_4dc46a0c2ce4e43c = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x67, 0xb6, 0x6b, 0xd1, 0x74, 0x95, 0x12, 0x57, 0x0d, 0x15, 0x4b, 0x5d, 0x44, 0xbb, 0x48,
	0x26, 0xd2, 0xae, 0x17, 0x41, 0xc4, 0x2e, 0x8b, 0x1e, 0x16, 0x59, 0xb2, 0xe2, 0x35, 0xcc, 0x26,
	0xaf, 0xe9, 0x60, 0x3a, 0x33, 0x9d, 0x99, 0xc4, 0xf6, 0xe6, 0x49, 0x14, 0x3c, 0x88, 0x78, 0xf2,
	0x3b, 0xf8, 0x05, 0xbc, 0x8b, 0x5e, 0xfd, 0x3a, 0x9e, 0x64, 0x66, 0x92, 0x5d, 0x8b, 0x1e, 0x3c,
	0x65, 0xe6, 0xfd, 0xfe, 0xbc, 0x5f, 0x5e, 0x5e, 0xbc, 0x3d, 0x60, 0x15, 0x5f, 0x45, 0xb0, 0xd4,
	0xc0, 0x14, 0xe5, 0x4c, 0x45, 0x53, 0x5a, 0x68, 0x90, 0x2a, 0x9a, 0x69, 0x2d, 0x22, 0xc9, 0x4b,
	0x0d, 0x32, 0xaa, 0xc6, 0xf5, 0x09, 0x0b, 0xc9, 0x35, 0xf7, 0xef, 0x58, 0x15, 0x3e, 0x53, 0xe1,
	0x5a, 0x85, 0x8d, 0x0a, 0xd7, 0xdc, 0x6a, 0xdc, 0xdb, 0x75, 0xf6, 0x29, 0x67, 0x53, 0x9a, 0x47,
	0x24, 0x4d, 0x41, 0xa9, 0x82, 0xe7, 0xc6, 0xf0, 0xf4, 0xe2, 0x3c, 0x7b, 0xfd, 0x9c, 0xf3, 0xbc,
	0x80, 0xc8, 0xde, 0x4e, 0xca, 0x69, 0xf4, 0x4a, 0x12, 0x21, 0x8c, 0xa7, 0xc3, 0x6f, 0x94, 0x99,
	0x20, 0x11, 0x61, 0x8c, 0x6b, 0xa2, 0x6d, 0x52, 0xa5, 0x89, 0x2e, 0x1b, 0xf8, 0xe6, 0x5f, 0x70,
	0x05, 0xd2, 0x64, 0xa3, 0xac, 0xe9, 0x70, 0xad, 0x22, 0x05, 0xcd, 0x88, 0x86, 0xa8, 0x39, 0x38,
	0x60, 0xe7, 0xcb, 0xa6, 0xd7, 0x8e, 0x6d, 0x66, 0xff, 0x91, 0x77, 0x31, 0x5b, 0x31, 0x32, 0xa7,
	0x69, 0x62, 0xec, 0x55, 0x80, 0x06, 0x68, 0xd8, 0x19, 0xf5, 0xb0, 0x4b, 0x87, 0x9b, 0x74, 0x78,
	0xc2, 0x79, 0xf1, 0x82, 0x14, 0x25, 0xc4, 0x5b, 0xb5, 0xe0, 0xd8, 0xf0, 0xfd, 0xa1, 0xd7, 0x55,
	0x9a, 0x48, 0x9d, 0xa4, 0x33, 0x5a, 0x64, 0x89, 0x12, 0x84, 0x05, 0x1b, 0x03, 0x34, 0x3c, 0x1f,
	0x5f, 0xb2, 0xf5, 0x7d, 0x53, 0x3e, 0x16, 0x84, 0xf9, 0x4f, 0xbc, 0xad, 0x52, 0x28, 0x2d, 0x81,
	0xcc, 0x93, 0x82, 0xe7, 0x41, 0x6b, 0xd0, 0x1a, 0x76, 0x46, 0xb7, 0xb0, 0x9b, 0xad, 0x1b, 0x19,
	0x3e, 0x9b, 0x52, 0x35, 0xc6, 0x8f, 0xed, 0xe5, 0x90, 0xe7, 0x71, 0xa7, 0x51, 0x1e, 0xf2, 0xdc,
	0xdf, 0xf3, 0xae, 0xaa, 0x52, 0x08, 0x09, 0x4a, 0x25, 0x56, 0x9c, 0xcc, 0x80, 0x64, 0x20, 0x55,
	0xb0, 0x69, 0x1b, 0x6f, 0x37, 0xe8, 0x81, 0x01, 0x9f, 0x3a, 0xcc, 0xff, 0x8e, 0xbc, 0x6d, 0xa5,
	0x25, 0x4d, 0x4d, 0x54, 0x48, 0x5f, 0x9e, 0x8a, 0xce, 0x0d, 0x5a, 0xc3, 0x0b, 0x93, 0x4f, 0xe8,
	0xd7, 0xe4, 0x03, 0xfa, 0x88, 0xde, 0xa3, 0x9d, 0x77, 0x48, 0xbe, 0x41, 0x71, 0x7f, 0x19, 0x5a,
	0xf3, 0xb0, 0x69, 0x1d, 0xca, 0x45, 0xa8, 0xe9, 0x1c, 0x78, 0xa9, 0xc3, 0xb9, 0x8a, 0x6f, 0xff,
	0x0b, 0x17, 0x20, 0x43, 0x2d, 0x57, 0x7f, 0xf2, 0x2e, 0x37, 0xbc, 0x39, 0x59, 0x86, 0x12, 0xb4,
	0xa4, 0xa0, 0xe2, 0x2b, 0x4d, 0xd1, 0x14, 0x56, 0x61, 0x2e, 0x45, 0x1a, 0x72, 0x16, 0x77, 0xd7,
	0xcb, 0x9c, 0xc5, 0xbe, 0x8b, 0xbc, 0x6f, 0x12, 0x37, 0x6f, 0xf2, 0xd0, 0xbb, 0x2e, 0x41, 0x09,
	0x48, 0x75, 0x02, 0x4b, 0xf3, 0x80, 0x2c, 0x91, 0x8b, 0xa4, 0xee, 0x19, 0xb4, 0xed, 0x10, 0x82,
	0x9a, 0x72, 0x50, 0x33, 0xe2, 0xc5, 0x73, 0x87, 0x3f, 0xb8, 0xf7, 0xf9, 0xdb, 0xdb, 0xfe, 0x5d,
	0x6f, 0x77, 0x6d, 0xee, 0x6e, 0x9f, 0xd7, 0xd7, 0x79, 0x84, 0xdd, 0x92, 0x4c, 0x9e, 0x7d, 0x7d,
	0xfd, 0xe3, 0x67, 0x7b, 0xa3, 0xbb, 0xe1, 0xdd, 0xa7, 0xdc, 0x7d, 0x2f, 0x21, 0xf9, 0x72, 0x85,
	0xff, 0xf3, 0xb7, 0x98, 0x74, 0x9c, 0xd1, 0x91, 0x59, 0xa6, 0x23, 0x74, 0xd2, 0xb6, 0x5b, 0x35,
	0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa6, 0x28, 0x8f, 0x8d, 0x03, 0x00, 0x00,
}
