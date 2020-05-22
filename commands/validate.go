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
		emoji.Printf(":exclamation: Could not connect to API (%s)\n", err.Error())
		emoji.Printf("Please go to https://www.transip.nl/cp/account/api/ and create a key pair. " +
			"Than update the configuration.\n")
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Connection successful.\n")
	}

	fmt.Printf(" - Verify access to domain\n")
	dom, err := tld.FindDomain()
	if err != nil {
		emoji.Printf(":exclamation: Could not find domain (%s)\n", config.Get().GetString("domain"))
		emoji.Printf("Please go to https://www.transip.nl/cp/ and verify you own that domain name.\n")
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Found domain\n")
		emoji.Printf("Renewal date: %s\n", dom.RenewalDate)
	}
	fmt.Printf("- Verify record exists to domain\n")
	entry, err := tld.FindRecord()
	if err != nil {
		emoji.Printf(":exclamation: Could not find record (%s) Create one manually or run the create command\n", err.Error())
		os.Exit(1)
	} else {
		emoji.Printf(":+1: Found record\n")
		emoji.Printf("Pointing to: %s (%s)\n", entry.Content, entry.Type)
	}
}
