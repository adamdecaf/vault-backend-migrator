.PHONY: build check test

linux: linux_amd64
linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/vault-backend-migrator-linux-amd64 github.com/adamdecaf/vault-backend-migrator

osx: osx_amd64
osx_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/vault-backend-migrator-osx-amd64 github.com/adamdecaf/vault-backend-migrator

win: win_64
win_64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/vault-backend-migrator-amd64.exe github.com/adamdecaf/vault-backend-migrator

dist: build linux osx win

check:
	go vet ./...
	go fmt ./...

test: check dist
	go test -v ./...

ci: check dist test

build: check
	go build -o cert-manage github.com/adamdecaf/vault-backend-migrator
	@chmod +x vault-backend-migrator
