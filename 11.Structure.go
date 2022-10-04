package main
import "fmt"

type Bro struct {
    name string;
    age int;
}


func main(){
    
    var b1 Bro;
    b1.name = "Uchiha";
    b1.age = 24;
    
    fmt.Println(b1.name);
    fmt.Println(b1.age);
}
