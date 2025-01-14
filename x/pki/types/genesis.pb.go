// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/genesis.proto

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

// GenesisState defines the pki module's genesis state.
type GenesisState struct {
	ApprovedCertificatesList                                []ApprovedCertificates                                `protobuf:"bytes,1,rep,name=approvedCertificatesList,proto3" json:"approvedCertificatesList"`
	ProposedCertificateList                                 []ProposedCertificate                                 `protobuf:"bytes,2,rep,name=proposedCertificateList,proto3" json:"proposedCertificateList"`
	ChildCertificatesList                                   []ChildCertificates                                   `protobuf:"bytes,3,rep,name=childCertificatesList,proto3" json:"childCertificatesList"`
	ProposedCertificateRevocationList                       []ProposedCertificateRevocation                       `protobuf:"bytes,4,rep,name=proposedCertificateRevocationList,proto3" json:"proposedCertificateRevocationList"`
	RevokedCertificatesList                                 []RevokedCertificates                                 `protobuf:"bytes,5,rep,name=revokedCertificatesList,proto3" json:"revokedCertificatesList"`
	UniqueCertificateList                                   []UniqueCertificate                                   `protobuf:"bytes,6,rep,name=uniqueCertificateList,proto3" json:"uniqueCertificateList"`
	ApprovedRootCertificates                                *ApprovedRootCertificates                             `protobuf:"bytes,7,opt,name=approvedRootCertificates,proto3" json:"approvedRootCertificates,omitempty"`
	RevokedRootCertificates                                 *RevokedRootCertificates                              `protobuf:"bytes,8,opt,name=revokedRootCertificates,proto3" json:"revokedRootCertificates,omitempty"`
	ApprovedCertificatesBySubjectList                       []ApprovedCertificatesBySubject                       `protobuf:"bytes,9,rep,name=approvedCertificatesBySubjectList,proto3" json:"approvedCertificatesBySubjectList"`
	RejectedCertificateList                                 []RejectedCertificate                                 `protobuf:"bytes,10,rep,name=rejectedCertificateList,proto3" json:"rejectedCertificateList"`
	PkiRevocationDistributionPointList                      []PkiRevocationDistributionPoint                      `protobuf:"bytes,11,rep,name=PkiRevocationDistributionPointList,proto3" json:"PkiRevocationDistributionPointList"`
	PkiRevocationDistributionPointsByIssuerSubjectKeyIDList []PkiRevocationDistributionPointsByIssuerSubjectKeyID `protobuf:"bytes,12,rep,name=PkiRevocationDistributionPointsByIssuerSubjectKeyIDList,proto3" json:"PkiRevocationDistributionPointsByIssuerSubjectKeyIDList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9478608499b59120, []int{0}
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

func (m *GenesisState) GetApprovedCertificatesList() []ApprovedCertificates {
	if m != nil {
		return m.ApprovedCertificatesList
	}
	return nil
}

func (m *GenesisState) GetProposedCertificateList() []ProposedCertificate {
	if m != nil {
		return m.ProposedCertificateList
	}
	return nil
}

func (m *GenesisState) GetChildCertificatesList() []ChildCertificates {
	if m != nil {
		return m.ChildCertificatesList
	}
	return nil
}

func (m *GenesisState) GetProposedCertificateRevocationList() []ProposedCertificateRevocation {
	if m != nil {
		return m.ProposedCertificateRevocationList
	}
	return nil
}

func (m *GenesisState) GetRevokedCertificatesList() []RevokedCertificates {
	if m != nil {
		return m.RevokedCertificatesList
	}
	return nil
}

func (m *GenesisState) GetUniqueCertificateList() []UniqueCertificate {
	if m != nil {
		return m.UniqueCertificateList
	}
	return nil
}

func (m *GenesisState) GetApprovedRootCertificates() *ApprovedRootCertificates {
	if m != nil {
		return m.ApprovedRootCertificates
	}
	return nil
}

func (m *GenesisState) GetRevokedRootCertificates() *RevokedRootCertificates {
	if m != nil {
		return m.RevokedRootCertificates
	}
	return nil
}

func (m *GenesisState) GetApprovedCertificatesBySubjectList() []ApprovedCertificatesBySubject {
	if m != nil {
		return m.ApprovedCertificatesBySubjectList
	}
	return nil
}

func (m *GenesisState) GetRejectedCertificateList() []RejectedCertificate {
	if m != nil {
		return m.RejectedCertificateList
	}
	return nil
}

func (m *GenesisState) GetPkiRevocationDistributionPointList() []PkiRevocationDistributionPoint {
	if m != nil {
		return m.PkiRevocationDistributionPointList
	}
	return nil
}

func (m *GenesisState) GetPkiRevocationDistributionPointsByIssuerSubjectKeyIDList() []PkiRevocationDistributionPointsByIssuerSubjectKeyID {
	if m != nil {
		return m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zigbeealliance.distributedcomplianceledger.pki.GenesisState")
}

func init() { proto.RegisterFile("pki/genesis.proto", fileDescriptor_9478608499b59120) }

var fileDescriptor_9478608499b59120 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0x36, 0x06, 0x64, 0x13, 0x12, 0x11, 0x88, 0x6a, 0x42, 0xd9, 0x18, 0x1c, 0x06,
	0xa8, 0x89, 0x34, 0x0e, 0x9c, 0xfb, 0x07, 0x8d, 0x89, 0x3f, 0xaa, 0x3a, 0x71, 0xe1, 0x40, 0x94,
	0xa4, 0x26, 0x33, 0xe9, 0x6a, 0x13, 0x3b, 0x13, 0x41, 0xe2, 0x80, 0x40, 0x70, 0xe5, 0x2b, 0x20,
	0xf1, 0x61, 0x76, 0xdc, 0x91, 0x13, 0x42, 0xed, 0x17, 0x41, 0x7e, 0xe3, 0xac, 0x69, 0xe3, 0x74,
	0x6b, 0xb4, 0x5b, 0xeb, 0x38, 0xcf, 0xf3, 0x7b, 0xfd, 0x3e, 0xaf, 0xa3, 0xdf, 0xa0, 0x21, 0xb6,
	0x03, 0x34, 0x44, 0x0c, 0x33, 0x8b, 0x46, 0x84, 0x13, 0xc3, 0xfa, 0x84, 0x03, 0x0f, 0x21, 0x77,
	0x30, 0xc0, 0xee, 0xd0, 0x47, 0x56, 0x1f, 0x33, 0x1e, 0x61, 0x2f, 0xe6, 0xa8, 0xef, 0x93, 0x43,
	0x9a, 0xae, 0x0e, 0x50, 0x3f, 0x40, 0x91, 0x45, 0x43, 0xbc, 0xbe, 0x21, 0x24, 0x5c, 0x4a, 0x23,
	0x72, 0x84, 0xfa, 0x8e, 0x8f, 0x22, 0x8e, 0xdf, 0x61, 0xdf, 0xe5, 0x48, 0x0a, 0xae, 0x9b, 0x62,
	0x03, 0x8d, 0x08, 0x25, 0x6c, 0x7a, 0x83, 0x7c, 0x7e, 0x47, 0x3c, 0xf7, 0x0f, 0xf0, 0x40, 0xf9,
	0xf6, 0x83, 0xb2, 0xb7, 0x9d, 0x08, 0x1d, 0x11, 0xdf, 0xe5, 0x98, 0x0c, 0xf3, 0x46, 0x62, 0x35,
	0x54, 0x83, 0x80, 0x51, 0x3c, 0xc4, 0x1f, 0x62, 0xa4, 0xc0, 0xb8, 0x3f, 0x55, 0x47, 0x44, 0x08,
	0x57, 0x69, 0xdc, 0xcb, 0x7b, 0x94, 0x6d, 0x7a, 0x58, 0x7a, 0x24, 0x8e, 0x97, 0x38, 0x2c, 0xf6,
	0xde, 0x23, 0x9f, 0x4f, 0x43, 0x8b, 0x15, 0xe5, 0xe9, 0x3c, 0x82, 0xfa, 0x43, 0x9c, 0x2b, 0xd7,
	0x39, 0x6d, 0x89, 0xf8, 0x43, 0x09, 0x1e, 0x66, 0x62, 0x4f, 0xcf, 0xb5, 0x19, 0x28, 0x30, 0x63,
	0x31, 0x8a, 0x32, 0x18, 0x27, 0x44, 0x89, 0x83, 0xfb, 0x52, 0xe6, 0x66, 0x40, 0x02, 0x02, 0x3f,
	0x6d, 0xf1, 0x2b, 0x5d, 0xdd, 0xfa, 0x71, 0x5d, 0x5f, 0xdb, 0x4d, 0xa3, 0xb2, 0xcf, 0x5d, 0x8e,
	0x8c, 0xef, 0x9a, 0x5e, 0xcf, 0xaa, 0x6c, 0xe7, 0x8a, 0x7c, 0x81, 0x19, 0xaf, 0x6b, 0x9b, 0x4b,
	0xdb, 0xab, 0x3b, 0x9d, 0x05, 0xd3, 0x64, 0x35, 0x15, 0x7a, 0xad, 0xe5, 0xe3, 0xbf, 0x1b, 0xb5,
	0x5e, 0xa9, 0x97, 0xf1, 0x55, 0xd3, 0x6f, 0x67, 0x11, 0xc9, 0x3d, 0x04, 0x8e, 0x4b, 0xc0, 0xd1,
	0x5e, 0x94, 0xa3, 0x5b, 0x94, 0x93, 0x18, 0x65, 0x4e, 0xc6, 0x67, 0xfd, 0x16, 0xa4, 0xb8, 0x70,
	0x14, 0x4b, 0x80, 0xd0, 0x5c, 0x14, 0xa1, 0x3d, 0x2b, 0x26, 0x01, 0xd4, 0x2e, 0xc6, 0x2f, 0x4d,
	0xbf, 0xab, 0x40, 0xeb, 0x9d, 0x26, 0x01, 0x58, 0x96, 0x81, 0xe5, 0xe5, 0x05, 0x1c, 0xc7, 0x44,
	0x58, 0x72, 0x9d, 0xed, 0x0e, 0x8d, 0x92, 0xc3, 0x53, 0x38, 0xa5, 0xcb, 0xd5, 0x1a, 0xd5, 0x2b,
	0xca, 0x65, 0x8d, 0x2a, 0x71, 0x12, 0x8d, 0x4a, 0x6f, 0x81, 0xd9, 0xac, 0xac, 0x54, 0x6b, 0xd4,
	0xeb, 0x59, 0xb1, 0xac, 0x51, 0x4a, 0x17, 0xe3, 0x5b, 0x6e, 0x6c, 0x7a, 0x84, 0xf0, 0x3c, 0x5f,
	0xfd, 0xca, 0xa6, 0xb6, 0xbd, 0xba, 0xf3, 0xac, 0xea, 0xd8, 0xcc, 0xea, 0xf5, 0x4a, 0x9d, 0x8c,
	0x2f, 0x93, 0x5e, 0x14, 0x28, 0xae, 0x02, 0xc5, 0x6e, 0xc5, 0x5e, 0x14, 0x20, 0xca, 0x7c, 0x20,
	0xb3, 0xaa, 0xa9, 0x6e, 0x25, 0xfb, 0xe9, 0xbd, 0x04, 0x6d, 0xb9, 0x56, 0x2d, 0xb3, 0xcd, 0x79,
	0xc2, 0x59, 0x66, 0xcf, 0x74, 0x97, 0x99, 0x4d, 0xef, 0xe7, 0xd9, 0xc0, 0xe8, 0x55, 0x33, 0x5b,
	0x90, 0x9b, 0x64, 0x56, 0xe9, 0x64, 0xfc, 0xd6, 0xf4, 0xad, 0x6e, 0x88, 0x27, 0xf3, 0xd4, 0xc9,
	0x5d, 0xeb, 0x5d, 0x71, 0xab, 0x03, 0xd0, 0x2a, 0x00, 0xbd, 0x5a, 0x78, 0xbc, 0xe7, 0x2a, 0x4b,
	0xb6, 0x73, 0xf8, 0x1b, 0x63, 0x4d, 0x7f, 0x32, 0x7f, 0x1b, 0x6b, 0x25, 0x7b, 0xf0, 0xe9, 0x91,
	0x67, 0xfc, 0x1c, 0x25, 0x7b, 0x1d, 0x60, 0x5f, 0x03, 0x76, 0xff, 0x62, 0xd9, 0x95, 0x76, 0xb2,
	0xa0, 0xaa, 0xa4, 0xad, 0xb7, 0xc7, 0x23, 0x53, 0x3b, 0x19, 0x99, 0xda, 0xbf, 0x91, 0xa9, 0xfd,
	0x1c, 0x9b, 0xb5, 0x93, 0xb1, 0x59, 0xfb, 0x33, 0x36, 0x6b, 0x6f, 0x3a, 0x01, 0xe6, 0x07, 0xb1,
	0x67, 0xf9, 0xe4, 0xd0, 0x4e, 0xeb, 0x68, 0x64, 0x85, 0xd8, 0xb9, 0x42, 0x1a, 0x93, 0x4a, 0x1a,
	0x69, 0x29, 0xf6, 0x47, 0xf1, 0xdd, 0xb6, 0x79, 0x42, 0x11, 0xf3, 0x56, 0xe0, 0x83, 0xfb, 0xf8,
	0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x57, 0xbb, 0x31, 0x9e, 0x09, 0x00, 0x00,
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
	if len(m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) > 0 {
		for iNdEx := len(m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.PkiRevocationDistributionPointList) > 0 {
		for iNdEx := len(m.PkiRevocationDistributionPointList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PkiRevocationDistributionPointList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.RejectedCertificateList) > 0 {
		for iNdEx := len(m.RejectedCertificateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RejectedCertificateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ApprovedCertificatesBySubjectList) > 0 {
		for iNdEx := len(m.ApprovedCertificatesBySubjectList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApprovedCertificatesBySubjectList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.RevokedRootCertificates != nil {
		{
			size, err := m.RevokedRootCertificates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.ApprovedRootCertificates != nil {
		{
			size, err := m.ApprovedRootCertificates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.UniqueCertificateList) > 0 {
		for iNdEx := len(m.UniqueCertificateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UniqueCertificateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.RevokedCertificatesList) > 0 {
		for iNdEx := len(m.RevokedCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RevokedCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ProposedCertificateRevocationList) > 0 {
		for iNdEx := len(m.ProposedCertificateRevocationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposedCertificateRevocationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ChildCertificatesList) > 0 {
		for iNdEx := len(m.ChildCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChildCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ProposedCertificateList) > 0 {
		for iNdEx := len(m.ProposedCertificateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposedCertificateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ApprovedCertificatesList) > 0 {
		for iNdEx := len(m.ApprovedCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApprovedCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ApprovedCertificatesList) > 0 {
		for _, e := range m.ApprovedCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposedCertificateList) > 0 {
		for _, e := range m.ProposedCertificateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChildCertificatesList) > 0 {
		for _, e := range m.ChildCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposedCertificateRevocationList) > 0 {
		for _, e := range m.ProposedCertificateRevocationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RevokedCertificatesList) > 0 {
		for _, e := range m.RevokedCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UniqueCertificateList) > 0 {
		for _, e := range m.UniqueCertificateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ApprovedRootCertificates != nil {
		l = m.ApprovedRootCertificates.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.RevokedRootCertificates != nil {
		l = m.RevokedRootCertificates.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ApprovedCertificatesBySubjectList) > 0 {
		for _, e := range m.ApprovedCertificatesBySubjectList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RejectedCertificateList) > 0 {
		for _, e := range m.RejectedCertificateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PkiRevocationDistributionPointList) > 0 {
		for _, e := range m.PkiRevocationDistributionPointList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) > 0 {
		for _, e := range m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList {
			l = e.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedCertificatesList", wireType)
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
			m.ApprovedCertificatesList = append(m.ApprovedCertificatesList, ApprovedCertificates{})
			if err := m.ApprovedCertificatesList[len(m.ApprovedCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedCertificateList", wireType)
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
			m.ProposedCertificateList = append(m.ProposedCertificateList, ProposedCertificate{})
			if err := m.ProposedCertificateList[len(m.ProposedCertificateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildCertificatesList", wireType)
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
			m.ChildCertificatesList = append(m.ChildCertificatesList, ChildCertificates{})
			if err := m.ChildCertificatesList[len(m.ChildCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedCertificateRevocationList", wireType)
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
			m.ProposedCertificateRevocationList = append(m.ProposedCertificateRevocationList, ProposedCertificateRevocation{})
			if err := m.ProposedCertificateRevocationList[len(m.ProposedCertificateRevocationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokedCertificatesList", wireType)
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
			m.RevokedCertificatesList = append(m.RevokedCertificatesList, RevokedCertificates{})
			if err := m.RevokedCertificatesList[len(m.RevokedCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniqueCertificateList", wireType)
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
			m.UniqueCertificateList = append(m.UniqueCertificateList, UniqueCertificate{})
			if err := m.UniqueCertificateList[len(m.UniqueCertificateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedRootCertificates", wireType)
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
			if m.ApprovedRootCertificates == nil {
				m.ApprovedRootCertificates = &ApprovedRootCertificates{}
			}
			if err := m.ApprovedRootCertificates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokedRootCertificates", wireType)
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
			if m.RevokedRootCertificates == nil {
				m.RevokedRootCertificates = &RevokedRootCertificates{}
			}
			if err := m.RevokedRootCertificates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedCertificatesBySubjectList", wireType)
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
			m.ApprovedCertificatesBySubjectList = append(m.ApprovedCertificatesBySubjectList, ApprovedCertificatesBySubject{})
			if err := m.ApprovedCertificatesBySubjectList[len(m.ApprovedCertificatesBySubjectList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectedCertificateList", wireType)
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
			m.RejectedCertificateList = append(m.RejectedCertificateList, RejectedCertificate{})
			if err := m.RejectedCertificateList[len(m.RejectedCertificateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkiRevocationDistributionPointList", wireType)
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
			m.PkiRevocationDistributionPointList = append(m.PkiRevocationDistributionPointList, PkiRevocationDistributionPoint{})
			if err := m.PkiRevocationDistributionPointList[len(m.PkiRevocationDistributionPointList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkiRevocationDistributionPointsByIssuerSubjectKeyIDList", wireType)
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
			m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = append(m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList, PkiRevocationDistributionPointsByIssuerSubjectKeyID{})
			if err := m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList[len(m.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
