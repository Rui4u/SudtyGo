package main

import (
	"fmt"
	"strconv"
)

func main() {

	BasicDataType()
}

func VarAndConst() {

	var (
		test = "aa"
		age  = 10
	) //多变量定义

	var num int = 10
	var _ int // _占位符， 避免定义后不使用的错误
	name := 1 // 简洁变量定义方式 相当于var name = 1
	//  定义一个变量  相当于var name = 1. 定义不用会报错

	// 常量
	const PI float32 = 3.1415926
	/*
		多常量定义，如果未赋值，按上一个类型和值走
	*/
	const (
		a = 10
		b
		c = "123"
		d
	)

	// iota 是一个特殊的常量,可以被编译器修改 自动递增
	const (
		ERR1 = iota
		ERR2 = iota
		ERR3 = iota
		ERR4 = iota
	)

	const (
		ERR5 = iota
		ERR6
		ERR7 = "123"
		ERR8
		ERR9 = iota
		ERR10
	)
	fmt.Println(name, age, test, num)
}

func BasicDataType() {
	var a byte
	var c rune
	a = 'a'
	c = '沙'

	b := 97

	fmt.Println(a)
	fmt.Printf("b=%c\n", b)
	fmt.Printf("c=%c\n", c)

	var d int32 = 12222
	var e = int8(d)
	var h = strconv.Itoa(int(e))
	fmt.Println(e)
	fmt.Println(h)

	f := "123"
	g, error := strconv.Atoi(f)
	fmt.Println(g)
	fmt.Println(error)
}

/*  go语言是静态语言
	1. 变量有零值， 也就是默认值
    2. 简洁变量定义方式，不能再全局使用
    3. 变量定义不使用会报错
*/
/*  go语言是静态语言

 */

/*

 */
