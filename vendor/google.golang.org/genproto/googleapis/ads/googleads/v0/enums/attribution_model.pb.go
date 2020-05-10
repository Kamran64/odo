// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/attribution_model.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The attribution model that describes how to distribute credit for a
// particular conversion across potentially many prior interactions.
type AttributionModelEnum_AttributionModel int32

const (
	// Not specified.
	AttributionModelEnum_UNSPECIFIED AttributionModelEnum_AttributionModel = 0
	// Used for return value only. Represents value unknown in this version.
	AttributionModelEnum_UNKNOWN AttributionModelEnum_AttributionModel = 1
	// Uses external attribution.
	AttributionModelEnum_EXTERNAL AttributionModelEnum_AttributionModel = 100
	// Attributes all credit for a conversion to its last click.
	AttributionModelEnum_GOOGLE_ADS_LAST_CLICK AttributionModelEnum_AttributionModel = 101
	// Attributes all credit for a conversion to its first click using Google
	// Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK AttributionModelEnum_AttributionModel = 102
	// Attributes credit for a conversion equally across all of its clicks using
	// Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_LINEAR AttributionModelEnum_AttributionModel = 103
	// Attributes exponentially more credit for a conversion to its more recent
	// clicks using Google Search attribution (half-life is 1 week).
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY AttributionModelEnum_AttributionModel = 104
	// Attributes 40% of the credit for a conversion to its first and last
	// clicks. Remaining 20% is evenly distributed across all other clicks. This
	// uses Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED AttributionModelEnum_AttributionModel = 105
	// Flexible model that uses machine learning to determine the appropriate
	// distribution of credit among clicks using Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN AttributionModelEnum_AttributionModel = 106
)

var AttributionModelEnum_AttributionModel_name = map[int32]string{
	0:   "UNSPECIFIED",
	1:   "UNKNOWN",
	100: "EXTERNAL",
	101: "GOOGLE_ADS_LAST_CLICK",
	102: "GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK",
	103: "GOOGLE_SEARCH_ATTRIBUTION_LINEAR",
	104: "GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY",
	105: "GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED",
	106: "GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN",
}
var AttributionModelEnum_AttributionModel_value = map[string]int32{
	"UNSPECIFIED":                              0,
	"UNKNOWN":                                  1,
	"EXTERNAL":                                 100,
	"GOOGLE_ADS_LAST_CLICK":                    101,
	"GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK":    102,
	"GOOGLE_SEARCH_ATTRIBUTION_LINEAR":         103,
	"GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY":     104,
	"GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED": 105,
	"GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN":    106,
}

func (x AttributionModelEnum_AttributionModel) String() string {
	return proto.EnumName(AttributionModelEnum_AttributionModel_name, int32(x))
}
func (AttributionModelEnum_AttributionModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_attribution_model_710acdb5f1419493, []int{0, 0}
}

// Container for enum representing the attribution model that describes how to
// distribute credit for a particular conversion across potentially many prior
// interactions.
type AttributionModelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributionModelEnum) Reset()         { *m = AttributionModelEnum{} }
func (m *AttributionModelEnum) String() string { return proto.CompactTextString(m) }
func (*AttributionModelEnum) ProtoMessage()    {}
func (*AttributionModelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_attribution_model_710acdb5f1419493, []int{0}
}
func (m *AttributionModelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributionModelEnum.Unmarshal(m, b)
}
func (m *AttributionModelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributionModelEnum.Marshal(b, m, deterministic)
}
func (dst *AttributionModelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributionModelEnum.Merge(dst, src)
}
func (m *AttributionModelEnum) XXX_Size() int {
	return xxx_messageInfo_AttributionModelEnum.Size(m)
}
func (m *AttributionModelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributionModelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AttributionModelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AttributionModelEnum)(nil), "google.ads.googleads.v0.enums.AttributionModelEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AttributionModelEnum_AttributionModel", AttributionModelEnum_AttributionModel_name, AttributionModelEnum_AttributionModel_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/attribution_model.proto", fileDescriptor_attribution_model_710acdb5f1419493)
}

