package labelinglog

import "io"

// SetIoWriter a
func (thisLabelingLogger *LabelingLogger) SetIoWriter(targetLevelFlgs LogLevel, writer io.Writer) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.Lock()
		logger.writer = writer
		logger.Unlock()
	}
}
