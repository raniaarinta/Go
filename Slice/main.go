package main

import(
	"fmt"
)

func main()  {
	var languages [] string
	languages= append(languages,"Go","#C","python")
	fmt.Println(languages)
	printData(languages)
	gameconsole:=[]string {"PS4",
							"Nintendo Switch",
	}
	foreachPrint(gameconsole)
	
}
func printData(input[] string)  {
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
	
}

func foreachPrint(input[] string) {
	for _,data:= range input{
		fmt.Println("data: ",data)
	}
	
}