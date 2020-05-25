.PHONY=all

test:
	export GO_ENV=test
	- soda create
	soda migrate
	GO_ENV=test go test -v -p=1 ./... -coverprofile=coverage.txt -covermode=atomic

cover:
	go tool cover -func=coverage.txt

dev:
	export ADDR=0.0.0.0
	export GO_ENV=development
	buffalo dev
