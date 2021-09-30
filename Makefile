.PHONY: generate build run test serve

generate:
	@go mod tidy
	@go generate ./...
	@echo "[OK] Generate all completed!"

nowTime=$(shell date +00%y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD)
gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
build: generate
	@CGO_ENABLED=0 go build -ldflags "-s -w -X main.build=${nowTime}.${gitCID}" -o "./bin/server"
	@echo "[OK] App binary was created!"

buildcross: generate
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -X main.build=${gitTime}.${gitCID}" -o "./bin/server.amd64"
	@echo "[OK] App amd64 binary was created!"
	@CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -ldflags "-s -w -X main.build=${gitTime}.${gitCID}" -o "./bin/server.arm64"
	@echo "[OK] App arm64 binary was created!"
	@CGO_ENABLED=0 GOARCH=mips64le GOOS=linux go build -ldflags "-s -w -X main.build=${gitTime}.${gitCID}" -o "./bin/server.mips64le"
	@echo "[OK] App mips64le binary was created!"

run:
	@./bin/server

test: 
	go test -v ./...

serve: build run