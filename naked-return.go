package main

import "fmt"

func foo(a, b int) (x, y int) {
	x = a * 3
	y = b * 4
	return
}

func main() {
	x, y := foo(1, 2)
	fmt.Println(x, y)
}
