package main

import "fmt"

func main() {
	if 3 > 2 {
		fmt.Println("true")
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 0 {
		}
	}

	// for 02_range

	nums := "abcdefg"

	for index, value := range nums {
		fmt.Println(index)
		fmt.Printf("%c", value)
	}
}
