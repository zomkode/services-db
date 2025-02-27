// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: protos/mysql.proto

package mysql_service

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
	MySQLService_StoreAndReplicateData_FullMethodName = "/mysql_service.MySQLService/StoreAndReplicateData"
)

// MySQLServiceClient is the client API for MySQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MySQLServiceClient interface {
	StoreAndReplicateData(ctx context.Context, in *StoreAndReplicateRequest, opts ...grpc.CallOption) (*StoreAndReplicateResponse, error)
}

type mySQLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMySQLServiceClient(cc grpc.ClientConnInterface) MySQLServiceClient {
	return &mySQLServiceClient{cc}
}

func (c *mySQLServiceClient) StoreAndReplicateData(ctx context.Context, in *StoreAndReplicateRequest, opts ...grpc.CallOption) (*StoreAndReplicateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreAndReplicateResponse)
	err := c.cc.Invoke(ctx, MySQLService_StoreAndReplicateData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLServiceServer is the server API for MySQLService service.
// All implementations must embed UnimplementedMySQLServiceServer
// for forward compatibility.
type MySQLServiceServer interface {
	StoreAndReplicateData(context.Context, *StoreAndReplicateRequest) (*StoreAndReplicateResponse, error)
	mustEmbedUnimplementedMySQLServiceServer()
}

// UnimplementedMySQLServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMySQLServiceServer struct{}

func (UnimplementedMySQLServiceServer) StoreAndReplicateData(context.Context, *StoreAndReplicateRequest) (*StoreAndReplicateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAndReplicateData not implemented")
}
func (UnimplementedMySQLServiceServer) mustEmbedUnimplementedMySQLServiceServer() {}
func (UnimplementedMySQLServiceServer) testEmbeddedByValue()                      {}

// UnsafeMySQLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MySQLServiceServer will
// result in compilation errors.
type UnsafeMySQLServiceServer interface {
	mustEmbedUnimplementedMySQLServiceServer()
}

func RegisterMySQLServiceServer(s grpc.ServiceRegistrar, srv MySQLServiceServer) {
	// If the following call pancis, it indicates UnimplementedMySQLServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MySQLService_ServiceDesc, srv)
}

func _MySQLService_StoreAndReplicateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreAndReplicateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).StoreAndReplicateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MySQLService_StoreAndReplicateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).StoreAndReplicateData(ctx, req.(*StoreAndReplicateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MySQLService_ServiceDesc is the grpc.ServiceDesc for MySQLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MySQLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mysql_service.MySQLService",
	HandlerType: (*MySQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreAndReplicateData",
			Handler:    _MySQLService_StoreAndReplicateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/mysql.proto",
}
