package main
import "fmt"

func main(){
    
    var mp = map[int]int { 1 : 10, 2 : 20, 3 : 30};
    
    mp[4] = 40;
    mp[5] = 50;
    
    fmt.Println(mp);
    fmt.Println(mp[1]);
    
    delete(mp, 2);
    fmt.Println(mp);
    
    x , ok := mp[3];
    
    fmt.Println(x, ok);
    
    for k, v := range mp {
        fmt.Println( k, v);
    }
}
