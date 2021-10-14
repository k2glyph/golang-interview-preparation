package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func sendValue(c chan string) {
	fmt.Println("Executing go routine")
	c <- "Hello world"
	fmt.Println("Finished executing channel")
}
func main() {
	// 	srv := api.NewServer()
	// 	http.ListenAndServe(":8080", srv.Router)

	fmt.Println("Concurrency with goroutine")
	// go compute(5)
	// go compute(5)
	// fmt.Scanln()

	values := make(chan string, 2)
	defer close(values)
	go sendValue(values)
	go sendValue(values)
	value := <-values
	fmt.Println(value)
}
