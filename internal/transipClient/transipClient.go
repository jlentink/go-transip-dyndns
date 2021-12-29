package transipClient

import (
	"github.com/kyokomi/emoji"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/gipify"
	"os"
)

var (
	_client repository.Client
	_IPv4   *gipify.IP
	_IPv6   *gipify.IP
)

func init() {
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:      config.GetUsername(),
		PrivateKeyReader: config.GetPrivateKey(),
	})

	if err != nil {
		emoji.Printf(":exclamation: Could not connect to API (%s)\n", err.Error())                   // nolint: errcheck
		emoji.Printf("Please go to https://www.transip.nl/cp/account/api/ and create a key pair. " + // nolint: errcheck
			"Than update the configuration.\n") // nolint: errcheck
		os.Exit(1)
	}
	_client = client
}

func ClearCache() {
	_IPv4 = nil
	_IPv6 = nil
}

func GetIPv4() (*gipify.IP, error) {
	if _IPv4 != nil {
		return _IPv4, nil
	}
	ip, err := gipify.GetIPv4()
	if err != nil {
		return nil, err
	}
	_IPv4 = ip
	return ip, nil
}

func GetIPv6() (*gipify.IP, error) {
	if _IPv6 != nil {
		return _IPv6, nil
	}
	ip, err := gipify.GetIPv6()
	if err != nil {
		return nil, err
	}
	_IPv6 = ip
	return ip, nil
}
