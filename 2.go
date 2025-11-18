package main

// 变量的定义和使用

import "fmt"

var globalV = "全局变量"

func getUserInfo() (string, int) {
	name, age := "lina", 10
	return name, age
}

func func2() {

	fmt.Print(globalV)

	/*
		1、var 声明变量
			 var  变量名称 类型
	*/
	var a string
	fmt.Println(a) // 变量声明后没有初始化的话，值为空

	/*
		2、go语言变量的定义和初始化
			 （1）、第一种初始化方式
			 			var 变量名称 类型
						变量名称 = 值
				(2）、第二种初始化方式
			 			var 变量名称 类型 = 值
				(3）、第三种初始化方式
			 			var 变量名称 = 值)
	*/
	var b string
	b = "hello"
	var c string = "world"
	var d = "!"
	fmt.Println(b, ",", c, d)

	/*
		3、一次定义多个变量
			(1)、var 变量名称1,变量名称2 类型  //类型都一样的
			(2）、var (
						变量名称1 类型1
						变量名称2 类型2
						变量名称3 类型3
					)  //类型不一样的

					var (
						变量名称1 类型1 = 值1
						变量名称2 类型2 = 值2
						变量名称3 类型3 = 值3
					)  //类型不一样的
	*/

	var e, f, g int
	e = 1
	f = 2
	g = 3
	fmt.Println(e, f, g)

	var (
		name string = "张三"
		age  int    = 10
		sex  string = "女"
	)
	fmt.Println(name, age, sex)

	var (
		name1 = "李四"
		age1  = 20
		sex1  = "男"
	)
	fmt.Println(name1, age1, sex1)

	/*
		4、短变量声明法 在函数内部，可以使用更简略的 := 方式声明并初始化变量，
			 注意：短变量只能用于声明局部变量，不能用于全局变量的声明
	*/
	username := "tingszhang"
	fmt.Println(username)
	fmt.Printf("%T\n", username)

	// 使用短变量一次声明多个变量，并初始化变量
	aa, bb, cc := 1, 2, "abc"
	fmt.Println(aa, bb, cc)
	fmt.Printf("aa的类型: %T,bb的类型: %T,cc的类型: %T\n", aa, bb, cc)

	/*
		匿名变量，在使用多重赋值的时候，如果想要忽略某个值，可以使用匿名变量（anonymous variable）,
		匿名变量用一个下划线表示: _
		匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	*/
	// var fName, fAge = getUserInfo()
	fName, _ := getUserInfo()
	fmt.Println(fName)

	var cName int
	cName = 123
	fmt.Println(cName)
}
