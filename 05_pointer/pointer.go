package main

import "fmt"

func main() {
	type Persiton struct {
		name string
		age  int
	}

	var p1 *Persiton = &Persiton{
		"13",
		19,
	}
	p1.name = "123"
	fmt.Println(p1)
	fmt.Printf("%p", p1)

	var p2 *Persiton // nil
	var p3 Persiton  // Persion

	fmt.Println(p2)
	fmt.Println(p3)

	// 指针最好make初始化

}
