package main

import (
	"context"
	"time"
)

type TradeFetch struct {
}

type ModelExecute struct {
}

type ResultAnalyzer struct {
}

func main() {
	an := custom.CustomResultAnalyzerTest1{}
	ctx, cancel := context.WithCancel(context.Background())

	an.Analyze(ctx, make(chan any))

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 10)
}
