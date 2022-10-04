package main
import (
	
	"fmt"
	"os"
	"io/ioutil"

) 

func main() {
    
	fname := "text.txt"
	strData := "Yoo Broo, How are you ?"

	// Create file
	file, err := os.Create(fname)
	
	if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
     
    defer file.Close()
    
	
	//Write to file
    len, err := file.WriteString(strData)
 
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }

	
	// //Read file
	data, err := ioutil.ReadFile(fname)

    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }

    fmt.Printf("\nData: %s", data)

    
	// //Append to file
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer s.Close()

	if _, err = s.WriteString(strData); err != nil {
		panic(err)
	}
}
