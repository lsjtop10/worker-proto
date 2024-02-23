package job

import (
	"context"
	"time"
)

type GeneratorJob struct {
	Job

	data []int

	in  chan any `type:"null"`
	out chan any `type:"int"`
}

func NewGeneratorJob() *GeneratorJob {
	return &GeneratorJob{
		data: []int{1, 2, 3, 4, 5},
		out:  make(chan any)}
}

func (j *GeneratorJob) Execute(ctx context.Context, msg chan Message) (outMsg chan Message) {
	go func() {
		for _, data := range j.data {
			select {
			case <-ctx.Done():
				return
			default:
				j.out <- data
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
