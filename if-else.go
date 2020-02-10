package main

import "fmt"

const Pi = 3.14

func conditionalAdd(x, y, n int) int {
	if v := x + y; v < n {
		return v
	}
	return n
}

func main() {
	// Should return 3
	x := conditionalAdd(1, 2, 5)
	fmt.Println(x)

	// Should return 5
	y := conditionalAdd(10, 20, 5)
	fmt.Println(y)
}
