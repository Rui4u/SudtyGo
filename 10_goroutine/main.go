package main

import "fmt"

/*
go 语言 goroutine 通信
不要通过共享内存来通信， 而是通过通信来实现内存共享
php python java 多线程编程，通过一个全局对象
go 希望提供队列方式，消费者与生产者之间的关系

channel +语法糖  使用channel更加简单
*/
// channel 是一个通道  底层是一个环形数组  多线程安全的
func main() {
	var msg chan string
	msg = make(chan string, 1) // 设置成0叫 无缓冲channel

	msg <- "sr" // 放值到chan中
	//msg <- "sr"   // 放值到chan中  // 放入的值超过缓存个数的时候，会阻塞
	data := <-msg // 取值

	fmt.Print(data)

	nobufferChannel()
}

/*
无缓冲channel适用于通知，B要第一时间通知A是否已经完成
有缓冲channel适合消费者和生产者之间的通讯
go channnel
1. 消息传递、消息过滤
2. 信号广播
3. 时间订阅和广播
4. 任务分发
5. 结果汇总
6. 并发控制
7. 同步和异步

默认情况下 channel是双向的。  但是数据传递希望他是单向的
*/
func nobufferChannel() {
	var msg chan string
	msg = make(chan string, 0)

	go func(msg chan string) { //happen before机制 可以保障
		fmt.Print(<-msg) //取值
	}(msg)
	msg <- "test"
}
