package main

import "fmt"

func main(){
	var a,b,c,pol,otr int
	pol = 0
	otr=0
    fmt.Scan(&a, &b, &c)
    if a > 0{
		pol+=1
	}else if a < 0{
        otr+=1
	}
	if b > 0{
		pol+=1
	}else if b < 0{
        otr+=1
	}
	if c > 0{
		pol+=1
	}else if c < 0{
        otr+=1
	}
	fmt.Println(pol, otr)
	
	
}	