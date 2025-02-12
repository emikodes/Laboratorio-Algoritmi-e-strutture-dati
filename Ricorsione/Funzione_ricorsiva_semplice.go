package main

import "fmt"

func multiply(x, y int) int {
	if y == 0 {
		return 0
	} else {
		return x + multiply(x, y-1)
	}
}

func main() {
	fmt.Println(multiply(5, 0))
}
