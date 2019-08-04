test:
	goimports -l ./
	go test  -cover ./... -v
