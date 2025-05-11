package labelinglog

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
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

	basePath string
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

	basePath := "/"
	{
		_, file, _, ok := runtime.Caller(0)
		if ok {
			for {
				parent := filepath.Dir(file)
				if parent == file {
					break
				}
				if filepath.Base(file) == "vendor" {
					basePath = parent
					break
				}
				file = parent
			}
		}
	}

	return &LabelingLogger{
		loggers:            loggers,
		enableLoggerLevels: enableLogLevels,
		enableFileame:      true,
		enableTimestamp:    true,

		basePath: basePath,
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

	fileName := l.getFileName()

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

	fileName := l.getFileName()

	msgLines := make([]string, 0)
	reader := bufio.NewReader(strings.NewReader(msg))
	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			internalLog(timestamp, fileName, err.Error())
			break
		}

		if err == io.EOF {
			if len(line) != 0 {
				line, _ = strings.CutSuffix(line, "\n")
				line, _ = strings.CutSuffix(line, "\r")
				msgLines = append(msgLines, line)
			}

			break
		}

		line, _ = strings.CutSuffix(line, "\n")
		line, _ = strings.CutSuffix(line, "\r")
		msgLines = append(msgLines, line)
	}

	for _, flg := range tartgetLogLevels {
		l.loggers[flg].logMultiLines(timestamp, fileName, msgLines)
	}
}

func (l *LabelingLogger) getFileName() string {
	if l.enableFileame {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			relPath, err := filepath.Rel(l.basePath, file)
			if err != nil {
				return "(unknown) "
			} else {
				return fmt.Sprintf("%s:%d", relPath, line) + " "
			}
		} else {
			return "(unknown) "
		}
	} else {
		return ""
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
