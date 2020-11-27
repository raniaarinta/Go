package main

import(
	"fmt"
)
func main()  {
	var numberA int =5
	var pointerA *int = &numberA
	 fmt.Println(numberA)
	 fmt.Println(pointerA)
	 var numberB int=6
	 fmt.Println(numberB)
	 result:=change(&numberB,5)
	 fmt.Println(result)
	 

}
func change( old_value *int, new_value int) int{
	*old_value= new_value
	return *old_value

}
