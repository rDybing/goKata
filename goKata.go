// Various small katas to test out Go
//
// v1.1: Added CLI menu and gcd
// v1.0: Added FizzBuzz
//
// CCBY roy.dybing@gmail.com 2016

package main

import (
	"fmt"
	"math"
)

func main() {

	var input string
	var quit bool

	help()

	for quit == false {

		fmt.Printf("Go!:> ")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "fizzbuzz":
			fizzBuzz()
		case "gcd":
			gcd()
		case "help":
			help()
		case "quit":
			quit = true
		}
	}
}

func help() {
	// display the available commands
	fmt.Println("Available Commands:")
	fmt.Println(" - fizzbuzz   | the drinking game")
	fmt.Println(" - gcd        | find greatest common denominator")
	fmt.Println(" - help       | list of commands")
	fmt.Println(" - quit       | exit this application")
	return
}

func gcd() {

	var high, low, first, second, counter, gcdenom int
	var exit bool

	counter = 0

	fmt.Printf("First (0..999):> ")
	fmt.Scanf("%3d\n", &first)
	fmt.Printf("Second(0..999):> ")
	fmt.Scanf("%3d\n", &second)

	if first > second {
		high = first
		low = second
	} else {
		high = second
		low = first
	}

	for exit != true {
		if high%2 == 0 && low%2 == 0 {
			high = high / 2
			low = low / 2
			fmt.Printf("%3d - %3d\n", high, low)
			counter++
		} else {
			exit = true
		}
		if counter > 99 {
			exit = true
		}
	}

	for high != low {
		switch {
		case high%2 == 0:
			high = high / 2
		case low%2 == 0:
			low = low / 2
		case high > low:
			high = (high - low) / 2
		default:
			low = (low - high) / 2
		}
	}
	gcdenom = round(float64(high) * math.Pow(2, float64(counter)))
	fmt.Printf("Greatest Common Denominator: %3d\n", gcdenom)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func fizzBuzz() {
	// fizzbuzz Drinking game
	var text string

	for i := 1; i <= 100; i++ {
		// when mod of i/x is zero...
		switch 0 {
		case i % 15:
			text = "FizzBuzz"
		case i % 5:
			text = "Fizz"
		case i % 3:
			text = "Buzz"
		default:
			text = ""
		}
		fmt.Printf("%3d : %s\n", i, text)
	}
}
