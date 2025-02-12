package main

import "fmt"

/*
INPUT:
R C
int int
int int
.
.
.
.
.
Dove R e C sono il numero di righe e colonne della griglia
Le successive righe, contengono due interi, che possono rappresentare
La bocca e la coda di un serpente (primo int > secondo int),
oppure la base e la cima di una scala (primo int < secondo int)

Output:
Numero minimo di mosse necessarie per raggiungere la casella n=R*C partendo dalla casella 1.
Una sequenza di lanci di lunghezza minima, che permettano di vincere la partita.
Numero minimo di mosse necessarie per raggiungere la casella n, partendo da una casella presa in input.
Numero minimo di mosse necessarie per raggiungere la casella n, partendo da una casella presa in input, senza utilizzare ne scale ne serpenti.

Esempio input:
5 6
27 1
21 9
17 4
19 7
11 26
3 22
20 29
5 8

Output:
3
2 6 2 (Potrebbe differire, ci sono diverse sequenze di lanci di lunghezza 3 che permettano di vincere la partita)

*/

func makeGrid(R int, C int) []int {
	gameGrid := make([]int, R*C) //vettore di lunghezza R*C, initializzato con tutti 0
	start, end := 0, 0
	for {
		nRead, _ := fmt.Scan(&start, &end)
		if nRead == 2 {
			gameGrid[start-1] = end //aggiungo al vettore, la scala/serpente. (In posizione della base/bocca, inserisco la coda/cima della scala, rappresentando il "balzo" della scala/serpente)
		} else {
			break
		}
	}
	return gameGrid
}

func playGame(gameGrid []int) []int {
	arrived := false
	startIndex := 0
	moves := []int{}
	arriveToBestMove := 0
	for !arrived {
		bestMove := 6
		//partendo da una casella,cerco nelle 6 successive la mossa migliore che potrei fare
		//questa mossa, deve essere maggiore di 6, cioè deve essere un balzo maggiore (una scala), altrimenti gioco 6
		for i := 0; i < 6 && startIndex+i < len(gameGrid); i++ {
			if gameGrid[startIndex+i] > bestMove {
				bestMove = gameGrid[startIndex+i]
				arriveToBestMove = i
			}
		}
		if bestMove > 6 && bestMove-1 < len(gameGrid) { //controllo se la scala mi fa "sforare" la fine della griglia
			moves = append(moves, arriveToBestMove) //una mossa per arrivare alla scala
			startIndex = bestMove - 1               //"percorro la scala" la mossa migliore
		} else if startIndex+bestMove < len(gameGrid) { //se giocando 6 non "sforo" la griglia
			moves = append(moves, 6)
			startIndex += bestMove //gioco +6
		} else { //se con +6 arrivo a n oppure sforo, vuol dire che ho vinto.
			moves = append(moves, len(gameGrid)-startIndex-1) //l'ultima mossa è il "rimanente" per vincere.
			arrived = true
		}
	}
	return moves
}

func main() {
	R, C := 0, 0
	fmt.Scan(&R, &C)
	gameGrid := makeGrid(R, C)
	winnerCombination := playGame(gameGrid)
	fmt.Printf("Numero mosse per vincere: %d\nMosse vincenti: %d", len(winnerCombination), winnerCombination)
}
