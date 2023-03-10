// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: api/jobs/v1/jobs.proto

package v1

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

// JobsClient is the client API for Jobs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsClient interface {
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsReply, error)
}

type jobsClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsClient(cc grpc.ClientConnInterface) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsReply, error) {
	out := new(ListJobsReply)
	err := c.cc.Invoke(ctx, "/api.Jobs.v1.Jobs/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServer is the serviceName API for Jobs service.
// All implementations must embed UnimplementedJobsServer
// for forward compatibility
type JobsServer interface {
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsReply, error)
	mustEmbedUnimplementedJobsServer()
}

// UnimplementedJobsServer must be embedded to have forward compatible implementations.
type UnimplementedJobsServer struct {
}

func (UnimplementedJobsServer) ListJobs(context.Context, *ListJobsRequest) (*ListJobsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (UnimplementedJobsServer) mustEmbedUnimplementedJobsServer() {}

// UnsafeJobsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobsServer will
// result in compilation errors.
type UnsafeJobsServer interface {
	mustEmbedUnimplementedJobsServer()
}

func RegisterJobsServer(s grpc.ServiceRegistrar, srv JobsServer) {
	s.RegisterService(&Jobs_ServiceDesc, srv)
}

func _Jobs_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Jobs.v1.Jobs/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jobs_ServiceDesc is the grpc.ServiceDesc for Jobs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jobs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Jobs.v1.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListJobs",
			Handler:    _Jobs_ListJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/jobs/v1/jobs.proto",
}
