package main

import (
	"fmt"
	"math"
)

func main() {
	if false {
		fmt.Println("The test is true")
	}
	statePopulations := map[string]int{
		"state1": 24324,
		"state2": 24324,
		"state3": 24324,
		"state4": 24324,
		"state5": 24324,
	}
	if pop, ok := statePopulations["state1"]; ok {
		fmt.Println("Value exists", pop, ok)
	}
	number := 50
	guess := 30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println(("The guess must be between 1 and 100"))
	} else {

		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You win !")
		}
		fmt.Println(number < guess, number > guess, number != guess)
	}

	myNum := 0.1
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
