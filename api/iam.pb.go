// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/iam.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{0}
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
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{1}
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

func (m *LoginResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyAccessTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyAccessTokenRequest) Reset()         { *m = VerifyAccessTokenRequest{} }
func (m *VerifyAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyAccessTokenRequest) ProtoMessage()    {}
func (*VerifyAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{2}
}

func (m *VerifyAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyAccessTokenRequest.Unmarshal(m, b)
}
func (m *VerifyAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *VerifyAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyAccessTokenRequest.Merge(m, src)
}
func (m *VerifyAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyAccessTokenRequest.Size(m)
}
func (m *VerifyAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyAccessTokenRequest proto.InternalMessageInfo

func (m *VerifyAccessTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyAccessTokenResponse struct {
	Status               bool                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ThingId              int64                `protobuf:"varint,3,opt,name=thing_id,json=thingId,proto3" json:"thing_id,omitempty"`
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VerifyAccessTokenResponse) Reset()         { *m = VerifyAccessTokenResponse{} }
func (m *VerifyAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyAccessTokenResponse) ProtoMessage()    {}
func (*VerifyAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{3}
}

func (m *VerifyAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyAccessTokenResponse.Unmarshal(m, b)
}
func (m *VerifyAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyAccessTokenResponse.Merge(m, src)
}
func (m *VerifyAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyAccessTokenResponse.Size(m)
}
func (m *VerifyAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyAccessTokenResponse proto.InternalMessageInfo

func (m *VerifyAccessTokenResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *VerifyAccessTokenResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *VerifyAccessTokenResponse) GetThingId() int64 {
	if m != nil {
		return m.ThingId
	}
	return 0
}

func (m *VerifyAccessTokenResponse) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type RefreshAccessTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshAccessTokenRequest) Reset()         { *m = RefreshAccessTokenRequest{} }
func (m *RefreshAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshAccessTokenRequest) ProtoMessage()    {}
func (*RefreshAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{4}
}

func (m *RefreshAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshAccessTokenRequest.Unmarshal(m, b)
}
func (m *RefreshAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *RefreshAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshAccessTokenRequest.Merge(m, src)
}
func (m *RefreshAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshAccessTokenRequest.Size(m)
}
func (m *RefreshAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshAccessTokenRequest proto.InternalMessageInfo

func (m *RefreshAccessTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshAccessTokenResponse struct {
	Status               bool                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ThingId              int64                `protobuf:"varint,3,opt,name=thing_id,json=thingId,proto3" json:"thing_id,omitempty"`
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RefreshAccessTokenResponse) Reset()         { *m = RefreshAccessTokenResponse{} }
func (m *RefreshAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshAccessTokenResponse) ProtoMessage()    {}
func (*RefreshAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dbcc987c0695300, []int{5}
}

func (m *RefreshAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshAccessTokenResponse.Unmarshal(m, b)
}
func (m *RefreshAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *RefreshAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshAccessTokenResponse.Merge(m, src)
}
func (m *RefreshAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshAccessTokenResponse.Size(m)
}
func (m *RefreshAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshAccessTokenResponse proto.InternalMessageInfo

func (m *RefreshAccessTokenResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *RefreshAccessTokenResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshAccessTokenResponse) GetThingId() int64 {
	if m != nil {
		return m.ThingId
	}
	return 0
}

func (m *RefreshAccessTokenResponse) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
	proto.RegisterType((*VerifyAccessTokenRequest)(nil), "api.VerifyAccessTokenRequest")
	proto.RegisterType((*VerifyAccessTokenResponse)(nil), "api.VerifyAccessTokenResponse")
	proto.RegisterType((*RefreshAccessTokenRequest)(nil), "api.RefreshAccessTokenRequest")
	proto.RegisterType((*RefreshAccessTokenResponse)(nil), "api.RefreshAccessTokenResponse")
}

func init() { proto.RegisterFile("api/iam.proto", fileDescriptor_4dbcc987c0695300) }

var fileDescriptor_4dbcc987c0695300 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x7d, 0x7d, 0x3c, 0xbe, 0xee, 0x7b, 0x2c, 0x98, 0xbc, 0x68, 0x99, 0x44, 0x21, 0x5d, 0xb1,
	0x2a, 0x88, 0x2b, 0x17, 0x26, 0xb2, 0x24, 0x91, 0x4d, 0x43, 0x74, 0x49, 0x06, 0x18, 0xca, 0x04,
	0xda, 0x19, 0x7b, 0xa7, 0x51, 0xff, 0x8e, 0xfe, 0x41, 0x7f, 0x82, 0xe9, 0x4c, 0x21, 0x4d, 0xa0,
	0xd1, 0xa5, 0xcb, 0x73, 0xcf, 0x3d, 0xa7, 0x3d, 0x67, 0x2e, 0xb4, 0x98, 0x12, 0x03, 0xc1, 0x22,
	0x5f, 0x25, 0x52, 0x4b, 0x52, 0x61, 0x4a, 0xd0, 0x6e, 0x28, 0x65, 0xb8, 0xe3, 0x03, 0x33, 0x5a,
	0xa4, 0xeb, 0x81, 0x16, 0x11, 0x47, 0xcd, 0x22, 0x65, 0xb7, 0xbc, 0x3b, 0xf8, 0x77, 0x2f, 0x43,
	0x11, 0x07, 0xfc, 0x29, 0xe5, 0xa8, 0xc9, 0x7f, 0xa8, 0xf2, 0x88, 0x89, 0x9d, 0xeb, 0xf4, 0x9c,
	0x7e, 0x33, 0xb0, 0x80, 0x50, 0x68, 0x28, 0x86, 0xf8, 0x2c, 0x93, 0x95, 0xfb, 0xdb, 0x10, 0x07,
	0xec, 0xdd, 0x42, 0x2b, 0x77, 0x40, 0x25, 0x63, 0xe4, 0xe4, 0x0c, 0x6a, 0xa8, 0x99, 0x4e, 0xd1,
	0x78, 0x34, 0x82, 0x1c, 0x65, 0xd6, 0x5a, 0x6e, 0x79, 0x9c, 0x3b, 0x58, 0xe0, 0x0d, 0xc1, 0x7d,
	0xe0, 0x89, 0x58, 0xbf, 0x8e, 0x97, 0x4b, 0x8e, 0x38, 0xcb, 0x86, 0x85, 0x9f, 0xb1, 0x0a, 0xa7,
	0xa8, 0x78, 0x73, 0xa0, 0x73, 0x42, 0xf2, 0xc5, 0xd7, 0xcf, 0xa1, 0x9e, 0x22, 0x4f, 0xe6, 0xc2,
	0x26, 0xa8, 0x04, 0xb5, 0x0c, 0x4e, 0x56, 0xa4, 0x03, 0x0d, 0xbd, 0x11, 0x71, 0x98, 0x31, 0x15,
	0xc3, 0xd4, 0x0d, 0x9e, 0xac, 0xc8, 0x0d, 0x00, 0x7f, 0x51, 0x22, 0xe1, 0x38, 0x67, 0xda, 0xfd,
	0xd3, 0x73, 0xfa, 0x7f, 0x47, 0xd4, 0xb7, 0x95, 0xfa, 0xfb, 0x4a, 0xfd, 0xd9, 0xbe, 0xd2, 0xa0,
	0x99, 0x6f, 0x8f, 0xb5, 0x77, 0x05, 0x9d, 0x80, 0xaf, 0x13, 0x8e, 0x9b, 0x6f, 0xe7, 0x7a, 0x77,
	0x80, 0x9e, 0xd2, 0xfc, 0xa8, 0x60, 0xa3, 0x0f, 0x07, 0x5a, 0x53, 0xb1, 0x65, 0x4a, 0xc6, 0x62,
	0x89, 0x93, 0xf1, 0x94, 0x0c, 0xa1, 0x6a, 0x0e, 0x80, 0xb4, 0x7d, 0xa6, 0x84, 0x5f, 0x3c, 0x27,
	0x4a, 0x8a, 0x23, 0x1b, 0xc4, 0xfb, 0x45, 0x66, 0xd0, 0x3e, 0x7a, 0x40, 0x72, 0x61, 0x56, 0xcb,
	0x6e, 0x81, 0x5e, 0x96, 0xd1, 0x07, 0xd7, 0x47, 0x20, 0xc7, 0xf5, 0x11, 0xab, 0x2b, 0x7d, 0x0b,
	0xda, 0x2d, 0xe5, 0xf7, 0xc6, 0x8b, 0x9a, 0x69, 0xe4, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0x25,
	0xf8, 0x83, 0xbe, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MikaponicsIAMClient is the client API for MikaponicsIAM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MikaponicsIAMClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	VerifyAccessToken(ctx context.Context, in *VerifyAccessTokenRequest, opts ...grpc.CallOption) (*VerifyAccessTokenResponse, error)
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
}

type mikaponicsIAMClient struct {
	cc *grpc.ClientConn
}

func NewMikaponicsIAMClient(cc *grpc.ClientConn) MikaponicsIAMClient {
	return &mikaponicsIAMClient{cc}
}

func (c *mikaponicsIAMClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsIAM/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mikaponicsIAMClient) VerifyAccessToken(ctx context.Context, in *VerifyAccessTokenRequest, opts ...grpc.CallOption) (*VerifyAccessTokenResponse, error) {
	out := new(VerifyAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsIAM/VerifyAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mikaponicsIAMClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/api.MikaponicsIAM/RefreshAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MikaponicsIAMServer is the server API for MikaponicsIAM service.
type MikaponicsIAMServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	VerifyAccessToken(context.Context, *VerifyAccessTokenRequest) (*VerifyAccessTokenResponse, error)
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

// UnimplementedMikaponicsIAMServer can be embedded to have forward compatible implementations.
type UnimplementedMikaponicsIAMServer struct {
}

func (*UnimplementedMikaponicsIAMServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedMikaponicsIAMServer) VerifyAccessToken(ctx context.Context, req *VerifyAccessTokenRequest) (*VerifyAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccessToken not implemented")
}
func (*UnimplementedMikaponicsIAMServer) RefreshAccessToken(ctx context.Context, req *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}

func RegisterMikaponicsIAMServer(s *grpc.Server, srv MikaponicsIAMServer) {
	s.RegisterService(&_MikaponicsIAM_serviceDesc, srv)
}

func _MikaponicsIAM_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsIAMServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsIAM/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsIAMServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MikaponicsIAM_VerifyAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsIAMServer).VerifyAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsIAM/VerifyAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsIAMServer).VerifyAccessToken(ctx, req.(*VerifyAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MikaponicsIAM_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikaponicsIAMServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MikaponicsIAM/RefreshAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikaponicsIAMServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MikaponicsIAM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MikaponicsIAM",
	HandlerType: (*MikaponicsIAMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _MikaponicsIAM_Login_Handler,
		},
		{
			MethodName: "VerifyAccessToken",
			Handler:    _MikaponicsIAM_VerifyAccessToken_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _MikaponicsIAM_RefreshAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/iam.proto",
}
