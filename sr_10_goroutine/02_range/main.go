package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int
	msg = make(chan int, 2)

	go func(msg chan int) {
		for data := range msg { //通过range来获取值
			fmt.Println(data)
		}
		fmt.Print("done")
	}(msg)

	msg <- 1
	msg <- 2

	fmt.Println("123")
	close(msg) // for 02_range 退出  close后不能放值，可以取值

	fmt.Println(<-msg)
	time.Sleep(time.Second * 10)
}
