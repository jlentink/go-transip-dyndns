package gipify

const (
	// UNKNOWN IP type
	UNKNOWN = 0
	// IPV4 IPv4 type
	IPV4 = 4
	// IPV6 IPv6 type
	IPV6 = 6
)

// IP description struct
type IP struct {
	IP   string `json:"ip"`
	Type int
}
