# Go programs are easy on memory

Go programs are fairly lightweight.
Each program includes a small amount of "extra" code that's included in the executable binary.
This extra code is called the Go Runtime.
One of the purposes of the Go runtime is to cleanup unused memory at runtime.

In other words, the Go compiler includes a small amount of extra logic in every Go program to make it easier for developers to write code that's memory efficient.

## Comparison

As a general rule Java programs use *more* memory than comparable Go programs because Go doesn't use an entire virtual machine to run its programs, just a small runtime.
The Go runtime is small enough that it is included directly in each program's compiled machine code.

As another general rule Rust and C++ programs use slightly less memory than Go programs because more control is given to the developer to optimize memory usage of the program. The Go runtime just handles it for us automatically.

```md
Q: Generally speaking, which language uses more memory?

* Java
* Go

A: Java

Q: What's one of the purposes of the Go runtime?

* To style Go code and make it easier to read
* To cleanup unused memory
* To cook fried chicken
* To compile Go code

A: To cleanup unused memory
```
