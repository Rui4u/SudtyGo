package main

import "fmt"

/*
*
支持interface嵌套
*/
type Duck interface {
	Run(string) error
	Gaga()
}

type pskDuck struct {
}

func (pds pskDuck) Gaga() {
	fmt.Println("gaga")
}

func (pds *pskDuck) Run(string string) (error error) {
	fmt.Println("run")
	return
}

func main() {
	var d Duck = &pskDuck{} // 当实现方法接收对象是指针的时候，这必须也是取值。
	d.Run("2")

	var p Persion = Student{PersionIm{}, "sha", 18}
	p.Info()
	a, ok := p.(Student)

	if !ok {
		fmt.Println("报错了")
	}
	fmt.Println(a)

	slice1 := []interface{}{1, 2, 3, 4, 5}
	InterfaceForBreak(slice1...)

	var slice2 = []string{"1", "2", "3"}
	var slice3 = []interface{}{}
	for _, s := range slice2 {
		slice3 = append(slice3, s)
	}
	InterfaceForBreak(slice3...) //如果定义的是[]string 的slice 此时会报错，必须转成[]interface 的slice
}

type Persion interface {
	Info()
}

type PersionIm struct {
}

func (p PersionIm) Info() {
	fmt.Println(" p  run")
}

type Student struct {
	Persion
	name string
	age  int
}

func InterfaceForBreak(datas ...interface{}) {
	for _, a := range datas {
		fmt.Println(a)
	}
}
