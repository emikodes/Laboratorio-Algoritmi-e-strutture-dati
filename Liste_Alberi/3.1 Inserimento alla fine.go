Considerate la seguente funzione (incompleta) che aggiunge un elemento e in fondo alla lista:
    1 func addNewNodeAtEnd (l * linkedListWithTail , val int) {
    2   if l. tail == nil {
    3     l. tail = newNode ( val )
    4     l. head = l. tail
    5   } else {
    6     // MISSING CODE
    7   }
    8 }

    [][][][][][tail]nil
Quale dei seguenti frammenti di codice completa la funzione addAtEnd? Spiegate quali
problemi si riscontrano scegliendo ciascuna delle altre opzioni.

A) l. tail . next = val
l. tail = val

Assegna a l.tail.next (puntatore a nodo), un intero. type mismatch

B) temp := newNode ( val )
l. tail . next = temp

Corretto ma manca un pezzo, non aggiorna valore della tail

C) temp := newNode ( val )
l. tail = temp

Sostituisce la coda, perdo dati

D) l. tail . next = newNode ( val )
l. tail = l. tail . next

Corretto

E) temp := l. head
for temp . next != nil {
temp = temp . next
}
temp . next = newNode ( val )

Corretto, ma inefficiente, e non aggiorna la tail


La risposta esatta è D