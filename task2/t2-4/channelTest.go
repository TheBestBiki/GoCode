package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。考察点 ：通道的基本使用、协程间通信。
*/

func main() {
	// 创建带缓冲的通道
	ch := make(chan int, 10)

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2) // 两个 goroutine

	// 生产者 goroutine
	go func() {
		defer wg.Done() // 完成时通知 WaitGroup
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("发送数字: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道。这里发送通道关闭后，不影响后面的消费者读取
		fmt.Println("通道已关闭")
	}()

	// 消费者 goroutine
	go func() {
		defer wg.Done()       // 完成时通知 WaitGroup
		for num := range ch { // 使用 range 优雅地处理通道关闭。这里 range 循环会自动处理通道关闭，通道一有值，这里也立马能读取到
			fmt.Printf("接收到数字: %d\n", num)
		}
		fmt.Println("所有数字接收完成")
	}()

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
