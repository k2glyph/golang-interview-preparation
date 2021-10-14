package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	defer fmt.Println("start")
	defer fmt.Println("middle") // execute at the end of the function
	defer fmt.Println("end")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	res, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	a := "start"
	defer fmt.Println("print a value", a)
	a = "end"

	// a1, b1 := 1, 0
	// ans := a1 / b1
	// fmt.Println(ans)
	// panic("something wrong happen")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello golang!"))
	// })
	// err = http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }
	panniker()

}
func panniker() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("Something wrong happen")
}
