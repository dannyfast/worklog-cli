
.PHONY: all build clean deps

BINARY_NAME=worklog
GO=go

all: build

build: deps
	$(GO) build -o $(BINARY_NAME)

deps:
	$(GO) get -u github.com/spf13/cobra
	$(GO) mod tidy

clean:
	$(GO) clean
	rm -f $(BINARY_NAME)
