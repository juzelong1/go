package main

import "fmt"

const (
	name = 0
	gao = 1
)

func main() {
	fmt.Println(max(10, 12))
	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}


	fmt.Println(name)

	var a = test()
	var b = test2()
	var c = test3()
	fmt.Println(a, b, c)
	var e, d, f = "你好", 1, 3
	fmt.Println(e, d, f)

	/* 定义局部变量 */
	var x int = 10

	/* 使用 if 语句判断布尔表达式 */
	if x < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("x 小于 20\n" )
	}
	fmt.Printf("x 的值为 : %d\n", x)
}

func test() string {
	var a string = "你好"
	return a
}

func test2() int  {
	var b int = 1
	return b
}

func test3() bool  {
	var c bool = true
	return c
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

