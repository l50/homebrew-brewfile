# homebrew-brewfile

[![Test Brewfile](https://github.com/l50/homebrew-brewfile/actions/workflows/tests.yaml/badge.svg)](https://github.com/l50/homebrew-brewfile/actions/workflows/tests.yaml)
[![License](https://img.shields.io/github/license/l50/homebrew-brewfile?label=License&style=flat&color=blue&logo=github)](https://github.com/l50/homebrew-brewfile/blob/main/LICENSE)

This repo is used to manage the packages installed on my Macs.

## Dependencies

- [Install asdf](https://asdf-vm.com/):

  ```bash
  git clone https://github.com/asdf-vm/asdf.git ~/.asdf
  ```

- Install and use asdf plugins to manage go, python, and ruby for this project:

  ```bash
  source .asdf
  ```

  Alternatively, you can pick and choose which plugins to install:

  ```bash
  # Employ asdf for this project's python:
  source .asdf python
  ```

- [Install pre-commit](https://pre-commit.com/):

  ```bash
  python3 -m pip install --upgrade pip
  python3 -m pip install pre-commit
  ```

- [Install Mage](https://magefile.org/):

  ```bash
  go install github.com/magefile/mage@latest
  ```

- Install homebrew:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

- Install [homebrew-file](https://github.com/rcmdnk/homebrew-file):

```bash
brew tap rcmdnk/file
brew install brew-file
```

- Install and run pre-commit hooks:

```bash
mage runprecommit
```

- Setup the Brewfile found in this repo on the local system:

```bash
mage setup
```

---

## Usage

- Update everything:

```go
mage update
```

---

## Debugging

If you're having trouble, please refer to the [debugging doc](docs/debugging.md).

---

### Test actions locally

```bash
act -P macos-latest=-self-hosted
```

---

## Useful Resources

- <https://thoughtbot.com/blog/brewfile-a-gemfile-but-for-homebrew>
- <https://coderwall.com/p/afmnbq/homebrew-s-new-feature-brewfiles>
- <https://medium.com/@satorusasozaki/automate-mac-os-x-configuration-by-using-brewfile-58a78ce5cc53>
- <https://github.com/timbru31/homebrew-brewfile/blob/master/Brewfile>
- <https://homebrew-file.readthedocs.io/en/latest/installation.html>
- <https://gist.github.com/ChristopherA/a579274536aab36ea9966f301ff14f3f>
