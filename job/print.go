package job

import "context"

type PrintJob struct {
	Job

	in  chan any `type:int`
	out chan any `type:null`
}

func NewPrintJob() *PrintJob {
	return &PrintJob{out: make(chan any)}
}

func (j *PrintJob) Execute(ctx context.Context, msg chan Message) (outMsg chan Message) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case num := <-j.in:
				println(num.(int))
			}
		}
	}()

	return
}

func (j *PrintJob) SetInput(input chan any) {
	j.in = input
}

func (j *PrintJob) Output() chan any {
	return j.out
}
