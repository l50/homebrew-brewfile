//go:build mage

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/l50/goutils/v2/dev/lint"
	mageutils "github.com/l50/goutils/v2/dev/mage"
	"github.com/l50/goutils/v2/fileutils"
	"github.com/l50/goutils/v2/git"
	"github.com/l50/goutils/v2/sys"
	"github.com/magefile/mage/mg"
)

func init() {
	os.Setenv("GO111MODULE", "on")
}

// InstallDeps Installs go dependencies
func InstallDeps() error {
	fmt.Println("Installing dependencies.")

	if err := lint.InstallGoPCDeps(); err != nil {
		return fmt.Errorf("failed to install pre-commit dependencies: %v", err)
	}

	if err := mageutils.InstallVSCodeModules(); err != nil {
		return fmt.Errorf("failed to install vscode-go modules: %v", err)
	}

	return nil
}

// RunPreCommit runs all pre-commit hooks locally
func RunPreCommit() error {
	fmt.Println("Updating pre-commit hooks.")
	if err := lint.UpdatePCHooks(); err != nil {
		return err
	}

	fmt.Println("Clearing the pre-commit cache to ensure we have a fresh start.")
	if err := lint.ClearPCCache(); err != nil {
		return err
	}

	fmt.Println("Running all pre-commit hooks locally.")
	if err := lint.RunPCHooks(); err != nil {
		return err
	}

	return nil
}

// RunTests runs all of the unit tests
func RunTests() error {
	mg.Deps(InstallDeps)

	if _, err := sys.RunCommand("bash", filepath.Join(".hooks", "run-bats-tests.sh")); err != nil {
		return fmt.Errorf("failed to run unit tests: %v", err)
	}

	return nil
}

// InstallRepo installs the brewfile repo
func InstallRepo() error {

	var tmpfile *os.File
	repoRoot, err := git.RepoRoot()
	if err != nil {
		return err
	}

	cwd := sys.Gwd()
	if cwd != repoRoot {
		if err := sys.Cd(repoRoot); err != nil {
			return err
		}

		output, err := sys.RunCommand("brew", "file", "set_repo", "-r", "l50/homebrew-brewfile", "-y")
		if err != nil {
			return fmt.Errorf("failed to set brewfile repo: %v", err)
		}

		// Create a temp file to write the output to
		// in the event that we need to debug
		tmpfile, err = os.CreateTemp("", "homebrew-brewfile")
		if err != nil {
			log.Printf("failed to create temp file: %v", err)
			return err
		}

		fmt.Println("Logging output to: ", tmpfile.Name())

		rf := fileutils.RealFile(tmpfile.Name())
		if err := rf.Write(output, "0644"); err != nil {
			log.Printf("failed to write to file: %v", err)
			return err
		}

		defer sys.Cd(cwd)
	}

	// Delete the temp file if it exists and we've
	// successfully installed the repo
	if err := fileutils.Delete(tmpfile.Name()); err != nil {
		fmt.Printf("failed to delete temp file: %v", err)
		return err
	}

	// Run Update to get everything in place and up-to-date
	if err := Update(); err != nil {
		return err
	}

	return nil
}

// Update executes the following actions:
// brew update
// brew upgrade
// brew cask upgrade
// brew install
// brew init
// brew pull
// brew push
func Update() error {
	// brew file update
	sys.RunCommand("brew", "file", "update")

	return nil

}
