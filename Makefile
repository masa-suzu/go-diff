test:
	gofmt -s -w .
	goimports -l -w ./
	golint ./
	go test  -cover ./... -v
