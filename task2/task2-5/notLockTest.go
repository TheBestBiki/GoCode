package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
题目：使用原子操作（sync/atomic包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点：原子操作、并发数据安全、无锁编程。
*/

func main() {
	// 使用 int64 类型的计数器
	var counter int64 = 0

	// 使用 WaitGroup 等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(10) // 10个协程

	fmt.Println("=== 原子操作无锁计数器演示 ===")
	fmt.Printf("初始计数器值: %d\n", atomic.LoadInt64(&counter))
	fmt.Printf("协程数量: %d, 每个协程递增次数: %d\n", 10, 1000)
	fmt.Println("开始并发执行...")

	startTime := time.Now()

	// 启动10个协程
	for i := 0; i < 10; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			// 每个协程执行1000次原子递增
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增操作

				// 每100次显示一次进度（只显示前3个协程的进度）
				if goroutineID < 3 && j%100 == 0 && j > 0 {
					currentValue := atomic.LoadInt64(&counter)
					fmt.Printf("协程%d: 已完成%d次递增，当前计数器值: %d\n",
						goroutineID, j, currentValue)
				}
			}

			finalValue := atomic.LoadInt64(&counter)
			fmt.Printf("协程%d: 完成1000次递增操作，当前计数器值: %d\n",
				goroutineID, finalValue)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	executionTime := time.Since(startTime)
	finalValue := atomic.LoadInt64(&counter)

	// 输出最终结果
	fmt.Println("\n=== 执行结果 ===")
	fmt.Printf("期望结果: %d\n", 10*1000)
	fmt.Printf("实际结果: %d\n", finalValue)
	fmt.Printf("执行时间: %v\n", executionTime)

	// 验证结果正确性
	if finalValue == 10*1000 {
		fmt.Println("✅ 结果正确！原子操作成功保护了共享数据")
	} else {
		fmt.Printf("❌ 结果错误！期望%d，实际%d\n", 10*1000, finalValue)
	}

	// 演示更多原子操作功能
	demonstrateAtomicOperations()

	// 对比 Mutex 和 Atomic 的性能
	comparePerformance()
}

// demonstrateAtomicOperations 演示各种原子操作
func demonstrateAtomicOperations() {
	fmt.Println("\n=== 原子操作功能演示 ===")

	var value int64 = 42

	// 原子加载
	fmt.Printf("原子加载: %d\n", atomic.LoadInt64(&value))

	// 原子存储
	atomic.StoreInt64(&value, 100)
	fmt.Printf("原子存储后: %d\n", atomic.LoadInt64(&value))

	// 原子交换
	oldValue := atomic.SwapInt64(&value, 200)
	fmt.Printf("原子交换: 旧值=%d, 新值=%d\n", oldValue, atomic.LoadInt64(&value))

	// 原子比较并交换 (CAS)
	// 如果当前值等于 200，则设置为 300
	if atomic.CompareAndSwapInt64(&value, 200, 300) {
		fmt.Printf("CAS成功: %d\n", atomic.LoadInt64(&value))
	}

	// 原子递增和递减
	atomic.AddInt64(&value, 10) // +10
	fmt.Printf("原子递增+10后: %d\n", atomic.LoadInt64(&value))

	atomic.AddInt64(&value, -5) // -5
	fmt.Printf("原子递减-5后: %d\n", atomic.LoadInt64(&value))
}

// comparePerformance 对比 Mutex 和 Atomic 的性能
func comparePerformance() {
	fmt.Println("\n=== Mutex vs Atomic 性能对比 ===")

	const iterations = 100000
	const goroutines = 10

	// 测试 Atomic 性能
	var atomicCounter int64
	var wg1 sync.WaitGroup
	wg1.Add(goroutines)

	start1 := time.Now()
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg1.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg1.Wait()
	time1 := time.Since(start1)

	// 测试 Mutex 性能
	var mutexCounter int64
	var mutex sync.Mutex
	var wg2 sync.WaitGroup
	wg2.Add(goroutines)

	start2 := time.Now()
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg2.Done()
			for j := 0; j < iterations; j++ {
				mutex.Lock()
				mutexCounter++
				mutex.Unlock()
			}
		}()
	}
	wg2.Wait()
	time2 := time.Since(start2)

	fmt.Printf("Atomic 方式: %v (结果: %d)\n", time1, atomic.LoadInt64(&atomicCounter))
	fmt.Printf("Mutex 方式: %v (结果: %d)\n", time2, mutexCounter)
	fmt.Printf("性能差异: Mutex 比 Atomic 慢 %.2fx\n", float64(time2)/float64(time1))

	fmt.Println("\n=== 原子操作优势总结 ===")
	fmt.Println("1. 无锁设计：不需要显式的锁机制")
	fmt.Println("2. 性能更好：避免了锁竞争和上下文切换")
	fmt.Println("3. 更简单：API 直观易用")
	fmt.Println("4. 适用场景：简单的数值操作（计数、标志位等）")
	fmt.Println("5. 局限性：只能处理基本数据类型，复杂逻辑仍需 Mutex")
}
