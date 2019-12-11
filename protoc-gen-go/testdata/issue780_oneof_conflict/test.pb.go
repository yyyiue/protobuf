// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue780_oneof_conflict/test.proto

package oneoftest

import (
	fmt "fmt"
	proto "github.com/yyyiue/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Foo struct {
	// Types that are valid to be assigned to Bar:
	//	*Foo_GetBar
	Bar                  isFoo_Bar `protobuf_oneof:"bar"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_48462cafc802a68e, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

type isFoo_Bar interface {
	isFoo_Bar()
}

type Foo_GetBar struct {
	GetBar string `protobuf:"bytes,1,opt,name=get_bar,json=getBar,oneof"`
}

func (*Foo_GetBar) isFoo_Bar() {}

func (m *Foo) GetBar() isFoo_Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func (m *Foo) GetGetBar() string {
	if x, ok := m.GetBar().(*Foo_GetBar); ok {
		return x.GetBar
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Foo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Foo_GetBar)(nil),
	}
}

func init() {
	proto.RegisterType((*Foo)(nil), "oneoftest.Foo")
}

func init() { proto.RegisterFile("issue780_oneof_conflict/test.proto", fileDescriptor_48462cafc802a68e) }

var fileDescriptor_48462cafc802a68e = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0xb7, 0x30, 0x88, 0xcf, 0xcf, 0x4b, 0xcd, 0x4f, 0x8b, 0x4f, 0xce, 0xcf, 0x4b, 0xcb,
	0xc9, 0x4c, 0x2e, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x04, 0x4b, 0x81, 0x04, 0x94, 0xd4, 0xb9, 0x98, 0xdd, 0xf2, 0xf3, 0x85, 0x24, 0xb9, 0xd8, 0xd3,
	0x53, 0x4b, 0xe2, 0x93, 0x12, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x3d, 0x18, 0x82, 0xd8,
	0xd2, 0x53, 0x4b, 0x9c, 0x12, 0x8b, 0x9c, 0x58, 0xb9, 0x98, 0x93, 0x12, 0x8b, 0x00, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x12, 0x66, 0x0c, 0x02, 0x58, 0x00, 0x00, 0x00,
}
