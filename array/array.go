package main

import "fmt"

func main() {
	grades := [3]int{1, 2, 2}
	fmt.Printf("Grades %v\n", grades)
	grades1 := [...]int{324, 324, 324}
	fmt.Printf("Grades %v\n", grades1)
	var students [4]string
	students[0] = "Dinesh"
	students[1] = "Ramesh"
	students[2] = "hari"
	students[3] = "gopal"
	fmt.Printf("Students %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 2}, [3]int{1, 4, 5}}

	fmt.Printf("Identity Matrix: %v", identityMatrix)
}
