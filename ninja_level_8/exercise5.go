package main

import (
	"sort"
	"fmt"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user
func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	printUsers(users)

	// your code goes here
	sort.Sort(ByAge(users))
	fmt.Println("After sorting ByAge")
	printUsers(users)

	sort.Sort(ByLast(users))
	fmt.Println("After sorting ByLast name")
	printUsers(users)

}

func printUsers(users []user) {
	for _, user := range users {
		fmt.Println("---", user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println(saying)
		}
	}
}
