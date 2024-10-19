package labelinglog

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

// LabelingLogger is the main
type LabelingLogger struct {
	loggers            map[LogLevel]*tLogger
	enableLoggerLevels []LogLevel
	enableFileame      bool
	enableTimestamp    bool
}

// New returns an initialized LabelingLogger
func New(prefix string, writer io.Writer) *LabelingLogger {
	loggers := make(map[LogLevel]*tLogger, len(logLevelLabel))
	enableLogLevels := make([]LogLevel, 0, len(logLevelLabel))

	for flg, label := range logLevelLabel {
		loggers[flg] = &tLogger{
			writer: writer,
			prefix: "[" + prefix + "][" + label + "] ",
		}

		enableLogLevels = append(enableLogLevels, flg)
	}

	return &LabelingLogger{
		loggers:            loggers,
		enableLoggerLevels: enableLogLevels,
		enableFileame:      true,
		enableTimestamp:    true,
	}
}

// Log outputs messages at the specified log level.
func (l *LabelingLogger) Log(targetLevelFlgs LogLevel, msg string) {
	tartgetLogLevels := make([]LogLevel, 0, len(l.enableLoggerLevels))
	for _, flg := range l.enableLoggerLevels {
		if targetLevelFlgs&flg != 0 {
			tartgetLogLevels = append(tartgetLogLevels, flg)
		}
	}
	if len(tartgetLogLevels) == 0 {
		return
	}

	var timestamp string
	if l.enableTimestamp {
		timestamp = time.Now().Format("2006/01/02 15:04:05.000") + " "
	} else {
		timestamp = ""
	}

	var fileName string
	if l.enableFileame {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			s := strings.Split(file, "/")
			fileName = fmt.Sprintf("%s line %3d", s[len(s)-1], line) + " "
		} else {
			fileName = "unknown "
		}
	} else {
		fileName = ""
	}

	for _, flg := range tartgetLogLevels {
		l.loggers[flg].log(timestamp, fileName, msg)
	}
}

// LogMultiLines outputs multi-line messages at the specified log level.
func (l *LabelingLogger) LogMultiLines(targetLevelFlgs LogLevel, msg string) {
	tartgetLogLevels := make([]LogLevel, 0, len(l.enableLoggerLevels))
	for _, flg := range l.enableLoggerLevels {
		if targetLevelFlgs&flg != 0 {
			tartgetLogLevels = append(tartgetLogLevels, flg)
		}
	}
	if len(tartgetLogLevels) == 0 {
		return
	}

	var timestamp string
	if l.enableTimestamp {
		timestamp = time.Now().Format("2006/01/02 15:04:05.000") + " "
	} else {
		timestamp = ""
	}

	var fileName string
	if l.enableFileame {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			s := strings.Split(file, "/")
			fileName = fmt.Sprintf("%s line %3d", s[len(s)-1], line) + " "
		} else {
			fileName = "unknown "
		}
	} else {
		fileName = ""
	}

	scanner := bufio.NewScanner(strings.NewReader(msg))
	for scanner.Scan() {
		for _, flg := range tartgetLogLevels {
			l.loggers[flg].log(timestamp, fileName, msg)
		}
	}

	if err := scanner.Err(); err != nil {
		internalLog(timestamp, fileName, err.Error())
	}
}

// SetEnableLevel enables only the output of the specified log level.
func (l *LabelingLogger) SetEnableLevel(targetLevelFlgs LogLevel) {
	l.enableLoggerLevels = nil

	for flg := range l.loggers {
		if targetLevelFlgs&flg != 0 {
			l.enableLoggerLevels = append(l.enableLoggerLevels, flg)
		}
	}
}

// SetIoWriter changes the output destination of the specified log level.
func (l *LabelingLogger) SetIoWriter(targetLevelFlgs LogLevel, writer io.Writer) {
	for flg, logger := range l.loggers {
		if targetLevelFlgs&flg != 0 {
			logger.setIoWriter(writer)
		}
	}
}

// DisableFilename disables the output of the file name of the log caller.
func (l *LabelingLogger) DisableFilename() {
	l.enableFileame = false
}

// EnableFilename enables output of the file name of the caller of the log.
func (l *LabelingLogger) EnableFilename() {
	l.enableFileame = true
}

// DisableTimestamp disables log timestamp output.
func (l *LabelingLogger) DisableTimestamp() {
	l.enableTimestamp = false
}

// EnableTimestamp enables log timestamp output.
func (l *LabelingLogger) EnableTimestamp() {
	l.enableTimestamp = true
}
