package main

/*
	## go中的内置函数
	* 执行顺序为 全局变量定义 -> init函数 -> main函数
	* 如果A引入了别的文件B，则B的上诉方法先按顺序执行，再到A的执行，可以用于初始化加载等等
*/

import (
	"fmt"
	"util/logs"
)

func buildinMakeDemo() {
	logs.Begin("内置的make演示（分配引用类型的指针）")
	fmt.Println(makeFbnSlice(25))
}

func makeFbnSlice(n uint) []uint64 {
	logs.Separate("make一个切片，存放斐波那契数列")
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < int(n); i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func buildinNewDemo() {
	logs.Begin("内置的new演示（分配值类型的指针）")
	num1 := 100
	fmt.Printf("num1的类型是%T, 值是%v, 地址是%v\n", num1, num1, &num1)

	num2 := new(int) // new了个指针
	fmt.Printf("num2的类型是%T, 值是%v, 地址是%v\n", num2, num2, &num2)
}

var Age1 int = 10 // 定义同时初始化，是1句，是OK的
//Age2 := 11 // 这里是错误的，因为这里等价于2句，而第2句赋值不能放在函数体外

var f1 = globalHaha()

func globalHaha() string {
	logs.Begin("这里是 globalHaha 函数")
	logs.End("这里是 globalHaha 函数")
	return "haha"
}

// init在main之前执行，可以做初始化工作
func init() {
	logs.Begin("这里是 init 函数")
	logs.End("这里是 init 函数")
}

func main() {
	logs.Begin("这里是 main 函数")
	logs.End("这里是 main 函数")
	buildinNewDemo()
	buildinMakeDemo()
}
