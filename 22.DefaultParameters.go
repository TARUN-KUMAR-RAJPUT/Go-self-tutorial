package main

import (
	"fmt"
)


func fun(s string, t string, d ...string) {

	if len(d) > 0 {

		fmt.Println(d[0])
		fmt.Println(d[1])
		fmt.Println(d[2])

	}

	fmt.Println(s + t)
}


func main() {
    
	// fun("Tarun ", "Kumar")
	fun("Tarun ", "Kumar ", "Rajput", , "Broo")
}