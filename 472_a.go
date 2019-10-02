package main

import (
	"fmt"
)

func main() {
var n int 
fmt.Scan(&n)
if (n%2==0){
	fmt.Println(8,n-8)
}else{
	fmt.Println(9,n-9)
}
}