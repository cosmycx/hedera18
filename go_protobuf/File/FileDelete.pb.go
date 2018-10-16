// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FileDelete.proto

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

// Delete the given file. After deletion, it will be marked as deleted and will have no contents. But information about it will continue to exist until it expires. A list of keys  was given when the file was created. All the keys on that list must sign transactions to create or modify the file, but any single one of them can be used to delete the file. Each "key" on that list may itself be a threshold key containing other keys (including other threshold keys).
type FileDeleteTransactionBody struct {
	FileID               *FileID  `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDeleteTransactionBody) Reset()         { *m = FileDeleteTransactionBody{} }
func (m *FileDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*FileDeleteTransactionBody) ProtoMessage()    {}
func (*FileDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ffe7db4276a630, []int{0}
}

func (m *FileDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDeleteTransactionBody.Unmarshal(m, b)
}
func (m *FileDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *FileDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDeleteTransactionBody.Merge(m, src)
}
func (m *FileDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_FileDeleteTransactionBody.Size(m)
}
func (m *FileDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_FileDeleteTransactionBody proto.InternalMessageInfo

func (m *FileDeleteTransactionBody) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func init() {
	proto.RegisterType((*FileDeleteTransactionBody)(nil), "proto.FileDeleteTransactionBody")
}

func init() { proto.RegisterFile("FileDelete.proto", fileDescriptor_45ffe7db4276a630) }

var fileDescriptor_45ffe7db4276a630 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x70, 0xcb, 0xcc, 0x49,
	0x75, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x09, 0x25, 0x27,
	0x2e, 0x49, 0x84, 0xe2, 0x90, 0xa2, 0xc4, 0xbc, 0xe2, 0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0x3c, 0xa7,
	0xfc, 0x94, 0x4a, 0x21, 0x55, 0x2e, 0xb6, 0xb4, 0xcc, 0x9c, 0x54, 0x4f, 0x17, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x5e, 0x88, 0x26, 0x3d, 0x37, 0xb0, 0x60, 0x10, 0x54, 0xd2, 0x49, 0x8d,
	0x4b, 0x29, 0x39, 0x3f, 0x57, 0x2f, 0x23, 0x35, 0x25, 0xb5, 0x28, 0x31, 0x23, 0xb1, 0x38, 0x23,
	0xbd, 0x28, 0xb1, 0x20, 0x43, 0x2f, 0xb1, 0x20, 0x13, 0xaa, 0x3e, 0x2b, 0xb1, 0x2c, 0x31, 0x80,
	0x31, 0x89, 0x0d, 0xcc, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x83, 0xe2, 0xdf, 0x9f,
	0x00, 0x00, 0x00,
}
