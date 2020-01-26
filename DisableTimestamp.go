package labelinglog

//DisableTimestamp a
func (thisLabelingLogger *LabelingLogger) DisableTimestamp() {
	thisLabelingLogger.enableTimestamp = false
}
