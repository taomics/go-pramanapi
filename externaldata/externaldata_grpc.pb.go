// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: externaldata/externaldata.proto

package externaldata

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
	ExternalDataService_StoreExternalData_FullMethodName = "/taomics.praman.externaldata.ExternalDataService/StoreExternalData"
)

// ExternalDataServiceClient is the client API for ExternalDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// # ExternalDataservice
//
// Common Errors:
//   - INTERNAL (13): Server is something wrong.
//   - UNAUTHENTICATED (16): Authorization header is something wrong.
type ExternalDataServiceClient interface {
	// Store external data.
	//
	// Errors:
	//   - INVALID_ARGUMENT (3): There is an invalid argument
	StoreExternalData(ctx context.Context, in *StoreExternalDataRequest, opts ...grpc.CallOption) (*StoreExternalDataResponse, error)
}

type externalDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalDataServiceClient(cc grpc.ClientConnInterface) ExternalDataServiceClient {
	return &externalDataServiceClient{cc}
}

func (c *externalDataServiceClient) StoreExternalData(ctx context.Context, in *StoreExternalDataRequest, opts ...grpc.CallOption) (*StoreExternalDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreExternalDataResponse)
	err := c.cc.Invoke(ctx, ExternalDataService_StoreExternalData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalDataServiceServer is the server API for ExternalDataService service.
// All implementations must embed UnimplementedExternalDataServiceServer
// for forward compatibility.
//
// # ExternalDataservice
//
// Common Errors:
//   - INTERNAL (13): Server is something wrong.
//   - UNAUTHENTICATED (16): Authorization header is something wrong.
type ExternalDataServiceServer interface {
	// Store external data.
	//
	// Errors:
	//   - INVALID_ARGUMENT (3): There is an invalid argument
	StoreExternalData(context.Context, *StoreExternalDataRequest) (*StoreExternalDataResponse, error)
	mustEmbedUnimplementedExternalDataServiceServer()
}

// UnimplementedExternalDataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExternalDataServiceServer struct{}

func (UnimplementedExternalDataServiceServer) StoreExternalData(context.Context, *StoreExternalDataRequest) (*StoreExternalDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreExternalData not implemented")
}
func (UnimplementedExternalDataServiceServer) mustEmbedUnimplementedExternalDataServiceServer() {}
func (UnimplementedExternalDataServiceServer) testEmbeddedByValue()                             {}

// UnsafeExternalDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalDataServiceServer will
// result in compilation errors.
type UnsafeExternalDataServiceServer interface {
	mustEmbedUnimplementedExternalDataServiceServer()
}

func RegisterExternalDataServiceServer(s grpc.ServiceRegistrar, srv ExternalDataServiceServer) {
	// If the following call pancis, it indicates UnimplementedExternalDataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExternalDataService_ServiceDesc, srv)
}

func _ExternalDataService_StoreExternalData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreExternalDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataServiceServer).StoreExternalData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataService_StoreExternalData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataServiceServer).StoreExternalData(ctx, req.(*StoreExternalDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalDataService_ServiceDesc is the grpc.ServiceDesc for ExternalDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taomics.praman.externaldata.ExternalDataService",
	HandlerType: (*ExternalDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreExternalData",
			Handler:    _ExternalDataService_StoreExternalData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "externaldata/externaldata.proto",
}
