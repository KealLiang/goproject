package main

import (
	"errors"
	"fmt"
	"strings"
	"util/logs"
)

/*
	go中的错误处理
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

func main() {
	panicDemo()
	readConfigDemo()
}
