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

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockServiceClient interface {
	GenesisSync(ctx context.Context, in *GenesisSyncRequest, opts ...grpc.CallOption) (*GenesisSyncResponse, error)
	BlockSync(ctx context.Context, in *BlockSyncRequest, opts ...grpc.CallOption) (BlockService_BlockSyncClient, error)
	BlockReceive(ctx context.Context, opts ...grpc.CallOption) (BlockService_BlockReceiveClient, error)
	BlockSearch(ctx context.Context, in *BlockSearchRequest, opts ...grpc.CallOption) (BlockService_BlockSearchClient, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GenesisSync(ctx context.Context, in *GenesisSyncRequest, opts ...grpc.CallOption) (*GenesisSyncResponse, error) {
	out := new(GenesisSyncResponse)
	err := c.cc.Invoke(ctx, "/rpc.v1.BlockService/GenesisSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) BlockSync(ctx context.Context, in *BlockSyncRequest, opts ...grpc.CallOption) (BlockService_BlockSyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[0], "/rpc.v1.BlockService/BlockSync", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceBlockSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_BlockSyncClient interface {
	Recv() (*BlockSyncResponse, error)
	grpc.ClientStream
}

type blockServiceBlockSyncClient struct {
	grpc.ClientStream
}

func (x *blockServiceBlockSyncClient) Recv() (*BlockSyncResponse, error) {
	m := new(BlockSyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockServiceClient) BlockReceive(ctx context.Context, opts ...grpc.CallOption) (BlockService_BlockReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[1], "/rpc.v1.BlockService/BlockReceive", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceBlockReceiveClient{stream}
	return x, nil
}

type BlockService_BlockReceiveClient interface {
	Send(*BlockReceiveRequest) error
	CloseAndRecv() (*BlockReceiveResponse, error)
	grpc.ClientStream
}

type blockServiceBlockReceiveClient struct {
	grpc.ClientStream
}

func (x *blockServiceBlockReceiveClient) Send(m *BlockReceiveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockServiceBlockReceiveClient) CloseAndRecv() (*BlockReceiveResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BlockReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockServiceClient) BlockSearch(ctx context.Context, in *BlockSearchRequest, opts ...grpc.CallOption) (BlockService_BlockSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[2], "/rpc.v1.BlockService/BlockSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceBlockSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_BlockSearchClient interface {
	Recv() (*BlockSearchResponse, error)
	grpc.ClientStream
}

type blockServiceBlockSearchClient struct {
	grpc.ClientStream
}

func (x *blockServiceBlockSearchClient) Recv() (*BlockSearchResponse, error) {
	m := new(BlockSearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockServiceServer is the server API for BlockService service.
// All implementations must embed UnimplementedBlockServiceServer
// for forward compatibility
type BlockServiceServer interface {
	GenesisSync(context.Context, *GenesisSyncRequest) (*GenesisSyncResponse, error)
	BlockSync(*BlockSyncRequest, BlockService_BlockSyncServer) error
	BlockReceive(BlockService_BlockReceiveServer) error
	BlockSearch(*BlockSearchRequest, BlockService_BlockSearchServer) error
	mustEmbedUnimplementedBlockServiceServer()
}

// UnimplementedBlockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (UnimplementedBlockServiceServer) GenesisSync(context.Context, *GenesisSyncRequest) (*GenesisSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenesisSync not implemented")
}
func (UnimplementedBlockServiceServer) BlockSync(*BlockSyncRequest, BlockService_BlockSyncServer) error {
	return status.Errorf(codes.Unimplemented, "method BlockSync not implemented")
}
func (UnimplementedBlockServiceServer) BlockReceive(BlockService_BlockReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method BlockReceive not implemented")
}
func (UnimplementedBlockServiceServer) BlockSearch(*BlockSearchRequest, BlockService_BlockSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method BlockSearch not implemented")
}
func (UnimplementedBlockServiceServer) mustEmbedUnimplementedBlockServiceServer() {}

// UnsafeBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockServiceServer will
// result in compilation errors.
type UnsafeBlockServiceServer interface {
	mustEmbedUnimplementedBlockServiceServer()
}

func RegisterBlockServiceServer(s grpc.ServiceRegistrar, srv BlockServiceServer) {
	s.RegisterService(&BlockService_ServiceDesc, srv)
}

func _BlockService_GenesisSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenesisSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GenesisSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.v1.BlockService/GenesisSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GenesisSync(ctx, req.(*GenesisSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockService_BlockSync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockSyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).BlockSync(m, &blockServiceBlockSyncServer{stream})
}

type BlockService_BlockSyncServer interface {
	Send(*BlockSyncResponse) error
	grpc.ServerStream
}

type blockServiceBlockSyncServer struct {
	grpc.ServerStream
}

func (x *blockServiceBlockSyncServer) Send(m *BlockSyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockService_BlockReceive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockServiceServer).BlockReceive(&blockServiceBlockReceiveServer{stream})
}

type BlockService_BlockReceiveServer interface {
	SendAndClose(*BlockReceiveResponse) error
	Recv() (*BlockReceiveRequest, error)
	grpc.ServerStream
}

type blockServiceBlockReceiveServer struct {
	grpc.ServerStream
}

func (x *blockServiceBlockReceiveServer) SendAndClose(m *BlockReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockServiceBlockReceiveServer) Recv() (*BlockReceiveRequest, error) {
	m := new(BlockReceiveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockService_BlockSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).BlockSearch(m, &blockServiceBlockSearchServer{stream})
}

type BlockService_BlockSearchServer interface {
	Send(*BlockSearchResponse) error
	grpc.ServerStream
}

type blockServiceBlockSearchServer struct {
	grpc.ServerStream
}

func (x *blockServiceBlockSearchServer) Send(m *BlockSearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// BlockService_ServiceDesc is the grpc.ServiceDesc for BlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.v1.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenesisSync",
			Handler:    _BlockService_GenesisSync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BlockSync",
			Handler:       _BlockService_BlockSync_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BlockReceive",
			Handler:       _BlockService_BlockReceive_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BlockSearch",
			Handler:       _BlockService_BlockSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "block.proto",
}