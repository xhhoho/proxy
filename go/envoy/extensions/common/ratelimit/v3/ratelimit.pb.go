// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/common/ratelimit/v3/ratelimit.proto

package envoy_extensions_common_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	// Descriptor entries.
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_584c20cb6057baa4, []int{0}
}

func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor.Unmarshal(m, b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor.Size(m)
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_584c20cb6057baa4, []int{0, 0}
}

func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Unmarshal(m, b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Size(m)
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.extensions.common.ratelimit.v3.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.extensions.common.ratelimit.v3.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/ratelimit/v3/ratelimit.proto", fileDescriptor_584c20cb6057baa4)
}

var fileDescriptor_584c20cb6057baa4 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xc9, 0xf4, 0xeb, 0x9f, 0x2f, 0x82, 0xc8, 0xb8, 0xb0, 0x16, 0x2a, 0x55, 0x5c, 0x14,
	0x91, 0x44, 0xa7, 0x82, 0xd0, 0x8d, 0x10, 0x74, 0x27, 0x52, 0xe6, 0x0d, 0x62, 0x7b, 0x91, 0x60,
	0x9b, 0x0c, 0xc9, 0x6d, 0xe8, 0xb8, 0x72, 0xe9, 0xda, 0xa5, 0x6f, 0xe0, 0x2b, 0xb8, 0x17, 0xdc,
	0xfa, 0x3a, 0x5d, 0xc9, 0x4c, 0x2c, 0x55, 0x54, 0xe8, 0x6e, 0xee, 0x3d, 0xf7, 0xfc, 0xe6, 0x70,
	0x42, 0x4f, 0x40, 0x7b, 0x93, 0x73, 0x98, 0x21, 0x68, 0xa7, 0x8c, 0x76, 0x7c, 0x68, 0x26, 0x13,
	0xa3, 0xb9, 0x95, 0x08, 0x63, 0x35, 0x51, 0xc8, 0x7d, 0x6f, 0x39, 0xb0, 0xcc, 0x1a, 0x34, 0xf1,
	0x7e, 0xe9, 0x62, 0x4b, 0x17, 0x0b, 0x2e, 0xb6, 0x3c, 0xf4, 0xbd, 0x56, 0x7b, 0x3a, 0xca, 0x24,
	0x97, 0x5a, 0x1b, 0x94, 0x58, 0xb2, 0x1d, 0x4a, 0x9c, 0xba, 0x00, 0x69, 0xed, 0xfe, 0x90, 0x3d,
	0xd8, 0x82, 0xa6, 0xf4, 0xcd, 0xe7, 0xc9, 0x96, 0x97, 0x63, 0x35, 0x92, 0x08, 0x7c, 0xf1, 0x11,
	0x84, 0xbd, 0xe7, 0x88, 0x6e, 0xa6, 0x12, 0xe1, 0xb2, 0xf8, 0xd7, 0x39, 0xb8, 0xa1, 0x55, 0x19,
	0x1a, 0x1b, 0x0f, 0x69, 0x1d, 0x34, 0x5a, 0x05, 0xae, 0x49, 0x3a, 0x95, 0xee, 0x5a, 0x72, 0xc6,
	0x56, 0x89, 0xca, 0x7e, 0x61, 0xb1, 0x0b, 0x8d, 0x36, 0x17, 0x8d, 0xb9, 0xa8, 0x3e, 0x92, 0xa8,
	0x41, 0xd2, 0x05, 0xb9, 0x75, 0x47, 0xab, 0xa5, 0x16, 0x6f, 0xd3, 0xca, 0x2d, 0xe4, 0x4d, 0xd2,
	0x21, 0xdd, 0xff, 0xa2, 0x3e, 0x17, 0xff, 0x6c, 0xd4, 0x21, 0x69, 0xb1, 0x8b, 0xdb, 0xb4, 0xea,
	0xe5, 0x78, 0x0a, 0xcd, 0xe8, 0xbb, 0x18, 0xb6, 0xfd, 0xd3, 0xa7, 0xd7, 0x87, 0x9d, 0x84, 0x1e,
	0x85, 0x70, 0x32, 0x53, 0xcc, 0x27, 0x5f, 0x12, 0xfd, 0x19, 0xa7, 0x7f, 0x5c, 0x18, 0x0f, 0xe9,
	0xc1, 0xea, 0x46, 0x71, 0xf5, 0x72, 0xff, 0xf6, 0x5e, 0x8b, 0x36, 0x22, 0x9a, 0x28, 0x13, 0xea,
	0xc8, 0xac, 0x99, 0xe5, 0x2b, 0x35, 0x23, 0xd6, 0xd3, 0xc5, 0x34, 0x28, 0x9a, 0x1f, 0x90, 0xeb,
	0x5a, 0xf9, 0x04, 0xbd, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x45, 0x24, 0x54, 0x3b, 0x02,
	0x00, 0x00,
}
