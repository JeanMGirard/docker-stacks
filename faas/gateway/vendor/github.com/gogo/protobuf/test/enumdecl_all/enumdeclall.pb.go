// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumdeclall.proto

/*
	Package enumdeclall is a generated protocol buffer package.

	It is generated from these files:
		enumdeclall.proto

	It has these top-level messages:
		Message
*/
package enumdeclall

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var MyEnum_name = map[int32]string{
	0: "A",
	1: "B",
}
var MyEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyEnum) String() string {
	return proto.EnumName(MyEnum_name, int32(x))
}
func (MyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnumdeclall, []int{0} }

type MyOtherEnum int32

const (
	C MyOtherEnum = 0
	D MyOtherEnum = 1
)

var MyOtherEnum_name = map[int32]string{
	0: "C",
	1: "D",
}
var MyOtherEnum_value = map[string]int32{
	"C": 0,
	"D": 1,
}

func (x MyOtherEnum) String() string {
	return proto.EnumName(MyOtherEnum_name, int32(x))
}
func (MyOtherEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnumdeclall, []int{1} }

type Message struct {
	EnumeratedField      MyEnum      `protobuf:"varint,1,opt,name=enumerated_field,json=enumeratedField,proto3,enum=enumdeclall.MyEnum" json:"enumerated_field,omitempty"`
	OtherenumeratedField MyOtherEnum `protobuf:"varint,2,opt,name=otherenumerated_field,json=otherenumeratedField,proto3,enum=enumdeclall.MyOtherEnum" json:"otherenumerated_field,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorEnumdeclall, []int{0} }

func (m *Message) GetEnumeratedField() MyEnum {
	if m != nil {
		return m.EnumeratedField
	}
	return A
}

func (m *Message) GetOtherenumeratedField() MyOtherEnum {
	if m != nil {
		return m.OtherenumeratedField
	}
	return C
}

func init() {
	proto.RegisterType((*Message)(nil), "enumdeclall.Message")
	proto.RegisterEnum("enumdeclall.MyEnum", MyEnum_name, MyEnum_value)
	proto.RegisterEnum("enumdeclall.MyOtherEnum", MyOtherEnum_name, MyOtherEnum_value)
}
func (this *Message) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Message")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Message but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Message but is not nil && this == nil")
	}
	if this.EnumeratedField != that1.EnumeratedField {
		return fmt.Errorf("EnumeratedField this(%v) Not Equal that(%v)", this.EnumeratedField, that1.EnumeratedField)
	}
	if this.OtherenumeratedField != that1.OtherenumeratedField {
		return fmt.Errorf("OtherenumeratedField this(%v) Not Equal that(%v)", this.OtherenumeratedField, that1.OtherenumeratedField)
	}
	return nil
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.EnumeratedField != that1.EnumeratedField {
		return false
	}
	if this.OtherenumeratedField != that1.OtherenumeratedField {
		return false
	}
	return true
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EnumeratedField != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEnumdeclall(dAtA, i, uint64(m.EnumeratedField))
	}
	if m.OtherenumeratedField != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEnumdeclall(dAtA, i, uint64(m.OtherenumeratedField))
	}
	return i, nil
}

func encodeVarintEnumdeclall(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMessage(r randyEnumdeclall, easy bool) *Message {
	this := &Message{}
	this.EnumeratedField = MyEnum([]int32{0, 1}[r.Intn(2)])
	this.OtherenumeratedField = MyOtherEnum([]int32{0, 1}[r.Intn(2)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyEnumdeclall interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEnumdeclall(r randyEnumdeclall) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEnumdeclall(r randyEnumdeclall) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneEnumdeclall(r)
	}
	return string(tmps)
}
func randUnrecognizedEnumdeclall(r randyEnumdeclall, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEnumdeclall(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEnumdeclall(dAtA []byte, r randyEnumdeclall, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEnumdeclall(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEnumdeclall(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.EnumeratedField != 0 {
		n += 1 + sovEnumdeclall(uint64(m.EnumeratedField))
	}
	if m.OtherenumeratedField != 0 {
		n += 1 + sovEnumdeclall(uint64(m.OtherenumeratedField))
	}
	return n
}

func sovEnumdeclall(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnumdeclall(x uint64) (n int) {
	return sovEnumdeclall(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnumdeclall
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnumeratedField", wireType)
			}
			m.EnumeratedField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnumdeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnumeratedField |= (MyEnum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherenumeratedField", wireType)
			}
			m.OtherenumeratedField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnumdeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherenumeratedField |= (MyOtherEnum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEnumdeclall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnumdeclall
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
func skipEnumdeclall(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnumdeclall
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
					return 0, ErrIntOverflowEnumdeclall
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnumdeclall
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEnumdeclall
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnumdeclall
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnumdeclall(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnumdeclall = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnumdeclall   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("enumdeclall.proto", fileDescriptorEnumdeclall) }

var fileDescriptorEnumdeclall = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcd, 0x2b, 0xcd,
	0x4d, 0x49, 0x4d, 0xce, 0x49, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x57,
	0x69, 0x06, 0x23, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x1d, 0x97, 0x00,
	0xc8, 0xa4, 0xd4, 0xa2, 0xc4, 0x92, 0xd4, 0x94, 0xf8, 0xb4, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x3e, 0x23, 0x61, 0x3d, 0x64, 0x5b, 0x7d, 0x2b, 0x5d, 0xf3, 0x4a, 0x73, 0x83,
	0xf8, 0x11, 0x8a, 0xdd, 0x40, 0x6a, 0x85, 0x7c, 0xb9, 0x44, 0xf3, 0x4b, 0x32, 0x52, 0x8b, 0x30,
	0x0c, 0x61, 0x02, 0x1b, 0x22, 0x81, 0x66, 0x88, 0x3f, 0x48, 0x2d, 0xd8, 0x24, 0x11, 0x34, 0x6d,
	0x60, 0xe3, 0xb4, 0x64, 0xb8, 0xd8, 0x20, 0x36, 0x09, 0xb1, 0x72, 0x31, 0x3a, 0x0a, 0x30, 0x80,
	0x28, 0x27, 0x01, 0x46, 0x29, 0x96, 0x8e, 0xc5, 0x72, 0x0c, 0x5a, 0xaa, 0x5c, 0xdc, 0x48, 0x46,
	0x80, 0xe4, 0x9c, 0x21, 0x4a, 0x5c, 0x04, 0x18, 0xa5, 0x38, 0x40, 0x4a, 0x0e, 0x2c, 0x91, 0x63,
	0x74, 0xd2, 0x79, 0xf0, 0x50, 0x8e, 0xf1, 0xc7, 0x43, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x77,
	0x3c, 0x92, 0x63, 0x3c, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0xfc, 0xf1, 0x48, 0x8e, 0xa1, 0xe1, 0xb1, 0x1c, 0xc3, 0x8e, 0xc7, 0x72,
	0x0c, 0x49, 0x6c, 0xe0, 0x40, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x91, 0xd9, 0xaf,
	0x65, 0x01, 0x00, 0x00,
}
