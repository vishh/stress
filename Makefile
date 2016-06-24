all: build docker

build:
	@GOBIN=`pwd` CGO_ENABLED=0 go install --ldflags '-extldflags "-static"'

docker:
	@docker build -t vish/stress .

.PHONY: docker build all
