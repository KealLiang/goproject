package main

import (
	"encoding/json"
	"fmt"
	"util/logs"
)

/**
##go中的结构体（对象）
* 结构体field在内存中是连续分布的
* 结构体能强转的前提是，两种结构体所有的field都一样

###结构体中tag的用法

##go中的方法绑定
* 语法 func (变量名 绑定的类) 方法 {}
* 自定义类型都可以有方法，不仅仅是在结构体中
* 方法绑定后，每次调用都会将变量值拷贝，为了避免这问题，可以定义为绑定指针 func (变量名 *绑定的类) 方法 {}
* 判断调用时是 值拷贝还是地址拷贝 看的是方法的定义，不是看调用方，别被语法糖混淆了（调用方可以写成 (&zhangsan).getGender()）

##go中的OOP
* go中用类似组合的方式实现继承
* 区别
	- 继承是匿名的
	- 组合是命名的 （go中推荐组合）

##go中的接口
* go中没有implement这样的关键字
* 接口和类是隐式关联上的，只要类实现了接口定义的方法即可
* 一个变量A实现了某个接口I的[全部方法]，才能说A实现了I
* 由于go中实现接口的特殊方式，go中的所有结构体都默认实现了空接口 interface{} （相当于Object）
* interface类型默认是一个指针（引用类型）
	- 如果没有对interface初始化就使用，会输出nil
* 类型断言
	1. 解决 判断当前接口的实现类是谁的问题
	1. 类型断言可以用在switch语句上：switch item.(type) { case int : ... }

*/

func TypeJudge(items ...interface{} /*根据go中的接口实现方式，所有类型都默认实现了空接口，因此这里能接收任何类型*/) {
	for index, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("第[%v]个数据[%v]是 bool 类型\n", index, item)
		case int, int8, int16, int32:
			fmt.Printf("第[%v]个数据[%v]是 int 类型\n", index, item)
		case float32, float64:
			fmt.Printf("第[%v]个数据[%v]是 float 类型\n", index, item)
		case string:
			fmt.Printf("第[%v]个数据[%v]是 string 类型\n", index, item)
		case Monster:
			fmt.Printf("第[%v]个数据[%v]是 Monster 类型\n", index, item)
		case *Monster:
			fmt.Printf("第[%v]个数据[%v]是 *Monster 类型\n", index, item)
		default:
			fmt.Printf("第[%v]个数据[%v]是 其他 类型\n", index, item)
		}
	}
}
func TypeJudgeDemo() {
	logs.Begin("go 中类型判断示例")
	var n1 float32 = 2.4
	var n2 int64 = 2
	var n3 int = 100000
	var name = "tom"
	address := "上海"
	var monster Monster = Monster{"金角", 600}
	var monsterPtr *Monster = &Monster{"银角", 600}

	TypeJudge(n1, n2, n3, name, address, monster, monsterPtr)
}

type HDMI interface {
	Working()
}

type USB interface {
	Working()
}
type Upan struct{}
type Camera struct{}

// go中的接口是隐式关联上的，只要结构体实现了接口定义的方法即可
func (u Upan) Working() {
	fmt.Println("读取U盘中的数据...")
}
func (c Camera) Working() {
	fmt.Println("摄像头开始拍摄你的帅气脸庞...")
}

// 定义一个电脑
type Computer struct{}

func (c Computer) PlugIn(usb USB) { //接口是引用类型，因此这里不会产生值拷贝问题
	if cam, ok := usb.(Camera); ok { //类型断言
		fmt.Println("类型断言：若实现类是camera则先检查是否有内存卡...", cam)
	}
	usb.Working() //电脑的插入方法，调用了usb接口的working()方法，
}

func interfaceDemo() {
	logs.Begin("go 中的接口示例")
	computer := Computer{}

	upan := Upan{}
	camera := Camera{}

	logs.Separate("U盘插入电脑")
	computer.PlugIn(upan) //Upan实现了Usb接口，这里才能这样传参
	logs.Separate("摄像头插入电脑")
	computer.PlugIn(camera)

	logs.Separate("不同接口，有同名的方法")
	var hdmi HDMI = camera
	hdmi.Working()
}

type Creature struct {
	Type string
}
type Character struct {
	Char string
}
type Animal struct {
	Creature //嵌套匿名结构体 实现继承
	Gender   bool
}

func (a Animal) Eat(food string) {
	fmt.Printf("[%v]吃了一个[%v]\n", a.Type, food)
}

type Dog struct {
	Ani        Animal //命名结构体 实现组合
	*Character        //避免值拷贝
	Name       string
	Color      string
}

