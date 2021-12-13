// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: application.proto

package types

import (
	context "context"
	fmt "fmt"
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

type ApplicationInfoResponse struct {
	Info *ApplicationInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (m *ApplicationInfoResponse) Reset()         { *m = ApplicationInfoResponse{} }
func (m *ApplicationInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ApplicationInfoResponse) ProtoMessage()    {}
func (*ApplicationInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc846aced8fe6ea6, []int{0}
}
func (m *ApplicationInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplicationInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInfoResponse.Merge(m, src)
}
func (m *ApplicationInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInfoResponse proto.InternalMessageInfo

func (m *ApplicationInfoResponse) GetInfo() *ApplicationInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ApplicationInfo struct {
	Checksum string `protobuf:"bytes,1,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	Commit   string `protobuf:"bytes,3,opt,name=Commit,proto3" json:"Commit,omitempty"`
}

func (m *ApplicationInfo) Reset()         { *m = ApplicationInfo{} }
func (m *ApplicationInfo) String() string { return proto.CompactTextString(m) }
func (*ApplicationInfo) ProtoMessage()    {}
func (*ApplicationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc846aced8fe6ea6, []int{1}
}
func (m *ApplicationInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplicationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInfo.Merge(m, src)
}
func (m *ApplicationInfo) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInfo proto.InternalMessageInfo

func (m *ApplicationInfo) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *ApplicationInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ApplicationInfo) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationInfoResponse)(nil), "wins.ApplicationInfoResponse")
	proto.RegisterType((*ApplicationInfo)(nil), "wins.ApplicationInfo")
}

func init() { proto.RegisterFile("application.proto", fileDescriptor_fc846aced8fe6ea6) }

var fileDescriptor_fc846aced8fe6ea6 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x28, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29,
	0xcf, 0xcc, 0x2b, 0x96, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0x85, 0x89, 0x29, 0xb9, 0x70, 0x89, 0x3b,
	0x22, 0x14, 0x7a, 0xe6, 0xa5, 0xe5, 0x07, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x69,
	0x72, 0xb1, 0x80, 0xf8, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xa2, 0x7a, 0x20, 0xdd, 0x7a,
	0xe8, 0x8a, 0xc1, 0x4a, 0x94, 0xe2, 0xb9, 0xf8, 0xd1, 0x24, 0x84, 0xa4, 0xb8, 0x38, 0x9c, 0x33,
	0x52, 0x93, 0xb3, 0x8b, 0x4b, 0x73, 0xc1, 0x26, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x12, 0x5c, 0xec,
	0x61, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x29, 0x18, 0x57, 0x48, 0x8c, 0x8b,
	0xcd, 0x39, 0x3f, 0x37, 0x37, 0xb3, 0x44, 0x82, 0x19, 0x2c, 0x01, 0xe5, 0x19, 0x79, 0x72, 0x09,
	0x21, 0x59, 0x10, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0x0c, 0x71, 0xa1, 0x10, 0x17,
	0xc4, 0x6d, 0x61, 0xf9, 0x99, 0x29, 0x52, 0xb2, 0xd8, 0xdd, 0x09, 0xf5, 0x94, 0x12, 0x83, 0x93,
	0xfc, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x96, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x43, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x6c, 0xb9, 0x50,
	0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ApplicationInfoResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ApplicationInfoResponse, error) {
	out := new(ApplicationInfoResponse)
	err := c.cc.Invoke(ctx, "/wins.ApplicationService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
type ApplicationServiceServer interface {
	Info(context.Context, *Void) (*ApplicationInfoResponse, error)
}

// UnimplementedApplicationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (*UnimplementedApplicationServiceServer) Info(ctx context.Context, req *Void) (*ApplicationInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wins.ApplicationService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Info(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wins.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _ApplicationService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func (m *ApplicationInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplicationInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplication(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApplicationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplicationInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commit) > 0 {
		i -= len(m.Commit)
		copy(dAtA[i:], m.Commit)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Commit)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	offset -= sovApplication(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApplicationInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovApplication(uint64(l))
	}
	return n
}

func (m *ApplicationInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.Commit)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	return n
}

func sovApplication(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApplicationInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ApplicationInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &ApplicationInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApplication
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
func (m *ApplicationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ApplicationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApplication
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
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApplication
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApplication
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApplication        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApplication = fmt.Errorf("proto: unexpected end of group")
)
