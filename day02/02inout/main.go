package main

import "fmt"

func output() {
	//输出函数
	name, age := "那天", 25
	//fmt.Println(name, reflect.TypeOf(name))
	//fmt.Println(age, reflect.TypeOf(age))

	//fmt.Printf("name:%s,age:%d\n", name, age)
	//fmt.Printf("name:%T,age:%T\n", name, age)
	//fmt.Printf("name:%v,age:%v\n", name, age)

	//Sprintf
	message := fmt.Sprintf("name:%s,age:%d", name, age)
	fmt.Printf("message:{%s}\n", message)
}

func input() {
	//(1)scan
	var name string
	n, _ := fmt.Scan(&name)
	fmt.Println("n:", n)
	fmt.Println("name:", name)
}
func main() {

	//output()
	input()

}
