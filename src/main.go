// *两个要点：
// 1 GO语言要求代码文件必须声明package包
// 2 main包下的main()方法才是可执行的
package main

/*
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
*/

import "fmt"

func main() {
	fmt.Println("hello world")
}
