package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    1,
		actorName: "Daniel",
		companions: []string{
			"cricket",
		},
	}
	fmt.Println("Doctor struct: ", aDoctor)
	fmt.Println("Doctor name: ", aDoctor.actorName)
	fmt.Println("Doctor name: ", aDoctor.companions[0])

	bDoctor := struct{ name string }{name: "dinesh"}
	fmt.Println(bDoctor)
	cDoctor := &bDoctor
	cDoctor.name = "Ramesh"
	fmt.Println(bDoctor)
	fmt.Println(cDoctor)

	b := Bird{}
	b.Name = "parrot"
	b.Origin = "jungle"
	b.SpeedKPH = 48
	b.CanFly = true
	fmt.Println("Bird: ", b)

	bc := Bird{
		Animal:   Animal{Name: "test", Origin: "hel"},
		SpeedKPH: 48,
		CanFly:   true,
	}
	fmt.Println(bc)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
