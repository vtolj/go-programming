package main

import (
	"fmt"
)
func main() {
	a := (12 == 11)
	b := (12 <= 11)
	c := (12 >= 11)
	d := (12 != 11)
	e := (12 < 11)
	f := (12 > 11)

	fmt.Println(a,b,c,d,e,f)
}
