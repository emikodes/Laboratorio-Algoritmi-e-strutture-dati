package main

import "fmt"

func merge(first []int, second []int) []int {
	merged := make([]int, len(first)+len(second)) //slice di dimensione definita a run-time.
	firstIndex, secondIndex, iterator := 0, 0, 0

	for firstIndex < len(first) && secondIndex < len(second) {
		if first[firstIndex] <= second[secondIndex] {
			merged[iterator] = first[firstIndex]
			firstIndex++
		} else {
			merged[iterator] = second[secondIndex]
			secondIndex++
		}
		iterator++
	}

	//se uno delle due slice non è finito
	if firstIndex < len(first) {
		for firstIndex < len(first) {
			merged[iterator] = first[firstIndex]
			firstIndex++
			iterator++
		}
	} else if secondIndex < len(second) {
		for secondIndex < len(second) {
			merged[iterator] = second[secondIndex]
			secondIndex++
			iterator++
		}
	}
	return merged
}

func mergeSort(slice []int) {
	if len(slice) > 1 {
		mid := len(slice) / 2
		mergeSort(slice[:mid])
		mergeSort(slice[mid:])
		copy(slice, merge(slice[:mid], slice[mid:]))
		/*slice = merge(slice[:mid], slice[mid:])
		La riga sopra non avrebbe funzionato, perchè in GoLang,
		tutto è passato per copia. Anche le slices,
		ma sappiamo che le slices in realtà sono strutture di 3 valori,
		un puntatore ad un array sottostante, len e cap.

		Alla funzione, è passata una copia di questa struttura, quindi un nuovo puntatore,
		che punta però allo stesso array sottostante. con slice= merge(),
		sto semplicemente puntando "slice", ad un array sottostante diverso,
		ma "slice" è la copia locale del puntatore, quindi la modifica non viene riflettuta nel chiamante!!
		*/
	}
}

func main() {
	arr := []int{14, 12, 3, 4, 43, 11}
	mergeSort(arr)
	fmt.Println(arr)
}
