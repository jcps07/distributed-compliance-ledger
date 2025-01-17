// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/proposed_certificate.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ProposedCertificate struct {
	Subject       string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId  string   `protobuf:"bytes,2,opt,name=subjectKeyId,proto3" json:"subjectKeyId,omitempty"`
	PemCert       string   `protobuf:"bytes,3,opt,name=pemCert,proto3" json:"pemCert,omitempty"`
	SerialNumber  string   `protobuf:"bytes,4,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	Owner         string   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Approvals     []*Grant `protobuf:"bytes,6,rep,name=approvals,proto3" json:"approvals,omitempty"`
	SubjectAsText string   `protobuf:"bytes,7,opt,name=subjectAsText,proto3" json:"subjectAsText,omitempty"`
	Rejects       []*Grant `protobuf:"bytes,8,rep,name=rejects,proto3" json:"rejects,omitempty"`
	Vid           int32    `protobuf:"varint,9,opt,name=vid,proto3" json:"vid,omitempty" validate:"gte=1,lte=65535"`
}

func (m *ProposedCertificate) Reset()         { *m = ProposedCertificate{} }
func (m *ProposedCertificate) String() string { return proto.CompactTextString(m) }
func (*ProposedCertificate) ProtoMessage()    {}
func (*ProposedCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_589d443e04a789ec, []int{0}
}
func (m *ProposedCertificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedCertificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedCertificate.Merge(m, src)
}
func (m *ProposedCertificate) XXX_Size() int {
	return m.Size()
}
func (m *ProposedCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedCertificate proto.InternalMessageInfo

func (m *ProposedCertificate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ProposedCertificate) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *ProposedCertificate) GetPemCert() string {
	if m != nil {
		return m.PemCert
	}
	return ""
}

func (m *ProposedCertificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *ProposedCertificate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProposedCertificate) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *ProposedCertificate) GetSubjectAsText() string {
	if m != nil {
		return m.SubjectAsText
	}
	return ""
}

func (m *ProposedCertificate) GetRejects() []*Grant {
	if m != nil {
		return m.Rejects
	}
	return nil
}

func (m *ProposedCertificate) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func init() {
	proto.RegisterType((*ProposedCertificate)(nil), "zigbeealliance.distributedcomplianceledger.pki.ProposedCertificate")
}

func init() { proto.RegisterFile("pki/proposed_certificate.proto", fileDescriptor_589d443e04a789ec) }

var fileDescriptor_589d443e04a789ec = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x6e, 0x13, 0x31,
	0x1c, 0xc6, 0x73, 0x84, 0x34, 0xc4, 0x80, 0x90, 0x4c, 0x07, 0x93, 0xc1, 0x44, 0x11, 0x43, 0x06,
	0x72, 0x07, 0x54, 0xc7, 0x80, 0xd4, 0xa1, 0x01, 0x09, 0x21, 0x24, 0x40, 0x57, 0x26, 0x06, 0x2a,
	0xdf, 0xf9, 0x8f, 0x31, 0xbd, 0x3b, 0x5b, 0xb6, 0x2f, 0xb4, 0x3c, 0x05, 0x0f, 0xc3, 0xce, 0xca,
	0x58, 0x31, 0x31, 0x21, 0x94, 0xbc, 0x01, 0x4f, 0x80, 0x7c, 0xbe, 0x28, 0xed, 0x88, 0xba, 0xf9,
	0xfb, 0xec, 0xef, 0xf3, 0xcf, 0xd6, 0x1f, 0x51, 0x7d, 0x2c, 0x13, 0x6d, 0x94, 0x56, 0x16, 0xf8,
	0x51, 0x01, 0xc6, 0xc9, 0x0f, 0xb2, 0x60, 0x0e, 0x62, 0x6d, 0x94, 0x53, 0x38, 0xfe, 0x22, 0x45,
	0x0e, 0xc0, 0xca, 0x52, 0xb2, 0xba, 0x80, 0x98, 0x4b, 0xeb, 0x8c, 0xcc, 0x1b, 0x07, 0xbc, 0x50,
	0x95, 0x0e, 0x6e, 0x09, 0x5c, 0x80, 0x89, 0xf5, 0xb1, 0x1c, 0xdf, 0x29, 0x94, 0xad, 0x94, 0x3d,
	0x6a, 0xd3, 0x49, 0x10, 0xa1, 0x6a, 0xbc, 0x2b, 0x94, 0x50, 0xc1, 0xf7, 0xab, 0xce, 0xbd, 0xe5,
	0x01, 0x84, 0x61, 0xb5, 0x0b, 0xc6, 0xf4, 0x7b, 0x1f, 0xdd, 0x7e, 0xd3, 0x01, 0x3d, 0xdd, 0xf2,
	0x60, 0x82, 0x86, 0xb6, 0xc9, 0x3f, 0x41, 0xe1, 0x48, 0x34, 0x89, 0x66, 0xa3, 0x6c, 0x23, 0xf1,
	0x14, 0xdd, 0xe8, 0x96, 0x2f, 0xe1, 0xf4, 0x05, 0x27, 0x57, 0xda, 0xed, 0x0b, 0x9e, 0x4f, 0x6b,
	0xa8, 0x7c, 0x1f, 0xe9, 0x87, 0x74, 0x27, 0xdb, 0x34, 0x18, 0xc9, 0xca, 0x57, 0x4d, 0x95, 0x83,
	0x21, 0x57, 0xbb, 0xf4, 0x39, 0x0f, 0xc7, 0x68, 0xa0, 0x3e, 0xd7, 0x60, 0xc8, 0xc0, 0x6f, 0x2e,
	0xc8, 0xcf, 0x6f, 0xf3, 0xdd, 0xee, 0x6d, 0x07, 0x9c, 0x1b, 0xb0, 0xf6, 0xd0, 0x19, 0x59, 0x8b,
	0x2c, 0x1c, 0xc3, 0x87, 0x68, 0xc4, 0xb4, 0x36, 0x6a, 0xc9, 0x4a, 0x4b, 0x76, 0x26, 0xfd, 0xd9,
	0xf5, 0x47, 0xe9, 0x7f, 0xfe, 0x64, 0xfc, 0xdc, 0xff, 0x49, 0xb6, 0xed, 0xc1, 0xf7, 0xd0, 0xcd,
	0xee, 0x49, 0x07, 0xf6, 0x2d, 0x9c, 0x38, 0x32, 0x6c, 0x49, 0x2f, 0x9a, 0xf8, 0x35, 0x1a, 0x1a,
	0xf0, 0xda, 0x92, 0x6b, 0x97, 0xb9, 0x78, 0xd3, 0x82, 0x1f, 0xa0, 0xfe, 0x52, 0x72, 0x32, 0x9a,
	0x44, 0xb3, 0xc1, 0x82, 0xfe, 0xfd, 0x7d, 0x77, 0xbc, 0x64, 0xa5, 0xe4, 0xcc, 0xc1, 0x93, 0xa9,
	0x70, 0xb0, 0xff, 0xf0, 0x7e, 0xe9, 0x60, 0xff, 0x71, 0x9a, 0xee, 0xa5, 0xd3, 0xcc, 0x1f, 0x5d,
	0xbc, 0xff, 0xb1, 0xa2, 0xd1, 0xd9, 0x8a, 0x46, 0x7f, 0x56, 0x34, 0xfa, 0xba, 0xa6, 0xbd, 0xb3,
	0x35, 0xed, 0xfd, 0x5a, 0xd3, 0xde, 0xbb, 0x67, 0x42, 0xba, 0x8f, 0x4d, 0x1e, 0x17, 0xaa, 0x4a,
	0x02, 0xd5, 0x7c, 0x83, 0x95, 0x9c, 0xc3, 0x9a, 0x6f, 0xb9, 0xe6, 0x01, 0x2c, 0x39, 0x49, 0xfc,
	0x9c, 0xb8, 0x53, 0x0d, 0x36, 0xdf, 0x69, 0x07, 0x65, 0xef, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0x05, 0x51, 0xbe, 0xbc, 0x02, 0x00, 0x00,
}

func (m *ProposedCertificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedCertificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedCertificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vid != 0 {
		i = encodeVarintProposedCertificate(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Rejects) > 0 {
		for iNdEx := len(m.Rejects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.SubjectAsText) > 0 {
		i -= len(m.SubjectAsText)
		copy(dAtA[i:], m.SubjectAsText)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.SubjectAsText)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposedCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedCertificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovProposedCertificate(uint64(l))
		}
	}
	l = len(m.SubjectAsText)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	if len(m.Rejects) > 0 {
		for _, e := range m.Rejects {
			l = e.Size()
			n += 1 + l + sovProposedCertificate(uint64(l))
		}
	}
	if m.Vid != 0 {
		n += 1 + sovProposedCertificate(uint64(m.Vid))
	}
	return n
}

func sovProposedCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedCertificate(x uint64) (n int) {
	return sovProposedCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedCertificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedCertificate
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
			return fmt.Errorf("proto: ProposedCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Grant{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAsText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectAsText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rejects = append(m.Rejects, &Grant{})
			if err := m.Rejects[len(m.Rejects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposedCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedCertificate
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
func skipProposedCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedCertificate
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
					return 0, ErrIntOverflowProposedCertificate
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
					return 0, ErrIntOverflowProposedCertificate
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
				return 0, ErrInvalidLengthProposedCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedCertificate = fmt.Errorf("proto: unexpected end of group")
)
