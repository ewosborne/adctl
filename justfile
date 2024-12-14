export bin := "adctl"

set dotenv-load := false

default:
    just --list

coverage:
    go test ./cmd -coverprofile=coverage.out
    go tool cover -html=coverage.out

run *ARGS: build
    ./$bin {{ ARGS }}

test: 
    go test ./cmd -test.v

testall: test testcli

testcli: build
    ./$bin status
    ./$bin status enable
    ./$bin status
    ./$bin status disable
    ./$bin status
    ./$bin status disable 15s
    ./$bin status
    ./$bin status enable
    ./$bin status
    ./$bin status toggle
    ./$bin status
    ./$bin status toggle
    ./$bin status
    ./$bin log get | jq '.oldest'
    ./$bin log get 42 | jq '.data | length'

fmt:
    just --unstable --fmt
    goimports -l -w .
    go fmt

#linux: 
#    #GOOS=linux GOARCH=amd64  go build -o build/adctl-linux -ldflags "-s -w" . 
#
#mac: 
#    #GOOS=darwin GOARCH=arm64  go build -o build/adctl-mac-arm -ldflags "-s -w" . 
#    ln -fs dist/adctl_darwin_arm64_v8.0/adctl ./$bin
#
#windows: build
#    #GOOS=windows GOARCH=amd64  go build -o build/adctl-amd64.exe -ldflags "-s -w" . 
#    GOOS=windows GOARCH=386  go build -o build/adctl-386.exe -ldflags "-s -w" . 

mac: test
    goreleaser build --single-target --snapshot --clean
    ln -fs dist/adctl_darwin_arm64_v8.0/adctl ./$bin

build: test
    goreleaser release --snapshot --clean
    ln -fs dist/adctl_darwin_arm64_v8.0/adctl ./$bin

clean:
    go clean -testcache
    go mod tidy
    rm -f $bin 
    rm -rf dist

install: test
    cp ./$bin ~/bin/
