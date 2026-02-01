.PHONY: build
build:
	tailwindcss -i ./pkg/tcp/routes/static/css/input.css -o ./pkg/tcp/routes/static/css/output.css
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
test: check build
	go test -v ./... -coverprofile=coverage.txt
	find ./pkg/tcp/routes/static/css/output.css
	git diff -- ./pkg/tcp/routes/static/css/output.css
	git diff --quiet -- ./pkg/tcp/routes/static/css/output.css

.PHONY: time
time:
	@touch kod
	sntp -K kod localhost
	@rm kod
