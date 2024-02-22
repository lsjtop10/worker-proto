package job

import (
	"context"
	"math/rand"
	"time"
)

type GeneratorJob struct {
	Job

	in  chan any `type:"null"`
	out chan any `type:"int"`
}

func NewGeneratorJob() *GeneratorJob {
	return &GeneratorJob{
		out: make(chan any)}
}

func (j *GeneratorJob) Execute(ctx context.Context, msg chan Message) (outMsg chan Message) {
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				j.out <- (int)(rand.Float32() * 10)
				time.Sleep(time.Second * 1)
			}
		}
	}()

	return
}

func (j *GeneratorJob) SetInput(input chan any) {
	j.in = input
}

func (j *GeneratorJob) Output() chan any {
	return j.out
}
