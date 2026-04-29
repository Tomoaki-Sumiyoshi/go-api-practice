package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(IsEven(11))
	fmt.Println(IsEven(4))
	fmt.Println(Max(2, 1))
	fmt.Println(Max(10, 10))
	fmt.Println(Max(1, 5))
}
