package main
import "fmt"

func main(){
    
    // slice_name := []datatype{values}
    
    arr1 := []int{1,2,3};
    arr2 := []int{4,5,6};
    var n = len(arr1);
    var c = cap(arr1);
    
    fmt.Println(n);
    fmt.Println(c);
    fmt.Println(arr1);
    
    
    // slice_name := make([]type, length, capacity)
    
    // arr2 := make([]int, 3, 3);
    // arr2 := make([]int, 3);  if capacity not passed then capacity will be length
    
    var n2 = len(arr2);
    var c2 = cap(arr2);
    
    fmt.Println(n2);
    fmt.Println(c2);
    
    arr2[0] = 1;
    
    fmt.Println(arr2);
    
    // append() method to add element at append
    
    // slice_name = append(slice_name, element1, element2, ...)
    
    arr1 = append(arr1, 4, 5, 6);
    
    fmt.Println(arr1);
    
    // append one slice to end of other
    // slice3 = append(slice1, slice2...)
    
    arr3 := append(arr1, arr2...);
    
    fmt.Println(arr3);
    
    arr4 := arr3[0:3];
    
    fmt.Println(arr4);
}
