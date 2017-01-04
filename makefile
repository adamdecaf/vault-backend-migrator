.PHONY: build test

vet:
	go tool vet .

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/vault-backend-migrator-linux .
osx:
	GOOS=darwin GOARCH=386 go build -o bin/vault-backend-migrator-osx .
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/vault-backend-migrator.exe .

build: vet osx linux win

test: build
	go test -v ./...
