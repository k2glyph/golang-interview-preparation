package main

import "fmt"

func main() {
	// statePopulations := make(map[string]int)
	statePopulations := map[string]int{
		"state1": 24324,
		"state2": 24324,
		"state3": 24324,
		"state4": 24324,
		"state5": 24324,
	}
	statePopulations["state6"] = 42
	delete(statePopulations, "state1")
	fmt.Println(statePopulations)
	_, ok := statePopulations["state1"]
	fmt.Printf(" state 1 and exists: %v \n", ok)
	fmt.Printf(" length %v\n", len(statePopulations))

}
