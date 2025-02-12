package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Point struct {
	x int
	y int
}

const righe = 100
const colonne = 100

func computeBasin(lowPoint Point, matrix [][]int) int {
	spawned := 1
	queue := []Point{lowPoint}
	iterator := 0
	for len(queue) != 0 && iterator < len(queue) {
		toVisit := queue[iterator]
		iterator++

		// Check all four neighbors: up, down, left, right
		if toVisit.x > 0 && matrix[toVisit.x-1][toVisit.y] > matrix[toVisit.x][toVisit.y] { // up is greater
			//se non è già nella coda, lo metto
			if !slices.Contains(queue, Point{toVisit.x - 1, toVisit.y}) && matrix[toVisit.x-1][toVisit.y] != 9 {
				queue = append(queue, Point{toVisit.x - 1, toVisit.y})
				spawned++
			}
		}
		if toVisit.x < righe-1 && matrix[toVisit.x+1][toVisit.y] > matrix[toVisit.x][toVisit.y] { // down is greater

			if !slices.Contains(queue, Point{toVisit.x + 1, toVisit.y}) && matrix[toVisit.x+1][toVisit.y] != 9 {
				queue = append(queue, Point{toVisit.x + 1, toVisit.y})
				spawned++
			}
		}
		if toVisit.y > 0 && matrix[toVisit.x][toVisit.y-1] > matrix[toVisit.x][toVisit.y] { // left is greater

			if !slices.Contains(queue, Point{toVisit.x, toVisit.y - 1}) && matrix[toVisit.x][toVisit.y-1] != 9 {
				queue = append(queue, Point{toVisit.x, toVisit.y - 1})
				spawned++
			}

		}
		if toVisit.y < colonne-1 && matrix[toVisit.x][toVisit.y+1] > matrix[toVisit.x][toVisit.y] { // right is greater
			if !slices.Contains(queue, Point{toVisit.x, toVisit.y + 1}) && matrix[toVisit.x][toVisit.y+1] != 9 {
				queue = append(queue, Point{toVisit.x, toVisit.y + 1})
				spawned++
			}

		}
	}

	return spawned
}

// Funzione per leggere una riga di numeri interi
func readIntegerLine(reader *bufio.Reader) []int {
	line := []int{}
	for {
		runa, _, err := reader.ReadRune()
		if err != nil || runa == '\n' {
			break
		}
		if runa != '\r' { // Ignora il carattere CR
			num, _ := strconv.Atoi(string(runa))
			line = append(line, num)
		}
	}
	return line
}

func main() {
	// Apertura del file di input
	basinDim := []int{}
	fileStream, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println("Errore nell'aprire il file:", err)
		return
	}
	defer fileStream.Close()

	// Inizializzazione lettura matrice
	reader := bufio.NewReader(fileStream)
	matrix := make([][]int, 100)
	for i := 0; i < 100; i++ {
		matrix[i] = readIntegerLine(reader)
	}

	// Array per salvare i punti bassi
	lowPoints := []Point{}

	// Computazione per identificare i punti bassi
	for i := 0; i < righe; i++ {
		for j := 0; j < colonne; j++ {
			current := matrix[i][j]
			isLow := true

			// Controllo dei vicini
			if i > 0 && matrix[i-1][j] <= current { // sopra
				isLow = false
			}
			if i < righe-1 && matrix[i+1][j] <= current { // sotto
				isLow = false
			}
			if j > 0 && matrix[i][j-1] <= current { // sinistra
				isLow = false
			}
			if j < colonne-1 && matrix[i][j+1] <= current { // destra
				isLow = false
			}

			// Aggiungi il punto basso alla lista
			if isLow {
				lowPoints = append(lowPoints, Point{i, j})
			}
		}
	}

	// Stampa i punti bassi e calcola il totale
	fmt.Println("Punti bassi:", lowPoints)

	total := 0
	for _, point := range lowPoints {
		total += matrix[point.x][point.y] + 1
	}
	fmt.Println("Totale:", total)

	for _, v := range lowPoints {
		basinDim = append(basinDim, computeBasin(v, matrix))
	}

	slices.Sort(basinDim)
	slices.Reverse(basinDim)

	fmt.Println(basinDim)
	fmt.Println(basinDim[0] * basinDim[1] * basinDim[2])
}
