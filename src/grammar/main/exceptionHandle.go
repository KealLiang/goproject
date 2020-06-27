package main

import (
	"errors"
	"fmt"
	"strings"
	"util/logs"
)

/*
	## go中的恐慌(panic)处理
	* 处理恐慌的一般方式 defer-recover 模板
	* defer关键字
		- 底层是runtime.deferproc，在函数返回前调用
		- return语句并非原子性的，可改写为如下三句（理解defer的关键）：
			1. 设置 返回值 = xxx
			1. 调用defer函数
			1. return
*/

func readConfigDemo() {
	logs.Begin("自定义错误场景演示")
	err := readConfig("haha.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("程序readConfigDemo()继续执行 (●'◡'●)")
}
func readConfig(name string) (err error) {
	if strings.HasSuffix(name, ".yml") {
		fmt.Println("配置读取成功！！")
		return nil
	} else {
		return errors.New("配置文件类型有误！！") //返回一个自定义的error类型错误
	}
}

func panicDemo() {
	logs.Begin("go中的异常(panic)处理")
	// 通过 defer-recover 处理错误
	someMethod()
	fmt.Println("异常方法后的语句 (●'◡'●)")
}
func someMethod() {
	//defer func() {
	//	recover()
	//}()
	defer func() {
		//err := recover() // recover()是一个内置的函数，可以捕获到异常
		if err := recover(); err != nil { // if前可以做一个语句
			fmt.Println("发生了异常！", err)
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println(res)
}

// ==================== 深入理解defer原理避免踩坑 ====================
//先来看看几个例子
//例1：
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//例2：
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//例3：
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	panicDemo()
	readConfigDemo()

	println(f1()) // 1
	println(f2()) // 5
	println(f3()) // 1
}
