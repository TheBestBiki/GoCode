package main

import (
	"fmt"
	"sync"
	"time"
)

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。考察点 ： go 关键字的使用、协程的并发执行
*/
/*
等待方式1：适合测试
*/
func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("打印奇数：", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("打印偶数：", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// 等待 goroutine 执行完成。 若不加这个暂停，则主程序一结束，goroutine哪怕没执行完，也会立即结束，所有 goroutine 都会被终止。加上这个才能看到上面的打印
	time.Sleep(2 * time.Second)
}

/*
*
等待方式2：sync.WaitGroup	精确控制，无需估计时间	需要管理计数	已知数量的任务
*/
func main2() {
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个 goroutine

	go func() {
		defer wg.Done() // 完成后通知
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("打印奇数：", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("打印偶数：", i)
			}
		}
	}()

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("所有 goroutine 执行完毕")
}

/*
等待方式2：channel	灵活，可以传递数据	代码稍复杂	需要通信或复杂同步
*/
func main3() {
	done := make(chan bool, 2) // 缓冲通道

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("打印奇数：", i)
			}
		}
		done <- true // 发送完成信号
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("打印偶数：", i)
			}
		}
		done <- true
	}()

	// 等待两个 goroutine 完成
	<-done
	<-done

	//优雅一点的写法
	// 等待所有 goroutine 完成
	for i := 0; i < 2; i++ {
		<-done
	}
	fmt.Println("所有任务完成")
}
