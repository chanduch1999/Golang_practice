package main

import (
	"fmt"
)

func main() {
	var n,mishka,chris,draw,mw,cws,i int
    fmt.Scan(&n)
    for  i=0;i<n;i++ {
        fmt.Scan(&mishka, &chris)
        if(mishka>chris){
            mw++;
        }else if(mishka<chris){
            cw++;
        }else
        {
            draw++;
        }
    }
    if(mw>cw){
        fmt.Println("Mishka")
    }else if(cw>mw){
        fmt.Println("Chris")
        }else{
			fmt.Println("Friendship is magic!^^")
        }
}