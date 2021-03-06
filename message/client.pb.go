// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HiMsg struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiMsg) Reset()         { *m = HiMsg{} }
func (m *HiMsg) String() string { return proto.CompactTextString(m) }
func (*HiMsg) ProtoMessage()    {}
func (*HiMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *HiMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiMsg.Unmarshal(m, b)
}
func (m *HiMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiMsg.Marshal(b, m, deterministic)
}
func (m *HiMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiMsg.Merge(m, src)
}
func (m *HiMsg) XXX_Size() int {
	return xxx_messageInfo_HiMsg.Size(m)
}
func (m *HiMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HiMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HiMsg proto.InternalMessageInfo

func (m *HiMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HiMsg)(nil), "message.HiMsg")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 77 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x55, 0x92, 0xe4, 0x62, 0xf5, 0xc8, 0xf4, 0x2d, 0x4e, 0x17, 0x12, 0xe0, 0x62, 0xce, 0x2d,
	0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x93, 0xd8, 0xc0, 0x4a, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xc9, 0x1c, 0xd4, 0x3a, 0x00, 0x00, 0x00,
}
