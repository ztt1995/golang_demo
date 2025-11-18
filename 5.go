package main

// 浮点型 float32、float64

import (
	"fmt"
	"unsafe"

	"github.com/shopspring/decimal"
)

func func5() {
	var a float32 = 3.12
	fmt.Printf("值:%f, 类型:%T, 字节:%v\n", a, a, unsafe.Sizeof(a))

	var b float64 = 3.12
	fmt.Printf("值:%f, 类型:%T, 字节:%v\n", b, b, unsafe.Sizeof(b))

	var c float64 = 3.14159265333
	fmt.Printf("值:%.4f, 值:%f, 类型:%T, 字节:%v\n", c, c, c, unsafe.Sizeof(c))

	// 64位的系统中Go语言浮点数默认是 float64
	d := 3.14159265333
	fmt.Printf("%T\n", d)

	// 科学计数法

	var f1 float32 = 3.14e-2
	fmt.Printf("%v--%f\n", f1, f1)

	//浮点数精度丢失的问题，引入第三方包
	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2)                                                //期望是4.4，结果打印出了4.399999999999999
	fmt.Println(decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))) //期望是4.4，结果打印出了4.399999999999999

	/* Go 语言中以 bool 类型进行声明布尔型数据，布尔型数据只有 true(真)和 false(假)两个
	值。
	注意:
	1. 布尔类型变量的默认值为 false。
	2. Go 语言中不允许将整型强制转换为布尔型.
	3. 布尔型无法参与数值运算，也无法与其他类型进行转换。 */

	var flag bool
	fmt.Printf("值:%v, 类型:%T, 字节:%v\n", flag, flag, unsafe.Sizeof(flag))

	if flag {
		fmt.Println("满足")
	} else {
		fmt.Println("不满足")
	}
	// string 类型的默认值为空
	var str string
	fmt.Printf("%v\n", str)

	// int 类型的默认值为0
	var num int
	fmt.Printf("%v\n", num)

	// float 类型的默认值为0
	var numf float64
	fmt.Printf("%v\n", numf)
}
