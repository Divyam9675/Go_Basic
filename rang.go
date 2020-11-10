package main

import "fmt"



func main(){

  arr := []int{10,20,30,40}

  for _ , j := range arr{

	 fmt.Printf(" %d value \n",  j)

  }
  
var Detail = map[string]string{"Name": "Divyam", "age": "20", "sallary": "400"}

for i,j := range Detail{

	fmt.Println(i,j)
}

// fmt.Println(Detail)
}


// range with map

 




