// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/compliance_history_item.proto

package types

import (
	fmt "fmt"
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

type ComplianceHistoryItem struct {
	SoftwareVersionCertificationStatus uint32 `protobuf:"varint,1,opt,name=softwareVersionCertificationStatus,proto3" json:"softwareVersionCertificationStatus,omitempty"`
	Date                               string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Reason                             string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	CDVersionNumber                    uint32 `protobuf:"varint,4,opt,name=cDVersionNumber,proto3" json:"cDVersionNumber,omitempty"`
}

func (m *ComplianceHistoryItem) Reset()         { *m = ComplianceHistoryItem{} }
func (m *ComplianceHistoryItem) String() string { return proto.CompactTextString(m) }
func (*ComplianceHistoryItem) ProtoMessage()    {}
func (*ComplianceHistoryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f19be5c9c9ea390, []int{0}
}
func (m *ComplianceHistoryItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceHistoryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceHistoryItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceHistoryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceHistoryItem.Merge(m, src)
}
func (m *ComplianceHistoryItem) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceHistoryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceHistoryItem.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceHistoryItem proto.InternalMessageInfo

func (m *ComplianceHistoryItem) GetSoftwareVersionCertificationStatus() uint32 {
	if m != nil {
		return m.SoftwareVersionCertificationStatus
	}
	return 0
}

func (m *ComplianceHistoryItem) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ComplianceHistoryItem) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ComplianceHistoryItem) GetCDVersionNumber() uint32 {
	if m != nil {
		return m.CDVersionNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*ComplianceHistoryItem)(nil), "zigbeealliance.distributedcomplianceledger.compliance.ComplianceHistoryItem")
}

func init() {
	proto.RegisterFile("compliance/compliance_history_item.proto", fileDescriptor_5f19be5c9c9ea390)
}

var fileDescriptor_5f19be5c9c9ea390 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x73, 0x5a, 0x0a, 0x1e, 0x88, 0x70, 0xa0, 0x64, 0x3a, 0x4a, 0xa7, 0x2c, 0x49, 0x06,
	0xf1, 0x0b, 0x58, 0x07, 0x45, 0xe8, 0x50, 0xc1, 0xc1, 0xa5, 0xdc, 0x25, 0xff, 0xa6, 0x7f, 0x48,
	0x72, 0xe1, 0xee, 0x1f, 0xb4, 0x7e, 0x0a, 0xbf, 0x90, 0xbb, 0x63, 0x47, 0x47, 0x49, 0xbe, 0x88,
	0x98, 0x54, 0x53, 0x9c, 0xba, 0xbd, 0xf7, 0x78, 0xf7, 0x8e, 0xff, 0x8f, 0x07, 0x89, 0x29, 0xaa,
	0x1c, 0x55, 0x99, 0x40, 0x3c, 0xc8, 0xe5, 0x1a, 0x1d, 0x19, 0xbb, 0x59, 0x22, 0x41, 0x11, 0x55,
	0xd6, 0x90, 0x11, 0x57, 0xaf, 0x98, 0x69, 0x00, 0x95, 0xf7, 0x95, 0x28, 0x45, 0x47, 0x16, 0x75,
	0x4d, 0x90, 0x0e, 0x0f, 0x73, 0x48, 0x33, 0xb0, 0xd1, 0x10, 0x4c, 0xdf, 0x19, 0x3f, 0x9f, 0xfd,
	0xd9, 0xdb, 0x7e, 0xf7, 0x8e, 0xa0, 0x10, 0x73, 0x3e, 0x75, 0x66, 0x45, 0xcf, 0xca, 0xc2, 0x23,
	0x58, 0x87, 0xa6, 0x9c, 0x81, 0x25, 0x5c, 0x61, 0xa2, 0x08, 0x4d, 0xf9, 0x40, 0x8a, 0x6a, 0xe7,
	0xb3, 0x09, 0x0b, 0x4e, 0x17, 0x07, 0x34, 0x85, 0xe0, 0xa3, 0x54, 0x11, 0xf8, 0x47, 0x13, 0x16,
	0x9c, 0x2c, 0x3a, 0x2d, 0x2e, 0xf8, 0xd8, 0x82, 0x72, 0xa6, 0xf4, 0x8f, 0xbb, 0x74, 0xe7, 0x44,
	0xc0, 0xcf, 0x92, 0x9b, 0xdd, 0xd6, 0xbc, 0x2e, 0x34, 0x58, 0x7f, 0xd4, 0x7d, 0xf4, 0x3f, 0xbe,
	0x86, 0x8f, 0x46, 0xb2, 0x6d, 0x23, 0xd9, 0x57, 0x23, 0xd9, 0x5b, 0x2b, 0xbd, 0x6d, 0x2b, 0xbd,
	0xcf, 0x56, 0x7a, 0x4f, 0xf7, 0x19, 0xd2, 0xba, 0xd6, 0x3f, 0xc7, 0xc6, 0x3d, 0x9b, 0xf0, 0x17,
	0x4e, 0xbc, 0x07, 0x27, 0x1c, 0x60, 0x84, 0x3d, 0x9e, 0xf8, 0x65, 0x0f, 0x75, 0x4c, 0x9b, 0x0a,
	0x9c, 0x1e, 0x77, 0x90, 0x2f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xec, 0xe6, 0xca, 0xb9, 0x90,
	0x01, 0x00, 0x00,
}

func (m *ComplianceHistoryItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceHistoryItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceHistoryItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CDVersionNumber != 0 {
		i = encodeVarintComplianceHistoryItem(dAtA, i, uint64(m.CDVersionNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintComplianceHistoryItem(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintComplianceHistoryItem(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x12
	}
	if m.SoftwareVersionCertificationStatus != 0 {
		i = encodeVarintComplianceHistoryItem(dAtA, i, uint64(m.SoftwareVersionCertificationStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceHistoryItem(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceHistoryItem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceHistoryItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SoftwareVersionCertificationStatus != 0 {
		n += 1 + sovComplianceHistoryItem(uint64(m.SoftwareVersionCertificationStatus))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovComplianceHistoryItem(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovComplianceHistoryItem(uint64(l))
	}
	if m.CDVersionNumber != 0 {
		n += 1 + sovComplianceHistoryItem(uint64(m.CDVersionNumber))
	}
	return n
}

func sovComplianceHistoryItem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceHistoryItem(x uint64) (n int) {
	return sovComplianceHistoryItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceHistoryItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceHistoryItem
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
			return fmt.Errorf("proto: ComplianceHistoryItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceHistoryItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionCertificationStatus", wireType)
			}
			m.SoftwareVersionCertificationStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceHistoryItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersionCertificationStatus |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceHistoryItem
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
				return ErrInvalidLengthComplianceHistoryItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceHistoryItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceHistoryItem
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
				return ErrInvalidLengthComplianceHistoryItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceHistoryItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CDVersionNumber", wireType)
			}
			m.CDVersionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceHistoryItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CDVersionNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceHistoryItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceHistoryItem
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
func skipComplianceHistoryItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceHistoryItem
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
					return 0, ErrIntOverflowComplianceHistoryItem
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
					return 0, ErrIntOverflowComplianceHistoryItem
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
				return 0, ErrInvalidLengthComplianceHistoryItem
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceHistoryItem
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceHistoryItem
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceHistoryItem        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceHistoryItem          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceHistoryItem = fmt.Errorf("proto: unexpected end of group")
)
