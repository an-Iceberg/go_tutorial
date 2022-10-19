package deferPanicAndRecover

import (
	"fmt"
)

func Go() {
	fmt.Println("  Defer, Panic and Recover")

	fmt.Println("start")
	defer fmt.Println("middle") // Defers are processed in a lifo (stack) order
	fmt.Println("end")

	/*
	response, error := http.Get("http://www.google.com/robots.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	robots, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s\n", robots)
	*/

	someString := "start"
	defer fmt.Println(someString) // This will print "start" coz defer evaluates variable values eagerly at the time of calling it
	someString = "end"

	// This will cause the runtime to panic
	// a, b := 1, 0
	// answer := a / b
	// fmt.Println(answer)

	// panic("something bad happened") // Panics happen after any defers have been executed
	// recover()
}
