package log

import (
	"encoding/json"
	"fmt"
	"maps"
	"net"
	"os"
	"strings"
	"time"

	"github.com/nice-pink/nicelog/log/config"
)

func (l *logger) verbose(logs ...any) {
	if l.cfg.LogLevel > config.LLVerbose {
		return
	}
	msg := getMsg(logs...)

	l.printLog(true, "VERBOSE", logs...)
	l.sendJsonWithSeverity(msg, nil, "VERBOSE")
}

func (l *logger) verboseD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLVerbose {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "VERBOSE", logs...)
	l.sendJsonWithSeverity(msg, data, "VERBOSE")
}

func (l *logger) debug(logs ...any) {
	if l.cfg.LogLevel > config.LLDebug {
		return
	}

	msg := getMsg(logs...)
	l.printLog(true, "DEBUG", logs...)
	l.sendJsonWithSeverity(msg, nil, "DEBUG")
}

func (l *logger) debugD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLDebug {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "DEBUG", logs...)
	l.sendJsonWithSeverity(msg, data, "DEBUG")
}

func (l *logger) info(logs ...any) {
	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "INFO", logs...)
	l.sendJsonWithSeverity(msg, nil, "INFO")
}

func (l *logger) infoD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLInfo {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "INFO", logs...)
	l.sendJsonWithSeverity(msg, data, "INFO")
}

func (l *logger) warn(logs ...any) {
	if l.cfg.LogLevel > config.LLWarn {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "WARN", logs...)
	l.sendJsonWithSeverity(msg, nil, "WARN")
}

func (l *logger) warnD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLWarn {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "WARN", logs...)
	l.sendJsonWithSeverity(msg, data, "WARN")
}

func (l *logger) error(logs ...any) {
	if l.cfg.LogLevel > config.LLError {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "ERROR", logs...)
	l.sendJsonWithSeverity(msg, nil, "ERROR")
}

func (l *logger) errorD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLError {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "ERROR", logs...)
	l.sendJsonWithSeverity(msg, data, "ERROR")
}

func (l *logger) critical(logs ...any) {
	if l.cfg.LogLevel > config.LLCritical {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "CRITICAL", logs...)
	l.sendJsonWithSeverity(msg, nil, "CRITICAL")
}

func (l *logger) criticalD(data map[string]interface{}, logs ...any) {
	if l.cfg.LogLevel > config.LLCritical {
		return
	}
	msg := getMsg(logs...)
	l.printLog(true, "CRITICAL", logs...)
	l.sendJsonWithSeverity(msg, data, "CRITICAL")
}

func (l *logger) logString(logs ...any) {
	msg := getMsg(logs...)
	l.sendString(msg)
	l.printLog(true, "STRING", logs...)
}

func (l *logger) print(logs ...any) {
	msg := getMsg(logs...)
	l.printLog(false, "", logs...)
	l.sendJsonWithSeverity(msg, nil, "")
}

func (l *logger) println(logs ...any) {
	msg := getMsg(logs...)
	l.printLog(true, "", logs...)
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

func (l *logger) sendJsonWithSeverity(msg string, add map[string]interface{}, severity string) {
	if l.cfg.Connection.Address == "" {
		if os.Getenv("GU_REMOTE_LOG_DEBUG") == "true" {
			fmt.Println("No address for remote logging.")
		}
		return
	}

	// create map
	data := map[string]interface{}{}
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

func (l *logger) sendJson(data map[string]interface{}) bool {
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

func (l *logger) printLog(newline bool, level string, logs ...any) {
	// add prefix
	if l.cfg.Prefix != "" {
		logs = append([]any{l.cfg.Prefix}, logs...)
	}

	// add level
	if level != "" {
		logs = append([]any{level + ":"}, logs...)
	}

	if l.cfg.LogTimestamp {
		logs = append([]any{l.getFormattedTimestamp()}, logs...)
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
