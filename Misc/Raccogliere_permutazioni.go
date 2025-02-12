/*Data una permutazione di 1..N, vogliamo raccogliere i numeri in ordine crescente cominciando
ad analizzarli da sinistra. Scrivete un programma che stabilisce quante volte avremmo bisogno di
tornare verso sinistra.
Esempio: Nella permutazione 4 5 1 3 6 2 l’output è 2, poiché 1 si trova andando sempre
verso destra, poi si prosegue verso destra per raccogliere 2, ma per raccogliere 3 bisogna tornare
indietro verso sinistra; bisogna tornare ancora indietro per raccogliere 4, dopodiché 5 e 6 si
trovano in ordine proseguendo verso destra.*/

package main

import (
	"fmt"
)

func main() {
	count, indexSucc := 0, 0
	arr := []int{4, 5, 1, 3, 6, 2}

	for i := 0; i < len(arr); i++ {
		indexSucc = len(arr)
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i]+1 {
				indexSucc = j
				break
			}
		}
		if i > indexSucc {
			count++
		}
	}

	fmt.Println(count)

}
