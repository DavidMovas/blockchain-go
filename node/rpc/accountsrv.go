package rpc

import (
	"blockchain-go/chain"
	v1 "blockchain-go/node/rpc/proto/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.AccountServiceServer = (*AccountSrv)(nil)

type BalanceChecker interface {
	Balance(address chain.Address) (uint64, bool)
}

type AccountSrv struct {
	v1.UnimplementedAccountServiceServer
	keyStoreDir    string
	balanceChecker BalanceChecker
}

func NewAccountSrv(keyStoreDir string, balanceChecker BalanceChecker) *AccountSrv {
	return &AccountSrv{keyStoreDir: keyStoreDir, balanceChecker: balanceChecker}
}

func (a *AccountSrv) AccountCreate(_ context.Context, request *v1.AccountCreateRequest) (*v1.AccountCreateResponse, error) {
	pass := []byte(request.Password)
	if len(pass) < 5 {
		return nil, status.Errorf(codes.InvalidArgument, "password too short")
	}

	acc, err := chain.NewAccount()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	err = acc.Write(a.keyStoreDir, pass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	return &v1.AccountCreateResponse{Address: string(acc.Address())}, nil
}

func (a *AccountSrv) AccountBalance(_ context.Context, request *v1.AccountBalanceRequest) (*v1.AccountBalanceResponse, error) {
	acc := request.Address
	balance, ok := a.balanceChecker.Balance(chain.Address(acc))
	if !ok {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	return &v1.AccountBalanceResponse{Balance: balance}, nil
}