var fileDescriptor_attribution_model_710acdb5f1419493 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x6a, 0xd5, 0x40,
	0x14, 0xc6, 0xbd, 0x11, 0x54, 0xa6, 0x82, 0xc3, 0x60, 0x17, 0x2e, 0x0a, 0xb6, 0x54, 0xa8, 0x20,
	0x93, 0x80, 0xb8, 0x19, 0x57, 0x27, 0xc9, 0xf4, 0x3a, 0x34, 0x9d, 0x5c, 0x92, 0xdc, 0xf8, 0x87,
	0xc0, 0x90, 0x9a, 0x18, 0x23, 0x37, 0x99, 0x72, 0x27, 0xb7, 0x0f, 0xe4, 0x52, 0xf0, 0x35, 0x5c,
	0xf8, 0x28, 0x7d, 0x0a, 0x49, 0xd2, 0x46, 0x28, 0xa4, 0x9b, 0xe1, 0x63, 0xbe, 0xdf, 0x77, 0x38,
	0x9c, 0x0f, 0xbd, 0xab, 0xb4, 0xae, 0x36, 0xa5, 0x9d, 0x17, 0xc6, 0x1e, 0x65, 0xaf, 0xae, 0x1c,
	0xbb, 0x6c, 0x77, 0x8d, 0xb1, 0xf3, 0xae, 0xdb, 0xd6, 0x17, 0xbb, 0xae, 0xd6, 0xad, 0x6a, 0x74,
	0x51, 0x6e, 0xe8, 0xe5, 0x56, 0x77, 0x9a, 0x1c, 0x8c, 0x2c, 0xcd, 0x0b, 0x43, 0xa7, 0x18, 0xbd,
	0x72, 0xe8, 0x10, 0x3b, 0xfa, 0x63, 0xa1, 0xe7, 0xf0, 0x3f, 0x7a, 0xde, 0x27, 0x79, 0xbb, 0x6b,
	0x8e, 0x7e, 0x5b, 0x08, 0xdf, 0x35, 0xc8, 0x33, 0xb4, 0xb7, 0x96, 0xf1, 0x8a, 0x7b, 0xe2, 0x54,
	0x70, 0x1f, 0x3f, 0x20, 0x7b, 0xe8, 0xf1, 0x5a, 0x9e, 0xc9, 0xf0, 0xa3, 0xc4, 0x0b, 0xf2, 0x14,
	0x3d, 0xe1, 0x9f, 0x12, 0x1e, 0x49, 0x08, 0x70, 0x41, 0x5e, 0xa0, 0xfd, 0x65, 0x18, 0x2e, 0x03,
	0xae, 0xc0, 0x8f, 0x55, 0x00, 0x71, 0xa2, 0xbc, 0x40, 0x78, 0x67, 0xb8, 0x24, 0xaf, 0xd1, 0xab,
	0x1b, 0x2b, 0xe6, 0x10, 0x79, 0x1f, 0x14, 0x24, 0x49, 0x24, 0xdc, 0x75, 0x22, 0x42, 0xa9, 0x4e,
	0x45, 0x34, 0xa1, 0xdf, 0xc8, 0x31, 0x7a, 0x39, 0x8f, 0x06, 0x42, 0x72, 0x88, 0x70, 0x45, 0x4e,
	0xd0, 0xf1, 0x3c, 0x95, 0x88, 0x73, 0xae, 0x7c, 0xee, 0xc1, 0x67, 0xfc, 0x9d, 0xbc, 0x41, 0x27,
	0xf3, 0xe4, 0x2a, 0x8c, 0xc5, 0x20, 0x5c, 0x88, 0xb9, 0x8f, 0xeb, 0xfb, 0x17, 0xf5, 0x21, 0x01,
	0xe5, 0x47, 0x22, 0xe5, 0x12, 0xff, 0x70, 0xaf, 0x17, 0xe8, 0xf0, 0xab, 0x6e, 0xe8, 0xbd, 0xe7,
	0x76, 0xf7, 0xef, 0x9e, 0x74, 0xd5, 0x97, 0xb4, 0x5a, 0x7c, 0x71, 0x6f, 0x72, 0x95, 0xde, 0xe4,
	0x6d, 0x45, 0xf5, 0xb6, 0xb2, 0xab, 0xb2, 0x1d, 0x2a, 0xbc, 0x6d, 0xfb, 0xb2, 0x36, 0x33, 0xe5,
	0xbf, 0x1f, 0xde, 0x9f, 0xd6, 0xc3, 0x25, 0xc0, 0x2f, 0xeb, 0x60, 0x39, 0x8e, 0x82, 0xc2, 0xd0,
	0x51, 0xf6, 0x2a, 0x75, 0x68, 0xdf, 0xab, 0xf9, 0x7b, 0xeb, 0x67, 0x50, 0x98, 0x6c, 0xf2, 0xb3,
	0xd4, 0xc9, 0x06, 0xff, 0xda, 0x3a, 0x1c, 0x3f, 0x19, 0x83, 0xc2, 0x30, 0x36, 0x11, 0x8c, 0xa5,
	0x0e, 0x63, 0x03, 0x73, 0xf1, 0x68, 0x58, 0xec, 0xed, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33,
	0x7f, 0x44, 0x3d, 0x94, 0x02, 0x00, 0x00,
}
