all: test vet fmt lint build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /docs/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /docs/ | xargs -L1 gofmt -l)

lint:
	go list ./... | grep -v /docs/ | xargs -L1 golint -set_exit_status

build:
	go build -o bin/app ./cmd
