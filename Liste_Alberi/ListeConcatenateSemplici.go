package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*1 Implementazione di liste concatenate semplici
Scrivete un programma con l’implementazione di una lista concatenata semplice, seguendo i
lucidi presentati a lezione. Definite i tipi listNode e linkedList e scrivete queste funzioni:
• newNode, che crea un nuovo nodo di lista;
• addNewNode, che inserisce un nuovo nodo in testa alla lista;
• printList, che stampa una lista;
• searchList, che cerca un elemento in una lista;
• removeItem, che cancella un item da una lista.
Per testare le vostre funzioni scrivete una funzione main che gestisca un insieme dinamico
(che variano nel tempo) di interi. Il main deve leggere da standard input una sequenza di istruzioni secondo il formato nella tabella qui sotto, dove n è un intero. I vari elementi sulla riga
sono separati da uno o più spazi. Quando una riga è letta, viene eseguita l’operazione associata;
le operazioni di stampa sono effettuate sullo standard output, e ogni operazione deve iniziare su
una nuova riga.


Istruzione Operazione
in input
+ n Se n non appartiene all’insieme lo inserisce, altrimenti non fa nulla
- n Se n appartiene all’insieme lo elimina, altrimenti non fa nulla
? n Stampa un messaggio che dichiara se n appartiene all’insieme
c Stampa il numero di elementi dell’insieme
p Stampa gli elementi dell’insieme (nell’ordine in cui compaiono nella lista)
o Stampa gli elementi dell’insieme nell’ordine inverso
d Cancella tutti gli elementi dell’insieme
f Termina l’esecuzione
Si assume che l’input sia inserito correttamente. Conviene scrivere le istruzioni di input in un
file in.txt ed eseguire il programma redirigendo lo standard input.*/

type listNode struct {
	data int
	next *listNode
}

type linkedList *listNode

func newNode(value int) linkedList {
	node := new(listNode)
	node.data = value
	node.next = nil
	return node
}

func addNewNode(head linkedList, value int) linkedList {
	if searchList(head, value) == false {
		newNode := newNode(value)
		newNode.next = head
		head = newNode
	}
	return head
}

func printList(head linkedList) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func length(head linkedList) int {
	i := 0
	for head != nil {
		i++
		head = head.next
	}
	return i
}

func searchList(head linkedList, value int) bool {
	for head != nil {
		if head.data == value {
			return true
		}
		head = head.next
	}
	return false
}

func removeItem(head linkedList, value int) linkedList {

	if head == nil {
		return head
	}

	if head.data == value {
		return head.next
	}

	iterator := head
	for iterator.next != nil {
		if iterator.next.data == value {
			iterator.next = iterator.next.next
			break
		}
		iterator = iterator.next
	}
	return head
}

func printListReverse(head linkedList) {
	if head == nil {
		return
	}
	printListReverse(head.next)

	fmt.Println(head.data)
}

func main() {
	//inizializzazione nuovo nodo
	scanner := bufio.NewScanner(os.Stdin)
	var head linkedList = nil

	for scanner.Scan() {
		char := scanner.Text()
		switch char {
		case "+":
			scanner.Scan()
			valore, _ := strconv.Atoi(scanner.Text())
			printList(head)
			head = addNewNode(head, valore)
			printList(head)
		case "-":
			scanner.Scan()
			valore, _ := strconv.Atoi(scanner.Text())
			printList(head)
			head = removeItem(head, valore)
			printList(head)
		case "?":
			scanner.Scan()
			valore, _ := strconv.Atoi(scanner.Text())
			fmt.Printf("%s%t", "Element found:", searchList(head, valore))
		case "c":
			fmt.Println(length(head))
		case "p":
			printList(head)
		case "o":
			printListReverse(head)
		case "d":
			head = nil
		case "f":
			os.Exit(0)
		}
	}
}
