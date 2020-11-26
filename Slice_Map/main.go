package main

import(
	"fmt"
)

func main()  {
	students:=[]map[string]string{
		{"id":"123a","name:":"rania","score":"80"},
		{"id":"123b","name:":"karla","score":"60"},
	}
	fmt.Println(students)
	for _, value:= range students{
		fmt.Println("ID: ",value["id"])
	}
}