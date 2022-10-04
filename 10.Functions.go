package main
import "fmt"

func fun() {
    fmt.Println("fun is called ");
}

func add(x int, y int) {
    fmt.Println(x + y);
}

func add1(x int, y int) int {
    return x + y;
}

func main(){
    
    fun();
    add(10, 5);
    fmt.Println(add1(10, 5));
}
