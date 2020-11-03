// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the auth module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// accounts are the accounts present at genesis.
	Accounts []*GenesisAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57ded559541890b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAccounts() []*GenesisAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Params defines the parameters for the auth module.
type Params struct {
	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty" yaml:"max_memo_characters"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty" yaml:"tx_sig_limit"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty" yaml:"tx_size_cost_per_byte"`
	SigVerifyCostED25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty" yaml:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256k1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty" yaml:"sig_verify_cost_secp256k1"`
	MaxTxGas               uint64 `protobuf:"varint,6,opt,name=max_tx_gas,json=maxTxGas,proto3" json:"max_tx_gas,omitempty" yaml:"max_tx_gas"`
	TxFees                 string `protobuf:"bytes,7,opt,name=tx_fees,json=txFees,proto3" json:"tx_fees,omitempty" yaml:"tx_fees"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57ded559541890b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxMemoCharacters() uint64 {
	if m != nil {
		return m.MaxMemoCharacters
	}
	return 0
}

func (m *Params) GetTxSigLimit() uint64 {
	if m != nil {
		return m.TxSigLimit
	}
	return 0
}

func (m *Params) GetTxSizeCostPerByte() uint64 {
	if m != nil {
		return m.TxSizeCostPerByte
	}
	return 0
}

func (m *Params) GetSigVerifyCostED25519() uint64 {
	if m != nil {
		return m.SigVerifyCostED25519
	}
	return 0
}

func (m *Params) GetSigVerifyCostSecp256k1() uint64 {
	if m != nil {
		return m.SigVerifyCostSecp256k1
	}
	return 0
}

func (m *Params) GetMaxTxGas() uint64 {
	if m != nil {
		return m.MaxTxGas
	}
	return 0
}

func (m *Params) GetTxFees() string {
	if m != nil {
		return m.TxFees
	}
	return ""
}

type GenesisAccount struct {
	Address           github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"address,omitempty"`
	Coins             github_com_cosmos_cosmos_sdk_types.Coins                      `protobuf:"bytes,2,opt,name=coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins,omitempty"`
	Sequence          uint64                                                        `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	AccountNumber     uint64                                                        `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	ModuleName        string                                                        `protobuf:"bytes,5,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ModulePermissions []string                                                      `protobuf:"bytes,6,rep,name=module_permissions,json=modulePermissions,proto3" json:"module_permissions,omitempty"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57ded559541890b, []int{2}
}
func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisAccount.Unmarshal(m, b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return xxx_messageInfo_GenesisAccount.Size(m)
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "heimdall.auth.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "heimdall.auth.v1beta1.Params")
	proto.RegisterType((*GenesisAccount)(nil), "heimdall.auth.v1beta1.GenesisAccount")
}

func init() { proto.RegisterFile("auth/v1beta/genesis.proto", fileDescriptor_d57ded559541890b) }

var fileDescriptor_d57ded559541890b = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x9d, 0x36, 0x9b, 0x26, 0xd3, 0xa5, 0x52, 0x66, 0xdb, 0x5d, 0x6f, 0x05, 0x9e, 0xc8,
	0x52, 0xa5, 0x48, 0x50, 0x5b, 0x49, 0x55, 0xa4, 0x06, 0x81, 0x14, 0x17, 0x28, 0x07, 0xa8, 0xaa,
	0x09, 0xe2, 0x00, 0x07, 0x6b, 0xe2, 0x4c, 0x1d, 0xab, 0x19, 0x4f, 0xea, 0x99, 0x14, 0xa7, 0x9f,
	0x80, 0x13, 0xe2, 0xc8, 0xb1, 0x5f, 0x85, 0x5b, 0xc5, 0xa9, 0x47, 0x4e, 0x16, 0x4a, 0x2f, 0x9c,
	0x73, 0xec, 0x09, 0x79, 0xc6, 0x4d, 0xff, 0x28, 0xd5, 0x9e, 0x12, 0xbf, 0xef, 0xf3, 0xfc, 0x9e,
	0x64, 0xde, 0xd7, 0x03, 0xde, 0x93, 0x89, 0x1c, 0xba, 0x17, 0xad, 0x3e, 0x95, 0xc4, 0x0d, 0x69,
	0x4c, 0x45, 0x24, 0x9c, 0x71, 0xc2, 0x25, 0x87, 0x5b, 0x43, 0x1a, 0xb1, 0x01, 0x19, 0x8d, 0x9c,
	0x5c, 0xe3, 0x68, 0x4d, 0x6b, 0x7b, 0x33, 0xe4, 0x21, 0x57, 0x0a, 0x37, 0xff, 0xa6, 0xc5, 0xf6,
	0xef, 0x25, 0xf0, 0xfa, 0x48, 0xdb, 0x7b, 0x92, 0x48, 0x0a, 0xbf, 0x00, 0x95, 0x31, 0x49, 0x08,
	0x13, 0x66, 0xa9, 0x51, 0x6a, 0xae, 0xb7, 0x3f, 0x71, 0x96, 0xe2, 0x9c, 0x13, 0x25, 0xf2, 0xca,
	0xd7, 0x19, 0x32, 0x70, 0x61, 0x81, 0x5d, 0x50, 0x25, 0x41, 0xc0, 0x27, 0xb1, 0x14, 0xe6, 0x4a,
	0x63, 0xb5, 0xb9, 0xde, 0xde, 0x79, 0xc1, 0x5e, 0x64, 0x76, 0xb5, 0x1a, 0x2f, 0x6c, 0xf6, 0x5f,
	0x65, 0x50, 0xd1, 0x6c, 0x78, 0x0c, 0xde, 0x30, 0x92, 0xfa, 0x8c, 0x32, 0xee, 0x07, 0x43, 0x92,
	0x90, 0x40, 0xd2, 0x44, 0xff, 0xae, 0xb2, 0x67, 0xcd, 0x33, 0xb4, 0x3d, 0x25, 0x6c, 0xd4, 0xb1,
	0x97, 0x88, 0x6c, 0x5c, 0x67, 0x24, 0xfd, 0x81, 0x32, 0x7e, 0xb8, 0xa8, 0xc1, 0x03, 0xf0, 0x5a,
	0xa6, 0xbe, 0x88, 0x42, 0x7f, 0x14, 0xb1, 0x48, 0x9a, 0x2b, 0x0a, 0xf4, 0x6e, 0x9e, 0xa1, 0x37,
	0x1a, 0xf4, 0xb8, 0x6b, 0x63, 0x20, 0xd3, 0x5e, 0x14, 0x7e, 0x9f, 0x3f, 0x40, 0x0c, 0xb6, 0x54,
	0xf3, 0x92, 0xfa, 0x01, 0x17, 0xd2, 0x1f, 0xd3, 0xc4, 0xef, 0x4f, 0x25, 0x35, 0x57, 0x15, 0xa3,
	0x31, 0xcf, 0xd0, 0xc7, 0x8f, 0x18, 0xcf, 0x65, 0x36, 0xae, 0xe7, 0xb0, 0x4b, 0x7a, 0xc8, 0x85,
	0x3c, 0xa1, 0x89, 0x37, 0x95, 0x14, 0x9e, 0x83, 0x77, 0x79, 0xda, 0x05, 0x4d, 0xa2, 0xd3, 0xa9,
	0xd6, 0xd3, 0x41, 0x7b, 0x7f, 0xbf, 0x75, 0x60, 0x96, 0x15, 0xb5, 0x33, 0xcb, 0xd0, 0x66, 0x2f,
	0x0a, 0x7f, 0x52, 0x8a, 0xdc, 0xfa, 0xcd, 0xd7, 0xaa, 0x3f, 0xcf, 0x90, 0xa5, 0xd3, 0x5e, 0x00,
	0xd8, 0x78, 0x53, 0x3c, 0xf1, 0xe9, 0x32, 0x9c, 0x82, 0xf7, 0xcf, 0x1d, 0x82, 0x06, 0xe3, 0xf6,
	0xfe, 0xe7, 0x67, 0x2d, 0xf3, 0x95, 0x0a, 0xfd, 0x6a, 0x96, 0xa1, 0xb7, 0x4f, 0x42, 0x7b, 0xf7,
	0x8a, 0x79, 0x86, 0x1a, 0xcb, 0x63, 0x17, 0x10, 0x1b, 0xbf, 0x15, 0x4b, 0xbd, 0x70, 0x0f, 0x80,
	0x7c, 0x4e, 0x32, 0xf5, 0x43, 0x22, 0xcc, 0x8a, 0xca, 0xda, 0x9a, 0x67, 0xa8, 0xfe, 0x30, 0x43,
	0xdd, 0xb3, 0x71, 0x95, 0x91, 0xf4, 0xc7, 0xf4, 0x88, 0x08, 0xf8, 0x29, 0x58, 0x93, 0xa9, 0x7f,
	0x4a, 0xa9, 0x30, 0xd7, 0x1a, 0xa5, 0x66, 0xcd, 0x83, 0xf3, 0x0c, 0x6d, 0x2c, 0x0e, 0x3a, 0x6f,
	0xd8, 0xb8, 0x22, 0xd3, 0x6f, 0x29, 0x15, 0x9d, 0xea, 0x9f, 0x57, 0xc8, 0xf8, 0xef, 0x0a, 0x19,
	0xf6, 0xdf, 0x2b, 0x60, 0xe3, 0xe9, 0x82, 0xc1, 0x5f, 0xc0, 0x1a, 0x19, 0x0c, 0x12, 0x2a, 0xf4,
	0xfe, 0xd4, 0xbc, 0xee, 0x5d, 0x86, 0xbe, 0x0c, 0x23, 0x39, 0x9c, 0xf4, 0x9d, 0x80, 0x33, 0x97,
	0x11, 0x19, 0x05, 0x31, 0x95, 0xbf, 0xf2, 0xe4, 0xcc, 0xbd, 0xdf, 0x59, 0x57, 0x4e, 0xc7, 0x54,
	0xb8, 0x01, 0x67, 0x8c, 0xc7, 0xce, 0x77, 0x45, 0xb5, 0xab, 0x41, 0xf8, 0x9e, 0x08, 0x3d, 0xf0,
	0x2a, 0xe0, 0x51, 0x2c, 0xd4, 0x46, 0xd5, 0xbc, 0xcf, 0xee, 0x32, 0xd4, 0x7c, 0x84, 0x0e, 0xb8,
	0x60, 0x5c, 0x14, 0x1f, 0xbb, 0x62, 0x70, 0xa6, 0xb1, 0xce, 0x61, 0xee, 0xc1, 0xda, 0x0a, 0xb7,
	0x41, 0x55, 0xd0, 0xf3, 0x09, 0x8d, 0x83, 0x62, 0xa9, 0xf0, 0xe2, 0x19, 0xee, 0x80, 0x8d, 0xe2,
	0xfd, 0xf0, 0xe3, 0x09, 0xeb, 0xd3, 0x44, 0x2f, 0x08, 0xfe, 0xa8, 0xa8, 0x1e, 0xab, 0x22, 0x44,
	0x60, 0x9d, 0xf1, 0xc1, 0x64, 0x44, 0xfd, 0x98, 0x30, 0xaa, 0xe6, 0x59, 0xc3, 0x40, 0x97, 0x8e,
	0x09, 0xa3, 0x70, 0x17, 0xc0, 0x42, 0x30, 0xa6, 0x09, 0x8b, 0x84, 0x88, 0x78, 0x9c, 0xcf, 0x62,
	0xb5, 0x59, 0xc3, 0x75, 0xdd, 0x39, 0x79, 0x68, 0x74, 0xca, 0xbf, 0x5d, 0x21, 0xc3, 0x3b, 0xba,
	0x9e, 0x59, 0xc6, 0xcd, 0xcc, 0x32, 0xfe, 0x9d, 0x59, 0xc6, 0x1f, 0xb7, 0x96, 0x71, 0x73, 0x6b,
	0x19, 0xff, 0xdc, 0x5a, 0xc6, 0xcf, 0xbb, 0x1f, 0x3c, 0xbe, 0xd4, 0x55, 0xd7, 0x94, 0xfa, 0xbb,
	0xfd, 0x8a, 0xba, 0x71, 0xf6, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xba, 0xdd, 0x75, 0x39, 0xbb,
	0x04, 0x00, 0x00,
}