/*
October 3rd, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
*/

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Enter a number to check if it is a factorial:")
	var revFactorial int
	_, error := fmt.Scan(&revFactorial)
	
	divisionResult := float64(revFactorial) //revFactorial / 1 outside of loop practically
	x := 1.0

	for divisionResult > 1{
		x += 1.0
		divisionResult = divisionResult / x
	}

	if divisionResult == float64(int64(divisionResult)){
		fmt.Println("The number is the factorial of ", x, "!")
	}else{
		fmt.Println("Not a factorial.")
	}

}

