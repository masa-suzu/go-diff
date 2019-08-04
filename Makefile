test:
	golangci-lint run --enable-all --disable gochecknoinits --disable gochecknoglobals
	go test  -cover ./... -v
