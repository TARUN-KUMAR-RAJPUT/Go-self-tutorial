package main

import (
	"fmt"

	"github.com/datasweet/datatable"
)

func main() {

	dt := datatable.New("mytabl")

	dt.AddColumn("Name", datatable.String, datatable.Values("Tarun", "Itachi", "Sasuke"))
	dt.AddColumn("Eid", datatable.String, datatable.Values("123", "456", "789"))
	dt.AddColumn("Email", datatable.String, datatable.Values("Tarun@123", "Itachi@123", "Sasuke@123"))

	fmt.Println(dt)
}