package transipClient

import (
	"fmt"
	"github.com/transip/gotransip/v6/domain"
)

//FindRecord Find record for domain.
func FindRecord(domainName, recordType, entry string) (*domain.DNSEntry, error) {

	records, err := getDomainRepo().GetDNSEntries(domainName)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if record.Name == entry && record.Type == recordType {
			return &record, nil
		}
	}
	return nil, fmt.Errorf("did not find record")
}
