
run:
	go run main.go

build: wire-job
	go build .

wire-job:
	wire ./job-manager/inject/wire.go