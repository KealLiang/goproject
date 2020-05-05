package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"unsafe"
	"util/logs"
)

/*
	##go中的变量
	* 变量的分类
		- 值类型 直接指向数据空间
		- 引用类型 指向数据空间的地址

	* go中的变量，通过首字母大小写表示共有或私有
		- 首字母大写 public
		- 首字母小写 private

	* go中的map
		- keyType 不能是slice\map\function 因为不能用 == 来比较
		- valueType 和keyType一样

	* 额外知识点 UTF-8 是 Unicode 编码的一种具体实现

	* go支持自定义数据类型（c语言的那套）
		- 语法 type myInt int // go中认为myInt和int是两个类型（即使他们都是int）

*/

func mapDemo() {
	logs.Begin("go 中的map演示")
	logs.Separate("第一种建立map的方式")
	var m1 map[string]int //声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
	m1 = make(map[string]int, 2)
	m1["思考"] = 1 //没分配空间会报错
	m1["行动"] = 2
	m1["反思"] = 3
	m1["创造"] = 4
	fmt.Println("新建了一个map，map是无序的：", m1)

	logs.Separate("第二种建立map的方式")
	var m2 = make(map[string]map[string]int) //省略大小默认就1个
	m2["广西"] = make(map[string]int)          //内层还是必须先make才能赋值
	m2["广西"]["南宁"] = 10
	fmt.Println("又新建了一个map：", m2)

	logs.Separate("第三种建立map的方式")
	m3 := map[string]string{
		"广东": "广州",
		"江苏": "南京",
	}
	m3["四川"] = "成都"
	fmt.Println("再新建了一个map：", m3)

	logs.Separate("map元素查找与删除")
	v, ok := m3["广东"]
	if ok {
		fmt.Println("是否找到：", ok, v)
	}
	delete(m3, "haha") //删除元素

	logs.Separate("map的遍历（只能 for-range）")
	for key, value := range m3 {
		fmt.Println(key, value)
	}
}

func mapSlice() {
	logs.Begin("go 中的map切片操作")
	// 声明一个map切片
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "紫霞仙子"
		monsters[1]["age"] = "1500"
	}

	// 通过append追加，避免越界
	monsters = append(monsters, map[string]string{"name": "火云邪神", "age": "100"})
	fmt.Println("怪物列表：", monsters)
}

func sliceDemo2() {
	logs.Begin("go 中的切片操作2")
	var s1 []int = []int{100, 200, 300}
	s2 := append(s1, 400, 500)
	fmt.Println("切片s1是：", s1)
	fmt.Println("切片s2是：", s2)
	fmt.Printf("append本质是底层new了一个新的数组并copy元素，地址和原来不同：\n&s1 = %p; &s2 = %p\n", &s1, &s2)

	logs.Separate("string进行切片处理")
	str := "lsrmod2014@163.com"
	fmt.Println("对字符串切片得到邮箱地址：", str[10:])
	runes := []rune(str) // []byte中文会有问题，[]rune都可以
	runes[0] = 'L'
	str = string(runes)
	fmt.Println("string是不可变的，想改需要先变为切片，再转回byte数组：", str)
}

func sliceDemo() {
	logs.Begin("go 中的切片（引用类型）")
	var hens [10]int // 8字节
	for i := 0; i < 10; i++ {
		hens[i] = rand.Intn(20)
	}
	fmt.Println("hens数组为：", hens)
	// 切片存的有三个部分： 起始地址, 长度, 容量
	logs.Separate("方式一：通过数组切出切片")
	slice := hens[1:5]
	fmt.Println("hens[1:5]切片为：", slice)
	fmt.Printf("这两个地址相等 &hens[1] = %p; &slice[0] = %p;\n", &hens[1], &slice[0])
	fmt.Println("此例中slice存的是： &hens[1], 4, 9")
	fmt.Println("slice切片长度为：", len(slice))
	fmt.Println("slice切片容量为：", cap(slice)) //切片的容量是动态变化的

	logs.Separate("方式二：使用make创建切片")
	var s1 []float64 = make([]float64, 5, 10) //底层还是数组，但这个数组是不可见的，只能通过切片访问
	fmt.Println("使用make([]float64, 5, 10)初始化的切片：", s1)

	logs.Separate("方式三：定义一个切片，直接就指定具体的数组")
	var s2 []string = []string{"tom", "jerry", "mike"}
	fmt.Println("s2 的元素为：", s2)
	fmt.Println("s2 的长度为：", len(s2))
	fmt.Println("s2 的容量为：", cap(s2))
}

func arrayDemo() {
	logs.Begin("go 中的数组（值类型）")
	var a1 []int //这里的a1，定义时没有写大小，其实是切片(slice)
	a2 := [...]int{1, 2}
	fmt.Printf("a1的类型是%T（没大小的是切片）\na2的类型是%T（go中长度是数组类型的一部分）\n", a1, a2)
	var array = [...]string{1: "张三", 0: "李四", 9: "王五"} //指定下标
	fmt.Println("array 的 length和值 分别是：", len(array), array)

	var hens [10]int // 8字节
	for i := 0; i < 10; i++ {
		hens[i] = rand.Intn(20)
	}
	fmt.Println("生成的数组为：", hens)
	fmt.Printf("hens的地址是：%p\n", &hens)
	fmt.Println("hens[0]的地址是：", &hens[0])
	fmt.Println("hens[1]的地址是，hens[1]-hens[0]是 8字节(64位)：", &hens[1])
	fmt.Println("hens[2]的地址是，hens[2]-hens[1]是 8字节(64位)：", &hens[2])
}

func defindSelfType() {
	logs.Begin("go 中自定义数量类型")
	type myInt int
	var i myInt = 10
	fmt.Println("自定义的 i 值为：", i)
	var i2 int = int(i)
	fmt.Println("即使myInt本身就是int，go也认为他们不同，必须显示转换 i2 值为：", i2)
}

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
	//varType()
	//showVarType()
	//printCharAndString()
	//typeMigration()
	//inferType()
	//defaultValueDemo()
	//ellipsisVar()
	//multiValue()
	//globalValue()
	//changeType()
	//defindSelfType()
	//arrayDemo()
	//sliceDemo()
	//sliceDemo2()
	mapDemo()
	mapSlice()
}
