package main

import(
	"fmt"
)
func main()  {
	fibonaci(6)

	
}
func fibonaci(number int)  {
	var t1,t2,next_term int
	t2=1
	t1=0
	next_term=0
	fmt.Println("Fibonaci series:")
	for i := 1; i <= number; i++ {
		if(i==1){
			fmt.Println(1)
			continue

		}
		if(i==2){
			fmt.Println(2)
			continue

		}
		next_term=t1+t2
		t1=t2
		t2=next_term
		fmt.Println(next_term)
	}

	
}