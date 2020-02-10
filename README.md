# Learning Go

## Resources

- [golang.org docs](https://golang.org/doc/#learning)

## Tour

### 101

Quoting heavily from the tour for my own notes.

Every Go program is made up of packages.

Programs begin running in the package `main`.

We import other packages using the following syntax:

```go
import (
    "fmt"
    "math/rand"
)
```

Parens surround imports, newline separated?

By convention, the package name is the same as the last element of the import path. For example, the package `math/rand` comprises files that include the statement `package rand`.

Altogether:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
```

- print -> `fmt.Println`
- must import `fmt`

The parenthesized style is called a "factored" import statement.

[`goimports`](https://godoc.org/golang.org/x/tools/cmd/goimports) can also format your import lines, and even remove missing imports or remove unreferenced ones! VSCode runs `goimports` for you, which is why my unreferenced imports got removed.

### Exports

A name is exported if its name begins with a capital letter.

When you use an imported package, you can only reference its exported names.

### Functions

A function an take zero or more arguments.

This function takes two args of type `int`:

```go
func add(x int, y int) int {
    return x + y
}
```

The type declaration happens on the right (`x int`), different than how C does it (`int x;`).

Many functional languages declare types like so:

```
x: int
```

Go just drops the `:` for brevity.

When two or more consecutive named function parameters share the same type, you can elide the argument names and just include the type at the end. For example these two are equivalent:

```go
func add(x int, y int) int {
    return x + y
}

func add(x, y int) int {
    return x + y
}
```

A function can return any number of results, comma-separated (`return-multiple-args.go`).

Go's return values may be named. Declare variables within the function and then issue a "naked return" at the end of the function.
