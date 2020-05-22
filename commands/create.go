package commands

import (
	"github.com/jlentink/go-transip-dyndns/internal/config"
	"github.com/jlentink/go-transip-dyndns/internal/gipify"
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/jlentink/go-transip-dyndns/internal/tld"
	"github.com/spf13/cobra"
)

// Create record with public IP
func Create(cmd *cobra.Command, args []string) {
	logger.SetVerbose(config.Get().GetBool("verbose"))
	IP, err := gipify.GetIP()
	if err != nil {
		logger.Get().Fatalf("Error getting IP address. (%s)", err.Error())
	}
	err = tld.InitTLD()
	if err != nil {
		logger.Get().Fatalf("Error accessing the API. please verify configuration (%s)", err.Error())
	}
	_, err = tld.FindRecord()
	if err == nil {
		logger.Get().Fatalf("Record already exists. Use update from now on.")
	}
	err = tld.CreateRecord(IP)
	if err != nil {
		logger.Get().Fatalf("Unable to create record. (%s)", err.Error())
	} else {
		logger.Get().Infof("Create record for %s.%s with ip %s.", config.Get().GetString("domain-entry"), config.Get().GetString("domain"), IP.IP)
	}
}
