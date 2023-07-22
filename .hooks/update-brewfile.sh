#!/bin/bash
set -e

output=$(brew file update --verbose debug --no-appstore)
exit_code=$?
echo "${output}"
exit ${exit_code}
