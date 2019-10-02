package main

import (
	"fmt"
	"math"
)
func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
	var n,a int 
	fmt.Scan(&n)
	for n>0 {
		fmt.Scan(&a)
		a = math.Sqrt(a)
		if (IsPrime(a)){
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
		n=n-1
	}
}