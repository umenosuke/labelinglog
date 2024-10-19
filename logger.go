package labelinglog

import (
	"fmt"
	"io"
	"sync"
)

// Wrapper for log output.
type tLogger struct {
	sync.Mutex
	isEnable bool
	writer   io.Writer
	prefix   string
	flg      LogLevel
}

func (thisLogger *tLogger) log(timestamp string, fileName string, msg string) {
	thisLogger.Lock()
	defer thisLogger.Unlock()

	_, err := fmt.Fprintln(thisLogger.writer, timestamp+thisLogger.prefix+" "+fileName+": "+msg)
	if err != nil {
		internalLogger.log(timestamp, fileName, err.Error())
	}
}
