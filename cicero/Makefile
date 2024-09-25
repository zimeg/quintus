.PHONY: build
build:
	go build -o cicero

.PHONY: check
check:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm cicero
	rm kod

.PHONY: start
start: build
	HTTP_PORT=8080 ./cicero

.PHONY: test
test: check
	go test -v ./...

.PHONY: time
time:
	@touch kod
	sntp -K kod localhost
	@rm kod
