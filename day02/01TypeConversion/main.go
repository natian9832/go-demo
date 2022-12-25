package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//(1)整型之间转化
	//var x int8 = 10
	//var y int16 = 20
	//fmt.Println(x + int8(y))

	//(2)字符串与整型转换
	//age := "32"
	//ageInt, err := strconv.Atoi(age)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(ageInt, reflect.TypeOf(ageInt))

	//price := 100
	//priceString := strconv.Itoa(price)
	//fmt.Println(priceString, reflect.TypeOf(priceString))

	//(3)parse
	//str-->int
	result1, _ := strconv.ParseInt("28", 10, 8)
	fmt.Println(result1, reflect.TypeOf(result1))
	//str-->float
	result2, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(result2, reflect.TypeOf(result2))
	//str-->bool
	result3, err := strconv.ParseBool("1")
	fmt.Println("err:", err)
	fmt.Println(result3, reflect.TypeOf(result3))

}
