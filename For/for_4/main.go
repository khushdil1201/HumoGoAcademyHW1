package main

import "fmt"

func main(){
   var a float64
   fmt.Println("Введите число:")
   fmt.Scan(&a)
   for i:=1; i<=10;i++{
	fmt.Println(a*float64(i))
   }
}