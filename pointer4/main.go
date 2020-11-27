package main

import(
	"fmt"
)
type Employee struct{
	id int
	Name string
	position string
}
func main()  {
	myEmployee:= Employee{1,"Kim Namjoon","idol"}
	fmt.Println(myEmployee)
	promotion(&myEmployee)
	fmt.Println(myEmployee)
}

func promotion(Data *Employee){
	Data.position= "entertainer"
}
