// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wallet.proto

package rpcpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ListTransactionsRequest struct {
	Addr   string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *ListTransactionsRequest) Reset()         { *m = ListTransactionsRequest{} }
func (m *ListTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsRequest) ProtoMessage()    {}
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_fe907f042cebfb95, []int{0}
}
func (m *ListTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTransactionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ListTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsRequest.Merge(dst, src)
}
func (m *ListTransactionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsRequest proto.InternalMessageInfo

func (m *ListTransactionsRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ListTransactionsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListTransactionsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListTransactionsResponse struct {
	Code         int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count        uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Transactions []*MsgTx `protobuf:"bytes,4,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *ListTransactionsResponse) Reset()         { *m = ListTransactionsResponse{} }
func (m *ListTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsResponse) ProtoMessage()    {}
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_fe907f042cebfb95, []int{1}
}
func (m *ListTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTransactionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ListTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsResponse.Merge(dst, src)
}
func (m *ListTransactionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsResponse proto.InternalMessageInfo

func (m *ListTransactionsResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListTransactionsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ListTransactionsResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListTransactionsResponse) GetTransactions() []*MsgTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Transaction struct {
	TxHash   string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	RawBytes []byte `protobuf:"bytes,2,opt,name=raw_bytes,json=rawBytes,proto3" json:"raw_bytes,omitempty"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_fe907f042cebfb95, []int{2}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Transaction) GetRawBytes() []byte {
	if m != nil {
		return m.RawBytes
	}
	return nil
}

type GetTransactionCountRequest struct {
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *GetTransactionCountRequest) Reset()         { *m = GetTransactionCountRequest{} }
func (m *GetTransactionCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionCountRequest) ProtoMessage()    {}
func (*GetTransactionCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_fe907f042cebfb95, []int{3}
}
func (m *GetTransactionCountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTransactionCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTransactionCountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetTransactionCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionCountRequest.Merge(dst, src)
}
func (m *GetTransactionCountRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTransactionCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionCountRequest proto.InternalMessageInfo

func (m *GetTransactionCountRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type GetTransactionCountResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count   uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *GetTransactionCountResponse) Reset()         { *m = GetTransactionCountResponse{} }
func (m *GetTransactionCountResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionCountResponse) ProtoMessage()    {}
func (*GetTransactionCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_fe907f042cebfb95, []int{4}
}
func (m *GetTransactionCountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTransactionCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTransactionCountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetTransactionCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionCountResponse.Merge(dst, src)
}
func (m *GetTransactionCountResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTransactionCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionCountResponse proto.InternalMessageInfo

func (m *GetTransactionCountResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetTransactionCountResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetTransactionCountResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*ListTransactionsRequest)(nil), "rpcpb.ListTransactionsRequest")
	proto.RegisterType((*ListTransactionsResponse)(nil), "rpcpb.ListTransactionsResponse")
	proto.RegisterType((*Transaction)(nil), "rpcpb.Transaction")
	proto.RegisterType((*GetTransactionCountRequest)(nil), "rpcpb.GetTransactionCountRequest")
	proto.RegisterType((*GetTransactionCountResponse)(nil), "rpcpb.GetTransactionCountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletCommandClient is the client API for WalletCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletCommandClient interface {
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	GetTransactionCount(ctx context.Context, in *GetTransactionCountRequest, opts ...grpc.CallOption) (*GetTransactionCountResponse, error)
}

type walletCommandClient struct {
	cc *grpc.ClientConn
}

func NewWalletCommandClient(cc *grpc.ClientConn) WalletCommandClient {
	return &walletCommandClient{cc}
}

func (c *walletCommandClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.WalletCommand/ListTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletCommandClient) GetTransactionCount(ctx context.Context, in *GetTransactionCountRequest, opts ...grpc.CallOption) (*GetTransactionCountResponse, error) {
	out := new(GetTransactionCountResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.WalletCommand/GetTransactionCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletCommandServer is the server API for WalletCommand service.
type WalletCommandServer interface {
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	GetTransactionCount(context.Context, *GetTransactionCountRequest) (*GetTransactionCountResponse, error)
}

func RegisterWalletCommandServer(s *grpc.Server, srv WalletCommandServer) {
	s.RegisterService(&_WalletCommand_serviceDesc, srv)
}

func _WalletCommand_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletCommandServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.WalletCommand/ListTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletCommandServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletCommand_GetTransactionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletCommandServer).GetTransactionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.WalletCommand/GetTransactionCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletCommandServer).GetTransactionCount(ctx, req.(*GetTransactionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.WalletCommand",
	HandlerType: (*WalletCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransactions",
			Handler:    _WalletCommand_ListTransactions_Handler,
		},
		{
			MethodName: "GetTransactionCount",
			Handler:    _WalletCommand_GetTransactionCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}

func (m *ListTransactionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTransactionsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if m.Offset != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Offset))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Limit))
	}
	return i, nil
}

