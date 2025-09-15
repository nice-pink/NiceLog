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
	if strings.EqualFold(level, "critical") {
		return LLCritical
	}
	if strings.EqualFold(level, "error") {
		return LLError
	}
	if strings.EqualFold(level, "warn") || strings.EqualFold(level, "warning") {
		return LLWarn
	}
	if strings.EqualFold(level, "info") {
		return LLInfo
	}
	if strings.EqualFold(level, "debug") {
		return LLDebug
	}
	return LLVerbose
}

func GetLogLevelPrefix(level LogLevel) string {
	if level == LLCritical {
		return "CRITICAL"
	}
	if level == LLError {
		return "ERROR"
	}
	if level == LLWarn {
		return "WARN"
	}
	if level == LLInfo {
		return "INFO"
	}
	if level == LLDebug {
		return "DEBUG"
	}
	return "VERBOSE"
}

func GetLogLevelColor(level LogLevel) string {
	if level == LLCritical {
		return Red
	}
	if level == LLError {
		return Red
	}
	if level == LLWarn {
		return Yellow
	}
	if level == LLInfo {
		return White
	}
	if level == LLDebug {
		return White
	}
	return Reset
}

// connection protocol

type ConnProtocol string

const (
	Tcp ConnProtocol = "tcp"
	Udp ConnProtocol = "udp"
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
