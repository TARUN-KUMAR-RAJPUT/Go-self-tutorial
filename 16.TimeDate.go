package main

import (

	"fmt"
	"time"
	"reflect"
	
)

func main() {
    
	// s := time.Now().Format("02-01-2006_15:04:06")
	s := time.Now()
	t := time.Now().Format("02-01-2006_15:04:06")

	ss := s.String()
	
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(ss))


	// fmt.Println(s)
}