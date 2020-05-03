package main

import (
	"fmt"
)

/*
	指针 - 类似c

	c语言中的指针有个概念上的坑点，go中避免了
	c中指针的重点：不要混淆 int *p = &num; 和 *p = &num; （前者等价于 int *p; p = &num;）
*/

func baseDemo() {
	fmt.Println("========== 指针基本使用 BEGIN ==========")

	var a int16 = 300
	var b int16 = 400
	var ptr *int16 = &a
	fmt.Println("a的值为：", a)
	fmt.Println("b的值为：", b)
	fmt.Printf("指针ptr的值为 [%v]\n", ptr)
	fmt.Println("----------- 指针修改后 -----------")
	*ptr = 100
	ptr = &b
	*ptr = 101
	fmt.Println("a现在的值为：", a)
	fmt.Println("b现在的值为：", b)
	fmt.Printf("指针ptr的值为 [%v]\n", ptr)

	fmt.Println("========== 指针基本使用 END ==========")
	fmt.Println()
}

func main() {
	baseDemo()
}
