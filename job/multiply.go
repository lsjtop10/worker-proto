package job

import (
	"context"
	"reflect"
)

type MultiplyJob struct {
	Job

	num int `default:"1.0"`

	in  chan any `type:"int"`
	out chan any `type:"int"`
}

func NewMultiplyJob(userParams map[string]float32) *MultiplyJob {
	return &MultiplyJob{
		num: int(userParams["mul"]),
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

func (j MultiplyJob) InputTypeTag() string {
	st := reflect.TypeOf(j)
	field, _ := st.FieldByName("in")
	return field.Tag.Get("type")
}

func (j MultiplyJob) OutputTypeTag() string {
	st := reflect.TypeOf(j)
	field, _ := st.FieldByName("out")
	return field.Tag.Get("type")
}
