package main

import (
	"fmt"
)

func main() {

   var a,b,c,res,largest int;
   fmt.Scan(&a,&b,&c)
   res=a+b+c;
   if res>largest{
       largest=res;
   }
   res=a*b*c;
   if res>largest {
       largest=res;
   }
   res=a*(b+c);
   if res>largest   {
       largest=res;
   }
   res=a+(b*c);
   if res>largest   {
       largest=res;
   }
   res=(a+b)*c;
   if res>largest   {
       largest=res;
   }
   res=(a*b)+c;
   if res>largest  {
       largest=res;
   }
	fmt.Println(largest)
}
