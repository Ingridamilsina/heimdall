// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/maticnetwork/heimdall/types"
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

// GenesisState defines the staking module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to staking
	Params           Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Validators       []*types.Validator  `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
	CurrentValSet    *types.ValidatorSet `protobuf:"bytes,3,opt,name=current_val_set,json=currentValSet,proto3" json:"current_val_set,omitempty" yaml:"current_val_set"`
	StakingSequences []string            `protobuf:"bytes,4,rep,name=staking_sequences,json=stakingSequences,proto3" json:"staking_sequences,omitempty" yaml:"staking_sequences"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_30376b0921a07e54, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
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

func (m *GenesisState) GetValidators() []*types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GenesisState) GetCurrentValSet() *types.ValidatorSet {
	if m != nil {
		return m.CurrentValSet
	}
	return nil
}

func (m *GenesisState) GetStakingSequences() []string {
	if m != nil {
		return m.StakingSequences
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "heimdall.staking.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("staking/v1beta1/genesis.proto", fileDescriptor_30376b0921a07e54) }

var fileDescriptor_30376b0921a07e54 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x5b, 0x20, 0x24, 0xb7, 0xdc, 0x9b, 0x7b, 0x6f, 0x63, 0x4c, 0x6d, 0xb0, 0x34, 0x5d,
	0xb1, 0x9a, 0x11, 0x5c, 0xe9, 0xc2, 0x45, 0x37, 0x46, 0x57, 0xa6, 0x4d, 0x58, 0xb8, 0xc1, 0xa1,
	0x9c, 0x94, 0x86, 0xb6, 0x83, 0x9d, 0x03, 0xca, 0x5b, 0xf8, 0x34, 0x3e, 0x03, 0x4b, 0x96, 0xae,
	0x88, 0x81, 0x37, 0xe0, 0x09, 0x8c, 0xed, 0x14, 0x0d, 0x86, 0xdd, 0xc9, 0x3f, 0xdf, 0xff, 0xcf,
	0x7f, 0x72, 0xb4, 0x53, 0x81, 0x6c, 0x1c, 0xa5, 0x21, 0x9d, 0x75, 0x06, 0x80, 0xac, 0x43, 0x43,
	0x48, 0x41, 0x44, 0x82, 0x4c, 0x32, 0x8e, 0x5c, 0x37, 0x46, 0x10, 0x25, 0x43, 0x16, 0xc7, 0x44,
	0x72, 0x44, 0x72, 0xa6, 0x55, 0xbe, 0x50, 0x9c, 0x4f, 0x40, 0xd0, 0x19, 0x8b, 0xa3, 0x21, 0x43,
	0x9e, 0x15, 0x4e, 0xb3, 0xb9, 0x1f, 0x3c, 0x61, 0x19, 0x4b, 0x64, 0xae, 0x79, 0x14, 0xf2, 0x90,
	0xe7, 0x23, 0xfd, 0x9c, 0x0a, 0xd5, 0x79, 0xad, 0x68, 0xbf, 0xaf, 0x8b, 0xff, 0x7d, 0x64, 0x08,
	0xfa, 0x95, 0x56, 0x2f, 0x6c, 0x86, 0x6a, 0xab, 0xed, 0x46, 0xd7, 0x26, 0x87, 0xfa, 0x90, 0xbb,
	0x9c, 0x73, 0x6b, 0x8b, 0x55, 0x4b, 0xf1, 0xa4, 0x4b, 0xbf, 0xd0, 0xb4, 0x5d, 0x2f, 0x61, 0x54,
	0xec, 0x6a, 0xbb, 0xd1, 0x3d, 0xf9, 0xca, 0xc8, 0x9b, 0x93, 0x5e, 0x49, 0x78, 0xdf, 0x60, 0xfd,
	0x41, 0xfb, 0x1b, 0x4c, 0xb3, 0x0c, 0x52, 0xec, 0xcf, 0x58, 0xdc, 0x17, 0x80, 0x46, 0x35, 0xef,
	0xd0, 0x3c, 0xe8, 0xf7, 0x01, 0x5d, 0x73, 0xbb, 0x6a, 0x1d, 0xcf, 0x59, 0x12, 0x5f, 0x3a, 0x7b,
	0x76, 0xc7, 0xfb, 0x23, 0x95, 0x1e, 0x8b, 0x7d, 0x40, 0xfd, 0x46, 0xfb, 0x2f, 0x97, 0xe8, 0x0b,
	0x78, 0x9c, 0x42, 0x1a, 0x80, 0x30, 0x6a, 0x76, 0xb5, 0xfd, 0xcb, 0x6d, 0x6e, 0x57, 0x2d, 0xa3,
	0x48, 0xf9, 0x81, 0x38, 0xde, 0x3f, 0xa9, 0xf9, 0xa5, 0xe4, 0xde, 0x2e, 0xd6, 0x96, 0xba, 0x5c,
	0x5b, 0xea, 0xfb, 0xda, 0x52, 0x5f, 0x36, 0x96, 0xb2, 0xdc, 0x58, 0xca, 0xdb, 0xc6, 0x52, 0xee,
	0xcf, 0xc2, 0x08, 0x47, 0xd3, 0x01, 0x09, 0x78, 0x42, 0x13, 0x86, 0x51, 0x90, 0x02, 0x3e, 0xf1,
	0x6c, 0x4c, 0x77, 0xe7, 0x7b, 0xa6, 0xe5, 0xa5, 0xf2, 0x75, 0x06, 0xf5, 0xfc, 0x16, 0xe7, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xce, 0xcc, 0x2b, 0x1a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakingSequences) > 0 {
		for iNdEx := len(m.StakingSequences) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StakingSequences[iNdEx])
			copy(dAtA[i:], m.StakingSequences[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakingSequences[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.CurrentValSet != nil {
		{
			size, err := m.CurrentValSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CurrentValSet != nil {
		l = m.CurrentValSet.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.StakingSequences) > 0 {
		for _, s := range m.StakingSequences {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &types.Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentValSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrentValSet == nil {
				m.CurrentValSet = &types.ValidatorSet{}
			}
			if err := m.CurrentValSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingSequences", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakingSequences = append(m.StakingSequences, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
