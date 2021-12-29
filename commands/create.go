package commands

import (
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/transipClient"
)

// Create record with public IP
func Create(_ *cobra.Command, _ []string) {
	transipClient.ClearCache()
	for _, record := range config.GetRecords() {
		dnsRecord, err := transipClient.FindRecord(record.Hostname, record.Type, record.Entry)
		if err != nil && err.Error() == "did not find record" {
			err = transipClient.CreateRecord(record)
			//nolint:errcheck
			emoji.Printf(":+1: created record! %s.%s [%s])\n",
				record.Entry, record.Hostname, record.Type)
			if err != nil {
				// nolint:errcheck
				emoji.Printf(":exclamation: Error creating record (%s)\n", err.Error())
			}
		} else if err != nil {
			//nolint:errcheck
			emoji.Printf(":exclamation: Error getting records for domain (%s) \n", record.Hostname)
		}
		if dnsRecord != nil {
			//nolint:errcheck
			emoji.Printf(":+1: Found record not creating... (%s.%s [%s])\n",
				record.Entry, record.Hostname, record.Type)
			//nolint:errcheck
			emoji.Printf("    To update these records use update.\n\n")
		}
	}
}
