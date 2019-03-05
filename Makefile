## Application makefile

start:
	go run application.go

check:
	go test ./...

setup:
	sh setup.sh
