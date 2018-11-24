run:
	find -type f | egrep -i "*.go|*.yml|*.js|*.html" | entr -r go run *.go

build:
	mkdir -p dist
	CGO_ENABLED=1 GOOS=linux go build -v -a -ldflags '-extldflags "-static"' -o dist/vertex