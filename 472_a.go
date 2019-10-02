package main

import (
	"fmt"
)

func main() {
var n int 
fmt.Scan(&n)
if (n%2==1){
	fmt.Println(9,n-9)
}else{
	fmt.Println(8,n-8)
}
}
