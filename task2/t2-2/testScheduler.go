package main

import (
	"fmt"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。考察点 ：协程原理、并发任务调度
3种方式实现简易版
*/
func main() {
	scheduler := getScheduler(3)
	scheduler.add(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("这是第1个任务")
	})
	scheduler.add(func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("这是第2个任务")
	})
	scheduler.add(func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("这是第3个任务")
	})

	//scheduler.run()
	//time.Sleep(3 * time.Second)

	// scheduler.run2()
	scheduler.run3()

}

type SchedulerFactory struct {
	maxWorkers int
	task       []func()
}

func getScheduler(maxWorkers int) *SchedulerFactory {
	return &SchedulerFactory{
		maxWorkers: maxWorkers,
	}
}

func (s *SchedulerFactory) add(function func()) {
	s.task = append(s.task, function)
}

func (s *SchedulerFactory) run() {
	for index, task := range s.task {
		go runTask(index, task)
	}
}

func (s *SchedulerFactory) run2() {
	var wg sync.WaitGroup
	wg.Add(len(s.task))
	for index, task := range s.task {
		go func(i int, t func()) {
			// defer 用于延迟执行函数调用，被 defer 的语句会在包含它的函数返回之前执行。
			/*
				func example() {
					fmt.Println("1. 开始执行")
					defer fmt.Println("3. defer 语句")  // 最后执行
					fmt.Println("2. 正常语句")
					// 函数即将返回时执行 defer 语句
				}
			*/
			defer wg.Done()
			runTask(i, t)
		}(index, task)
	}
	wg.Wait()
}

func (s *SchedulerFactory) run3() {
	done := make(chan bool, len(s.task))
	for index, task := range s.task {
		go func(i int, t func()) {
			runTask(i, t)
			done <- true
		}(index, task)

		// 下面这种是错的，done <- true应该放在函数里面，所以可以在runTask外面再包一层
		// go runTask(index, task)
		// done <- true
	}

	for i := 0; i < len(s.task); i++ {
		<-done
	}
}

func runTask(index int, task func()) {
	now := time.Now()
	task()
	fmt.Println("执行任务", index+1, "花费的时间为：", time.Now().Sub(now).Milliseconds())
}
