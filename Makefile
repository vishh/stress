all: build docker

build:
	@GOBIN=`pwd` CGO_ENABLED=0 go install --ldflags '-extldflags "-static"'
	@chmod a+rx memstress

docker:
	@docker build -t vish/memstress .

.PHONY: docker build all
