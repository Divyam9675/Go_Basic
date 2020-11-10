package main

import "fmt"


type Basic struct{

	firstName string
	lastName string
	Age int
	city string
}

func (b Basic) access()string{

	return "Hello "+ b.firstName + "how are you" + b.lastName

}

func (b *Basic) change(){

	b.Age++
	b.lastName = "sharma"

}



func main(){

	obj1 := Basic{"Divyam","Gupta",23,"BSR"}
	obj2 := Basic{"Ramesh","Gupta",22,"goa"}
	obj3 := Basic{"Suresh","Gupta",24,"goa"}
	obj4 := Basic{"Dinesh","Gupta",25,"goa"}


	fmt.Println(obj1,obj2,obj3,obj4)

	fmt.Println(obj1.firstName)

	obj1.city = "goa"

	fmt.Println(obj1)

	fmt.Println(obj1.access())

	obj3.change()

	fmt.Println(obj3)


}