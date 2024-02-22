package jobManager

import (
	"encoding/json"
	"worker-proto-01/job"
)

type JobProvider func(userParams map[string]float32) job.Job
type JobTags map[string]string

type JobProviderManager struct {
	Jobs map[string]JobProvider
}

func NewJobProviderManager() *JobProviderManager {
	m := &JobProviderManager{
		Jobs: make(map[string]JobProvider),
	}
	m.registerJobs(managedJobs)
	return m
}

// 주어진 조건을 만족하는 provider를 return
func (m *JobProviderManager) Search(tags JobTags) JobProvider {
	tagsToId(tags)

	provider, ok := m.Jobs[tagsToId(tags)]

	if ok {
		return provider
	} else {
		return nil
	}

}

func (m *JobProviderManager) registerJobs(jobs []taggedProvider) {
	for _, element := range jobs {
		m.Jobs[tagsToId(element.Tags)] = element.Provider
	}
}

func tagsToId(tags map[string]string) string {
	jsonBytes, _ := json.Marshal(tags)
	return string(jsonBytes)
}
