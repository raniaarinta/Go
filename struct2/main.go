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
	result:= getItemNamebyID(userItem, 1)
	fmt.Println(result)
	result_price:= getPricebyID(userItem, 1)
	fmt.Println(result_price)

}




func getItemNamebyID(data Item, id int) string  {

	if (data.ID == id){
		return data.Item_name
	} else {
		return"data is not found"
	}
	
}
func getPricebyID(data Item,id int) int  {
	if (data.ID == id){
		return data.Item_price
	} else {
		return 0
	}
	
}