package rpc

import (
	"blockchain-go/chain"
	v1 "blockchain-go/node/rpc/proto/v1"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
)

var _ v1.BlockServiceServer = (*BlockSrv)(nil)

type BlockApplier interface {
	ApplyBlockToState(blk chain.SigBlock) error
}

type BlockRelayer interface {
	RelayBlock(blk chain.SigBlock)
}

type BlockSrv struct {
	v1.UnimplementedBlockServiceServer
	blockStoreDir string
	//eventPub      chain.EventPublisher
	blockApplier BlockApplier
	blockRelayer BlockRelayer
}

func NewBlockSrv(
	blockStoreDir string,
	blkApplier BlockApplier,
	blkRelayer BlockRelayer,
) *BlockSrv {
	return &BlockSrv{
		blockStoreDir: blockStoreDir,
		blockApplier:  blkApplier,
		blockRelayer:  blkRelayer,
	}
}

func (b *BlockSrv) GenesisSync(_ context.Context, request *v1.GenesisSyncRequest) (*v1.GenesisSyncResponse, error) {
	jgen, err := chain.ReadGenesisBytes(b.blockStoreDir)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &v1.GenesisSyncResponse{Genesis: jgen}, nil
}

func (b *BlockSrv) BlockSync(request *v1.BlockSyncRequest, stream v1.BlockService_BlockSyncServer) error {
	blocks, closeBlocks, err := chain.ReadBlocksBytes(b.blockStoreDir)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}
	defer closeBlocks()
	num, i := int(request.Number), 1
	for err, jdlk := range blocks {
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		if i >= num {
			res := &v1.BlockSyncResponse{Block: jdlk}
			err = stream.Send(res)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
		}
		i++
	}
	return nil
}

func (b *BlockSrv) BlockReceive(stream v1.BlockService_BlockReceiveServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &v1.BlockReceiveResponse{}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		var blk chain.SigBlock
		err = json.Unmarshal(req.Block, &blk)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("<== Block receive\n%v", blk)
		err = b.blockApplier.ApplyBlockToState(blk)
		if err != nil {
			fmt.Print(err)
			continue
		}
		err = blk.Write(b.blockStoreDir)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if b.blockRelayer != nil {
			b.blockRelayer.RelayBlock(blk)
		}
		//if s.eventPub != nil {
		//	s.publishBlockAndTxs(blk)
		//}
	}
}

func (b *BlockSrv) BlockSearch(request *v1.BlockSearchRequest, stream v1.BlockService_BlockSearchServer) error {
	blocks, closeBlocks, err := chain.ReadBlocks(b.blockStoreDir)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}
	defer closeBlocks()
	prefix := strings.HasPrefix
	for err, blk := range blocks {
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		if request.Number != 0 && blk.Number == request.Number ||
			len(request.Hash) > 0 && prefix(blk.Hash().String(), request.Hash) ||
			len(request.Parent) > 0 && prefix(blk.Parent.String(), request.Parent) {
			jblk, err := json.Marshal(blk)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
			res := &v1.BlockSearchResponse{Block: jblk}
			err = stream.Send(res)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
			break
		}
	}
	return nil
}

//func (s *BlockSrv) publishBlockAndTxs(blk chain.SigBlock) {
//	jblk, _ := json.Marshal(blk)
//	event := chain.NewEvent(chain.EvBlock, "validated", jblk)
//	s.eventPub.PublishEvent(event)
//	for _, tx := range blk.Txs {
//		jtx, _ := json.Marshal(tx)
//		event := chain.NewEvent(chain.EvTx, "validated", jtx)
//		s.eventPub.PublishEvent(event)
//	}
//}
