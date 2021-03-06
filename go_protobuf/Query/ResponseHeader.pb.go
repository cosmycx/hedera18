// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ResponseHeader.proto

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

// Every query receives a response containing the QueryResponseHeader. Either or both of the cost and stateProof fields may be blank, if the responseType didn��t ask for the cost or stateProof.
type ResponseHeader struct {
	NodeTransactionPrecheckCode NodeTransactionPrecheckCode `protobuf:"varint,1,opt,name=nodeTransactionPrecheckCode,proto3,enum=proto.NodeTransactionPrecheckCode" json:"nodeTransactionPrecheckCode,omitempty"`
	ResponseType                ResponseType                `protobuf:"varint,2,opt,name=responseType,proto3,enum=proto.ResponseType" json:"responseType,omitempty"`
	Cost                        uint64                      `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	StateProof                  []byte                      `protobuf:"bytes,4,opt,name=stateProof,proto3" json:"stateProof,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                    `json:"-"`
	XXX_unrecognized            []byte                      `json:"-"`
	XXX_sizecache               int32                       `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b707db0c2015df73, []int{0}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetNodeTransactionPrecheckCode() NodeTransactionPrecheckCode {
	if m != nil {
		return m.NodeTransactionPrecheckCode
	}
	return NodeTransactionPrecheckCode_OK
}

func (m *ResponseHeader) GetResponseType() ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ResponseType_ANSWER_ONLY
}

func (m *ResponseHeader) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *ResponseHeader) GetStateProof() []byte {
	if m != nil {
		return m.StateProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseHeader)(nil), "proto.ResponseHeader")
}

func init() { proto.RegisterFile("ResponseHeader.proto", fileDescriptor_b707db0c2015df73) }

var fileDescriptor_b707db0c2015df73 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8d, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x89, 0xae, 0x16, 0xc3, 0x71, 0x60, 0xb4, 0x58, 0x4f, 0x90, 0x65, 0x0b, 0xd9, 0x6a,
	0x0b, 0x2d, 0xec, 0xb5, 0xb1, 0x92, 0x35, 0xdc, 0x0b, 0x8c, 0xc9, 0x68, 0x4e, 0x31, 0x13, 0x26,
	0x51, 0xb8, 0xe7, 0xf5, 0x45, 0x84, 0xdc, 0x0a, 0xb9, 0x42, 0xab, 0x64, 0x66, 0xbe, 0xff, 0xfb,
	0xe1, 0xcc, 0x50, 0x8a, 0x1c, 0x12, 0x3d, 0x10, 0x3a, 0x92, 0x31, 0x0a, 0x67, 0xd6, 0x47, 0xe5,
	0x59, 0x9d, 0xaf, 0x05, 0x43, 0x42, 0x9b, 0x37, 0x1c, 0x7e, 0xb9, 0x1d, 0xb1, 0x3a, 0x79, 0xfa,
	0x24, 0xd9, 0xd6, 0xa1, 0xfe, 0x5b, 0xc1, 0x72, 0xdf, 0xa6, 0x1d, 0x5c, 0x04, 0x76, 0x54, 0x69,
	0x26, 0x21, 0xeb, 0xc9, 0xbe, 0xdf, 0xb3, 0xa3, 0x56, 0x75, 0x6a, 0x58, 0x5e, 0xf7, 0xbb, 0xfc,
	0xf8, 0xf8, 0x37, 0x69, 0xfe, 0xd3, 0xe8, 0x5b, 0x58, 0xc8, 0xdc, 0xbb, 0xde, 0x46, 0x6a, 0x0f,
	0x8a, 0xf6, 0x74, 0xd6, 0x9a, 0xea, 0x64, 0xf6, 0x40, 0xad, 0xa1, 0xb1, 0x9c, 0x72, 0x7b, 0xd8,
	0xa9, 0xa1, 0x31, 0xe5, 0xaf, 0x2f, 0x01, 0x52, 0xc6, 0x4c, 0x93, 0x30, 0xbf, 0xb4, 0x4d, 0xa7,
	0x86, 0x85, 0xa9, 0x36, 0x77, 0x57, 0xd0, 0x5b, 0xfe, 0x18, 0x3d, 0x39, 0x12, 0xf4, 0x98, 0xfc,
	0xab, 0x60, 0xf4, 0x23, 0xc6, 0xcd, 0xdc, 0xf7, 0x86, 0x5f, 0x38, 0xa9, 0xe7, 0xe3, 0x32, 0xdd,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x07, 0xf2, 0x78, 0x61, 0x01, 0x00, 0x00,
}
