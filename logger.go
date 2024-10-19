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

func (l *tLogger) log(timestamp string, fileName string, msg string) {
	l.RLock()
	defer l.RUnlock()

	_, err := fmt.Fprintln(l.writer, timestamp+l.prefix+" "+fileName+": "+msg)
	if err != nil {
		internalLog(timestamp, fileName, err.Error())
	}
}

func (l *tLogger) logMultiLines(timestamp string, fileName string, msgLines []string) {
	l.Lock()
	defer l.Unlock()

	for _, msg := range msgLines {
		_, err := fmt.Fprintln(l.writer, timestamp+l.prefix+" "+fileName+": "+msg)
		if err != nil {
			internalLog(timestamp, fileName, err.Error())
		}
	}
}

func (l *tLogger) setIoWriter(writer io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.writer = writer
}
