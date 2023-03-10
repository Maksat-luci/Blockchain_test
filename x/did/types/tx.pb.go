// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spider/did/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateDID struct {
	Did                  string       `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Document             *DIDDocument `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	VerificationMethodId string       `protobuf:"bytes,3,opt,name=verificationMethodId,proto3" json:"verificationMethodId,omitempty"`
	Signature            []byte       `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	FromAddress          string       `protobuf:"bytes,5,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
}

func (m *MsgCreateDID) Reset()         { *m = MsgCreateDID{} }
func (m *MsgCreateDID) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDID) ProtoMessage()    {}
func (*MsgCreateDID) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb2559689e71f8ff, []int{0}
}
func (m *MsgCreateDID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDID.Merge(m, src)
}
func (m *MsgCreateDID) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDID) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDID.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDID proto.InternalMessageInfo

func (m *MsgCreateDID) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *MsgCreateDID) GetDocument() *DIDDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *MsgCreateDID) GetVerificationMethodId() string {
	if m != nil {
		return m.VerificationMethodId
	}
	return ""
}

func (m *MsgCreateDID) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *MsgCreateDID) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type MsgCreateDIDResponse struct {
}

func (m *MsgCreateDIDResponse) Reset()         { *m = MsgCreateDIDResponse{} }
func (m *MsgCreateDIDResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDIDResponse) ProtoMessage()    {}
func (*MsgCreateDIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb2559689e71f8ff, []int{1}
}
func (m *MsgCreateDIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDIDResponse.Merge(m, src)
}
func (m *MsgCreateDIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDIDResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateDID)(nil), "spider.did.MsgCreateDID")
	proto.RegisterType((*MsgCreateDIDResponse)(nil), "spider.did.MsgCreateDIDResponse")
}

func init() { proto.RegisterFile("spider/did/tx.proto", fileDescriptor_cb2559689e71f8ff) }

var fileDescriptor_cb2559689e71f8ff = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xad, 0xbf, 0x7c, 0x20, 0xe2, 0x76, 0x40, 0xa6, 0x02, 0xab, 0x02, 0x2b, 0xea, 0x94, 0x01,
	0xa5, 0x52, 0xfa, 0x04, 0x40, 0x24, 0x94, 0x21, 0x0c, 0x19, 0x59, 0x50, 0xe8, 0xbd, 0x0d, 0x1e,
	0x1a, 0x47, 0xb6, 0x8b, 0xca, 0x5b, 0xf0, 0x58, 0xb0, 0x75, 0x64, 0x44, 0xc9, 0x8b, 0x20, 0x92,
	0x96, 0x64, 0x80, 0xcd, 0x3a, 0x7f, 0xd6, 0x39, 0x97, 0x9e, 0x98, 0x52, 0x02, 0xea, 0x19, 0x48,
	0x98, 0xd9, 0x4d, 0x50, 0x6a, 0x65, 0x15, 0xa3, 0x2d, 0x18, 0x80, 0x84, 0xc9, 0x45, 0x4f, 0x00,
	0x12, 0x1e, 0x40, 0x2d, 0xd6, 0x2b, 0x2c, 0x6c, 0x2b, 0x9d, 0xbe, 0x13, 0x3a, 0x4a, 0x4c, 0x7e,
	0xa3, 0x31, 0xb3, 0x18, 0xc5, 0x11, 0x3b, 0xa6, 0x0e, 0x48, 0xe0, 0xc4, 0x23, 0xbe, 0x9b, 0x7e,
	0x3f, 0xd9, 0x9c, 0x1e, 0xed, 0x4d, 0xfc, 0x9f, 0x47, 0xfc, 0x61, 0x78, 0x16, 0x74, 0x1f, 0x04,
	0x51, 0x1c, 0x45, 0x3b, 0x3a, 0xfd, 0x11, 0xb2, 0x90, 0x8e, 0x9f, 0x51, 0xcb, 0xa5, 0x5c, 0x64,
	0x56, 0xaa, 0x22, 0x41, 0xfb, 0xa4, 0x20, 0x06, 0xee, 0x34, 0xb9, 0xbf, 0x72, 0xec, 0x9c, 0xba,
	0x46, 0xe6, 0x45, 0x66, 0xd7, 0x1a, 0xf9, 0x7f, 0x8f, 0xf8, 0xa3, 0xb4, 0x03, 0x98, 0x47, 0x87,
	0x4b, 0xad, 0x56, 0x57, 0x00, 0x1a, 0x8d, 0xe1, 0x07, 0x4d, 0x50, 0x1f, 0x9a, 0x9e, 0xd2, 0x71,
	0xbf, 0x4a, 0x8a, 0xa6, 0x54, 0x85, 0xc1, 0xf0, 0x8e, 0x3a, 0x89, 0xc9, 0xd9, 0x2d, 0x75, 0xbb,
	0x9a, 0xbc, 0x5f, 0xa1, 0xef, 0x9a, 0x78, 0x7f, 0x31, 0xfb, 0xbc, 0xeb, 0xcb, 0xb7, 0x4a, 0x90,
	0x6d, 0x25, 0xc8, 0x67, 0x25, 0xc8, 0x6b, 0x2d, 0x06, 0xdb, 0x5a, 0x0c, 0x3e, 0x6a, 0x31, 0xb8,
	0x67, 0xbb, 0xb1, 0x37, 0xed, 0x3d, 0x5e, 0x4a, 0x34, 0x8f, 0x87, 0xcd, 0xd0, 0xf3, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0x9c, 0xaa, 0x3d, 0xaa, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateDID(ctx context.Context, in *MsgCreateDID, opts ...grpc.CallOption) (*MsgCreateDIDResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDID(ctx context.Context, in *MsgCreateDID, opts ...grpc.CallOption) (*MsgCreateDIDResponse, error) {
	out := new(MsgCreateDIDResponse)
	err := c.cc.Invoke(ctx, "/spider.did.Msg/CreateDID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateDID(context.Context, *MsgCreateDID) (*MsgCreateDIDResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateDID(ctx context.Context, req *MsgCreateDID) (*MsgCreateDIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDID not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spider.did.Msg/CreateDID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDID(ctx, req.(*MsgCreateDID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spider.did.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDID",
			Handler:    _Msg_CreateDID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spider/did/tx.proto",
}

func (m *MsgCreateDID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VerificationMethodId) > 0 {
		i -= len(m.VerificationMethodId)
		copy(dAtA[i:], m.VerificationMethodId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VerificationMethodId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Document != nil {
		{
			size, err := m.Document.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateDIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateDID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Document != nil {
		l = m.Document.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VerificationMethodId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateDIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateDID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateDID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Document == nil {
				m.Document = &DIDDocument{}
			}
			if err := m.Document.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethodId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethodId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateDIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateDIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
