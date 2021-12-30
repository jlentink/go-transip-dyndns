package config

import "strings"

type ConfigObject struct {
	General General  `toml:"general"`
	Account Account  `toml:"account"`
	Record  []Record `toml:"record"`
}
type General struct {
	Verbose bool `toml:"verbose"`
	IPv4    bool `toml:"IPv4"`
	IPv6    bool `toml:"IPv6"`
}
type Account struct {
	Username   string `toml:"username"`
	PrivateKey string `toml:"private-key"`
}
type Record struct {
	Hostname string `toml:"hostname"`
	Entry    string `toml:"entry"`
	Content  string `toml:"content"`
	TTL      int    `toml:"ttl"`
	Type     string `toml:"type"`
}

func (r *Record) GetType() string {
	return strings.ToUpper(r.Type)
}
