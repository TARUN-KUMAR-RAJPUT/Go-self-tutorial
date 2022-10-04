package main
import "fmt"

func main(){


// 1. Single variable  

//  var variablename type = value
// 	variablename := value

	var s string = "Uchiha" //type is string
    var t = "Itachi" //type is inferred
    x := 2 //type is inferred
    
    var flag = "true";
    
    var y int;
    y = 3;
    
    fmt.Println(s);
    fmt.Println(t);
    fmt.Println(x);
    fmt.Println(flag);
    fmt.Println(y);


// 2. Multiple variable with same type    
    
    var a, b, c, d int = 1, 3, 5, 7;
    
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    
// 3. Multiple variable with different type

   var p, q = 5, "hello";
   
   fmt.Println(p);
   fmt.Println(q);

// 4. Constants
// const CONSTNAME type = value
   
   const PI = 3.14
  
}
