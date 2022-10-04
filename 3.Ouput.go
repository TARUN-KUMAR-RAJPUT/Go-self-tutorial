package main
import "fmt"

func main(){

    // Print() prints its arguments with their default format.
    // Println() same but add new line
    // Printf()
    
    var x, y = 2, 3;
    
    fmt.Print(x, " ", y);
    fmt.Print(x, y);
    fmt.Print(x, "\n");
    fmt.Println(x, y);
    
    var i string = "Hello"
    var j int = 15

    fmt.Printf("i has value: %v and type: %T\n", i, i)
    fmt.Printf("j has value: %v and type: %T", j, j)
  
}
