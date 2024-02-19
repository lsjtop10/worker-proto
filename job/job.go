package job

import "context"

type Job interface {
	Execute(ctx context.Context, msg chan Message) (outMsg chan Message)

	SetInput(chan any)
	Output() chan any
}
