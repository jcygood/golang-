// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package registerproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RegisterReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterRsp struct {
	Errorid              int32    `protobuf:"varint,1,opt,name=errorid,proto3" json:"errorid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Userid               int32    `protobuf:"varint,3,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRsp) Reset()         { *m = RegisterRsp{} }
func (m *RegisterRsp) String() string { return proto.CompactTextString(m) }
func (*RegisterRsp) ProtoMessage()    {}
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRsp.Unmarshal(m, b)
}
func (m *RegisterRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRsp.Marshal(b, m, deterministic)
}
func (m *RegisterRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRsp.Merge(m, src)
}
func (m *RegisterRsp) XXX_Size() int {
	return xxx_messageInfo_RegisterRsp.Size(m)
}
func (m *RegisterRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRsp proto.InternalMessageInfo

func (m *RegisterRsp) GetErrorid() int32 {
	if m != nil {
		return m.Errorid
	}
	return 0
}

func (m *RegisterRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRsp) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "registerproto.RegisterReq")
	proto.RegisterType((*RegisterRsp)(nil), "registerproto.RegisterRsp")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0xf1, 0xc1, 0x5c,
	0x25, 0x45, 0x2e, 0xee, 0x20, 0xa8, 0x40, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x14, 0x8c, 0xa4, 0xa4, 0xb8,
	0x40, 0x48, 0x82, 0x8b, 0x3d, 0xb5, 0xa8, 0x28, 0xbf, 0x28, 0x33, 0x05, 0xac, 0x8a, 0x35, 0x08,
	0xc6, 0x85, 0x6b, 0x66, 0x42, 0x68, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x2d, 0x4e, 0x05, 0x29, 0x66,
	0x06, 0x2b, 0x86, 0xf2, 0x8c, 0xc2, 0xb9, 0xf8, 0x61, 0x86, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x0a, 0xb9, 0x70, 0x71, 0xc0, 0x84, 0x84, 0xa4, 0xf4, 0x50, 0x9c, 0xa9, 0x87, 0xe4, 0x46,
	0x29, 0x9c, 0x72, 0xc5, 0x05, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0x20, 0xee, 0xa2, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error)
}

type registerServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServiceClient(cc *grpc.ClientConn) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/registerproto.RegisterService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
type RegisterServiceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRsp, error)
}

// UnimplementedRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (*UnimplementedRegisterServiceServer) Register(ctx context.Context, req *RegisterReq) (*RegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegisterServiceServer(s *grpc.Server, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerproto.RegisterService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registerproto.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegisterService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}