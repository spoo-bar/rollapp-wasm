// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/cwerrors/v1/cwerrors.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// ModuleErrors defines the module level error codes
type ModuleErrors int32

const (
	// ERR_UNKNOWN is the default error code
	ModuleErrors_ERR_UNKNOWN ModuleErrors = 0
	// ERR_CALLBACK_EXECUTION_FAILED is the error code for when the error callback fails
	ModuleErrors_ERR_CALLBACK_EXECUTION_FAILED ModuleErrors = 1
)

var ModuleErrors_name = map[int32]string{
	0: "ERR_UNKNOWN",
	1: "ERR_CALLBACK_EXECUTION_FAILED",
}

var ModuleErrors_value = map[string]int32{
	"ERR_UNKNOWN":                   0,
	"ERR_CALLBACK_EXECUTION_FAILED": 1,
}

func (x ModuleErrors) String() string {
	return proto.EnumName(ModuleErrors_name, int32(x))
}

func (ModuleErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5547f0c109cd175, []int{0}
}

// SudoError defines the sudo message for the error callback
type SudoError struct {
	// module_name is the name of the module throwing the error
	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// error_code is the module level error code
	ErrorCode int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// contract_address is the address of the contract which will receive the
	// error callback
	ContractAddress string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// input_payload is any input which caused the error
	InputPayload string `protobuf:"bytes,4,opt,name=input_payload,json=inputPayload,proto3" json:"input_payload,omitempty"`
	// error_message is the error message
	ErrorMessage string `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (m *SudoError) Reset()         { *m = SudoError{} }
func (m *SudoError) String() string { return proto.CompactTextString(m) }
func (*SudoError) ProtoMessage()    {}
func (*SudoError) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5547f0c109cd175, []int{0}
}
func (m *SudoError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SudoError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SudoError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SudoError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SudoError.Merge(m, src)
}
func (m *SudoError) XXX_Size() int {
	return m.Size()
}
func (m *SudoError) XXX_DiscardUnknown() {
	xxx_messageInfo_SudoError.DiscardUnknown(m)
}

var xxx_messageInfo_SudoError proto.InternalMessageInfo

func (m *SudoError) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *SudoError) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *SudoError) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *SudoError) GetInputPayload() string {
	if m != nil {
		return m.InputPayload
	}
	return ""
}

func (m *SudoError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("archway.cwerrors.v1.ModuleErrors", ModuleErrors_name, ModuleErrors_value)
	proto.RegisterType((*SudoError)(nil), "archway.cwerrors.v1.SudoError")
}

func init() {
	proto.RegisterFile("archway/cwerrors/v1/cwerrors.proto", fileDescriptor_d5547f0c109cd175)
}

var fileDescriptor_d5547f0c109cd175 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xc1, 0x4e, 0xc2, 0x30,
	0x18, 0xc0, 0xf1, 0x55, 0xc5, 0x84, 0x82, 0x81, 0x4c, 0x0f, 0x8b, 0x09, 0x13, 0xf1, 0x82, 0x26,
	0xb2, 0x10, 0xe3, 0x03, 0x8c, 0x39, 0x13, 0x02, 0x0c, 0x33, 0x25, 0x1a, 0x2f, 0x4b, 0x59, 0x9b,
	0x41, 0xb2, 0xad, 0x4b, 0xbb, 0x01, 0xf3, 0x29, 0x7c, 0x24, 0x8f, 0x1e, 0x39, 0x7a, 0x34, 0xf0,
	0x22, 0x66, 0xad, 0xe2, 0xed, 0xcb, 0x6f, 0xff, 0x7c, 0xdb, 0x5a, 0xd8, 0x42, 0xcc, 0x9f, 0x2d,
	0x51, 0x6e, 0xf8, 0x4b, 0xc2, 0x18, 0x65, 0xdc, 0x58, 0x74, 0x77, 0x73, 0x27, 0x61, 0x34, 0xa5,
	0xea, 0xf1, 0x6f, 0xd3, 0xd9, 0xf9, 0xa2, 0x7b, 0x7a, 0x12, 0xd0, 0x80, 0x8a, 0xe7, 0x46, 0x31,
	0xc9, 0xb4, 0xf5, 0x01, 0x60, 0xf9, 0x31, 0xc3, 0xd4, 0x2e, 0x3a, 0xf5, 0x0c, 0x56, 0x22, 0x8a,
	0xb3, 0x90, 0x78, 0x31, 0x8a, 0x88, 0x06, 0x9a, 0xa0, 0x5d, 0x76, 0xa1, 0x24, 0x07, 0x45, 0x44,
	0x6d, 0x40, 0x28, 0x36, 0x7a, 0x3e, 0xc5, 0x44, 0xdb, 0x6b, 0x82, 0x76, 0xc9, 0x2d, 0x0b, 0xb1,
	0x28, 0x26, 0xea, 0x25, 0xac, 0xfb, 0x34, 0x4e, 0x19, 0xf2, 0x53, 0x0f, 0x61, 0xcc, 0x08, 0xe7,
	0xda, 0xbe, 0x58, 0x52, 0xfb, 0x73, 0x53, 0xb2, 0x7a, 0x01, 0x8f, 0xe6, 0x71, 0x92, 0xa5, 0x5e,
	0x82, 0xf2, 0x90, 0x22, 0xac, 0x1d, 0x88, 0xae, 0x2a, 0xf0, 0x41, 0x5a, 0x11, 0xc9, 0xd7, 0x45,
	0x84, 0x73, 0x14, 0x10, 0xad, 0x24, 0x23, 0x81, 0x23, 0x69, 0x57, 0x3d, 0x58, 0x1d, 0x89, 0x2f,
	0x14, 0xff, 0xc0, 0xd5, 0x1a, 0xac, 0xd8, 0xae, 0xeb, 0x4d, 0x9c, 0x81, 0x33, 0x7e, 0x76, 0xea,
	0x8a, 0x7a, 0x0e, 0x1b, 0x05, 0x58, 0xe6, 0x70, 0xd8, 0x33, 0xad, 0x81, 0x67, 0xbf, 0xd8, 0xd6,
	0xe4, 0xa9, 0x3f, 0x76, 0xbc, 0x7b, 0xb3, 0x3f, 0xb4, 0xef, 0xea, 0xa0, 0x37, 0xfe, 0xdc, 0xe8,
	0x60, 0xbd, 0xd1, 0xc1, 0xf7, 0x46, 0x07, 0xef, 0x5b, 0x5d, 0x59, 0x6f, 0x75, 0xe5, 0x6b, 0xab,
	0x2b, 0xaf, 0xb7, 0xc1, 0x3c, 0x9d, 0x65, 0xd3, 0x8e, 0x4f, 0x23, 0x03, 0xe7, 0x11, 0x89, 0xf9,
	0x9c, 0xc6, 0xab, 0xfc, 0xcd, 0x60, 0x34, 0x0c, 0x51, 0x92, 0x5c, 0x2f, 0x11, 0x8f, 0x8c, 0xd5,
	0xff, 0x75, 0xa4, 0x79, 0x42, 0xf8, 0xf4, 0x50, 0x1c, 0xef, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x29, 0xe3, 0x74, 0xaf, 0x01, 0x00, 0x00,
}

func (m *SudoError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SudoError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SudoError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintCwerrors(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.InputPayload) > 0 {
		i -= len(m.InputPayload)
		copy(dAtA[i:], m.InputPayload)
		i = encodeVarintCwerrors(dAtA, i, uint64(len(m.InputPayload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCwerrors(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ErrorCode != 0 {
		i = encodeVarintCwerrors(dAtA, i, uint64(m.ErrorCode))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintCwerrors(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCwerrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovCwerrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SudoError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovCwerrors(uint64(l))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovCwerrors(uint64(m.ErrorCode))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCwerrors(uint64(l))
	}
	l = len(m.InputPayload)
	if l > 0 {
		n += 1 + l + sovCwerrors(uint64(l))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovCwerrors(uint64(l))
	}
	return n
}

func sovCwerrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCwerrors(x uint64) (n int) {
	return sovCwerrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SudoError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCwerrors
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
			return fmt.Errorf("proto: SudoError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SudoError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwerrors
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
				return ErrInvalidLengthCwerrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwerrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwerrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwerrors
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
				return ErrInvalidLengthCwerrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwerrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwerrors
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
				return ErrInvalidLengthCwerrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwerrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputPayload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwerrors
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
				return ErrInvalidLengthCwerrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwerrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCwerrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCwerrors
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
func skipCwerrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCwerrors
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
					return 0, ErrIntOverflowCwerrors
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
					return 0, ErrIntOverflowCwerrors
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
				return 0, ErrInvalidLengthCwerrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCwerrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCwerrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCwerrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCwerrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCwerrors = fmt.Errorf("proto: unexpected end of group")
)
