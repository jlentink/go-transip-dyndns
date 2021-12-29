package main

import (
	_ "embed"
	"fmt"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"go-transip-dyndns/commands"
	"os"
	"path/filepath"
)

var (
	//ApplicationVersion Application version.
	ApplicationVersion = "1.0.0"
	CommitHash         = ""
	BuildDate          = ""
	verbose            bool
	keepRunning        bool

	//go:embed example.go-transip-dyndns.toml
	ConfigTemplate string
)

func init() {

}

func main() {
	commands.ConfigTemplate = ConfigTemplate
	rootCmd := &cobra.Command{
		Use:     "go-transip-dyndns [-v]",
		Short:   "Update ip address on Transip DNS to current public ip ",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Version: ApplicationVersion,
		Run:     commands.Update,
	}
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	validateCmd := &cobra.Command{
		Use:              "validate",
		Short:            "Validate the the setup",
		Long:             "Run validation to verify the setup is correct",
		Example:          "",
		PersistentPreRun: commands.PreRun,
		Run:              commands.Validate,
	}

	initCmd := &cobra.Command{
		Use:     "init",
		Short:   "Write config file to disk",
		Long:    "Get a sample configuration file tog et started.",
		Example: "",
		Run:     commands.Init,
	}

	createCmd := &cobra.Command{
		Use:              "create",
		Short:            "One time create record for updating",
		Long:             "Create a record for this configuration",
		Example:          "",
		PersistentPreRun: commands.PreRun,
		Run:              commands.Create,
	}

	updateCmd := &cobra.Command{
		Use:     "update",
		Short:   "Update ip address on Transip DNS to current public ip (default command)",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Example: "",

		Run: commands.Update,
	}
	updateCmd.PersistentFlags().BoolVarP(&keepRunning, "keep-alive", "k", false, "keep running continuously.")

	versionCmd := &cobra.Command{
		Use:              "version",
		Short:            "Show the current version of the application",
		Long:             "Show the current version of this application.",
		Example:          "",
		PersistentPreRun: commands.PreRun,
		Run: func(cmd *cobra.Command, args []string) {
			executableName := filepath.Base(os.Args[0])
			fmt.Printf("%s: \n - Version: %s\n - Git Hash: %s\n - Build date: %s\n",
				executableName, ApplicationVersion, CommitHash, BuildDate)
		},
	}

	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(versionCmd)
	err := rootCmd.Execute() // nolint: errcheck
	if err != nil {
		emoji.Printf(":exclamation: Could not execute root command!")
		os.Exit(1)
	}
}