func extendDemo() {
	logs.Begin("go 中的继承与初始化示例")
	// 两种实例化的方式：命名与非命名的方式不能混用
	dog := Dog{
		Ani: Animal{
			Creature: Creature{Type: "哺乳动物"},
			Gender:   false,
		},
		Name:  "小飞侠",
		Color: "黄色",
	}
	dog2 := Dog{
		Animal{Creature{"节肢动物"}, false},
		&Character{"温顺的"},
		"大牛",
		"黑色",
	}
	fmt.Println("实例化的dog对象：", dog)
	dog.Ani.Eat("骨头")
	fmt.Println("实例化的dog2对象：", dog2)
	fmt.Println("实例化的dog2对象，发现Character是个指针，需要这样取值：", dog2, *dog2.Character)
	fmt.Println("dog2的种类和性格分别是：", dog2.Ani.Type, dog2.Char)
}

type integer int

func (i *integer) oddOrEven() string {
	if *i%2 == 0 {
		return "偶数"
	} else {
		return "奇数"
	}
}
func selfMethodDemo() {
	logs.Begin("自定义类型的方法")
	var num integer = 12345
	fmt.Printf("num = %v 是 [%v]\n", num, num.oddOrEven())
}

type Person struct {
	Name    string
	Age     uint
	Gender  bool
	Hobbies []string
}

//将此方法绑定在Person结构体上
//func (p Person) getGender() string { //注意1：这里的p和调用是传入的不是同一个p，每次调用都会值拷贝一次，效率较低
//	fmt.Printf("getGender方法中的p[%p]\n", &p)
func (p *Person) getGender() string { //注意2：定义成指针就解决了上诉问题，p.Gender是语法糖
	fmt.Printf("getGender方法中的p[%p]\n", p)
	if p.Gender {
		return "男"
	} else {
		return "女"
	}
}
func (p *Person) String() string {
	return fmt.Sprintf("自定义toString方法 -> Person{%v,%v,%v,%v}", p.Name, p.Age, p.Gender, p.Hobbies)
}
func structMethodDemo() {
	logs.Begin("结构体的方法")
	zhangsan := Person{"张三", 19, true, nil}
	fmt.Println("张三的人设是：", zhangsan)
	fmt.Println("张三的人设是：", &zhangsan) //实现了String方法后可以这么调用到
	//若使用了[注意2]的写法，这里的正确调用方式应该是 (&zhangsan).getGender()
	fmt.Println("张三的性别是：", (&zhangsan).getGender())
	zhangsan.Gender = !zhangsan.Gender
	fmt.Println("张三变性后性别是：", zhangsan.getGender()) //zhangsan.getGender()其实是语法糖
	fmt.Printf("structMethodDemo中的zhangsan[%p]\n", &zhangsan)
}

// 通过加tag的方式给field起别名
type Monster struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func structSerialization() {
	logs.Begin("结构体的序列化")
	monster := Monster{"铁扇公主", 600}
	json, err := json.Marshal(monster) // Marshal中使用了反射
	if err != nil {
		fmt.Println("序列化异常：", monster)
	}
	fmt.Println("返回的json为：", string(json))
}

type A struct {
	Num int
}
type B struct {
	Num int
	//Price float64
}

func structConvert() {
	logs.Begin("结构体的类型转换")
	var a A
	var b B
	a = A(b) //结构体能转换的前提是A B 所有的field都一样
	fmt.Printf("a[%p] b[%p]\n", &a, &b)
}

type Cat struct {
	Name  string
	Age   uint
	Color string
}

func createStruct() {
	logs.Begin("创建结构体的4种方式")
	logs.Separate("第一种方式")
	var cat1 Cat
	cat1.Name = "小花"
	cat1.Age = 3
	cat1.Color = "黄白黑"
	fmt.Println("第一只：", cat1)

	logs.Separate("第二种方式（推荐）")
	var cat2 Cat = Cat{"BlackPrince", 4, "黑色"}
	fmt.Println("第二只：", cat2)

	logs.Separate("第三种方式")
	cat3 := new(Cat) // new出来是个指针
	fmt.Printf("cat3的类型是[%T]\n", cat3)
	(*cat3).Name = "史密斯" //正确的赋值方式
	cat3.Age = 10        //语法糖，编译器底层加上了(*ptr)
	cat3.Color = "黄色"
	fmt.Println("第三只：", cat3)

	logs.Separate("第四种方式")
	cat4 := &Cat{"混乱小魔王", 1, "纯白"}
	fmt.Printf("cat4的类型是[%T]\n", cat4)
	(*cat4).Age = 2
	cat4.Color = "米白" //同第三种，也是语法糖
	fmt.Println("第四只：", cat4)
}

func main() {
	createStruct()
	structConvert()
	structSerialization()
	structMethodDemo()
	selfMethodDemo()
	extendDemo()
	interfaceDemo()
	TypeJudgeDemo()
}
