package main

import "fmt"

func main(){
   var n int
   fmt.Println("Введите размер массива: ")
   fmt.Scan(&n)
   array:=make([]int,n)

   fmt.Println("Введите элементы массива: ")
   for i:=0;i<n;i++{
	fmt.Scan(&array[i])
   }

   max:=array[0]
   min:=array[0]
   for i:=0; i < n;i++{
	  if array[i]>max{
		max=array[i]
	  }
	  if array[i]<min{
		min=array[i]
	  }
	}
    
	for i:=0; i < n;i++{
		if array[i]==max{
		  array[i]=min
		}
	  }
    
	fmt.Println("Ответ: ")
	fmt.Println(array)

}