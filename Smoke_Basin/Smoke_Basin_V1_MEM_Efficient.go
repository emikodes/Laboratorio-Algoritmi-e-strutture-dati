package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readIntegerLine(reader *bufio.Reader) []int {
	line := []int{}
	for {
		runa, _, err := reader.ReadRune()
		if runa != 13 && err == nil { //13 is CR
			integer, _ := strconv.Atoi(string(runa))
			line = append(line, integer)
		} else {
			_, _, _ = reader.ReadRune() //consume LF
			break
		}
	}
	return line

}

func main() {
	fileStream, _ := os.Open("Input.txt")
	reader := bufio.NewReader(fileStream)
	matrix := [3][]int{}
	lowPoints := []int{}
	first := true

	//read first three matrix
	for i := 0; i < 3; i++ {
		matrix[i] = readIntegerLine(reader)
	}

	//computazione
	for {
		for i := 0; i < 2; i++ {
			if first != true {
				i = 1
			}
			for j := 0; j < len(matrix[i]); j++ {
				if j == 0 {
					//se è un angolo SX, e sotto e dx sono minori di esso
					if i == 0 && matrix[i+1][j] > matrix[i][j] && matrix[i][j+1] > matrix[i][j] {
						lowPoints = append(lowPoints, matrix[i][j])

					} else if i != 0 && matrix[i+1][j] > matrix[i][j] && matrix[i-1][j] > matrix[i][j] && matrix[i][j+1] > matrix[i][j] {
						//primo ma non della prima riga, controllo sopra sotto e dx
						lowPoints = append(lowPoints, matrix[i][j])
					}
				} else if j == len(matrix[i])-1 {
					//se è un angolo DX, e sotto e sx sono minori di esso
					if i == 0 && matrix[i+1][j] > matrix[i][j] && matrix[i][j-1] > matrix[i][j] {
						lowPoints = append(lowPoints, matrix[i][j])
					} else if matrix[i+1][j] > matrix[i][j] && matrix[i-1][j] > matrix[i][j] && matrix[i][j-1] > matrix[i][j] { //uktimo ma non della prima riga, controllo sopra sotto e sx
						lowPoints = append(lowPoints, matrix[i][j])
					}
				} else if i == 0 {
					if matrix[i+1][j] > matrix[i][j] && matrix[i][j-1] > matrix[i][j] && matrix[i][j+1] > matrix[i][j] {
						//è un nodo cornice superiore, controllo sx,dx,sotto
						lowPoints = append(lowPoints, matrix[i][j])
					}
				} else if matrix[i+1][j] > matrix[i][j] && matrix[i-1][j] > matrix[i][j] && matrix[i][j-1] > matrix[i][j] && matrix[i][j+1] > matrix[i][j] { //mid, controllo sopra sotto dx e sx
					lowPoints = append(lowPoints, matrix[i][j])
				}
			}
		}
		tmp := readIntegerLine(reader)
		if len(tmp) == 0 {
			break
		} else {
			matrix[0] = matrix[1]
			matrix[1] = matrix[2]
			matrix[2] = tmp
			first = false
		}
	}
	//controllo ultima riga

	for i := 0; i < len(matrix[2]); i++ {
		//se è un angolo SX, e sotto e dx sono minori di esso
		if i == 0 {
			if matrix[1][i] > matrix[2][i] && matrix[2][i+1] > matrix[2][i] {
				lowPoints = append(lowPoints, matrix[2][i])
			}
		} else if i == len(matrix[2])-1 {
			if matrix[1][i] > matrix[2][i] && matrix[2][i-1] > matrix[2][i] {
				lowPoints = append(lowPoints, matrix[2][i])
				//angolo dx
			}
		} else if matrix[1][i] > matrix[2][i] && matrix[2][i+1] > matrix[2][i] && matrix[2][i-1] > matrix[2][i] {
			lowPoints = append(lowPoints, matrix[2][i])
		} //mid
	}

	fmt.Println(lowPoints)

	total := 0
	for _, v := range lowPoints {
		total += v + 1
	}
	fmt.Println(total)
}
