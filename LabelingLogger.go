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
	loggers         []*tLogger
	enableFileame   bool
	enableTimestamp bool
}

// New returns an initialized LabelingLogger
func New(prefix string, writer io.Writer) *LabelingLogger {
	loggers := make([]*tLogger, 0)

	for flg, label := range logLevelLabel {
		loggers = append(loggers, &tLogger{
			isEnable: true,
			writer:   writer,
			prefix:   "[" + prefix + "][" + label + "] ",
			flg:      flg,
		})
	}

	return &LabelingLogger{
		loggers:         loggers,
		enableFileame:   true,
		enableTimestamp: true,
	}
}

// Log outputs messages at the specified log level.
func (thisLabelingLogger *LabelingLogger) Log(targetLevelFlgs LogLevel, msg string) {
	if !thisLabelingLogger.isActive(targetLevelFlgs) {
		return
	}

	var timestamp string
	if thisLabelingLogger.enableTimestamp {
		timestamp = time.Now().Format("2006/01/02 15:04:05.000") + " "
	} else {
		timestamp = ""
	}

	var fileName string
	if thisLabelingLogger.enableFileame {
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

	for _, logger := range thisLabelingLogger.loggers {
		if logger.isEnable {
			if targetLevelFlgs&logger.flg != 0 {
				logger.log(timestamp, fileName, msg)
			}
		}
	}
}

// LogMultiLines outputs multi-line messages at the specified log level.
func (thisLabelingLogger *LabelingLogger) LogMultiLines(targetLevelFlgs LogLevel, msg string) {
	if !thisLabelingLogger.isActive(targetLevelFlgs) {
		return
	}

	var timestamp string
	if thisLabelingLogger.enableTimestamp {
		timestamp = time.Now().Format("2006/01/02 15:04:05.000") + " "
	} else {
		timestamp = ""
	}

	var fileName string
	if thisLabelingLogger.enableFileame {
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
		for _, logger := range thisLabelingLogger.loggers {
			if logger.isEnable {
				if targetLevelFlgs&logger.flg != 0 {
					logger.log(timestamp, fileName, scanner.Text())
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		internalLogger.log(timestamp, fileName, err.Error())
	}
}

func (thisLabelingLogger *LabelingLogger) isActive(targetLevelFlgs LogLevel) bool {
	for _, logger := range thisLabelingLogger.loggers {
		if logger.isEnable {
			if targetLevelFlgs&logger.flg != 0 {
				return true
			}
		}
	}
	return false
}

// SetEnableLevel enables only the output of the specified log level.
func (thisLabelingLogger *LabelingLogger) SetEnableLevel(targetLevelFlgs LogLevel) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.isEnable = targetLevelFlgs&logger.flg != 0
	}
}

// SetIoWriter changes the output destination of the specified log level.
func (thisLabelingLogger *LabelingLogger) SetIoWriter(targetLevelFlgs LogLevel, writer io.Writer) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.Lock()
		logger.writer = writer
		logger.Unlock()
	}
}

// DisableFilename disables the output of the file name of the log caller.
func (thisLabelingLogger *LabelingLogger) DisableFilename() {
	thisLabelingLogger.enableFileame = false
}

// EnableFilename enables output of the file name of the caller of the log.
func (thisLabelingLogger *LabelingLogger) EnableFilename() {
	thisLabelingLogger.enableFileame = true
}

// DisableTimestamp disables log timestamp output.
func (thisLabelingLogger *LabelingLogger) DisableTimestamp() {
	thisLabelingLogger.enableTimestamp = false
}

// EnableTimestamp enables log timestamp output.
func (thisLabelingLogger *LabelingLogger) EnableTimestamp() {
	thisLabelingLogger.enableTimestamp = true
}
