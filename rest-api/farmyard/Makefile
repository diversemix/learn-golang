.PHONY: clean build

clean:
	go clean
	go clean -testcache

build:
	go build -o farmyard-server ./main.go ./config.go
