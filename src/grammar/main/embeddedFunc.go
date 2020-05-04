package main

/*
	## go中的内置函数
	* 执行顺序为 全局变量定义 -> init函数 -> main函数
	* 如果A引入了别的文件B，则B的上诉方法先按顺序执行，再到A的执行，可以用于初始化加载等等
*/

import "util/logs"

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
}
