package main

import (
	"fmt"
	"github.com/rodaine/table"
)

func main() {

	tbl := table.New("ID", "Name", "Score")
	tbl.AddRow(1, "Tarun", 15)
	tbl.AddRow(2, "Kumar", 25)
	tbl.AddRow(3, "Rajput", 35)

	tbl.Print()

}