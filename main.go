package main

import (
	"time"
	"worker-proto-01/job"
	"worker-proto-01/pipeline"
)

func main() {
	p := pipeline.NewPipeline(
		job.NewGeneratorJob(),
		job.NewMultiplyJob(2),
		job.NewMultiplyJob(1),
		job.NewPrintJob())

	p.Run()
	time.Sleep(time.Second * 10)
	p.Stop()
}
