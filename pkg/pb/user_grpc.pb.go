// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/pb/user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	AddName(ctx context.Context, in *AddNameRequest, opts ...grpc.CallOption) (*AddNameResponse, error)
	AddPhone(ctx context.Context, in *AddPhoneRequest, opts ...grpc.CallOption) (*AddPhoneResponse, error)
	IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminResponse, error)
	GetUsersByChatId(ctx context.Context, in *GetUsersByChatIdRequest, opts ...grpc.CallOption) (*GetUsersByChatIdResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) AddName(ctx context.Context, in *AddNameRequest, opts ...grpc.CallOption) (*AddNameResponse, error) {
	out := new(AddNameResponse)
	err := c.cc.Invoke(ctx, "/user.User/AddName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddPhone(ctx context.Context, in *AddPhoneRequest, opts ...grpc.CallOption) (*AddPhoneResponse, error) {
	out := new(AddPhoneResponse)
	err := c.cc.Invoke(ctx, "/user.User/AddPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminResponse, error) {
	out := new(IsAdminResponse)
	err := c.cc.Invoke(ctx, "/user.User/IsAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersByChatId(ctx context.Context, in *GetUsersByChatIdRequest, opts ...grpc.CallOption) (*GetUsersByChatIdResponse, error) {
	out := new(GetUsersByChatIdResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUsersByChatId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	AddName(context.Context, *AddNameRequest) (*AddNameResponse, error)
	AddPhone(context.Context, *AddPhoneRequest) (*AddPhoneResponse, error)
	IsAdmin(context.Context, *IsAdminRequest) (*IsAdminResponse, error)
	GetUsersByChatId(context.Context, *GetUsersByChatIdRequest) (*GetUsersByChatIdResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) AddName(context.Context, *AddNameRequest) (*AddNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddName not implemented")
}
func (UnimplementedUserServer) AddPhone(context.Context, *AddPhoneRequest) (*AddPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhone not implemented")
}
func (UnimplementedUserServer) IsAdmin(context.Context, *IsAdminRequest) (*IsAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAdmin not implemented")
}
func (UnimplementedUserServer) GetUsersByChatId(context.Context, *GetUsersByChatIdRequest) (*GetUsersByChatIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByChatId not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_AddName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/AddName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddName(ctx, req.(*AddNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/AddPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddPhone(ctx, req.(*AddPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_IsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).IsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/IsAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).IsAdmin(ctx, req.(*IsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersByChatId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByChatIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersByChatId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUsersByChatId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersByChatId(ctx, req.(*GetUsersByChatIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddName",
			Handler:    _User_AddName_Handler,
		},
		{
			MethodName: "AddPhone",
			Handler:    _User_AddPhone_Handler,
		},
		{
			MethodName: "IsAdmin",
			Handler:    _User_IsAdmin_Handler,
		},
		{
			MethodName: "GetUsersByChatId",
			Handler:    _User_GetUsersByChatId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/user.proto",
}