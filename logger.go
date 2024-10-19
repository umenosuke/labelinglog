package labelinglog

import (
	"fmt"
	"io"
	"sync"
)

// Wrapper for log output.
type tLogger struct {
	sync.RWMutex
	writer io.Writer
	prefix string
}

func (thisLogger *tLogger) log(timestamp string, fileName string, msg string) {
	thisLogger.RLock()
	defer thisLogger.RUnlock()

	_, err := fmt.Fprintln(thisLogger.writer, timestamp+thisLogger.prefix+" "+fileName+": "+msg)
	if err != nil {
		internalLog(timestamp, fileName, err.Error())
	}
}

func (thisLogger *tLogger) setIoWriter(writer io.Writer) {
	thisLogger.Lock()
	defer thisLogger.Unlock()
	thisLogger.writer = writer
}
