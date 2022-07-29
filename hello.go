package main

import (
	"fmt"
	// "log"
	"time"

	// "example.com/greetings"
	// "rsc.io/quote"
)

// func main() {

// 	log.SetPrefix("grettings: ")
// 	log.SetFlags(0)

// 		c := make(chan string)
// 		go count("fish", c)

// 		for msg := range c {
// 			fmt.Println(msg)
// 		}

// 		message, err := greetings.Hello("Dwight")

// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		fmt.Println(message, quote.Go())
	

// }

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }



//make the channel

func main() {
	c := make(chan string)
	go count("lion", c)

	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println(shouldEscape('?'))
	fmt.Println(calculator(20,5))
}
	
func count(thing string, c chan string) {
		c <- thing
		time.Sleep(time.Millisecond * 500)
		close(c)

}

func shouldEscape(input byte) bool {
	switch input {
	case '?', '&', ' ', '=', '#', '+', '%':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}

func calculator( a int, b int) (mul int, div int) {
	mul = a * b
	div = a / b
	return
}
