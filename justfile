export bin := "adctl"

set dotenv-load := false

default:
    just --list

coverage:
    go test ./cmd -coverprofile=coverage.out
    go tool cover -html=coverage.out

run *ARGS: build
    ./$bin {{ ARGS }}

test: build
    go test ./cmd -test.v

testAll: test testCLI

testCLI: build
    ./$bin status
    ./$bin enable
    ./$bin status
    ./$bin disable
    ./$bin status
    ./$bin disable 15s
    ./$bin status
    ./$bin enable
    ./$bin status
    ./$bin toggle
    ./$bin status
    ./$bin toggle
    ./$bin status
    ./$bin getlog | jq '.oldest'
    ./$bin getlog 42 | jq '.data | length'

fmt:
    just --unstable --fmt
    goimports -l -w .
    go fmt

build: fmt
    go build -ldflags "-s -w" .

clean:
    go clean -testcache
    go mod tidy
    rm -f $bin

install: build test
    cp ./$bin ~/bin/
