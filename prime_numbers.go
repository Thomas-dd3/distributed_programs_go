package main

import (
    "fmt"
	"flag"
	//"time"
	"strconv"
)

func filter(k int, c <-chan int, r chan<- int) {
	next := false
	canalNext := make(chan int)
	for {
		num := <- c
		// si on reçoit 0, on envoie le résultat (valeur de sa go routine) dans le canal
		if num == 0 {
			r <- k
			if next {
				canalNext <- num
			}else{
				// FIN de l'algo et de la recherche des nombres premiers
				r <- 0
			}
			return
		}
		// si le modulo est différent de 0 alors on n'est pas un multiple
		if num%k != 0 {
			// si on a déjà un canal actif vers le filtre suivant on transmet la valeur
			if next {
				canalNext <- num
			} else {
				// lancement du filtre suivant
				next = true
				go filter(num, canalNext, r)
			}
		}		
	}
}

func main() {
	// Récupération de la valeur entrée dans la console
	p_num := flag.Int("n", 5, "nombre")
	flag.Parse()

	if ( *p_num < 3 ) {
		fmt.Print("Veuillez entrer un nombre supérieur ou égale à 3\n")
		return
	}

	canal := make(chan int)
	resultat := make(chan int)
	// Lancement du 1er filtre
	go filter(2, canal, resultat)

	// Envoie des nombres jusque p_num dans le canal
	for i:=3 ; i<=*p_num ; i++ {
		canal <- i
	}
	
	// Demande des résulats, peut être fait sans attendre car le canal est FIFO
	canal <- 0
	resultStr := ""

	// Récupération des résulats
	/*
	// Version 1
	go func(){
		result
		for {
			result := <- resultat
			resultStr += strconv.Itoa(result) + ", "
		}
	}()
	// On laisse le temps à la go routine de récupérer tous les résultats, 
	// plus le nombre entré dans la console est grand et plus il faudra laisser de temps
	//time.Sleep(1 * time.Second)
	
	*/
	
	// Version 2 - plus rapide
	result := -1
	for result != 0 {
		result = <- resultat
		resultStr += strconv.Itoa(result) + ", "
	}

	
	// Affichage des résultats
	resultStr = resultStr[:len(resultStr)-5] //On retire le 0
	fmt.Print("Liste des nombres premiers jusqu'à ", *p_num, " : ", resultStr, ".\n")

}