package jobManager

import "worker-proto-01/job-manager/inject"

type taggedProvider struct {
	Tags     map[string]string
	Provider JobProvider
}

// 여기에서 tag와 provider의 관계를 정의한다.
var managedJobs = []taggedProvider{
	{
		map[string]string{
			"type": "operation",
			"op":   "mul",
		},
		inject.InitializeMultiplyJob,
	},
}