// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/a.proto

package import_public // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/import_public"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import sub "github.com/gogo/protobuf/protoc-gen-gogo/testdata/import_public/sub"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// M from public import import_public/sub/a.proto
type M = sub.M

// E from public import import_public/sub/a.proto
type E = sub.E

var E_name = sub.E_name
var E_value = sub.E_value

const E_ZERO = E(sub.E_ZERO)

// Ignoring public import of Local from import_public/b.proto

type Public struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    sub.E    `protobuf:"varint,2,opt,name=e,proto3,enum=goproto.test.import_public.sub.E" json:"e,omitempty"`
	Local                *Local   `protobuf:"bytes,3,opt,name=local" json:"local,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Public) Reset()         { *m = Public{} }
func (m *Public) String() string { return proto.CompactTextString(m) }
func (*Public) ProtoMessage()    {}
func (*Public) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_2f23fa0a8a46426d, []int{0}
}
func (m *Public) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_Public.Unmarshal(m, b)
}
func (m *Public) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Public.Marshal(b, m, deterministic)
}
func (dst *Public) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Public.Merge(dst, src)
}
func (m *Public) XXX_Size() int {
	return xxx_messageInfo_Public.Size(m)
}
func (m *Public) XXX_DiscardUnknown() {
	xxx_messageInfo_Public.DiscardUnknown(m)
}

var xxx_messageInfo_Public proto.InternalMessageInfo

func (m *Public) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Public) GetE() sub.E {
	if m != nil {
		return m.E
	}
	return sub.E_ZERO
}

func (m *Public) GetLocal() *Local {
	if m != nil {
		return m.Local
	}
	return nil
}

func init() {
	proto.RegisterType((*Public)(nil), "goproto.test.import_public.Public")
}

func init() { proto.RegisterFile("import_public/a.proto", fileDescriptor_a_2f23fa0a8a46426d) }

var fileDescriptor_a_2f23fa0a8a46426d = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x4f, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x4a, 0xcf, 0x07, 0x33, 0xf4, 0x4a, 0x52, 0x8b, 0x4b, 0xf4, 0x50, 0xd4, 0x48,
	0x49, 0xa2, 0x6a, 0x29, 0x2e, 0x4d, 0x82, 0x69, 0x93, 0x42, 0x33, 0x2d, 0x09, 0x22, 0xac, 0xb4,
	0x98, 0x91, 0x8b, 0x2d, 0x00, 0x2c, 0x24, 0xa4, 0xcf, 0xc5, 0x98, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0xa4, 0xa8, 0x87, 0xdb, 0x12, 0xbd, 0xe2, 0xd2, 0x24, 0x3d, 0xdf, 0x20, 0xc6, 0x5c,
	0x90, 0x86, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0xc2, 0x1a, 0x5c, 0x83, 0x18, 0x53, 0x85,
	0xcc, 0xb9, 0x58, 0x73, 0xf2, 0x93, 0x13, 0x73, 0x24, 0x98, 0x09, 0xdb, 0xe2, 0x03, 0x52, 0x18,
	0x04, 0x51, 0xef, 0xe4, 0x18, 0x65, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x9f, 0x9e, 0x9f, 0x9e, 0xaf, 0x0f, 0xd6, 0x98, 0x54, 0x9a, 0x06, 0x61, 0x24, 0xeb, 0xa6,
	0xa7, 0xe6, 0xe9, 0x82, 0x25, 0x40, 0x66, 0xa5, 0x24, 0x96, 0x24, 0xea, 0xa3, 0x98, 0x17, 0xc0,
	0x10, 0xc0, 0x98, 0xc4, 0x06, 0x56, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x63, 0x98,
	0xb1, 0x5a, 0x01, 0x00, 0x00,
}
