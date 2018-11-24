run:
	find -type f | egrep -i "*.go|*.yml" | entr -r go run *.go

build:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -v -a \
		-ldflags '-extldflags "-static" -X main.sha1ver=$(shell git rev-parse HEAD) -X main.buildTime=$(shell date +'%Y.%m.%d-%H:%M:%S')' \
		-o dist/vertex