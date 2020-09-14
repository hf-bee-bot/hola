REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')

pack:
	pack --no-color build hferentschik/hello:$(REV) --builder paketobuildpacks/builder:tiny --publish

build:
	go build -o hola-server ./...