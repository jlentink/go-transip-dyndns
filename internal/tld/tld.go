package tld

import (
	"fmt"
	"github.com/jlentink/go-transip-dyndns/internal/config"
	"github.com/jlentink/go-transip-dyndns/internal/gipify"
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/repository"
)

var _transip *repository.Client

// InitTLD setup TransAPI client
func InitTLD() error {
	if _transip != nil {
		return nil
	}
	transipClient, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    config.Get().GetString("username"),
		PrivateKeyPath: config.Get().GetString("private-key"),
	})
	_transip = &transipClient
	return err
	//logger.Get().Fatalf("Could not connect to transIP API. (%s)", err.Error())

}

// CreateRecord creates a record in the domain
func CreateRecord(ip *gipify.IP) error {
	recordType := "A"

	switch ip.Type {
	case gipify.IPV4:
		recordType = "A"
	case gipify.IPV6:
		recordType = "AAAA"
	default:
		logger.Get().Fatalf("Not a valid IPv4 or IPv6 address\n")
	}

	logger.Get().Debug("Create Record..\n")
	repo := domain.Repository{Client: *_transip}
	return repo.AddDNSEntry(config.Get().GetString("domain"), domain.DNSEntry{
		Name:    config.Get().GetString("domain-entry"),
		Expire:  config.Get().GetInt("domain-ttl"),
		Type:    recordType,
		Content: ip.IP,
	})
}

// UpdateRecord updates an existing record
func UpdateRecord(ip *gipify.IP) error {
	recordType := "A"

	switch ip.Type {
	case gipify.IPV4:
		recordType = "A"
	case gipify.IPV6:
		recordType = "AAAA"
	default:
		logger.Get().Fatalf("Not a valid IPv4 or IPv6 address\n")
	}

	logger.Get().Debug("Update Record..\n")
	repo := domain.Repository{Client: *_transip}
	return repo.UpdateDNSEntry(config.Get().GetString("domain"), domain.DNSEntry{
		Name:    config.Get().GetString("domain-entry"),
		Expire:  config.Get().GetInt("domain-ttl"),
		Type:    recordType,
		Content: ip.IP,
	})
}

// FindDomain Find domain in the API
func FindDomain() (domain.Domain, error) {
	repo := domain.Repository{Client: *_transip}
	return repo.GetByDomainName(config.Get().GetString("domain"))
}

// FindRecord finds the record in the given domain
func FindRecord() (domain.DNSEntry, error) {
	repo := domain.Repository{Client: *_transip}
	entries, err := repo.GetDNSEntries(config.Get().GetString("domain"))
	if err != nil {
		return domain.DNSEntry{}, err
	}

	for _, entry := range entries {
		if entry.Name == config.Get().GetString("domain-entry") {
			return entry, nil
		}
	}
	return domain.DNSEntry{}, fmt.Errorf("could not find entry (%s)", config.Get().GetString("domain-entry"))
}
