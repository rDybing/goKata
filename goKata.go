// Various small katas to test out Go

package main

import (
	"fmt"
)

func main() {

	var text string

	for i := 1; i <= 100; i++ {
		// when mod of i/x is zero...
		switch 0 {
		case i % 15:
			text = "FizzBuzz"
			break
		case i % 5:
			text = "Fizz"
		case i % 3:
			text = "Buzz"
		default:
			text = ""
		}
		fmt.Printf("%v : %s\n", i, text)
	}
}
