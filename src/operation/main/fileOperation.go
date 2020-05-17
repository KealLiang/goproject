package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"util/logs"
)

/*
	##go中的文件操作
	* 同java都是以流的形式操作
	* 文件对象 == 文件指针 == 文件句柄

	###判断文件路径是否存在
	* 见PathExist()函数，go中需要根据os.Stat(path)返回的err来判断
*/

func PathExist(path string) (bool, error) {
	// 封装判断文件是否存在的函数
	_, err := os.Stat(path)
	if err == nil { // 文件或目录是存在的
		return true, nil
	}
	if os.IsNotExist(err) { // 文件或目录不存在
		return false, nil
	}
	return false, err
}

func readAndWriteDemo() {
	logs.Begin("读写文件案例")
	fileName := "src/operation/resource/sunflower"                              // 文件的相对路径就不能省略src了
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) // 第三个参数FileMode只在linux下有效
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// 读
	logs.Separate("读")
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Print(readString)
		if err != nil {
			break
		}
	}

	// 写
	logs.Separate("写")
	content := "Hello, 世界，时间啊停止吧！"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(content + "\n")
	}

	// 因为缓存的存在，需要flush才能落盘
	writer.Flush()
}

func readFileWithBufferDemo() {
	logs.Begin("使用缓冲读取文件")
	fileName := "src/operation/resource/demo.txt" // 文件的相对路径就不能省略src了
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	logs.Separate("用缓冲读，默认4096字节")
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Print(readString)
		if err == io.EOF { //这里的err是预判的，换言之此时readString还有最后一行的内容
			fmt.Println("")
			break
		}
	}

	logs.Separate("一次性读入整个文件（不用显示关闭）")
	readFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", readFile)
	fmt.Printf("%s\n", readFile)

	logs.End("使用缓冲读取文件")
}

func baseDemo() {
	logs.Begin("文件的基本操作")
	fileName := "D:/coding/Go/fundamental/.gitignore"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件出错！", err)
	}

	fmt.Printf("File's Content = %v\n", file) // 可以看到就是个指针

	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件出错！", err)
	}
}

func main() {
	//baseDemo()
	//readFileWithBufferDemo()
	readAndWriteDemo()
}
