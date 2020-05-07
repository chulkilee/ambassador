// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package envoy_config_filter_http_transcoder_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof" json:"proto_descriptor,omitempty"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof" json:"proto_descriptor_bin,omitempty"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_540f2f6de8b0585c)
}

var fileDescriptor_540f2f6de8b0585c = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0x13, 0x3d,
	0x14, 0xc5, 0xeb, 0xf6, 0x6b, 0xbe, 0xa9, 0x9b, 0xfe, 0xc1, 0x40, 0x3b, 0x44, 0x6d, 0x08, 0x08,
	0x50, 0x24, 0xa4, 0x19, 0x29, 0x05, 0x15, 0xc1, 0x02, 0x35, 0xa2, 0xb4, 0x45, 0x02, 0xc2, 0x00,
	0x62, 0x69, 0xb9, 0x33, 0x6e, 0x6a, 0xc8, 0xd8, 0xae, 0xed, 0x49, 0x9b, 0x1d, 0x6f, 0x80, 0xc4,
	0x0a, 0xf1, 0x06, 0xbc, 0x02, 0x4f, 0xd0, 0x25, 0xbc, 0x01, 0xea, 0x1b, 0xc0, 0x0a, 0xb1, 0x40,
	0xc8, 0x76, 0x1a, 0x86, 0x76, 0xd3, 0xdd, 0xd8, 0xe7, 0x9c, 0x7b, 0xe7, 0x5e, 0xff, 0xe0, 0x2a,
	0xe5, 0x7d, 0x31, 0x88, 0x53, 0xc1, 0x77, 0x58, 0x37, 0xde, 0x61, 0x3d, 0x43, 0x55, 0xbc, 0x6b,
	0x8c, 0x8c, 0x8d, 0x22, 0x5c, 0xa7, 0x22, 0xa3, 0x2a, 0xee, 0xb7, 0x4a, 0xa7, 0x48, 0x2a, 0x61,
	0x04, 0xba, 0xe1, 0x82, 0x91, 0x0f, 0x46, 0x3e, 0x18, 0xd9, 0x60, 0x54, 0xb2, 0xf6, 0x5b, 0xb5,
	0x7a, 0x91, 0x49, 0x12, 0x13, 0xce, 0x85, 0x21, 0x86, 0x09, 0xae, 0xe3, 0x9c, 0x75, 0x15, 0x31,
	0xd4, 0xd7, 0xa9, 0x2d, 0x9f, 0xd2, 0xb5, 0x21, 0xa6, 0xd0, 0x43, 0x79, 0xb1, 0x4f, 0x7a, 0x2c,
	0x23, 0x86, 0xc6, 0xc7, 0x1f, 0x5e, 0xb8, 0xfa, 0xa9, 0x02, 0xd1, 0x86, 0x92, 0xe9, 0x23, 0x2d,
	0xf8, 0x8b, 0x51, 0x47, 0x74, 0x13, 0xce, 0x3b, 0x1d, 0x67, 0x54, 0xa7, 0x8a, 0x49, 0x23, 0x54,
	0x08, 0x1a, 0xa0, 0x39, 0xb5, 0x39, 0x96, 0xcc, 0x39, 0xe5, 0xc1, 0x48, 0x40, 0x2d, 0x78, 0xe1,
	0xa4, 0x19, 0x6f, 0x33, 0x1e, 0xfe, 0xd7, 0x00, 0xcd, 0xea, 0xe6, 0x58, 0x82, 0x4e, 0x04, 0xda,
	0x8c, 0xa3, 0x6b, 0x30, 0xd0, 0x54, 0xf5, 0x59, 0x4a, 0x75, 0x38, 0xde, 0x98, 0x68, 0x4e, 0xb5,
	0x83, 0x5f, 0xed, 0xc9, 0xf7, 0x60, 0x3c, 0x00, 0xc9, 0x48, 0x41, 0x3d, 0x38, 0x23, 0x15, 0xe3,
	0x06, 0x0b, 0xe9, 0x86, 0x0a, 0x27, 0x1a, 0xa0, 0x39, 0xdd, 0xda, 0x88, 0xce, 0xb6, 0xb5, 0xe8,
	0xf4, 0x64, 0x51, 0xc7, 0xd6, 0x7b, 0xea, 0xcb, 0x25, 0x55, 0x59, 0x3a, 0xa1, 0xfb, 0x70, 0x29,
	0x27, 0x26, 0xdd, 0xc5, 0x8c, 0xa7, 0x22, 0x67, 0xbc, 0x8b, 0x15, 0xdd, 0x2b, 0xa8, 0x36, 0x58,
	0x89, 0xc2, 0xd0, 0x70, 0xb2, 0x01, 0x9a, 0x41, 0x72, 0xc9, 0x79, 0xb6, 0x86, 0x96, 0xc4, 0x3b,
	0x12, 0x6b, 0x40, 0x77, 0x60, 0xc8, 0xba, 0x5c, 0x28, 0x9a, 0xe1, 0xbd, 0x82, 0xaa, 0x01, 0x96,
	0x44, 0x91, 0x9c, 0x1a, 0xaa, 0x74, 0x58, 0xb1, 0x43, 0x26, 0x0b, 0x43, 0xfd, 0x99, 0x95, 0x3b,
	0x23, 0x15, 0x5d, 0x81, 0x55, 0x52, 0x18, 0x81, 0x73, 0x22, 0x25, 0xe3, 0xdd, 0xf0, 0x7f, 0xd7,
	0x6a, 0xda, 0xde, 0x3d, 0xf6, 0x57, 0x68, 0x1d, 0x5e, 0xf6, 0x61, 0x5c, 0xf0, 0x37, 0x5c, 0xec,
	0xf3, 0xd3, 0x3d, 0x02, 0x97, 0x5a, 0xf2, 0xb6, 0x97, 0xde, 0x75, 0xb2, 0x53, 0x04, 0xcf, 0xa7,
	0x82, 0xf7, 0xa9, 0x32, 0xb8, 0xab, 0x64, 0x8a, 0x3d, 0x26, 0xe1, 0x94, 0x8b, 0x9e, 0x1b, 0x4a,
	0x76, 0x6f, 0xcf, 0x9d, 0x50, 0xfb, 0x0e, 0x60, 0xb5, 0xbc, 0x33, 0x74, 0x1d, 0xce, 0x92, 0x2c,
	0xc3, 0xfb, 0xbb, 0xcc, 0x50, 0x2d, 0x49, 0x4a, 0x1d, 0x18, 0x41, 0x32, 0x43, 0xb2, 0xec, 0xd5,
	0xe8, 0x12, 0xad, 0xc1, 0x65, 0xd2, 0xdb, 0x27, 0x03, 0x8d, 0xfd, 0x0b, 0x4a, 0xc5, 0x72, 0x66,
	0x58, 0x9f, 0xe2, 0x1d, 0x46, 0x7b, 0x99, 0x7d, 0x75, 0x9b, 0xaa, 0x79, 0x93, 0xeb, 0xd0, 0x39,
	0xb6, 0x3c, 0x74, 0x0e, 0x74, 0x17, 0xd6, 0xfe, 0x29, 0x41, 0x79, 0x91, 0x6b, 0x4c, 0x34, 0x66,
	0xdc, 0x78, 0x14, 0x82, 0x64, 0xa1, 0x94, 0x5f, 0xb7, 0xfa, 0x9a, 0xde, 0xe2, 0x46, 0xa3, 0x7b,
	0xb0, 0x26, 0x15, 0xb5, 0x20, 0x51, 0xec, 0xe1, 0x74, 0x6d, 0x31, 0x27, 0x39, 0xd5, 0x8e, 0xcc,
	0x20, 0x59, 0x3c, 0x76, 0x74, 0xac, 0xc1, 0x35, 0x7d, 0x62, 0xe5, 0xf6, 0x45, 0x38, 0x5b, 0x42,
	0x59, 0x53, 0x83, 0x26, 0x7e, 0xb6, 0x41, 0xfb, 0x23, 0x38, 0x3c, 0xaa, 0x83, 0x2f, 0x47, 0x75,
	0xf0, 0xed, 0xa8, 0x0e, 0x7e, 0x7c, 0xf8, 0xfd, 0x6e, 0x72, 0x15, 0xdd, 0xf6, 0x28, 0xd2, 0x03,
	0x43, 0xb9, 0xb6, 0x0b, 0x1a, 0xe2, 0xa8, 0x3d, 0x8f, 0x6e, 0xc7, 0xaf, 0xb5, 0xe0, 0xb8, 0x4c,
	0xe6, 0xca, 0xe7, 0xb7, 0x87, 0x5f, 0x2b, 0xe3, 0xf3, 0x63, 0xf0, 0x16, 0x13, 0x1e, 0x66, 0xa9,
	0xc4, 0xc1, 0xe0, 0x8c, 0x5c, 0xb7, 0xe7, 0xfe, 0x02, 0xed, 0xfe, 0xbd, 0x03, 0xb6, 0x2b, 0x6e,
	0xca, 0x95, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x65, 0xf6, 0x01, 0x8a, 0x04, 0x00, 0x00,
}

