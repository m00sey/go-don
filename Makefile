.PHONY: all
all: clean unit-test

ROOT=$(abspath .)

.PHONY: unit-test
unit-test:
	@scripts/check_unit.sh

.PHONY: clean
clean:
	rm -f bin/*

.PHONY: build
build:
	@echo 'Building api...'
	cd cmd/api && go build -o $(ROOT)/bin/api