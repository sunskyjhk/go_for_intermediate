package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

// coolFunc take in a function
// as a parameter

func coolFunc(a func()) {
	// it then immediately calls that function
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	// here we call our coolFunc function
	// passing in myFunc
	coolFunc(myFunc)
}
