# ansible-collection-workstation/magefiles

`magefiles` provides utilities that would normally be managed
and executed with a `Makefile`. Instead of being written in the make language,
magefiles are crafted in Go and leverage the [Mage](https://magefile.org/) library.

---

## Table of contents

- [Functions](#functions)
- [Contributing](#contributing)
- [License](#license)

---

## Functions

### GeneratePackageDocs()

```go
GeneratePackageDocs() error
```

GeneratePackageDocs creates documentation for the various packages
in the project.

Example usage:

```go
mage generatepackagedocs
```

**Returns:**

error: An error if any issue occurs during documentation generation.

---

### InstallDeps()

```go
InstallDeps() error
```

InstallDeps installs the Go dependencies necessary for developing
on the project.

Example usage:

```bash
mage installdeps
```

**Returns:**

error: An error if any issue occurs while trying to
install the dependencies.

---

### Run()

```go
Run() error
```

Run runs the Brewfile on the local system.

Example usage:

```bash
mage run
```

**Returns:**

error: An error if any issue occurs while trying to
run the Brewfile.

---

### RunPreCommit()

```go
RunPreCommit() error
```

RunPreCommit updates, clears, and executes all pre-commit hooks
locally. The function follows a three-step process:

First, it updates the pre-commit hooks.
Next, it clears the pre-commit cache to ensure a clean environment.
Lastly, it executes all pre-commit hooks locally.

Example usage:

```bash
mage runprecommit
```

**Returns:**

error: An error if any issue occurs at any of the three stages
of the process.

---

### Setup()

```go
Setup() error
```

Setup initializes and configures the homebrew-brewfile repo
on the local system.

Example usage:

```bash
mage setup
```

**Returns:**

error: An error if any issue occurs while trying to
set up the brewfile repo.

---

### Update()

```go
Update() error
```

Update updates various components of
the Brewfile on the local system.

Example usage:

```bash
mage update
```

**Returns:**

error: An error if any issue occurs while trying to
update the Brewfile.

---

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what
you would like to change.

---

## License

This project is licensed under the MIT
License - see the [LICENSE](../LICENSE)
file for details.
