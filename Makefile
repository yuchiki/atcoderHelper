VERSION_PATH := github.com/yuchiki/atcoderHelper/internal/cmd/ach/version
LD_FLAGS := -X '$(VERSION_PATH).version=manual-build'
LD_FLAGS += -X '$(VERSION_PATH).commit=$(shell git rev-parse HEAD)'
LD_FLAGS += -X '$(VERSION_PATH).edited=$(shell if git diff HEAD --exit-code > /dev/null; then echo "HEAD"; else echo "edited"; fi)'
LD_FLAGS += -X '$(VERSION_PATH).date=$(shell date '+%Y/%m/%d %H:%M:%S %Z')'
FLAGS := -ldflags "$(LD_FLAGS)"

.PHONY: install clean build ach test lint yamllint generate-docs dry-release

default: build test lint yamllint generate-docs

all: build test yamllint lint generate-docs dry-release

install:
	go install ./cmd/ach


build: ach gendocs


ach:
	go build ${FLAGS} -o ./bin/ach ./cmd/ach/main.go

gendocs:
	go build -o ./bin/gendocs ./cmd/gendocs/main.go

generate-docs: gendocs
	./bin/gendocs

dry-release:
	goreleaser --snapshot --skip-publish --rm-dist


test:
	go test ./...

lint:
	golangci-lint run
	go mod tidy

yamllint:
	yamllint .

clean:
	rm -rf bin dist
