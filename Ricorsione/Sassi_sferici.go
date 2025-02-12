package main

import "fmt"

func sassi(height int) int {
	if height == 1 {
		return 1
	} else {
		return (height * height) + sassi(height-1)
	}
}

func main() {
	fmt.Println(sassi(1))
}
