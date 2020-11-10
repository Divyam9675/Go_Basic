package main

import "fmt"

func calc() func(int)int {

	sum := 0
	

	fmt.Print("Hello coder")

	return func(x int) int{

		

		sum += x

		return sum

	}

}

func main(){

	sum:= calc()


	//fmt.Println(sum(2))

for i:=0; i<10; i++{

	fmt.Println(sum(i))

}
	
}