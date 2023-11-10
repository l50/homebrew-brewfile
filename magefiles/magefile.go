//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/l50/goutils/v2/dev/lint"
	mageutils "github.com/l50/goutils/v2/dev/mage"
	"github.com/l50/goutils/v2/docs"
	fileutils "github.com/l50/goutils/v2/file/fileutils"
	"github.com/l50/goutils/v2/git"
	"github.com/l50/goutils/v2/sys"
	"github.com/spf13/afero"
)

func init() {
	os.Setenv("GO111MODULE", "on")
}

// InstallDeps installs the Go dependencies necessary for developing
// on the project.
//
// Example usage:
//
// ```bash
// mage installdeps
// ```
//
// **Returns:**
//
// error: An error if any issue occurs while trying to
// install the dependencies.
func InstallDeps() error {
	fmt.Println("Installing dependencies.")

	cwd := sys.Gwd()
	if err := sys.Cd("magefiles"); err != nil {
		return fmt.Errorf("failed to cd into magefiles directory: %v", err)
	}

	if err := mageutils.Tidy(); err != nil {
		return fmt.Errorf("failed to install dependencies: %v", err)
	}

	if err := sys.Cd(cwd); err != nil {
		return fmt.Errorf("failed to cd into project root directory: %v", err)
	}

	if err := lint.InstallGoPCDeps(); err != nil {
		return fmt.Errorf("failed to install pre-commit dependencies: %v", err)
	}

	if err := mageutils.InstallVSCodeModules(); err != nil {
		return fmt.Errorf("failed to install vscode-go modules: %v", err)
	}

	return nil
}

// RunPreCommit updates, clears, and executes all pre-commit hooks
// locally. The function follows a three-step process:
//
// First, it updates the pre-commit hooks.
// Next, it clears the pre-commit cache to ensure a clean environment.
// Lastly, it executes all pre-commit hooks locally.
//
// Example usage:
//
// ```bash
// mage runprecommit
// ```
//
// **Returns:**
//
// error: An error if any issue occurs at any of the three stages
// of the process.
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

// RunTests runs all of the bats test for custom pre-commit hooks
//
// Example usage:
//
// ```bash
// mage runtests
// ```
//
// **Returns:**
//
// error: An error if any issue occurs while trying to
func RunTests() error {
	if _, err := sys.RunCommand("bash", filepath.Join(".hooks", "run-bats-tests.sh")); err != nil {
		return fmt.Errorf("failed to run unit tests: %v", err)
	}

	return nil
}

// Setup initializes and configures the homebrew-brewfile repo
// on the local system.
//
// Example usage:
//
// ```bash
// mage setup
// ```
//
// **Returns:**
//
// error: An error if any issue occurs while trying to
// set up the brewfile repo.
func Setup() error {
	// Make sure we are in the repo root
	repoRoot, err := git.RepoRoot()
	if err != nil {
		return err
	}

	// Get the current working directory
	// so that we can return to it
	cwd := sys.Gwd()

	if cwd != repoRoot {
		if err := sys.Cd(repoRoot); err != nil {
			return err
		}
	}

	var cdErr error
	defer func() {
		cdErr = sys.Cd(cwd)
	}()
	if cdErr != nil {
		return fmt.Errorf("failed to cd back to %s: %v", cwd, cdErr)
	}

	home, err := sys.GetHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	srcBrewFile := filepath.Join(repoRoot, "Brewfile")

	// Read the contents of the srcBrewFile into a byte slice
	srcBrewFileContents, err := os.ReadFile(srcBrewFile)
	if err != nil {
		return fmt.Errorf("failed to read brewfile contents: %v", err)
	}

	brewfileInstallPath := filepath.Join(home, ".brewfile", "Brewfile")

	// Copy contents of srcBrewFile to the .brewfile directory - this is necessary
	// for running brew bundle.
	if err := fileutils.Create(brewfileInstallPath, srcBrewFileContents, fileutils.CreateFile); err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	// // Copy Brewfile to the .brewfile directory - this is necessary
	// // for running brew bundle.
	// if err := sys.Cp(filepath.Join(repoRoot, "Brewfile"),
	// 	filepath.Join(home, ".brewfile", "Brewfile", "Brewfile")); err != nil {
	// 	return fmt.Errorf("failed to set brewfile repo: %v", err)
	// }

	return nil
}

// Update updates various components of
// the Brewfile on the local system.
//
// Example usage:
//
// ```bash
// mage update
// ```
//
// **Returns:**
//
// error: An error if any issue occurs while trying to
// update the Brewfile.
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

// Run runs the Brewfile on the local system.
//
// Example usage:
//
// ```bash
// mage run
// ```
//
// **Returns:**
//
// error: An error if any issue occurs while trying to
// run the Brewfile.
func Run() error {
	home, err := sys.GetHomeDir()
	if err != nil {
		return err
	}

	fmt.Println("Installing brew packages, please wait...")

	sys.Cd(filepath.Join(home, ".brewfile", "Brewfile"))
	if _, err := sys.RunCommand("brew", "bundle"); err != nil {
		return fmt.Errorf("failed to run brewfile: %v", err)
	}

	return nil
}

// GeneratePackageDocs creates documentation for the various packages
// in the project.
//
// Example usage:
//
// ```go
// mage generatepackagedocs
// ```
//
// **Returns:**
//
// error: An error if any issue occurs during documentation generation.
func GeneratePackageDocs() error {
	fs := afero.NewOsFs()

	repoRoot, err := git.RepoRoot()
	if err != nil {
		return fmt.Errorf("failed to get repo root: %v", err)
	}
	sys.Cd(repoRoot)

	repo := docs.Repo{
		Owner: "l50",
		Name:  "homebrew-brewfile",
	}

	templatePath := filepath.Join("magefiles", "tmpl", "README.md.tmpl")
	if err := docs.CreatePackageDocs(fs, repo, templatePath); err != nil {
		return fmt.Errorf("failed to create package docs: %v", err)
	}

	return nil
}
