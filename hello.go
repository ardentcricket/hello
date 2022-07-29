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
}
	
func count(thing string, c chan string) {
		c <- thing
		time.Sleep(time.Millisecond * 500)
		close(c)

}

func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}

