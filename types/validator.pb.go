// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/types/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_maticnetwork_heimdall_types_common "github.com/maticnetwork/heimdall/types/common"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorID int32

const (
	DEFAULT ValidatorID = 0
)

var ValidatorID_name = map[int32]string{
	0: "DEFAULT",
}

var ValidatorID_value = map[string]int32{
	"DEFAULT": 0,
}

func (ValidatorID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{0}
}

type ValidatorSet struct {
	Validators       []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer         *Validator   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower int64        `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (m *ValidatorSet) Reset()      { *m = ValidatorSet{} }
func (*ValidatorSet) ProtoMessage() {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{0}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSet.Unmarshal(m, b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return xxx_messageInfo_ValidatorSet.Size(m)
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

type Validator struct {
	ID               ValidatorID                                                   `protobuf:"varint,1,opt,name=ID,proto3,enum=heimdall.types.ValidatorID" json:"ID,omitempty"`
	StartEpoch       uint64                                                        `protobuf:"varint,2,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch         uint64                                                        `protobuf:"varint,3,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Nonce            uint64                                                        `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	VotingPower      int64                                                         `protobuf:"varint,5,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	PubKey           github_com_maticnetwork_heimdall_types_common.PubKey          `protobuf:"bytes,6,opt,name=pub_key,json=pubKey,proto3,casttype=github.com/maticnetwork/heimdall/types/common.PubKey" json:"pub_key,omitempty"`
	Signer           github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,7,opt,name=signer,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"signer,omitempty"`
	LastUpdated      string                                                        `protobuf:"bytes,8,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Jailed           bool                                                          `protobuf:"varint,9,opt,name=Jailed,proto3" json:"Jailed,omitempty"`
	ProposerPriority int64                                                         `protobuf:"varint,10,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (m *Validator) Reset()      { *m = Validator{} }
func (*Validator) ProtoMessage() {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

type MinimalVal struct {
	// option (gogoproto.goproto_stringer)        = false;
	ID          ValidatorID                                                   `protobuf:"varint,1,opt,name=ID,proto3,enum=heimdall.types.ValidatorID" json:"ID,omitempty"`
	VotingPower uint64                                                        `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	Signer      github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,3,opt,name=signer,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"signer,omitempty"`
}

func (m *MinimalVal) Reset()         { *m = MinimalVal{} }
func (m *MinimalVal) String() string { return proto.CompactTextString(m) }
func (*MinimalVal) ProtoMessage()    {}
func (*MinimalVal) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{2}
}
func (m *MinimalVal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinimalVal.Unmarshal(m, b)
}
func (m *MinimalVal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinimalVal.Marshal(b, m, deterministic)
}
func (m *MinimalVal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinimalVal.Merge(m, src)
}
func (m *MinimalVal) XXX_Size() int {
	return xxx_messageInfo_MinimalVal.Size(m)
}
func (m *MinimalVal) XXX_DiscardUnknown() {
	xxx_messageInfo_MinimalVal.DiscardUnknown(m)
}

var xxx_messageInfo_MinimalVal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("heimdall.types.ValidatorID", ValidatorID_name, ValidatorID_value)
	proto.RegisterType((*ValidatorSet)(nil), "heimdall.types.ValidatorSet")
	proto.RegisterType((*Validator)(nil), "heimdall.types.Validator")
	proto.RegisterType((*MinimalVal)(nil), "heimdall.types.MinimalVal")
}

func init() { proto.RegisterFile("heimdall/types/validator.proto", fileDescriptor_30e71641a391e76e) }

var fileDescriptor_30e71641a391e76e = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x5f, 0x6b, 0xd3, 0x50,
	0x18, 0xc6, 0x93, 0x65, 0x6b, 0xd3, 0x37, 0x65, 0xcc, 0xc3, 0x90, 0xb8, 0x41, 0x1a, 0x87, 0x48,
	0x70, 0x92, 0xc2, 0x54, 0xd0, 0x81, 0x60, 0x47, 0x27, 0xd6, 0x3f, 0x50, 0xa3, 0x2b, 0xe8, 0x4d,
	0x38, 0x6d, 0x0e, 0xed, 0x61, 0x49, 0x4e, 0x38, 0x39, 0xed, 0xe8, 0x37, 0xd8, 0xa5, 0x97, 0x5e,
	0x0e, 0xf4, 0xc2, 0x8f, 0xe0, 0x8d, 0xf7, 0xde, 0xb9, 0x4b, 0xaf, 0x86, 0xb4, 0xdf, 0x62, 0x57,
	0x92, 0x93, 0xb6, 0x6b, 0x07, 0x8a, 0x03, 0xaf, 0x92, 0xf3, 0xfc, 0xce, 0xfb, 0xe4, 0x7d, 0xdf,
	0x87, 0x80, 0xd5, 0x23, 0x34, 0x0a, 0x70, 0x18, 0x56, 0xc5, 0x30, 0x21, 0x69, 0x75, 0x80, 0x43,
	0x1a, 0x60, 0xc1, 0xb8, 0x9b, 0x70, 0x26, 0x18, 0x5a, 0x9d, 0x72, 0x57, 0xf2, 0x8d, 0xf5, 0x2e,
	0xeb, 0x32, 0x89, 0xaa, 0xd9, 0x5b, 0x7e, 0x6b, 0xeb, 0xab, 0x0a, 0xe5, 0xd6, 0xb4, 0xf2, 0x0d,
	0x11, 0xe8, 0x11, 0xc0, 0xcc, 0x29, 0x35, 0x55, 0x5b, 0x73, 0x8c, 0x9d, 0x1b, 0xee, 0xa2, 0x97,
	0x3b, 0xab, 0xf0, 0xe6, 0x2e, 0xa3, 0x07, 0xa0, 0x27, 0x9c, 0x25, 0x2c, 0x25, 0xdc, 0x5c, 0xb2,
	0xd5, 0xbf, 0x17, 0xce, 0xae, 0xa2, 0xbb, 0x80, 0x04, 0x13, 0x38, 0xf4, 0x07, 0x4c, 0xd0, 0xb8,
	0xeb, 0x27, 0xec, 0x88, 0x70, 0x53, 0xb3, 0x55, 0x47, 0xf3, 0xd6, 0x24, 0x69, 0x49, 0xd0, 0xcc,
	0xf4, 0x5d, 0xfd, 0xf8, 0xa4, 0xa2, 0x7c, 0x3c, 0xa9, 0x28, 0x5b, 0x3f, 0x34, 0x28, 0xcd, 0xfc,
	0xd0, 0x36, 0x2c, 0x35, 0xea, 0xa6, 0x6a, 0xab, 0xce, 0xea, 0xce, 0xe6, 0x1f, 0x3f, 0xdb, 0xa8,
	0x7b, 0x4b, 0x8d, 0x3a, 0xaa, 0x80, 0x91, 0x0a, 0xcc, 0x85, 0x4f, 0x12, 0xd6, 0xe9, 0xc9, 0x66,
	0x97, 0x3d, 0x90, 0xd2, 0x7e, 0xa6, 0xa0, 0x4d, 0x28, 0x91, 0x38, 0x98, 0x60, 0x4d, 0x62, 0x9d,
	0xc4, 0x41, 0x0e, 0xd7, 0x61, 0x25, 0x66, 0x71, 0x87, 0x98, 0xcb, 0x12, 0xe4, 0x07, 0x74, 0x13,
	0xca, 0x0b, 0x03, 0xac, 0xc8, 0x01, 0x8c, 0xc1, 0x45, 0xef, 0xe8, 0x35, 0x14, 0x93, 0x7e, 0xdb,
	0x3f, 0x24, 0x43, 0xb3, 0x60, 0xab, 0x4e, 0x79, 0xef, 0xe1, 0xf9, 0x59, 0xe5, 0x7e, 0x97, 0x8a,
	0x5e, 0xbf, 0xed, 0x76, 0x58, 0x54, 0x8d, 0xb0, 0xa0, 0x9d, 0x98, 0x88, 0x23, 0xc6, 0x0f, 0xab,
	0x97, 0xf2, 0xed, 0xb0, 0x28, 0x62, 0xb1, 0xdb, 0xec, 0xb7, 0x5f, 0x90, 0xa1, 0x57, 0x48, 0xe4,
	0x13, 0xbd, 0x83, 0x42, 0x4a, 0xbb, 0x31, 0xe1, 0x66, 0x51, 0x3a, 0xd6, 0xce, 0xcf, 0x2a, 0x8f,
	0xaf, 0xe6, 0xf8, 0x6c, 0xa2, 0xd6, 0x82, 0x80, 0x93, 0x34, 0xf5, 0x26, 0x86, 0xd9, 0x40, 0x21,
	0x4e, 0x85, 0xdf, 0x4f, 0x02, 0x2c, 0x48, 0x60, 0xea, 0xb6, 0xea, 0x94, 0x3c, 0x23, 0xd3, 0x0e,
	0x72, 0x09, 0x5d, 0x87, 0xc2, 0x73, 0x4c, 0x43, 0x12, 0x98, 0x25, 0x5b, 0x75, 0x74, 0x6f, 0x72,
	0x42, 0xdb, 0x70, 0x6d, 0x1a, 0xaf, 0x9f, 0x70, 0xca, 0x38, 0x15, 0x43, 0x13, 0xf2, 0x44, 0xa7,
	0xa0, 0x39, 0xd1, 0xe7, 0x12, 0xfd, 0xa6, 0x02, 0xbc, 0xa2, 0x31, 0x8d, 0x70, 0xd8, 0xc2, 0xe1,
	0xd5, 0x22, 0xbd, 0xbc, 0xfe, 0x3c, 0xd3, 0x85, 0xf5, 0x5f, 0xec, 0x4a, 0xfb, 0xcf, 0xbb, 0xda,
	0x5d, 0xce, 0x66, 0xb8, 0x73, 0x0b, 0x8c, 0xb9, 0xb6, 0x90, 0x01, 0xc5, 0xfa, 0xfe, 0xd3, 0xda,
	0xc1, 0xcb, 0xb7, 0x6b, 0xca, 0x86, 0x7e, 0xfc, 0xc9, 0x52, 0xbe, 0x7c, 0xb6, 0x94, 0xbd, 0x27,
	0xdf, 0x47, 0x96, 0x72, 0x3a, 0xb2, 0x94, 0x5f, 0x23, 0x4b, 0xf9, 0x30, 0xb6, 0x94, 0xd3, 0xb1,
	0xa5, 0xfc, 0x1c, 0x5b, 0xca, 0xfb, 0xdb, 0xff, 0xd6, 0x4c, 0xbb, 0x20, 0xff, 0xdd, 0x7b, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0x22, 0xcf, 0xa7, 0x03, 0x04, 0x00, 0x00,
}