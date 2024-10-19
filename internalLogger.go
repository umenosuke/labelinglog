package labelinglog

import (
	"fmt"
	"io"
	"os"
)

// Use for logger's own log output.
type tInternalLogger struct {
	writer io.Writer
	prefix string
	flg    LogLevel
}

func (thisLogger *tInternalLogger) log(timestamp string, fileName string, msg string) {
	fmt.Fprintln(thisLogger.writer, timestamp+thisLogger.prefix+" "+fileName+": "+msg)
}

var internalLogger = &tInternalLogger{
	writer: os.Stderr,
	prefix: "[logger][FATAL]",
	flg:    FlgsetAll,
}
