// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

type ChangePasswordRequest struct {
	CurrentPassword      string   `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordRequest) Reset()         { *m = ChangePasswordRequest{} }
func (m *ChangePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordRequest) ProtoMessage()    {}
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{2}
}

func (m *ChangePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordRequest.Unmarshal(m, b)
}
func (m *ChangePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordRequest.Marshal(b, m, deterministic)
}
func (m *ChangePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordRequest.Merge(m, src)
}
func (m *ChangePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordRequest.Size(m)
}
func (m *ChangePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordRequest proto.InternalMessageInfo

func (m *ChangePasswordRequest) GetCurrentPassword() string {
	if m != nil {
		return m.CurrentPassword
	}
	return ""
}

func (m *ChangePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordResponse) Reset()         { *m = ChangePasswordResponse{} }
func (m *ChangePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordResponse) ProtoMessage()    {}
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{3}
}

func (m *ChangePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordResponse.Unmarshal(m, b)
}
func (m *ChangePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordResponse.Marshal(b, m, deterministic)
}
func (m *ChangePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordResponse.Merge(m, src)
}
func (m *ChangePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordResponse.Size(m)
}
func (m *ChangePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
	proto.RegisterType((*ChangePasswordRequest)(nil), "user.ChangePasswordRequest")
	proto.RegisterType((*ChangePasswordResponse)(nil), "user.ChangePasswordResponse")
}

func init() { proto.RegisterFile("proto/v1/user/user.proto", fileDescriptor_a8fd1c8d22ccfa2e) }

var fileDescriptor_a8fd1c8d22ccfa2e = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4b, 0x4e, 0x2a, 0x41,
	0x14, 0x4d, 0xf3, 0x7b, 0xbc, 0x82, 0xf7, 0x1e, 0xd4, 0x53, 0xe8, 0xb4, 0x0c, 0x4c, 0x4f, 0x34,
	0x24, 0x74, 0xf9, 0x49, 0x1c, 0x30, 0x32, 0xb0, 0x01, 0x83, 0x3a, 0x71, 0x82, 0x45, 0x73, 0xd3,
	0x54, 0x02, 0x55, 0x65, 0x55, 0x35, 0x0c, 0x4d, 0xdc, 0x82, 0xbb, 0x71, 0xe4, 0x0a, 0x9c, 0xb8,
	0x05, 0x27, 0x6e, 0xc2, 0x18, 0xba, 0xbb, 0xf0, 0x03, 0x71, 0xd2, 0xb9, 0xe9, 0x73, 0xee, 0x3d,
	0xe7, 0x9e, 0xba, 0xc8, 0x95, 0x4a, 0x18, 0x41, 0xe6, 0x87, 0x24, 0xd6, 0xa0, 0x92, 0x4f, 0x90,
	0xfc, 0xc2, 0x85, 0x65, 0xed, 0xb5, 0x22, 0x21, 0xa2, 0x29, 0x10, 0x2a, 0x19, 0xa1, 0x9c, 0x0b,
	0x43, 0x0d, 0x13, 0x5c, 0xa7, 0x1c, 0xaf, 0x39, 0xa7, 0x53, 0x36, 0xa6, 0x06, 0x88, 0x2d, 0x52,
	0xc0, 0xbf, 0x46, 0xf5, 0xbe, 0x02, 0x6a, 0xe0, 0x52, 0x83, 0x1a, 0xc0, 0x4d, 0x0c, 0xda, 0x60,
	0x1f, 0x15, 0x61, 0x46, 0xd9, 0xd4, 0x75, 0x76, 0x9d, 0xfd, 0xdf, 0xbd, 0xea, 0xc3, 0xeb, 0x63,
	0xfe, 0x97, 0x2a, 0xd6, 0xf2, 0xee, 0x9b, 0x33, 0x48, 0x21, 0xbc, 0x87, 0xca, 0x92, 0x6a, 0xbd,
	0x10, 0x6a, 0xec, 0xe6, 0x12, 0x5a, 0x65, 0x49, 0x2b, 0xa9, 0x42, 0xad, 0xec, 0x9e, 0x0e, 0x56,
	0xa0, 0xbf, 0x85, 0xf0, 0x67, 0x05, 0x2d, 0x05, 0xd7, 0xe0, 0xdf, 0xa2, 0xed, 0xfe, 0x84, 0xf2,
	0x08, 0xce, 0x32, 0x9e, 0xd5, 0x3e, 0x41, 0xb5, 0x30, 0x56, 0x0a, 0xb8, 0x19, 0xae, 0xe6, 0x3b,
	0xeb, 0xf3, 0xff, 0x65, 0x24, 0xdb, 0x8e, 0x03, 0x54, 0xe5, 0xb0, 0x18, 0xfe, 0xe4, 0xa9, 0xc2,
	0x61, 0x61, 0xf9, 0xbe, 0x8b, 0x1a, 0xdf, 0x0d, 0xa4, 0xd6, 0x8e, 0x9e, 0x1c, 0x54, 0x59, 0x7a,
	0x3d, 0x07, 0x35, 0x67, 0x21, 0xe0, 0x0b, 0x84, 0x3e, 0x16, 0xc0, 0xcd, 0x20, 0x89, 0x7e, 0x2d,
	0x34, 0xcf, 0x5d, 0x07, 0xb2, 0x5d, 0xff, 0xdf, 0x3d, 0xbf, 0xdc, 0xe7, 0xfe, 0xf8, 0x65, 0xfb,
	0x7a, 0x5d, 0xa7, 0x8d, 0x19, 0xfa, 0xfb, 0x55, 0x1f, 0xef, 0x64, 0x03, 0x36, 0xc5, 0xe2, 0xb5,
	0x36, 0x83, 0x99, 0x42, 0x2b, 0x51, 0x68, 0x78, 0xf5, 0xd5, 0x7d, 0xd8, 0x1c, 0xba, 0x4e, 0xbb,
	0x77, 0x70, 0x15, 0x44, 0xcc, 0x4c, 0xe2, 0x51, 0x10, 0x8a, 0x19, 0x61, 0xda, 0xe8, 0x09, 0x89,
	0x44, 0x27, 0x52, 0x32, 0xec, 0x68, 0x3a, 0x93, 0xc9, 0xd5, 0x48, 0x22, 0x47, 0xb6, 0x7b, 0x54,
	0x4a, 0x8e, 0xe3, 0xf8, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x59, 0x33, 0x7e, 0x75, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user/user.proto",
}
