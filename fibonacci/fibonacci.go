package main

import "fmt"

func main() {
	var num = 5

	var result = fibonacci(num)
	fmt.Println(result)
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}

	return fibonacci(num - 1) + fibonacci(num - 2)
}

