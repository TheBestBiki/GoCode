package main

import (
	"fmt"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。考察点 ：协程原理、并发任务调度
任务调度器-基础版本
*/
func main() {
	// 创建调度器（最大3个并发）
	scheduler := NewScheduler(3)

	// 添加示例任务
	tasks := createSampleTasks()
	for _, task := range tasks {
		scheduler.AddTask(task)
	}

	fmt.Println("开始执行任务...")
	start := time.Now()

	// 执行所有任务
	results := scheduler.Run()

	totalTime := time.Since(start)

	// 打印统计信息
	scheduler.PrintStatistics()
	fmt.Printf("\n调度器总执行时间: %v\n", totalTime.Round(time.Millisecond))

	// 输出详细结果
	fmt.Println("\n=== 详细结果 ===")
	for _, result := range results {
		fmt.Printf("任务%d: 耗时=%v, 错误=%v\n",
			result.TaskID, result.Duration, result.Err)
	}
}

// Task 定义任务类型
type Task func() error

// TaskResult 任务执行结果
type TaskResult struct {
	TaskID    int           // 任务ID
	Duration  time.Duration // 执行时间
	Err       error         // 错误信息
	StartTime time.Time     // 开始时间
	EndTime   time.Time     // 结束时间
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks       []Task          // 任务列表
	maxWorkers  int             // 最大并发数
	wg          sync.WaitGroup  // 等待组
	results     chan TaskResult // 结果通道
	resultsLock sync.Mutex      // 结果锁
	taskResults []TaskResult    // 最终结果
}

// NewScheduler 创建新的调度器
func NewScheduler(maxWorkers int) *Scheduler {
	if maxWorkers <= 0 {
		maxWorkers = 3 // 默认3个并发
	}

	return &Scheduler{
		maxWorkers: maxWorkers,
		results:    make(chan TaskResult, 100), // 缓冲通道
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// worker 工作协程
func (s *Scheduler) worker(taskID int, task Task) {
	defer s.wg.Done()

	start := time.Now()

	// 执行任务
	err := task()

	end := time.Now()
	duration := end.Sub(start)

	// 发送结果到通道
	s.results <- TaskResult{
		TaskID:    taskID,
		Duration:  duration,
		Err:       err,
		StartTime: start,
		EndTime:   end,
	}
}

// collectResults 收集结果
func (s *Scheduler) collectResults() {
	for result := range s.results {
		s.resultsLock.Lock()
		s.taskResults = append(s.taskResults, result)
		s.resultsLock.Unlock()
	}
}

// Run 执行所有任务
func (s *Scheduler) Run() []TaskResult {
	totalTasks := len(s.tasks)

	// 启动结果收集协程
	go s.collectResults()

	// 创建任务通道
	taskChan := make(chan struct {
		ID   int
		Task Task
	}, totalTasks)

	// 启动工作协程
	s.wg.Add(s.maxWorkers)
	for i := 0; i < s.maxWorkers; i++ {
		go func(workerID int) {
			defer s.wg.Done()
			for item := range taskChan {
				s.worker(item.ID, item.Task)
			}
		}(i)
	}

	// 发送任务到通道
	for i, task := range s.tasks {
		taskChan <- struct {
			ID   int
			Task Task
		}{ID: i, Task: task}
	}
	close(taskChan)

	// 等待所有工作协程完成
	s.wg.Wait()

	// 关闭结果通道，结束结果收集
	close(s.results)

	// 等待一小段时间确保结果收集完成
	time.Sleep(100 * time.Millisecond)

	return s.taskResults
}

// PrintStatistics 打印统计信息
func (s *Scheduler) PrintStatistics() {
	fmt.Println("\n=== 任务执行统计 ===")

	var totalDuration time.Duration
	successCount := 0
	failCount := 0

	for _, result := range s.taskResults {
		totalDuration += result.Duration
		if result.Err == nil {
			successCount++
		} else {
			failCount++
		}

		status := "✓ 成功"
		if result.Err != nil {
			status = fmt.Sprintf("✗ 失败: %v", result.Err)
		}

		fmt.Printf("任务 %d: %s | 耗时: %v | 开始: %s | 结束: %s\n",
			result.TaskID,
			status,
			result.Duration.Round(time.Millisecond),
			result.StartTime.Format("15:04:05.000"),
			result.EndTime.Format("15:04:05.000"),
		)
	}

	fmt.Printf("\n总计: 成功 %d, 失败 %d, 总耗时: %v\n",
		successCount, failCount, totalDuration.Round(time.Millisecond))

	if len(s.taskResults) > 0 {
		avgDuration := totalDuration / time.Duration(len(s.taskResults))
		fmt.Printf("平均耗时: %v\n", avgDuration.Round(time.Millisecond))
	}
}

// 示例任务函数
func createSampleTasks() []Task {
	return []Task{
		func() error {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务1: 数据处理完成")
			return nil
		},
		func() error {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务2: 网络请求完成")
			return nil
		},
		func() error {
			time.Sleep(150 * time.Millisecond)
			fmt.Println("任务3: 文件读取完成")
			return nil
		},
		func() error {
			time.Sleep(400 * time.Millisecond)
			return fmt.Errorf("任务4: 数据库连接失败")
		},
		func() error {
			time.Sleep(250 * time.Millisecond)
			fmt.Println("任务5: 计算完成")
			return nil
		},
		func() error {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("任务6: 缓存更新完成")
			return nil
		},
	}
}
