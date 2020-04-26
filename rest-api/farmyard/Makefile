.PHONY: clean build test

clean:
	go clean
	go clean -testcache

build:
	go build -o farmyard-server ./main.go ./config.go

test:
	go vet ./...
	go clean -testcache
	go test -v --cover ./...