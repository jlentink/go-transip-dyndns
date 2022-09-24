package gipify

import (
	"encoding/json"
	"fmt"
	"go-transip-dyndns/internal/logger"
	"io"
	"net/http"
	"regexp"
)

const (
	regexIPv6 = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	regexIPv4 = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
)

var (
	IPv4URL = "https://api.ipify.org?format=json"
	IPv6URL = "https://api64.ipify.org?format=json"
)

// GetIPv4 from the ipify service IPv4
func GetIPv4() (*IP, error) {
	resp, err := http.Get(IPv4URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status should be 200 is %d", resp.StatusCode)
	}
	return parse(resp.Body)
}

// GetIPv6 from the ipify service IPv6
func GetIPv6() (*IP, error) {
	resp, err := http.Get(IPv6URL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status should be 200 is %d", resp.StatusCode)
	}
	return parse(resp.Body)
}

func parse(i io.Reader) (*IP, error) {
	respIp := IP{Type: UNKNOWN}
	resp, err := io.ReadAll(i)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &respIp)
	if err != nil {
		return nil, err
	}
	respIp.Type = ipType(respIp.IP)

	logger.Get().Debugf("Found IP address. (%s)\n", respIp.IP)
	return &respIp, nil
}

func ipType(i string) int {
	ipv6, _ := regexp.MatchString(regexIPv6, i)
	if ipv6 {
		return IPV6
	}

	ipv4, _ := regexp.MatchString(regexIPv4, i)
	if ipv4 {
		return IPV4
	}

	return UNKNOWN
}
