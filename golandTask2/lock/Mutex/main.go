/* 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。 */

package main
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // 创建一个互斥锁
	counter := 0

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock() // 上锁
				counter++ // 对计数器进行递增操作
				mu.Unlock() // 解锁
			}
		}()
	}

	wg.Wait() // 等待所有协程完成

	fmt.Println("Final Counter Value:", counter) // 输出计数器的最终值
}