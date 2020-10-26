// Author : chandu
package main

import (
	"fmt"
)
func main() {


var a,b,d,rem,sum int
    fmt.Scan( &a,&b)
     sum = a;
    for a >= b{
        d = a / b
        sum += d
        rem=a % b
		a = d+ rem  
	  }
	fmt.Println(sum)
}
