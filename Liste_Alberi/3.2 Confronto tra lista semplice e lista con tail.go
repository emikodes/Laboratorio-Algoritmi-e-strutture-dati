3.2 Confronto tra lista semplice e lista con tail
Nella lista concatenata linkedListWithTail abbiamo i riferimenti all’inizio e alla fine della lista.
Confrontate questa implementazione con quella di una lista semplice che non ha un
riferimento alla fine della lista, cioè definita come segue:

type linkedList struct {
    head * listNode
}

Per quali delle seguenti operazioni si ha un tempo migliore con linkedListWithTail piuttosto
che con linkedList? Scegliete tutte le opzioni corrette e giustificate la risposta.
A) restituisci 1 se la lista contiene un dato elemento //Il tempo rimane uguale, O(n) (Non posso iterare dalla coda, non è doppiamente linkata)
B) cancella l’ultimo elemento della lista //O(1) vs O(n)
C) aggiungi un elemento all’inizio della lista //O(1) per entrambe
D) aggiungi un elemento alla fine della lista //O(1) vs O(n)