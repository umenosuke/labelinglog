package labelinglog

import "io"

// SetEnableLevel a
func (thisLabelingLogger *LabelingLogger) SetEnableLevel(targetLevelFlgs LogLevel) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.isEnable = targetLevelFlgs&logger.flg != 0
	}
}

// SetIoWriter a
func (thisLabelingLogger *LabelingLogger) SetIoWriter(targetLevelFlgs LogLevel, writer io.Writer) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.Lock()
		logger.writer = writer
		logger.Unlock()
	}
}

//DisableFilename is
func (thisLabelingLogger *LabelingLogger) DisableFilename() {
	thisLabelingLogger.enableFileame = false
}

//EnableFilename a
func (thisLabelingLogger *LabelingLogger) EnableFilename() {
	thisLabelingLogger.enableFileame = true
}

//DisableTimestamp a
func (thisLabelingLogger *LabelingLogger) DisableTimestamp() {
	thisLabelingLogger.enableTimestamp = false
}

//EnableTimestamp a
func (thisLabelingLogger *LabelingLogger) EnableTimestamp() {
	thisLabelingLogger.enableTimestamp = true
}
