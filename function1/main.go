package main

import(
	"fmt"
)
func main()  {
	printresult()
	result:=printstring()
	sentence:=printsentence("rania")
	fmt.Println(result)
	fmt.Println(sentence)
	
}
func printresult()  {
	fmt.Println("print result")
}
func printstring() string{
	return "print string"
}
func printsentence(input string) string{
	 newsentence:="hello there "+ input
	 return  newsentence
}