func (m *ListTransactionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTransactionsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Count != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Count))
	}
	if len(m.Transactions) > 0 {
		for _, msg := range m.Transactions {
			dAtA[i] = 0x22
			i++
			i = encodeVarintWallet(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if len(m.RawBytes) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.RawBytes)))
		i += copy(dAtA[i:], m.RawBytes)
	}
	return i, nil
}

func (m *GetTransactionCountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTransactionCountRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	return i, nil
}

func (m *GetTransactionCountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTransactionCountResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Count != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func encodeVarintWallet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ListTransactionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovWallet(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovWallet(uint64(m.Limit))
	}
	return n
}

func (m *ListTransactionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovWallet(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovWallet(uint64(m.Count))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovWallet(uint64(l))
		}
	}
	return n
}

func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.RawBytes)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	return n
}

func (m *GetTransactionCountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	return n
}

func (m *GetTransactionCountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovWallet(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovWallet(uint64(m.Count))
	}
	return n
}

func sovWallet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWallet(x uint64) (n int) {
	return sovWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListTransactionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListTransactionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTransactionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListTransactionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListTransactionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTransactionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &MsgTx{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawBytes = append(m.RawBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.RawBytes == nil {
				m.RawBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTransactionCountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTransactionCountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTransactionCountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTransactionCountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTransactionCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTransactionCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWallet
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWallet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWallet
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWallet(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWallet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWallet   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_wallet_fe907f042cebfb95) }

var fileDescriptor_wallet_fe907f042cebfb95 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x1c, 0xc5, 0x73, 0x69, 0x93, 0x92, 0x7f, 0x5d, 0x09, 0x1d, 0x88, 0x9e, 0x1c, 0x30, 0xe6, 0x90,
	0x50, 0xd4, 0x21, 0x2e, 0x65, 0xeb, 0xd8, 0x0c, 0x30, 0xc0, 0x62, 0x55, 0x62, 0x60, 0xa8, 0x2e,
	0xf6, 0xd5, 0xb1, 0x64, 0xdf, 0x19, 0xdf, 0x05, 0x9b, 0x15, 0xf1, 0x01, 0x22, 0xf1, 0xa5, 0x18,
	0x23, 0xb1, 0x30, 0xa2, 0x84, 0x0f, 0x82, 0x7c, 0x76, 0x90, 0x81, 0x24, 0x13, 0xdb, 0x3d, 0xfb,
	0xf9, 0xdd, 0xfb, 0xff, 0xfc, 0x07, 0xab, 0x60, 0x49, 0xc2, 0xf5, 0x38, 0xcb, 0xa5, 0x96, 0xb8,
	0x97, 0x67, 0x41, 0x36, 0xb5, 0x1f, 0x46, 0x52, 0x46, 0x09, 0xf7, 0x58, 0x16, 0x7b, 0x4c, 0x08,
	0xa9, 0x99, 0x8e, 0xa5, 0x50, 0xb5, 0xc9, 0xb6, 0x02, 0x99, 0xa6, 0x52, 0xd4, 0x8a, 0xbe, 0x83,
	0xd3, 0xd7, 0xb1, 0xd2, 0xd7, 0x39, 0x13, 0x8a, 0x05, 0xc6, 0xe7, 0xf3, 0xf7, 0x73, 0xae, 0x34,
	0xc6, 0x70, 0xc8, 0xc2, 0x30, 0x27, 0xc8, 0x45, 0xa3, 0x81, 0x6f, 0xce, 0xf8, 0x01, 0xf4, 0xe5,
	0xed, 0xad, 0xe2, 0x9a, 0x74, 0x5d, 0x34, 0x3a, 0xf1, 0x1b, 0x85, 0xef, 0x43, 0x2f, 0x89, 0xd3,
	0x58, 0x93, 0x03, 0xf3, 0xb8, 0x16, 0x74, 0x81, 0x80, 0xfc, 0x9b, 0xae, 0x32, 0x29, 0x14, 0xaf,
	0xe2, 0x03, 0x19, 0x72, 0x13, 0xdf, 0xf3, 0xcd, 0x19, 0x13, 0x38, 0x4a, 0xb9, 0x52, 0x2c, 0xe2,
	0x26, 0x7f, 0xe0, 0x6f, 0x64, 0x75, 0x41, 0x20, 0xe7, 0xe2, 0xf7, 0x05, 0x46, 0xe0, 0x73, 0xb0,
	0x74, 0x2b, 0x9b, 0x1c, 0xba, 0x07, 0xa3, 0xe3, 0x0b, 0x6b, 0x6c, 0x38, 0x8c, 0xdf, 0xa8, 0xe8,
	0xba, 0xf4, 0xff, 0x70, 0xd0, 0x09, 0x1c, 0xb7, 0xda, 0xe0, 0x53, 0x38, 0xd2, 0xe5, 0xcd, 0x8c,
	0xa9, 0x59, 0x33, 0x66, 0x5f, 0x97, 0xaf, 0x98, 0x9a, 0xe1, 0x21, 0x0c, 0x72, 0x56, 0xdc, 0x4c,
	0x3f, 0x6a, 0xae, 0x4c, 0x17, 0xcb, 0xbf, 0x93, 0xb3, 0xe2, 0xaa, 0xd2, 0xf4, 0x1c, 0xec, 0x97,
	0xbc, 0x3d, 0xd5, 0xa4, 0x6a, 0xb3, 0x87, 0x1b, 0x65, 0x30, 0xdc, 0xfa, 0xc5, 0xff, 0x63, 0x71,
	0xb1, 0xe8, 0xc2, 0xc9, 0x5b, 0xb3, 0x0d, 0x13, 0x99, 0xa6, 0x4c, 0x84, 0xb8, 0x84, 0xbb, 0x7f,
	0xd3, 0xc7, 0x4e, 0xc3, 0x66, 0xc7, 0x4f, 0xb7, 0x1f, 0xef, 0x7c, 0x5f, 0x57, 0xa5, 0x4f, 0x3f,
	0x7d, 0xfb, 0xf9, 0xa5, 0xfb, 0x88, 0x12, 0xef, 0xc3, 0x73, 0xaf, 0x48, 0xb4, 0x97, 0xc4, 0x4a,
	0xb7, 0x11, 0x5f, 0xa2, 0x33, 0xfc, 0x19, 0xc1, 0xbd, 0x2d, 0xf3, 0xe2, 0x27, 0x4d, 0xfa, 0x6e,
	0x7a, 0x36, 0xdd, 0x67, 0x69, 0x3a, 0x3c, 0x33, 0x1d, 0x5c, 0x3a, 0xdc, 0x74, 0x88, 0x78, 0xbb,
	0x82, 0xe1, 0x71, 0x89, 0xce, 0xae, 0xc8, 0xd7, 0x95, 0x83, 0x96, 0x2b, 0x07, 0xfd, 0x58, 0x39,
	0x68, 0xb1, 0x76, 0x3a, 0xcb, 0xb5, 0xd3, 0xf9, 0xbe, 0x76, 0x3a, 0xd3, 0xbe, 0xd9, 0xfe, 0x17,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x86, 0x89, 0x17, 0x40, 0x03, 0x00, 0x00,
}
