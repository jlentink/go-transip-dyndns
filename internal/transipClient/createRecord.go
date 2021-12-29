package transipClient

import (
	"bytes"
	"fmt"
	"github.com/transip/gotransip/v6/domain"
	"go-transip-dyndns/internal/config"
	"go-transip-dyndns/internal/gipify"
	"html/template"
	"strings"
)

type recordContent struct {
	IPv4 string
	IPv6 string
}

func validRecordType(recordType string) bool {
	recordType = strings.ToUpper(recordType)
	if recordType == "A" || recordType == "AAAA" || recordType == "CNAME" ||
		recordType == "MX" || recordType == "TXT" || recordType == "SRV" || recordType == "SSHFP" ||
		recordType == "TLSA" {
		return true
	}
	return false
}

// CreateRecord creates a record in the domain
func CreateRecord(record config.Record) error {
	entry, err := createDNSEntry(record)
	if err != nil {
		return err
	}
	getDomainRepo().AddDNSEntry(record.Hostname, entry)

	return nil
}

func createDNSEntry(record config.Record) (domain.DNSEntry, error) {
	if !validRecordType(record.Type) {
		return domain.DNSEntry{}, fmt.Errorf("record type %s is not valid. "+
			"Allowed are 'A', 'AAAA', 'CNAME', 'MX', 'NS', 'TXT', 'SRV', 'SSHFP' and 'TLSA'", record.Type)
	}

	var entry = domain.DNSEntry{}
	var err error

	switch strings.ToUpper(record.Type) {
	case "A":
		entry, err = createAEntry(record)
	case "AAAA":
		entry, err = createAAAAEntry(record)
	case "CNAME":
		fallthrough
	case "MX":
		fallthrough
	case "TXT":
		fallthrough
	case "SRV":
		fallthrough
	case "SSHFP":
		fallthrough
	case "TLSA":
		entry, err = createGenericEntry(record)
	}
	return entry, err
}

func createAEntry(record config.Record) (domain.DNSEntry, error) {
	ip, err := GetIPv4()
	if err != nil {
		return domain.DNSEntry{}, err
	}
	entry := domain.DNSEntry{
		Name:    record.Entry,
		Expire:  record.TTL,
		Type:    record.Type,
		Content: ip.IP,
	}
	return entry, nil
}

func createAAAAEntry(record config.Record) (domain.DNSEntry, error) {
	ip, err := GetIPv6()
	if err != nil {
		return domain.DNSEntry{}, err
	}
	if ip.Type == gipify.IPV4 {
		return createAEntry(record)
	}

	entry := domain.DNSEntry{
		Name:    record.Entry,
		Expire:  record.TTL,
		Type:    record.Type,
		Content: ip.IP,
	}
	return entry, nil
}

func createContentObject() (recordContent, error) {
	ipv6, err := GetIPv6()
	if err != nil {
		return recordContent{}, err
	}

	ipv4, err := GetIPv4()
	if err != nil {
		return recordContent{}, err
	}
	content := recordContent{
		IPv4: ipv4.IP,
		IPv6: ipv6.IP,
	}
	return content, nil
}

func createGenericEntry(record config.Record) (domain.DNSEntry, error) {
	var buf = &bytes.Buffer{}
	contentObj, err := createContentObject()
	if err != nil {
		return domain.DNSEntry{}, err
	}

	tmpl, err := template.New("content").Parse(record.Content)
	if err != nil {
		return domain.DNSEntry{}, err
	}

	err = tmpl.Execute(buf, contentObj)
	if err != nil {
		return domain.DNSEntry{}, err
	}

	entry := domain.DNSEntry{
		Name:    record.Entry,
		Expire:  record.TTL,
		Type:    record.Type,
		Content: buf.String(),
	}
	return entry, nil
}
