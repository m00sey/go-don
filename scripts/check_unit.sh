#!/bin/bash

set -e
echo -n > coverage.txt
echo "Running $0"

GO_TEST_CMD="go test"

ROOT=$(pwd)
touch "$ROOT"/coverage.out

amend_coverage_file () {
if [ -f profile.out ]; then
     cat profile.out >> "$ROOT"/coverage.out
     rm profile.out
fi
}

PKGS=$(go list github.com/m00sey/go-don/... 2> /dev/null | grep -v vendor | grep -v mocks | grep -v cmd)
$GO_TEST_CMD $PKGS -count=1 -race -coverprofile=profile.out -covermode=atomic -timeout=10m
if [ -f profile.out ]; then
  cat profile.out >>coverage.txt
  rm profile.out
fi