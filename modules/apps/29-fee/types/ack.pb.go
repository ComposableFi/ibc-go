// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/ack.proto

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

// IncentivizedAcknowledgement is the acknowledgement format to be used by applications wrapped in the fee middleware
type IncentivizedAcknowledgement struct {
	// the underlying app acknowledgement bytes
	AppAcknowledgement []byte `protobuf:"bytes,1,opt,name=app_acknowledgement,json=appAcknowledgement,proto3" json:"app_acknowledgement,omitempty" yaml:"app_acknowledgement"`
	// the relayer address which submits the recv packet message
	ForwardRelayerAddress string `protobuf:"bytes,2,opt,name=forward_relayer_address,json=forwardRelayerAddress,proto3" json:"forward_relayer_address,omitempty" yaml:"forward_relayer_address"`
	// success flag of the base application callback
	UnderlyingAppSuccess bool `protobuf:"varint,3,opt,name=underlying_app_success,json=underlyingAppSuccess,proto3" json:"underlying_app_success,omitempty" yaml:"underlying_app_successl"`
}

func (m *IncentivizedAcknowledgement) Reset()         { *m = IncentivizedAcknowledgement{} }
func (m *IncentivizedAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*IncentivizedAcknowledgement) ProtoMessage()    {}
func (*IncentivizedAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2834946fb65ea4, []int{0}
}
func (m *IncentivizedAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentivizedAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentivizedAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentivizedAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentivizedAcknowledgement.Merge(m, src)
}
func (m *IncentivizedAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *IncentivizedAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentivizedAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_IncentivizedAcknowledgement proto.InternalMessageInfo

func (m *IncentivizedAcknowledgement) GetAppAcknowledgement() []byte {
	if m != nil {
		return m.AppAcknowledgement
	}
	return nil
}

func (m *IncentivizedAcknowledgement) GetForwardRelayerAddress() string {
	if m != nil {
		return m.ForwardRelayerAddress
	}
	return ""
}

func (m *IncentivizedAcknowledgement) GetUnderlyingAppSuccess() bool {
	if m != nil {
		return m.UnderlyingAppSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*IncentivizedAcknowledgement)(nil), "ibc.applications.fee.v1.IncentivizedAcknowledgement")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/ack.proto", fileDescriptor_ab2834946fb65ea4) }

var fileDescriptor_ab2834946fb65ea4 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0xc2, 0x40,
	0x1c, 0xc6, 0x29, 0x26, 0x46, 0x1b, 0xa7, 0x8a, 0x42, 0x30, 0x39, 0xb0, 0x13, 0x0b, 0xbd, 0xa0,
	0x61, 0xd0, 0x0d, 0x36, 0x27, 0x92, 0xba, 0x18, 0x96, 0xe6, 0x7a, 0xfd, 0x53, 0x2f, 0x5c, 0xef,
	0x2e, 0x77, 0x6d, 0x49, 0x7d, 0x0a, 0x1f, 0xc2, 0x87, 0x71, 0x64, 0x74, 0x22, 0x06, 0xde, 0x80,
	0x27, 0x30, 0xa5, 0x26, 0xa2, 0xc1, 0xed, 0xf2, 0x7d, 0xbf, 0xfc, 0x72, 0xf9, 0x7f, 0xf6, 0x35,
	0x0b, 0x29, 0x26, 0x4a, 0x71, 0x46, 0x49, 0xca, 0xa4, 0x30, 0x78, 0x06, 0x80, 0xf3, 0x01, 0x26,
	0x74, 0xee, 0x29, 0x2d, 0x53, 0xe9, 0x34, 0x59, 0x48, 0xbd, 0x7d, 0xc4, 0x9b, 0x01, 0x78, 0xf9,
	0xa0, 0xdd, 0x88, 0x65, 0x2c, 0x77, 0x0c, 0x2e, 0x5f, 0x15, 0xee, 0xbe, 0xd5, 0xed, 0xab, 0x07,
	0x41, 0x41, 0xa4, 0x2c, 0x67, 0x2f, 0x10, 0x8d, 0xe8, 0x5c, 0xc8, 0x05, 0x87, 0x28, 0x86, 0x04,
	0x44, 0xea, 0x4c, 0xec, 0x73, 0xa2, 0x54, 0x40, 0x7e, 0xc7, 0x2d, 0xab, 0x6b, 0xf5, 0xce, 0xc6,
	0x68, 0xbb, 0xea, 0xb4, 0x0b, 0x92, 0xf0, 0x7b, 0xf7, 0x00, 0xe4, 0xfa, 0x0e, 0x51, 0xea, 0xaf,
	0x70, 0x6a, 0x37, 0x67, 0x52, 0x2f, 0x88, 0x8e, 0x02, 0x0d, 0x9c, 0x14, 0xa0, 0x03, 0x12, 0x45,
	0x1a, 0x8c, 0x69, 0xd5, 0xbb, 0x56, 0xef, 0x74, 0xec, 0x6e, 0x57, 0x1d, 0x54, 0x49, 0xff, 0x01,
	0x5d, 0xff, 0xe2, 0xbb, 0xf1, 0xab, 0x62, 0x54, 0xe5, 0xce, 0x93, 0x7d, 0x99, 0x89, 0x08, 0x34,
	0x2f, 0x98, 0x88, 0x83, 0xf2, 0x4b, 0x26, 0xa3, 0xb4, 0x54, 0x1f, 0x75, 0xad, 0xde, 0xc9, 0xbe,
	0xfa, 0x30, 0xc7, 0x5d, 0xbf, 0xf1, 0xd3, 0x8c, 0x94, 0x7a, 0xac, 0xf2, 0xf1, 0xe4, 0x7d, 0x8d,
	0xac, 0xe5, 0x1a, 0x59, 0x9f, 0x6b, 0x64, 0xbd, 0x6e, 0x50, 0x6d, 0xb9, 0x41, 0xb5, 0x8f, 0x0d,
	0xaa, 0x4d, 0x87, 0x31, 0x4b, 0x9f, 0xb3, 0xd0, 0xa3, 0x32, 0xc1, 0x54, 0x9a, 0x44, 0x1a, 0xcc,
	0x42, 0xda, 0x8f, 0x25, 0xce, 0x87, 0x38, 0x91, 0x51, 0xc6, 0xc1, 0x94, 0x93, 0x19, 0x7c, 0x73,
	0xd7, 0x2f, 0xd7, 0x4a, 0x0b, 0x05, 0x26, 0x3c, 0xde, 0x9d, 0xff, 0xf6, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x92, 0x0d, 0x35, 0xd2, 0x01, 0x00, 0x00,
}

