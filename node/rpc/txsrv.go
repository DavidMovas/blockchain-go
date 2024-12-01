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
	"path/filepath"
	"strings"
)

var _ v1.TxServiceServer = (*TxSrv)(nil)

type TxApplier interface {
	Nonce(acc chain.Address) uint64
	ApplyTx(tx chain.SigTx) error
}

type TxRelayer interface {
	RelayTx(tx chain.SigTx)
}

type TxSrv struct {
	v1.UnimplementedTxServiceServer
	keyStoreDir   string
	blockStoreDir string
	txApplier     TxApplier
	txRelayer     TxRelayer
}

func (t TxSrv) TxSign(_ context.Context, request *v1.TxSignRequest) (*v1.TxSignResponse, error) {
	path := filepath.Join(t.keyStoreDir, request.From)
	acc, err := chain.ReadAccount(path, []byte(request.Password))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	tx := chain.NewTx(chain.Address(request.From), chain.Address(request.To), request.Value, t.txApplier.Nonce(chain.Address(request.From))+1)

	stx, err := acc.SignTx(tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	jtx, err := json.Marshal(stx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.TxSignResponse{Tx: jtx}, nil
}

func (t TxSrv) TxSend(_ context.Context, request *v1.TxSendRequest) (*v1.TxSendResponse, error) {
	var tx chain.SigTx
	err := json.Unmarshal(request.Tx, &tx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = t.txApplier.ApplyTx(tx)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if t.txRelayer != nil {
		t.txRelayer.RelayTx(tx)
	}

	return &v1.TxSendResponse{Hash: tx.Hash().String()}, nil
}

func (t TxSrv) TxReceive(stream v1.TxService_TxReceiveServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&v1.TxReceiveResponse{})
		}

		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		var tx chain.SigTx
		err = json.Unmarshal(req.Tx, &tx)
		if err != nil {
			fmt.Printf("invalid tx: %v\n", err)
			continue
		}

		fmt.Printf("<== Tx received: %v\n", tx)
		err = t.txApplier.ApplyTx(tx)
		if err != nil {
			fmt.Printf("failed to apply tx: %v\n", err)
			continue
		}

		if t.txRelayer != nil {
			t.txRelayer.RelayTx(tx)
		}
	}
}

func (t TxSrv) TxSearch(request *v1.TxSearchRequest, stream v1.TxService_TxSearchServer) error {
	blocks, closeBlocks, err := chain.ReadBlocks(t.blockStoreDir)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}
	defer closeBlocks()
	prefix := strings.HasPrefix
block:
	for err, blk := range blocks {
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		for _, tx := range blk.Txs {
			if len(request.Hash) > 0 && prefix(tx.Hash().String(), request.Hash) {
				err = sendTxSearchResponse(blk, tx, stream)
				if err != nil {
					return status.Errorf(codes.Internal, err.Error())
				}
				break block
			}
			if len(request.From) > 0 && prefix(string(tx.From), request.From) ||
				len(request.To) > 0 && prefix(string(tx.To), request.To) ||
				len(request.Account) > 0 &&
					(prefix(string(tx.From), request.From) || prefix(string(tx.To), request.To)) {
				err := sendTxSearchResponse(blk, tx, stream)
				if err != nil {
					return status.Errorf(codes.Internal, err.Error())
				}
			}
		}
	}
	return nil
}

func (t TxSrv) TxProve(_ context.Context, request *v1.TxProveRequest) (*v1.TxProveResponse, error) {
	blocks, closeBlocks, err := chain.ReadBlocks(t.blockStoreDir)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	defer closeBlocks()
	for err, blk := range blocks {
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		for _, tx := range blk.Txs {
			if tx.Hash().String() == request.Hash {
				merkleTree, err := chain.MerkleHash(
					blk.Txs, chain.TxHash, chain.TxPairHash,
				)
				if err != nil {
					return nil, status.Errorf(codes.Internal, err.Error())
				}
				merkleProof, err := chain.MerkleProve(tx.Hash(), merkleTree)
				if err != nil {
					return nil, status.Errorf(codes.Internal, err.Error())
				}
				jmp, err := json.Marshal(merkleProof)
				if err != nil {
					return nil, status.Errorf(codes.Internal, err.Error())
				}
				res := &v1.TxProveResponse{MerkleProof: jmp}
				return res, nil
			}
		}
	}
	return nil, status.Errorf(
		codes.NotFound, fmt.Sprintf("transaction %v not found", request.Hash),
	)
}

func (t TxSrv) TxVerify(ctx context.Context, request *v1.TxVerifyRequest) (*v1.TxVerifyResponse, error) {
	txh, err := chain.DecodeHash(request.Hash)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	var merkleProof []chain.Proof[chain.Hash]
	err = json.Unmarshal(request.MerkleProof, &merkleProof)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	merkleRoot, err := chain.DecodeHash(request.MerkleRoot)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	valid := chain.MerkleVerify(txh, merkleProof, merkleRoot, chain.TxPairHash)
	res := &v1.TxVerifyResponse{Valid: valid}
	return res, nil
}

func sendTxSearchResponse(blk chain.SigBlock, tx chain.SigTx, stream v1.TxService_TxSearchServer) error {
	stx := chain.NewSearchTx(tx, blk.Number, blk.Hash(), blk.MerkleRoot)

	jtx, err := json.Marshal(stx)
	if err != nil {
		return err
	}
	res := &v1.TxSearchResponse{Tx: jtx}
	err = stream.Send(res)
	return err
}
