#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SRC_DIR="$(dirname "$DIR")"

cd "$SRC_DIR/app/pkg"

# echo "Running unit tests for app (go test)"
# go test -v ./...

echo "Running behavioral tests for app (ginkgo)"
GINKGO_OUT="$SRC_DIR/app/tmp/ginkgo"
 
ginkgo -v -r --output-dir=$GINKGO_OUT --json-report=report.ginkgo.json

echo "Generating summary table from report.ginkgo.json"
jq -r '.[] | [.SuiteDescription, .SuiteSucceeded, .RunTime] | @tsv' $GINKGO_OUT/report.ginkgo.json | column -t -s $'\t'

echo "Testing executed completed"
