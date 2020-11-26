package main
import(
	"fmt"
)
func main()  {
	for i:=1;i<=5;i++{
		fmt.Println(i)

	}
	result:=sqrt(2,3)
	fmt.Println("square root: ", result)
	while_loop()
	foreach_string()
}
func sqrt(input float32, power int) float32{
	var temp float32=1
	for i := 0; i <power; i++ {
		temp=temp*input
	}
	return temp;
}
func while_loop()  {
	i:=1
	for (i <=5){
		fmt.Println("while loop :",i)
		i++
	}
	
}
func foreach_string(){
	title:="how to do for each string"

	for index,letter:= range title{
		fmt.Println("index:  ", index,"letter: ",string(letter))
	}
	for _,letter:= range title{
		fmt.Println("just the letter without index: ",string(letter))
	}
}
