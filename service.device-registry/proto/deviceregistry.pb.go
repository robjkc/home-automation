// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.device-registry/proto/deviceregistry.proto

package deviceregistryproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/jakewright/home-automation/tools/protoc-gen-jrpc/proto"
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

type BarRequest struct {
	Baz                  string   `protobuf:"bytes,1,opt,name=baz,proto3" json:"baz,omitempty"`
	Test                 string   `protobuf:"bytes,2,opt,name=test,proto3" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_149d5a60e0054bdc, []int{0}
}

func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRequest.Unmarshal(m, b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
}
func (m *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(m, src)
}
func (m *BarRequest) XXX_Size() int {
	return xxx_messageInfo_BarRequest.Size(m)
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

func (m *BarRequest) GetBaz() string {
	if m != nil {
		return m.Baz
	}
	return ""
}

func (m *BarRequest) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

type BarResponse struct {
	Bat                  int32    `protobuf:"varint,1,opt,name=bat,proto3" json:"bat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarResponse) Reset()         { *m = BarResponse{} }
func (m *BarResponse) String() string { return proto.CompactTextString(m) }
func (*BarResponse) ProtoMessage()    {}
func (*BarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_149d5a60e0054bdc, []int{1}
}

func (m *BarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarResponse.Unmarshal(m, b)
}
func (m *BarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarResponse.Marshal(b, m, deterministic)
}
func (m *BarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarResponse.Merge(m, src)
}
func (m *BarResponse) XXX_Size() int {
	return xxx_messageInfo_BarResponse.Size(m)
}
func (m *BarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BarResponse proto.InternalMessageInfo

func (m *BarResponse) GetBat() int32 {
	if m != nil {
		return m.Bat
	}
	return 0
}

func init() {
	proto.RegisterType((*BarRequest)(nil), "deviceregistryproto.BarRequest")
	proto.RegisterType((*BarResponse)(nil), "deviceregistryproto.BarResponse")
}

func init() {
	proto.RegisterFile("service.device-registry/proto/deviceregistry.proto", fileDescriptor_149d5a60e0054bdc)
}

var fileDescriptor_149d5a60e0054bdc = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0xba, 0x45, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0x45,
	0x95, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x10, 0x51, 0x98, 0xa0, 0x1e, 0x58, 0x50, 0x48,
	0x18, 0x55, 0x14, 0x2c, 0x28, 0xa5, 0x56, 0x92, 0x9f, 0x9f, 0x53, 0x0c, 0xd1, 0x96, 0xac, 0x9b,
	0x9e, 0x9a, 0xa7, 0x9b, 0x55, 0x54, 0x90, 0x0c, 0x35, 0x06, 0xc4, 0x84, 0x68, 0x56, 0x32, 0xe2,
	0xe2, 0x72, 0x4a, 0x2c, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x4e,
	0x4a, 0xac, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x84, 0xb8, 0x58, 0x4a,
	0x52, 0x8b, 0x4b, 0x24, 0x98, 0xc0, 0x42, 0x60, 0xb6, 0x92, 0x3c, 0x17, 0x37, 0x58, 0x4f, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x44, 0x53, 0x09, 0x58, 0x13, 0x2b, 0x48, 0x53, 0x89, 0x51, 0x1f,
	0x23, 0x17, 0x9f, 0x0b, 0xd8, 0x51, 0x41, 0x50, 0x47, 0x09, 0xc5, 0x72, 0x31, 0x3b, 0x25, 0x16,
	0x09, 0xc9, 0xeb, 0x61, 0x71, 0xac, 0x1e, 0xc2, 0x05, 0x52, 0x0a, 0xb8, 0x15, 0x40, 0xac, 0x53,
	0x12, 0x3e, 0xf4, 0x59, 0x82, 0x95, 0x8b, 0xd9, 0xdd, 0x35, 0xe4, 0xd0, 0x67, 0x09, 0x36, 0x21,
	0x16, 0xfd, 0xb4, 0xfc, 0x7c, 0x29, 0xd9, 0xa6, 0xad, 0x12, 0x92, 0x5c, 0xe2, 0x38, 0x42, 0xcf,
	0x29, 0x32, 0x2a, 0x3c, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x2b,
	0x31, 0x3b, 0xb5, 0xbc, 0x28, 0x33, 0x3d, 0xa3, 0x44, 0x3f, 0x23, 0x3f, 0x37, 0x55, 0x37, 0xb1,
	0xb4, 0x24, 0x3f, 0x37, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x1f, 0x6f, 0xf0, 0x5b, 0x63, 0x71, 0x5a,
	0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x37, 0x52, 0x29, 0xba, 0x01,
	0x00, 0x00,
}