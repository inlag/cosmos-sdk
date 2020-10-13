// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/convo/convo.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
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

// MsgConvo defines a Msg for a sender on one chain to send a message to a receiver
// on another chain connected by an IBC channel
type MsgConvo struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the sender address
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// the message sender sends to receiver
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,7,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,8,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgConvo) Reset()         { *m = MsgConvo{} }
func (m *MsgConvo) String() string { return proto.CompactTextString(m) }
func (*MsgConvo) ProtoMessage()    {}
func (*MsgConvo) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e93517810dd02f, []int{0}
}
func (m *MsgConvo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvo.Merge(m, src)
}
func (m *MsgConvo) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvo)(nil), "ibc.applications.convo.v1.MsgConvo")
}

func init() {
	proto.RegisterFile("ibc/applications/convo/convo.proto", fileDescriptor_52e93517810dd02f)
}

var fileDescriptor_52e93517810dd02f = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xbf, 0x8e, 0xd3, 0x30,
	0x18, 0x4f, 0xe0, 0xe8, 0x05, 0x9f, 0xee, 0x04, 0x16, 0x9c, 0x7c, 0x11, 0x24, 0x55, 0xa6, 0x2e,
	0xc4, 0x2a, 0x0c, 0x48, 0xc7, 0x82, 0xd2, 0x05, 0x06, 0x24, 0x14, 0xc1, 0xc2, 0x52, 0x12, 0xd7,
	0x4a, 0x2c, 0x92, 0x38, 0xb2, 0xdd, 0x88, 0xbe, 0x01, 0x62, 0xe2, 0x11, 0xfa, 0x38, 0x1d, 0x3b,
	0x32, 0x45, 0xa8, 0x5d, 0x98, 0xfb, 0x04, 0x28, 0xb1, 0x5b, 0xda, 0xe1, 0x16, 0x7f, 0xdf, 0xef,
	0xdf, 0x27, 0xeb, 0xb3, 0x41, 0xc0, 0x52, 0x82, 0x93, 0xba, 0x2e, 0x18, 0x49, 0x14, 0xe3, 0x95,
	0xc4, 0x84, 0x57, 0x0d, 0xd7, 0x67, 0x58, 0x0b, 0xae, 0x38, 0xbc, 0x61, 0x29, 0x09, 0x8f, 0x3d,
	0xa1, 0x56, 0x9b, 0xb1, 0xfb, 0x24, 0xe3, 0x19, 0xef, 0x5d, 0xb8, 0xeb, 0x74, 0xc0, 0xf5, 0xbb,
	0xa1, 0x84, 0x0b, 0x8a, 0x49, 0xc1, 0x68, 0xa5, 0x70, 0x33, 0x36, 0x9d, 0x36, 0x04, 0x3f, 0xef,
	0x03, 0xe7, 0x83, 0xcc, 0x26, 0xdd, 0x18, 0xf8, 0x1a, 0x5c, 0x48, 0x3e, 0x17, 0x84, 0x4e, 0x6b,
	0x2e, 0x14, 0xb2, 0x87, 0xf6, 0xe8, 0x61, 0x74, 0xbd, 0x6b, 0x7d, 0xb8, 0x48, 0xca, 0xe2, 0x36,
	0x38, 0x12, 0x83, 0x18, 0x68, 0xf4, 0x91, 0x0b, 0x05, 0xdf, 0x82, 0x2b, 0xa3, 0x91, 0x3c, 0xa9,
	0x2a, 0x5a, 0xa0, 0x7b, 0x7d, 0xf6, 0x66, 0xd7, 0xfa, 0x4f, 0x4f, 0xb2, 0x46, 0x0f, 0xe2, 0x4b,
	0x4d, 0x4c, 0x34, 0x86, 0xd7, 0x60, 0x20, 0x69, 0x35, 0xa3, 0x02, 0x9d, 0x75, 0xc9, 0xd8, 0x20,
	0xe8, 0x02, 0x47, 0x50, 0x42, 0x59, 0x43, 0x05, 0x7a, 0xd0, 0x2b, 0x07, 0x0c, 0x11, 0x38, 0x2f,
	0xa9, 0x94, 0x49, 0x46, 0xd1, 0xa0, 0x97, 0xf6, 0x10, 0x7e, 0x05, 0x57, 0x8a, 0x95, 0x94, 0xcf,
	0xd5, 0x34, 0xa7, 0x2c, 0xcb, 0x15, 0x3a, 0x1f, 0xda, 0xa3, 0x8b, 0x97, 0x6e, 0xd8, 0x2d, 0xb0,
	0xdb, 0x47, 0x68, 0xb6, 0xd0, 0x8c, 0xc3, 0x77, 0xbd, 0x23, 0x7a, 0xbe, 0x6a, 0x7d, 0xeb, 0xff,
	0x7d, 0x4f, 0xf3, 0x41, 0x7c, 0x69, 0x08, 0xed, 0x86, 0xef, 0xc1, 0xe3, 0xbd, 0xa3, 0xab, 0x52,
	0x25, 0x65, 0x8d, 0x9c, 0xa1, 0x3d, 0x3a, 0x8b, 0x9e, 0xed, 0x5a, 0x1f, 0x9d, 0x0e, 0x39, 0x58,
	0x82, 0xf8, 0x91, 0xe1, 0x3e, 0xed, 0xa9, 0x5b, 0xe7, 0xc7, 0xd2, 0xb7, 0xfe, 0x2e, 0x7d, 0x2b,
	0xfa, 0xbc, 0xda, 0x78, 0xf6, 0x7a, 0xe3, 0xd9, 0x7f, 0x36, 0x9e, 0xfd, 0x6b, 0xeb, 0x59, 0xeb,
	0xad, 0x67, 0xfd, 0xde, 0x7a, 0xd6, 0x97, 0x37, 0x19, 0x53, 0xf9, 0x3c, 0x0d, 0x09, 0x2f, 0x31,
	0xe1, 0xb2, 0xe4, 0xd2, 0x94, 0x17, 0x72, 0xf6, 0x0d, 0x7f, 0xc7, 0x77, 0xfc, 0x1d, 0xb5, 0xa8,
	0xa9, 0x4c, 0x07, 0xfd, 0x53, 0xbf, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xbf, 0x39, 0x58,
	0x62, 0x02, 0x00, 0x00,
}

func (m *MsgConvo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintConvo(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConvo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintConvo(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintConvo(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintConvo(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintConvo(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintConvo(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConvo(dAtA []byte, offset int, v uint64) int {
	offset -= sovConvo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovConvo(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovConvo(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovConvo(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovConvo(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovConvo(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovConvo(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovConvo(uint64(m.TimeoutTimestamp))
	}
	return n
}

func sovConvo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConvo(x uint64) (n int) {
	return sovConvo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConvo
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
			return fmt.Errorf("proto: MsgConvo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
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
				return ErrInvalidLengthConvo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConvo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConvo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConvo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConvo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConvo
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
func skipConvo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConvo
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
					return 0, ErrIntOverflowConvo
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
					return 0, ErrIntOverflowConvo
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
				return 0, ErrInvalidLengthConvo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConvo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConvo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConvo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConvo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConvo = fmt.Errorf("proto: unexpected end of group")
)