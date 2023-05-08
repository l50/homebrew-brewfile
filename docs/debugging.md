# Debugging

If you're encountering an error while running:

```bash
brew file update --verbose debug --appstore 0
```

Try the following in order:

```bash
# Reset taps
brew update-reset

# Update taps
brew update

# Cleanup old taps
brew cleanup

# Remove old taps that are not linked
brew autoremove

# Run brew doctor
brew doctor

# https://github.com/Homebrew/homebrew-cask/issues/140701
export HOMEBREW_NO_INSTALL_FROM_API=1
brew file update --verbose debug --appstore 0
```
