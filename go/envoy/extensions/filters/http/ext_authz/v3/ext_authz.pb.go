// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/ext_authz/v3/ext_authz.proto

package envoy_extensions_filters_http_ext_authz_v3

import (
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v32 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	v31 "github.com/cilium/proxy/go/envoy/type/v3"
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

// [#next-free-field: 11]
type ExtAuthz struct {
	// External authorization service configuration.
	//
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services isExtAuthz_Services `protobuf_oneof:"services"`
	//  Changes filter's behaviour on errors:
	//
	//  1. When set to true, the filter will *accept* client request even if the communication with
	//  the authorization service has failed, or if the authorization service has returned a HTTP 5xx
	//  error.
	//
	//  2. When set to false, ext-authz will *reject* client requests and return a *Forbidden*
	//  response if the communication with the authorization service has failed, or if the
	//  authorization service has returned a HTTP 5xx error.
	//
	// Note that errors can be *always* tracked in the :ref:`stats
	// <config_http_filters_ext_authz_stats>`.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Enables filter to buffer the client request body and send it within the authorization request.
	// A ``x-envoy-auth-partial-body: false|true`` metadata header will be added to the authorization
	// request message indicating if the body data is partial.
	WithRequestBody *BufferSettings `protobuf:"bytes,5,opt,name=with_request_body,json=withRequestBody,proto3" json:"with_request_body,omitempty"`
	// Clears route cache in order to allow the external authorization service to correctly affect
	// routing decisions. Filter clears all cached routes when:
	//
	// 1. The field is set to *true*.
	//
	// 2. The status returned from the authorization service is a HTTP 200 or gRPC 0.
	//
	// 3. At least one *authorization response header* is added to the client request, or is used for
	// altering another client request header.
	//
	ClearRouteCache bool `protobuf:"varint,6,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	// Sets the HTTP status that is returned to the client when there is a network error between the
	// filter and the authorization server. The default status is HTTP 403 Forbidden.
	StatusOnError *v31.HttpStatus `protobuf:"bytes,7,opt,name=status_on_error,json=statusOnError,proto3" json:"status_on_error,omitempty"`
	// Specifies a list of metadata namespaces whose values, if present, will be passed to the
	// ext_authz service as an opaque *protobuf::Struct*.
	//
	// For example, if the *jwt_authn* filter is used and :ref:`payload_in_metadata
	// <envoy_api_field_extensions.filters.http.jwt_authn.v3.JwtProvider.payload_in_metadata>` is set,
	// then the following will pass the jwt payload to the authorization server.
	//
	// .. code-block:: yaml
	//
	//    metadata_context_namespaces:
	//    - envoy.filters.http.jwt_authn
	//
	MetadataContextNamespaces []string `protobuf:"bytes,8,rep,name=metadata_context_namespaces,json=metadataContextNamespaces,proto3" json:"metadata_context_namespaces,omitempty"`
	// Specifies if the filter is enabled.
	//
	// If :ref:`runtime_key <envoy_api_field_config.core.v3.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests to filter.
	//
	// If this field is not specified, the filter will be enabled for all requests.
	FilterEnabled *v3.RuntimeFractionalPercent `protobuf:"bytes,9,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies if the peer certificate is sent to the external service.
	//
	// When this field is true, Envoy will include the peer X.509 certificate, if available, in the
	// :ref:`certificate<envoy_api_field_service.auth.v3.AttributeContext.Peer.certificate>`.
	IncludePeerCertificate bool     `protobuf:"varint,10,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *v3.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func (m *ExtAuthz) GetWithRequestBody() *BufferSettings {
	if m != nil {
		return m.WithRequestBody
	}
	return nil
}

func (m *ExtAuthz) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *ExtAuthz) GetStatusOnError() *v31.HttpStatus {
	if m != nil {
		return m.StatusOnError
	}
	return nil
}

func (m *ExtAuthz) GetMetadataContextNamespaces() []string {
	if m != nil {
		return m.MetadataContextNamespaces
	}
	return nil
}

