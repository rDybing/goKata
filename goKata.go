// Various small katas to test out Go
//
// v1.2: Added Word CLock
// v1.1: Added CLI menu and gcd
// v1.0: Added FizzBuzz
//
// CCBY roy.dybing@gmail.com 2016

package main

import (
	"fmt"
	"math"
	"time"
)

type wordStruct struct {
	words   [12]string
	minutes [20]string
	hours   [12]string
	output  string
}

type timeStruct struct {
	hour   int
	minute int
}

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
		case "clock":
			wordClock()
		case "help":
			help()
		case "quit":
			quit = true
		}
	}
}

// display the available commands
func help() {
	fmt.Println("Available Commands:")
	fmt.Println(" - fizzbuzz   | the drinking game")
	fmt.Println(" - gcd        | find greatest common denominator")
	fmt.Println(" - clock      | word clock")
	fmt.Println(" - help       | list of commands")
	fmt.Println(" - quit       | exit this application")
}

// get greatest common denominator between two numbers
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

// the infamous fizzbuzz Drinking game
func fizzBuzz() {
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

// word clock, display time as words, not numbers
func wordClock() {

	ws := new(wordStruct)
	ts := new(timeStruct)

	// populate structs
	ws.words[0] = "It is"
	ws.words[1] = "half"
	ws.words[2] = "to"
	ws.words[3] = "past"
	ws.words[4] = "o'clock"
	ws.words[5] = "in the"
	ws.words[6] = "afternoon!"
	ws.words[7] = "noon!"
	ws.words[8] = "midnight!"
	ws.words[9] = "morning!"
	ws.words[10] = "at night!"
	ws.words[11] = "evening!"

	ws.minutes[0] = "one"
	ws.minutes[1] = "two"
	ws.minutes[2] = "three"
	ws.minutes[3] = "four"
	ws.minutes[4] = "five"
	ws.minutes[5] = "six"
	ws.minutes[6] = "seven"
	ws.minutes[7] = "eight"
	ws.minutes[8] = "nine"
	ws.minutes[9] = "ten"
	ws.minutes[10] = "eleven"
	ws.minutes[11] = "twelve"
	ws.minutes[12] = "thirteen"
	ws.minutes[13] = "fourteen"
	ws.minutes[14] = "quarter"
	ws.minutes[15] = "sixteen"
	ws.minutes[16] = "seventeen"
	ws.minutes[17] = "eighteen"
	ws.minutes[18] = "nineteen"
	ws.minutes[19] = "twenty"

	ws.hours[0] = "one"
	ws.hours[1] = "two"
	ws.hours[2] = "three"
	ws.hours[3] = "four"
	ws.hours[4] = "five"
	ws.hours[5] = "six"
	ws.hours[6] = "seven"
	ws.hours[7] = "eight"
	ws.hours[8] = "nine"
	ws.hours[9] = "ten"
	ws.hours[10] = "even"
	ws.hours[11] = "twelve"

	// get stuff done
	getTime(ts)
	showTime(ws, ts)
	fmt.Println(ws.output)
}

// get the time doh!
func getTime(t *timeStruct) {
	now := time.Now()
	t.hour, t.minute, _ = now.Clock()
}

// translate time in numbers to string
func showTime(w *wordStruct, t *timeStruct) {
	plusHour := t.hour

	// out>>it is
	addWordToString(w.words[0], w)

	// out>>minutes
	if t.minute == 0 {
		if t.hour == 0 {
			addWordToString(w.words[8], w)
		} else if t.hour == 12 {
			addWordToString(w.words[7], w)
		} else {
			addWordToString(w.words[4], w)
		}
	} else {
		if t.minute <= 20 { // one .. twenty
			addWordToString(w.minutes[t.minute-1], w)
		} else if t.minute < 30 { // twenty one .. twenty nine
			addWordToString(w.minutes[19], w)
			addWordToString(w.minutes[t.minute-21], w)
		} else if t.minute == 30 { // half
			addWordToString(w.words[1], w)
		} else if t.minute < 40 { // twenty nine .. twenty one
			addWordToString(w.minutes[19], w) // twenty
			addWordToString(w.minutes[60-t.minute-21], w)
		} else { // twenty .. one
			addWordToString(w.minutes[60-t.minute-1], w)
		}
		if t.minute <= 30 { // past
			addWordToString(w.words[3], w)
		} else { // to
			addWordToString(w.words[2], w)
			plusHour++
		}
	}

	// out>>hours
	if !(t.minute == 0 && (t.hour == 0 || t.hour == 12)) {
		// Hours
		if plusHour == 0 {
			addWordToString(w.hours[11], w)
		} else if plusHour <= 12 {
			addWordToString(w.hours[plusHour-1], w)
		} else {
			addWordToString(w.hours[plusHour-13], w)
		}
		if plusHour == 11 || plusHour == 23 {
			//addWordToString(w_el)
		}
		// Time of day
		if t.hour < 6 { // at night
			addWordToString(w.words[10], w)
		} else if t.hour < 12 { // in the morning
			addWordToString(w.words[5], w)
			addWordToString(w.words[9], w)
		} else if t.hour < 18 { // in the afternoon
			addWordToString(w.words[5], w)
			addWordToString(w.words[6], w)
		} else { // at night (again)
			addWordToString(w.words[10], w)
		}
	}
}

// append new word to output string
func addWordToString(addWord string, w *wordStruct) {
	w.output += addWord + " "
}
