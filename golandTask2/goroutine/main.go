package main

import (
	"fmt"
	"time"
)

/* Goroutine
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。 */

func printOddNumbers() {
	fmt.Println("Odd:")
	for i := 1; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond) // 模拟延时
		fmt.Println("Odd:", i)
	}
}

func printEvenNumbers() {
	for i := 2; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond) // 模拟延时
		fmt.Println("Even:", i)
	}
}


func main() {
	fmt.Println("Starting Goroutines for Odd and Even Numbers")

	go printOddNumbers()  // 启动协程打印奇数
	
	go printEvenNumbers() // 启动协程打印偶数
	

	// 等待一段时间确保所有goroutine完成
    time.Sleep(time.Second)

} 

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	for i := 2; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond) // 模拟延时
		fmt.Println("Even:", i)
	}
}
/* 
func main() {

	go say("in goroutine: world")
	say("hello")

	
} */
