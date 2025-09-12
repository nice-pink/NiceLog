package config

import "strings"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// log level

type LogLevel int

const (
	LLVerbose LogLevel = iota
	LLDebug
	LLInfo
	LLWarn
	LLError
	LLCritical
)

func GetLogLevel(level string) LogLevel {
	if strings.ToLower(level) == "critical" {
		return LLCritical
	}
	if strings.ToLower(level) == "error" {
		return LLError
	}
	if strings.ToLower(level) == "warn" || strings.ToLower(level) == "warning" {
		return LLWarn
	}
	if strings.ToLower(level) == "info" {
		return LLInfo
	}
	if strings.ToLower(level) == "debug" {
		return LLDebug
	}
	return LLVerbose
}

// connection protocol

type ConnProtocol int

const (
	Tcp ConnProtocol = iota
	Udp
)

func GetNetwork(protocol ConnProtocol) string {
	if protocol == Udp {
		return "udp"
	}
	return "tcp"
}

// common keys

type Keys struct {
	Timestamp string
	Message   string
	Severity  string
}
