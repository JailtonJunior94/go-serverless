.PHONY: build deploy

build: 
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/api api/main.go

deploy: build
	sls deploy --verbose
