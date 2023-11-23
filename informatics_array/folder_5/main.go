package main

import "fmt"

func main(){
   var n,min int 
   fmt.Scan(&n)
   array:=make([]int, n)
   for i:=0;i<n;i++{
	  fmt.Scan(&array[i])
   }
   for i:=0;i<n;i++{
	  if array[i]%2!=0{
		min=array[i]
		break
	  }
   }
   for i:=0;i<n;i++{
	if array[i]%2!=0{
	  if array[i]<min{
		min=array[i]
	  }
	}
 }
   fmt.Println(min)

}