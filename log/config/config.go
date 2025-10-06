package config

import "time"

type ConnectionConfig struct {
	Address     string
	Protocol    ConnProtocol
	Timeout     time.Duration
	ContentType string
	QueryParams string
	IsHttpPost  bool
}

func GetConnectionConfig(address, protocol, contentType, queryParams string, timeout time.Duration, isHttpPost bool) ConnectionConfig {
	return ConnectionConfig{
		Address:     address,
		Protocol:    ConnProtocol(protocol),
		Timeout:     timeout,
		ContentType: contentType,
		QueryParams: queryParams,
		IsHttpPost:  isHttpPost,
	}
}

type Config struct {
	// std
	Prefix   string
	LogLevel LogLevel
	// http
	IsHttpPost  bool
	ContentType string
	QueryParams string
	// timestamp
	LogTimestamp bool
	TimeFormat   string
	IsUtc        bool
	// remote
	Connection ConnectionConfig
	// structured data
	Keys       Keys
	CommonData map[string]any
}

var config *Config

func init() {
	config = defaultConfig()
}

func DefaultConfig() *Config {
	return config
}

// default config

func defaultConfig() *Config {
	return &Config{
		Prefix:       "",
		LogLevel:     LLDebug,
		LogTimestamp: false,
		TimeFormat:   time.DateTime,
		IsUtc:        false,
		Connection: ConnectionConfig{
			Address:  "",
			Protocol: Tcp,
			Timeout:  3 * time.Second,
		},
		Keys: Keys{
			Timestamp: "timestamp",
			Message:   "message",
			Severity:  "severity",
		},
		CommonData: map[string]any{},
	}
}
