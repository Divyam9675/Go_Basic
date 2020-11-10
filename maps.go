package main

import "fmt"

func main(){

	// Creation map

	//email := make(map[string]string)

	// Assign key value


//	email["Bob"] = "g@gmail"
//	email["Rohit"] = "r@gmail"




	// 2nd way define map and assign value

	email := map[string]string{"Bob": "b@gmail.com", "Ramesh": "r@gmail"}
	
	fmt.Println(email)
	fmt.Println(len(email))
	
	email["suresh"] = "s@gmail.com"

	fmt.Println(email)

	delete(email, "suresh")

	fmt.Println(email)






	

}