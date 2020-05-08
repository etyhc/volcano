// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

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

type SERVICE int32

const (
	SERVICE_NONE  SERVICE = 0
	SERVICE_LOBBY SERVICE = 1
	SERVICE_MATCH SERVICE = 2
	SERVICE_ROOM  SERVICE = 3
)

var SERVICE_name = map[int32]string{
	0: "NONE",
	1: "LOBBY",
	2: "MATCH",
	3: "ROOM",
}

var SERVICE_value = map[string]int32{
	"NONE":  0,
	"LOBBY": 1,
	"MATCH": 2,
	"ROOM":  3,
}

func (x SERVICE) String() string {
	return proto.EnumName(SERVICE_name, int32(x))
}

func (SERVICE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

type CallReqMsg struct {
	Call                 string   `protobuf:"bytes,1,opt,name=call,proto3" json:"call,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallReqMsg) Reset()         { *m = CallReqMsg{} }
func (m *CallReqMsg) String() string { return proto.CompactTextString(m) }
func (*CallReqMsg) ProtoMessage()    {}
func (*CallReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *CallReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReqMsg.Unmarshal(m, b)
}
func (m *CallReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReqMsg.Marshal(b, m, deterministic)
}
func (m *CallReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReqMsg.Merge(m, src)
}
func (m *CallReqMsg) XXX_Size() int {
	return xxx_messageInfo_CallReqMsg.Size(m)
}
func (m *CallReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CallReqMsg proto.InternalMessageInfo

func (m *CallReqMsg) GetCall() string {
	if m != nil {
		return m.Call
	}
	return ""
}

func (m *CallReqMsg) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type CallResMsg struct {
	Raw                  []byte   `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResMsg) Reset()         { *m = CallResMsg{} }
func (m *CallResMsg) String() string { return proto.CompactTextString(m) }
func (*CallResMsg) ProtoMessage()    {}
func (*CallResMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *CallResMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResMsg.Unmarshal(m, b)
}
func (m *CallResMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResMsg.Marshal(b, m, deterministic)
}
func (m *CallResMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResMsg.Merge(m, src)
}
func (m *CallResMsg) XXX_Size() int {
	return xxx_messageInfo_CallResMsg.Size(m)
}
func (m *CallResMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CallResMsg proto.InternalMessageInfo

func (m *CallResMsg) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.SERVICE", SERVICE_name, SERVICE_value)
	proto.RegisterType((*CallReqMsg)(nil), "message.CallReqMsg")
	proto.RegisterType((*CallResMsg)(nil), "message.CallResMsg")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x55, 0x32, 0xe1, 0xe2, 0x72, 0x4e, 0xcc, 0xc9, 0x09, 0x4a, 0x2d, 0xf4, 0x2d, 0x4e, 0x17,
	0x12, 0xe2, 0x62, 0x49, 0x4e, 0xcc, 0xc9, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x41, 0x62, 0x89, 0x45, 0xe9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x20, 0x31, 0x10, 0x5b, 0x49, 0x0e,
	0xa6, 0xab, 0x18, 0xa4, 0x4b, 0x80, 0x8b, 0xb9, 0x28, 0xb1, 0x1c, 0xac, 0x89, 0x27, 0x08, 0xc4,
	0xd4, 0x32, 0xe6, 0x62, 0x0f, 0x76, 0x0d, 0x0a, 0xf3, 0x74, 0x76, 0x15, 0xe2, 0xe0, 0x62, 0xf1,
	0xf3, 0xf7, 0x73, 0x15, 0x60, 0x10, 0xe2, 0xe4, 0x62, 0xf5, 0xf1, 0x77, 0x72, 0x8a, 0x14, 0x60,
	0x04, 0x31, 0x7d, 0x1d, 0x43, 0x9c, 0x3d, 0x04, 0x98, 0x40, 0xf2, 0x41, 0xfe, 0xfe, 0xbe, 0x02,
	0xcc, 0x49, 0x6c, 0x60, 0xa7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x83, 0xd0, 0x24, 0x96,
	0xaa, 0x00, 0x00, 0x00,
}