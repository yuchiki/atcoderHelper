VERSION_PATH := github.com/yuchiki/atcoderHelper/cmd/atcoderHelper/cmd
LD_FLAGS := -X '$(VERSION_PATH).version=manual-build'
LD_FLAGS += -X '$(VERSION_PATH).commit=$(shell git rev-parse HEAD)'
LD_FLAGS += -X '$(VERSION_PATH).edited=$(shell if git diff HEAD --exit-code > /dev/null; then echo "HEAD"; else echo "edited"; fi)'
LD_FLAGS += -X '$(VERSION_PATH).date=$(shell date '+%Y/%m/%d %H:%M:%S %Z')'
FLAGS := -ldflags "$(LD_FLAGS)"

.PHONY: ach test lint yamllint clean

all: yamllint lint ach test

ach: test
	go build ${FLAGS} -o ./bin/ach ./cmd/atcoderHelper/main.go

# go vet は保守的なcheckをするので、厳しすぎるかもしれない。問題があるようなら、suggestionだけして、CIの成否には関与しないように変更する。
test:
	go test ./...

lint:
	golangci-lint run

yamllint:
	yamllint .


clean:
	rm -rf bin
