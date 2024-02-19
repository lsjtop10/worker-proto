package job

import (
	"context"
)

type MultiplyJob struct {
	Job

	num int

	in  chan any
	out chan any
}

func NewMultiplyJob(num int) *MultiplyJob {
	return &MultiplyJob{
		num: num,
		out: make(chan any)}
}

func (j *MultiplyJob) Execute(ctx context.Context, msg chan Message) (outMsg chan Message) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case num := <-j.in:
				j.out <- num.(int) * j.num
			}
		}
	}()

	return
}

func (j *MultiplyJob) SetInput(input chan any) {
	j.in = input
}

func (j *MultiplyJob) Output() chan any {
	return j.out
}
