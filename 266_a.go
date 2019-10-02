package main
import (
	"fmt"
	"math"
)
func main() {

 var n, m, a, b int
    
    fmt.Scan( &n ,&m,&a,&b)
    if (m * a <= b){
        fmt.Println( n * a)
	}else {
		fmt.Println((n/m) * b + math.Min((n%m)*a, b)) 
	}

}