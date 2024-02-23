
run:
	go run main.go

build:
	go build .

wire-job:
	wire ./job-manager/inject/wire.go