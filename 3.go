package main

// 常量、命名规则、代码风格

import "fmt"

func func3() {

	const pi = 3.1415976
	fmt.Println(pi)
	fmt.Printf("%T\n", pi)

	const A = "A"
	fmt.Println(A)

	const (
		a = "a"
		b = "b"
		c = "c"
	)
	fmt.Println(a, b, c)

	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	const n4, n5, n6 = 10, 11, 12
	fmt.Println(n4, n5, n6)

	/*
		iota是golang语言的常量计数器，只能在常量的表达式中使用。

		iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量 声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	*/
	const a0 = iota
	const (
		a1 = iota // 每次const出现时，都会让iota初始化为0（自增长）
		_         // 1
		a3        // 2
		a4        //3
	)
	fmt.Println(a0, a1, a3, a4) // 0 1 2 3

	const (
		b1 = iota // 每次const出现时，都会让iota初始化为0（自增长）
		b2 = 100  // 1
		b3 = iota // 2
		b4        // 3
	)
	fmt.Println(b1, b2, b3, b4) // 0 100 2 3

	const (
		c1, c2 = iota + 1, iota + 2 // 1 2
		c3, c4                      // 2 3
		c5, c6                      // 3 4
	)
	fmt.Println(c1, c2, c3, c4, c5, c6) // 1 2 2 3 3 4

	/*
		 (1)、代码每一行结束后不用写分号( ;)
		（2）、强制的代码风格
				左括号必须紧接着语句不换行，这个特性刚开始会使开发者不习惯，但随着对 Go 语言的 不断熟悉，就会发现风格统一让大家在阅读代码时把注意力集中到了解决问题上，而不是代码风格上
	*/
}
