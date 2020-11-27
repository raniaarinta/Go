package main

import(
	"fmt"
)
type Item struct {
	ID int
	Item_name string
	Item_price int

}

func main()  {
	myItem:= Item{}
	myItem.ID =1
	myItem.Item_name ="Water Bottle"
	myItem.Item_price=4
	fmt.Println(myItem)

	userItem:=Item{ ID:1,Item_name:"books",Item_price:4, }
	fmt.Println(userItem)

}

	
