package log

import (
	"net/http"
	"sync"
	"time"

	"github.com/nice-pink/NiceLog/log/config"
)

// And just go global.
var defaultLogger *logger

type logger struct {
	mu         sync.Mutex
	cfg        *config.Config
	httpClient *http.Client
}

func newLogger() *logger {
	return &logger{
		mu:  sync.Mutex{},
		cfg: config.DefaultConfig(),
	}
}

func init() {
	defaultLogger = newLogger()
}

/*** configure ***/

// prefix
func SetPrefix(prefix string) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.Prefix = prefix
}

// content type
func SetContentType(contentType string) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.ContentType = contentType
}

// query params
func SetQueryParams(queryParams string) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.QueryParams = queryParams
}

// is http
func SetIsHttpPost(isHttpPost bool) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.IsHttpPost = isHttpPost
	defaultLogger.httpClient = &http.Client{
		Timeout: 1 * time.Second,
	}
}

// log level

func SetLogLevelDebug() {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogLevel = config.LLDebug
}

func SetLogLevelInfo() {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogLevel = config.LLInfo
}

func SetLogLevelWarn() {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogLevel = config.LLWarn
}

func SetLogLevelError() {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogLevel = config.LLError
}

func SetLogLevelCritical() {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogLevel = config.LLCritical
}

// timestamp

func SetLogTimestamp(logTimestamp bool) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.LogTimestamp = logTimestamp
}

func SetTimeFormat(timeFormat string) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.TimeFormat = timeFormat
}

func SetIsUtc(isUtc bool) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.IsUtc = isUtc
}

// keys

func SetKeys(keys config.Keys) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.Keys = keys
}

// common data

func SetCommonData(commonData map[string]any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.cfg.CommonData = commonData
}

// connect to remote log sink

type Connection struct {
	Address     string
	Protocol    string
	ContentType string
	QueryParams string
	Timeout     time.Duration
	IsHttp      bool
}

func Connect(cfg Connection) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	connectionConfig := config.GetConnectionConfig(cfg.Address, cfg.Protocol, cfg.ContentType, cfg.QueryParams, cfg.Timeout, cfg.IsHttp)
	if cfg.Protocol == "http" {
		connectionConfig.IsHttpPost = true
		defaultLogger.httpClient = &http.Client{
			Timeout: 1 * time.Second,
		}
	} else {
		defaultLogger.connect(connectionConfig)
	}
}

/*** log ***/

func Verbose(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.verbose(logs...)
}

func VerboseD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.verboseD(data, logs...)
}

func Debug(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.debug(logs...)
}

func DebugD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.debugD(data, logs...)
}

func Info(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.info(logs...)
}

func InfoD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.infoD(data, logs...)
}

func Warn(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.warn(logs...)
}

func WarnD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.warnD(data, logs...)
}

func Error(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.error(logs...)
}

func ErrorD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.errorD(data, logs...)
}

func Critical(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.critical(logs...)
}

func CriticalD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.criticalD(data, logs...)
}

// special

func LogString(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.logString(logs...)
}

func Success(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.success(logs...)
}

func SuccessD(data map[string]interface{}, logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.successD(data, logs...)
}

// std

func Print(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.print(logs...)
}

func Println(logs ...any) {
	defaultLogger.mu.Lock()
	defer defaultLogger.mu.Unlock()
	defaultLogger.println(logs...)
}
