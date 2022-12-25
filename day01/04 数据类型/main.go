package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 1 << 7
	fmt.Println("x:", x)
	fmt.Println(reflect.TypeOf(x))
	y := 2e3
	fmt.Println(y, reflect.TypeOf(y))
	//fmt.Println(1 < 2)
}
