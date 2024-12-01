package cli

import (
	"context"
	"github.com/spf13/cobra"
)

func ChainCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "chain",
		Short:         "chain commands",
		Version:       "0.1.0",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().String("node", "", "target node address host:port")
	_ = cmd.MarkFlagRequired("node")
	cmd.AddCommand(accountCmd(ctx))
	return cmd
}
