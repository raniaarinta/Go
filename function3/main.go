package main

import(
	"fmt"
)
func main()  {
	a,p:=rectangle(4,3)
	fmt.Println("area of rectagle: ", a)
	fmt.Println("Perimeter: ",p)

	
}

func rectangle(length int , width int) (int, int){
	area:=length*width
	Perimeter:=2*(width+length)
	return area,Perimeter

}