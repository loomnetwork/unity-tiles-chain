// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	TileMapTx
	TileMapState
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

type TileMapTx struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TileMapTx) Reset()                    { *m = TileMapTx{} }
func (m *TileMapTx) String() string            { return proto.CompactTextString(m) }
func (*TileMapTx) ProtoMessage()               {}
func (*TileMapTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *TileMapTx) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type TileMapState struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TileMapState) Reset()                    { *m = TileMapState{} }
func (m *TileMapState) String() string            { return proto.CompactTextString(m) }
func (*TileMapState) ProtoMessage()               {}
func (*TileMapState) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *TileMapState) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*TileMapTx)(nil), "TileMapTx")
	proto.RegisterType((*TileMapState)(nil), "TileMapState")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0xf2, 0x5c, 0x9c, 0x21, 0x99,
	0x39, 0xa9, 0xbe, 0x89, 0x05, 0x21, 0x15, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x12, 0x17, 0x0f, 0x54, 0x41, 0x70, 0x49,
	0x62, 0x49, 0x2a, 0x36, 0x35, 0x49, 0x6c, 0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0xf6, 0xbc, 0x38, 0x60, 0x00, 0x00, 0x00,
}