package main

import (
	"fmt"
	"time"
)

func main() {
	var x int = 10
	fmt.Println(x)
	time.Sleep(3 * time.Second)
	fmt.Println(x + 3)
}