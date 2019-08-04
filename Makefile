test:
	goimports -l -w ./
	go test  -cover ./... -v
