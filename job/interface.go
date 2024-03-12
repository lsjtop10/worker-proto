package job

import (
	"context"
)

type UserParams map[string]float32

type Job interface {
	//
	Execute(ctx context.Context)

	//
	SetInputChan(chan any)
	//
	OutputChan() chan any
}

/*
고루틴의 소유권은 각 Job이 가져야 한다.
*/
