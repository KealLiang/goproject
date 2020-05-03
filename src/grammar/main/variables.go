package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
	##go中的变量

	* go中的变量，通过首字母大小写表示共有或私有
		- 首字母大写 public
		- 首字母小写 private

	* 额外知识点 UTF-8 是 Unicode 编码的一种具体实现
*/

func varType() {
	fmt.Println("=== GO中的变量类型 BEGIN ===")
	fmt.Println()
	fmt.Println("--- 基本数据类型 ---")
	var a, b, c, d = 100, -60.5, 'c', true
	fmt.Println("数值、字符（没有，用byte来保存单个字符）、布尔、字符串 a, b, c(注意), d =", a, b, c, d)
	fmt.Println("数值类型如下：")
	fmt.Println("int\tint8\tint16\tint32\tint64\nuint\tuint8\tuint16\tuint32\tuint64\tuintptr")
	fmt.Println("byte(uint8的别名)\trune(int32的别名，表示一个unicode码点)")
	fmt.Println("float32\tfloat64\ncomplex64\tcomplex128")
	fmt.Println()
	fmt.Println("--- 派生/复杂数据类型 ---")
	fmt.Println("1 指针")
	fmt.Println("2 数组")
	fmt.Println("3 结构体")
	fmt.Println("4 管道")
	fmt.Println("5 函数（对，这也是数据类型）")
	fmt.Println("6 切片")
	fmt.Println("7 接口")
	fmt.Println("8 map")
	fmt.Println("=== GO中的变量类型 END ===")
	fmt.Println()
}

func showVarType() {
	fmt.Println("========== 查看变量的类型 BEGIN ==========")
	var i = 100               // 默认是int64
	var i2 int32 = 1000000000 // go遵守保小不保大的原则，即在保证程序正确运行的情况下，尽量使用占用空间小的类型
	i3 := -39.12              // 浮点数都是有符号的
	i4 := 5.12e12             // 科学计数法
	var bb bool = true        // go中的Boolean类型只能取true或false，不能为空
	i5 := 3.14
	fmt.Printf("i=%d 的数据类型是 %T; 占用的字节数是 %d \n", i, i, unsafe.Sizeof(i))
	fmt.Printf("i2=%d 的数据类型是 %T; 占用的字节数是 %d \n", i2, i2, unsafe.Sizeof(i2))
	fmt.Printf("i3=%f 的数据类型是 %T; 占用的字节数是 %d \n", i3, i3, unsafe.Sizeof(i3))
	fmt.Printf("i4=%f 的数据类型是 %T; 占用的字节数是 %d \n", i4, i4, unsafe.Sizeof(i4))
	fmt.Printf("bb=%t 的数据类型是 %T; 占用的字节数是 %d \n", bb, bb, unsafe.Sizeof(bb))
	fmt.Printf("i5=%v 的数据类型是 %T; 占用的字节数是 %v \n", i5, i5, unsafe.Sizeof(i5)) // %v 是按照原始数据类型输出
	fmt.Println("========== 查看变量的类型 END ==========")
	fmt.Println()
}

func printCharAndString() {
	fmt.Println("========== 字符变量 BEGIN ==========")
	var a byte = 'A'
	//var b byte = '北' // 存不下了 constant 21271 overflows byte
	var b int = '北' // 换int就能放下了
	fmt.Printf("a = %c 占用的字节数是 %d;\n", a, unsafe.Sizeof(a))
	fmt.Printf("b = %c 占用的字节数是 %b; b 中存的是%d\n", b, unsafe.Sizeof(b), b)

	fmt.Println("----------- 字符串 -----------")
	var name string = "法外狂徒张三在此"
	//name[0] = 'c' //go中的字符串一旦赋值就不能再修改了，类似java
	fmt.Printf("name = %s; name中的第一个字符是：%c\n", name, name[0])
	str := `abc\ndef`
	code := `
		var i = 100 // 默认是int64
		var i2 int32 = 1000000000 // go遵守保小不保大的原则，即在保证程序正确运行的情况下，尽量使用占用空间小的类型
		`
	fmt.Println("使用反引号（esc下面的）输出带有特殊字符的字符串str：", str)
	fmt.Println("使用反引号（esc下面的）输出带有特殊字符的字符串code：", code)

	fmt.Println("========== 字符变量 END ==========")
	fmt.Println()
}

