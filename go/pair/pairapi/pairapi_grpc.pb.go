// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/pairapi.proto

package pairapi

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

// PairAPIClient is the client API for PairAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PairAPIClient interface {
	// Account //
	CreateAccount(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*AccountCreateResp, error)
	GetAccount(ctx context.Context, in *AccountGetReq, opts ...grpc.CallOption) (*AccountGetResp, error)
	UpdateAccount(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*AccountUpdateResp, error)
	// Pair //
	StartPair(ctx context.Context, in *PairStartReq, opts ...grpc.CallOption) (*PairStartResp, error)
	FinishPair(ctx context.Context, in *PairFinishReq, opts ...grpc.CallOption) (*PairFinishResp, error)
}

type pairAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPairAPIClient(cc grpc.ClientConnInterface) PairAPIClient {
	return &pairAPIClient{cc}
}

func (c *pairAPIClient) CreateAccount(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*AccountCreateResp, error) {
	out := new(AccountCreateResp)
	err := c.cc.Invoke(ctx, "/pairapi.PairAPI/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairAPIClient) GetAccount(ctx context.Context, in *AccountGetReq, opts ...grpc.CallOption) (*AccountGetResp, error) {
	out := new(AccountGetResp)
	err := c.cc.Invoke(ctx, "/pairapi.PairAPI/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairAPIClient) UpdateAccount(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*AccountUpdateResp, error) {
	out := new(AccountUpdateResp)
	err := c.cc.Invoke(ctx, "/pairapi.PairAPI/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairAPIClient) StartPair(ctx context.Context, in *PairStartReq, opts ...grpc.CallOption) (*PairStartResp, error) {
	out := new(PairStartResp)
	err := c.cc.Invoke(ctx, "/pairapi.PairAPI/StartPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairAPIClient) FinishPair(ctx context.Context, in *PairFinishReq, opts ...grpc.CallOption) (*PairFinishResp, error) {
	out := new(PairFinishResp)
	err := c.cc.Invoke(ctx, "/pairapi.PairAPI/FinishPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PairAPIServer is the server API for PairAPI service.
// All implementations must embed UnimplementedPairAPIServer
// for forward compatibility
type PairAPIServer interface {
	// Account //
	CreateAccount(context.Context, *AccountCreateReq) (*AccountCreateResp, error)
	GetAccount(context.Context, *AccountGetReq) (*AccountGetResp, error)
	UpdateAccount(context.Context, *AccountUpdateReq) (*AccountUpdateResp, error)
	// Pair //
	StartPair(context.Context, *PairStartReq) (*PairStartResp, error)
	FinishPair(context.Context, *PairFinishReq) (*PairFinishResp, error)
	mustEmbedUnimplementedPairAPIServer()
}

// UnimplementedPairAPIServer must be embedded to have forward compatible implementations.
type UnimplementedPairAPIServer struct {
}

func (UnimplementedPairAPIServer) CreateAccount(context.Context, *AccountCreateReq) (*AccountCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedPairAPIServer) GetAccount(context.Context, *AccountGetReq) (*AccountGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedPairAPIServer) UpdateAccount(context.Context, *AccountUpdateReq) (*AccountUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedPairAPIServer) StartPair(context.Context, *PairStartReq) (*PairStartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPair not implemented")
}
func (UnimplementedPairAPIServer) FinishPair(context.Context, *PairFinishReq) (*PairFinishResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishPair not implemented")
}
func (UnimplementedPairAPIServer) mustEmbedUnimplementedPairAPIServer() {}

// UnsafePairAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PairAPIServer will
// result in compilation errors.
type UnsafePairAPIServer interface {
	mustEmbedUnimplementedPairAPIServer()
}

func RegisterPairAPIServer(s grpc.ServiceRegistrar, srv PairAPIServer) {
	s.RegisterService(&PairAPI_ServiceDesc, srv)
}

func _PairAPI_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairAPIServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairapi.PairAPI/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairAPIServer).CreateAccount(ctx, req.(*AccountCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairAPI_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairAPIServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairapi.PairAPI/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairAPIServer).GetAccount(ctx, req.(*AccountGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairAPI_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairAPIServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairapi.PairAPI/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairAPIServer).UpdateAccount(ctx, req.(*AccountUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairAPI_StartPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairStartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairAPIServer).StartPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairapi.PairAPI/StartPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairAPIServer).StartPair(ctx, req.(*PairStartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairAPI_FinishPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairFinishReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairAPIServer).FinishPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairapi.PairAPI/FinishPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairAPIServer).FinishPair(ctx, req.(*PairFinishReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PairAPI_ServiceDesc is the grpc.ServiceDesc for PairAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PairAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pairapi.PairAPI",
	HandlerType: (*PairAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _PairAPI_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _PairAPI_GetAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _PairAPI_UpdateAccount_Handler,
		},
		{
			MethodName: "StartPair",
			Handler:    _PairAPI_StartPair_Handler,
		},
		{
			MethodName: "FinishPair",
			Handler:    _PairAPI_FinishPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pairapi.proto",
}
