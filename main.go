package labelinglog

import (
	"io"
)

// LogLevel a
type LogLevel uint16

// FlgAll a
const (
	FlgFatal  = 1 << 0
	FlgError  = 1 << 1
	FlgWarn   = 1 << 2
	FlgNotice = 1 << 3
	FlgInfo   = 1 << 4
	FlgDebug  = 1 << 5
)

// FlgAll a
const (
	FlgsetAll    = 0xffff
	FlgsetCommon = FlgFatal | FlgError | FlgWarn | FlgNotice
)

// LabelingLogger a
type LabelingLogger struct {
	loggers         []*tLogger
	enableFileame   bool
	enableTimestamp bool
}

// New a
func New(prefix string, writer io.Writer) *LabelingLogger {
	loggers := make([]*tLogger, 0)
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][FATAL] ",
		flg:      FlgFatal,
	})
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][ERROR] ",
		flg:      FlgError,
	})
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][WARN]  ",
		flg:      FlgWarn,
	})
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][NOTICE]",
		flg:      FlgNotice,
	})
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][INFO]  ",
		flg:      FlgInfo,
	})
	loggers = append(loggers, &tLogger{
		isEnable: true,
		writer:   writer,
		prefix:   "[" + prefix + "][DEBUG] ",
		flg:      FlgDebug,
	})

	return &LabelingLogger{
		loggers:         loggers,
		enableFileame:   true,
		enableTimestamp: true,
	}
}
