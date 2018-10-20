package main

import (
	"encoding/json"
	"fmt"
)


type user struct {
	First string
	Age int
}

func main() {
	u1 := user{
		First: "Raj",
		Age: 75,
	}

	u2 := user{
		First: "Ambi",
		Age: 60,
	}

	u3 := user{
		First: "Uppi",
		Age: 40,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	xb, err := json.Marshal(users)
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}