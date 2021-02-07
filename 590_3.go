package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	var j,s,a,n float64 
	fmt.Scan(&t)
	for t>0{
		fmt.Scan(&n)
		s=0
		for j=0;j<n;j++{
			fmt.Scan(&a)
			s+=a
		}
		s = s/n
		fmt.Println(math.Ceil(s))
		t=t-1
	}
}
