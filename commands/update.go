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
	/*
		logger.SetVerbose(config.Get().GetBool("verbose"))
		IP, err := gipify.GetIPv4()
		if err != nil {
			logger.Get().Fatalf("Error getting IP address. (%s)", err.Error())
		}
		err = tld.InitTLD(config.Get().GetString("username"), config.Get().GetString("private-key"))
		if err != nil {
			logger.Get().Fatalf("Error accessing the API. please verify configuration (%s)", err.Error())
		}

		tld.SetRecordInformation(
			config.Get().GetString("domain"),
			config.Get().GetString("domain-entry"),
			config.Get().GetInt("domain-ttl"),
		)

		changed, err := tld.UpdateRecord(IP)
		if err != nil {
			logger.Get().Fatalf("Unable to create record. (%s)", err.Error())
		} else if changed {
			logger.Get().Infof("Updating record for %s.%s with ip %s.", config.Get().GetString("domain-entry"), config.Get().GetString("domain"), IP.IP)
		} else {
			logger.Get().Infof("Record is up to date %s.%s with ip %s.", config.Get().GetString("domain-entry"), config.Get().GetString("domain"), IP.IP)
		}
	*/
}
