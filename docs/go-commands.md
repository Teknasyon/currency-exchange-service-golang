# Golang Commands

## go get

The `go get` command adds dependencies to current module and installs them. It adds the dependency to go.mod file if the module does not exist in the GOPATH. If the module already exists in the GOPATH, it will update the version in go.mod file.

Example:

```bash
# Install the latest version of a package
go get github.com/gorilla/mux
# Install a specific version of a package
go get github.com/gorilla/mux@version
# Install a specific commit of a package
go get github.com/gorilla/mux@commit
```

## go install

The `go install` command compiles and installs packages and dependencies.

Example:

```bash
# Install the current package
go install
# Install a specific package
go install github.com/gorilla/mux
```

## go run

The `go run` command compiles and runs Go program.

Example:

```bash
# Run the current package
go run .
# Run a specific package
go run github.com/gorilla/mux
```

## go build

The `go build` command compiles packages and dependencies.

Example:

```bash
# Build the current package
go build .
# Build for linux
GOOS=linux GOARCH=amd64 go build .
# Build for windows
GOOS=windows GOARCH=amd64 go build .
# Build for mac
GOOS=darwin GOARCH=amd64 go build .
```

## Go Modules

Go modules are a dependency management system built directly into the go command. They allow versioned packages to be installed and updated in a predictable way. Modules are the future of dependency management in Go.

### go mod init

The `go mod init` command initializes a new module, creating a go.mod file in the current directory and in effect, turning the current directory into the root of a new module.

### go mod tidy

The `go mod tidy` command adds any missing modules necessary to build the current module's packages and dependencies. It then removes unused modules that don't provide any relevant packages. In effect, `go mod tidy` cleans up the current module's dependencies.

### go mod vendor

The `go mod vendor` command makes a copy of all the dependencies of the current module into a subdirectory named vendor. This subdirectory is then used to build the current module, as if the source code of the dependencies had been copied directly into the module.

### go mod download

The `go mod download` command downloads the modules needed to build the current module, as well as the modules needed to build the tests for the current module. It does not build the current module or its tests.

### go mod verify

The `go mod verify` command verifies that the dependencies of the current module have not been modified since go.mod was created or updated.

### go mod graph

The `go mod graph` command prints the dependency graph of the current module. The graph shows the dependencies of the current module, with the required modules listed first, followed by the dependencies of those modules, and so on. Each module is listed only once, even if multiple modules depend on it.

### go mod why

The `go mod why` command prints the module requirement graph that explains why each module is needed in the current module's dependency graph. The graph shows the dependencies of the current module, with the required modules listed first, followed by the dependencies of those modules, and so on. Each module is listed only once, even if multiple modules depend on it.
