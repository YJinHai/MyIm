// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google/api"
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

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "proto.LoginResponse")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2d, 0x4e, 0x2d,
	0xd2, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x4a, 0x4a, 0x26, 0x3d,
	0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1,
	0x24, 0x33, 0x3f, 0xaf, 0x18, 0xa2, 0x48, 0xc9, 0x81, 0x8b, 0xc7, 0x27, 0x3f, 0x3d, 0x33, 0x2f,
	0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x84, 0x8b, 0xd5, 0x35, 0x37, 0x31, 0x33, 0x47,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x92, 0xe2, 0xe2, 0x08, 0x48, 0x2c, 0x2e,
	0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x02, 0x4b, 0xc0, 0xf9, 0x4a, 0x8a, 0x5c, 0xbc, 0x50, 0x13,
	0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x43, 0x33, 0x53, 0xa0, 0x06, 0x80,
	0x98, 0x46, 0x01, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x1e, 0x5c, 0x1c, 0x41, 0x05, 0xc9,
	0x60, 0xd5, 0x42, 0xc2, 0x10, 0x07, 0xe8, 0x21, 0xdb, 0x2e, 0x25, 0x82, 0x2a, 0x08, 0x31, 0x50,
	0x49, 0xb0, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xdc, 0x4a, 0x6c, 0xfa, 0x39, 0x20, 0x71, 0x2b, 0x46,
	0xad, 0x24, 0x36, 0xb0, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xe9, 0xbb, 0x05,
	0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	RpcLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) RpcLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.User/RpcLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	RpcLogin(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) RpcLogin(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcLogin not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_RpcLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RpcLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/RpcLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RpcLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcLogin",
			Handler:    _User_RpcLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
