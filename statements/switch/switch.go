package main

import "fmt"

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one five or ten")
	case 2, 4, 6:
		fmt.Println("two four six")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not one or two")
	}
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}
	var it interface{} = "1"
	switch it.(type) {
	case int:
		fmt.Println("i is an int")
		break
		// fmt.Println("This will print too")
	case float64:
		fmt.Println("is is a float")
	case string:
		fmt.Println("is a string")
	default:
		fmt.Println("i is another type")
	}

}
