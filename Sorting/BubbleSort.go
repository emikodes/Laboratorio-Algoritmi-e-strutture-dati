package main

import "fmt"

func BubbleSort(arr []int) {
	scambiato := true
	for scambiato {
		scambiato = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				scambiato = true
			}
		}
	}
}

func main() {
	arr := []int{14, 12, 3, 4, 43, 11}
	BubbleSort(arr)
	fmt.Println(arr)
}
