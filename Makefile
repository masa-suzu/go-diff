.PHONY: test
test:
	gofmt -s -w .
	go run "golang.org/x/tools/cmd/goimports" -l -w .
	go run "github.com/golangci/golangci-lint/cmd/golangci-lint" run --enable-all --disable gochecknoinits --disable gochecknoglobals
	go test  -cover ./... -v
