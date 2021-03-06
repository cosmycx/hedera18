// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Response.proto

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

// A single response, which is returned from the node to the client, after the client sent the node a query. This includes all responses.
type Response struct {
	// Types that are valid to be assigned to Response:
	//	*Response_GetByKey
	//	*Response_GetBySolidityID
	//	*Response_ContractCallLocal
	//	*Response_ContractGetBytecodeResponse
	//	*Response_ContractGetInfo
	//	*Response_ContractGetRecordsResponse
	//	*Response_CryptogetAccountBalance
	//	*Response_CryptoGetAccountRecords
	//	*Response_CryptoGetInfo
	//	*Response_CryptoGetClaim
	//	*Response_CryptoGetProxyStakers
	//	*Response_FileGetContents
	//	*Response_FileGetInfo
	//	*Response_TransactionGetReceipt
	//	*Response_TransactionGetRecord
	Response             isResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c4c7055b2f7bc3f, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type isResponse_Response interface {
	isResponse_Response()
}

type Response_GetByKey struct {
	GetByKey *GetByKeyResponse `protobuf:"bytes,1,opt,name=getByKey,proto3,oneof"`
}

type Response_GetBySolidityID struct {
	GetBySolidityID *GetBySolidityIDResponse `protobuf:"bytes,2,opt,name=getBySolidityID,proto3,oneof"`
}

type Response_ContractCallLocal struct {
	ContractCallLocal *ContractCallLocalResponse `protobuf:"bytes,3,opt,name=contractCallLocal,proto3,oneof"`
}

type Response_ContractGetBytecodeResponse struct {
	ContractGetBytecodeResponse *ContractGetBytecodeResponse `protobuf:"bytes,5,opt,name=contractGetBytecodeResponse,proto3,oneof"`
}

type Response_ContractGetInfo struct {
	ContractGetInfo *ContractGetInfoResponse `protobuf:"bytes,4,opt,name=contractGetInfo,proto3,oneof"`
}

type Response_ContractGetRecordsResponse struct {
	ContractGetRecordsResponse *ContractGetRecordsResponse `protobuf:"bytes,6,opt,name=contractGetRecordsResponse,proto3,oneof"`
}

type Response_CryptogetAccountBalance struct {
	CryptogetAccountBalance *CryptoGetAccountBalanceResponse `protobuf:"bytes,7,opt,name=cryptogetAccountBalance,proto3,oneof"`
}

type Response_CryptoGetAccountRecords struct {
	CryptoGetAccountRecords *CryptoGetAccountRecordsResponse `protobuf:"bytes,8,opt,name=cryptoGetAccountRecords,proto3,oneof"`
}

type Response_CryptoGetInfo struct {
	CryptoGetInfo *CryptoGetInfoResponse `protobuf:"bytes,9,opt,name=cryptoGetInfo,proto3,oneof"`
}

type Response_CryptoGetClaim struct {
	CryptoGetClaim *CryptoGetClaimResponse `protobuf:"bytes,10,opt,name=cryptoGetClaim,proto3,oneof"`
}

type Response_CryptoGetProxyStakers struct {
	CryptoGetProxyStakers *CryptoGetStakersResponse `protobuf:"bytes,11,opt,name=cryptoGetProxyStakers,proto3,oneof"`
}

type Response_FileGetContents struct {
	FileGetContents *FileGetContentsResponse `protobuf:"bytes,12,opt,name=fileGetContents,proto3,oneof"`
}

type Response_FileGetInfo struct {
	FileGetInfo *FileGetInfoResponse `protobuf:"bytes,13,opt,name=fileGetInfo,proto3,oneof"`
}

type Response_TransactionGetReceipt struct {
	TransactionGetReceipt *TransactionGetReceiptResponse `protobuf:"bytes,14,opt,name=transactionGetReceipt,proto3,oneof"`
}

type Response_TransactionGetRecord struct {
	TransactionGetRecord *TransactionGetRecordResponse `protobuf:"bytes,15,opt,name=transactionGetRecord,proto3,oneof"`
}

func (*Response_GetByKey) isResponse_Response() {}

func (*Response_GetBySolidityID) isResponse_Response() {}

func (*Response_ContractCallLocal) isResponse_Response() {}

func (*Response_ContractGetBytecodeResponse) isResponse_Response() {}

func (*Response_ContractGetInfo) isResponse_Response() {}

func (*Response_ContractGetRecordsResponse) isResponse_Response() {}

func (*Response_CryptogetAccountBalance) isResponse_Response() {}

func (*Response_CryptoGetAccountRecords) isResponse_Response() {}

func (*Response_CryptoGetInfo) isResponse_Response() {}

func (*Response_CryptoGetClaim) isResponse_Response() {}

func (*Response_CryptoGetProxyStakers) isResponse_Response() {}

func (*Response_FileGetContents) isResponse_Response() {}

func (*Response_FileGetInfo) isResponse_Response() {}

func (*Response_TransactionGetReceipt) isResponse_Response() {}

func (*Response_TransactionGetRecord) isResponse_Response() {}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetGetByKey() *GetByKeyResponse {
	if x, ok := m.GetResponse().(*Response_GetByKey); ok {
		return x.GetByKey
	}
	return nil
}

func (m *Response) GetGetBySolidityID() *GetBySolidityIDResponse {
	if x, ok := m.GetResponse().(*Response_GetBySolidityID); ok {
		return x.GetBySolidityID
	}
	return nil
}

func (m *Response) GetContractCallLocal() *ContractCallLocalResponse {
	if x, ok := m.GetResponse().(*Response_ContractCallLocal); ok {
		return x.ContractCallLocal
	}
	return nil
}

func (m *Response) GetContractGetBytecodeResponse() *ContractGetBytecodeResponse {
	if x, ok := m.GetResponse().(*Response_ContractGetBytecodeResponse); ok {
		return x.ContractGetBytecodeResponse
	}
	return nil
}

func (m *Response) GetContractGetInfo() *ContractGetInfoResponse {
	if x, ok := m.GetResponse().(*Response_ContractGetInfo); ok {
		return x.ContractGetInfo
	}
	return nil
}

func (m *Response) GetContractGetRecordsResponse() *ContractGetRecordsResponse {
	if x, ok := m.GetResponse().(*Response_ContractGetRecordsResponse); ok {
		return x.ContractGetRecordsResponse
	}
	return nil
}

func (m *Response) GetCryptogetAccountBalance() *CryptoGetAccountBalanceResponse {
	if x, ok := m.GetResponse().(*Response_CryptogetAccountBalance); ok {
		return x.CryptogetAccountBalance
	}
	return nil
}

func (m *Response) GetCryptoGetAccountRecords() *CryptoGetAccountRecordsResponse {
	if x, ok := m.GetResponse().(*Response_CryptoGetAccountRecords); ok {
		return x.CryptoGetAccountRecords
	}
	return nil
}

func (m *Response) GetCryptoGetInfo() *CryptoGetInfoResponse {
	if x, ok := m.GetResponse().(*Response_CryptoGetInfo); ok {
		return x.CryptoGetInfo
	}
	return nil
}

func (m *Response) GetCryptoGetClaim() *CryptoGetClaimResponse {
	if x, ok := m.GetResponse().(*Response_CryptoGetClaim); ok {
		return x.CryptoGetClaim
	}
	return nil
}

func (m *Response) GetCryptoGetProxyStakers() *CryptoGetStakersResponse {
	if x, ok := m.GetResponse().(*Response_CryptoGetProxyStakers); ok {
		return x.CryptoGetProxyStakers
	}
	return nil
}

func (m *Response) GetFileGetContents() *FileGetContentsResponse {
	if x, ok := m.GetResponse().(*Response_FileGetContents); ok {
		return x.FileGetContents
	}
	return nil
}

func (m *Response) GetFileGetInfo() *FileGetInfoResponse {
	if x, ok := m.GetResponse().(*Response_FileGetInfo); ok {
		return x.FileGetInfo
	}
	return nil
}

func (m *Response) GetTransactionGetReceipt() *TransactionGetReceiptResponse {
	if x, ok := m.GetResponse().(*Response_TransactionGetReceipt); ok {
		return x.TransactionGetReceipt
	}
	return nil
}

func (m *Response) GetTransactionGetRecord() *TransactionGetRecordResponse {
	if x, ok := m.GetResponse().(*Response_TransactionGetRecord); ok {
		return x.TransactionGetRecord
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Response_OneofMarshaler, _Response_OneofUnmarshaler, _Response_OneofSizer, []interface{}{
		(*Response_GetByKey)(nil),
		(*Response_GetBySolidityID)(nil),
		(*Response_ContractCallLocal)(nil),
		(*Response_ContractGetBytecodeResponse)(nil),
		(*Response_ContractGetInfo)(nil),
		(*Response_ContractGetRecordsResponse)(nil),
		(*Response_CryptogetAccountBalance)(nil),
		(*Response_CryptoGetAccountRecords)(nil),
		(*Response_CryptoGetInfo)(nil),
		(*Response_CryptoGetClaim)(nil),
		(*Response_CryptoGetProxyStakers)(nil),
		(*Response_FileGetContents)(nil),
		(*Response_FileGetInfo)(nil),
		(*Response_TransactionGetReceipt)(nil),
		(*Response_TransactionGetRecord)(nil),
	}
}

func _Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Response)
	// response
	switch x := m.Response.(type) {
	case *Response_GetByKey:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetByKey); err != nil {
			return err
		}
	case *Response_GetBySolidityID:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetBySolidityID); err != nil {
			return err
		}
	case *Response_ContractCallLocal:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractCallLocal); err != nil {
			return err
		}
	case *Response_ContractGetBytecodeResponse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractGetBytecodeResponse); err != nil {
			return err
		}
	case *Response_ContractGetInfo:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractGetInfo); err != nil {
			return err
		}
	case *Response_ContractGetRecordsResponse:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractGetRecordsResponse); err != nil {
			return err
		}
	case *Response_CryptogetAccountBalance:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptogetAccountBalance); err != nil {
			return err
		}
	case *Response_CryptoGetAccountRecords:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptoGetAccountRecords); err != nil {
			return err
		}
	case *Response_CryptoGetInfo:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptoGetInfo); err != nil {
			return err
		}
	case *Response_CryptoGetClaim:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptoGetClaim); err != nil {
			return err
		}
	case *Response_CryptoGetProxyStakers:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptoGetProxyStakers); err != nil {
			return err
		}
	case *Response_FileGetContents:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileGetContents); err != nil {
			return err
		}
	case *Response_FileGetInfo:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileGetInfo); err != nil {
			return err
		}
	case *Response_TransactionGetReceipt:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransactionGetReceipt); err != nil {
			return err
		}
	case *Response_TransactionGetRecord:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransactionGetRecord); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Response.Response has unexpected type %T", x)
	}
	return nil
}

func _Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Response)
	switch tag {
	case 1: // response.getByKey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetByKeyResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_GetByKey{msg}
		return true, err
	case 2: // response.getBySolidityID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetBySolidityIDResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_GetBySolidityID{msg}
		return true, err
	case 3: // response.contractCallLocal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractCallLocalResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_ContractCallLocal{msg}
		return true, err
	case 5: // response.contractGetBytecodeResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractGetBytecodeResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_ContractGetBytecodeResponse{msg}
		return true, err
	case 4: // response.contractGetInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractGetInfoResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_ContractGetInfo{msg}
		return true, err
	case 6: // response.contractGetRecordsResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractGetRecordsResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_ContractGetRecordsResponse{msg}
		return true, err
	case 7: // response.cryptogetAccountBalance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoGetAccountBalanceResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_CryptogetAccountBalance{msg}
		return true, err
	case 8: // response.cryptoGetAccountRecords
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoGetAccountRecordsResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_CryptoGetAccountRecords{msg}
		return true, err
	case 9: // response.cryptoGetInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoGetInfoResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_CryptoGetInfo{msg}
		return true, err
	case 10: // response.cryptoGetClaim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoGetClaimResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_CryptoGetClaim{msg}
		return true, err
	case 11: // response.cryptoGetProxyStakers
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoGetStakersResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_CryptoGetProxyStakers{msg}
		return true, err
	case 12: // response.fileGetContents
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileGetContentsResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_FileGetContents{msg}
		return true, err
	case 13: // response.fileGetInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileGetInfoResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_FileGetInfo{msg}
		return true, err
	case 14: // response.transactionGetReceipt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionGetReceiptResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_TransactionGetReceipt{msg}
		return true, err
	case 15: // response.transactionGetRecord
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionGetRecordResponse)
		err := b.DecodeMessage(msg)
		m.Response = &Response_TransactionGetRecord{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Response)
	// response
	switch x := m.Response.(type) {
	case *Response_GetByKey:
		s := proto.Size(x.GetByKey)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_GetBySolidityID:
		s := proto.Size(x.GetBySolidityID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_ContractCallLocal:
		s := proto.Size(x.ContractCallLocal)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_ContractGetBytecodeResponse:
		s := proto.Size(x.ContractGetBytecodeResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_ContractGetInfo:
		s := proto.Size(x.ContractGetInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_ContractGetRecordsResponse:
		s := proto.Size(x.ContractGetRecordsResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CryptogetAccountBalance:
		s := proto.Size(x.CryptogetAccountBalance)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CryptoGetAccountRecords:
		s := proto.Size(x.CryptoGetAccountRecords)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CryptoGetInfo:
		s := proto.Size(x.CryptoGetInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CryptoGetClaim:
		s := proto.Size(x.CryptoGetClaim)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CryptoGetProxyStakers:
		s := proto.Size(x.CryptoGetProxyStakers)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_FileGetContents:
		s := proto.Size(x.FileGetContents)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_FileGetInfo:
		s := proto.Size(x.FileGetInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_TransactionGetReceipt:
		s := proto.Size(x.TransactionGetReceipt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_TransactionGetRecord:
		s := proto.Size(x.TransactionGetRecord)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("Response.proto", fileDescriptor_2c4c7055b2f7bc3f) }

var fileDescriptor_2c4c7055b2f7bc3f = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x5b, 0x60, 0xa3, 0x7b, 0x65, 0x9d, 0x66, 0x56, 0x6a, 0x3a, 0x06, 0xa3, 0xa0, 0x89,
	0x53, 0x0f, 0x20, 0xae, 0x48, 0xb4, 0x13, 0xed, 0x80, 0x43, 0x95, 0x21, 0x21, 0x24, 0x2e, 0x9e,
	0xe3, 0xb6, 0x81, 0xcc, 0x8e, 0x1c, 0x83, 0xc8, 0x9f, 0xce, 0x6d, 0xaa, 0x6b, 0xa7, 0xb6, 0x93,
	0xf6, 0x54, 0xe5, 0x7d, 0xef, 0xfb, 0xbd, 0xfa, 0xf9, 0xf9, 0x41, 0x27, 0x62, 0x79, 0x26, 0x78,
	0xce, 0x86, 0x99, 0x14, 0x4a, 0xa0, 0x3d, 0xfd, 0xd3, 0xef, 0x4c, 0x98, 0x1a, 0x15, 0x5f, 0x58,
	0xb1, 0x0e, 0xf7, 0xbb, 0xfa, 0xfb, 0x5a, 0xa4, 0x49, 0x9c, 0xa8, 0xe2, 0xea, 0xd2, 0x84, 0x7b,
	0x63, 0xc1, 0x95, 0x24, 0x54, 0x8d, 0x49, 0x9a, 0x7e, 0x15, 0x94, 0xa4, 0x46, 0x78, 0x6a, 0x05,
	0xed, 0x53, 0x8c, 0x8a, 0x98, 0x59, 0x94, 0x23, 0x5d, 0xf1, 0xb9, 0x30, 0x61, 0xec, 0x84, 0x23,
	0x46, 0x85, 0x8c, 0x73, 0xa3, 0x9c, 0x8d, 0x65, 0x91, 0x29, 0x31, 0x61, 0xea, 0x23, 0xa5, 0xe2,
	0x0f, 0x57, 0x23, 0x92, 0x12, 0x4e, 0xd9, 0x36, 0xd9, 0x77, 0x3f, 0x2e, 0x65, 0xa7, 0xd8, 0x49,
	0x19, 0x1c, 0xa7, 0x24, 0xb9, 0x35, 0xd1, 0x27, 0x65, 0xf4, 0x5a, 0x91, 0xdf, 0x4c, 0x5a, 0x44,
	0xf7, 0x53, 0x92, 0xb2, 0x55, 0xae, 0xe0, 0x8a, 0x71, 0x65, 0xc3, 0xc7, 0x26, 0xec, 0x70, 0x4f,
	0xbf, 0x49, 0xc2, 0x73, 0x42, 0x55, 0x22, 0xf8, 0xfa, 0x1c, 0x2c, 0xc9, 0x94, 0x11, 0xfb, 0x15,
	0x51, 0xc8, 0x78, 0xad, 0x0d, 0xfe, 0x1f, 0x40, 0xcb, 0xde, 0x04, 0x7a, 0x0f, 0xad, 0x85, 0x69,
	0x3f, 0x6e, 0x9e, 0x37, 0xdf, 0xb4, 0xdf, 0xf6, 0xd6, 0x69, 0x43, 0x7b, 0x2b, 0x36, 0x75, 0xda,
	0x88, 0xca, 0x54, 0xf4, 0x19, 0x8e, 0x16, 0xfe, 0x2d, 0xe1, 0x7b, 0xda, 0xfd, 0xdc, 0x75, 0x6f,
	0x54, 0x07, 0x12, 0x1a, 0xd1, 0x0c, 0x8e, 0x69, 0x78, 0xb5, 0xf8, 0xbe, 0xa6, 0x9d, 0x1b, 0x5a,
	0xe5, 0xea, 0x1d, 0x5e, 0xd5, 0x8c, 0xe6, 0x70, 0x4a, 0xab, 0x33, 0x61, 0x3d, 0x78, 0x4f, 0xb3,
	0x07, 0x01, 0xbb, 0x26, 0x73, 0xda, 0x88, 0x76, 0x81, 0x56, 0x5d, 0xa0, 0xfe, 0x80, 0xe1, 0x07,
	0x5e, 0x17, 0x82, 0xf1, 0x73, 0xbb, 0x10, 0x18, 0x11, 0x85, 0x3e, 0xad, 0x4c, 0x65, 0xf9, 0x97,
	0xf7, 0x35, 0xf6, 0x65, 0x15, 0x1b, 0x24, 0x4e, 0x1b, 0xd1, 0x0e, 0x0c, 0xba, 0x81, 0x1e, 0xd5,
	0x73, 0xb7, 0x08, 0x07, 0x1c, 0x3f, 0xd4, 0x15, 0x2e, 0x6c, 0x85, 0xfa, 0x67, 0xe0, 0x94, 0xd9,
	0x06, 0xda, 0xd4, 0xa8, 0xbc, 0x12, 0xdc, 0xda, 0x59, 0xa3, 0x7a, 0x94, 0x6d, 0x20, 0x74, 0x09,
	0x87, 0xd4, 0x7d, 0x6a, 0xf8, 0x40, 0x93, 0x9f, 0x85, 0xe4, 0xa0, 0xe9, 0xbe, 0x09, 0x4d, 0xa0,
	0x43, 0xbd, 0xb7, 0x89, 0x41, 0x63, 0xce, 0x42, 0x8c, 0x16, 0x1d, 0x4e, 0x60, 0x43, 0xdf, 0xa1,
	0x5b, 0x46, 0x66, 0x52, 0xfc, 0x2b, 0xcc, 0x9b, 0xc6, 0x6d, 0xcd, 0x7b, 0x11, 0xf2, 0x8c, 0xec,
	0x10, 0xeb, 0xfd, 0xab, 0x01, 0x9b, 0xfb, 0xfb, 0x00, 0x3f, 0xf2, 0x06, 0x2c, 0xd8, 0x16, 0xee,
	0x80, 0x05, 0x46, 0xf4, 0x01, 0xda, 0xf3, 0xcd, 0x12, 0xc1, 0x87, 0x9a, 0xd3, 0xf7, 0x39, 0x41,
	0xbf, 0x5c, 0x03, 0xfa, 0x09, 0x5d, 0x55, 0xb7, 0x71, 0x70, 0x47, 0x93, 0x5e, 0x1b, 0x52, 0xed,
	0x56, 0x72, 0x4f, 0x5a, 0x0b, 0x41, 0x3f, 0xe0, 0x44, 0xd5, 0xac, 0x2c, 0x7c, 0xa4, 0xe1, 0xaf,
	0xb6, 0xc1, 0x85, 0x8c, 0x1d, 0x76, 0x2d, 0x62, 0x04, 0xd0, 0x92, 0x26, 0x67, 0x74, 0x01, 0x03,
	0x2a, 0x6e, 0x87, 0x4b, 0x16, 0x33, 0x49, 0x96, 0x24, 0x5f, 0x2e, 0x24, 0xc9, 0x96, 0x43, 0x92,
	0x25, 0xa6, 0xc2, 0x2f, 0xf2, 0x97, 0xcc, 0x9a, 0x37, 0xfb, 0xfa, 0xeb, 0xdd, 0x5d, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x14, 0x40, 0xb7, 0xb3, 0x06, 0x00, 0x00,
}
