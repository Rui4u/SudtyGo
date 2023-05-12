package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}

func getCPUInfo(ctx context.Context) {
	defer func() {
		wg.Done()
	}()
	fmt.Printf("%s\n", ctx.Value("key")) // traceID
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("获取CPU信息")
		}
	}
}

func getCPUInfo2(ctx context.Context) {
	defer func() {
		wg2.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束2")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("获取CPU信息2")
		}
	}
}
func main() {
	wg.Add(1)
	//wg2.Add(1)

	//  context 提供了三个函数  withCancel  withTimeOut WithValue
	//  如果你的goroutine  函数 ctx前加上ctx。 因为是个interface 所以可以任意实现
	//  context 是树状结构， 调用父interface中的接口，会调用所有子节点里面的接口  cancel具有传递性
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx)
	ctx3 := context.WithValue(ctx2, "key", "value")
	go getCPUInfo(ctx3)
	//go getCPUInfo2(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	//wg2.Wait()
}
