package commands

import (
	"github.com/spf13/cobra"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/logger"
)

func PreRun(cmd *cobra.Command, args []string) {
	config.Get().BindPFlag("general.verbose", cmd.Flags().Lookup("verbose")) // nolint: errcheck
	logger.SetVerbose(config.Get().GetBool("general.verbose"))
}
