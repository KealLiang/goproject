package main

/*
	##go中引包示例
	* 注意：需要将 D:\coding\Go\fundamental 加入Project GOPATH GoLand才有提示
		- 将当前目录加入GOPATH后，跳过src写文件夹名（斜杠分隔）即可import
		- eg.
			"grammar/constant"
			_package "grammar/package" // 包起别名
			"util/logs"

	* import时，路径从 $GOPATH 的src下开始，不用带src（命令行可以 echo %GOPATH% 查看）
	* import的是文件夹，使用也是[包名.变量名]，不是文件名！
*/

import (
	"fmt"
	"grammar/constant"         // import的是文件夹名，文件里的包名可以和文件夹名不同，使用时用文件里的包名做前缀
	_package "grammar/package" // 包起别名
	"util/logs"
)

func referenceConst() {
	logs.Begin("引用了不同目录下的方法")
	fmt.Printf("引用了同目录下的常量，注意是[包名.变量名] 派 = %v, 普朗克常数 = %v, PubVar = %v\n", fix.PI, fix.H, _package.PubVar)
	logs.End("引用了不同目录下的方法")
}

func main() {
	referenceConst()
}
