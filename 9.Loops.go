package main
import "fmt"

func main(){
    
    for i := 0; i < 5; i++ {
        fmt.Println(i);
        
        if i == 3 {
            break;
        }
    }
    
    
    var arr = [5] int {10, 20, 30, 40, 50};
    
    for ix, val := range arr {
        fmt.Println(ix, val);
    }
    
}
