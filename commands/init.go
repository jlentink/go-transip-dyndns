package commands

import (
	"errors"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var ConfigTemplate string

// Init dump a config file.
func Init(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("go-transip-dyndns.toml"); errors.Is(err, os.ErrNotExist) {
		err = ioutil.WriteFile("go-transip-dyndns.toml", []byte(ConfigTemplate), 0640)
		if err != nil {
			emoji.Printf(":exclamation: Config file could not be written. exiting...") // nolint: errcheck
			os.Exit(1)
		}
		emoji.Printf(":+1: Written config file.")
	} else {
		emoji.Printf(":exclamation: Config file already exists. exiting... " +
			"Remove the old one to write a fresh one.") // nolint: errcheck
		os.Exit(1)
	}
}
