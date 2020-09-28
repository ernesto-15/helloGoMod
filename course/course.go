package course

import (
	"fmt"
)

//Course struct
type Course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

//New is the constructor for course struct
func New(name string, price float64, isFree bool) *Course {
	if price <= 0 {
		price = 30
	}
	return &Course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

//SetPrice method. This is a function with a pointer receiver
func (course *Course) SetPrice(price float64) {
	course.price = price
}
//Price getter method of price
func (course *Course) Price() float64{
	return course.price 
}


//SetClasses setter method of classes
func (course *Course) SetClasses(classes map[uint]string) {
	course.classes = classes
}
//Classes getter method of classes
func (course *Course) Classes() map[uint]string {
	return course.classes
}


//Name getter method of name
func (course *Course) Name() string {
	return course.name
}
//SetName setter method of name
func (course *Course) SetName(name string) {
	course.name = name
}


//IsFree getter method of IsFree
func (course *Course) IsFree() bool {
	return course.isFree
}
//SetIsFree setter method of IsFree
func (course *Course) SetIsFree(isFree bool) {
	course.isFree = isFree
}

//UserIDs getter method of UserIDs
func (course *Course) UserIDs() []uint {
	return course.userIDs
}
//SetUserIDs setter method of UserIDs
func (course *Course) SetUserIDs(userIDs []uint) {
	course.userIDs = userIDs
}


//PrintClasses method
func (course *Course) PrintClasses() {
	text := "The classes of the " + course.name + " course " + "are:\n"
	for _, class := range course.classes {
		text += "*) " + class + "\n"
	}
	fmt.Println(text[:len(text)-2])
}

