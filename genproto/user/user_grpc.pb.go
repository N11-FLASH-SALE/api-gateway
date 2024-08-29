// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user.proto

package user

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
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	GetUSerByEmail(ctx context.Context, in *GetUSerByEmailReq, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*Void, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error)
	IsUserExist(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, "/user.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/user.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUSerByEmail(ctx context.Context, in *GetUSerByEmailReq, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUSerByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.User/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.User/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) IsUserExist(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/user.User/IsUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
	GetUSerByEmail(context.Context, *GetUSerByEmailReq) (*GetUserResponse, error)
	GetUserById(context.Context, *UserId) (*GetUserResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*Void, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Void, error)
	DeleteUser(context.Context, *UserId) (*Void, error)
	IsUserExist(context.Context, *UserId) (*Void, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) GetUSerByEmail(context.Context, *GetUSerByEmailReq) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUSerByEmail not implemented")
}
func (UnimplementedUserServer) GetUserById(context.Context, *UserId) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServer) ResetPassword(context.Context, *ResetPasswordReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *UserId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) IsUserExist(context.Context, *UserId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserExist not implemented")
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

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUSerByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUSerByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUSerByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUSerByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUSerByEmail(ctx, req.(*GetUSerByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_IsUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).IsUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/IsUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).IsUserExist(ctx, req.(*UserId))
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
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "GetUSerByEmail",
			Handler:    _User_GetUSerByEmail_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _User_UpdatePassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _User_ResetPassword_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "IsUserExist",
			Handler:    _User_IsUserExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsClient interface {
	CreateNotification(ctx context.Context, in *CreateNotificationsReq, opts ...grpc.CallOption) (*CreateNotificationsRes, error)
	GetAllNotifications(ctx context.Context, in *GetNotificationsReq, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	GetAndMarkNotificationAsRead(ctx context.Context, in *GetAndMarkNotificationAsReadReq, opts ...grpc.CallOption) (*GetAndMarkNotificationAsReadRes, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) CreateNotification(ctx context.Context, in *CreateNotificationsReq, opts ...grpc.CallOption) (*CreateNotificationsRes, error) {
	out := new(CreateNotificationsRes)
	err := c.cc.Invoke(ctx, "/user.Notifications/CreateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetAllNotifications(ctx context.Context, in *GetNotificationsReq, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/user.Notifications/GetAllNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetAndMarkNotificationAsRead(ctx context.Context, in *GetAndMarkNotificationAsReadReq, opts ...grpc.CallOption) (*GetAndMarkNotificationAsReadRes, error) {
	out := new(GetAndMarkNotificationAsReadRes)
	err := c.cc.Invoke(ctx, "/user.Notifications/GetAndMarkNotificationAsRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
// All implementations must embed UnimplementedNotificationsServer
// for forward compatibility
type NotificationsServer interface {
	CreateNotification(context.Context, *CreateNotificationsReq) (*CreateNotificationsRes, error)
	GetAllNotifications(context.Context, *GetNotificationsReq) (*GetNotificationsResponse, error)
	GetAndMarkNotificationAsRead(context.Context, *GetAndMarkNotificationAsReadReq) (*GetAndMarkNotificationAsReadRes, error)
	mustEmbedUnimplementedNotificationsServer()
}

// UnimplementedNotificationsServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (UnimplementedNotificationsServer) CreateNotification(context.Context, *CreateNotificationsReq) (*CreateNotificationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationsServer) GetAllNotifications(context.Context, *GetNotificationsReq) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedNotificationsServer) GetAndMarkNotificationAsRead(context.Context, *GetAndMarkNotificationAsReadReq) (*GetAndMarkNotificationAsReadRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndMarkNotificationAsRead not implemented")
}
func (UnimplementedNotificationsServer) mustEmbedUnimplementedNotificationsServer() {}

// UnsafeNotificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServer will
// result in compilation errors.
type UnsafeNotificationsServer interface {
	mustEmbedUnimplementedNotificationsServer()
}

func RegisterNotificationsServer(s grpc.ServiceRegistrar, srv NotificationsServer) {
	s.RegisterService(&Notifications_ServiceDesc, srv)
}

func _Notifications_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Notifications/CreateNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).CreateNotification(ctx, req.(*CreateNotificationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetAllNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetAllNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Notifications/GetAllNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetAllNotifications(ctx, req.(*GetNotificationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetAndMarkNotificationAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAndMarkNotificationAsReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetAndMarkNotificationAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Notifications/GetAndMarkNotificationAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetAndMarkNotificationAsRead(ctx, req.(*GetAndMarkNotificationAsReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifications_ServiceDesc is the grpc.ServiceDesc for Notifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _Notifications_CreateNotification_Handler,
		},
		{
			MethodName: "GetAllNotifications",
			Handler:    _Notifications_GetAllNotifications_Handler,
		},
		{
			MethodName: "GetAndMarkNotificationAsRead",
			Handler:    _Notifications_GetAndMarkNotificationAsRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// CardsClient is the client API for Cards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardsClient interface {
	CreateCard(ctx context.Context, in *CreateCardReq, opts ...grpc.CallOption) (*CreateCardRes, error)
	GetCardsOfUser(ctx context.Context, in *GetCardsOfUserReq, opts ...grpc.CallOption) (*GetCardsOfUserRes, error)
	GetCardAmount(ctx context.Context, in *GetCardAmountReq, opts ...grpc.CallOption) (*GetCardAmountRes, error)
	UpdateCardAmount(ctx context.Context, in *UpdateCardAmountReq, opts ...grpc.CallOption) (*UpdateCardAmountRes, error)
}

type cardsClient struct {
	cc grpc.ClientConnInterface
}

func NewCardsClient(cc grpc.ClientConnInterface) CardsClient {
	return &cardsClient{cc}
}

func (c *cardsClient) CreateCard(ctx context.Context, in *CreateCardReq, opts ...grpc.CallOption) (*CreateCardRes, error) {
	out := new(CreateCardRes)
	err := c.cc.Invoke(ctx, "/user.Cards/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsClient) GetCardsOfUser(ctx context.Context, in *GetCardsOfUserReq, opts ...grpc.CallOption) (*GetCardsOfUserRes, error) {
	out := new(GetCardsOfUserRes)
	err := c.cc.Invoke(ctx, "/user.Cards/GetCardsOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsClient) GetCardAmount(ctx context.Context, in *GetCardAmountReq, opts ...grpc.CallOption) (*GetCardAmountRes, error) {
	out := new(GetCardAmountRes)
	err := c.cc.Invoke(ctx, "/user.Cards/GetCardAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsClient) UpdateCardAmount(ctx context.Context, in *UpdateCardAmountReq, opts ...grpc.CallOption) (*UpdateCardAmountRes, error) {
	out := new(UpdateCardAmountRes)
	err := c.cc.Invoke(ctx, "/user.Cards/UpdateCardAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardsServer is the server API for Cards service.
// All implementations must embed UnimplementedCardsServer
// for forward compatibility
type CardsServer interface {
	CreateCard(context.Context, *CreateCardReq) (*CreateCardRes, error)
	GetCardsOfUser(context.Context, *GetCardsOfUserReq) (*GetCardsOfUserRes, error)
	GetCardAmount(context.Context, *GetCardAmountReq) (*GetCardAmountRes, error)
	UpdateCardAmount(context.Context, *UpdateCardAmountReq) (*UpdateCardAmountRes, error)
	mustEmbedUnimplementedCardsServer()
}

// UnimplementedCardsServer must be embedded to have forward compatible implementations.
type UnimplementedCardsServer struct {
}

func (UnimplementedCardsServer) CreateCard(context.Context, *CreateCardReq) (*CreateCardRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedCardsServer) GetCardsOfUser(context.Context, *GetCardsOfUserReq) (*GetCardsOfUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardsOfUser not implemented")
}
func (UnimplementedCardsServer) GetCardAmount(context.Context, *GetCardAmountReq) (*GetCardAmountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardAmount not implemented")
}
func (UnimplementedCardsServer) UpdateCardAmount(context.Context, *UpdateCardAmountReq) (*UpdateCardAmountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardAmount not implemented")
}
func (UnimplementedCardsServer) mustEmbedUnimplementedCardsServer() {}

// UnsafeCardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardsServer will
// result in compilation errors.
type UnsafeCardsServer interface {
	mustEmbedUnimplementedCardsServer()
}

func RegisterCardsServer(s grpc.ServiceRegistrar, srv CardsServer) {
	s.RegisterService(&Cards_ServiceDesc, srv)
}

func _Cards_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Cards/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServer).CreateCard(ctx, req.(*CreateCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cards_GetCardsOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardsOfUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServer).GetCardsOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Cards/GetCardsOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServer).GetCardsOfUser(ctx, req.(*GetCardsOfUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cards_GetCardAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardAmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServer).GetCardAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Cards/GetCardAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServer).GetCardAmount(ctx, req.(*GetCardAmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cards_UpdateCardAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardAmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServer).UpdateCardAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Cards/UpdateCardAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServer).UpdateCardAmount(ctx, req.(*UpdateCardAmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cards_ServiceDesc is the grpc.ServiceDesc for Cards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Cards",
	HandlerType: (*CardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCard",
			Handler:    _Cards_CreateCard_Handler,
		},
		{
			MethodName: "GetCardsOfUser",
			Handler:    _Cards_GetCardsOfUser_Handler,
		},
		{
			MethodName: "GetCardAmount",
			Handler:    _Cards_GetCardAmount_Handler,
		},
		{
			MethodName: "UpdateCardAmount",
			Handler:    _Cards_UpdateCardAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