func (m *ExtAuthz) GetFilterEnabled() *v3.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *ExtAuthz) GetIncludePeerCertificate() bool {
	if m != nil {
		return m.IncludePeerCertificate
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

// Configuration for buffering the request data.
type BufferSettings struct {
	// Sets the maximum size of a message body that the filter will hold in memory. Envoy will return
	// *HTTP 413* and will *not* initiate the authorization process when buffer reaches the number
	// set in this field. Note that this setting will have precedence over :ref:`failure_mode_allow
	// <envoy_api_field_extensions.filters.http.ext_authz.v3.ExtAuthz.failure_mode_allow>`.
	MaxRequestBytes uint32 `protobuf:"varint,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	// When this field is true, Envoy will buffer the message until *max_request_bytes* is reached.
	// The authorization request will be dispatched and no 413 HTTP error will be returned by the
	// filter.
	AllowPartialMessage  bool     `protobuf:"varint,2,opt,name=allow_partial_message,json=allowPartialMessage,proto3" json:"allow_partial_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{1}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

func (m *BufferSettings) GetAllowPartialMessage() bool {
	if m != nil {
		return m.AllowPartialMessage
	}
	return false
}

// HttpService is used for raw HTTP communication between the filter and the authorization service.
// When configured, the filter will parse the client request and use these attributes to call the
// authorization server. Depending on the response, the filter may reject or accept the client
// request. Note that in any of these events, metadata can be added, removed or overridden by the
// filter:
//
// *On authorization request*, a list of allowed request headers may be supplied. See
// :ref:`allowed_headers
// <envoy_api_field_extensions.filters.http.ext_authz.v3.AuthorizationRequest.allowed_headers>`
// for details. Additional headers metadata may be added to the authorization request. See
// :ref:`headers_to_add
// <envoy_api_field_extensions.filters.http.ext_authz.v3.AuthorizationRequest.headers_to_add>` for
// details.
//
// On authorization response status HTTP 200 OK, the filter will allow traffic to the upstream and
// additional headers metadata may be added to the original client request. See
// :ref:`allowed_upstream_headers
// <envoy_api_field_extensions.filters.http.ext_authz.v3.AuthorizationResponse.allowed_upstream_headers>`
// for details.
//
// On other authorization response statuses, the filter will not allow traffic. Additional headers
// metadata as well as body may be added to the client's response. See :ref:`allowed_client_headers
// <envoy_api_field_extensions.filters.http.ext_authz.v3.AuthorizationResponse.allowed_client_headers>`
// for details.
// [#next-free-field: 9]
type HttpService struct {
	// Sets the HTTP server URI which the authorization requests must be sent to.
	ServerUri *v3.HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	// Sets a prefix to the value of authorization request header *Path*.
	PathPrefix string `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// Settings used for controlling authorization request metadata.
	AuthorizationRequest *AuthorizationRequest `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	// Settings used for controlling authorization response metadata.
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{2}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *v3.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	// Authorization request will include the client request headers that have a correspondent match
	// in the :ref:`list <envoy_api_msg_type.matcher.v3.ListStringMatcher>`. Note that in addition to the
	// user's supplied matchers:
	//
	// 1. *Host*, *Method*, *Path* and *Content-Length* are automatically included to the list.
	//
	// 2. *Content-Length* will be set to 0 and the request to the authorization service will not have
	// a message body. However, the authorization request can include the buffered client request body
	// (controlled by :ref:`with_request_body
	// <envoy_api_field_extensions.filters.http.ext_authz.v3.ExtAuthz.with_request_body>` setting),
	// consequently the value of *Content-Length* of the authorization request reflects the size of
	// its payload size.
	//
	AllowedHeaders *v32.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	// Sets a list of headers that will be included to the request to authorization service. Note that
	// client request of the same key will be overridden.
	HeadersToAdd         []*v3.HeaderValue `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{3}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*v3.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	// When this :ref:`list <envoy_api_msg_type.matcher.v3.ListStringMatcher>` is set, authorization
	// response headers that have a correspondent match will be added to the original client request.
	// Note that coexistent headers will be overridden.
	AllowedUpstreamHeaders *v32.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	// When this :ref:`list <envoy_api_msg_type.matcher.v3.ListStringMatcher>`. is set, authorization
	// response headers that have a correspondent match will be added to the client's response. Note
	// that when this list is *not* set, all the authorization response headers, except *Authority
	// (Host)* will be in the response to the client. When a header is included in this list, *Path*,
	// *Status*, *Content-Length*, *WWWAuthenticate* and *Location* are automatically added.
	AllowedClientHeaders *v32.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{4}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

// Extra settings on a per virtualhost/route/weighted-cluster level.
type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{5}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

// Extra settings for the check request. You can use this to provide extra context for the
// external authorization server on specific virtual hosts \ routes. For example, adding a context
// extension on the virtual host level can give the ext-authz server information on what virtual
// host is used without needing to parse the host header. If CheckSettings is specified in multiple
// per-filter-configs, they will be merged in order, and the result will be used.
type CheckSettings struct {
	// Context extensions to set on the CheckRequest's
	// :ref:`AttributeContext.context_extensions<envoy_api_field_service.auth.v3.AttributeContext.context_extensions>`
	//
	// Merge semantics for this field are such that keys from more specific configs override.
	//
	// .. note::
	//
	//   These settings are only applied to a filter configured with a
	//   :ref:`grpc_service<envoy_api_field_extensions.filters.http.ext_authz.v3.ExtAuthz.grpc_service>`.
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{6}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.extensions.filters.http.ext_authz.v3.ExtAuthz")
	proto.RegisterType((*BufferSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3.BufferSettings")
	proto.RegisterType((*HttpService)(nil), "envoy.extensions.filters.http.ext_authz.v3.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.extensions.filters.http.ext_authz.v3.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.extensions.filters.http.ext_authz.v3.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.extensions.filters.http.ext_authz.v3.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.extensions.filters.http.ext_authz.v3.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ext_authz/v3/ext_authz.proto", fileDescriptor_9a074478ef6deb0a)
}

var fileDescriptor_9a074478ef6deb0a = []byte{
	// 1175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xef, 0xda, 0x4e, 0xba, 0x99, 0xd4, 0xb1, 0x33, 0xdf, 0x24, 0xdf, 0x6d, 0x51, 0x69, 0x6a,
	0x84, 0x88, 0x2a, 0xb4, 0x86, 0x9a, 0x96, 0xe0, 0xb6, 0x08, 0x3b, 0xa4, 0x8d, 0x22, 0x0a, 0x66,
	0x43, 0x38, 0x21, 0x56, 0x93, 0xdd, 0x67, 0x7b, 0xe8, 0x7a, 0x77, 0x99, 0x99, 0x75, 0xed, 0x5e,
	0xe0, 0x88, 0x38, 0x72, 0xe4, 0xc2, 0x9f, 0xc0, 0x8d, 0x03, 0x77, 0x24, 0xae, 0x5c, 0xf9, 0x3f,
	0x10, 0xa8, 0x27, 0x34, 0x3f, 0xd6, 0x3f, 0x5a, 0x23, 0xe2, 0x72, 0xdb, 0x79, 0x3f, 0x3e, 0xef,
	0xbd, 0xcf, 0x7b, 0xf3, 0x66, 0x51, 0x13, 0xe2, 0x61, 0x32, 0xae, 0xc3, 0x48, 0x40, 0xcc, 0x69,
	0x12, 0xf3, 0x7a, 0x97, 0x46, 0x02, 0x18, 0xaf, 0xf7, 0x85, 0x48, 0xa5, 0xdc, 0x27, 0x99, 0xe8,
	0x3f, 0xa9, 0x0f, 0x1b, 0xd3, 0x83, 0x9b, 0xb2, 0x44, 0x24, 0xf8, 0x86, 0xf2, 0x75, 0xa7, 0xbe,
	0xae, 0xf1, 0x75, 0xa5, 0xaf, 0x3b, 0x35, 0x1f, 0x36, 0xae, 0x5c, 0xd3, 0x71, 0x82, 0x24, 0xee,
	0xd2, 0x5e, 0x3d, 0x48, 0x18, 0x48, 0xc4, 0x33, 0xc2, 0x41, 0x83, 0x5d, 0x79, 0x6d, 0xa1, 0x41,
	0x8f, 0xa5, 0x81, 0xcf, 0x81, 0x0d, 0x69, 0x90, 0x1b, 0xbe, 0xb2, 0xd0, 0x50, 0x46, 0xf4, 0x33,
	0x46, 0x8d, 0x51, 0x4d, 0x1b, 0x89, 0x71, 0x0a, 0xf5, 0x01, 0x11, 0x41, 0x1f, 0x98, 0xb4, 0xe2,
	0x82, 0xd1, 0xb8, 0x67, 0x6c, 0xae, 0xcd, 0xd8, 0xe4, 0x08, 0x5c, 0x10, 0x91, 0xf1, 0xf9, 0x48,
	0x24, 0x8e, 0x13, 0x41, 0x84, 0x22, 0x27, 0x84, 0x94, 0x41, 0xa0, 0x0e, 0xc6, 0xe8, 0x6a, 0x16,
	0xa6, 0x64, 0xce, 0x66, 0x0e, 0xe3, 0xfa, 0x73, 0xea, 0x21, 0x30, 0x49, 0xd6, 0x34, 0x8f, 0xff,
	0x0f, 0x49, 0x44, 0x43, 0x22, 0xa0, 0x9e, 0x7f, 0x68, 0x45, 0xed, 0xf7, 0x15, 0x64, 0x1f, 0x8e,
	0x44, 0x4b, 0x72, 0x88, 0xef, 0xa3, 0x4b, 0xb3, 0x64, 0x38, 0xd6, 0xae, 0xb5, 0xb7, 0x7e, 0xf3,
	0xba, 0xab, 0x7b, 0xa0, 0xd9, 0x70, 0x25, 0x1b, 0xee, 0xb0, 0xe1, 0x3e, 0x60, 0x69, 0x70, 0xa2,
	0x0d, 0x8f, 0x2e, 0x78, 0xeb, 0xbd, 0xe9, 0x11, 0x7f, 0x86, 0x2e, 0xe9, 0x4a, 0x0d, 0x4e, 0x51,
	0xe1, 0xbc, 0xed, 0x9e, 0xbf, 0x97, 0xee, 0x91, 0x10, 0xe9, 0x0c, 0x7a, 0x7f, 0x7a, 0xc4, 0xaf,
	0x23, 0xdc, 0x25, 0x34, 0xca, 0x18, 0xf8, 0x83, 0x24, 0x04, 0x9f, 0x44, 0x51, 0xf2, 0xd8, 0x29,
	0xec, 0x5a, 0x7b, 0xb6, 0x57, 0x35, 0x9a, 0x87, 0x49, 0x08, 0x2d, 0x29, 0xc7, 0x5d, 0xb4, 0xf9,
	0x98, 0x8a, 0xbe, 0xcf, 0xe0, 0xcb, 0x0c, 0xb8, 0xf0, 0xcf, 0x92, 0x70, 0xec, 0xac, 0xa8, 0x84,
	0x9a, 0xcb, 0x24, 0xd4, 0xce, 0xba, 0x5d, 0x60, 0x27, 0x20, 0x04, 0x8d, 0x7b, 0xdc, 0xab, 0x48,
	0x50, 0x4f, 0x63, 0xb6, 0x93, 0x70, 0x8c, 0x6f, 0xa0, 0xcd, 0x20, 0x02, 0xc2, 0x7c, 0x96, 0x64,
	0x02, 0xfc, 0x80, 0x04, 0x7d, 0x70, 0x56, 0x55, 0x52, 0x15, 0xa5, 0xf0, 0xa4, 0xfc, 0x40, 0x8a,
	0x71, 0x0b, 0x55, 0x74, 0x03, 0xfd, 0x24, 0xf6, 0x81, 0xb1, 0x84, 0x39, 0x17, 0x55, 0x46, 0x97,
	0x4d, 0x46, 0x72, 0x5e, 0x26, 0x2c, 0x28, 0x4b, 0xaf, 0xac, 0x3d, 0x3e, 0x8a, 0x0f, 0xa5, 0x3d,
	0x7e, 0x17, 0xbd, 0x34, 0x00, 0x41, 0x42, 0x22, 0x88, 0x1f, 0x24, 0xb1, 0x90, 0xb9, 0xc6, 0x64,
	0x00, 0x3c, 0x25, 0x01, 0x70, 0xc7, 0xde, 0x2d, 0xee, 0xad, 0x79, 0x97, 0x73, 0x93, 0x03, 0x6d,
	0xf1, 0xe1, 0xc4, 0x00, 0x9f, 0xa2, 0x0d, 0x5d, 0xab, 0x0f, 0x31, 0x39, 0x8b, 0x20, 0x74, 0xd6,
	0x54, 0x06, 0xee, 0xe2, 0x66, 0x7b, 0x59, 0x2c, 0xe8, 0x00, 0xee, 0x33, 0x12, 0xc8, 0x09, 0x23,
	0x51, 0x07, 0x58, 0x00, 0xb1, 0xf0, 0xca, 0x1a, 0xe5, 0x50, 0x83, 0xe0, 0x7d, 0xe4, 0xd0, 0x38,
	0x88, 0xb2, 0x10, 0xfc, 0x14, 0x80, 0xf9, 0x01, 0x30, 0x41, 0xbb, 0x34, 0x20, 0x02, 0x1c, 0xa4,
	0xc8, 0xd8, 0x31, 0xfa, 0x0e, 0x00, 0x3b, 0x98, 0x6a, 0x9b, 0xb7, 0xbe, 0xff, 0xe5, 0x9b, 0x97,
	0xdf, 0x40, 0xf3, 0xe1, 0x35, 0xf8, 0x73, 0xdd, 0xb8, 0xe9, 0xe6, 0x23, 0xdb, 0x46, 0xc8, 0x36,
	0x53, 0xc6, 0x8f, 0x4b, 0x76, 0xa9, 0xba, 0xe2, 0xad, 0x65, 0x5c, 0xce, 0x44, 0xda, 0x27, 0xb5,
	0x9f, 0x2c, 0xb4, 0x31, 0xdf, 0x37, 0xdc, 0x40, 0x9b, 0x03, 0x32, 0x9a, 0x4e, 0xc3, 0x58, 0x00,
	0x57, 0x73, 0x5e, 0x6e, 0x5f, 0x7c, 0xda, 0x2e, 0xdd, 0x28, 0xec, 0x5e, 0xf0, 0x2a, 0x03, 0x32,
	0xca, 0x5b, 0x2b, 0xf5, 0xf8, 0x26, 0xda, 0x56, 0x43, 0xe6, 0xa7, 0x84, 0x09, 0x4a, 0x22, 0x7f,
	0x00, 0x9c, 0x93, 0x1e, 0x98, 0xa1, 0xfb, 0x9f, 0x52, 0x76, 0xb4, 0xee, 0xa1, 0x56, 0x35, 0xef,
	0xc8, 0x7a, 0x6e, 0xa3, 0xb7, 0xce, 0x57, 0xcf, 0x7c, 0x96, 0xb5, 0x1f, 0x8b, 0x68, 0x7d, 0xe6,
	0x06, 0xe0, 0xbb, 0x08, 0xc9, 0x2a, 0x81, 0xc9, 0xf5, 0x63, 0xae, 0xe5, 0xd5, 0xc5, 0x9d, 0x92,
	0x6e, 0xa7, 0x8c, 0x7a, 0x6b, 0xda, 0xe1, 0x94, 0x51, 0x7c, 0x0d, 0xad, 0xa7, 0x44, 0xf4, 0xfd,
	0x94, 0x41, 0x97, 0x8e, 0x54, 0xd2, 0x6b, 0x1e, 0x92, 0xa2, 0x8e, 0x92, 0xe0, 0x0c, 0x6d, 0xcb,
	0x4c, 0x12, 0x46, 0x9f, 0xa8, 0x05, 0x92, 0xd3, 0x63, 0xa6, 0xf2, 0xbd, 0x65, 0xee, 0x49, 0x6b,
	0x16, 0xc8, 0xb0, 0xe8, 0x6d, 0x91, 0x05, 0x52, 0x3c, 0x42, 0x3b, 0xcf, 0x86, 0xe5, 0x69, 0x12,
	0x73, 0x70, 0x6c, 0x15, 0xb7, 0xf5, 0x1f, 0xe2, 0x6a, 0x20, 0x6f, 0x9b, 0x2c, 0x12, 0x37, 0xf7,
	0x65, 0x73, 0x1a, 0xe8, 0xcd, 0xf3, 0x35, 0x67, 0xa6, 0x13, 0xc7, 0x25, 0xbb, 0x58, 0x2d, 0xe9,
	0x49, 0x3b, 0x2e, 0xd9, 0x2b, 0xd5, 0xd5, 0xe3, 0x92, 0xbd, 0x5a, 0xbd, 0x58, 0xfb, 0xc3, 0x42,
	0x5b, 0x8b, 0x4a, 0xc7, 0x1f, 0xa3, 0x8a, 0x1a, 0x0f, 0x08, 0xfd, 0x3e, 0x90, 0x10, 0x18, 0x37,
	0xfd, 0xdb, 0x9b, 0xbd, 0xeb, 0xe6, 0xfd, 0x90, 0x85, 0x7c, 0x40, 0xb9, 0x38, 0x51, 0x6f, 0xc8,
	0x43, 0x2d, 0xf4, 0x36, 0x0c, 0xc0, 0x91, 0xf6, 0xc7, 0x0f, 0xd0, 0x86, 0x81, 0xf2, 0x45, 0xe2,
	0x93, 0x30, 0x74, 0x0a, 0xbb, 0xc5, 0x7f, 0x5e, 0xd4, 0xda, 0xed, 0x53, 0x12, 0x65, 0xe0, 0x5d,
	0x32, 0x8e, 0x9f, 0x24, 0xad, 0x30, 0x6c, 0xb6, 0x24, 0x0d, 0x77, 0xcd, 0xfb, 0xfc, 0xaf, 0x34,
	0x2c, 0x2a, 0xaf, 0xf6, 0x43, 0x01, 0x6d, 0x2f, 0xa4, 0x1e, 0x9f, 0x21, 0x27, 0x2f, 0x3c, 0x4b,
	0xb9, 0x60, 0x40, 0x06, 0x2f, 0xcc, 0xc0, 0x8e, 0x41, 0x3a, 0x35, 0x40, 0x39, 0x13, 0x9f, 0xa3,
	0x5c, 0xe3, 0x07, 0x11, 0x85, 0x58, 0x4c, 0x22, 0x14, 0x96, 0x8c, 0xb0, 0x65, 0x70, 0x0e, 0x14,
	0x8c, 0xc1, 0x6f, 0xb6, 0x25, 0x41, 0xf7, 0xd0, 0x9d, 0x17, 0x22, 0x48, 0xf3, 0x50, 0xfb, 0xd3,
	0x42, 0xd5, 0x7c, 0x5d, 0x75, 0x40, 0x3f, 0x03, 0xf8, 0x55, 0x64, 0x87, 0x94, 0xeb, 0xc5, 0x2b,
	0xc9, 0xb0, 0xd5, 0xf6, 0xf9, 0xa2, 0x60, 0x5b, 0x47, 0x17, 0xbc, 0x89, 0x0a, 0x47, 0x68, 0x23,
	0xe8, 0x43, 0xf0, 0xc8, 0xe7, 0x66, 0x33, 0x98, 0xba, 0xde, 0x59, 0xe6, 0x66, 0x1c, 0x48, 0x84,
	0x7c, 0xb5, 0xb4, 0xed, 0xa7, 0xed, 0x95, 0x6f, 0xad, 0x42, 0x55, 0x06, 0x2a, 0x07, 0xb3, 0xaa,
	0xe6, 0x3d, 0x59, 0xed, 0x3e, 0xba, 0xbd, 0xdc, 0x0a, 0xce, 0x6b, 0x6a, 0x57, 0x90, 0x9d, 0x0c,
	0x81, 0x31, 0x1a, 0x02, 0x2e, 0xfe, 0xd5, 0xb6, 0x6a, 0xdf, 0x15, 0x50, 0x79, 0x2e, 0x38, 0xfe,
	0x0a, 0xe1, 0xfc, 0xb1, 0x9a, 0xa6, 0xee, 0x58, 0x6a, 0x7a, 0x3b, 0x2f, 0x5c, 0x93, 0x6b, 0x9e,
	0xb7, 0xc3, 0x89, 0xe7, 0x61, 0x2c, 0xd8, 0xd8, 0xdb, 0x0c, 0x9e, 0x95, 0x5f, 0x79, 0x1f, 0xed,
	0x2c, 0x36, 0xc6, 0x55, 0x54, 0x7c, 0x04, 0x63, 0xd5, 0x8c, 0x35, 0x4f, 0x7e, 0xe2, 0x2d, 0xb4,
	0x32, 0x94, 0x97, 0xc6, 0x2c, 0x4c, 0x7d, 0x68, 0x16, 0xf6, 0xad, 0x66, 0x53, 0x12, 0x75, 0x0b,
	0x35, 0xce, 0x47, 0xd4, 0x3c, 0xff, 0x27, 0x3f, 0x7f, 0xfd, 0xeb, 0x6f, 0xab, 0x85, 0x6a, 0x01,
	0xed, 0xd3, 0x44, 0x97, 0x9c, 0xb2, 0x64, 0x34, 0x5e, 0xa2, 0xfa, 0x76, 0x79, 0xc2, 0xbd, 0xfc,
	0x87, 0xeb, 0x58, 0x67, 0xab, 0xea, 0x67, 0xae, 0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f,
	0x33, 0xe9, 0x7c, 0x6a, 0x0b, 0x00, 0x00,
}
