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

##go中的运算符
* 算数运算符
* 赋值运算符
* 比较运算符
* 逻辑运算符
* 位运算符
* 其他运算符 &(取址) *(取值)

##go中引包示例
* 注意：需要将 D:\coding\Go\fundamental 加入Project GOPATH GoLand才有提示
    - 将当前目录加入GOPATH后，跳过src写文件夹名（斜杠分隔）即可import
    - eg.
        "grammar/constant"
        _package "grammar/package" // 包起别名
        "util/logs"

* import的是包，使用也是[包名.变量名]，不是文件名！