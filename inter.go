package main

import "fmt"

type inter interface{

	calc()int 
   
}

type student1 struct{

	maths int
	english int
}

func (s1 student1) calc() int{

	sum := (s1.english + s1.maths)/2

	fmt.Println(sum)

	return sum

}



type student2 struct{

	maths int
	english int
}

func (s2 student2) calc() int{

	sum := (s2.english + s2.maths)/2

	fmt.Println(sum)

	return sum

}



func show(i inter) int{

	return i.calc()

	
	
}


func main(){

	obj1 := student1{50,25}
	obj2 := student2{60,30}

	fmt.Println("sum is",show(obj1))
	fmt.Println("sum is",show(obj2))


}