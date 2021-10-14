package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	ab := [3]int{1, 2, 3}
	bc := &ab[0]
	ca := &ab[1]
	fmt.Printf("%v %p %p\n", ab, bc, ca)
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
	(*ms).foo = 82 // ms.foo will also work
	fmt.Println((*ms).foo)
	ms.foo = 81
	fmt.Println(ms.foo)

	a1 := []int{1, 2, 3}
	b1 := a1
	fmt.Println(a1, b1)
	a1[1] = 43
	fmt.Println(a1, b1)

	a2 := map[string]string{"foo": "bar"}
	b2 := a2
	fmt.Println(a2, b2)
	b2["quick"] = "balue"
	fmt.Println(a2, b2)
}
