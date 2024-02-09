package pipeline

import (
	"context"
	"worker-proto-01/vo"
)

type Fetcher interface {
	Fetch(ctx context.Context) chan []vo.ProductAggregate
}

type Executer interface {
	ExecuteModel(ctx context.Context, data chan []vo.ProductAggregate) chan any
}

type ResultAnalyzer interface {
	Analyze(ctx context.Context, data chan any) chan vo.Action
}

type Transmitter interface {
	Transmit(ctx context.Context, action chan vo.Action)
}
