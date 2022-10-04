package main
import (
	"fmt"
	"github.com/fatih/structs"
	"strconv"
)

type MyTable struct {
    
	PROFORMA_ID string 
	PROFORMA_STATUS string
	FK_COMPANYID string
	GLID string 
	PROFORMA_ISSUEDATE string
}
  
func main() {

	var tb [] MyTable

	tb = append(tb, MyTable{PROFORMA_ID: "1564712", PROFORMA_STATUS: "P", FK_COMPANYID: "9688884", 
	             GLID : "67491298", PROFORMA_ISSUEDATE: "01-02-22"})
	
	tb = append(tb, MyTable{PROFORMA_ID: "1564712", PROFORMA_STATUS: "P", FK_COMPANYID: "9688884", 
	GLID : "67491298", PROFORMA_ISSUEDATE: "01-02-22"})
    

	fmt.Println(len(tb))
	s := strconv.Itoa(len(tb))
	fmt.Println(s)			

	
	fmt.Println(tb[0])			
	
	fmt.Println(tb[0].PROFORMA_ID)			
    
	for _, row := range tb {
		fmt.Println(row.PROFORMA_ID)
	}


	Columns := structs.Names(&MyTable{})
	
	for _, ColumnName := range Columns {
		fmt.Println(ColumnName)		
	}

	fmt.Println(Columns)

	fmt.Println("Yoo")

}

