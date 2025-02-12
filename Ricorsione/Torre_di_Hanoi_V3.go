package main

import "fmt"

func hanoi(n int, from *string, temp *string, to *string) {
	if n == 1 {
		*to = *to + string((*from)[len(*from)-1])
		*from = string([]rune(*from)[:len(*from)-1])
		fmt.Println(*from + "," + *temp + "," + *to)
	} else {
		hanoi(n-1, from, to, temp)
		hanoi(1, from, temp, to)
		hanoi(n-1, temp, from, to)
	}
}

func main() {
	from := "ABC"
	to := ""
	temp := ""
	hanoi(3, &from, &temp, &to)
}
