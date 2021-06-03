// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package netwdevpb

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

// CacheUpdateClient is the client API for CacheUpdate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheUpdateClient interface {
	Update(ctx context.Context, in *CacheUpdateRequest, opts ...grpc.CallOption) (*CacheUpdateReply, error)
}

type cacheUpdateClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheUpdateClient(cc grpc.ClientConnInterface) CacheUpdateClient {
	return &cacheUpdateClient{cc}
}

func (c *cacheUpdateClient) Update(ctx context.Context, in *CacheUpdateRequest, opts ...grpc.CallOption) (*CacheUpdateReply, error) {
	out := new(CacheUpdateReply)
	err := c.cc.Invoke(ctx, "/netwdevpb.CacheUpdate/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheUpdateServer is the server API for CacheUpdate service.
// All implementations must embed UnimplementedCacheUpdateServer
// for forward compatibility
type CacheUpdateServer interface {
	Update(context.Context, *CacheUpdateRequest) (*CacheUpdateReply, error)
	mustEmbedUnimplementedCacheUpdateServer()
}

// UnimplementedCacheUpdateServer must be embedded to have forward compatible implementations.
type UnimplementedCacheUpdateServer struct {
}

func (UnimplementedCacheUpdateServer) Update(context.Context, *CacheUpdateRequest) (*CacheUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCacheUpdateServer) mustEmbedUnimplementedCacheUpdateServer() {}

// UnsafeCacheUpdateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheUpdateServer will
// result in compilation errors.
type UnsafeCacheUpdateServer interface {
	mustEmbedUnimplementedCacheUpdateServer()
}

func RegisterCacheUpdateServer(s grpc.ServiceRegistrar, srv CacheUpdateServer) {
	s.RegisterService(&CacheUpdate_ServiceDesc, srv)
}

func _CacheUpdate_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheUpdateServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netwdevpb.CacheUpdate/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheUpdateServer).Update(ctx, req.(*CacheUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheUpdate_ServiceDesc is the grpc.ServiceDesc for CacheUpdate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheUpdate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netwdevpb.CacheUpdate",
	HandlerType: (*CacheUpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _CacheUpdate_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configMessage.proto",
}

// CacheStatusClient is the client API for CacheStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheStatusClient interface {
	Request(ctx context.Context, in *CacheStatusRequest, opts ...grpc.CallOption) (*CacheStatusReply, error)
}

type cacheStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheStatusClient(cc grpc.ClientConnInterface) CacheStatusClient {
	return &cacheStatusClient{cc}
}

func (c *cacheStatusClient) Request(ctx context.Context, in *CacheStatusRequest, opts ...grpc.CallOption) (*CacheStatusReply, error) {
	out := new(CacheStatusReply)
	err := c.cc.Invoke(ctx, "/netwdevpb.CacheStatus/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheStatusServer is the server API for CacheStatus service.
// All implementations must embed UnimplementedCacheStatusServer
// for forward compatibility
type CacheStatusServer interface {
	Request(context.Context, *CacheStatusRequest) (*CacheStatusReply, error)
	mustEmbedUnimplementedCacheStatusServer()
}

// UnimplementedCacheStatusServer must be embedded to have forward compatible implementations.
type UnimplementedCacheStatusServer struct {
}

func (UnimplementedCacheStatusServer) Request(context.Context, *CacheStatusRequest) (*CacheStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedCacheStatusServer) mustEmbedUnimplementedCacheStatusServer() {}

// UnsafeCacheStatusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheStatusServer will
// result in compilation errors.
type UnsafeCacheStatusServer interface {
	mustEmbedUnimplementedCacheStatusServer()
}

func RegisterCacheStatusServer(s grpc.ServiceRegistrar, srv CacheStatusServer) {
	s.RegisterService(&CacheStatus_ServiceDesc, srv)
}

func _CacheStatus_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStatusServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netwdevpb.CacheStatus/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStatusServer).Request(ctx, req.(*CacheStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheStatus_ServiceDesc is the grpc.ServiceDesc for CacheStatus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheStatus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netwdevpb.CacheStatus",
	HandlerType: (*CacheStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _CacheStatus_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configMessage.proto",
}

// DeviationClient is the client API for Deviation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviationClient interface {
	Update(ctx context.Context, in *DeviationUpdate, opts ...grpc.CallOption) (*DeviationUpdateReply, error)
}

type deviationClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviationClient(cc grpc.ClientConnInterface) DeviationClient {
	return &deviationClient{cc}
}

func (c *deviationClient) Update(ctx context.Context, in *DeviationUpdate, opts ...grpc.CallOption) (*DeviationUpdateReply, error) {
	out := new(DeviationUpdateReply)
	err := c.cc.Invoke(ctx, "/netwdevpb.Deviation/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviationServer is the server API for Deviation service.
// All implementations must embed UnimplementedDeviationServer
// for forward compatibility
type DeviationServer interface {
	Update(context.Context, *DeviationUpdate) (*DeviationUpdateReply, error)
	mustEmbedUnimplementedDeviationServer()
}

// UnimplementedDeviationServer must be embedded to have forward compatible implementations.
type UnimplementedDeviationServer struct {
}

func (UnimplementedDeviationServer) Update(context.Context, *DeviationUpdate) (*DeviationUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeviationServer) mustEmbedUnimplementedDeviationServer() {}

// UnsafeDeviationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviationServer will
// result in compilation errors.
type UnsafeDeviationServer interface {
	mustEmbedUnimplementedDeviationServer()
}

func RegisterDeviationServer(s grpc.ServiceRegistrar, srv DeviationServer) {
	s.RegisterService(&Deviation_ServiceDesc, srv)
}

func _Deviation_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviationUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netwdevpb.Deviation/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviationServer).Update(ctx, req.(*DeviationUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// Deviation_ServiceDesc is the grpc.ServiceDesc for Deviation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deviation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netwdevpb.Deviation",
	HandlerType: (*DeviationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Deviation_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configMessage.proto",
}
