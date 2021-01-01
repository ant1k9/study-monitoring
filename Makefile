.PHONY=all

test:
	export GO_ENV=test
	- soda create -e test
	soda migrate -e test
	GO_ENV=test go test -v -p=1 ./... -coverprofile=coverage.txt -covermode=atomic

cover:
	go tool cover -func=coverage.txt

dev:
	export ADDR=0.0.0.0
	export PORT=3000
	export GO_ENV=development
	buffalo dev
