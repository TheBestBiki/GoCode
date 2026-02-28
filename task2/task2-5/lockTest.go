package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点：sync.Mutex 的使用、并发数据安全。
*/

// Counter 安全计数器结构体
type Counter struct {
	value int
	mutex sync.Mutex
}

// Increment 安全递增方法
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// GetValue 获取当前值（只读操作也需要保护）
func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

// SetValue 设置值（写操作需要保护）
func (c *Counter) SetValue(val int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = val
}

func main() {
	// 创建计数器实例
	counter := &Counter{value: 0}

	// 使用 WaitGroup 等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(10) // 10个协程

	fmt.Println("=== 并发安全计数器演示 ===")
	fmt.Printf("初始计数器值: %d\n", counter.GetValue())
	fmt.Printf("协程数量: %d, 每个协程递增次数: %d\n", 10, 1000)
	fmt.Println("开始并发执行...")

	startTime := time.Now()

	// 启动10个协程
	for i := 0; i < 10; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			// 每个协程执行1000次递增
			for j := 0; j < 1000; j++ {
				counter.Increment()

				// 每100次显示一次进度（只显示前3个协程的进度）
				if goroutineID < 3 && j%100 == 0 && j > 0 {
					fmt.Printf("协程%d: 已完成%d次递增\n", goroutineID, j)
				}
			}

			fmt.Printf("协程%d: 完成1000次递增操作\n", goroutineID)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	executionTime := time.Since(startTime)

	// 输出最终结果
	fmt.Println("\n=== 执行结果 ===")
	fmt.Printf("期望结果: %d\n", 10*1000)
	fmt.Printf("实际结果: %d\n", counter.GetValue())
	fmt.Printf("执行时间: %v\n", executionTime)

	// 验证结果正确性
	if counter.GetValue() == 10*1000 {
		fmt.Println("✅ 结果正确！Mutex 成功保护了共享数据")
	} else {
		fmt.Printf("❌ 结果错误！期望%d，实际%d\n", 10*1000, counter.GetValue())
	}

	// 演示不使用 Mutex 的问题
	demonstrateRaceCondition()
}

// demonstrateRaceCondition 演示不使用 Mutex 会出现的竞争条件
func demonstrateRaceCondition() {
	fmt.Println("\n=== 竞争条件演示 ===")
	fmt.Println("下面演示不使用 Mutex 时可能出现的问题：")

	unsafeCounter := 0
	var unsafeWg sync.WaitGroup
	unsafeWg.Add(2)

	// 启动两个协程同时修改同一个变量
	go func() {
		defer unsafeWg.Done()
		for i := 0; i < 1000; i++ {
			unsafeCounter++ // 没有保护的并发访问
		}
	}()

	go func() {
		defer unsafeWg.Done()
		for i := 0; i < 1000; i++ {
			unsafeCounter++ // 没有保护的并发访问
		}
	}()

	unsafeWg.Wait()

	fmt.Printf("不使用 Mutex 的结果: %d (期望: %d)\n", unsafeCounter, 2000)
	if unsafeCounter != 2000 {
		fmt.Println("⚠️ 出现了竞争条件！多个协程同时修改同一变量导致数据丢失")
	}

	fmt.Println("\n=== Mutex 重要性总结 ===")
	fmt.Println("1. Mutex 提供互斥锁，确保同一时刻只有一个协程访问临界区")
	fmt.Println("2. Lock() 和 Unlock() 必须成对出现，通常使用 defer 确保 Unlock 被调用")
	fmt.Println("3. 所有对共享数据的读写操作都需要适当的同步保护")
	fmt.Println("4. 即使是看似简单的操作（如 ++）在并发环境下也可能出现问题")
}
