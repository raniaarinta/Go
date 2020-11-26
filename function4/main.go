package main

import(
	"fmt"
)
func main()  {
	a,p:=square(5)
	fmt.Println("area of square: ", a)
	fmt.Println("Perimeter: ",p)

	
}

func square(length int) (area int,  perimeter int){
	area =length*length
	perimeter =4*length
	return 

}