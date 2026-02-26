&Aa 返回变量指针地址
*Aa 返回变量的值，一般来说，此时的Aa是指针地址
defer 用于延迟执行函数调用，被 defer 的语句会在包含它的函数返回之前执行。


#常用函数
time.Now()   // 获取当前时间

done := make(chan bool, 2) // 缓冲通道



#最佳实践
1.多线程处理任务示例，创建多个goroutine处理任务
task2/t2-2/testScheduler.go

2.关于结构体需要实现接口的所有方法才算实现接口的解释
task2/t2-3/ShapeTest2.go

3.goroutine和channel的配合使用
task2/t2-4/channelTest.go
