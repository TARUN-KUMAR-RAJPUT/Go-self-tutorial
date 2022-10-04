package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	t := time.Now().Format("02-01-2006")
	fmt.Println(reflect.TypeOf(t))
}