package main
import(
	"fmt"
)

func main(){
	number:= 2
	switch number {
	case 1: 
		fmt.Println("you choose number 1")
	case 2:
		fmt.Println("you choose number 2")
	default:
		fmt.Println("null")
	}
	var input string ="cat"
	switch  {
		case input =="cat": 
			fmt.Println("you choose cat")
		case  input == "fish":
			fmt.Println("you choose fish")
		default:
			fmt.Println("null")
			
	}

	

}
