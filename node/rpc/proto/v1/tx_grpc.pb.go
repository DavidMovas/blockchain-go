// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// TxServiceClient is the client API for TxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxServiceClient interface {
	TxSign(ctx context.Context, in *TxSignRequest, opts ...grpc.CallOption) (*TxSignResponse, error)
	TxSend(ctx context.Context, in *TxSendRequest, opts ...grpc.CallOption) (*TxSendResponse, error)
	TxReceive(ctx context.Context, opts ...grpc.CallOption) (TxService_TxReceiveClient, error)
	TxSearch(ctx context.Context, in *TxSearchRequest, opts ...grpc.CallOption) (TxService_TxSearchClient, error)
	TxProve(ctx context.Context, in *TxProveRequest, opts ...grpc.CallOption) (*TxProveResponse, error)
	TxVerify(ctx context.Context, in *TxVerifyRequest, opts ...grpc.CallOption) (*TxVerifyResponse, error)
}

type txServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTxServiceClient(cc grpc.ClientConnInterface) TxServiceClient {
	return &txServiceClient{cc}
}

func (c *txServiceClient) TxSign(ctx context.Context, in *TxSignRequest, opts ...grpc.CallOption) (*TxSignResponse, error) {
	out := new(TxSignResponse)
	err := c.cc.Invoke(ctx, "/rpc.v1.TxService/TxSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) TxSend(ctx context.Context, in *TxSendRequest, opts ...grpc.CallOption) (*TxSendResponse, error) {
	out := new(TxSendResponse)
	err := c.cc.Invoke(ctx, "/rpc.v1.TxService/TxSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) TxReceive(ctx context.Context, opts ...grpc.CallOption) (TxService_TxReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &TxService_ServiceDesc.Streams[0], "/rpc.v1.TxService/TxReceive", opts...)
	if err != nil {
		return nil, err
	}
	x := &txServiceTxReceiveClient{stream}
	return x, nil
}

type TxService_TxReceiveClient interface {
	Send(*TxReceiveRequest) error
	CloseAndRecv() (*TxReceiveResponse, error)
	grpc.ClientStream
}

type txServiceTxReceiveClient struct {
	grpc.ClientStream
}

func (x *txServiceTxReceiveClient) Send(m *TxReceiveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *txServiceTxReceiveClient) CloseAndRecv() (*TxReceiveResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TxReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *txServiceClient) TxSearch(ctx context.Context, in *TxSearchRequest, opts ...grpc.CallOption) (TxService_TxSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &TxService_ServiceDesc.Streams[1], "/rpc.v1.TxService/TxSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &txServiceTxSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TxService_TxSearchClient interface {
	Recv() (*TxSearchResponse, error)
	grpc.ClientStream
}

type txServiceTxSearchClient struct {
	grpc.ClientStream
}

func (x *txServiceTxSearchClient) Recv() (*TxSearchResponse, error) {
	m := new(TxSearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *txServiceClient) TxProve(ctx context.Context, in *TxProveRequest, opts ...grpc.CallOption) (*TxProveResponse, error) {
	out := new(TxProveResponse)
	err := c.cc.Invoke(ctx, "/rpc.v1.TxService/TxProve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) TxVerify(ctx context.Context, in *TxVerifyRequest, opts ...grpc.CallOption) (*TxVerifyResponse, error) {
	out := new(TxVerifyResponse)
	err := c.cc.Invoke(ctx, "/rpc.v1.TxService/TxVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxServiceServer is the server API for TxService service.
// All implementations must embed UnimplementedTxServiceServer
// for forward compatibility
type TxServiceServer interface {
	TxSign(context.Context, *TxSignRequest) (*TxSignResponse, error)
	TxSend(context.Context, *TxSendRequest) (*TxSendResponse, error)
	TxReceive(TxService_TxReceiveServer) error
	TxSearch(*TxSearchRequest, TxService_TxSearchServer) error
	TxProve(context.Context, *TxProveRequest) (*TxProveResponse, error)
	TxVerify(context.Context, *TxVerifyRequest) (*TxVerifyResponse, error)
	mustEmbedUnimplementedTxServiceServer()
}

// UnimplementedTxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTxServiceServer struct {
}

func (UnimplementedTxServiceServer) TxSign(context.Context, *TxSignRequest) (*TxSignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxSign not implemented")
}
func (UnimplementedTxServiceServer) TxSend(context.Context, *TxSendRequest) (*TxSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxSend not implemented")
}
func (UnimplementedTxServiceServer) TxReceive(TxService_TxReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method TxReceive not implemented")
}
func (UnimplementedTxServiceServer) TxSearch(*TxSearchRequest, TxService_TxSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method TxSearch not implemented")
}
func (UnimplementedTxServiceServer) TxProve(context.Context, *TxProveRequest) (*TxProveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxProve not implemented")
}
func (UnimplementedTxServiceServer) TxVerify(context.Context, *TxVerifyRequest) (*TxVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxVerify not implemented")
}
func (UnimplementedTxServiceServer) mustEmbedUnimplementedTxServiceServer() {}

// UnsafeTxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxServiceServer will
// result in compilation errors.
type UnsafeTxServiceServer interface {
	mustEmbedUnimplementedTxServiceServer()
}

func RegisterTxServiceServer(s grpc.ServiceRegistrar, srv TxServiceServer) {
	s.RegisterService(&TxService_ServiceDesc, srv)
}

func _TxService_TxSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).TxSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.v1.TxService/TxSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).TxSign(ctx, req.(*TxSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_TxSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).TxSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.v1.TxService/TxSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).TxSend(ctx, req.(*TxSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_TxReceive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TxServiceServer).TxReceive(&txServiceTxReceiveServer{stream})
}

type TxService_TxReceiveServer interface {
	SendAndClose(*TxReceiveResponse) error
	Recv() (*TxReceiveRequest, error)
	grpc.ServerStream
}

type txServiceTxReceiveServer struct {
	grpc.ServerStream
}

func (x *txServiceTxReceiveServer) SendAndClose(m *TxReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *txServiceTxReceiveServer) Recv() (*TxReceiveRequest, error) {
	m := new(TxReceiveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TxService_TxSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TxServiceServer).TxSearch(m, &txServiceTxSearchServer{stream})
}

type TxService_TxSearchServer interface {
	Send(*TxSearchResponse) error
	grpc.ServerStream
}

type txServiceTxSearchServer struct {
	grpc.ServerStream
}

func (x *txServiceTxSearchServer) Send(m *TxSearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TxService_TxProve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxProveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).TxProve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.v1.TxService/TxProve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).TxProve(ctx, req.(*TxProveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_TxVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).TxVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.v1.TxService/TxVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).TxVerify(ctx, req.(*TxVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxService_ServiceDesc is the grpc.ServiceDesc for TxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.v1.TxService",
	HandlerType: (*TxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TxSign",
			Handler:    _TxService_TxSign_Handler,
		},
		{
			MethodName: "TxSend",
			Handler:    _TxService_TxSend_Handler,
		},
		{
			MethodName: "TxProve",
			Handler:    _TxService_TxProve_Handler,
		},
		{
			MethodName: "TxVerify",
			Handler:    _TxService_TxVerify_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TxReceive",
			Handler:       _TxService_TxReceive_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TxSearch",
			Handler:       _TxService_TxSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tx.proto",
}
