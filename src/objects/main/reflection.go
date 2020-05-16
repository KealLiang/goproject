package main

import (
	"fmt"
	"reflect"
	"util/logs"
)

/*
	##go中的反射
	* 反射的本质是运行时操作程序
	* 学反射就是学Type和Value，这两者可以相互获取
		- Type	是个接口（reflect.TypeOf()）
		- Value	是个结构体（reflect.ValueOf()）

	###Type和Kind区分
	* Type 类型
	* Kind 类别
	* Type和Kind可能相同也可能不同
		- 例1：var num int = 10 那么num的Type是 int，Kind也是 int
		- 例2：var stu Student 那么stu的Type是 包名.Student，Kind是struct

	###反射涉及的类型转换（见convertValueDemo()方法）
	* 过程中主要有三种类型要区分
		- 变量
		- interface{}
		- reflect.Value()

	###通过反射修改变量
	* 当使用SetXXX方法，需要通过对应的指针类型才能改变传入的值
	* 同时配合reflect.Value().Elem()方法（类比 *ptr = newVal）
*/

func modifyValueVarReflectDemo() {
	logs.Begin("通过反射修改变量的值（值拷贝问题）")
	//var s int = 10
	//of := reflect.ValueOf(s)
	//of.SetInt(20) //报错，of需要是指针类型
	var name string = "法外狂徒-张三"
	modifyValue(&name) // 为了实现在方法中修改变量，需要传入地址
	fmt.Println("modifyValueVarReflectDemo() --> ", name)
}
func modifyValue(v interface{}) {
	of := reflect.ValueOf(v)
	//of.SetString(of.String() + " 已经伏法") //这里of是指针，运行时报错：reflect.Value.SetString using unaddressable value
	elem := of.Elem() // 通过Elem()又拿到指针指向的真正的值 类比 *ptr
	elem.SetString(elem.String() + " 已经伏法")
	fmt.Println("modifyValue() --> ", elem)
}

func convertValueDemo() {
	logs.Begin("反射涉及的类型转换")
	var a int = 100
	rVal := reflect.ValueOf(a)
	vInter := rVal.Interface()
	//vInter++ //error 编译时并不知道（虽然运行时可以知道）
	i := vInter.(int)
	i++

	i2 := rVal.Int() //直接调用方法转
	i2++

	//类型转换报错：call of reflect.Value.Float on int Value
	//f1 := rVal.Float()
	//fmt.Printf("f1[%T]=%v;\n", f1,f1)

	fmt.Printf("a[%T]=%v;\nrVal[%T]=%v;\nvInter[%T]=%v;\ni[%T]=%v;\n", a, a, rVal, rVal, vInter, vInter, i, i)
	fmt.Printf("i2[%T]=%v;\n", i2, i2)
}

func main() {
	//convertValueDemo()
	modifyValueVarReflectDemo()
}
