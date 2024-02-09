package pipeline

import "context"

type WorkerPipeline struct {
	tradeFetch     Fetcher
	modelExecute   Executer
	resultAnalysis ResultAnalyzer
	transmit       Transmitter

	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorkerPipeline(
	fetcher Fetcher,
	executer Executer,
	resultAnalyzer ResultAnalyzer,
	transmit Transmitter) *WorkerPipeline {
	return &WorkerPipeline{
		tradeFetch:     fetcher,
		modelExecute:   executer,
		resultAnalysis: resultAnalyzer,
		transmit:       transmit,
	}

}

func (w *WorkerPipeline) Run() {
	w.ctx, w.cancel = context.WithCancel(context.Background())

	data := w.tradeFetch.Fetch(w.ctx)
	res := w.modelExecute.ExecuteModel(w.ctx, data)
	action := w.resultAnalysis.Analyze(w.ctx, res)
	w.transmit.Transmit(w.ctx, action)
}

func (w *WorkerPipeline) Stop() {
	w.cancel()
}
