# homebrew-brewfile

[![Test Brewfile](https://github.com/l50/homebrew-brewfile/actions/workflows/test.yml/badge.svg)](https://github.com/l50/homebrew-brewfile/actions/workflows/test.yml)
[![License](https://img.shields.io/github/license/l50/homebrew-brewfile?label=license&style=flat&color=blue&logo=github)](https://github.com/l50/homebrew-brewfile/blob/master/LICENSE)

Used to install packages that I like to have on my Macs.

## Setup

Install homebrew:

```bash
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Install dependencies from Brewfile:

```bash
brew bundle

# If you prefer verbose output:
brew bundle -v
```

Install pre-commit hooks:

```bash
pre-commit install
```

## Useful Resources

- <https://thoughtbot.com/blog/brewfile-a-gemfile-but-for-homebrew>
- <https://coderwall.com/p/afmnbq/homebrew-s-new-feature-brewfiles>
- <https://medium.com/@satorusasozaki/automate-mac-os-x-configuration-by-using-brewfile-58a78ce5cc53>
- <https://github.com/timbru31/homebrew-brewfile/blob/master/Brewfile>
- <https://homebrew-file.readthedocs.io/en/latest/installation.html>
- <https://gist.github.com/ChristopherA/a579274536aab36ea9966f301ff14f3f>
