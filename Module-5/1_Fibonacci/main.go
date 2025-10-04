package main

import "fmt"

func main() {
	result := fibonacci(9)
	fmt.Println(result)
}

func fibonacci(n int) int {
	r := make([]int, n+2)
	r[0] = 0
	r[1] = 1

	for i := 2; i <= n; i++ {
		r[i] = r[i-1] + r[i-2]
	}
	return r[n]
}
