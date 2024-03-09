package pipelineBuilder

import (
	"worker-proto-01/job"
	jobManager "worker-proto-01/job-manager"
	"worker-proto-01/pipeline"
)

type PipelineOpts struct {
	IsRealtimeTrade bool
	ModelType       string
	AnalyzeStrategy string
}

type PipelineBuilder struct {
	jobs *jobManager.JobProviderManager

	fetchJob     job.Job
	modelExecJob job.Job
	adaptJob     job.Job
	resDecodeJob job.Job
	transmitJob  job.Job
}

func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		jobs: jobManager.NewJobProviderManager(),
	}
}

func (b *PipelineBuilder) Build(opts PipelineOpts, UserParams map[string]float32) (*pipeline.Pipeline, error) {

	b.jobs.Search(jobManager.JobTags{
		"Type": "fetch",
	})

	return pipeline.NewPipeline(
		b.fetchJob,
		b.modelExecJob,
		b.adaptJob,
		b.resDecodeJob,
		b.transmitJob,
	)
}
