package main

import (
	"fmt"
)

func  main()  {
	var languages [3] string
	languages[0]="go"
	languages[1]="ruby"
	languages[2]="rust"
	fmt.Println(languages)
	for i := 0; i <len(languages); i++ {
		if (languages[i]=="go") {
			languages[i]="python"
		}
		
	}
	fmt.Println(languages)
	lang:= [...]string {"ruby",
						"go",
						"python",
						"#C",
					}
	for index,la:= range lang{
		fmt.Println("index number: ",index,"value: ",la)
	}
	
	var avg [6] float64 
	avg [0]= 2.00
	avg [1]= 2.00
	avg [2]= 3.00
	avg [3]= 3.00
	avg [4]= 4.00
	avg [5]= 4.00
	avg_result:=average(avg)
	for _,a:= range avg{
		fmt.Println("data: ",a)
	}
	fmt.Println("average: ",avg_result)
	
}

func average( data[6] float64) float64{
	length:= len(data)
	var temp float64
	for i := 0; i < length; i++ {
		temp=temp + data[i]
		
	}
	return temp/ float64(length)
}