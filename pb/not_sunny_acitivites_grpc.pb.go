// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// NotSunnyActivitiesClient is the client API for NotSunnyActivities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotSunnyActivitiesClient interface {
	GetAllWeatherActivities(ctx context.Context, in *NotSunnyActivitiesParams, opts ...grpc.CallOption) (*Activity, error)
}

type notSunnyActivitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotSunnyActivitiesClient(cc grpc.ClientConnInterface) NotSunnyActivitiesClient {
	return &notSunnyActivitiesClient{cc}
}

func (c *notSunnyActivitiesClient) GetAllWeatherActivities(ctx context.Context, in *NotSunnyActivitiesParams, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/template.NotSunnyActivities/GetAllWeatherActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotSunnyActivitiesServer is the server API for NotSunnyActivities service.
// All implementations must embed UnimplementedNotSunnyActivitiesServer
// for forward compatibility
type NotSunnyActivitiesServer interface {
	GetAllWeatherActivities(context.Context, *NotSunnyActivitiesParams) (*Activity, error)
	mustEmbedUnimplementedNotSunnyActivitiesServer()
}

// UnimplementedNotSunnyActivitiesServer must be embedded to have forward compatible implementations.
type UnimplementedNotSunnyActivitiesServer struct {
}

func (UnimplementedNotSunnyActivitiesServer) GetAllWeatherActivities(context.Context, *NotSunnyActivitiesParams) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWeatherActivities not implemented")
}
func (UnimplementedNotSunnyActivitiesServer) mustEmbedUnimplementedNotSunnyActivitiesServer() {}

// UnsafeNotSunnyActivitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotSunnyActivitiesServer will
// result in compilation errors.
type UnsafeNotSunnyActivitiesServer interface {
	mustEmbedUnimplementedNotSunnyActivitiesServer()
}

func RegisterNotSunnyActivitiesServer(s grpc.ServiceRegistrar, srv NotSunnyActivitiesServer) {
	s.RegisterService(&NotSunnyActivities_ServiceDesc, srv)
}

func _NotSunnyActivities_GetAllWeatherActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotSunnyActivitiesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotSunnyActivitiesServer).GetAllWeatherActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.NotSunnyActivities/GetAllWeatherActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotSunnyActivitiesServer).GetAllWeatherActivities(ctx, req.(*NotSunnyActivitiesParams))
	}
	return interceptor(ctx, in, info, handler)
}

// NotSunnyActivities_ServiceDesc is the grpc.ServiceDesc for NotSunnyActivities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotSunnyActivities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.NotSunnyActivities",
	HandlerType: (*NotSunnyActivitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllWeatherActivities",
			Handler:    _NotSunnyActivities_GetAllWeatherActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "not_sunny_acitivites.proto",
}