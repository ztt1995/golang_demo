package main

//fmt 包

import "fmt"

func main() {
	// 换行输出
	//  fmt.Println("Hello, World!")
	// //  不换行输出
	//  fmt.Print("Hello, World!")
	//  fmt.Printf("Hello, World!%v", "\n")

	//  var a = "aaa"
	//  fmt.Printf("a is %v", a)

	var a int = 10
	var b int = 3
	var c int = 5

	fmt.Printf("a + b = %v\n", a+b)
	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	//类型推导方式定义变量
	d := 20
	f := 30
	fmt.Printf("d=%v\n", d)
	fmt.Printf("f的类型是%T\n", f)

}
