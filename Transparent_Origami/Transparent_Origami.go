package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	fileStream, _ := os.Open("TransparentOrigamiInput.txt")
	defer fileStream.Close() // Defer closing the file

	scanner := bufio.NewScanner(fileStream)
	points := []Point{}

	// SALVATAGGIO PUNTI
	for scanner.Scan() && scanner.Text() != "" { // Riga vuota per dividere punti da folds
		coordinates := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		points = append(points, Point{x, y})
	}

	// FOLD
	for scanner.Scan() {
		fold := strings.Split(scanner.Text(), "=")
		foldCoordinate, _ := strconv.Atoi(fold[1])

		if fold[0][len(fold[0])-1] == 'x' { // Fold along x
			fmt.Println("Fold along x")
			for i, v := range points {
				if v.x > foldCoordinate {
					offset := v.x - foldCoordinate
					points[i] = Point{foldCoordinate - offset, v.y} // Riflette il punto
				}
			}
		} else { // Fold along y
			fmt.Println("Fold along y")
			for i, v := range points {
				if v.y > foldCoordinate {
					offset := v.y - foldCoordinate
					points[i] = Point{v.x, foldCoordinate - offset} // Riflette il punto
				}
			}
		}

		// Rimuovi duplicati
		points = removeDuplicate(points)
	}

	// Stampa punti
	for _, p := range points {
		fmt.Printf("(%d, %d)\n", p.x, p.y)
	}

	fmt.Println(len(points))

	//Stampa matrice con numero
	// Trovo la riga massima (asse y)
	maxLine := 0
	for _, v := range points {
		if v.y > maxLine {
			maxLine = v.y
		}
	}
	// Trovo la colonna massima (asse x)
	maxColumn := 0
	for _, v := range points {
		if v.x > maxColumn {
			maxColumn = v.x
		}
	}

	// Stampa matrice
	for riga := 0; riga <= maxLine; riga++ { // deve essere <= per includere l'ultima riga
		for colonna := 0; colonna <= maxColumn; colonna++ { // deve essere <= per includere l'ultima colonna
			trovato := false
			for _, v := range points {
				if v.x == colonna && v.y == riga { // Cambia asse per disegnare correttamente
					fmt.Print("#")
					trovato = true
					break
				}
			}
			if !trovato {
				fmt.Print(".")
			}
		}
		fmt.Println() // Passa alla riga successiva
	}

}
