package main

import (
    "fmt"
    "database/sql"
)

func main() {

	var s sql.NullString
    fmt.Println(s.Valid)
}