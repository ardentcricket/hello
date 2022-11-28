package main

import "fmt"

func main() {
	a, b := 0, 0
	value := Add(a)
	valuePtr := AddPtr(&b)
	fmt.Println("a: ", a, "b: ", " value for A: ", value, b, " value for B ptr:", valuePtr)
}

func Add(x int) int {
	 x++
	 return x
}

func AddPtr(x *int) int {
	*x++
	return *x
}
