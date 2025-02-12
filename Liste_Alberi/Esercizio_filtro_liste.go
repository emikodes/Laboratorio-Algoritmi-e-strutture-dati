1) Analizzate la seguente funzione ricorsiva, che ha per argomenti il puntatore p al primo nodo
di una lista concatenata e un intero k:
func f( p *node , k int ) int {
var a int
if p == nil {
return 0
}
a = 1 + f(p.next , k)
if a == k {
fmt.Println(p.val)
}
return a
}
Rispondete alle seguenti domande, dove p punta al nodo con valore 3 della lista rappresentata
qui sotto:
• Cosa stampa la chiamata f(p,k) se k vale 1? Stampa 7
• Cosa stampa la chiamata f(p,k) se k vale 5? Stampa 3
• Cosa stampa la chiamata f(p,k) se k vale 10? Non stampa nulla
• Completate la seguente frase: “Ricevendo il puntatore alla testa di una lista e un intero k la funzione f stampa il valore in posizione k a partire dalla fine della lista, e restituisce la lunghezza della lista, partendo dal nodo puntato da p, fino alla fine (nil). Se K è maggiore della lunghezza della lista, o negativo, non stampa nulla.”. Includete nella descrizione anche i casi
particolari/limite, se ve ne sono di rilevanti.
• Qual è la complessità della funzione? La funzione scorre tutta la lista, fino a trovare nil, quindi la complessità è O(n), dove n è la lunghezza della lista a partire dal nodo puntato da p.

2) Definite un tipo strutturato node compatibile con la funzione f qui sopra.
type node struct{
    next *node
    val int
}

3) Definite un nuovo tipo circNode che serva a a rappresentare i nodi di una lista circolare
contenente numeri interi. Una lista circolare è un tipo particolare di lista doppiamente concatenata, in cui l’ultimo nodo della lista e il primo sono collegati in modo da formare un
circolo.

type circNode struct{
    prev *circNode
    next *circNode
    val int
    //la gestione dei collegamenti circolari, avviene tramite funzioni apposite
}

4) Scrivete una funzione stampaDaZero(p *circNode) che stampi una lista circolare partendo
dal nodo di valore 0 e proseguendo circolarmente. Potete assumere che la lista circolare di
cui fa parte p contenga esattamente un nodo con valore 0. Ad esempio, se p punta ad un
nodo qualunque della lista circolare disegnata sopra, allora la funzione stampaDaZero deve
stampare i numeri 0 1 7 3 -2.

func stampaDaZero(p *circNode){
    for p.val !=0{
        p=p.next
    }
    //ho trovato lo 0
    fmt.Println(p.val)
    p=p.next
    for p.val!=0{
        fmt.Println(p.val)
        p=p.next
    }
    
}

5) Scrivete una funzione main che legge una serie di interi (uno per riga), crea una lista circolare
contenente questi interi, e la stampi invocando la funzione stampaDaZero. Potete assumere
che il vettore contenga esattamente uno zero.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type circNode struct {
    val  int
    next *circNode
}

func stampaDaZero(p *circNode) {
    // Trova il nodo con valore 0
    for p.val != 0 {
        p = p.next
    }
    // Stampa il nodo con valore 0
    fmt.Println(p.val)
    p = p.next
    // Stampa i valori successivi
    for p.val != 0 {
        fmt.Println(p.val)
        p = p.next
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var head *circNode = nil
    var lastNode *circNode = nil // Per tenere traccia dell'ultimo nodo

    // Legge i numeri interi finché non viene inserito uno zero
    for {
        scanner.Scan()
        value, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Errore nella conversione dell'input")
            return
        }

        // Crea un nuovo nodo
        tmp := &circNode{val: value}

        // Collega il nuovo nodo alla lista
        if head == nil {
            head = tmp // Se è il primo nodo, diventa head
        } else {
            lastNode.next = tmp // Collega l'ultimo nodo al nuovo nodo
        }

        lastNode = tmp // Aggiorna lastNode

        // Esci se il valore è 0
        if value == 0 {
            break
        }
    }

    // Completa la lista circolare
    if lastNode != nil {
        lastNode.next = head // L'ultimo nodo punta al primo nodo per formare il cerchio
    }

    // Stampa i valori della lista circolare partendo da 0
    stampaDaZero(head)
}