func (m *IncentivizedAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentivizedAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentivizedAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnderlyingAppSuccess {
		i--
		if m.UnderlyingAppSuccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ForwardRelayerAddress) > 0 {
		i -= len(m.ForwardRelayerAddress)
		copy(dAtA[i:], m.ForwardRelayerAddress)
		i = encodeVarintAck(dAtA, i, uint64(len(m.ForwardRelayerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AppAcknowledgement) > 0 {
		i -= len(m.AppAcknowledgement)
		copy(dAtA[i:], m.AppAcknowledgement)
		i = encodeVarintAck(dAtA, i, uint64(len(m.AppAcknowledgement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAck(dAtA []byte, offset int, v uint64) int {
	offset -= sovAck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncentivizedAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppAcknowledgement)
	if l > 0 {
		n += 1 + l + sovAck(uint64(l))
	}
	l = len(m.ForwardRelayerAddress)
	if l > 0 {
		n += 1 + l + sovAck(uint64(l))
	}
	if m.UnderlyingAppSuccess {
		n += 2
	}
	return n
}

func sovAck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAck(x uint64) (n int) {
	return sovAck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncentivizedAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAck
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
			return fmt.Errorf("proto: IncentivizedAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentivizedAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAcknowledgement", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
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
				return ErrInvalidLengthAck
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAcknowledgement = append(m.AppAcknowledgement[:0], dAtA[iNdEx:postIndex]...)
			if m.AppAcknowledgement == nil {
				m.AppAcknowledgement = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardRelayerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
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
				return ErrInvalidLengthAck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForwardRelayerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderlyingAppSuccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAck
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
			m.UnderlyingAppSuccess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAck
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
func skipAck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAck
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
					return 0, ErrIntOverflowAck
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
					return 0, ErrIntOverflowAck
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
				return 0, ErrInvalidLengthAck
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAck
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAck
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAck        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAck          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAck = fmt.Errorf("proto: unexpected end of group")
)
