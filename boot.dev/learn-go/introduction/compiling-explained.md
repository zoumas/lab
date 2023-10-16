# Compiling Explained

Computers don't know how to do anything unless we as programmers tell them what to do.

Unfortunately computers don't understand human language.
In fact, they don't even understand uncompiled computer programs.

For example, the code:
```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world!")
}
```
means *nothing* to a computer.

## Computers need machine code 

A computer's CPU only understands its own *instruction* set, which we call "machine code".

Instructions are basic math operations like addition, subtraction, multiplication, and the ability to save data temporarily.

For example, an ARM processor uses the ADD instruction when supplied with the number `0100` in binary.

## Go, C, Rust, etc.

Go, C and Rust are all languages where the code is first converted to machine code by the compiler before it's executed.

```md
Q: Do computer processors understand English instructions like 'open the browser'?

A: No
```
