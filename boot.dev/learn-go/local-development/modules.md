# Modules

Go programs are organized into packages.
A package is a directory of Go code that's all compiled together.
Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package (directory).

A repository contains one or more modules. A module is a collection of Go packages that are released together.

## A Go repository typically contains only one module, located at the root of the repository

A file name `go.mod` at the root of a project declares the module. It contains:

* The module path
* The version of the Go language your project requires
* Optionally, any external package dependencies your project has

The module path is just the import path prefix for all packages within the module.

Here's an example of a `go.mod` file:

```md
module github.com/bootdovdev/exampleproject

go 1.20

require github.com/google/examplepackage v1.3.0
```

Each module's path not only serves as an import path prefix for the packages within
but also indicates where the go command should look to download it.
For example, to download the module `golang.org/x/tools`, the go command would consult
the repository located at `https://golang.org/x/tools`.

```md
A "import path" is a string used to import a package.
A package's import path is its module path joined with its subdirectory within the module.
For example, the module `github.com/google/go-cmp` contains a package in the directory `cmp/`.
That packages's import path is `github.com/google/go-cmp/cmp`.
Packages in the standard library do not have a module path prefix.
```

```md
Q: What is a Go module?

* An executable main package
* A collection of packages that are released together
* A file of Go code
* A library package

A: A collection of packages that are released together

Q: Do packages in the standard library have a module path prefix?
A: No

Q: What is an import path?
* An HTTP connection
* A module path + package subdirectory
* A RESTful server
```
