package main
import(
	"fmt"
)

func main()  {
	 age:= 9
	if age > 10 {
		fmt.Println("young adult")
	}else
	{
		fmt.Println("Teen")
	}
	var grade_result string = grade(80)
	fmt.Println(grade_result)
	var score_result int = score("A")
	fmt.Println(score_result)
}

func grade( score int ) string{
	var grade string
	if( score <= 50){
		grade="E"
	} else if(score <=60) {
		grade="D"
	} else if(score <=70) {
		grade="C"
	} else if(score <=80) {
		grade="B"
	} else if(score <=90) {
		grade="A"
	} else {
		grade="Null"
	}
	return grade
}

func score (grade string) int{
	var score int
	if(grade=="E"){
		score=50
	}else if (grade=="D") {
		score=60
	}else if (grade=="C") {
		score=70
	}else if (grade=="B") {
		score=80
	}else if (grade=="A") {
		score=90
	}
	return score
}