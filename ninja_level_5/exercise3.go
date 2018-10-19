package main

import (
	"fmt"
)

type vehical struct {
	doors int
	color string
}

type truck struct {
	vehical
	fourWheel bool
}

type sedan struct {
	vehical
	luxury bool
}
func main(){
	t1 := truck {
		vehical: vehical{
			doors : 2,
			color : "Black",
		}, 
		fourWheel : true,
	}

	s1 := sedan {
		vehical : vehical{
			doors : 4,
			color : "Red",
		},
		luxury : true,
	}

	fmt.Println(t1)
	fmt.Println(s1)
}