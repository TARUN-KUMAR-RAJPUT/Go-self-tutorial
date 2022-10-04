package main
import "fmt"

func main(){
    
    var x, y = 10, 10;
    
    if x > y {
        fmt.Println("x is big");
    } else if x < y {
        fmt.Println("y is big");
    } else {
        fmt.Println("equal");
    }
    

    // wrong..  else should start right to closing bracket of if.

    // if x > y {
    //     fmt.Println("x is big");
    // } 
    // else if x < y {
    //     fmt.Println("y is big");
    // } 
    // else {
    //     fmt.Println("equal");
    // }
    
    
}
