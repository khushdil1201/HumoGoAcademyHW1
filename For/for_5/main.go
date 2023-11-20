package main

import "fmt"

func main(){
   var a float64
   fmt.Println("Введите число:")
   fmt.Scan(&a)
   for i:=0.1; i<=1;i+=0.1{
	fmt.Println(a*i)
   }
}