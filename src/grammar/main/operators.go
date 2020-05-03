package main

import (
	"fmt"
	"util/logs"
)

/*
	##go中的运算符
		* 算数运算符
		* 赋值运算符
		* 比较运算符
		* 逻辑运算符
		* 位运算符
		* 其他运算符

	##go中的顺序分支循环
		* if
		* switch
		* for
		* 流程控制（break和continue和goto都可以跳到标签定义的地方(label:)）
*/

func labelDemo() {
	logs.Begin("流程控制 - label演示")
here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Printf("i = %v; j = %v;\n", i, j)
		}
	}
}

func forDemo() {
	logs.Begin("循环 - for演示")
	for i := 0; i < 3; i++ {
		fmt.Println("你好，世界", i)
	}
	logs.Separate("死循环 for-break")
	for {
		fmt.Println("哈哈")
		break
	}
	logs.Separate("遍历字符串 for-range")
	var str string = "hello world!你好" // for-range语法可以按照字符（而非字节）取，因此能正确取出中文
	for index, val := range str {
		fmt.Printf("[%v]%c ", index, val)
	}
}

func switchDemo(dayOfWeek uint) {
	logs.Begin("分支 - switch演示")
	// go中的 switch-case 不用加break
	// switch 后面可以不写表达式，直接写 case 的判断式，类似 if-else 的使用方式
	// switch 语句还有一种用法 type-switch 即类型判断
	switch dayOfWeek % 7 {
	case 1:
		fmt.Println("今天星期一 月曜日")
	case 2:
		fmt.Println("今天星期二 火曜日")
	case 3:
		fmt.Println("今天星期三 水曜日")
		fallthrough // case 穿透，默认只穿透一层
	case 4:
		fmt.Println("今天星期四 木曜日")
	case 5:
		fmt.Println("今天星期五 金曜日")
	case 6:
		fmt.Println("今天星期六 土曜日")
	case 7, 0:
		fmt.Println("今天星期日 日曜日") //case 后面可以带多个表达式，逗号分隔
	default:
		fmt.Println("输入的dayOfWeek有误！！")
	}
}

func ifDemo(age uint) {
	logs.Begin("分支 - if演示")
	if age >= 18 && age < 60 {
		fmt.Println("成年人")
	} else if age < 18 && age >= 0 {
		fmt.Println("小孩")
	} else if age >= 60 {
		fmt.Println("老人")
	} else {
		fmt.Println("年龄不合法！")
	}
}

func baseDemo() {
	logs.Begin("基本数据运算演示")
	var n1 float32 = 10 / 4
	var n2 float32 = 10.0 / 4
	fmt.Println("运算得到结果，不管接收方什么类型 n1 =", n1)
	fmt.Println("想要小数运算在计算时指明10.0即可 n2 =", n2)
	logs.Separate("取模运算的公式：a % b = a - a/b*b")
	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("10 % -3 =", 10%-3)
	fmt.Println("-10 % 3 =", -10%3)
	fmt.Println("-10 % -3 =", -10%-3)
	logs.Separate("go中的 ++ -- 运算符只能作为单独的语句使用，aa := i++ 这样是不允许的")
	logs.End("基本数据运算演示")
}

func main() {
	baseDemo()
	ifDemo(22)
	switchDemo(10)
	forDemo()
	labelDemo()
}
