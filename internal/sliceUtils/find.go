package sliceUtils

import (
	"fmt"
	"github.com/transip/gotransip/v6/domain"
)

// FindRecord find record based on characteristics
func FindRecord(hayStack []domain.DNSEntry, hostname string, recordEntry string, recordType string) (domain.DNSEntry, error) {
	for _, entry := range hayStack {
		if entry.Name == recordEntry && entry.Type == recordType {
			return entry, nil
		}
	}
	return domain.DNSEntry{}, fmt.Errorf("could not find record")
}
