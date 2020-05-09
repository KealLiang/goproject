## 运行go程序
* 编译方式
    - cd到.go文件目录
    - 执行 go build xxx.go
    - 运行 xxx.exe
* 直接run方式
    - cd到.go文件目录
    - 执行 go run xxx.go
    
| 比较 | 本质 | 性能 | 平台差异 |
| :-----| ----: | ----: | ----: |  
| 编译方式 | 二进制 | 执行快 | 会编译成可执行文件 |
| run方式 | 脚本形式 | 每次都要隐式编译 | 一样 |

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

##go中的函数
* 参数传递机制
    - 基本数据类型和数组：值传递（值拷贝）；希望修改的话可以传递指针 &变量名

* go不支持传统的重载
    - 通过空接口实现

* go中函数也是一个变量（函数式编程）

##go内存模型
* 逻辑上：栈、堆、代码区

##go中包管理
* 注意：需要将 D:\coding\Go\fundamental 加入Project GOPATH GoLand才有提示
    - 将当前目录加入GOPATH后，跳过src写文件夹名（斜杠分隔）即可import
    - eg.
        "grammar/constant"
        _package "grammar/package" // 包起别名
        "util/logs"

* import时，路径从 $GOPATH 的src下开始，不用带src（命令行可以 echo %GOPATH% 查看）
* import的是文件夹，使用也是[包名.变量名]，不是文件名！
* 如果要编译为一个可执行文件，就需要将包声明为main

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