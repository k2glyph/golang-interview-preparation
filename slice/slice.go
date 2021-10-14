package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 4
	fmt.Printf("Array a: %v\n", a)
	fmt.Printf("Array b: %v\n", b)
	slice := make([]int, 3, 100)
	fmt.Printf("Length: %v \n", len(slice))
	fmt.Printf("Capacity: %v \n", cap(slice))

	newSlice := []int{}
	newSlice = append(newSlice, 10)
	fmt.Printf("Values: %v\n", newSlice)
	fmt.Printf("Length: %v \n", len(newSlice))
	fmt.Printf("Capacity: %v \n", cap(newSlice))
	newSlice = append(newSlice, 10, 1, 2, 3, 4, 5)
	fmt.Printf("Values: %v\n", newSlice)
	fmt.Printf("Length: %v \n", len(newSlice))
	fmt.Printf("Capacity: %v \n", cap(newSlice))

	newSlice = append(newSlice, a...)
	fmt.Printf("Values: %v\n", newSlice)
	fmt.Printf("Length: %v \n", len(newSlice))
	fmt.Printf("Capacity: %v \n", cap(newSlice))

	ab := []int{1, 2, 3, 4, 5}
	// ba := ab[:len(ab)-1]
	ab = append(ab[:2], ab[3:]...)
	fmt.Printf("Values: %v \n", ab)
}
