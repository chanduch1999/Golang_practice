package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	var i,s,a,n float64 
	fmt.Scan(&t)
	for t>0{
		fmt.Scan(&n)
		s=0
		for i=0;i<n;i++{
			fmt.Scan(&a)
			s+=a
		}
		s = s/n
		fmt.Println(math.Ceil(s))
		t=t-1
	}
}