package pipeline

import (
	"context"
	"errors"
	"worker-proto-01/job"
)

var ErrTypeNotMatch = errors.New("pipeline: cannot build a pipeline because the types are not compatible between the jobs")

type Pipeline struct {
	fetch      job.Job
	modelExec  job.Job
	adapt      job.Job
	resAnalyze job.Job
	transmit   job.Job

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPipeline(fetch, modelExec, adapt, resAnalyze, transmit job.Job) (*Pipeline, error) {

	return &Pipeline{
		fetch:      fetch,
		modelExec:  modelExec,
		adapt:      adapt,
		resAnalyze: resAnalyze,
		transmit:   transmit,
	}, nil
}

func (p *Pipeline) Run() {
	p.ctx, p.cancel = context.WithCancel(context.Background())
	//TODO: 하나의 컴포넌트에서 발생한 메시지를 다른 데에 전달하는 좋은 방법 개발
	msg := make(chan job.Message)

	p.modelExec.SetInput(p.fetch.Output())
	p.resAnalyze.SetInput(p.modelExec.Output())
	p.transmit.SetInput(p.resAnalyze.Output())

	p.fetch.Execute(p.ctx, msg)
	p.modelExec.Execute(p.ctx, msg)
	p.resAnalyze.Execute(p.ctx, msg)
	p.transmit.Execute(p.ctx, msg)

}

func (p *Pipeline) Stop() {
	p.cancel()
}
