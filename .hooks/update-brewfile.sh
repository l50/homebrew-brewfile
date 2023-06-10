#!/bin/bash
set -ex

output=$(brew file update --verbose 5 --appstore 0 2>&1)
exit_code=$?
echo "${output}"
exit ${exit_code}
