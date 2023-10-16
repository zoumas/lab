# Compiled vs Interpreted

Compiled programs can be run without access to the original source code, and without access to a compiler.

This is different from interpreted languages like Python and JavaScript.
With Python and JavaScript the code is interpreted at runtime by a separate program known as the "interpreter".
Distributing code for users to run can be a pain because they need to have an interpreter installed,
and they need access to the original source code.

## Examples of Compiled Languages

* Go
* C
* C++
* Rust
* Haskell
* Zig

## Examples of Interpreted Languages

* JavaScript
* Python
* PHP
* Ruby
* Perl

**Deploying compiled programs is generally simpler**

```md
Q: Do users of compiled programs need access to source code?
A: No

Q: Which language is interpreted?
* Go
* C++
* Python
* Rust

A: Python

Q: Why is it more simple to deploy a compiled server program?
* Compiled code is more memory efficient
* Because docker exists
* There are no runtime dependencies
* Compiled code is faster

A: There are no runtime dependencies
```
