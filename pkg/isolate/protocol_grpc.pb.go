// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package isolate

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DirektivIsolateClient is the client API for DirektivIsolate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirektivIsolateClient interface {
	RunIsolate(ctx context.Context, in *RunIsolateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type direktivIsolateClient struct {
	cc grpc.ClientConnInterface
}

func NewDirektivIsolateClient(cc grpc.ClientConnInterface) DirektivIsolateClient {
	return &direktivIsolateClient{cc}
}

func (c *direktivIsolateClient) RunIsolate(ctx context.Context, in *RunIsolateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/isolate.DirektivIsolate/RunIsolate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirektivIsolateServer is the server API for DirektivIsolate service.
// All implementations must embed UnimplementedDirektivIsolateServer
// for forward compatibility
type DirektivIsolateServer interface {
	RunIsolate(context.Context, *RunIsolateRequest) (*empty.Empty, error)
	mustEmbedUnimplementedDirektivIsolateServer()
}

// UnimplementedDirektivIsolateServer must be embedded to have forward compatible implementations.
type UnimplementedDirektivIsolateServer struct {
}

func (UnimplementedDirektivIsolateServer) RunIsolate(context.Context, *RunIsolateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunIsolate not implemented")
}
func (UnimplementedDirektivIsolateServer) mustEmbedUnimplementedDirektivIsolateServer() {}

// UnsafeDirektivIsolateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirektivIsolateServer will
// result in compilation errors.
type UnsafeDirektivIsolateServer interface {
	mustEmbedUnimplementedDirektivIsolateServer()
}

func RegisterDirektivIsolateServer(s grpc.ServiceRegistrar, srv DirektivIsolateServer) {
	s.RegisterService(&DirektivIsolate_ServiceDesc, srv)
}

func _DirektivIsolate_RunIsolate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunIsolateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIsolateServer).RunIsolate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/isolate.DirektivIsolate/RunIsolate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIsolateServer).RunIsolate(ctx, req.(*RunIsolateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirektivIsolate_ServiceDesc is the grpc.ServiceDesc for DirektivIsolate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirektivIsolate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "isolate.DirektivIsolate",
	HandlerType: (*DirektivIsolateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunIsolate",
			Handler:    _DirektivIsolate_RunIsolate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/isolate/protocol.proto",
}