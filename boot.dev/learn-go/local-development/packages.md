# Packages

Every Go program is made up of packages.

A package named "main" has an entrypoint at the `main()` function.
A `main` package is compiled into an executable program.

A package by any other name is a "library package".
Libraries have no entry point. Libraries simply export functionality that can be used by other packages.
