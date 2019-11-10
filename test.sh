#!/bin/bash


if ! [ -x "$(command -v ginkgo)" ]; then
  echo "'ginkgo' required to run tests: https://onsi.github.io/ginkgo/"
  exit 1
fi

ginkgo -mod vendor -r
