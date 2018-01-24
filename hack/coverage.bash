#!/bin/bash
# https://stackoverflow.com/questions/45776043/codeclimate-test-coverage-formatter-for-golang#45776499
# https://github.com/chartmogul/chartmogul-go/blob/master/.test-coverage.sh

for pkg in $(go list ./... | grep -v vendor); do
    go test -coverprofile=$(echo $pkg | tr / -).cover $pkg
done
echo "mode: set" > c.out
grep -h -v "^mode:" ./*.cover >> c.out
rm -f *.cover
