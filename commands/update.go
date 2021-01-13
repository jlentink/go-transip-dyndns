package commands

import (
	"github.com/jlentink/go-transip-dyndns/internal/config"
	"github.com/jlentink/go-transip-dyndns/internal/gipify"
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/jlentink/go-transip-dyndns/internal/tld"
	"github.com/spf13/cobra"
)

// Update record with public IP
func Update(cmd *cobra.Command, args []string) {
	logger.SetVerbose(config.Get().GetBool("verbose"))
	IP, err := gipify.GetIP()
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
}
