package commands

import (
	"fmt"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/repository"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/gipify"
	"go-transip-dyndns/internal/logger"
	"go-transip-dyndns/internal/sliceUtils"
	"os"
)

var (
	transIpClient repository.Client
)

func validateApi() {
	fmt.Printf(" - Verify access to API.\n")
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:      config.GetUsername(),
		PrivateKeyReader: config.GetPrivateKey(),
	})

	if err != nil {
		emoji.Printf(":exclamation: Could not connect to API (%s)\n", err.Error())
		emoji.Printf("Please go to https://www.transip.nl/cp/account/api/ and create a key pair. " +
			"Than update the configuration.\n")
		os.Exit(1)
	} else {
		transIpClient = client
		emoji.Printf(":+1: Connection successful.\n") // nolint: errcheck
	}
}

func validateIP() {
	fmt.Printf(" - Verify IP fetch\n") // nolint: errcheck
	if config.IsIPv4Enabled() {
		ip, err := gipify.GetIPv4()
		if nil != err {
			emoji.Printf(":exclamation: Could not get IPv4 address\n", err.Error())
			os.Exit(1)
		} else {
			emoji.Printf(":+1: IPv4 fetch successful - %s.\n", ip.IP) // nolint: errcheck
		}
	} else {
		emoji.Printf(":+1: IPv4 disabled.\n") // nolint: errcheck
	}
	if config.IsIPv6Enabled() {
		ip, err := gipify.GetIPv6()
		if nil != err {
			emoji.Printf(":exclamation: Could not get IPv4 address\n", err.Error()) // nolint: errcheck
			os.Exit(1)
		} else {
			emoji.Printf(":+1: IPv6 fetch successful - %s.\n", ip.IP) // nolint: errcheck
		}
	} else {
		emoji.Printf(":+1: IPv6 disabled.\n") // nolint: errcheck
	}
}

func validateDomains() {
	errState := false
	fmt.Printf(" - Verify access to domain(s)\n") // nolint: errcheck
	repo := domain.Repository{Client: transIpClient}
	domains, err := repo.GetAll()
	if err != nil {
		emoji.Printf(":exclamation: Could not read the domains...\n", err.Error()) // nolint: errcheck
		os.Exit(1)
	}

	for _, cDomain := range config.GetDomains() {
		index := sliceUtils.ContainsDomain(domains, cDomain)
		if -1 == index {
			emoji.Printf(":exclamation: Could not find domain (%s) in your account...\n", cDomain)
			emoji.Printf("    Please go to https://www.transip.nl/cp/ and verify you own that domain name.\n")
			errState = true
		} else {
			emoji.Printf(":+1: Found domain: %s\n", cDomain)
			emoji.Printf("    Renewal date: %s\n", domains[index].RenewalDate)
			logger.Get().Debugf("Domain: %+v", domains[index])
		}
	}
	if errState {
		os.Exit(1)
	}
}

func validateRecords() {
	errState := false
	fmt.Printf(" - Verify target domain record(s)\n") // nolint: errcheck
	repo := domain.Repository{Client: transIpClient}

	for _, cRecord := range config.GetRecords() {
		logger.Get().Debugf("Config record: %+v\n", cRecord)
		entries, err := repo.GetDNSEntries(cRecord.Hostname)
		if err != nil {
			emoji.Printf(":exclamation: Could not read the entries for domain %s...\n",
				cRecord.Hostname, err.Error()) // nolint: errcheck
			errState = true
		}
		record, err := sliceUtils.FindRecord(entries, cRecord.Hostname, cRecord.Entry, cRecord.Type)
		logger.Get().Debugf("Record: %+v\n", record)
		if err != nil {
			emoji.Printf(":exclamation: Could not not find the entry %s (%s) for domain %s...\n",
				cRecord.Entry, cRecord.Type, cRecord.Hostname) // nolint: errcheck
			emoji.Printf(":exclamation: Run create to create it\n") // nolint: errcheck
			errState = true
		} else {
			emoji.Printf(":+1: Found record: %s.%s (%s)\n",
				cRecord.Entry, cRecord.Hostname, cRecord.Type) // nolint: errcheck
			logger.Get().Debugf("Record found details: %+v\n", cRecord)
		}
	}
	if errState {
		os.Exit(1)
	}

}

// Validate the setup so see all is alright
func Validate(cmd *cobra.Command, args []string) {
	validateApi()
	validateIP()
	validateDomains()
	validateRecords()
}
