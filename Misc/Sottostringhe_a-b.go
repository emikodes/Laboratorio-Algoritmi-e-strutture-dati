/*Data una stringa di N caratteri nell’alfabeto a b c, stampare il numero di sottostringhe che
iniziano con a e finiscono con b (tali sottostringhe possono sovrapporsi).
Esempio: nella stringa ccbaacbabbcbab il numero di sottostringhe a-b è 15. Notate che
ciascuna delle prime due a (cioè, le due più a sinistra) appaiono ciascuna in 5 sottostringhe
a-b).

ccbaacbabbcbab
b=6
a=4

Numero combinazioni senza ripetizioni: B! /A! * (B-A)! Con B>A

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	stringa := "ccbaacbabbcbab"
	nSottostringhe := 0

	for i, v := range stringa {
		if v == 'a' {
			nSottostringhe += strings.Count(stringa[i:], "b")
		}
	}

	fmt.Println(nSottostringhe)

}
