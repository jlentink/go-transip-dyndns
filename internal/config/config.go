package config

import (
	"bufio"
	"github.com/kyokomi/emoji"
	"github.com/spf13/viper"
	"go-transip-dyndns/internal/logger"
	"go-transip-dyndns/internal/sliceUtils"
	"io"
	"os"
	"strings"
	"time"
)

var _config *viper.Viper
var _configObject ConfigObject

// Init the config setup
func _init() {
	_config = viper.New()
	_config.SetConfigType("toml")
	_config.AddConfigPath("/etc")
	_config.AddConfigPath("/etc/go-transip-dyndns.d")
	_config.AddConfigPath(".")
	_config.SetConfigName("go-transip-dyndns")

	if err := _config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Get().Fatalf("%s\n", err.Error())
		} else {
			logger.Get().Fatalf("Error parsing config file. (%s)\n", err.Error())
		}
	}
	err := _config.Unmarshal(&_configObject)
	if err != nil {
		logger.Get().Fatalf("Error parsing config file. (%s)\n", err.Error())
	}
}

// Get the config instance
func Get() *viper.Viper {
	if _config == nil {
		_init()
	}
	return _config
}

func GetUsername() string {
	return Get().GetString("account.username")
}

func GetPrivateKey() io.Reader {
	privateKeyEntry := Get().GetString("account.private-key")
	if len(privateKeyEntry) > 27 && privateKeyEntry[0:27] == "-----BEGIN PRIVATE KEY-----" {
		return strings.NewReader(privateKeyEntry)
	}
	file, err := os.Open(privateKeyEntry)
	if err != nil {
		_, _ = emoji.Printf(":exclamation: Could not read config file (%s)\n", err.Error()) //nolint:errcheck
	}
	return bufio.NewReader(file)
}

func IsIPv4Enabled() bool {
	return Get().GetBool("general.IPv4")
}

func IsIPv6Enabled() bool {
	return Get().GetBool("general.IPv6")
}

func GetRecords() []Record {
	return _configObject.Record
}

func GetDomains() []string {
	var domains []string
	for _, record := range GetRecords() {
		if !sliceUtils.Contains(domains, record.Hostname) {
			domains = append(domains, record.Hostname)
		}
	}
	return domains
}

func GetUpdateFrequency() time.Duration {
	updateFrequency := Get().GetInt("general.update-frequency")
	if updateFrequency < 1 {
		updateFrequency = 1
	}
	return time.Duration(updateFrequency) * time.Minute
}
