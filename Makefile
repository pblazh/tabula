.PHONY: build test clean install lint

build: test build-darwin-arm64 build-darwin-amd64 build-linux-arm64 build-linux-amd64 build-linux-386 build-windows-arm64 build-windows-amd64 build-windows-386 build-wasm

build-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/csvss ./cmd/cli

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/csvss ./cmd/cli


build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/csvss ./cmd/cli

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/csvss ./cmd/cli

build-linux-386:
	env GOOS=linux GOARCH=386 go build -o bin/linux/amd64/csvss ./cmd/cli


build-windows-arm64:
	env GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/csvss.exe ./cmd/cli

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/csvss.exe ./cmd/cli

build-windows-386:
	env GOOS=windows GOARCH=386 go build -o bin/windows/386/csvss.exe ./cmd/cli


build-wasm:
	env GOOS=js GOARCH=wasm go build -o bin/wasm/csvss ./cmd/cli


test:
	go test -v -cover ./...

clean:
	rm -rf bin/

install:
	go install ./cmd/cli

lint:
	golangci-lint run

fmt:
	@echo "TODO: Implement format target"
	go fmt ./...

# TODO: Add benchmark target
benchmark:
	@echo "TODO: Implement benchmark tests"
	# go test -bench=. ./...
