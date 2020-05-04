package main

/*
	##go中的函数
	* 参数传递机制
		- 值传递（值拷贝）：基本数据类型、数组、结构体；希望修改的话可以传递指针 &变量名
		- 引用传递：指针、slice切片、map、chan管道、interface等

	* go不支持传统的重载
		- 通过空接口实现

	* go中函数也是一个变量（函数式编程）

	##go内存模型
	* 逻辑上：栈、堆、代码区
*/

import (
	"fmt"
	"strings"
	"util/logs"
)

// defer最佳实践场景
func deferScene() {
	//file = openfile("文件名")
	//defer file.close() //备忘关闭
	// 其他代码
}
func deferDemo() {
	logs.Begin("go 中的defer延迟栈演示")
	deferAdd := func(n1 int, n2 int) int {
		defer fmt.Println("deferAdd() 第1个入栈 n1 =", n1) //可以认为是压入了一个defer栈，先进后出
		defer fmt.Println("deferAdd() 第2个入栈 n2 =", n2) //值也会被拷贝入栈，如这里n1和n2
		// 验证值是 拷贝 入栈的
		n1++
		n2++
		res := n1 + n2
		fmt.Println("deferAdd() res =", res)
		logs.Separate("函数执行完成，在最后return前会依次将defer栈中的内容弹出，执行之")
		return res
	}

	out := deferAdd(10, 20)
	fmt.Println("deferDemo() out =", out)
}

// 闭包实践：判断传入的文件名是否有指定的后缀，若没有则加上
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 返回的匿名函数，持有外部变量n的引用，因此这个匿名函数和n形成一个整体，就是闭包
// 函数和她使用的n构成闭包，就如同一个 类，其中有域变量和操作它的方法
func addClosure() func(int) int {
	var n int = 0 //n的生命被延长了
	return func(x int) int {
		n += x
		return n
	}
}
func closureDemo() {
	logs.Begin("go 中的闭包演示")
	add := addClosure()
	fmt.Println("闭包调用1次：", add(5)) //5
	fmt.Println("闭包调用2次：", add(5)) //10
	fmt.Println("闭包调用3次：", add(5)) //15
	logs.Separate("判断传入的文件名是否有指定的后缀，若没有则加上")
	make := makeSuffix(".jpg") // 使用闭包的话，suffix只用定义一次，且避免使用全局变量
	fmt.Printf("文件[%v]经过makeSuffix()处理后：%v\n", "haha.jpg", make("haha.jpg"))
	fmt.Printf("文件[%v]经过makeSuffix()处理后：%v\n", "hehe123", make("hehe123"))
}

var (
	// minus是全局匿名函数
	minus = func(n1 int, n2 int) int {
		return n1 - n2
	}
)

func anonymousFunc() {
	logs.Begin("go 中的匿名函数")
	n1 := 30
	n2 := 40

	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(n1, n2)
	fmt.Println("通过匿名函数计算 n1 + n2 =", res)
	fmt.Println("定义全局匿名函数赋值给minus，并计算 n1 - n2 =", minus(n1, n2))
}

func variableParameters() {
	logs.Begin("go 中的可变参数示例")
	fmt.Println("求和 2,4,6 权重10：", summing(10, 2, 4, 6))
}
func summing(weight int, args ...int) (sum int) {
	// 这里的args是个slice
	for _, val := range args {
		sum += val
	}
	sum *= weight
	return
}

func functionPrograming() {
	logs.Begin("go 中的函数式编程")
	f := add
	fmt.Printf("f 的数据类型是 %T, add 的数据类型是 %T\n", f, add)
	call(15, 25, f)
}
func call(n1 int, n2 int, method myFuncType) {
	logs.Separate("方法作为参数传递")
	fmt.Printf("method(%v, %v) 的调用结果是：%v\n", n1, n2, method(n1, n2))
}

type myFuncType func(int, int) int // 用自定义类型 简化编写
func add(n1 int, n2 int) int {
	return n1 + n2
}

func testMaster() {
	logs.Begin("函数调用的值传递演示")
	var n1 int = 99
	var n2 int = 49
	fmt.Printf("调用testSlave(n1, &n2) 前testMaster() 中的 n1 = %v, n2 = %v\n", n1, n2)
	testSlave(n1, &n2)
	fmt.Printf("调用testSlave(n1, &n2) 后testMaster() 中的 n1 = %v, n2 = %v\n", n1, n2)
}
func testSlave(n1 int, ptr *int) {
	n1 += 1
	*ptr += 1
	fmt.Printf("testSlave() 中的 n1 = %v, ptr[%v] = %v\n", n1, ptr, *ptr)
}

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
	fmt.Printf("[%s]的年龄是%v，司龄%v，年薪%v\n", name, age, compAge, float64(compAge+12)*salary)
	logs.End("获取键盘输入")
}

// *函数返回多个值
// 示例：返回两个数的和与差
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// go支持对返回值命名，返回值定义、顺序一次搞定
func getSumAndSubNamedResult(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func showResult() {
	logs.Begin("返回两个数的和与差")
	i1, i2 := getSumAndSub(10, 100)
	fmt.Println(i1)
	fmt.Println(i2)
	sum, sub := getSumAndSubNamedResult(10, 1)
	fmt.Println("10 与 1 的和与差是：", sum, sub)
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
	//scanDemo()
	testMaster()
	functionPrograming()
	variableParameters()
	anonymousFunc()
	closureDemo()
	deferDemo()
}
