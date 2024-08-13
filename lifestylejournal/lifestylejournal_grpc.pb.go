// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: lifestylejournal/lifestylejournal.proto

package lifestylejournal

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
	LifestyleJournalService_KeepJournal_FullMethodName             = "/taomics.praman.lifestylejournal.LifestyleJournalService/KeepJournal"
	LifestyleJournalService_EvaluateChallengeScores_FullMethodName = "/taomics.praman.lifestylejournal.LifestyleJournalService/EvaluateChallengeScores"
)

// LifestyleJournalServiceClient is the client API for LifestyleJournalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// # LifestyleJournalService
//
// Common Errors:
//   - INTERNAL (13): Server is something wrong.
//   - UNAUTHENTICATED (16): Authorization header is something wrong.
type LifestyleJournalServiceClient interface {
	// Keep a journal.
	//
	// Errors:
	//   - INVALID_ARGUMENT (3): There is an invalid argument
	KeepJournal(ctx context.Context, in *LifestyleJournalKeepingRequest, opts ...grpc.CallOption) (*LifestyleJournalKeepingResponse, error)
	// Evaluate the challenge scores based on the latest journal records.
	// It requires a minimum number of journal records to perform an evaluation.
	//
	// Errors:
	//   - FAILED_PRECONDITION (9): Not enough journal records to evaluate the
	//     challenge scores.
	EvaluateChallengeScores(ctx context.Context, in *LifestyleJournalEvaluateChallengeScoresRequest, opts ...grpc.CallOption) (*LifestyleJournalEvaluateChallengeScoresResponse, error)
}

type lifestyleJournalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLifestyleJournalServiceClient(cc grpc.ClientConnInterface) LifestyleJournalServiceClient {
	return &lifestyleJournalServiceClient{cc}
}

func (c *lifestyleJournalServiceClient) KeepJournal(ctx context.Context, in *LifestyleJournalKeepingRequest, opts ...grpc.CallOption) (*LifestyleJournalKeepingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LifestyleJournalKeepingResponse)
	err := c.cc.Invoke(ctx, LifestyleJournalService_KeepJournal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifestyleJournalServiceClient) EvaluateChallengeScores(ctx context.Context, in *LifestyleJournalEvaluateChallengeScoresRequest, opts ...grpc.CallOption) (*LifestyleJournalEvaluateChallengeScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LifestyleJournalEvaluateChallengeScoresResponse)
	err := c.cc.Invoke(ctx, LifestyleJournalService_EvaluateChallengeScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifestyleJournalServiceServer is the server API for LifestyleJournalService service.
// All implementations must embed UnimplementedLifestyleJournalServiceServer
// for forward compatibility.
//
// # LifestyleJournalService
//
// Common Errors:
//   - INTERNAL (13): Server is something wrong.
//   - UNAUTHENTICATED (16): Authorization header is something wrong.
type LifestyleJournalServiceServer interface {
	// Keep a journal.
	//
	// Errors:
	//   - INVALID_ARGUMENT (3): There is an invalid argument
	KeepJournal(context.Context, *LifestyleJournalKeepingRequest) (*LifestyleJournalKeepingResponse, error)
	// Evaluate the challenge scores based on the latest journal records.
	// It requires a minimum number of journal records to perform an evaluation.
	//
	// Errors:
	//   - FAILED_PRECONDITION (9): Not enough journal records to evaluate the
	//     challenge scores.
	EvaluateChallengeScores(context.Context, *LifestyleJournalEvaluateChallengeScoresRequest) (*LifestyleJournalEvaluateChallengeScoresResponse, error)
	mustEmbedUnimplementedLifestyleJournalServiceServer()
}

// UnimplementedLifestyleJournalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLifestyleJournalServiceServer struct{}

func (UnimplementedLifestyleJournalServiceServer) KeepJournal(context.Context, *LifestyleJournalKeepingRequest) (*LifestyleJournalKeepingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepJournal not implemented")
}
func (UnimplementedLifestyleJournalServiceServer) EvaluateChallengeScores(context.Context, *LifestyleJournalEvaluateChallengeScoresRequest) (*LifestyleJournalEvaluateChallengeScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluateChallengeScores not implemented")
}
func (UnimplementedLifestyleJournalServiceServer) mustEmbedUnimplementedLifestyleJournalServiceServer() {
}
func (UnimplementedLifestyleJournalServiceServer) testEmbeddedByValue() {}

// UnsafeLifestyleJournalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifestyleJournalServiceServer will
// result in compilation errors.
type UnsafeLifestyleJournalServiceServer interface {
	mustEmbedUnimplementedLifestyleJournalServiceServer()
}

func RegisterLifestyleJournalServiceServer(s grpc.ServiceRegistrar, srv LifestyleJournalServiceServer) {
	// If the following call pancis, it indicates UnimplementedLifestyleJournalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LifestyleJournalService_ServiceDesc, srv)
}

func _LifestyleJournalService_KeepJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifestyleJournalKeepingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleJournalServiceServer).KeepJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleJournalService_KeepJournal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleJournalServiceServer).KeepJournal(ctx, req.(*LifestyleJournalKeepingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifestyleJournalService_EvaluateChallengeScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifestyleJournalEvaluateChallengeScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleJournalServiceServer).EvaluateChallengeScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleJournalService_EvaluateChallengeScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleJournalServiceServer).EvaluateChallengeScores(ctx, req.(*LifestyleJournalEvaluateChallengeScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LifestyleJournalService_ServiceDesc is the grpc.ServiceDesc for LifestyleJournalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LifestyleJournalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taomics.praman.lifestylejournal.LifestyleJournalService",
	HandlerType: (*LifestyleJournalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeepJournal",
			Handler:    _LifestyleJournalService_KeepJournal_Handler,
		},
		{
			MethodName: "EvaluateChallengeScores",
			Handler:    _LifestyleJournalService_EvaluateChallengeScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lifestylejournal/lifestylejournal.proto",
}
