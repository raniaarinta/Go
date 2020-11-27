package main

import(
	"fmt"
)
type Item struct {
	ID int
	Item_name string
	Item_price int

}
type Catalog struct{
	id_catalog int
	Catalog_name string
	myItem []Item
}
func (data Catalog) Printall(){
	fmt.Printf("Catalog name : %s",data.Catalog_name)
	fmt.Println("\n =======Item========")
	for _,item_data:= range data.myItem{
		fmt.Println("Item: ",item_data.Item_name)
		fmt.Println("Price: $ ",item_data.Item_price)
	}

}


func main()  {
	

	userItem:=Item{ ID:1,Item_name:"Harry potter",Item_price:4,}
	userItem2:=Item{ ID:2,Item_name:"IT",Item_price:4,}
	userItem3:=Item{ ID:3,Item_name:"Twilight",Item_price:4,}
	all_item := []Item{userItem,userItem2,userItem3}
	userCatalog:=Catalog{1,"books",all_item,}

	userItem4:=Item{ ID:4,Item_name:"Nintendo Switch",Item_price:4,}
	userItem5:=Item{ ID:5,Item_name:"PS4",Item_price:4,}
	userItem6:=Item{ ID:6,Item_name:"PS5",Item_price:4,}
	video_games_item:= [] Item{userItem4,userItem5,userItem6}
	userCatalog2:=Catalog{2,"Video Games",video_games_item,}

	fmt.Println("=======catalog========")
	userCatalog.Printall()
	fmt.Println("=======search Item by Name========")
	searchItem:=searchItemName(userCatalog,"Twilight")
	fmt.Println(searchItem)
	fmt.Println("========search Catalog by Name==========")
	searchCatalogbyName(userCatalog2, "Video Games")
	

}

func searchItemName(data Catalog,item string) string  {
	var result string
	for _,item_data:= range data.myItem{
		if (item_data.Item_name == item) {
			result= fmt.Sprintf("Item Name : %s, Item Price: %d",item_data.Item_name,item_data.Item_price)
		}else{
			result= "data not found"
		}
	}
	return result
}

func searchCatalogbyName(data Catalog, search string)  {
	if(data.Catalog_name== search){
		for _,item_data:= range data.myItem{
			fmt.Println("Item: ",item_data.Item_name)
			fmt.Println("Price: $ ",item_data.Item_price)
		}

	}else {
		fmt.Println("data not found")
	}
	
	
}

