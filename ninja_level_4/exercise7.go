package main

import (
	"fmt"
)


func main () {
	bondDetails := []string{"James", "Bond", "Shaken not stirred"}
	moneyPennyDetails := []string{"Miss", "MoneyPenny", "Helloooo Jamesss"}

	xxs := [][]string{bondDetails, moneyPennyDetails}
	fmt.Println(xxs)
	fmt.Printf("%T\n", xxs)

	for io, vo := range xxs {
		fmt.Println("record ar index ", io, "is", vo)
		for ii, vi := range vo {
			fmt.Println("\tindex position", ii, "value = ", vi)	
		}
	}
}