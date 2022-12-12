#!/usr/bin/bash

# Run Python tests
tests=$(find /src/test/ -type f -iname "*_test.py")
for test in $tests; do
  python3 "$test"
done
