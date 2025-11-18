package main

// go语言的基本数据结构类型
/*
Go语言中数据类型分为:基本数据类型和复合数据类型
（1）基本数据类型有:
 	整型、浮点型、布尔型、字符串
（2）复合数据类型有:
	数组、切片、结构体、函数、map、通道(channel)、接口等。
*/

import (
	"fmt"
	"unsafe"
)

func func4() {

	var num int = 10
	fmt.Printf("num=%v 类型: %T\n", num, num)

	// (-128 到  127)
	var num1 int8 = 127
	var num2 int32 = 11227
	fmt.Printf("num1=%v, 类型: %T, 字节: %v \n", num1, num1, unsafe.Sizeof(num1))
	fmt.Printf("字节: %v \n", unsafe.Sizeof(num2))

	var a uint8 = 200
	fmt.Println(a)

	//不同长度之间的转换, 高位向低位转的时候要注意溢出的问题
	var a1 int32 = 10000
	var a2 int8 = 127
	a3 := a1 + int32(a2)
	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	//数字字面量语法
	// 十进制
	var aNum int
	aNum = 11
	fmt.Printf("%v \n", aNum) // 原样输出
	fmt.Printf("%d \n", aNum) // 十进制
	fmt.Printf("%b \n", aNum) // 二进制
	fmt.Printf("%o \n", aNum) // 八进制
	fmt.Printf("%x \n", aNum) // 十六进制
	fmt.Printf("%T\n", aNum)
	fmt.Printf("%v\n", unsafe.Sizeof(aNum))

}
