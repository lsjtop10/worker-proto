package main

import (
	"time"
	"worker-proto-01/job"
	"worker-proto-01/pipeline"
)

func main() {
	p, err := pipeline.NewPipeline(
		job.NewGeneratorJob(),
		job.NewMultiplyJob(job.UserParams{
			"mul": 2.0,
		}),
		job.NewMultiplyJob(job.UserParams{
			"mul": 3.0,
		}),
		job.NewPrintJob(),
	)

	if err != nil {
		println(err)
		return
	}

	p.Run()

	time.Sleep(6 * time.Second)

	p.Stop()

	time.Sleep(1 * time.Second)

}
