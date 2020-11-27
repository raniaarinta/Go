package main

import(
	"fmt"
)
func main()  {
	numberA:=5
	pointerA:=&numberA
	fmt.Println("value: ",numberA)
	fmt.Println("pointer adress: ",pointerA)
	fmt.Println("get pointer value: ",*pointerA)
	*pointerA=10
	fmt.Println("new pointer value: ",*pointerA)

}
