// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/compliance_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ComplianceInfo struct {
	Vid                                int32                    `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Pid                                int32                    `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	SoftwareVersion                    uint32                   `protobuf:"varint,3,opt,name=softwareVersion,proto3" json:"softwareVersion,omitempty"`
	CertificationType                  string                   `protobuf:"bytes,4,opt,name=certificationType,proto3" json:"certificationType,omitempty"`
	SoftwareVersionString              string                   `protobuf:"bytes,5,opt,name=softwareVersionString,proto3" json:"softwareVersionString,omitempty"`
	CDVersionNumber                    uint32                   `protobuf:"varint,6,opt,name=cDVersionNumber,proto3" json:"cDVersionNumber,omitempty"`
	SoftwareVersionCertificationStatus uint32                   `protobuf:"varint,7,opt,name=softwareVersionCertificationStatus,proto3" json:"softwareVersionCertificationStatus,omitempty"`
	Date                               string                   `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Reason                             string                   `protobuf:"bytes,9,opt,name=reason,proto3" json:"reason,omitempty"`
	Owner                              string                   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	History                            []*ComplianceHistoryItem `protobuf:"bytes,11,rep,name=history,proto3" json:"history,omitempty"`
	ProgramTypeVersion                 string                   `protobuf:"bytes,12,opt,name=programTypeVersion,proto3" json:"programTypeVersion,omitempty"`
	CertificationID                    string                   `protobuf:"bytes,13,opt,name=certificationID,proto3" json:"certificationID,omitempty"`
	FamilyID                           string                   `protobuf:"bytes,14,opt,name=familyID,proto3" json:"familyID,omitempty"`
	SupportedClusters                  string                   `protobuf:"bytes,15,opt,name=supportedClusters,proto3" json:"supportedClusters,omitempty"`
	CompliancePlatformUsed             string                   `protobuf:"bytes,16,opt,name=compliancePlatformUsed,proto3" json:"compliancePlatformUsed,omitempty"`
	CompliancePlatformVersion          string                   `protobuf:"bytes,17,opt,name=compliancePlatformVersion,proto3" json:"compliancePlatformVersion,omitempty"`
}

func (m *ComplianceInfo) Reset()         { *m = ComplianceInfo{} }
func (m *ComplianceInfo) String() string { return proto.CompactTextString(m) }
func (*ComplianceInfo) ProtoMessage()    {}
func (*ComplianceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96709118384d66e8, []int{0}
}
func (m *ComplianceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceInfo.Merge(m, src)
}
func (m *ComplianceInfo) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceInfo proto.InternalMessageInfo

func (m *ComplianceInfo) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *ComplianceInfo) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ComplianceInfo) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *ComplianceInfo) GetCertificationType() string {
	if m != nil {
		return m.CertificationType
	}
	return ""
}

func (m *ComplianceInfo) GetSoftwareVersionString() string {
	if m != nil {
		return m.SoftwareVersionString
	}
	return ""
}

func (m *ComplianceInfo) GetCDVersionNumber() uint32 {
	if m != nil {
		return m.CDVersionNumber
	}
	return 0
}

func (m *ComplianceInfo) GetSoftwareVersionCertificationStatus() uint32 {
	if m != nil {
		return m.SoftwareVersionCertificationStatus
	}
	return 0
}

func (m *ComplianceInfo) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ComplianceInfo) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ComplianceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ComplianceInfo) GetHistory() []*ComplianceHistoryItem {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *ComplianceInfo) GetProgramTypeVersion() string {
	if m != nil {
		return m.ProgramTypeVersion
	}
	return ""
}

func (m *ComplianceInfo) GetCertificationID() string {
	if m != nil {
		return m.CertificationID
	}
	return ""
}

func (m *ComplianceInfo) GetFamilyID() string {
	if m != nil {
		return m.FamilyID
	}
	return ""
}

func (m *ComplianceInfo) GetSupportedClusters() string {
	if m != nil {
		return m.SupportedClusters
	}
	return ""
}

func (m *ComplianceInfo) GetCompliancePlatformUsed() string {
	if m != nil {
		return m.CompliancePlatformUsed
	}
	return ""
}

func (m *ComplianceInfo) GetCompliancePlatformVersion() string {
	if m != nil {
		return m.CompliancePlatformVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*ComplianceInfo)(nil), "zigbeealliance.distributedcomplianceledger.compliance.ComplianceInfo")
}

func init() { proto.RegisterFile("compliance/compliance_info.proto", fileDescriptor_96709118384d66e8) }

var fileDescriptor_96709118384d66e8 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x49, 0x93, 0xb6, 0x53, 0xfa, 0x1a, 0x41, 0x35, 0xcd, 0xc2, 0x8a, 0xba, 0xf2, 0x82,
	0xd8, 0x12, 0xaf, 0x15, 0x1b, 0x9a, 0x2c, 0x88, 0x40, 0x15, 0x4a, 0x81, 0x05, 0x9b, 0x68, 0xe2,
	0xb9, 0x76, 0x47, 0xb2, 0x3d, 0xd6, 0xcc, 0x98, 0x12, 0xbe, 0x82, 0x35, 0xdf, 0xc1, 0x47, 0xb0,
	0xac, 0x58, 0xb1, 0x44, 0xc9, 0x8f, 0x20, 0xcf, 0x24, 0x75, 0xe4, 0xa4, 0x12, 0xea, 0x6e, 0xee,
	0x39, 0xf7, 0x1c, 0xdf, 0x99, 0x73, 0x8d, 0xba, 0xa1, 0x48, 0xf3, 0x84, 0xd3, 0x2c, 0x84, 0xa0,
	0x3a, 0x8e, 0x79, 0x16, 0x09, 0x3f, 0x97, 0x42, 0x0b, 0xfc, 0xe2, 0x1b, 0x8f, 0x27, 0x00, 0x34,
	0xb1, 0x94, 0xcf, 0xb8, 0xd2, 0x92, 0x4f, 0x0a, 0x0d, 0xac, 0x12, 0x24, 0xc0, 0x62, 0x90, 0x7e,
	0x05, 0x74, 0xbc, 0xcd, 0xc6, 0x57, 0x5c, 0x69, 0x21, 0xa7, 0x63, 0xae, 0x21, 0xb5, 0x1f, 0xe8,
	0x9c, 0x86, 0x42, 0xa5, 0x42, 0x8d, 0x4d, 0x15, 0xd8, 0xc2, 0x52, 0x67, 0x3f, 0xda, 0xe8, 0xa0,
	0x7f, 0x2b, 0x1e, 0x66, 0x91, 0xc0, 0x47, 0xa8, 0xf9, 0x85, 0x33, 0xe2, 0x74, 0x1d, 0xaf, 0x35,
	0x2a, 0x8f, 0x25, 0x92, 0x73, 0x46, 0x1e, 0x58, 0x24, 0xe7, 0x0c, 0x7b, 0xe8, 0x50, 0x89, 0x48,
	0x5f, 0x53, 0x09, 0x9f, 0x40, 0x2a, 0x2e, 0x32, 0xd2, 0xec, 0x3a, 0xde, 0xfe, 0xa8, 0x0e, 0xe3,
	0x27, 0xe8, 0x38, 0x04, 0xa9, 0x79, 0xc4, 0x43, 0xaa, 0xb9, 0xc8, 0x3e, 0x4c, 0x73, 0x20, 0x5b,
	0x5d, 0xc7, 0xdb, 0x1d, 0xad, 0x13, 0xf8, 0x39, 0x7a, 0x5c, 0x33, 0xb8, 0xd4, 0x92, 0x67, 0x31,
	0x69, 0x19, 0xc5, 0x66, 0xb2, 0x9c, 0x26, 0x1c, 0x2c, 0xa0, 0x8b, 0x22, 0x9d, 0x80, 0x24, 0x6d,
	0x3b, 0x4d, 0x0d, 0xc6, 0x17, 0xe8, 0xac, 0x66, 0xd1, 0x5f, 0x9d, 0xe1, 0x52, 0x53, 0x5d, 0x28,
	0xb2, 0x6d, 0xc4, 0xff, 0xd1, 0x89, 0x31, 0xda, 0x62, 0x54, 0x03, 0xd9, 0x31, 0xe3, 0x99, 0x33,
	0x3e, 0x41, 0x6d, 0x09, 0x54, 0x89, 0x8c, 0xec, 0x1a, 0x74, 0x51, 0x61, 0x1f, 0xb5, 0xc4, 0x75,
	0x06, 0x92, 0xa0, 0x12, 0x3e, 0x27, 0xbf, 0x7f, 0xf6, 0x1e, 0x2d, 0xb2, 0x78, 0xcd, 0x98, 0x04,
	0xa5, 0xec, 0x75, 0x46, 0xb6, 0x0d, 0x47, 0x68, 0x7b, 0x91, 0x25, 0xd9, 0xeb, 0x36, 0xbd, 0xbd,
	0xa7, 0xef, 0xfc, 0x7b, 0x2d, 0x8a, 0x5f, 0xe5, 0xfb, 0xc6, 0xfa, 0x0d, 0x35, 0xa4, 0xa3, 0xa5,
	0x39, 0xf6, 0x11, 0xce, 0xa5, 0x88, 0x25, 0x4d, 0xcb, 0x08, 0x96, 0x71, 0x3e, 0x34, 0xb3, 0x6f,
	0x60, 0xcc, 0x6b, 0xaf, 0x3e, 0xc5, 0x70, 0x40, 0xf6, 0x4d, 0x73, 0x1d, 0xc6, 0x1d, 0xb4, 0x13,
	0xd1, 0x94, 0x27, 0xd3, 0xe1, 0x80, 0x1c, 0x98, 0x96, 0xdb, 0xba, 0xdc, 0x0b, 0x55, 0xe4, 0xb9,
	0x90, 0x1a, 0x58, 0x3f, 0x29, 0x94, 0x06, 0xa9, 0xc8, 0xa1, 0xdd, 0x8b, 0x35, 0x02, 0xbf, 0x44,
	0x27, 0xd5, 0x85, 0xde, 0x27, 0x54, 0x47, 0x42, 0xa6, 0x1f, 0x15, 0x30, 0x72, 0x64, 0x24, 0x77,
	0xb0, 0xf8, 0x15, 0x3a, 0x5d, 0x67, 0x96, 0x57, 0x3c, 0x36, 0xd2, 0xbb, 0x1b, 0xce, 0xe1, 0xd7,
	0xcc, 0x75, 0x6e, 0x66, 0xae, 0xf3, 0x77, 0xe6, 0x3a, 0xdf, 0xe7, 0x6e, 0xe3, 0x66, 0xee, 0x36,
	0xfe, 0xcc, 0xdd, 0xc6, 0xe7, 0xb7, 0x31, 0xd7, 0x57, 0xc5, 0xa4, 0x7c, 0xe5, 0xc0, 0x86, 0xd2,
	0x5b, 0xa6, 0x12, 0xac, 0xa4, 0xd2, 0xab, 0xbc, 0x7b, 0x36, 0x97, 0xe0, 0xeb, 0xca, 0xbf, 0x1a,
	0xe8, 0x69, 0x0e, 0x6a, 0xd2, 0x36, 0xbf, 0xe2, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x35, 0x0e, 0xce, 0x2a, 0x04, 0x00, 0x00,
}

func (m *ComplianceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CompliancePlatformVersion) > 0 {
		i -= len(m.CompliancePlatformVersion)
		copy(dAtA[i:], m.CompliancePlatformVersion)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CompliancePlatformVersion)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.CompliancePlatformUsed) > 0 {
		i -= len(m.CompliancePlatformUsed)
		copy(dAtA[i:], m.CompliancePlatformUsed)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CompliancePlatformUsed)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.SupportedClusters) > 0 {
		i -= len(m.SupportedClusters)
		copy(dAtA[i:], m.SupportedClusters)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.SupportedClusters)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.FamilyID) > 0 {
		i -= len(m.FamilyID)
		copy(dAtA[i:], m.FamilyID)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.FamilyID)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.CertificationID) > 0 {
		i -= len(m.CertificationID)
		copy(dAtA[i:], m.CertificationID)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CertificationID)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ProgramTypeVersion) > 0 {
		i -= len(m.ProgramTypeVersion)
		copy(dAtA[i:], m.ProgramTypeVersion)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.ProgramTypeVersion)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.History) > 0 {
		for iNdEx := len(m.History) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.History[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintComplianceInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x42
	}
	if m.SoftwareVersionCertificationStatus != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.SoftwareVersionCertificationStatus))
		i--
		dAtA[i] = 0x38
	}
	if m.CDVersionNumber != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.CDVersionNumber))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SoftwareVersionString) > 0 {
		i -= len(m.SoftwareVersionString)
		copy(dAtA[i:], m.SoftwareVersionString)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.SoftwareVersionString)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CertificationType) > 0 {
		i -= len(m.CertificationType)
		copy(dAtA[i:], m.CertificationType)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CertificationType)))
		i--
		dAtA[i] = 0x22
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.Pid != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x10
	}
	if m.Vid != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovComplianceInfo(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovComplianceInfo(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovComplianceInfo(uint64(m.SoftwareVersion))
	}
	l = len(m.CertificationType)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.SoftwareVersionString)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	if m.CDVersionNumber != 0 {
		n += 1 + sovComplianceInfo(uint64(m.CDVersionNumber))
	}
	if m.SoftwareVersionCertificationStatus != 0 {
		n += 1 + sovComplianceInfo(uint64(m.SoftwareVersionCertificationStatus))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	if len(m.History) > 0 {
		for _, e := range m.History {
			l = e.Size()
			n += 1 + l + sovComplianceInfo(uint64(l))
		}
	}
	l = len(m.ProgramTypeVersion)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CertificationID)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.FamilyID)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.SupportedClusters)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CompliancePlatformUsed)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CompliancePlatformVersion)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	return n
}

func sovComplianceInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceInfo(x uint64) (n int) {
	return sovComplianceInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceInfo
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
			return fmt.Errorf("proto: ComplianceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SoftwareVersionString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CDVersionNumber", wireType)
			}
			m.CDVersionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionCertificationStatus", wireType)
			}
			m.SoftwareVersionCertificationStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.History = append(m.History, &ComplianceHistoryItem{})
			if err := m.History[len(m.History)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramTypeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProgramTypeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FamilyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FamilyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedClusters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedClusters = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompliancePlatformUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompliancePlatformUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompliancePlatformVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompliancePlatformVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceInfo
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
func skipComplianceInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceInfo
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
					return 0, ErrIntOverflowComplianceInfo
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
					return 0, ErrIntOverflowComplianceInfo
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
				return 0, ErrInvalidLengthComplianceInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceInfo = fmt.Errorf("proto: unexpected end of group")
)
