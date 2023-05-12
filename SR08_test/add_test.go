package ch08

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("错误了")
	}
}

func TestAdd2(t *testing.T) {

	if testing.Short() { //针对short模式跳过
		t.Skipf("跳过")
	}

	result := add(1, 2)
	if result != 3 {
		t.Errorf("错误了")
	}
}

func BenchmarkAdd(bb *testing.B) {

	a := 123
	b := 456
	c := 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Println("错了")
		}
	}
}

func BenchmarkAdd2(bb *testing.B) {

	a := 123
	b := 456
	c := 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Println("错了")
		}
	}
}
