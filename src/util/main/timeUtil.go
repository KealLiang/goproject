package main

import (
	"fmt"
	"math/rand"
	"time"
	"util/logs"
)

/*
	go中的时间相关方法
*/

func sleepDemo() {
	logs.Begin("go 中的睡眠")
	rand.Seed(time.Now().Unix()) // 随机种子
	for i := 0; i < 10; i++ {
		fmt.Println(">> 输出一个随机数", rand.Intn(10))
		time.Sleep(50 * time.Millisecond)
	}
}

func unitTime() {
	logs.Begin("获取当前的unix时间戳（随机种子）")
	now := time.Now()
	fmt.Println("当前的unix秒时间为：", now.Unix())
	fmt.Println("当前的unix纳秒时间为：", now.UnixNano())
}

func formatTime() {
	logs.Begin("格式化时间输出")
	now := time.Now()
	// 必须使用这个时间当格式，才能得到正确的结果 "2006-01-02 15:04:05"
	fmt.Println("当前时间是：", now.Format("2006-01-02 15:04:05.000"))
}

func main() {
	formatTime()
	unitTime()
	sleepDemo()
}
