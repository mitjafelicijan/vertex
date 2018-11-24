VERSION = 0.1
BUILD_TIME = $(shell date +'%Y.%m.%d-%H:%M:%S')
SHA1_VER = $(shell git rev-parse HEAD)

run:
	find -type f \( -name "*.go" -o -name "*.js" \) | entr -r go run *.go

build:
	mkdir -p dist
	cp vertex.yml dist/
	CGO_ENABLED=0 GOOS=linux go build -v -a \
		-ldflags '-extldflags "-static"' \
		-ldflags '-X main.buildTime=$(BUILD_TIME) -X main.sha1ver=$(SHA1_VER)' \
		-o dist/vertex

publish:
	git tag v$(VERSION)
	git push origin --tags
	goreleaser release --rm-dist
