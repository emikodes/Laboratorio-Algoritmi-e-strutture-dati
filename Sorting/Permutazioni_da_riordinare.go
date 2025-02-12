package main

import "fmt"

type struttura struct {
	name string
	key  int
}

func main() {
	arr := []struttura{
		{name: "Andrea", key: 1},
		{name: "Beatrice", key: 2},
		{name: "Francesco", key: 6},
		{name: "Elisa", key: 5},
		{name: "Henry", key: 8},
		{name: "Irene", key: 9},
		{name: "Dino", key: 4},
		{name: "Carlo", key: 3},
		{name: "Giorgia", key: 7},
	}

	for i := 0; i < len(arr); i++ {
		if i != len(arr)-arr[i].key {
			arr[len(arr)-arr[i].key], arr[i] = arr[i], arr[len(arr)-arr[i].key]
		}
	}

	fmt.Println(arr)

}
