package main

import (
	"fmt"
)

func main() {
	var k,n,w int
	fmt.Scan(&k,&n,&w)
	w = k*w*(w+1)/2
	fmt.Println(w-n)
}
