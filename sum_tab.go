package main

import (
    "fmt"
	"flag"
	"runtime"
)

func remplirTableau(tab []int){
    for i:=0 ; i < len(tab) ; i++ {
        tab[i] = i
    }
}

func afficherTableau(tab []int) {
    for index, valeur := range (tab) {
        fmt.Print(" tabint[", index, "] = ", valeur, "\n")
    }
}

func creationTabportion(tabportion []map[string]int, size int, nbcpu int){
	for i:=0 ; i < len(tabportion) ; i++ {
		tabportion[i] = make(map[string]int)
		if i == 0 {
			tabportion[i]["start"] = 0
		} else {
			tabportion[i]["start"] = tabportion[i-1]["end"]
		}
		tabportion[i]["end"] = int((float64(size) / float64(nbcpu)) * float64(i+1))
	}
}

func sommer(tab []int, c chan<- int) {
	sum := 0
	for _, valeur := range tab {
		sum += valeur
	}
	c <- sum // send sum to c
}

func main() {
	// Récupération de la taille du tableau donnée en ligne de commande
    p_num := flag.Int("n", 5, "nombre")
    flag.Parse()

    // Tableau de taille dynamique, allant de 0 à *p_num+1
	tabnum := make([]int, *p_num)
	
	remplirTableau(tabnum)
	//afficherTableau(tabnum)

	nbcpu := runtime.NumCPU()
	fmt.Print("Nombre de processeurs : ", nbcpu, "\n\n")

	tabportion := make([]map[string]int, nbcpu)
	creationTabportion(tabportion, len(tabnum), nbcpu)

	resultat := make(chan int)

	for i:=0 ; i < nbcpu ; i++ {
		go sommer(tabnum[tabportion[i]["start"]:tabportion[i]["end"]], resultat)
	}

	fmt.Print("Voici les résulats intermédiaires :\n")
	tabresultat := make([]int, nbcpu)
	for i:=0 ; i < nbcpu ; i++ {
		tabresultat[i] = <- resultat
		fmt.Print(tabresultat[i], "\n")
	}

	go sommer(tabresultat, resultat)
	total := <- resultat

	fmt.Print("\nLa somme total est ", total, "\n")

    //fmt.Print(*resultat)
}