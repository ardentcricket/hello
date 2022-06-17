package main

import ("fmt"
		"log"
		"example.com/greetings"
)

import "rsc.io/quote"

func main() {

	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Dwight")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message, quote.Go())
}