export bin := "adctl"

set dotenv-load := false

# test for quoted args to work, didn't do anything. doesn't matter much.
#set positional-arguments

default:
    just --list

coverage:
    go test ./cmd -coverprofile=coverage.out
    go tool cover -html=coverage.out

# TODO Need to clean all this up.

# run *ARGS: mac-notest
run *ARGS: build
    ./$bin {{ ARGS }}

qbuild:
    goreleaser build --single-target --snapshot --clean
    ln -fs dist/adctl_darwin_arm64_v8.0/adctl ./$bin

qrun *ARGS: qbuild
    ./$bin {{ ARGS }}

qinstall: qbuild
    cp ./$bin ~/bin/

# TODO I hate that I need to install this in my path first
#  but I can't figure out how to get tescript to use ./adctl and stop searching my path
#  also tried 'env PATH=$PATH:$PWD' and that didn't work

# urgh.
test:
    go test ./cmd

testv:
    go test ./cmd -test.v

testall: test testcli

testcli: mac
    ./$bin status

#    ./$bin status enable
#    ./$bin status
#    ./$bin status disable
#    ./$bin status
#    ./$bin status disable 15s
#    ./$bin status
#    ./$bin status enable
#    ./$bin status
#    ./$bin status toggle
#    ./$bin status
#    ./$bin status toggle
#    ./$bin status
#    ./$bin log get | jq '.oldest'
#    ./$bin log get 42 | jq '.data | length'

fmt:
    just --unstable --fmt
    goimports -l -w .
    go fmt

mac: test
    goreleaser build --single-target --snapshot --clean
    ln -fs dist/adctl_darwin_arm64_v8.0/adctl ./$bin

mac-notest:
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

install: mac
    cp ./$bin ~/bin/

# TODO: prompt for a tag here?
# not for now
# git tag -a v0.1.0 -m "first release, test of goreleaser"
# git push origin v0.1.0
# takes two arguments. first is tag (v0.1.0), second is tag description.

# TODO: what do I do if I have uncommitted changes?
release arg1 arg2: testall
    rm -rf dist/
    #git tag -a {{ arg1 }} -m "{{ arg2 }}"
    git tag {{ arg1 }}
    git push origin {{ arg1 }}
    goreleaser release
