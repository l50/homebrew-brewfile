#!/bin/bash
set -e

# Find the project root
root_dir=$(git rev-parse --show-toplevel)

# Run bats tests
output=$(bats "${root_dir}/tests/brew_upgrade_tests.bats" 2>&1)
exit_code=$?

# Check if the error message indicates another brew update is in progress
if [[ $output == *"Another active Homebrew update process is already in progress"* ]]; then
	echo "Another Homebrew update process is currently running. Please wait for it to finish before retrying."
	exit 1
fi

echo "${output}"
exit ${exit_code}
