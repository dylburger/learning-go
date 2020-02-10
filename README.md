# Learning Go

## Questions

- Why does the binary I get from running `go build` take so long to run the first time? My optional for loop took forever.

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

### Variables

`var` declares a list of variables, and the type declaration follows:

```go
var x, y, z int
```

You can include initializers in the variable declaration. If you initialize, you can omit the type since it's inferred:

```go
var x, y int = 1, 2
```

is equivalent to

```go
var x, y = 1, 2
```

`:=` is the "short assignment" syntax. You can only use this within a function. Outside of a function, `var`, `func` and other keywords must be used.

Variable declarations can also be factored, like import statements:

```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

### Zero values

Variables declared without a specific initial value are given a "zero value":

- `0` for numeric types
- `false` for `bool`
- `""` for strings

### Type conversion

`T(v)` converts the value `v` into the type `T`.

### Type inference

When a type is absent from a declaration and the right hand side contains an untyped number, the variable is typed according to the precision of the number:

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants

Constants are declared using the `const` keyword:

```go
const Pi = 3.14
```

Constants cannot be declared using short assigment (`:=`) syntax.

### For loops

Go only has one looping construct: the `for` loop.

Basic for loops have three, semi-colon delimited statements:

- The init statement: executed before the first iteration.
- The condition expression: evaluated _before_ every iteration.
- The post statement: executed at the end of every iteration.

The loop will stop iterating once the boolean condition expression evaluates to false.

The init and post statements are optional.

### Infinite loops

An infinite loop is very compact:

```go
for {
}
```
