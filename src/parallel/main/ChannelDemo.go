package main

/*
	##go中的管道
	* 管道是线程（协程）安全的
	* 协程间通讯可以通过管道
	* 管道是 FIFO

	###管道的死锁
	* 提示：fatal error: all goroutines are asleep - deadlock!
	* 编译器底层会分析：当管道内的数据__不流动__时，才会报死锁（这点非常重要，仔细品go语言的管道的设计理念）
		- 生产慢消费快 不会死锁
		- 生产快消费慢 不会死锁
		- 修改例子writeAndReadDemo()中的rate就能看到效果

	###多线程工具类比
	* CountDownLatch -> WaitGroup

	###协程的ID
	* 在go语言中，谷歌开发者不建议大家获取协程ID，主要是为了GC更好的工作，滥用协程ID会导致GC不能及时回收内存

*/

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
	"util/logs"
)

func channelSelectDemo() {
	logs.Begin("使用select解决从管道读数据阻塞问题")
	intChan := make(chan int, 5)
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		intChan <- i
	}
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("string数据 %v", i)
	}

label:
	for {
		// 管道不close，一直读会导致死锁，但是实际开发中close的时机不是那么容易确认的，这时就要用另一种方式
		//fmt.Println(<-intChan, <-stringChan)
		time.Sleep(200 * time.Millisecond)
		select {
		case v := <-intChan:
			fmt.Println("从intChan取到数据：", v)
		case v := <-stringChan:
			fmt.Println("从stringChan取到数据：", v)
		default:
			fmt.Println("取不到数据了，但是我也不会报错\"all goroutines are asleep - deadlock!\"哦 (*^_^*)")
			break label
		}
	}
	logs.End("使用select解决从管道读数据阻塞问题")
}

func channelReadWriteDemo() {
	logs.Begin("只读只写管道示例")
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 3)
	go send(ch, exitChan) // 一个生产者，两个消费者
	go recv(ch, exitChan)
	go recv(ch, exitChan)

	total := 0
	for _ = range exitChan {
		total++
		if total == 3 {
			break
		}
	}
	logs.End("只读只写管道示例")
}

func recv(ch <-chan int, exitChan chan struct{}) {
	// ch只读
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("接收者[%v]接收到了：%v\n", goID(), v)
		//time.Sleep(time.Nanosecond)
	}
	exitChan <- struct{}{}
}

func send(ch chan<- int, exitChan chan struct{}) {
	// ch只写
	for i := 0; i < 6; i++ {
		fmt.Println("写入：", i)
		ch <- i
		time.Sleep(time.Nanosecond)
	}
	close(ch)
	exitChan <- struct{}{} //放入一个空结构体表示已处理完成
}

// 获取协程id（不建议）
func goID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)] // 这里的b拿到是：goroutine 7 [running]: main.goID(...) D:/coding/Go/fundamental/

	b = b[:bytes.IndexByte(b, '[')-1]
	return string(b)

	//b = bytes.TrimPrefix(b, []byte("goroutine "))
	//b = b[:bytes.IndexByte(b, ' ')]
	//n, _ := strconv.ParseUint(string(b), 10, 64)
	//return n
}

func findPrimeNumberDemo() {
	logs.Begin("找素数的案例")
	upper := 100000
	start := time.Now()
	//primeNumbers := findPrimeNumber(upper) //单线程的方式
	primeNumbers := findPrimeNumberParallel(upper, 4) //4个协程大概就是单线程的4倍
	fmt.Println("计算耗时：", time.Since(start))
	fmt.Printf("2~%v 中有 %v 个素数\n", upper, len(primeNumbers))
	for _, v := range primeNumbers {
		fmt.Printf("%v ", v)
	}
	fmt.Println("")
}

// 多协程
func findPrimeNumberParallel(upper int, routines int) []int {
	intChan := make(chan int, upper)
	primeChan := make(chan int, upper)
	exitChan := make(chan bool, routines) // 可以用WaitGroup代替exitChan

	go initNumbers(intChan, upper)
	for i := 0; i < routines; i++ {
		go calcPrimeNumbers(intChan, primeChan, exitChan)
	}

	// 匿名函数，若exitChan退够4个则完成
	go func() {
		for i := 0; i < routines; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()

	// 主线程必须要用到，否则立马退出
	primeNums := make([]int, 0, upper/2)
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		primeNums = append(primeNums, res)
	}
	fmt.Println("")
	fmt.Println("多协程计算完成")
	return primeNums
}

func calcPrimeNumbers(intChan chan int, primeChan chan int, exitChan chan bool) {
outer:
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				continue outer
			}
		}
		primeChan <- num
	}
	// 渣算法
	//for {
	//	num, ok := <-intChan
	//	if !ok {
	//		break
	//	}
	//	flag := true
	//	for i:=2;i<num;i++ {
	//		if num%i == 0 {
	//			flag = false
	//		}
	//	}
	//	if flag {
	//		primeChan<- num
	//	}
	//}
	exitChan <- true
}

func initNumbers(intChan chan int, upper int) {
	for i := 2; i <= upper; i++ {
		intChan <- i
	}
	close(intChan)
}

// 单线程
func findPrimeNumber(upper int) []int {
	primeNums := make([]int, 0, upper/2)
outer:
	for i := 2; i <= upper; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue outer
			}
		}
		primeNums = append(primeNums, i)
	}
	// 渣算法
	//for i:=2;i<=upper;i++ {
	//	flag := true
	//	for j:=2;j<i;j++ {
	//		if i%j == 0 {
	//			flag = false
	//		}
	//	}
	//	if flag {
	//		primeNums = append(primeNums, i)
	//	}
	//}
	return primeNums
}

func writeAndReadDemo() {
	logs.Begin("生产者与消费者的基本 协程-管道 演示")
	// 独写的协程，通过channel通讯
	// 先创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	// 开始生产与消费
	go writeData(intChan, 20)
	go readData(intChan, exitChan, 30)

	// 如果不用标记的方式就只好让主线程sleep了
	//time.Sleep(10 * time.Second)

	// 主线程中如何知道何时结束，就通过exitChan
	//return // 作为对比直接return的话，程序就是秒返回
	// 为什么不直接使用变量代替exitChan，为了协程安全与通讯
	// 标记一：
	//_, ok := <-exitChan
	//if !ok {
	//	return // 管道被关闭后主线程退出
	//}
	// 标记二：
	v := <-exitChan
	if v {
		return // exitChan中拿到true标记后退出
	}
}

func writeData(intChan chan int, rate int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写数据：", i)
		time.Sleep(time.Duration(rate) * time.Millisecond)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool, rate int) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读到数据：", v)
		time.Sleep(time.Duration(rate) * time.Millisecond)
	}
	// 读取完数据，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//writeAndReadDemo()
	//findPrimeNumberDemo()
	//channelReadWriteDemo()
	channelSelectDemo()
}
