package log

import (
	"encoding/json"
	"fmt"
	"maps"
	"net"
	"os"
	"strings"
	"time"

	"github.com/nice-pink/NiceLog/log/config"
)

func (l *logger) verbose(logs ...any) {
	if l.cfg.LogLevel > config.LLVerbose {
		return
	}
	msg := getMsg(logs...)

	prefix := config.GetLogLevelPrefix(config.LLVerbose)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLVerbose),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, "VERBOSE")
}

func (l *logger) verboseD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLVerbose {
		return
	}
	msg := getMsg(logs...)
	level := config.GetLogLevelPrefix(config.LLVerbose)
	l.printLog(
		true,
		level,
		config.GetLogLevelColor(config.LLVerbose),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, "VERBOSE")
}

func (l *logger) debug(logs ...any) {
	if l.cfg.LogLevel > config.LLDebug {
		return
	}

	msg := getMsg(logs...)
	prefix := config.GetLogLevelPrefix(config.LLDebug)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLDebug),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, "DEBUG")
}

func (l *logger) debugD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLDebug {
		return
	}
	msg := getMsg(logs...)
	prefix := config.GetLogLevelPrefix(config.LLDebug)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLDebug),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, "DEBUG")
}

func (l *logger) info(logs ...any) {
	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	msg := getMsg(logs...)
	prefix := config.GetLogLevelPrefix(config.LLInfo)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLInfo),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, prefix)
}

func (l *logger) infoD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	msg := getMsg(logs...)
	prefix := config.GetLogLevelPrefix(config.LLInfo)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLInfo),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, "INFO")
}

func (l *logger) warn(logs ...any) {
	if l.cfg.LogLevel > config.LLWarn {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLWarn)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLWarn),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, "WARN")
}

func (l *logger) warnD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLWarn {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLWarn)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLWarn),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, prefix)
}

func (l *logger) error(logs ...any) {
	if l.cfg.LogLevel > config.LLError {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLError)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLError),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, prefix)
}

func (l *logger) errorD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLError {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLError)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLError),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, prefix)
}

func (l *logger) critical(logs ...any) {
	if l.cfg.LogLevel > config.LLCritical {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLCritical)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLCritical),
		logs...,
	)
	l.sendJsonWithSeverity(msg, nil, prefix)
}

func (l *logger) criticalD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLCritical {
		return
	}
	prefix := config.GetLogLevelPrefix(config.LLCritical)
	msg := getMsg(logs...)
	l.printLog(
		true,
		prefix,
		config.GetLogLevelColor(config.LLCritical),
		logs...,
	)
	l.sendJsonWithSeverity(msg, data, prefix)
}

func (l *logger) success(logs ...any) {

	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	level := config.GetLogLevelPrefix(config.LLInfo)
	msg := getMsg(logs...)
	l.printLog(true, level, config.Green, logs...)
	l.sendJsonWithSeverity(msg, nil, "SUCCESS")
}

func (l *logger) successD(data map[string]any, logs ...any) {
	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	level := config.GetLogLevelPrefix(config.LLInfo)
	msg := getMsg(logs...)
	l.printLog(true, level, config.Green, logs...)
	l.sendJsonWithSeverity(msg, data, "SUCCESS")
}

func (l *logger) logString(logs ...any) {
	msg := getMsg(logs...)
	l.sendString(msg)
	l.printLog(true, "STRING", config.Reset, logs...)
}

func (l *logger) print(logs ...any) {
	msg := getMsg(logs...)
	l.printLog(false, "", config.Reset, logs...)
	l.sendJsonWithSeverity(msg, nil, "")
}

func (l *logger) println(logs ...any) {
	msg := getMsg(logs...)
	l.printLog(true, "", config.Reset, logs...)
	l.sendJsonWithSeverity(msg, nil, "")
}

// private

func (l *logger) connect(cfg config.ConnectionConfig) net.Conn {
	network := config.GetNetwork(cfg.Protocol)
	conn, err := net.Dial(network, cfg.Address)
	if err != nil {
		// Err(err, "dial to network", network, "address", l.Address)
		return nil
	}

	// update deadlines
	deadline := time.Now().Add(cfg.Timeout * time.Second)
	conn.SetDeadline(deadline)
	conn.SetWriteDeadline(deadline)
	conn.SetReadDeadline(deadline)

	return conn
}

func (l *logger) sendJsonWithSeverity(msg string, add map[string]any, severity string) {
	if l.cfg.Connection.Address == "" {
		if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
			fmt.Println("No address for remote logging.")
		}
		return
	}

	// create map
	data := map[string]any{}
	data[l.cfg.Keys.Message] = msg

	// add severity
	if severity != "" {
		data[l.cfg.Keys.Severity] = severity
	}

	// add timestamp
	data[l.cfg.Keys.Timestamp] = l.getFormattedTimestamp()

	// copy additional
	if add != nil {
		maps.Copy(data, add)
	}
	if l.cfg.CommonData != nil {
		maps.Copy(data, l.cfg.CommonData)
	}

	go l.sendJson(data)
}

func (l *logger) sendJson(data map[string]any) bool {
	if l.cfg.Connection.Address == "" {
		if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
			fmt.Println("No address for remote logging.")
		}
		return false
	}

	conn := l.connect(l.cfg.Connection)
	if conn == nil {
		if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
			fmt.Println("No connection for remote logging.")
		}
		return false
	}
	defer conn.Close()

	payload, err := json.Marshal(data)
	if err != nil {
		print("ERROR", l.cfg.Prefix, "cannot marshal data", data, err.Error())
		return false
	}

	if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
		fmt.Println(string(payload))
	}

	_, err = conn.Write(payload)
	if err != nil {
		print("ERROR", l.cfg.Prefix, "cannot write payload", err.Error())
		return false
	}
	return true
}

func (l *logger) sendString(data string) bool {
	if l.cfg.Connection.Address == "" {
		return false
	}

	if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
		fmt.Println(data)
	}

	conn := l.connect(l.cfg.Connection)
	if conn == nil {
		return false
	}
	defer conn.Close()

	_, err := conn.Write([]byte(data))
	if err != nil {
		print("ERROR", l.cfg.Prefix, "cannot write string", err.Error())
		return false
	}
	return true
}

func getMsg(logs ...any) string {
	msg := fmt.Sprintln(logs...)
	return strings.TrimSuffix(msg, "\n")
}

func (l *logger) printLog(newline bool, level, color string, logs ...any) {
	// add prefix
	if l.cfg.Prefix != "" {
		logs = append([]any{l.cfg.Prefix}, logs...)
	}

	// add timestamp
	if l.cfg.LogTimestamp {
		logs = append([]any{l.getFormattedTimestamp()}, logs...)
	}

	// add level
	if level != "" {
		logs = append([]any{color + level + ":" + config.Reset}, logs...)
	}

	// add timestamp
	if newline {
		fmt.Println(logs...)
	} else {
		fmt.Print(logs...)
	}
}

func (l *logger) getFormattedTimestamp() string {
	// timestamp
	var ts time.Time
	if l.cfg.IsUtc {
		ts = time.Now().UTC()
	} else {
		ts = time.Now()
	}
	return ts.Format(l.cfg.TimeFormat)
}
