package commands

import (
	"errors"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/transipClient"
	"time"
)

// Update record with public IP
func Update(cmd *cobra.Command, _ []string) {
	keepAlive, _ := cmd.Flags().GetBool("keep-alive")
	timeStamp := ""
	for {
		if keepAlive {
			currentTime := time.Now()
			timeStamp = currentTime.Format("2006-01-02 15:04:05") + " - "
		}
		transipClient.ClearCache()
		for _, record := range config.GetRecords() {
			entry, err := transipClient.UpdateRecord(record)
			if err != nil {
				if errors.Is(err, transipClient.NotFoundError) {
					emoji.Printf("%s:exclamation: Could not find the record to update %s.%s [%s]\n",
						timeStamp, record.Entry, record.Hostname, record.Type) // nolint: errcheck
				} else if errors.Is(err, transipClient.NotChangedError) {
					emoji.Printf("%s:+1: Record still up to date: %s.%s [%s]\n",
						timeStamp, record.Entry, record.Hostname, record.Type) // nolint: errcheck
				} else {
					emoji.Printf("%s:exclamation: Could not update the record to update %s.%s [%s] - %s\n",
						timeStamp, record.Entry, record.Hostname, record.Type, err) // nolint: errcheck
				}
			} else {
				emoji.Printf("%s:+1: Record updated %s.%s [%s] to %s\n",
					timeStamp, record.Entry, record.Hostname, record.Type, entry.Content) // nolint: errcheck
			}
		}
		if keepAlive {
			time.Sleep(config.GetUpdateFrequency())
		} else {
			break
		}
	}
}
