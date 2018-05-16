// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/examples/plugins/evmexample/types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/examples/plugins/evmexample/types/types.proto

It has these top-level messages:
	Dummy
	WrapValue
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Dummy struct {
}

func (m *Dummy) Reset()                    { *m = Dummy{} }
func (m *Dummy) String() string            { return proto.CompactTextString(m) }
func (*Dummy) ProtoMessage()               {}
func (*Dummy) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type WrapValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *WrapValue) Reset()                    { *m = WrapValue{} }
func (m *WrapValue) String() string            { return proto.CompactTextString(m) }
func (*WrapValue) ProtoMessage()               {}
func (*WrapValue) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *WrapValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Dummy)(nil), "Dummy")
	proto.RegisterType((*WrapValue)(nil), "WrapValue")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/examples/plugins/evmexample/types/types.proto", fileDescriptorTypes)
}

var fileDescriptorTypes = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52,
	0x8b, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x53, 0xcb, 0x72, 0xa1, 0x62, 0xfa,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x52, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x89, 0x9d, 0x8b,
	0xd5, 0xa5, 0x34, 0x37, 0xb7, 0x52, 0x49, 0x91, 0x8b, 0x33, 0xbc, 0x28, 0xb1, 0x20, 0x2c, 0x31,
	0xa7, 0x34, 0x55, 0x48, 0x84, 0x8b, 0xb5, 0x0c, 0xc4, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x82, 0x70, 0x92, 0xd8, 0xc0, 0x5a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xd3, 0xec,
	0xd2, 0x82, 0x00, 0x00, 0x00,
}
