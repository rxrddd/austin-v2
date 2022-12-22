// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: api/files/v1/files.proto

package v1

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

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesClient interface {
	GetOssStsToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OssStsTokenResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
}

type filesClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesClient(cc grpc.ClientConnInterface) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) GetOssStsToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OssStsTokenResponse, error) {
	out := new(OssStsTokenResponse)
	err := c.cc.Invoke(ctx, "/api.files.v1.Files/GetOssStsToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/api.files.v1.Files/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesServer is the server API for Files service.
// All implementations must embed UnimplementedFilesServer
// for forward compatibility
type FilesServer interface {
	GetOssStsToken(context.Context, *emptypb.Empty) (*OssStsTokenResponse, error)
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	mustEmbedUnimplementedFilesServer()
}

// UnimplementedFilesServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServer struct {
}

func (UnimplementedFilesServer) GetOssStsToken(context.Context, *emptypb.Empty) (*OssStsTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOssStsToken not implemented")
}
func (UnimplementedFilesServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFilesServer) mustEmbedUnimplementedFilesServer() {}

// UnsafeFilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServer will
// result in compilation errors.
type UnsafeFilesServer interface {
	mustEmbedUnimplementedFilesServer()
}

func RegisterFilesServer(s grpc.ServiceRegistrar, srv FilesServer) {
	s.RegisterService(&Files_ServiceDesc, srv)
}

func _Files_GetOssStsToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).GetOssStsToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.files.v1.Files/GetOssStsToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).GetOssStsToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.files.v1.Files/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Files_ServiceDesc is the grpc.ServiceDesc for Files service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Files_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.files.v1.Files",
	HandlerType: (*FilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOssStsToken",
			Handler:    _Files_GetOssStsToken_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Files_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/files/v1/files.proto",
}
