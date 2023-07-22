//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/l50/goutils/v2/dev/lint"
	mageutils "github.com/l50/goutils/v2/dev/mage"
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

// Setup initializes and configures the homebrew-brewfile repo
func Setup() error {
	// Make sure we are in the repo root
	repoRoot, err := git.RepoRoot()
	if err != nil {
		return err
	}

	// Get the current working directory
	// so that we can return to it in the event
	// that we are not in the repo root
	cwd := sys.Gwd()
	if cwd != repoRoot {
		if err := sys.Cd(repoRoot); err != nil {
			return err
		}
	}
	defer sys.Cd(cwd)

	home, err := sys.GetHomeDir()
	if err != nil {
		return err
	}

	err = sys.Cp(filepath.Join(repoRoot, "Brewfile"),
		filepath.Join(home, ".brewfile", "Brewfile", "Brewfile"))
	if err != nil {
		return fmt.Errorf("failed to set brewfile repo: %v", err)
	}

	return nil
}

// Update updates various components of
// the Brewfile on the local system.
func Update() error {
	commands := []string{
		"update",
		"upgrade",
	}

	for _, cmd := range commands {
		fmt.Println("Running brew", cmd)
		if _, err := sys.RunCommand("brew", cmd); err != nil {
			return fmt.Errorf("failed to update brewfile: %v", err)
		}
	}

	return nil
}
