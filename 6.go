package main

// 字符串

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "hello"
	var str2 = "world"
	str3 := "go语言"
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

	var str4 = `hello
	hello
	hello
				tingszhang
	`
	fmt.Println(str4)

	var str5 = "world"
	fmt.Println(len(str5))

	//拼接

	str6 := str5 + str1
	str7 := fmt.Sprintf("%v %v", str1, str2)

	fmt.Println(str6)
	fmt.Println(str7)

	// 分割 strings.Split
	var str8 = "123-456-789"
	arr := strings.Split(str8, "-") //切片
	fmt.Println(str8, arr)

	strNew := strings.Join(arr, "%")
	fmt.Println(strNew)

	arr1 := []string{"123", "java", "go"}
	fmt.Println(arr1)

	str8 = "123-456-45789"
	str2 = "45"

	flag := strings.Contains(str8, str2)
	fmt.Println(flag)

	flag = strings.HasPrefix(str8, str2)
	fmt.Println(flag)

	flag = strings.HasSuffix(str8, str2)
	fmt.Println(flag)

	num := strings.Index(str8, str2)
	fmt.Println(num)

	num = strings.LastIndex(str8, str2)
	fmt.Println(num)

}
