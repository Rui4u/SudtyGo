package main

func main() {
	var ch1 chan int //双向

	/**
	var ch2 01_chan<- float64 //单向   只能存
	var ch3 <-01_chan float64 //单向  只能读取
	*/
	ch1 = make(chan int, 2)
	go func(msg <-chan int) {

	}(ch1)

}
