package labelinglog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type tLogger struct {
	sync.Mutex
	isEnable bool
	writer   io.Writer
	prefix   string
	flg      LogLevel
}

func (thisLogger *tLogger) logSub(timestamp string, fileName string, msg string) {
	thisLogger.Lock()
	defer thisLogger.Unlock()

	_, err := fmt.Fprintln(thisLogger.writer, timestamp+thisLogger.prefix+" "+fileName+": "+msg)
	if err != nil {
		selfLogger.log(timestamp, fileName, err.Error())
	}
}

type tSelfLogger struct {
	sync.Mutex
	writer io.Writer
	prefix string
	flg    LogLevel
}

func (thisLogger *tSelfLogger) log(timestamp string, fileName string, msg string) {
	thisLogger.Lock()
	defer thisLogger.Unlock()
	fmt.Fprintln(thisLogger.writer, timestamp+thisLogger.prefix+" "+fileName+": "+msg)
}

var selfLogger = &tSelfLogger{
	writer: os.Stderr,
	prefix: "[logger][FATAL]",
	flg:    FlgsetAll,
}
