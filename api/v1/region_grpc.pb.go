// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: region.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RegionClient is the client API for Region service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionClient interface {
	Create(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*RegionResponse, error)
	Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error)
}

type regionClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionClient(cc grpc.ClientConnInterface) RegionClient {
	return &regionClient{cc}
}

func (c *regionClient) Create(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*RegionResponse, error) {
	out := new(RegionResponse)
	err := c.cc.Invoke(ctx, "/video.v1.Region/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/video.v1.Region/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Delete(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/video.v1.Region/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Get(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error) {
	out := new(GetRegionResponse)
	err := c.cc.Invoke(ctx, "/video.v1.Region/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServer is the server API for Region service.
// All implementations must embed UnimplementedRegionServer
// for forward compatibility
type RegionServer interface {
	Create(context.Context, *UpdateRegionRequest) (*RegionResponse, error)
	Update(context.Context, *UpdateRegionRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteRegionRequest) (*emptypb.Empty, error)
	Get(context.Context, *GetRegionRequest) (*GetRegionResponse, error)
	mustEmbedUnimplementedRegionServer()
}

// UnimplementedRegionServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServer struct {
}

func (UnimplementedRegionServer) Create(context.Context, *UpdateRegionRequest) (*RegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRegionServer) Update(context.Context, *UpdateRegionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRegionServer) Delete(context.Context, *DeleteRegionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRegionServer) Get(context.Context, *GetRegionRequest) (*GetRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRegionServer) mustEmbedUnimplementedRegionServer() {}

// UnsafeRegionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServer will
// result in compilation errors.
type UnsafeRegionServer interface {
	mustEmbedUnimplementedRegionServer()
}

func RegisterRegionServer(s grpc.ServiceRegistrar, srv RegionServer) {
	s.RegisterService(&Region_ServiceDesc, srv)
}

func _Region_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.v1.Region/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Create(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.v1.Region/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Update(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.v1.Region/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Delete(ctx, req.(*DeleteRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.v1.Region/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Get(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Region_ServiceDesc is the grpc.ServiceDesc for Region service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Region_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.v1.Region",
	HandlerType: (*RegionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Region_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Region_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Region_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Region_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "region.proto",
}
