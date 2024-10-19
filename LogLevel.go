package labelinglog

// LogLevel is used to specify output and configuration targets.
type LogLevel uint16

// Flg** is a log level bit flag.
const (
	FlgFatal  LogLevel = 1 << 0
	FlgError  LogLevel = 1 << 1
	FlgWarn   LogLevel = 1 << 2
	FlgNotice LogLevel = 1 << 3
	FlgInfo   LogLevel = 1 << 4
	FlgDebug  LogLevel = 1 << 5
)

// Flgset** is a preset log level bit flag set.
const (
	FlgsetAll    LogLevel = 0xffff
	FlgsetCommon LogLevel = FlgFatal | FlgError | FlgWarn | FlgNotice
)

var logLevelLabel = map[LogLevel]string{
	FlgFatal:  "FATAL",
	FlgError:  "ERROR",
	FlgWarn:   "WARN",
	FlgNotice: "NOTICE",
	FlgInfo:   "INFO",
	FlgDebug:  "DEBUG",
}
