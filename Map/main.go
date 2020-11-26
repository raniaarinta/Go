package main

import(
	"fmt"
)
func main()  {
	var myMap map[string] int
	myMap= map[string] int{}
	myMap["satu"]=1
	myMap["dua"]=2
	fmt.Println(myMap)
	fmt.Println("=================================================")
	Mapstring:=map[string]string{"language":"Go","numbers":"1",}
	fmt.Println(Mapstring["numbers"])
	fmt.Println("=================================================")
	for key, value:= range Mapstring{
		fmt.Println("key: ",key,"Value: ",value)
	}
	fmt.Println("=================================================")
	delete(myMap, "satu")
	for _, value:= range myMap{
		fmt.Println("Value: ", value)
	}
	fmt.Println("=================================================")
	value,isAvailable:=Mapstring["cool"]
	fmt.Println(value)
	fmt.Println(isAvailable)
}
