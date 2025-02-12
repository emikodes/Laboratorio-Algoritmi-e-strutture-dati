package main

import (
	"fmt"
)

/*Insieme di liste di adiacenza (linkedList)
0[] -> 1[] -> 3[]
1[] -> 2[]
2[] -> 3[]
3[] -> 0[] -> 4[]
4[] -> 2[]
*/

type graph struct {
	n     int //numero di vertici
	nodes []linkedList
}

//Singolo nodo della lista

type listNode struct {
	item int
	next *listNode
}

// Inizio della lista (alias di head)
type linkedList struct {
	head *listNode
}

/*Restituisce un nuovo grafo vuoto, composto da n nodi*/
func nuovoGrafo(n int) *graph {
	graph := new(graph)
	graph.n = n
	graph.nodes = make([]linkedList, n)
	return graph
}

func inserisciInTesta(list *listNode, to int) *listNode {
	newNode := new(listNode)
	newNode.item = to
	newNode.next = list
	list = newNode
	return list
}

/*Crea un nuovo grafo composto da n nodi, legge una coppia di interi del tipo 0 1, rappresentante un arco che va dal nodo 0 al nodo 1*/
func leggiGrafo(n int) *graph {
	graph := nuovoGrafo(n)
	from, to := 0, 0
	for {
		val, _ := fmt.Scan(&from, &to)
		if val != 2 {
			break
		} else {
			graph.nodes[from].head = inserisciInTesta(graph.nodes[from].head, to)
		}
	}

	return graph
}

func stampaGrafo(grafo *graph) {
	for i := 0; i < grafo.n; i++ {
		iterator := grafo.nodes[i].head
		fmt.Print("Archi uscenti dal nodo: " + string(i+48) + ": ")
		for iterator != nil {
			fmt.Print(string(iterator.item+48) + " ")
			iterator = iterator.next
		}
		fmt.Println()
	}
}

func cercaInLista(element int, head *listNode) bool {
	iterator := head
	for iterator != nil {
		if iterator.item == element {
			return true
		}
		iterator = iterator.next
	}
	return false
}

func esisteArco(x int, y int, grafo *graph) bool {
	return cercaInLista(y, grafo.nodes[x].head)
}

func main() {
	n := 0
	fmt.Println("Inserisci numero di nodi: ")
	fmt.Scan(&n)
	grafo := leggiGrafo(n)
	stampaGrafo(grafo)

	//esiste arco
	fmt.Printf("Esiste arco tra 3 e 4? : %t\n", esisteArco(3, 4, grafo)) //%t is the printing verb used to print boolean types
	//esiste arco
	fmt.Printf("Esiste arco tra 4 e 2? : %t\n", esisteArco(4, 2, grafo))
	//esiste arco
	fmt.Printf("Esiste arco tra 4 e 2? : %t\n", esisteArco(0, 4, grafo))
}
