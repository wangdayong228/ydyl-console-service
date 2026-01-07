package enums

type PipelineStatus string

const (
	PipelineStatusRunning PipelineStatus = "running"
	PipelineStatusSuccess PipelineStatus = "success"
	PipelineStatusFailed  PipelineStatus = "failed"
)
