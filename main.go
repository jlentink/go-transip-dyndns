package main

import (
	"github.com/jlentink/go-transip-dyndns/commands"
	"github.com/jlentink/go-transip-dyndns/internal/config"
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "dyndns",
		Short:   "Update ip address on Transip DNS to current public ip ",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Version: "1.0.0",
		Run:     commands.Update,
	}
	userName string
	keyFile  string
	domain   string
	verbose  bool
)

func init() {
	logger.Init()
	config.Init()
	rootCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "Transip username")
	rootCmd.PersistentFlags().StringVarP(&keyFile, "key", "k", "", "Transip password key file")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "The domain (A|AAAA record) for which the ip must be set. (including optional subdomain)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Log level verbose")

	rootCmd.AddCommand(&cobra.Command{
		Use:     "validate",
		Short:   "Validate the the setup",
		Long:    "Run validation to verify the setup is correct",
		Example: "",
		Run:     commands.Validate,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:     "create",
		Short:   "One time create record for updating",
		Long:    "Create a record for this configuration",
		Example: "",
		Run:     commands.Create,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:     "update",
		Short:   "Update ip address on Transip DNS to current public ip (default command)",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Example: "",
		Run:     commands.Update,
	})

	config.Get().BindPFlag("username", rootCmd.PersistentFlags().Lookup("username")) // nolint: errcheck
	config.Get().BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))           // nolint: errcheck
	config.Get().BindPFlag("domain", rootCmd.PersistentFlags().Lookup("domain"))     // nolint: errcheck
	config.Get().BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))   // nolint: errcheck
}

func main() {
	rootCmd.Execute() // nolint: errcheck
}
