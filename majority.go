package main

import (
	"fmt"
)

func main() {
	var flag,ind,t,a,i,n int
	fmt.Scan(&t)
	for t>0 {
		fmt.Scan(&n)
		 flag =0 
		 ind=-1
		s :=make([]int ,100)
		for i=0;i<n;i++ {
			fmt.Scan(&a)
			s[a]++
		}
		for i=0;i<n;i++ {
			if s[i]>=n/2{
				flag = 1
				ind =i
			}
		}
		if(flag==1){
			fmt.Println(s[ind])
		}else {
			fmt.Println(-1)
		}
		t=t-1
	}
}
