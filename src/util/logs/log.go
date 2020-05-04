package logs

import "fmt"

var Var1 string = "共有变量"
var var2 string = "私有变量"

func Begin(str string) {
	fmt.Println("=============== " + str + " BEGIN ===============")
}

func Separate(str string) {
	fmt.Println("----------- " + str + " -----------")
}

func End(str string) {
	fmt.Println("=============== " + str + " END ===============")
	fmt.Println()
}
