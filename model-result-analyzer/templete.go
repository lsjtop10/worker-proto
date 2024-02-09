package resultAnalyzer

import (
	"context"
	"worker-proto-01/vo"
)

// 탬플릿 메서드 패턴 사용
type CustomAnalyzer interface {
	decode(ctx context.Context, data chan any) chan any
	decide(ctx context.Context, data chan any) chan vo.Action
}

type ResultAnalyzer struct {
	customAnalyzer CustomAnalyzer
}

func NewResultAnalyzer(custom CustomAnalyzer) *ResultAnalyzer {
	return &ResultAnalyzer{
		customAnalyzer: custom,
	}
}

func (a *ResultAnalyzer) Analyze(ctx context.Context, data chan any) chan vo.Action {
	res := a.customAnalyzer.decode(ctx, data)
	return a.customAnalyzer.decide(ctx, res)
}
