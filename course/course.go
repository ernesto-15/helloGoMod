package course

import (
	"fmt"
)

//Course struct
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//PrintClasses method
func (course Course) PrintClasses() {
	text := "The classes of the " + course.Name + " course " + "are:\n"
	for _, class := range course.Classes {
		text += "*) " + class + "\n"
	}
	fmt.Println(text[:len(text)-2])
}

//ChangePrice method. This is a function with a pointer receiver
func (course *Course) ChangePrice(price float64) {
	course.Price = price
}