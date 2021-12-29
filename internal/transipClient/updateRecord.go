package transipClient

import (
	"github.com/transip/gotransip/v6/domain"
	"go-transip-dyndns/internal/config"
)

func UpdateRecord(record config.Record) (*domain.DNSEntry, error) {
	entry, err := createDNSEntry(record)
	if err != nil {
		return nil, err
	}

	dnsRecord, err := FindRecord(record.Hostname, record.Type, record.Entry)
	if err != nil {
		return nil, NotFoundError
	}

	if dnsRecord.Content == entry.Content {
		return nil, NotChangedError
	}

	return &entry, getDomainRepo().UpdateDNSEntry(record.Hostname, entry)
}
