//go:build wireinject
// +build wireinject

package inject

import (
	"worker-proto-01/job"

	"github.com/google/wire"
)

// 여기에서 Job의 DI를 수행한다. 여기서는 Job과 다른 어뎁터, 엔티티간의 의존 관계를 설정한다.
func InitializeMultiplyJob(userParams map[string]float32) job.Job {
	wire.Build(wire.NewSet(
		job.NewMultiplyJob,
		wire.Bind(new(job.Job), new(*job.MultiplyJob))))
	return &job.MultiplyJob{}
}