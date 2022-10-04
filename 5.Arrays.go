package main
import "fmt"

func main(){
    
    // var array_name = [length]datatype{values} // here length is defined
    // var array_name = [...]datatype{values} // here length is inferred
    
    var arr1 = [3] int {1, 2, 3};
    var arr2 = [...]int{4, 5, 6, 7, 8};
	var n = len(arr1);
    
    fmt.Println(n);

    fmt.Println(arr1);
    fmt.Println(arr2);
    
    var arr3 = [5] int {};
    
    arr3[0] = 1;
    
    fmt.Println(arr3);
}
