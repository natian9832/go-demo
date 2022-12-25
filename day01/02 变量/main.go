package main

import "fmt"

func main() {
	//(1)先声明再赋值
	var age int
	fmt.Println("age:", age)
	age = 25
	fmt.Println("age:", age)

	//(2)声明并赋值
	//var name string = "natian"
	//可以替换为
	var name = "那天"
	fmt.Println("name:", name)

	//(3)一行声明多个变量
	//var x, y int
	var (
		a int
		b string
		c bool
	)
	fmt.Println(a, b, c)

	//(4)声明并赋值的简介写法,不能使用于全局变量
	addr := "北京"
	fmt.Println("addr:", addr)

	//(5)一行声明赋值多个变量
	//var x, y, z = "那天", 25, "北京"
	x, y, z := "那天", 25, true
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
