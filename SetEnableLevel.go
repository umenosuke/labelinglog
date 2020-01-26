package labelinglog

// SetEnableLevel a
func (thisLabelingLogger *LabelingLogger) SetEnableLevel(targetLevelFlgs LogLevel) {
	for _, logger := range thisLabelingLogger.loggers {
		logger.isEnable = targetLevelFlgs&logger.flg != 0
	}
}
