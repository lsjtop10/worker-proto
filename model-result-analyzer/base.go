package resultAnalyzer

import (
	"context"
	"worker-proto-01/vo"
)

type DecodeFunc func(ctx context.Context, data chan any) chan any
type DecideFunc func(ctx context.Context, data chan any) chan vo.Action

type ResultAnalyzerBase struct {
	decodeFunc DecodeFunc
	decideFunc DecideFunc
}

func NewResultAnalyzer(decode DecodeFunc, decide DecideFunc) *ResultAnalyzerBase {
	return &ResultAnalyzerBase{
		decodeFunc: decode,
		decideFunc: decide,
	}
}

func (a *ResultAnalyzerBase) Analyze(ctx context.Context, data chan any) chan vo.Action {

	res := a.decodeFunc(ctx, data)
	action := a.decideFunc(ctx, res)

	return action
}
