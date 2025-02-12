package main

import "fmt"

func hanoi(n int, from int, temp int, to int) {
	if n == 1 {
		fmt.Printf("%d -> %d\n", from, to)
	} else {
		hanoi(n-1, from, to, temp)
		hanoi(1, from, temp, to)
		hanoi(n-1, temp, from, to)
	}
}

func main() {
	hanoi(3, 0, 1, 2)
}
