package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bssatya/learn_to_code_go/learn_go_with_tests/http_server"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Lets Play Poker")
	fmt.Println("Type {name} wins to reord a win")

	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
