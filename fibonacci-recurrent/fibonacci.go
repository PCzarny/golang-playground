package main

import "fmt"

func fib(a int) int {
	if a == 1 || a == 0 {
		return 1
	}
	return fib(a - 1) + fib(a - 2)
}

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x, fib(x))
	}
}