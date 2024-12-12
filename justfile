export bin := "adctl"

set dotenv-load := false

default:
    just --list

coverage:
    go test ./cmd -coverprofile=coverage.out
    go tool cover -html=coverage.out

run *ARGS: build
    ./$bin {{ ARGS }}

test: mac
    go test ./cmd -test.v

testAll: test testCLI

# TODO: this needs to be rewritten
testCLI: build

#    ./$bin status
#    ./$bin enable
#    ./$bin status
#    ./$bin disable
#    ./$bin status
#    ./$bin disable 15s
#    ./$bin status
#    ./$bin enable
#    ./$bin status
#    ./$bin toggle
#    ./$bin status
#    ./$bin toggle
#    ./$bin status
#    ./$bin getlog | jq '.oldest'
#    ./$bin getlog 42 | jq '.data | length'

fmt:
    just --unstable --fmt
    goimports -l -w .
    go fmt

linux:
    GOOS=linux GOARCH=amd64  go build -o build/adctl-linux -ldflags "-s -w" . 

mac:
    GOOS=darwin GOARCH=arm64  go build -o build/adctl-mac-arm -ldflags "-s -w" . 
    ln -fs build/adctl-mac-arm ./$bin

windows:
    GOOS=windows GOARCH=amd64  go build -o build/adctl-amd64.exe -ldflags "-s -w" . 
    GOOS=windows GOARCH=386  go build -o build/adctl-386.exe -ldflags "-s -w" . 

build: fmt mac

multibuild: fmt mac linux windows

clean:
    go clean -testcache
    go mod tidy
    rm -f $bin adctl-mac-arm adctl-linux

install: test
    cp ./$bin ~/bin/
