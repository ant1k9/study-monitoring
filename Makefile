.PHONY=test
test:
	export GO_ENV=test
	- soda create
	soda migrate
	GO_ENV=test go test -v -count=1 -p=1 ./...

.PHONY=dev
dev:
	export GO_ENV=development
	buffalo dev
