// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blob/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgPayForBlobs pays for the inclusion of a blob in the block.
type MsgPayForBlobs struct {
	Signer       string   `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NamespaceIds [][]byte `protobuf:"bytes,2,rep,name=namespace_ids,json=namespaceIds,proto3" json:"namespace_ids,omitempty"`
	BlobSizes    []uint32 `protobuf:"varint,3,rep,packed,name=blob_sizes,json=blobSizes,proto3" json:"blob_sizes,omitempty"`
	// share_commitments is a list of share commitments (one per blob).
	ShareCommitments [][]byte `protobuf:"bytes,4,rep,name=share_commitments,json=shareCommitments,proto3" json:"share_commitments,omitempty"`
	// share_versions are the versions of the share format that the blobs
	// associated with this message should use when included in a block. The
	// share_versions specified must match the share_versions used to generate the
	// share_commitment in this message.
	ShareVersions []uint32 `protobuf:"varint,8,rep,packed,name=share_versions,json=shareVersions,proto3" json:"share_versions,omitempty"`
}

func (m *MsgPayForBlobs) Reset()         { *m = MsgPayForBlobs{} }
func (m *MsgPayForBlobs) String() string { return proto.CompactTextString(m) }
func (*MsgPayForBlobs) ProtoMessage()    {}
func (*MsgPayForBlobs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{0}
}
func (m *MsgPayForBlobs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPayForBlobs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPayForBlobs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPayForBlobs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayForBlobs.Merge(m, src)
}
func (m *MsgPayForBlobs) XXX_Size() int {
	return m.Size()
}
func (m *MsgPayForBlobs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayForBlobs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayForBlobs proto.InternalMessageInfo

func (m *MsgPayForBlobs) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgPayForBlobs) GetNamespaceIds() [][]byte {
	if m != nil {
		return m.NamespaceIds
	}
	return nil
}

func (m *MsgPayForBlobs) GetBlobSizes() []uint32 {
	if m != nil {
		return m.BlobSizes
	}
	return nil
}

func (m *MsgPayForBlobs) GetShareCommitments() [][]byte {
	if m != nil {
		return m.ShareCommitments
	}
	return nil
}

func (m *MsgPayForBlobs) GetShareVersions() []uint32 {
	if m != nil {
		return m.ShareVersions
	}
	return nil
}

// MsgPayForBlobsResponse describes the response returned after the submission
// of a PayForBlobs
type MsgPayForBlobsResponse struct {
}

func (m *MsgPayForBlobsResponse) Reset()         { *m = MsgPayForBlobsResponse{} }
func (m *MsgPayForBlobsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPayForBlobsResponse) ProtoMessage()    {}
func (*MsgPayForBlobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{1}
}
func (m *MsgPayForBlobsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPayForBlobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPayForBlobsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPayForBlobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayForBlobsResponse.Merge(m, src)
}
func (m *MsgPayForBlobsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPayForBlobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayForBlobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayForBlobsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPayForBlobs)(nil), "blob.MsgPayForBlobs")
	proto.RegisterType((*MsgPayForBlobsResponse)(nil), "blob.MsgPayForBlobsResponse")
}

func init() { proto.RegisterFile("blob/tx.proto", fileDescriptor_f945cb94fe124aae) }

var fileDescriptor_f945cb94fe124aae = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x2d, 0xcb, 0x98, 0x7a, 0x6b, 0x99, 0x76, 0x31, 0x46, 0x08, 0x57, 0x18, 0x95, 0x82,
	0xa1, 0x54, 0x2a, 0xed, 0x1b, 0xb8, 0x50, 0x68, 0xc1, 0x10, 0x14, 0xc8, 0x21, 0x17, 0xb3, 0x92,
	0x37, 0xeb, 0x05, 0x49, 0xb3, 0x68, 0x36, 0xc1, 0xce, 0x31, 0x4f, 0x10, 0xc8, 0x0b, 0xe5, 0x98,
	0xa3, 0x21, 0x97, 0x1c, 0x83, 0x9d, 0x07, 0x09, 0x2b, 0x39, 0x4e, 0x0c, 0xb9, 0xfd, 0xfb, 0xfd,
	0x3b, 0x3f, 0xff, 0x30, 0xc4, 0x49, 0x32, 0x48, 0x22, 0xbd, 0x0c, 0x55, 0x09, 0x1a, 0x68, 0xcb,
	0x3c, 0xbd, 0xbe, 0x00, 0x01, 0x15, 0x88, 0x8c, 0xaa, 0x3d, 0x6f, 0x28, 0x00, 0x44, 0xc6, 0x23,
	0xa6, 0x64, 0xc4, 0x8a, 0x02, 0x34, 0xd3, 0x12, 0x0a, 0xac, 0xdd, 0xe0, 0xd6, 0x22, 0xbd, 0x29,
	0x8a, 0x23, 0xb6, 0xfa, 0x0b, 0xe5, 0x24, 0x83, 0x04, 0xe9, 0x80, 0xb4, 0x51, 0x8a, 0x82, 0x97,
	0xae, 0x35, 0xb2, 0xc6, 0x9d, 0x78, 0xf7, 0xa2, 0x5f, 0x89, 0x53, 0xb0, 0x9c, 0xa3, 0x62, 0x29,
	0x9f, 0xc9, 0x39, 0xba, 0xcd, 0x91, 0x3d, 0xee, 0xc6, 0xdd, 0x3d, 0xfc, 0x37, 0x47, 0xfa, 0x85,
	0x10, 0xd3, 0x65, 0x86, 0xf2, 0x92, 0xa3, 0x6b, 0x8f, 0xec, 0xb1, 0x13, 0x77, 0x0c, 0x39, 0x36,
	0x80, 0x7e, 0x27, 0x9f, 0x71, 0xc1, 0x4a, 0x3e, 0x4b, 0x21, 0xcf, 0xa5, 0xce, 0x79, 0xa1, 0xd1,
	0x6d, 0x55, 0x39, 0x9f, 0x2a, 0xe3, 0xcf, 0x2b, 0xa7, 0xdf, 0x48, 0xaf, 0xfe, 0x7c, 0xc1, 0x4b,
	0x34, 0x9d, 0xdd, 0x0f, 0x55, 0x9e, 0x53, 0xd1, 0x93, 0x1d, 0x0c, 0x5c, 0x32, 0x38, 0xdc, 0x20,
	0xe6, 0xa8, 0xa0, 0x40, 0xfe, 0x2b, 0x23, 0xf6, 0x14, 0x05, 0xe5, 0xe4, 0xe3, 0xdb, 0xfd, 0xfa,
	0xa1, 0xe9, 0x13, 0x1e, 0xce, 0x78, 0xc3, 0xf7, 0xe8, 0x4b, 0x52, 0x10, 0x5c, 0xdd, 0x3f, 0xdd,
	0x34, 0x87, 0xd4, 0x8b, 0x52, 0x9e, 0x71, 0xd4, 0x92, 0x45, 0xd5, 0x05, 0x14, 0x5b, 0x9d, 0x41,
	0x69, 0x24, 0x4e, 0xfe, 0xdf, 0x6d, 0x7c, 0x6b, 0xbd, 0xf1, 0xad, 0xc7, 0x8d, 0x6f, 0x5d, 0x6f,
	0xfd, 0xc6, 0x7a, 0xeb, 0x37, 0x1e, 0xb6, 0x7e, 0xe3, 0xf4, 0xa7, 0x90, 0x7a, 0x71, 0x9e, 0x84,
	0x29, 0xe4, 0xfb, 0x79, 0x28, 0xc5, 0x5e, 0xff, 0x60, 0x4a, 0x45, 0xcb, 0x3a, 0x51, 0xaf, 0x14,
	0xc7, 0xa4, 0x5d, 0x5d, 0xe7, 0xf7, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x3f, 0xb4, 0x67,
	0xe8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PayForBlobs allows the user to pay for the inclusion of one or more blobs
	PayForBlobs(ctx context.Context, in *MsgPayForBlobs, opts ...grpc.CallOption) (*MsgPayForBlobsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PayForBlobs(ctx context.Context, in *MsgPayForBlobs, opts ...grpc.CallOption) (*MsgPayForBlobsResponse, error) {
	out := new(MsgPayForBlobsResponse)
	err := c.cc.Invoke(ctx, "/blob.Msg/PayForBlobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PayForBlobs allows the user to pay for the inclusion of one or more blobs
	PayForBlobs(context.Context, *MsgPayForBlobs) (*MsgPayForBlobsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PayForBlobs(ctx context.Context, req *MsgPayForBlobs) (*MsgPayForBlobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayForBlobs not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PayForBlobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPayForBlobs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PayForBlobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.Msg/PayForBlobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PayForBlobs(ctx, req.(*MsgPayForBlobs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blob.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayForBlobs",
			Handler:    _Msg_PayForBlobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blob/tx.proto",
}

func (m *MsgPayForBlobs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPayForBlobs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPayForBlobs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShareVersions) > 0 {
		dAtA2 := make([]byte, len(m.ShareVersions)*10)
		var j1 int
		for _, num := range m.ShareVersions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ShareCommitments) > 0 {
		for iNdEx := len(m.ShareCommitments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ShareCommitments[iNdEx])
			copy(dAtA[i:], m.ShareCommitments[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.ShareCommitments[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BlobSizes) > 0 {
		dAtA4 := make([]byte, len(m.BlobSizes)*10)
		var j3 int
		for _, num := range m.BlobSizes {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintTx(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NamespaceIds) > 0 {
		for iNdEx := len(m.NamespaceIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NamespaceIds[iNdEx])
			copy(dAtA[i:], m.NamespaceIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.NamespaceIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPayForBlobsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPayForBlobsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPayForBlobsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPayForBlobs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.NamespaceIds) > 0 {
		for _, b := range m.NamespaceIds {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.BlobSizes) > 0 {
		l = 0
		for _, e := range m.BlobSizes {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	if len(m.ShareCommitments) > 0 {
		for _, b := range m.ShareCommitments {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.ShareVersions) > 0 {
		l = 0
		for _, e := range m.ShareVersions {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgPayForBlobsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPayForBlobs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPayForBlobs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPayForBlobs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceIds", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceIds = append(m.NamespaceIds, make([]byte, postIndex-iNdEx))
			copy(m.NamespaceIds[len(m.NamespaceIds)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BlobSizes = append(m.BlobSizes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BlobSizes) == 0 {
					m.BlobSizes = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BlobSizes = append(m.BlobSizes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BlobSizes", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCommitments", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareCommitments = append(m.ShareCommitments, make([]byte, postIndex-iNdEx))
			copy(m.ShareCommitments[len(m.ShareCommitments)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ShareVersions = append(m.ShareVersions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ShareVersions) == 0 {
					m.ShareVersions = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ShareVersions = append(m.ShareVersions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareVersions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPayForBlobsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPayForBlobsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPayForBlobsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
