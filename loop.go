package main

import "fmt"

func main(){

	

	for x:= 1; x<100; x++{


		if x%15 == 0{

			fmt.Println("FIZZ BUZZ ", x)


		}else if x%3 ==0{

			fmt.Println(" FIZZ ", x)

		}else{

			fmt.Println(" BUZZ", x)

		}

		
	}
}	

		

		