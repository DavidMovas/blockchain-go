package cli

import (
	v1 "blockchain-go/node/rpc/proto/v1"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func accountCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "account commands",
	}
	cmd.AddCommand(accountCreateCmd(ctx), accountBalanceCmd(ctx))
	return cmd
}

func grpcAccountCreate(ctx context.Context, addr, ownerPass string) (string, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}

	defer func() {
		_ = conn.Close()
	}()

	client := v1.NewAccountServiceClient(conn)
	req := &v1.AccountCreateRequest{Password: ownerPass}
	res, err := client.AccountCreate(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Address, nil
}

func accountCreateCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create account",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, _ := cmd.Flags().GetString("node")
			ownerPass, _ := cmd.Flags().GetString("ownerpass")
			acc, err := grpcAccountCreate(ctx, addr, ownerPass)
			if err != nil {
				return err
			}
			fmt.Printf("Account: %s\n", acc)
			return nil
		},
	}

	cmd.Flags().String("ownerpass", "password", "owner password")
	return cmd
}

func grpcAccountBalance(ctx context.Context, addr, acc string) (uint64, error) {
	conn, err := grpc.NewClient(
		addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = conn.Close()
	}()

	client := v1.NewAccountServiceClient(conn)
	req := &v1.AccountBalanceRequest{Address: acc}
	res, err := client.AccountBalance(ctx, req)
	if err != nil {
		return 0, err
	}
	return res.Balance, nil
}

func accountBalanceCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Returns the balance of an account",
		RunE: func(cmd *cobra.Command, _ []string) error {
			addr, _ := cmd.Flags().GetString("node")
			acc, _ := cmd.Flags().GetString("account")
			balance, err := grpcAccountBalance(ctx, addr, acc)
			if err != nil {
				return err
			}
			fmt.Printf("Account: %v: %v\n", acc, balance)
			return nil
		},
	}
	cmd.Flags().String("account", "acc", "account address")
	return cmd
}
