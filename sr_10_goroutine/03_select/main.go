package main

import (
	"fmt"
	"time"
)

/*
03_select 类似于switch case  03_select 的功能和linux里提供的io的select、poll、epoll基本相同
主要作用于多个channel

有个需求：
两个gorountine都在执行， 希望在主grountine中，知道什么时候立即完成，是哪一个完成
*/

func main() {

	chan1 := make(chan struct{})
	chan2 := make(chan struct{})

	go go1(chan1)
	go go2(chan2)

	// 监控2个返回值，任何一个channel返回都知道
	// 某一个分支就绪了就执行该分支，， 因为正常的<-chan 会阻塞
	// 时间相同，就随机，

	tc := time.NewTimer(1 * time.Second) //设置超时时间

	select {
	case <-chan1:
		fmt.Println("g1 done")
	case <-chan2:
		fmt.Println("g2 done")
	case <-tc.C:
		fmt.Println("time out")
	}
}

func go1(msg chan struct{}) {
	time.Sleep(3 * time.Second)
	msg <- struct{}{}
}

func go2(msg chan struct{}) {
	time.Sleep(2 * time.Second)
	msg <- struct{}{}
}
