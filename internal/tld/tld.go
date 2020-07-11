package tld

import (
	"fmt"
	"github.com/jlentink/go-transip-dyndns/internal/gipify"
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/repository"
)

var _transip *repository.Client
var _tld string
var _entry string
var _ttl int

// SetDomainName Sets Domain name to use for up switching
func SetDomainName(domainName string) {
	_tld = domainName
}

// SetEntry Sets Entry to update in domain
func SetEntry(entry string) {
	_entry = entry
}

// SetTTL Sets TTL for entry
func SetTTL(ttl int) {
	_ttl = ttl
}

// SetRecordInformation Sets all needed Record information in one go.
func SetRecordInformation(domainName, entry string, ttl int) {
	SetDomainName(domainName)
	SetEntry(entry)
	SetTTL(ttl)
}

// InitTLD setup TransAPI client
func InitTLD(accountName, privateKeyPath string) error {
	if _transip != nil {
		return nil
	}
	transipClient, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    accountName,
		PrivateKeyPath: privateKeyPath,
	})
	_transip = &transipClient
	return err
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
	return repo.AddDNSEntry(_tld, domain.DNSEntry{
		Name:    _entry,
		Expire:  _ttl,
		Type:    recordType,
		Content: ip.IP,
	})
}

// UpdateRecord updates an existing record
func UpdateRecord(ip *gipify.IP) (bool, error) {
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

	change, err := isIPChanged(ip)
	if err != nil {
		return false, err
	}

	if !change {
		logger.Get().Debug("IP is unchanged...\n")
		return false, nil
	}

	return true, repo.UpdateDNSEntry(_tld, domain.DNSEntry{
		Name:    _entry,
		Expire:  _ttl,
		Type:    recordType,
		Content: ip.IP,
	})
}

func isIPChanged(ip *gipify.IP) (bool, error) {
	record, err := FindRecord()

	if err != nil {
		return false, err
	}

	if record.Content != ip.IP {
		return true, nil
	}
	return false, nil
}

// FindDomain Find domain in the API
func FindDomain() (domain.Domain, error) {
	repo := domain.Repository{Client: *_transip}
	return repo.GetByDomainName(_tld)
}

// FindRecord finds the record in the given domain
func FindRecord() (domain.DNSEntry, error) {
	repo := domain.Repository{Client: *_transip}
	entries, err := repo.GetDNSEntries(_tld)
	if err != nil {
		return domain.DNSEntry{}, err
	}

	for _, entry := range entries {
		if entry.Name == _entry {
			return entry, nil
		}
	}
	return domain.DNSEntry{}, fmt.Errorf("could not find entry (%s)", _entry)
}
