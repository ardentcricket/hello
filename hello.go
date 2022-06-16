package main

import ("fmt"
		"example.com/greetings"
)

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Dwight")
	fmt.Println(message)
}