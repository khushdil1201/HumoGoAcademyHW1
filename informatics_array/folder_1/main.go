package main

import "fmt"

func main(){
   var n int 
   var check bool
   fmt.Scan(&n)
   array:=make([]int, n)
   for i:=0;i<n;i++{
	  fmt.Scan(&array[i])
   }
   for i:=1;i<n;i++{
	if (array[i]>0 && array[i-1]>0) || (array[i]<0 && array[i-1]<0){
		check=true
		break
	}
   }
   if check {
	fmt.Println("YES")
   }else{
	fmt.Println("NO")
   }

}