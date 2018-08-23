GOPATH=$(shell realpath "./../..")
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell GOPATH=$(GOPATH) go list ./...)
GOFILES := $(shell find . -name "*.go" -type f)
TESTOUTPUT := coverage.out
OUTPUT := bin/currency-converter
VERSION := 1.0.0-beta1

all: build

build: fmt-check test
	@echo "Building sources"; \
	go build -ldflags='-X main.version=$(VERSION)' -o $(OUTPUT);

testing:
	@echo "Run unit test"; \
	echo "mode: count" > $(TESTOUTPUT); \
	for PKG in $(PACKAGES); do \
		go test -v -cover -covermode=atomic -coverprofile=profile.out $$PKG; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode:" >> $(TESTOUTPUT); \
			rm profile.out; \
		fi; \
	done;

report:
	@echo "Run report"; \
	go tool cover -html=$(TESTOUTPUT); \
	

fmt:
	@echo "Formating sources"; \
	$(GOFMT) -w $(GOFILES);

fmt-check:
	@echo "Check source code formatting"; \
	diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

clean:
	@echo "Cleaning output files"; \
	rm $(TESTOUTPUT); \
	rm $(OUTPUT);