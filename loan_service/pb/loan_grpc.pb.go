// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: loan.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LoanService_IssueLoan_FullMethodName  = "/LoanService/IssueLoan"
	LoanService_ReturnLoan_FullMethodName = "/LoanService/ReturnLoan"
)

// LoanServiceClient is the client API for LoanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoanServiceClient interface {
	IssueLoan(ctx context.Context, in *IssueLoanRequest, opts ...grpc.CallOption) (*IssueLoanResponse, error)
	ReturnLoan(ctx context.Context, in *ReturnLoanRequest, opts ...grpc.CallOption) (*ReturnLoanResponse, error)
}

type loanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoanServiceClient(cc grpc.ClientConnInterface) LoanServiceClient {
	return &loanServiceClient{cc}
}

func (c *loanServiceClient) IssueLoan(ctx context.Context, in *IssueLoanRequest, opts ...grpc.CallOption) (*IssueLoanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IssueLoanResponse)
	err := c.cc.Invoke(ctx, LoanService_IssueLoan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) ReturnLoan(ctx context.Context, in *ReturnLoanRequest, opts ...grpc.CallOption) (*ReturnLoanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReturnLoanResponse)
	err := c.cc.Invoke(ctx, LoanService_ReturnLoan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoanServiceServer is the server API for LoanService service.
// All implementations must embed UnimplementedLoanServiceServer
// for forward compatibility
type LoanServiceServer interface {
	IssueLoan(context.Context, *IssueLoanRequest) (*IssueLoanResponse, error)
	ReturnLoan(context.Context, *ReturnLoanRequest) (*ReturnLoanResponse, error)
	mustEmbedUnimplementedLoanServiceServer()
}

// UnimplementedLoanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoanServiceServer struct {
}

func (UnimplementedLoanServiceServer) IssueLoan(context.Context, *IssueLoanRequest) (*IssueLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueLoan not implemented")
}
func (UnimplementedLoanServiceServer) ReturnLoan(context.Context, *ReturnLoanRequest) (*ReturnLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnLoan not implemented")
}
func (UnimplementedLoanServiceServer) mustEmbedUnimplementedLoanServiceServer() {}

// UnsafeLoanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoanServiceServer will
// result in compilation errors.
type UnsafeLoanServiceServer interface {
	mustEmbedUnimplementedLoanServiceServer()
}

func RegisterLoanServiceServer(s grpc.ServiceRegistrar, srv LoanServiceServer) {
	s.RegisterService(&LoanService_ServiceDesc, srv)
}

func _LoanService_IssueLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueLoanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).IssueLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoanService_IssueLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).IssueLoan(ctx, req.(*IssueLoanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_ReturnLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnLoanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).ReturnLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoanService_ReturnLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).ReturnLoan(ctx, req.(*ReturnLoanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoanService_ServiceDesc is the grpc.ServiceDesc for LoanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LoanService",
	HandlerType: (*LoanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueLoan",
			Handler:    _LoanService_IssueLoan_Handler,
		},
		{
			MethodName: "ReturnLoan",
			Handler:    _LoanService_ReturnLoan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loan.proto",
}
