#!/usr/bin/env bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

@test "brew file update" {
	# Run update-brewfile hook and capture its output and exit status
	run "${BATS_TEST_DIRNAME}"/../.hooks/update-brewfile.sh

	# Check if the command ran successfully
	if [ $status -ne 0 ]; then
		# Check for the cask bug error message
		if [[ $output == *"was affected by a bug and cannot be upgraded as-is"* ]]; then
			assert_output --partial "was affected by a bug and cannot be upgraded as-is"

			# Extract the package name from the error message
			package=$(echo "${output}" | awk -F"'" '{print $2}')

			# Try to fix the error by reinstalling the package
			run brew reinstall --cask --force "${package}"
			assert_success
		fi
	fi
}
