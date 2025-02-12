package main

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// Salva l'elemento corrente
		key := arr[i]
		j := i - 1

		// Sposta gli elementi a destra finché sono maggiori di "key"
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Inserisci "key" nella posizione corretta
		arr[j+1] = key
	}
}

func main() {
	arr := []int{14, 12, 3, 4, 43, 11}
	InsertionSort(arr)
	fmt.Println(arr)
}
