package main

import (
	"fmt"
)

func main () {
	favSport := "football"

	switch (favSport) {
	case "cricket":
		fmt.Println("go to local ground")
	case "football":
		fmt.Println("go to local club")
	case "skiing":
		fmt.Println("go to mountains")
	}
}