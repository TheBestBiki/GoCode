package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
消费者协程从通道中接收这些整数并打印。
考察点：通道的缓冲机制、生产者-消费者模式
*/

func main() {
	// 创建带缓冲的通道，缓冲区大小为10
	bufferedChan := make(chan int, 10)

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2) // 两个 goroutine：生产者和消费者

	fmt.Println("=== 带缓冲通道的生产者-消费者示例 ===")
	fmt.Printf("缓冲区大小: %d\n", cap(bufferedChan))
	fmt.Printf("初始通道长度: %d\n\n", len(bufferedChan))

	// 生产者 goroutine
	go func() {
		defer wg.Done()

		fmt.Println("生产者开始工作...")
		for i := 1; i <= 100; i++ {
			// 发送数据到缓冲通道
			bufferedChan <- i

			// 显示通道状态
			if i <= 20 || i%20 == 0 { // 前20个和每20个显示一次状态
				fmt.Printf("生产者发送: %d, 当前通道长度: %d/%d\n",
					i, len(bufferedChan), cap(bufferedChan))
			}

			// 模拟生产耗时
			if i%10 == 0 {
				time.Sleep(10 * time.Millisecond)
			}
		}

		close(bufferedChan) // 生产完成，关闭通道
		fmt.Println("\n生产者工作完成，通道已关闭")
	}()

	// 消费者 goroutine
	go func() {
		defer wg.Done()

		fmt.Println("消费者开始工作...")
		count := 0

		// 使用 range 接收数据直到通道关闭
		for num := range bufferedChan {
			count++
			fmt.Printf("消费者接收: %d", num)

			// 显示通道状态
			if count <= 20 || count%20 == 0 {
				fmt.Printf(" (当前通道剩余: %d)", len(bufferedChan))
			}
			fmt.Println()

			// 模拟消费耗时
			if count%25 == 0 {
				time.Sleep(15 * time.Millisecond)
			}
		}

		fmt.Printf("\n消费者工作完成，共处理 %d 个数字\n", count)
	}()

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("程序执行完毕")

	// 演示缓冲机制的优势
	demonstrateBufferingAdvantage()
}

// demonstrateBufferingAdvantage 演示缓冲机制的优势
func demonstrateBufferingAdvantage() {
	fmt.Println("\n=== 缓冲机制优势演示 ===")

	// 无缓冲通道示例
	fmt.Println("1. 无缓冲通道:")
	unbuffered := make(chan int)

	go func() {
		fmt.Print("   发送 1...")
		unbuffered <- 1 // 会阻塞，直到有人接收
		fmt.Println(" 发送完成")
	}()

	time.Sleep(100 * time.Millisecond) // 让发送者阻塞一会儿
	fmt.Print("   接收...")
	<-unbuffered
	fmt.Println(" 接收完成")

	// 带缓冲通道示例
	fmt.Println("2. 带缓冲通道:")
	buffered := make(chan int, 2)

	fmt.Print("   发送 1...")
	buffered <- 1 // 不会阻塞，因为有缓冲
	fmt.Println(" 发送完成")

	fmt.Print("   发送 2...")
	buffered <- 2 // 不会阻塞
	fmt.Println(" 发送完成")

	fmt.Print("   发送 3...") // 这次会阻塞，因为缓冲满了
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(" 接收...")
		<-buffered
		fmt.Println(" 接收完成")
	}()

	buffered <- 3
	fmt.Println(" 发送完成")

	close(unbuffered)
	close(buffered)
}
