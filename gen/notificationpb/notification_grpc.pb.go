// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: chat/notification.proto

package notificationpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NotificationService_GetNotifications_FullMethodName     = "/notificationpb.NotificationService/GetNotifications"
	NotificationService_MarkNotificationRead_FullMethodName = "/notificationpb.NotificationService/MarkNotificationRead"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	MarkNotificationRead(ctx context.Context, in *MarkReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkNotificationRead(ctx context.Context, in *MarkReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkReadResponse)
	err := c.cc.Invoke(ctx, NotificationService_MarkNotificationRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility.
type NotificationServiceServer interface {
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	MarkNotificationRead(context.Context, *MarkReadRequest) (*MarkReadResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) MarkNotificationRead(context.Context, *MarkReadRequest) (*MarkReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkNotificationRead not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}
func (UnimplementedNotificationServiceServer) testEmbeddedByValue()                             {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	// If the following call pancis, it indicates UnimplementedNotificationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkNotificationRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkNotificationRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_MarkNotificationRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkNotificationRead(ctx, req.(*MarkReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notificationpb.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifications",
			Handler:    _NotificationService_GetNotifications_Handler,
		},
		{
			MethodName: "MarkNotificationRead",
			Handler:    _NotificationService_MarkNotificationRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/notification.proto",
}
