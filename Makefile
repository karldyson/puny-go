VERSION = $(shell git describe)
NAME = puny

all: linux darwin_intel darwin_apple

linux:
	GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o build/linux/amd64/$(NAME) -buildvcs=true -ldflags "-X main.versionString=$(VERSION)" $(NAME).go

darwin_intel:
	GOOS=darwin GOARCH=amd64 /usr/local/go/bin/go build -o build/darwin/amd64/$(NAME) -buildvcs=true -ldflags "-X main.versionString=$(VERSION)" $(NAME).go

darwin_apple:
	GOOS=darwin GOARCH=arm64 /usr/local/go/bin/go build -o build/darwin/arm64/$(NAME) -buildvcs=true -ldflags "-X main.versionString=$(VERSION)" $(NAME).go

