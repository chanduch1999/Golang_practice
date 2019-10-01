package main

import (
	"fmt"
)

func main() {

	var a,b,c,de,largest int;
   fmt.Scan(&a,&b,&c)
   de=a+b+c;
   if de>largest{
       largest=de;
   }
   de=a*b*c;
   if de>largest {
       largest=de;
   }
   de=a*(b+c);
   if de>largest   {
       largest=de;
   }
   de=a+(b*c);
   if de>largest   {
       largest=de;
   }
   de=(a+b)*c;
   if de>largest   {
       largest=de;
   }
   de=(a*b)+c;
   if de>largest  {
       largest=de;
   }
	fmt.Println(largest)
}