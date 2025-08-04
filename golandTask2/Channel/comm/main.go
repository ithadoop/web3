/* 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。 */
package main

import (
	"fmt"
	"time"
)

func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 发送整数到通道
		time.Sleep(100 * time.Millisecond) // 模拟延时
	}
	close(ch) // 关闭通道
}

func printNumbers(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num) // 从通道接收整数并打印
	}
}

func main() {
	ch := make(chan int) // 创建通道

	go generateNumbers(ch) // 启动生成数字的协程
	go printNumbers(ch)    // 启动打印数字的协程

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)
	fmt.Println("All numbers processed.")
}