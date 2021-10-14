package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	imain := 0
	for ; imain < 5; imain++ {
		fmt.Println(imain)

	}
	// forever looping
	// for i := 0; i < 5; {
	// 	fmt.Println(i)

	// }
	fmt.Println("Print check using break")
	check := 0
	for {
		fmt.Println(check)
		check++
		if check > 10 {
			break
		}
	}
	fmt.Println("Print odd number using continue")
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Printing double loop ")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break
			}
		}
	}

	fmt.Println("Printing double loop  with break")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j > 3 {
				break
			}
		}
	}
	fmt.Println("Printing double loop break label")
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}
	statePopulations := map[string]int{
		"state1": 24324,
		"state2": 24324,
		"state3": 24324,
		"state4": 24324,
		"state5": 24324,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	son := "hello guys"
	for k, v := range son {
		fmt.Println(k, string(v))
	}
	for _, v := range son {
		fmt.Println(string(v))
	}
	fmt.Println(reverseString("Here is what you are missing in your benchmark: your solution is faster because it doesn't preserve combining characters. It is just unfair to compare them"))
}

func reverseString(input string) string {
	output := ""
	for _, v := range input {
		output = string(v) + output
	}
	return output
}
