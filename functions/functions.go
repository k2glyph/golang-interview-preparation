package main

import (
	"fmt"
	"log"
)

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
func (g *greeter) greetWitPointer() {
	fmt.Println((*g).greeting, (*g).name)
}
func main() {
	// for i := 0; i < 5; i++ {
	// 	sayMessage("hello go!", i)
	// }
	greeting := "Hello"
	name := "Dinesh"
	sayGreeting(&greeting, &name)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 8)
	sum := sumReturn(1, 2, 3, 4, 5)
	fmt.Println("Sum is", *sum)
	fmt.Println(sumReturnWithName(1, 2, 3, 4, 5, 6, 7, 8))
	d, err := divide(10, 1.0)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(d)
	for i := 0; i < 5; i++ {
		var f func(int) = func(i int) {
			fmt.Println(i)
		}
		f(i)
	}

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	dresult, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dresult)

	g := greeter{
		greeting: "Welcome to the world of golang ",
		name:     "Dinesh",
	}
	g.greet()
	g.greetWitPointer()
}
func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}

func sayGreeting(greeting, name *string) {
	*name = "Ted"
	fmt.Println(*greeting, *name)

}
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}

func sumReturn(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sumReturnWithName(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		// panic("Cannot provide zero as second value")
		return 0.0, fmt.Errorf("Cannot provide zero as second value")
	}
	return a / b, nil
}
