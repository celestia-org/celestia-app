// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qgb/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params represent the Quantum Gravity Bridge genesis and store parameters.
type Params struct {
	DataCommitmentWindow uint64 `protobuf:"varint,1,opt,name=data_commitment_window,json=dataCommitmentWindow,proto3" json:"data_commitment_window,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_afeb526ae8d4446d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDataCommitmentWindow() uint64 {
	if m != nil {
		return m.DataCommitmentWindow
	}
	return 0
}

// GenesisState struct, containing all persistant data required by the Gravity module
type GenesisState struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_afeb526ae8d4446d, []int{1}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

// GravityCounters contains the many noces and counters required to maintain the bridge state in the genesis
type GravityNonces struct {
	// the nonce of the last generated validator set
	LatestValsetNonce uint64 `protobuf:"varint,1,opt,name=latest_valset_nonce,json=latestValsetNonce,proto3" json:"latest_valset_nonce,omitempty"`
}

func (m *GravityNonces) Reset()         { *m = GravityNonces{} }
func (m *GravityNonces) String() string { return proto.CompactTextString(m) }
func (*GravityNonces) ProtoMessage()    {}
func (*GravityNonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_afeb526ae8d4446d, []int{2}
}
func (m *GravityNonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GravityNonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GravityNonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GravityNonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GravityNonces.Merge(m, src)
}
func (m *GravityNonces) XXX_Size() int {
	return m.Size()
}
func (m *GravityNonces) XXX_DiscardUnknown() {
	xxx_messageInfo_GravityNonces.DiscardUnknown(m)
}

var xxx_messageInfo_GravityNonces proto.InternalMessageInfo

func (m *GravityNonces) GetLatestValsetNonce() uint64 {
	if m != nil {
		return m.LatestValsetNonce
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "qgb.Params")
	proto.RegisterType((*GenesisState)(nil), "qgb.GenesisState")
	proto.RegisterType((*GravityNonces)(nil), "qgb.GravityNonces")
}

func init() { proto.RegisterFile("qgb/genesis.proto", fileDescriptor_afeb526ae8d4446d) }

var fileDescriptor_afeb526ae8d4446d = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x13, 0xfe, 0x25, 0xc5, 0xec, 0xaf, 0xb2, 0x71, 0x11, 0x49, 0x31, 0x2c, 0xb1, 0xb1,
	0x31, 0x03, 0xae, 0x95, 0x8d, 0xa0, 0xc2, 0x62, 0x23, 0xb2, 0x82, 0x82, 0x4d, 0x98, 0x64, 0x2f,
	0xe3, 0x40, 0x26, 0x93, 0x64, 0xae, 0xbb, 0x6e, 0xe7, 0x23, 0xf8, 0x58, 0x96, 0x5b, 0x5a, 0x58,
	0x48, 0xf2, 0x22, 0x92, 0x49, 0x48, 0x77, 0xef, 0xfd, 0xce, 0x99, 0x33, 0x1c, 0x32, 0x29, 0x45,
	0xc2, 0x04, 0xe4, 0x60, 0xa4, 0x89, 0x8a, 0x4a, 0xa3, 0xf6, 0xff, 0x95, 0x22, 0x09, 0xa6, 0x42,
	0x0b, 0x6d, 0x77, 0xd6, 0x4e, 0x1d, 0x0a, 0x0e, 0x5a, 0x35, 0x6e, 0x0b, 0xe8, 0xb5, 0xc1, 0x7e,
	0x7b, 0x50, 0x46, 0xf4, 0x7b, 0x78, 0x4b, 0xbc, 0x07, 0x5e, 0x71, 0x65, 0xfc, 0x0b, 0x72, 0xb4,
	0xe2, 0xc8, 0xe3, 0x54, 0x2b, 0x25, 0x51, 0x41, 0x8e, 0xf1, 0x46, 0xe6, 0x2b, 0xbd, 0x39, 0x76,
	0x67, 0xee, 0xe9, 0x68, 0x39, 0x6d, 0xe9, 0xcd, 0x00, 0x9f, 0x2d, 0xbb, 0x1c, 0x7d, 0xfc, 0xcc,
	0x9c, 0x70, 0x4e, 0xfe, 0x2f, 0xba, 0x2f, 0x3d, 0x22, 0x47, 0xf0, 0x4f, 0x88, 0x57, 0xd8, 0x57,
	0xad, 0x77, 0x7c, 0x3e, 0x8e, 0x4a, 0x91, 0x44, 0x5d, 0xd0, 0xb2, 0x47, 0xe1, 0x15, 0xd9, 0x5b,
	0x54, 0x7c, 0x2d, 0x71, 0x7b, 0xaf, 0xf3, 0x14, 0x8c, 0x1f, 0x91, 0xc3, 0x8c, 0x23, 0x18, 0x8c,
	0xd7, 0x3c, 0x33, 0x80, 0x71, 0xde, 0xde, 0xfb, 0xf8, 0x49, 0x87, 0x9e, 0x2c, 0xb1, 0x86, 0xeb,
	0xbb, 0xaf, 0x9a, 0xba, 0xbb, 0x9a, 0xba, 0xbf, 0x35, 0x75, 0x3f, 0x1b, 0xea, 0xec, 0x1a, 0xea,
	0x7c, 0x37, 0xd4, 0x79, 0x61, 0x42, 0xe2, 0xeb, 0x5b, 0x12, 0xa5, 0x5a, 0xb1, 0x14, 0x32, 0x30,
	0x28, 0xb9, 0xae, 0xc4, 0x30, 0x9f, 0xf1, 0xa2, 0x60, 0xef, 0x6c, 0x28, 0x27, 0xf1, 0x6c, 0x1b,
	0xf3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x5f, 0x53, 0x3b, 0x5e, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DataCommitmentWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.DataCommitmentWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if m.Params != nil {
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
	}
	return len(dAtA) - i, nil
}

func (m *GravityNonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GravityNonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GravityNonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestValsetNonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LatestValsetNonce))
		i--
		dAtA[i] = 0x8
	}
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DataCommitmentWindow != 0 {
		n += 1 + sovGenesis(uint64(m.DataCommitmentWindow))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GravityNonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestValsetNonce != 0 {
		n += 1 + sovGenesis(uint64(m.LatestValsetNonce))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataCommitmentWindow", wireType)
			}
			m.DataCommitmentWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataCommitmentWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *GravityNonces) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GravityNonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GravityNonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestValsetNonce", wireType)
			}
			m.LatestValsetNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestValsetNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
