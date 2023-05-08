package main

import "fmt"

func main() {
	dictFunction()
}

func dictFunction() {
	var dict map[string]string = map[string]string{"a": "a"}
	dict["key"] = "123"
	fmt.Println("", dict["key"])

	dict2 := make(map[string]string, 2)
	dict2["key"] = "1234"
	fmt.Println("", dict2["key"])

}
