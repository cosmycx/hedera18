// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ContractGetBytecode.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Get the bytecode for a smart contract instance
type ContractGetBytecodeQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractGetBytecodeQuery) Reset()         { *m = ContractGetBytecodeQuery{} }
func (m *ContractGetBytecodeQuery) String() string { return proto.CompactTextString(m) }
func (*ContractGetBytecodeQuery) ProtoMessage()    {}
func (*ContractGetBytecodeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d8843f452292c3, []int{0}
}

func (m *ContractGetBytecodeQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetBytecodeQuery.Unmarshal(m, b)
}
func (m *ContractGetBytecodeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetBytecodeQuery.Marshal(b, m, deterministic)
}
func (m *ContractGetBytecodeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetBytecodeQuery.Merge(m, src)
}
func (m *ContractGetBytecodeQuery) XXX_Size() int {
	return xxx_messageInfo_ContractGetBytecodeQuery.Size(m)
}
func (m *ContractGetBytecodeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetBytecodeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetBytecodeQuery proto.InternalMessageInfo

func (m *ContractGetBytecodeQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetBytecodeQuery) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

// Response when the client sends the node ContractGetBytecodeQuery
type ContractGetBytecodeResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Bytecode             []byte          `protobuf:"bytes,6,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ContractGetBytecodeResponse) Reset()         { *m = ContractGetBytecodeResponse{} }
func (m *ContractGetBytecodeResponse) String() string { return proto.CompactTextString(m) }
func (*ContractGetBytecodeResponse) ProtoMessage()    {}
func (*ContractGetBytecodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d8843f452292c3, []int{1}
}

func (m *ContractGetBytecodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetBytecodeResponse.Unmarshal(m, b)
}
func (m *ContractGetBytecodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetBytecodeResponse.Marshal(b, m, deterministic)
}
func (m *ContractGetBytecodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetBytecodeResponse.Merge(m, src)
}
func (m *ContractGetBytecodeResponse) XXX_Size() int {
	return xxx_messageInfo_ContractGetBytecodeResponse.Size(m)
}
func (m *ContractGetBytecodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetBytecodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetBytecodeResponse proto.InternalMessageInfo

func (m *ContractGetBytecodeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetBytecodeResponse) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractGetBytecodeQuery)(nil), "proto.ContractGetBytecodeQuery")
	proto.RegisterType((*ContractGetBytecodeResponse)(nil), "proto.ContractGetBytecodeResponse")
}

func init() { proto.RegisterFile("ContractGetBytecode.proto", fileDescriptor_72d8843f452292c3) }

var fileDescriptor_72d8843f452292c3 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x60, 0x91, 0xd1, 0x83, 0x1b, 0x14, 0x62, 0xbd, 0x2c, 0x3d, 0xc8, 0x22, 0x18,
	0x50, 0xdf, 0xa0, 0x2b, 0xa8, 0x37, 0x0d, 0xbe, 0xc0, 0x34, 0x1d, 0x4c, 0x05, 0x9b, 0x90, 0x44,
	0x21, 0x6f, 0x2f, 0xa4, 0xb1, 0x54, 0xe9, 0x29, 0xcc, 0xcc, 0x37, 0xdf, 0xfc, 0x81, 0x8b, 0xbd,
	0x1d, 0xa3, 0x47, 0x1d, 0x1f, 0x29, 0xb6, 0x29, 0x92, 0xb6, 0x3d, 0x49, 0xe7, 0x6d, 0xb4, 0xfc,
	0x30, 0x3f, 0xf5, 0x69, 0x8b, 0x61, 0xd0, 0x6f, 0xc9, 0x51, 0x98, 0x06, 0xf5, 0xe6, 0xf5, 0x8b,
	0x7c, 0x7a, 0x22, 0xec, 0xc9, 0x97, 0xd6, 0x99, 0xa2, 0xe0, 0xec, 0x18, 0x68, 0xd9, 0x6d, 0x12,
	0x88, 0x15, 0x7d, 0xde, 0xe6, 0xd7, 0x50, 0x99, 0xcc, 0x0a, 0xb6, 0x65, 0xbb, 0xe3, 0x3b, 0x3e,
	0xed, 0xc8, 0x85, 0x5b, 0x15, 0x82, 0xdf, 0x02, 0xe8, 0xe2, 0x79, 0x7e, 0x10, 0x07, 0x99, 0xdf,
	0x14, 0x7e, 0x3f, 0x0f, 0xd4, 0x02, 0x6a, 0x0c, 0x5c, 0xae, 0x9c, 0xfe, 0x4d, 0xc9, 0x6f, 0xfe,
	0x5d, 0x3f, 0x2f, 0xb6, 0xbf, 0xdf, 0x98, 0x03, 0xd4, 0x70, 0xd4, 0x15, 0x85, 0xa8, 0xb6, 0x6c,
	0x77, 0xa2, 0xe6, 0xba, 0xbd, 0x82, 0x46, 0xdb, 0x4f, 0x69, 0xa8, 0x27, 0x8f, 0x06, 0x83, 0x79,
	0xf7, 0xe8, 0x8c, 0x44, 0x37, 0x14, 0xe7, 0x07, 0x7e, 0xe3, 0x0b, 0xeb, 0xaa, 0x5c, 0xdd, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa7, 0x85, 0x4c, 0x72, 0x01, 0x00, 0x00,
}