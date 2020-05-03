package main

import (
	"fmt"
	"util/logs"
)

// 获取键盘输入
func scanDemo() {
	logs.Begin("获取键盘输入")
	var name string
	var age, compAge uint
	var salary float64
	fmt.Println("请输入名字：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄和司龄和薪水 空格分隔：")
	fmt.Scanf("%d %d %f", &age, &compAge, &salary)
	fmt.Printf("[%s]的年龄是%v，司龄%v，年薪%v", name, age, compAge, float64(compAge+12)*salary)
	logs.End("获取键盘输入")
}

// *函数返回多个值
// 示例：返回两个数的和与差
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func showResult() {
	logs.Begin("返回两个数的和与差")
	i1, i2 := getSumAndSub(10, 100)
	fmt.Println(i1)
	fmt.Println(i2)
}

// 交换变量，不用中间变量
func swapVar() {
	logs.Begin("交换变量，不用中间变量")
	a := 1
	b := 100
	fmt.Printf("BEFORE: a = %v, b = %v\n", a, b)
	a, b = b, a
	fmt.Printf("AFTER: a = %v, b = %v\n", a, b)
}

func main() {
	showResult()
	swapVar()
	scanDemo()
}
