package main

import (
	"fmt"
)

func main() {
	var t,n,i,j,flag int 
	fmt.Scan(&t)
	for t>0{
		fmt.Scan(&n)
		s := make([]int, n)
		for i=0;i<n;i++{
		fmt.Scan(&s[i])
		}

		for i=0;i<n;i++{
			 flag = 0 
			for j=i+1;j<n;j++ {
				if s[i]<s[j] {
					flag=1
				}
			}
			if flag==0{
				fmt.Println(s[i])
			}
		}
		t=t-1
	}
}