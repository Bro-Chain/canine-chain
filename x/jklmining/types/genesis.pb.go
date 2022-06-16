// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jklmining/genesis.proto

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

// GenesisState defines the jklmining module's genesis state.
type GenesisState struct {
	Params           Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SaveRequestsList []SaveRequests `protobuf:"bytes,2,rep,name=saveRequestsList,proto3" json:"saveRequestsList"`
	MinersList       []Miners       `protobuf:"bytes,3,rep,name=minersList,proto3" json:"minersList"`
	MinedList        []Mined        `protobuf:"bytes,4,rep,name=minedList,proto3" json:"minedList"`
	MinedCount       uint64         `protobuf:"varint,5,opt,name=minedCount,proto3" json:"minedCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a6b599b8814acc, []int{0}
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

func (m *GenesisState) GetSaveRequestsList() []SaveRequests {
	if m != nil {
		return m.SaveRequestsList
	}
	return nil
}

func (m *GenesisState) GetMinersList() []Miners {
	if m != nil {
		return m.MinersList
	}
	return nil
}

func (m *GenesisState) GetMinedList() []Mined {
	if m != nil {
		return m.MinedList
	}
	return nil
}

func (m *GenesisState) GetMinedCount() uint64 {
	if m != nil {
		return m.MinedCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "jackaldao.canine.jklmining.GenesisState")
}

func init() { proto.RegisterFile("jklmining/genesis.proto", fileDescriptor_78a6b599b8814acc) }

var fileDescriptor_78a6b599b8814acc = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0x3a, 0x31,
	0x1c, 0xc4, 0xb7, 0xc0, 0x8f, 0xe4, 0x57, 0x3c, 0x98, 0xc6, 0x3f, 0x64, 0x13, 0x2b, 0x72, 0xda,
	0x8b, 0x6d, 0x82, 0x2f, 0x60, 0x30, 0x46, 0x0e, 0x9a, 0x18, 0xb8, 0x71, 0x31, 0x85, 0x6d, 0xd6,
	0x02, 0xdb, 0xe2, 0xb6, 0x18, 0x7d, 0x0b, 0x1f, 0x8b, 0x23, 0x47, 0x2f, 0x1a, 0x03, 0x2f, 0x62,
	0x68, 0xab, 0xbb, 0x86, 0x88, 0xb7, 0xcd, 0x64, 0xe6, 0x33, 0xb3, 0xfd, 0xc2, 0xc3, 0xd1, 0x78,
	0x92, 0x0a, 0x29, 0x64, 0x42, 0x13, 0x2e, 0xb9, 0x16, 0x9a, 0x4c, 0x33, 0x65, 0x14, 0x0a, 0x47,
	0x6c, 0x38, 0x66, 0x93, 0x98, 0x29, 0x32, 0x64, 0x52, 0x48, 0x4e, 0xbe, 0x9d, 0xe1, 0x5e, 0xa2,
	0x12, 0x65, 0x6d, 0x74, 0xfd, 0xe5, 0x12, 0xe1, 0x41, 0x8e, 0x9a, 0xb2, 0x8c, 0xa5, 0x9e, 0x14,
	0x1e, 0xe5, 0xba, 0x66, 0x8f, 0xfc, 0x2e, 0xe3, 0x0f, 0x33, 0xae, 0x8d, 0xde, 0x8c, 0xa5, 0x42,
	0xf2, 0xec, 0x4b, 0xdf, 0xff, 0xa9, 0xc7, 0x4e, 0x6e, 0xbe, 0x95, 0xe0, 0xce, 0x95, 0x5b, 0xda,
	0x33, 0xcc, 0x70, 0x74, 0x0e, 0xab, 0xae, 0xae, 0x0e, 0x1a, 0x20, 0xaa, 0xb5, 0x9a, 0xe4, 0xf7,
	0xe5, 0xe4, 0xd6, 0x3a, 0xdb, 0x95, 0xf9, 0xfb, 0x71, 0xd0, 0xf5, 0x39, 0xd4, 0x87, 0xbb, 0xeb,
	0x61, 0x5d, 0xbf, 0xeb, 0x5a, 0x68, 0x53, 0x2f, 0x35, 0xca, 0x51, 0xad, 0x15, 0x6d, 0x63, 0xf5,
	0x0a, 0x19, 0x4f, 0xdc, 0xe0, 0xa0, 0x0e, 0x84, 0xee, 0xaf, 0x2c, 0xb5, 0x6c, 0xa9, 0x5b, 0x17,
	0xde, 0x58, 0xb7, 0xe7, 0x15, 0xb2, 0xe8, 0x12, 0xfe, 0xb7, 0xef, 0x60, 0x41, 0x15, 0x0b, 0x3a,
	0xf9, 0x0b, 0x14, 0x7b, 0x4e, 0x9e, 0x44, 0xd8, 0x0d, 0x8a, 0x2f, 0xd4, 0x4c, 0x9a, 0xfa, 0xbf,
	0x06, 0x88, 0x2a, 0xdd, 0x82, 0xd2, 0xee, 0xcc, 0x97, 0x18, 0x2c, 0x96, 0x18, 0x7c, 0x2c, 0x31,
	0x78, 0x59, 0xe1, 0x60, 0xb1, 0xc2, 0xc1, 0xeb, 0x0a, 0x07, 0x7d, 0x92, 0x08, 0x73, 0x3f, 0x1b,
	0x90, 0xa1, 0x4a, 0xa9, 0xeb, 0x3d, 0x8d, 0x99, 0xa2, 0xae, 0x98, 0x3e, 0xd1, 0xfc, 0x5e, 0xe6,
	0x79, 0xca, 0xf5, 0xa0, 0x6a, 0x0f, 0x76, 0xf6, 0x19, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x0d, 0x33,
	0xbc, 0x63, 0x02, 0x00, 0x00,
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
	if m.MinedCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinedCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MinedList) > 0 {
		for iNdEx := len(m.MinedList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinedList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.MinersList) > 0 {
		for iNdEx := len(m.MinersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SaveRequestsList) > 0 {
		for iNdEx := len(m.SaveRequestsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SaveRequestsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SaveRequestsList) > 0 {
		for _, e := range m.SaveRequestsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MinersList) > 0 {
		for _, e := range m.MinersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MinedList) > 0 {
		for _, e := range m.MinedList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MinedCount != 0 {
		n += 1 + sovGenesis(uint64(m.MinedCount))
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
				return fmt.Errorf("proto: wrong wireType = %d for field SaveRequestsList", wireType)
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
			m.SaveRequestsList = append(m.SaveRequestsList, SaveRequests{})
			if err := m.SaveRequestsList[len(m.SaveRequestsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinersList", wireType)
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
			m.MinersList = append(m.MinersList, Miners{})
			if err := m.MinersList[len(m.MinersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinedList", wireType)
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
			m.MinedList = append(m.MinedList, Mined{})
			if err := m.MinedList[len(m.MinedList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinedCount", wireType)
			}
			m.MinedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinedCount |= uint64(b&0x7F) << shift
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
