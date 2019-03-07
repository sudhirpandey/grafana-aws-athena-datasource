// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/account_budget_proposal_status.proto

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

// The possible statuses of an AccountBudgetProposal.
type AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus int32

const (
	// Not specified.
	AccountBudgetProposalStatusEnum_UNSPECIFIED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AccountBudgetProposalStatusEnum_UNKNOWN AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 1
	// The proposal is pending approval.
	AccountBudgetProposalStatusEnum_PENDING AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 2
	// The proposal has been approved but the corresponding billing setup
	// has not.  This can occur for proposals that set up the first budget
	// when signing up for billing or when performing a change of bill-to
	// operation.
	AccountBudgetProposalStatusEnum_APPROVED_HELD AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 3
	// The proposal has been approved.
	AccountBudgetProposalStatusEnum_APPROVED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 4
	// The proposal has been cancelled by the user.
	AccountBudgetProposalStatusEnum_CANCELLED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 5
	// The proposal has been rejected by the user, e.g. by rejecting an
	// acceptance email.
	AccountBudgetProposalStatusEnum_REJECTED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 6
)

var AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED_HELD",
	4: "APPROVED",
	5: "CANCELLED",
	6: "REJECTED",
}
var AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"PENDING":       2,
	"APPROVED_HELD": 3,
	"APPROVED":      4,
	"CANCELLED":     5,
	"REJECTED":      6,
}

func (x AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus) String() string {
	return proto.EnumName(AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name, int32(x))
}
func (AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_status_517bcbbece97130f, []int{0, 0}
}

// Message describing AccountBudgetProposal statuses.
type AccountBudgetProposalStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetProposalStatusEnum) Reset()         { *m = AccountBudgetProposalStatusEnum{} }
func (m *AccountBudgetProposalStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalStatusEnum) ProtoMessage()    {}
func (*AccountBudgetProposalStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_status_517bcbbece97130f, []int{0}
}
func (m *AccountBudgetProposalStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Unmarshal(m, b)
}
func (m *AccountBudgetProposalStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Marshal(b, m, deterministic)
}
func (dst *AccountBudgetProposalStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalStatusEnum.Merge(dst, src)
}
func (m *AccountBudgetProposalStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Size(m)
}
func (m *AccountBudgetProposalStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccountBudgetProposalStatusEnum)(nil), "google.ads.googleads.v1.enums.AccountBudgetProposalStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus", AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name, AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/account_budget_proposal_status.proto", fileDescriptor_account_budget_proposal_status_517bcbbece97130f)
}

var fileDescriptor_account_budget_proposal_status_517bcbbece97130f = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xb5, 0x45, 0x51, 0x17, 0x89, 0xb5, 0x47, 0x95, 0x28, 0x7c, 0xc0, 0x36, 0x8d, 0xb7, 0xf5,
	0xb4, 0x6d, 0x57, 0x44, 0x49, 0x69, 0x40, 0x6a, 0x62, 0x9a, 0x34, 0x0b, 0x6d, 0x1a, 0x12, 0xd8,
	0x6d, 0xd8, 0x2d, 0x9f, 0xe0, 0x87, 0x78, 0xe4, 0x53, 0xfc, 0x14, 0xef, 0xde, 0x4d, 0xb7, 0x94,
	0x9b, 0x5c, 0x36, 0x6f, 0xe7, 0xcd, 0xcc, 0x9b, 0x79, 0x03, 0x9c, 0x8c, 0xf3, 0x6c, 0x99, 0x5a,
	0x34, 0x11, 0x56, 0x05, 0x4b, 0xb4, 0xb1, 0xad, 0x94, 0x15, 0x2b, 0x61, 0xd1, 0xf9, 0x9c, 0x17,
	0x4c, 0xc6, 0xb3, 0x22, 0xc9, 0x52, 0x19, 0xe7, 0x6b, 0x9e, 0x73, 0x41, 0x97, 0xb1, 0x90, 0x54,
	0x16, 0x02, 0xe6, 0x6b, 0x2e, 0xb9, 0xd9, 0xa9, 0x0a, 0x21, 0x4d, 0x04, 0xdc, 0xf7, 0x80, 0x1b,
	0x1b, 0xaa, 0x1e, 0xd7, 0xb7, 0xb5, 0x44, 0xbe, 0xb0, 0x28, 0x63, 0x5c, 0x52, 0xb9, 0xe0, 0x6c,
	0x57, 0xdc, 0xdb, 0x6a, 0xe0, 0x0e, 0x57, 0x2a, 0x8e, 0x12, 0x09, 0x76, 0x1a, 0x13, 0x25, 0x41,
	0x58, 0xb1, 0xea, 0x7d, 0x6a, 0xe0, 0xe6, 0x40, 0x8e, 0x79, 0x09, 0x5a, 0x53, 0x7f, 0x12, 0x10,
	0x77, 0xf0, 0x34, 0x20, 0x9e, 0x71, 0x64, 0xb6, 0xc0, 0xe9, 0xd4, 0x7f, 0xf5, 0x47, 0xef, 0xbe,
	0xa1, 0x95, 0x9f, 0x80, 0xf8, 0xde, 0xc0, 0xef, 0x1b, 0xba, 0x79, 0x05, 0xda, 0x38, 0x08, 0xc6,
	0xa3, 0x90, 0x78, 0xf1, 0x33, 0x19, 0x7a, 0x46, 0xc3, 0xbc, 0x00, 0x67, 0x75, 0xc8, 0x38, 0x36,
	0xdb, 0xe0, 0xdc, 0xc5, 0xbe, 0x4b, 0x86, 0x43, 0xe2, 0x19, 0x27, 0x25, 0x39, 0x26, 0x2f, 0xc4,
	0x7d, 0x23, 0x9e, 0xd1, 0x74, 0x7e, 0x35, 0xd0, 0x9d, 0xf3, 0x15, 0x3c, 0xb8, 0xb0, 0x73, 0x7f,
	0x60, 0xd6, 0xa0, 0x5c, 0x3a, 0xd0, 0x3e, 0x76, 0xbe, 0xc3, 0x8c, 0x2f, 0x29, 0xcb, 0x20, 0x5f,
	0x67, 0x56, 0x96, 0x32, 0x65, 0x49, 0x7d, 0x87, 0x7c, 0x21, 0xfe, 0x39, 0xcb, 0xa3, 0x7a, 0xbf,
	0xf4, 0x46, 0x1f, 0xe3, 0xad, 0xde, 0xe9, 0x57, 0xad, 0x70, 0x22, 0x60, 0x05, 0x4b, 0x14, 0xda,
	0xb0, 0xf4, 0x4e, 0x7c, 0xd7, 0x7c, 0x84, 0x13, 0x11, 0xed, 0xf9, 0x28, 0xb4, 0x23, 0xc5, 0xff,
	0xe8, 0xdd, 0x2a, 0x88, 0x10, 0x4e, 0x04, 0x42, 0xfb, 0x0c, 0x84, 0x42, 0x1b, 0x21, 0x95, 0x33,
	0x6b, 0xaa, 0xc1, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x9f, 0xc3, 0xa1, 0x2e, 0x02,
	0x00, 0x00,
}
