.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy
