GO_BUILD=CGO_ENABLED=0 go build -i --ldflags="-w"

lint:
	golangci-lint run

all: ci
ci: lint build

.PHONY: all lint ci

#################################################
# Building
#################################################

build: bin/shipit

bin/shipit: shipit.go
	$(GO_BUILD) -o bin/shipit

.PHONY: build

#################################################
# Cleaning
#################################################

clean:
	rm bin/shipit