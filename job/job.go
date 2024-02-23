package job

import (
	"context"
	"reflect"
)

type UserParams map[string]float32

type Job interface {
	Execute(ctx context.Context, msg chan Message) (outMsg chan Message)

	SetInput(chan any)
	Output() chan any
}

func InputTypeTag(job Job) string {
	t := reflect.TypeOf(job).Elem()

	//내부 구조를 참조하는 문제가 있음. 캡슐화 원칙에 어긋나서 문제가 생길 수도...
	field, ok := t.FieldByName("in")
	if !ok {
		panic("cannot get 'in' field of " + t.String())
	}

	tag := field.Tag.Get("type")
	if tag == "" {
		panic("cannot get type tag of input channel of " + t.String())
	}

	return tag
}

func OutputTypeTag(job Job) string {
	t := reflect.TypeOf(job).Elem()
	//내부 구조를 참조하는 문제가 있음. 캡슐화 원칙에 어긋나서 문제가 생길 수도...
	field, ok := t.FieldByName("out")

	if !ok {
		panic("cannot get 'in' field of " + t.String())
	}

	tag := field.Tag.Get("type")
	if tag == "" {
		panic("cannot get type tag of input channel of " + t.String())
	}

	return tag
}
