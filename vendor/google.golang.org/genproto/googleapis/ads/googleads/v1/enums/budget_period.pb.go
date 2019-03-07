// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/budget_period.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible period of a Budget.
type BudgetPeriodEnum_BudgetPeriod int32

const (
	// Not specified.
	BudgetPeriodEnum_UNSPECIFIED BudgetPeriodEnum_BudgetPeriod = 0
	// Used for return value only. Represents value unknown in this version.
	BudgetPeriodEnum_UNKNOWN BudgetPeriodEnum_BudgetPeriod = 1
	// Daily budget.
	BudgetPeriodEnum_DAILY BudgetPeriodEnum_BudgetPeriod = 2
	// Custom budget.
	BudgetPeriodEnum_CUSTOM BudgetPeriodEnum_BudgetPeriod = 3
	// Fixed daily budget.
	BudgetPeriodEnum_FIXED_DAILY BudgetPeriodEnum_BudgetPeriod = 4
)

var BudgetPeriodEnum_BudgetPeriod_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DAILY",
	3: "CUSTOM",
	4: "FIXED_DAILY",
}
var BudgetPeriodEnum_BudgetPeriod_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DAILY":       2,
	"CUSTOM":      3,
	"FIXED_DAILY": 4,
}

func (x BudgetPeriodEnum_BudgetPeriod) String() string {
	return proto.EnumName(BudgetPeriodEnum_BudgetPeriod_name, int32(x))
}
func (BudgetPeriodEnum_BudgetPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_budget_period_0817578d922cd855, []int{0, 0}
}

// Message describing Budget period.
type BudgetPeriodEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BudgetPeriodEnum) Reset()         { *m = BudgetPeriodEnum{} }
func (m *BudgetPeriodEnum) String() string { return proto.CompactTextString(m) }
func (*BudgetPeriodEnum) ProtoMessage()    {}
func (*BudgetPeriodEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_budget_period_0817578d922cd855, []int{0}
}
func (m *BudgetPeriodEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BudgetPeriodEnum.Unmarshal(m, b)
}
func (m *BudgetPeriodEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BudgetPeriodEnum.Marshal(b, m, deterministic)
}
func (dst *BudgetPeriodEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetPeriodEnum.Merge(dst, src)
}
func (m *BudgetPeriodEnum) XXX_Size() int {
	return xxx_messageInfo_BudgetPeriodEnum.Size(m)
}
func (m *BudgetPeriodEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetPeriodEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetPeriodEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BudgetPeriodEnum)(nil), "google.ads.googleads.v1.enums.BudgetPeriodEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.BudgetPeriodEnum_BudgetPeriod", BudgetPeriodEnum_BudgetPeriod_name, BudgetPeriodEnum_BudgetPeriod_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/budget_period.proto", fileDescriptor_budget_period_0817578d922cd855)
}

var fileDescriptor_budget_period_0817578d922cd855 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0x76, 0x9d, 0x4e, 0xcc, 0x04, 0x6b, 0x8f, 0xe2, 0x0e, 0xdb, 0x0f, 0x48, 0x08, 0xde, 0xe2,
	0x29, 0xdd, 0xba, 0x51, 0xd4, 0xae, 0xb0, 0x0f, 0x3f, 0x28, 0x8c, 0xce, 0x94, 0x58, 0x58, 0x93,
	0xd2, 0xb4, 0xfb, 0x41, 0x1e, 0xfd, 0x29, 0xfe, 0x10, 0x0f, 0xfe, 0x0a, 0x69, 0x62, 0xcb, 0x2e,
	0x7a, 0x09, 0x0f, 0xef, 0xf3, 0x91, 0xe7, 0x7d, 0x01, 0xe6, 0x52, 0xf2, 0x5d, 0x82, 0x62, 0xa6,
	0x90, 0x81, 0x35, 0xda, 0x63, 0x94, 0x88, 0x2a, 0x53, 0x68, 0x5b, 0x31, 0x9e, 0x94, 0x9b, 0x3c,
	0x29, 0x52, 0xc9, 0x60, 0x5e, 0xc8, 0x52, 0x3a, 0x03, 0xa3, 0x83, 0x31, 0x53, 0xb0, 0xb5, 0xc0,
	0x3d, 0x86, 0xda, 0x72, 0x75, 0xdd, 0x24, 0xe6, 0x29, 0x8a, 0x85, 0x90, 0x65, 0x5c, 0xa6, 0x52,
	0x28, 0x63, 0x1e, 0xbd, 0x01, 0xdb, 0xd5, 0x99, 0xa1, 0x8e, 0xf4, 0x44, 0x95, 0x8d, 0x96, 0xe0,
	0xfc, 0x70, 0xe6, 0x5c, 0x80, 0xfe, 0x2a, 0x58, 0x84, 0xde, 0xd8, 0x9f, 0xfa, 0xde, 0xc4, 0x3e,
	0x72, 0xfa, 0xe0, 0x74, 0x15, 0xdc, 0x05, 0xf3, 0xc7, 0xc0, 0xee, 0x38, 0x67, 0xe0, 0x64, 0x42,
	0xfd, 0xfb, 0x67, 0xdb, 0x72, 0x00, 0xe8, 0x8d, 0x57, 0x8b, 0xe5, 0xfc, 0xc1, 0xee, 0xd6, 0xa6,
	0xa9, 0xff, 0xe4, 0x4d, 0x36, 0x86, 0x3c, 0x76, 0xbf, 0x3a, 0x60, 0xf8, 0x2a, 0x33, 0xf8, 0x6f,
	0x5b, 0xf7, 0xf2, 0xf0, 0xe7, 0xb0, 0xae, 0x18, 0x76, 0x5e, 0xdc, 0x5f, 0x0f, 0x97, 0xbb, 0x58,
	0x70, 0x28, 0x0b, 0x8e, 0x78, 0x22, 0xf4, 0x02, 0xcd, 0x91, 0xf2, 0x54, 0xfd, 0x71, 0xb3, 0x5b,
	0xfd, 0xbe, 0x5b, 0xdd, 0x19, 0xa5, 0x1f, 0xd6, 0x60, 0x66, 0xa2, 0x28, 0x53, 0xd0, 0xc0, 0x1a,
	0xad, 0x31, 0xac, 0x37, 0x57, 0x9f, 0x0d, 0x1f, 0x51, 0xa6, 0xa2, 0x96, 0x8f, 0xd6, 0x38, 0xd2,
	0xfc, 0xb7, 0x35, 0x34, 0x43, 0x42, 0x28, 0x53, 0x84, 0xb4, 0x0a, 0x42, 0xd6, 0x98, 0x10, 0xad,
	0xd9, 0xf6, 0x74, 0xb1, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xa5, 0x03, 0x16, 0xcb,
	0x01, 0x00, 0x00,
}
