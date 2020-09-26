package main

import (
	"fmt"
	"github.com/ernesto-15/helloGoMod/course"
)

func main() {

	// fmt.Println(greet.English())
	// fmt.Println(quote.HelloV3())
	// fmt.Println(quote.Concurrency())

	goCourse := course.Course{
		IsFree:  false,
		Name:    "Go from zero to hero",
		Price:   12.34,
		UserIDs: []uint{12, 56, 32},
		Classes: map[uint]string{
			1: "Intro",
			2: "Struct",
			3: "Maps",
		},
	}

	nose := new(course.Course)

	goCourse.PrintClasses()
	goCourse.ChangePrice(104.5)
	fmt.Println(goCourse.Price)
	fmt.Printf("%T --- %+v", nose, nose)

}