package main

import "fmt"

func main() {
	//(1)值拷贝
	var x = 10
	var y = x
	x = 20
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	//(2)
	var a = 1 + 1
	fmt.Println("a:", a)
	b := x * y
	fmt.Println("b:", b)
}
