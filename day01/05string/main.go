package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "那天"
	s := "natian"
	//fmt.Println("s:", s)
	//fmt.Println(string(s[0]))
	//fmt.Println(len(s))
	//for index, value := range s {
	//	fmt.Println(index, string(value))
	//}

	//切片
	//fmt.Println(s[:])
	//fmt.Println(s[1:3])
	//fmt.Println(s[:3])
	//fmt.Println(s[1:])

	toUpperS := strings.ToUpper(s)
	fmt.Println(toUpperS)
	toLowerS := strings.ToLower(toUpperS)
	fmt.Println(toLowerS)

	fmt.Println(strings.HasPrefix(s, "na"))
	fmt.Println(strings.HasPrefix(s, "n1a"))

	fmt.Println(strings.HasSuffix(s, "an"))
	fmt.Println(strings.HasSuffix(s, "an1"))

	fmt.Println(strings.Contains(s, "tian"))
	fmt.Println(strings.Contains(s, "jiu"))

	fmt.Println(strings.Index(s, "tian"))
	fmt.Println(strings.Index(s, "tian1"))

	s1 := "    jjt    a"
	//fmt.Println(strings.Trim(s1, " "))
	fmt.Println(strings.TrimSpace(s1))

	fmt.Println(strings.Replace(s, "tian", "666", 1))

	date := "1998.31.2"
	arr := strings.Split(date, ".")
	fmt.Println(arr)
	date1 := strings.Join(arr, "/")
	fmt.Println(date1)
}
