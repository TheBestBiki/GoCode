package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。考察点 ：协程原理、并发任务调度
任务调度器-简洁版本
*/
func main() {
	// 创建调度器，最大并发3个任务
	scheduler := NewTaskScheduler(3)

	// 添加任务
	scheduler.AddTask(func() error {
		time.Sleep(200 * time.Millisecond)
		return nil
	})

	scheduler.AddTask(func() error {
		time.Sleep(300 * time.Millisecond)
		return nil
	})

	scheduler.AddTask(func() error {
		time.Sleep(150 * time.Millisecond)
		return fmt.Errorf("模拟错误")
	})

	scheduler.AddTask(func() error {
		time.Sleep(400 * time.Millisecond)
		return nil
	})

	scheduler.AddTask(func() error {
		time.Sleep(250 * time.Millisecond)
		return nil
	})

	fmt.Println("开始执行任务...")
	results := scheduler.Run()

	// 统计
	fmt.Println("\n=== 执行结果统计 ===")
	var total time.Duration
	for taskID, duration := range results {
		fmt.Printf("任务%d: %v\n", taskID, duration)
		total += duration
	}

	fmt.Printf("\n所有任务完成，总耗时: %v\n", total)
	if len(results) > 0 {
		avg := total / time.Duration(len(results))
		fmt.Printf("平均任务耗时: %v\n", avg.Round(time.Millisecond))
	}
}

// TaskScheduler 简洁版任务调度器
type TaskScheduler struct {
	tasks   []func() error
	workers int
}

// NewTaskScheduler 创建调度器
func NewTaskScheduler(workers int) *TaskScheduler {
	return &TaskScheduler{
		workers: workers,
	}
}

// AddTask 添加任务
func (ts *TaskScheduler) AddTask(task func() error) {
	ts.tasks = append(ts.tasks, task)
}

// Run 执行所有任务并返回结果
func (ts *TaskScheduler) Run() map[int]time.Duration {
	// 使用 sync.WaitGroup 等待所有任务完成
	var wg sync.WaitGroup
	result := make(map[int]time.Duration)
	// 使用 sync.Mutex 保护共享数据。用于加锁
	var mu sync.Mutex

	// 控制并发数
	semaphore := make(chan struct{}, ts.workers)

	for i, task := range ts.tasks {
		wg.Add(1)

		go func(taskID int, t func() error) {
			defer wg.Done()

			// 获取信号量（控制并发）
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// 执行并计时
			start := time.Now()
			err := t()
			duration := time.Since(start)

			// 保存结果
			mu.Lock()
			result[taskID] = duration
			mu.Unlock()

			// 输出结果
			status := "成功"
			if err != nil {
				status = fmt.Sprintf("失败: %v", err)
			}
			fmt.Printf("任务%d %s, 耗时: %v\n",
				taskID, status, duration.Round(time.Millisecond))
		}(i, task)
	}

	wg.Wait()
	return result
}
