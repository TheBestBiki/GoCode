package main

import (
	"fmt"
	"sync"
)

/*
简化版有锁计数器
*/
func main() {
	lockCount := &LockCount{count: 0}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lockCount.increment()
			}
		}()
	}

	wg.Wait()

	if lockCount.getValue() == 10*1000 {
		fmt.Println("有锁并发安全")
	} else {
		fmt.Println("有锁并发不安全: ", lockCount.getValue())
	}

}

type LockCount struct {
	count int
	mutex sync.Mutex
}

func (lockCount *LockCount) increment() {
	lockCount.mutex.Lock()
	defer lockCount.mutex.Unlock()
	lockCount.count++
}

func (lockCount *LockCount) setValue(value int) {
	lockCount.mutex.Lock()
	defer lockCount.mutex.Unlock()
	lockCount.count = value
}

func (lockCount *LockCount) getValue() int {
	lockCount.mutex.Lock()
	defer lockCount.mutex.Unlock()
	return lockCount.count
}
