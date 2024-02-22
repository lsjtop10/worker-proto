package job

import (
	"context"
)

type UserParams map[string]float32

type Job interface {
	Execute(ctx context.Context, msg chan Message) (outMsg chan Message)

	SetInput(chan any)
	Output() chan any
}
