# Package Naming

## Naming Convention

By convention, a package's name is the same as the last element of its import path.
For instance, the `math/rand` package comprises files that begin with:
```go
package rand
```

That said, package names aren't required to match their import path.
For example, I could write a new package with the path `github.com/mailio/rand` and name the package `random`:
```go
package random
```

While the above is possible, it is discouraged for the sake of consistency.

## One Package / Directory

A directory of Go code can have **at most** one package.
All `.go` files in a single directory must all belong to the same package.
If they don't an error will be thrown by the compiler.
This is true for main and library packages alike.

```md
Q: What would be the conventional package name of a package with the path github.com/wagslane/parser ?

* parser
* wagslane
* github.com
* go-parser

A: parser, the last element of the import path
```

```md
Q: Given the import path of path/to/rand, which of these is a valid package name?

* random
* rand
* spam
* path
* Any of these

A: Any of these, but comform to the conventions, it should be rand
```
