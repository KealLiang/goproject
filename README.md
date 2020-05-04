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

* go中的变量，通过首字母大小写表示共有或私有
    - 首字母大写 public
    - 首字母小写 private

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