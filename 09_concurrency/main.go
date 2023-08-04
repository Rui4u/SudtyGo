package main

import (
	"sync"
)

/*
多线程：经过操作系统管理。 线程切换代价高
go 内部自己管理。所以内存和性能比较好 GMP机制  节省线程切换，节省context 以及regiter file的存储

	生成多个goruntine, 通过调度器映射到线程上。 M:N  M是线程 N是任务  固定线程池。
	如果其中一个goruntine有耗时操作，处理器会挂起操作。执行后面的goruntine
	调度器是多个processor的组合
	G：goroutine   M: thread  P：processor

多线程保证数据安全：
1. 锁、
2. 原子化
*/

/*
 */

var total int64 = 0
var wg = sync.WaitGroup{}
var lock = sync.Mutex{}

func main() {
	// atomic.AddInt64(&total,1)   原子性相加
	// 监控多少个goruntine
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	print(total)
}

func add() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		lock.Lock()
		total++
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		lock.Lock()
		total--
		lock.Unlock()
	}
}
