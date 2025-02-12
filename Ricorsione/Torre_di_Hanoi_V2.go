package main

import "fmt"

func hanoi(n int, from int, temp int, to int) int {
	if n == 1 {
		return 1
	} else {
		return hanoi(n-1, from, to, temp) + 1 + hanoi(n-1, temp, from, to)
	}
}

func main() {
	fmt.Println(hanoi(64, 0, 1, 2))
}
