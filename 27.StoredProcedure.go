package main

import (

	"fmt"
	// "io/ioutil"
	// "net/http"
	"os"
	// "strconv"
	// "strings"
	// "time"
	// "reflect"
	// "github.com/google/uuid"
	// "strconv"
	// "github.com/fatih/structs"

	


	db "../components/db"
	setup "../setup"
	// mailer "../api/mail"
)


func testQuery() {
	
	mainConnection, err := db.GetDatabaseConnection("main")

	if err != nil {
		fmt.Println(err)
	}

	
	query := `
	          select * from test_table_uchiha
	    	 `

	params := make([]interface{}, 0)
	
	result, err := db.SelectQuerySql(mainConnection, query, params)
	
	if err != nil {
		fmt.Println(err)
	}
    
	fmt.Println(result)

}

func storPro() {

	mainConnection, err := db.GetDatabaseConnection("main")
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	x := 11
	y := "bro3"
	z := "bro3@123"
    
	fmt.Println(x, y, z)

	query := `

	    call  TEST_SP_UCHIHA(:1, :2, :3)

		`

	// query := `

	// 	DECLARE 

	// 	a NUMBER; 
	// 	b VARCHAR2(20); 
	// 	c VARCHAR2(20); 
		
	// 	BEGIN 
		
	// 	a := :x; 
	// 	b := :y; 
	// 	c := :z; 
		
	// 	TEST_SP_UCHIHA(MY_EID => a, 
	// 			MY_NAME => b, 
	// 			MY_EMAIL => c); 
								
	// 	END ;`

	
	
	// query := `
	
	// 	DECLARE 
		
	// 	IN_WO_ID NUMBER; 
	// 	IN_LAST_COMMENTS VARCHAR2(200); 
	// 	IN_LAST_MODIFIED_BY_EMPID NUMBER;
	// 	IN_WIP_COST NUMBER;
	// 	IN_WO_CANCEL_REASON_ID NUMBER;
		
	// 	OUT_SUCCESS_FLAG NUMBER;                         Output parameter
		
	// 	BEGIN 

	// 	IN_WO_ID := :WOID; 
	// 	IN_LAST_COMMENTS := :LastComments;
	// 	IN_LAST_MODIFIED_BY_EMPID := :LastModifiedByEmp;
	// 	IN_WIP_COST := :WIP_Cost;
	// 	IN_WO_CANCEL_REASON_ID := 0;
	
	// 	TEST_SP_CANCEL_WO(IN_WO_ID => IN_WO_ID,
	// 		IN_LAST_COMMENTS => IN_LAST_COMMENTS,
	// 		IN_LAST_MODIFIED_BY_EMPID => IN_LAST_MODIFIED_BY_EMPID,
	// 		OUT_SUCCESS_FLAG => OUT_SUCCESS_FLAG,
	// 		IN_WIP_COST => IN_WIP_COST,
	// 		IN_WO_CANCEL_REASON_ID => IN_WO_CANCEL_REASON_ID); 
								
	// 	END ;`

	





	fmt.Println(query)
	params := make([]interface{}, 0)
	params = append(params, x, y, z)
	

	result, err := db.ExecuteQuerySql(mainConnection, query, params, true)

	if err != nil {
		fmt.Println(err)
	}
    
	fmt.Println(result)

}

func main() {

	setup.StartSetup()
	dbMain, err := db.GetDatabaseConnection("main")

	if err != nil {
		fmt.Println("Could not start the main database because of following err " + err.Error() + " ...!!!")
		os.Exit(1)
	}
	err = dbMain.Ping()
	if err != nil {
		fmt.Println("Could not start the main database because of following err " + err.Error() + " ...!!!")
		os.Exit(1)
	} else {
		fmt.Println("Connection to the main-database established ...!!!")
	}

	defer func() {
		fmt.Println("DB Close")
		err = dbMain.Close()
		if err != nil {
			fmt.Println("Could not close the main database because of following err " + err.Error() + " ...!!!")
		}
	}()

	fmt.Println("Hello world..Connection Succesfull")
	
	// practice()
	storPro()
    
}