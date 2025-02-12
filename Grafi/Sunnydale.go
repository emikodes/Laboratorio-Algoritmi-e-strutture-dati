package main

import (
	"fmt"
	"sort"
)

/*INPUT:
N M H S (N=Numero nodi, M=Numero archi, H=Nodo Start, S=Nodo End)
.
.
.
.
.
(M righe)

Struttura della singola riga: A B C, dove A B sono nodi connessi da un arco di peso C
Output: Numero di gallerie necessarie x passare da Start a End, -1 se durante la costruzione del percorso, passo per un nodo già visitato

Criterio di scelta del prossimo arco da percorrere: arco a peso minore (greedy)

Modello un grafo non orientato, con N nodi e M archi.
Partendo dal nodo H, scelgo per ogni nodo l'arco di peso minore, e lo percorro.
Stop quando ho raggiunto nodo S, oppure se finisco su un nodo già visitato*/

type nodo struct {
	valore int
	peso   int
}

type listaAdiacenti []nodo

type grafo struct {
	n    int
	nodi []listaAdiacenti
}

func nuovoGrafo(n int) *grafo {
	grafo := new(grafo)
	grafo.n = n
	grafo.nodi = make([]listaAdiacenti, n)
	return grafo
}

func leggiArchi(numeroNodi int, numeroArchi int) *grafo {
	from, to, peso := 0, 0, 0
	grafo := nuovoGrafo(numeroNodi)
	for i := 0; i < numeroArchi; i++ {
		fmt.Scan(&from, &to, &peso)
		//inserisco gli archi due volte, in "entrambe" le direzioni,
		//nelle liste del nodo di partenza e del nodo di arrivo
		grafo.nodi[from-1] = append(grafo.nodi[from-1], nodo{to - 1, peso})
		grafo.nodi[to-1] = append(grafo.nodi[to-1], nodo{from - 1, peso})
	}
	return grafo
}

func numeroGallerie(grafoDaVisitare *grafo, Harmony int, Sarah int) int {
	current, gallerie, prev := Harmony, 0, -1
	for current != Sarah {
		//ordino la slice dei nodi adiacenti, considerando come "criterio", il peso degli archi.
		sort.Slice(grafoDaVisitare.nodi[current], func(i, j int) bool {
			return grafoDaVisitare.nodi[current][i].peso < grafoDaVisitare.nodi[current][j].peso
		})

		//se non sono tornato indietro, proseguo
		if prev != grafoDaVisitare.nodi[current][0].valore {
			prev = current
			current = grafoDaVisitare.nodi[current][0].valore
			gallerie++
		} else { //altrimenti, Sarah non è raggiungibile col criterio stabilito (scelgo per ogni nodo l'arco di peso minore)
			return -1
		}
	}
	return gallerie
}

func main() {
	N, M, H, S := 0, 0, 0, 0
	fmt.Scan(&N, &M, &H, &S)
	grafo := leggiArchi(N, M)
	fmt.Println(numeroGallerie(grafo, H-1, S-1))
}
