package main

import "fmt"

func test() {
	collection1 := []int{1, 2, 3, 4, 5}    // type:[]int
	collection2 := [...]int{1, 2, 3, 4, 5} //type:[5]int
	collection3 := []int{1, 2, 3, 4, 5}

	for i, i2 := range collection1 {
		fmt.Print(i)
		fmt.Println(i2)
	}

	for i, i2 := range collection2 {
		fmt.Print(i)
		fmt.Println(i2)
	}

	collection1 = append(collection1, 4)

	courseSlice := collection1[0:3] // 0..<3

	fmt.Print(courseSlice)

	courseSlice2 := make([]string, 1)
	var _ = make([]string, 3)

	fmt.Printf("%t", courseSlice2)
	courseSlice2[0] = "1"
	//courseSlice2[1] = "2"
	fmt.Println(courseSlice2)

	fmt.Println(collection1[0:3])
	fmt.Println(collection1[0:])
	fmt.Println(collection1[:3])

	collection1 = append(collection1, collection3...) // ... 将collection3打散

	// slice 的默认值是nil   map也一样
	var sliceA []int32                // 是nil
	var sliceB []int = make([]int, 0) //是empty

	sliceA = append(sliceA, 3)
	fmt.Println(sliceA)
	fmt.Println(sliceB)

}
