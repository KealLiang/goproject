package main

import (
	"fmt"
	"strconv"
	"strings"
	"util/logs"
)

/*
	go中的字符串操作
*/

func baseDemo() {
	logs.Begin("基础字符串操作的函数")
	str := "hello世界"
	fmt.Printf("go中的编码统一为utf8，len是按照字节数统计的：%v len = %v\n", str, len(str))
	logs.Separate("-")

	fmt.Println("取出字符串中的字符的时候，需要先转为rune的切片 []rune ，否则是按照字节取：")
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("[%c] ", str2[i])
	}
	fmt.Println()
	logs.Separate("-")

	atoi, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误！", err)
	} else {
		fmt.Println("strconv包下的 string转数字：", atoi)
	}
	logs.Separate("-")

	fmt.Println("字符串转[]byte：", []byte(str))
	fmt.Println("[]byte转字符串：", string([]byte{97, 98, 99}))
	logs.Separate("-")

	fmt.Println("十进制17转 2 8 16进制：", strconv.FormatInt(17, 2),
		strconv.FormatInt(17, 8),
		strconv.FormatInt(17, 16))
	logs.Separate("-")

	haystack := "mississippi"
	needle := "ssis"
	fmt.Printf("查找字串[%v]是否在指定字符串[%v]中：%t\n", needle, haystack, strings.Contains(haystack, needle))
	fmt.Printf("查找字串[%v]在指定字符串[%v]中的位置：%v\n", needle, haystack, strings.Index(haystack, needle))
}

func stringCompare() {
	logs.Begin("GO中 字符串比较可以使用双等号")
	haystack := "mississippi"
	needle := "mississippi"
	fmt.Printf("[%v] == [%v]：%t\n", needle, haystack, haystack == needle) // 考虑大小写

	logs.Separate("忽略大小写比较")
	fmt.Println("ABC 与 abc 比较，不区分大小写：", strings.EqualFold("ABC", "abc"))
}

func main() {
	stringCompare()
	baseDemo()
}