func (m *GrpcJsonTranscoder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConvertGrpcStatus {
		i--
		if m.ConvertGrpcStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.IgnoreUnknownQueryParameters {
		i--
		if m.IgnoreUnknownQueryParameters {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.AutoMapping {
		i--
		if m.AutoMapping {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for iNdEx := len(m.IgnoredQueryParameters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IgnoredQueryParameters[iNdEx])
			copy(dAtA[i:], m.IgnoredQueryParameters[iNdEx])
			i = encodeVarintTranscoder(dAtA, i, uint64(len(m.IgnoredQueryParameters[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MatchIncomingRequestRoute {
		i--
		if m.MatchIncomingRequestRoute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.DescriptorSet != nil {
		{
			size := m.DescriptorSet.Size()
			i -= size
			if _, err := m.DescriptorSet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.PrintOptions != nil {
		{
			size, err := m.PrintOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTranscoder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Services[iNdEx])
			copy(dAtA[i:], m.Services[iNdEx])
			i = encodeVarintTranscoder(dAtA, i, uint64(len(m.Services[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtoDescriptor)
	copy(dAtA[i:], m.ProtoDescriptor)
	i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptor)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProtoDescriptorBin != nil {
		i -= len(m.ProtoDescriptorBin)
		copy(dAtA[i:], m.ProtoDescriptorBin)
		i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptorBin)))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PreserveProtoFieldNames {
		i--
		if m.PreserveProtoFieldNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AlwaysPrintEnumsAsInts {
		i--
		if m.AlwaysPrintEnumsAsInts {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AlwaysPrintPrimitiveFields {
		i--
		if m.AlwaysPrintPrimitiveFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.AddWhitespace {
		i--
		if m.AddWhitespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTranscoder(dAtA []byte, offset int, v uint64) int {
	offset -= sovTranscoder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GrpcJsonTranscoder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DescriptorSet != nil {
		n += m.DescriptorSet.Size()
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.PrintOptions != nil {
		l = m.PrintOptions.Size()
		n += 1 + l + sovTranscoder(uint64(l))
	}
	if m.MatchIncomingRequestRoute {
		n += 2
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for _, s := range m.IgnoredQueryParameters {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.AutoMapping {
		n += 2
	}
	if m.IgnoreUnknownQueryParameters {
		n += 2
	}
	if m.ConvertGrpcStatus {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptor)
	n += 1 + l + sovTranscoder(uint64(l))
	return n
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProtoDescriptorBin != nil {
		l = len(m.ProtoDescriptorBin)
		n += 1 + l + sovTranscoder(uint64(l))
	}
	return n
}
func (m *GrpcJsonTranscoder_PrintOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddWhitespace {
		n += 2
	}
	if m.AlwaysPrintPrimitiveFields {
		n += 2
	}
	if m.AlwaysPrintEnumsAsInts {
		n += 2
	}
	if m.PreserveProtoFieldNames {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTranscoder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTranscoder(x uint64) (n int) {
	return sovTranscoder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcJsonTranscoder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: GrpcJsonTranscoder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcJsonTranscoder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrintOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrintOptions == nil {
				m.PrintOptions = &GrpcJsonTranscoder_PrintOptions{}
			}
			if err := m.PrintOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptorBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{v}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchIncomingRequestRoute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.MatchIncomingRequestRoute = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredQueryParameters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredQueryParameters = append(m.IgnoredQueryParameters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoMapping", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.AutoMapping = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreUnknownQueryParameters", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.IgnoreUnknownQueryParameters = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConvertGrpcStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.ConvertGrpcStatus = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: PrintOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrintOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddWhitespace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.AddWhitespace = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintPrimitiveFields", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.AlwaysPrintPrimitiveFields = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintEnumsAsInts", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.AlwaysPrintEnumsAsInts = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreserveProtoFieldNames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
			m.PreserveProtoFieldNames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTranscoder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
				return 0, ErrInvalidLengthTranscoder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTranscoder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTranscoder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTranscoder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTranscoder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTranscoder = fmt.Errorf("proto: unexpected end of group")
)
