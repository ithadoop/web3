/* 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。 */

package main
import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i // 发送整数到通道
		time.Sleep(50 * time.Millisecond) // 模拟生产延时
	}
	close(ch) // 关闭通道
}
//
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num) // 从通道接收整数并打印
		time.Sleep(100 * time.Millisecond) // 模拟消费延时
	}
}

func main() {
	ch := make(chan int, 10) // 创建一个带缓冲的通道，缓冲区大小为10

	go producer(ch) // 启动生产者协程
	go consumer(ch) // 启动消费者协程

	// 等待一段时间确保所有goroutine完成
	time.Sleep(10 * time.Second)
	fmt.Println("All numbers processed.")
}

