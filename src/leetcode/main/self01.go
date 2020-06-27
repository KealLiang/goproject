package main

import (
	"fmt"
	"time"
)

/*
	计算n!
	1、递归
	2、动态规划
*/
var num int = 1

func fibonacciDynamic(n uint64) uint64 {
	slot := make([]uint64, n)
	slot[0] = 1
	slot[1] = 1
	for i := 2; i < int(n); i++ {
		slot[i] = slot[i-1] + slot[i-2]
	}
	return slot[n-1]
}

func fibonacciRecursion(n uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

func factorialDynamic(n uint64) uint64 {
	slot := make([]uint64, n+1)
	slot[0] = 1
	slot[1] = 1
	var i uint64 = 2
	for ; i <= n; i++ {
		slot[i] = i * slot[i-1]
	}
	return slot[n]
}

func factorialRecursion(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorialRecursion(n-1)
}

func run(n uint64) {
	now := time.Now()
	for i := 0; i < num; i++ {
		//factorialRecursion(n)
		//factorialDynamic(n)
		println(fibonacciRecursion(n))
		//println(fibonacciDynamic(n))
	}
	fmt.Println("计算耗时（毫秒）：", (time.Now().UnixNano()-now.UnixNano())/1000/1000)
}

func main() {
	run(48)
}
