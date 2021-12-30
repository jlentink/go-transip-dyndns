package transipClient

import (
	"github.com/transip/gotransip/v6/domain"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/logger"
	"net"
)

func constructHostname(entry domain.DNSEntry, record config.Record) string {
	hostname := ""
	if entry.Name != "@" {
		hostname += entry.Name + "."
	}
	hostname += record.Hostname
	return hostname
}

func dnsAChanged(entry domain.DNSEntry, record config.Record) bool {
	ips, err := net.LookupIP(constructHostname(entry, record))
	if err != nil {
		logger.Get().Debug("Error fetching DNS record. (%s)\n", err.Error())
		return true
	}
	for _, ip := range ips {
		if record.GetType() == "A" && ip.String() == entry.Content {
			logger.Get().Debug("DNS [A] record has not been changed.\n")
			return false
		}
		if record.GetType() == "AAAA" && ip.String() == entry.Content {
			logger.Get().Debug("DNS [AAAA] record has not been changed.\n")
			return false
		}

	}
	return true
}

func dnsTXTChanged(entry domain.DNSEntry, record config.Record) bool {
	txts, err := net.LookupTXT(constructHostname(entry, record))
	if err != nil {
		logger.Get().Debug("Error fetching DNS record. (%s)\n", err.Error())
		return true
	}
	for _, txt := range txts {
		if txt == entry.Content {
			logger.Get().Debug("DNS [TXT] record has not been changed.\n")
			return false
		}

	}
	return true
}

func dnsChanged(entry domain.DNSEntry, record config.Record) bool {
	switch entry.Type {
	case "A":
		fallthrough
	case "AAAA":
		return dnsAChanged(entry, record)
	case "TXT":
		return dnsTXTChanged(entry, record)
	default:
		return true
	}
}

func UpdateRecord(record config.Record) (*domain.DNSEntry, error) {
	entry, err := createDNSEntry(record)
	if err != nil {
		return nil, err
	}

	if !dnsChanged(entry, record) {
		return nil, NotChangedError
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
