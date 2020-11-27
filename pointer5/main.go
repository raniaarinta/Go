package main

import(
	"fmt"
)
type Employee struct{
	id int
	Name string
	position string
}

func(data *Employee) promotion(){
	data.position= "entertainer"
}


func main()  {
	myEmployee:= Employee{1,"Kim Namjoon","idol"}
	fmt.Println(myEmployee)
	myEmployee.promotion()
	fmt.Println(myEmployee)
}


