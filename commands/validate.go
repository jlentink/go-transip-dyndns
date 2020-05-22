package commands

import (
	"fmt"
	"github.com/jlentink/go-transip-dyndns/internal/config"
	"github.com/jlentink/go-transip-dyndns/internal/tld"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"os"
)

// Validate the setup so see all is alright
func Validate(cmd *cobra.Command, args []string) {
	fmt.Printf(" - Verify access to API.\n")
	err := tld.InitTLD()
	if err != nil {
		emoji.Printf(":exclamation: Could not connect to API (%s)\n", err.Error())                   // nolint: errcheck
		emoji.Printf("Please go to https://www.transip.nl/cp/account/api/ and create a key pair. " + // nolint: errcheck
			"Than update the configuration.\n") // nolint: errcheck
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Connection successful.\n") // nolint: errcheck
	}

	fmt.Printf(" - Verify access to domain\n") // nolint: errcheck
	dom, err := tld.FindDomain()
	if err != nil {
		emoji.Printf(":exclamation: Could not find domain (%s)\n", config.Get().GetString("domain"))   // nolint: errcheck
		emoji.Printf("Please go to https://www.transip.nl/cp/ and verify you own that domain name.\n") // nolint: errcheck
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Found domain\n")                 // nolint: errcheck
		emoji.Printf("Renewal date: %s\n", dom.RenewalDate) // nolint: errcheck
	}
	fmt.Printf("- Verify record exists to domain\n")
	entry, err := tld.FindRecord()
	if err != nil {
		emoji.Printf(":exclamation: Could not find record (%s) Create one manually or run the create command\n", err.Error()) // nolint: errcheck
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Found record\n")                               // nolint: errcheck
		emoji.Printf("Pointing to: %s (%s)\n", entry.Content, entry.Type) // nolint: errcheck
	}
}
