# 注意：
window: go 1.20  
macOS: go 1.25.0

# 基础
### 常用
&Aa 返回变量指针地址  
*Aa 返回变量的值，一般来说，此时的Aa是指针地址  
defer 用于延迟执行函数调用，被 defer 的语句会在包含它的函数返回之前执行。  

### 常用函数  
time.Now()   // 获取当前时间  
number2CountMap := make(map[int]int)  
done := make(chan bool, 2) // 缓冲通道  

### 通道的缓冲和无缓冲
无缓冲  
bufferedChan := make(chan int)  
有缓冲  
bufferedChan := make(chan int, 10)  
区别是，有缓冲的通道，就跟java线程池的核心线程一样，多个线程可以同时处理任务，当缓冲区别被占满时，比如这里有10个都在处理中，  
那么第11个就会进入等待；  
但是无缓冲的通道，只能处理一个任务，其他任务会等待。  
注意：  
关键理解点  
每次 <-channel 只取一个值  
按照 FIFO（先进先出）顺序消费  
**如果通道为空时执行 <-channel 会阻塞**  
**如果通道满时执行 channel <- value 会阻塞**  

### for value := range buffered解读
关键特性总结  
实时响应：一旦通道有数据就立即消费  
自动阻塞：通道为空时自动阻塞等待  
优雅退出：检测到通道关闭后自动退出循环  
FIFO顺序：严格按照**先进先出**的顺序消费  



# 最佳实践
### 1.多线程处理任务示例，创建多个goroutine处理任务
task2/t2-2/testScheduler.go

### 2.关于结构体需要实现接口的所有方法才算实现接口的解释
task2/t2-3/ShapeTest2.go

### 3.goroutine和channel的配合使用
task2/t2-4/channelTest.go

### 4.多线程万能公式：goroutine和WaitGroup
task2/t2-4/channelTest.go
