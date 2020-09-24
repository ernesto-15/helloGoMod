package main

import (
	"github.com/ernesto-15/helloGoMod/greet"
	"rsc.io/quote/v3"
	"fmt"
)

func main() {
	fmt.Println(greet.English())
	fmt.Println(quote.HelloV3())
	fmt.Println(quote.Concurrency())
}