VERSION_PATH := github.com/yuchiki/atcoderHelper/cmd/atcoderHelper/cmd
LD_FLAGS := -X '$(VERSION_PATH).version=manual-build'
LD_FLAGS += -X '$(VERSION_PATH).commit=$(shell git rev-parse HEAD)'
LD_FLAGS += -X '$(VERSION_PATH).edited=$(shell git diff HEAD || echo "edited")'
LD_FLAGS += -X '$(VERSION_PATH).date=$(shell date '+%Y/%m/%d %H:%M:%S %Z')'
FLAGS := -ldflags "$(LD_FLAGS)"

.PHONY: atcoderHelper


atcoderHelper:
	go build ${FLAGS} -o ./bin/ach ./cmd/atcoderHelper/main.go
