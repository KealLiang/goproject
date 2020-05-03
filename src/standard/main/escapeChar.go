package main

/*
	* go中的转义字符
	* go中的注释：行注释、块注释
		- 官方的包里基本都是行注释，参考
*/

import "fmt"

func main() {
	fmt.Println("go中的转义字符：")
	fmt.Println("制表符\t111")
	fmt.Println("换行符\n222")
	fmt.Println("斜杠\\333")
	fmt.Println("一个回车\r444") // 从当前行的最前面输出，覆盖之前的内容

	fmt.Println("===============================")
	fmt.Println("多行字符串aodfnofdngoisdjfgoijadfgjdjfgoi" +
		"lkjalsdjflajogdfnafhngonaomvonraovnoraivgalkjgsdfsdf" +
		"ljaonojfnoajdfjhpoqangopanmsjfsaljdfojasojwowanddfhf")
}
