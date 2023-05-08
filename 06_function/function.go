package main

import "fmt"

type SRInt int // Main.SRint

type SR2Int = int // int

func (this SRInt) addOne() (result int) {
	result = int(this) + 1
	return
}

func main() {
	var i SRInt = 10
	var j SR2Int = 20
	i.addOne()

	fmt.Printf("%t", i)
	fmt.Printf("%t", j)

}
