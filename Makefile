test:
	gofmt -s -w .
	goimports -l -w .
	golangci-lint run --enable-all --disable gochecknoinits --disable gochecknoglobals
	go test  -cover ./... -v
