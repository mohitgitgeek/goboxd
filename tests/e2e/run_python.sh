#!/usr/bin/env bash
set -euo pipefail

output="$(python3 -c 'print(1+1)')"
if [ "$output" != "2" ]; then
  echo "expected python output to be 2, got: $output"
  exit 1
fi

echo "python e2e stub passed"
