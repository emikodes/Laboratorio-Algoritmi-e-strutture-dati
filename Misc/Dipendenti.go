package main

import "fmt"

type azienda struct {
	dipendenti [](*dipendente)
}

type dipendente struct {
	nome         string
	supervisor   *dipendente
	subordinates [](*dipendente) //array di puntatori a dipendente
}

func (dip *dipendente) stampaSubordinati() { //dato un dipendente, stampa l'elenco dei suoi subordinati
	for _, singleSubordinate := range dip.subordinates {
		fmt.Println(singleSubordinate.nome)
	}
}

func (dip *dipendente) visitaSingoloDipendente() int { //è una DFS, incrementa contatore se nodo in esame non ha subordinati
	total := 0
	if dip != nil {
		if len(dip.subordinates) == 0 {
			return 1
		}
		for _, sub := range dip.subordinates {
			total += sub.visitaSingoloDipendente()
		}
	}
	return total
}

func (azienda *azienda) quantiSenzaSubordinati() int {
	total := 0
	for _, v := range azienda.dipendenti {
		total += v.visitaSingoloDipendente()
	}
	return total
}

func (dip *dipendente) supervisore() string {
	if dip.supervisor == nil {
		return "L'impiegato selezionato è di massimo livello: non ha supervisori.\n"
	}
	return dip.supervisor.nome
}

func (dip *dipendente) stampaImpiegatiSopra() {
	for dip.supervisor != nil {
		fmt.Println(dip.supervisor.nome)
		dip = dip.supervisor
	}
}

func (dip *dipendente) DFSSingoloDipendente() {
	if dip != nil {
		if len(dip.subordinates) == 0 {
			fmt.Print(dip.nome)
			return
		}
		for _, v := range dip.subordinates {
			v.DFSSingoloDipendente()
			fmt.Println(" è supervisionato da " + dip.nome)
		}
		if dip.supervisor != nil {
			fmt.Print(dip.nome)
		}
	}
}

func (azienda *azienda) stampaImpiegatiConSupervisore() {
	//DFS
	if azienda != nil {
		for _, v := range azienda.dipendenti {
			v.DFSSingoloDipendente()
			fmt.Println()
		}
	}

}

func nuovoDipendente(nome string, supervisor *dipendente) *dipendente {
	node := new(dipendente)
	node.nome = nome
	node.supervisor = supervisor
	return node
}

func main() {
	algore := azienda{}

	algore.dipendenti = append(algore.dipendenti, nuovoDipendente("Anna", nil))
	algore.dipendenti = append(algore.dipendenti, nuovoDipendente("Gianni", nil))
	algore.dipendenti[1].subordinates = append(algore.dipendenti[1].subordinates, nuovoDipendente("Harry", algore.dipendenti[1])) //attacco harry a gianni

	algore.dipendenti[0].subordinates = append(algore.dipendenti[0].subordinates, nuovoDipendente("Bruno", algore.dipendenti[0]))   //attacco bruno ad anna
	algore.dipendenti[0].subordinates = append(algore.dipendenti[0].subordinates, nuovoDipendente("Carlo", algore.dipendenti[0]))   //attacco carlo ad anna
	algore.dipendenti[0].subordinates = append(algore.dipendenti[0].subordinates, nuovoDipendente("Daniela", algore.dipendenti[0])) //attacco daniela ad anna

	algore.dipendenti[0].subordinates[0].subordinates = append(algore.dipendenti[0].subordinates[0].subordinates, nuovoDipendente("Enrico", algore.dipendenti[0].subordinates[0]))    //attacco enrico a bruno
	algore.dipendenti[0].subordinates[0].subordinates = append(algore.dipendenti[0].subordinates[0].subordinates, nuovoDipendente("Francesco", algore.dipendenti[0].subordinates[0])) //attacco francesco a bruno

	algore.dipendenti[0].subordinates[0].subordinates[1].subordinates = append(algore.dipendenti[0].subordinates[0].subordinates[1].subordinates, nuovoDipendente("Irene", algore.dipendenti[0].subordinates[0].subordinates[1])) //attacco irene a francesco

	algore.dipendenti[0].stampaSubordinati() //subordinati di Anna, stampa Bruno,Carlo,Daniela
	fmt.Println()
	fmt.Println(algore.quantiSenzaSubordinati()) //stampa 5
	fmt.Println()
	fmt.Println(algore.dipendenti[0].supervisore()) //stampa "non ha supervisori"
	algore.stampaImpiegatiConSupervisore()          //stampa intero albero
}
