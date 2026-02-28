package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
简化版无锁计数器
*/
func main() {
	// var声明，需要具体的数据类型，跟初始化
	// sync/atomic 包的函数需要特定的类型 func AddInt64(addr *int64, delta int64) int64
	var count int64 = 0
	// 不能用这个,因为这个是短变量声明，是一个局部变量。短变量声明没法指定具体类型
	// count :=0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()

	finalValue := atomic.LoadInt64(&count)
	if finalValue == 10*1000 {
		fmt.Println("有锁并发安全")
	} else {
		fmt.Println("有锁并发不安全: ", finalValue)
	}

}
