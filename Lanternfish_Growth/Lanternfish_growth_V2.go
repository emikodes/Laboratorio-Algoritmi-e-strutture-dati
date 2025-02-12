package main

import "fmt"

func main() {
	input := []int{5, 1, 1, 3, 1, 1, 5, 1, 2, 1, 5, 2, 5, 1, 1, 1, 4, 1, 1, 5, 1, 1, 4, 1, 1, 1, 3, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4, 4, 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 5, 1, 1, 1, 4, 1, 1, 1, 1, 1, 3, 1, 1, 4, 1, 4, 1, 1, 2, 3, 1, 1, 1, 1, 4, 1, 2, 2, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 4, 4, 1, 4, 2, 1, 1, 1, 1, 1, 4, 3, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 3, 1, 1, 1, 2, 1, 1, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 3, 1, 2, 1, 1, 4, 1, 1, 5, 3, 1, 1, 1, 2, 4, 1, 1, 2, 4, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 4, 3, 1, 2, 1, 2, 1, 5, 1, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 3, 1, 1, 5, 1, 1, 1, 1, 5, 1, 4, 1, 1, 1, 4, 1, 3, 4, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 5, 1, 3, 1, 1, 1, 1, 4, 1, 5, 3, 1, 1, 1, 1, 1, 5, 1, 1, 1, 2, 2}
	days, count := 256, 0
	stadiPesci := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, v := range input {
		stadiPesci[v]++
	}

	for i := 0; i < days; i++ {
		stadiPesci = append(stadiPesci[1:], stadiPesci[0])
		stadiPesci[6] += stadiPesci[8]
	}

	for _, v := range stadiPesci {
		count += v
	}
	fmt.Println(count)

}

/*
Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,5,6,7,8
After  4 days: 6,0,6,4,5,6,7,8,8
After  5 days: 5,6,5,3,4,5,6,7,7,8
After  6 days: 4,5,4,2,3,4,5,6,6,7
After  7 days: 3,4,3,1,2,3,4,5,5,6
After  8 days: 2,3,2,0,1,2,3,4,4,5
After  9 days: 1,2,1,6,0,1,2,3,3,4,8

Each day, a 0 becomes a 6 and adds a new 8 to the end of the list,
while each other number decreases by 1 if it was present at the start of the day.

011210000
112100000
121000101
210001111
100011312
000113221
001132210

Salvo il valore della prima cifra, shifto a destra di uno, aggiungo il valore della prima cifra alla fine, aggiungo il valore della prima cifra al count dei valori a 6.
*/
