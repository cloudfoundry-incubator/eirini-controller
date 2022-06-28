#!/bin/bash

set -euo pipefail

BASEDIR="$(cd "$(dirname "$0")"/.. && pwd)"
readonly BASEDIR
if [ -z ${EIRINIUSER_PASSWORD+x} ]; then
  EIRINIUSER_PASSWORD="$(pass eirini/docker-hub)"
fi

main() {
  pushd "$BASEDIR"/tests/eats >/dev/null || exit 1
  go run github.com/onsi/ginkgo/v2/ginkgo -p -r --keep-going --randomize-all --randomize-suites --timeout=20m --slow-spec-threshold=30s "$@"
  popd >/dev/null || exit 1
}

main "$@"
