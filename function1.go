package main

import "fmt"

func calc(p , q int)(y,z int) {


	y = p+q
	z = p*q
	
return

}



func main(){

	x, y:= 40,50


	res1, res2 := calc(x, y) 
	

	fmt.Println("print the sum of the no.", res1, res2)

}