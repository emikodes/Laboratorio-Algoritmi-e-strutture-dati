package main

import "fmt"

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}

func main() {
	numbers := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
	fmt.Println(largest(numbers))
}

/*
Stack di esecuzione ricorsivo:
{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
{1, 2, 5, 7, -2, 10, 9, 21, 3}
{1, 2, 5, 7, -2, 10, 9, 21}
{1, 2, 5, 7, -2, 10, 9}
{1, 2, 5, 7, -2, 10}
{1, 2, 5, 7, -2}
{1, 2, 5, 7}
{1, 2, 5}
{1, 2}
{1} //restituisce 1

La funzione max, viene eseguita con 1 e 2 come argomenti.
L'ultima volta, viene eseguita con 21 e 8.
*/
