package main

import (
	"fmt"
	"github.com/ernesto-15/helloGoMod/course"
)

func main() {

	// fmt.Println(greet.English())
	// fmt.Println(quote.HelloV3())
	// fmt.Println(quote.Concurrency())

	goCourse := course.New("Go from zero to hero", 45.3, false)

  classes := map[uint]string{
    1: "Intro",
    2: "Struct",
    3: "Maps",
  }
  goCourse.SetClasses(classes)
  userIDs := []uint {1,2,3,4,5}
  goCourse.SetUserIDs(userIDs)
  goCourse.SetIsFree(true)
  goCourse.SetName("Go for dummies")
  goCourse.SetPrice(69.99)

  fmt.Println("Name: ", goCourse.Name())
  fmt.Println("Price: ", goCourse.Price())
  fmt.Println("Classes: ", goCourse.Classes())
  fmt.Println("UserIDs: ", goCourse.UserIDs())
  fmt.Println("IsFree: ", goCourse.IsFree())
  
  fmt.Printf("%T\n%+v", goCourse, goCourse)

  
  // nose := new(course.Course)
  // fmt.Printf("%T --- %+v", nose, nose)
}