func typeMigration() {
	fmt.Println("========== go中的数据类型只能显示转换 BEGIN ==========")
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 float64 = float64(n1)
	fmt.Println("int 转换为 float:", n1)
	fmt.Println("float32 转换为 float64（低精度到高精度，也需要显示转换）:", n2)
	fmt.Printf("注意：转换只是值的转换，原来的 i 的数据类型不会变化：%T\n", i)

	fmt.Println("----------- 基本数据类型转字符串 -----------")
	var v1 int = 10
	var v2 float32 = 3.14
	var v3 bool = true
	var v4 byte = 'C'
	var str string
	// 方式一 fmt.Sprintf()
	str = fmt.Sprintf("%d", v1)
	fmt.Println("fmt.Sprintf()转为string", str)
	str = fmt.Sprintf("%f", v2)
	fmt.Println("fmt.Sprintf()转为string", str)
	// 方式二 strconv包的函数
	fmt.Println("strconv包函数转为string", strconv.FormatBool(v3))
	str = fmt.Sprintf("%c", v4)
	fmt.Println("strconv包函数转为string", strconv.FormatInt(int64(v4), 10)) //第二个参数是进制
	fmt.Println("strconv包函数Itoa()转int为string", strconv.Itoa(v1))

	fmt.Println("----------- 字符串转基本数据类型 -----------")
	var str1 string = "true"
	var str2 string = "8.99"
	bb, _ := strconv.ParseBool(str1) //我只想用第一个返回值，因此使用 _ 占位消掉第二个返回
	ff, _ := strconv.ParseFloat(str2, 64)
	fmt.Println("strconv包函数转string为基本类型", bb)
	fmt.Println("strconv包函数转string为基本类型", ff)

	fmt.Println("========== go中的数据类型只能显示转换 END ==========")
	fmt.Println()
}

func inferType() {
	var i = 10.22 // 省略类型定义
	fmt.Println("省略类型定义，自动推导变量类型：i =", i)
}

func defaultValueDemo() {
	fmt.Println("go中所有的数据类型都有默认值，赞！")
	var i int
	var s string
	var f float32
	var b bool
	fmt.Println("变量声明后不赋值，就会使用默认值：i =", i)
	fmt.Println("变量声明后不赋值，就会使用默认值：s =", s)
	fmt.Println("变量声明后不赋值，就会使用默认值：f =", f)
	fmt.Println("变量声明后不赋值，就会使用默认值：b =", b)
}

func ellipsisVar() {
	name := "张三"
	fmt.Println("省略var来定义变量，注意冒号不能省略 name =", name)
}

func multiValue() {
	var n1, n2, n3 int
	fmt.Println("一次声明多个变量 n1, n2, n3 =", n1, n2, n3)
	var a, b, c = 100, "Tom", 99.9
	fmt.Println("一次声明多个变量，并赋值 a, b, c =", a, b, c)
	d, e, f := 101, "Jerry", 98.9
	fmt.Println("一次声明多个变量，并赋值，省略var d, e, f =", d, e, f)
}

//全局变量，必须用var定义
var x, y, z = 102, "Marry", 100.0

func globalValue() {
	fmt.Println("全局声明多个变量，并赋值 x, y, z =", x, y, z)
}

func changeType() {
	i := 10
	fmt.Println("变量改变类型测试，原值为 i =", i)
	//i = 10.5
	fmt.Println("被赋值后变量就不能改变类型，否则将编译错误 constant 10.5 truncated to integer i =", i)
}

func main() {
	varType()
	showVarType()
	printCharAndString()
	typeMigration()
	inferType()
	defaultValueDemo()
	ellipsisVar()
	multiValue()
	globalValue()
	changeType()
}
