package main

import (
	"fmt"
	"first_part/Calculation"

)

func main()  {
	fmt.Println("hello world")
	// test := Test_String()
	// fmt.Println(test)
	add_result:=Calculation.Add(3,4)
	minus_result :=Calculation.Minus(10,5)
	var div_result float32= Calculation.Div(8,7)
	var multiply_result float32=Calculation.Multiply(5,5)
	fmt.Println("add result: ",add_result)
	fmt.Println("minus result:",minus_result)
	fmt.Println("division result: ",div_result)
	fmt.Println("mulitply result: ",multiply_result)
}