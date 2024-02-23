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
	resAnalyze job.Job
	transmit   job.Job

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPipeline(fetch, modelExec, resAnalyze, transmit job.Job) (*Pipeline, error) {

	if !(isTypeCompatible(fetch, modelExec) &&
		isTypeCompatible(modelExec, resAnalyze) &&
		isTypeCompatible(resAnalyze, transmit)) {

		return nil, ErrTypeNotMatch
	}

	return &Pipeline{
		fetch:      fetch,
		modelExec:  modelExec,
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

func isTypeCompatible(prev, next job.Job) bool {

	// 다음 job의 입력 타입이 any이거나
	// 이전 job의 출력 타입과 다음 job의 출력 타입이 같고
	// 이전 job의 출력 타입과 다음 job의 타입의 입력 타입이 모두 Null이 아닌지 검사
	return job.InputTypeTag(next) == "any" ||
		(job.OutputTypeTag(prev) == job.InputTypeTag(next)) &&
			!(job.OutputTypeTag(prev) == "null" || job.InputTypeTag(next) == "null")

}
