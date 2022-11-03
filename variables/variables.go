package variables

import (
	"fmt"
	"strconv"
)

/* Global variables must be declared using the 'var' keyword */

/* Block declaration */
var (
	kilometersToMiles float32 = 0.621_371
	milesToKilometers float32 = 1.609_34
)

func Go() {
	fmt.Println("  Variables")

	var one float32
	one = 69
	one *= kilometersToMiles
	one = 49

	one = milesToKilometers

	var two int = 42

	three := 512. // The . at the end signifies that it's a float, without it it would just be an int

	four := 50

	five := float32(four) // Type casting

	six := strconv.Itoa(four)

	fmt.Printf(
		"one: %v %T\ntwo: %v %T\nthree: %v %T\nfour: %v %T\nfive: %v %T\nsix: %v %T\n",
		one, one,
		two, two,
		three, three,
		four, four,
		five, five,
		six, six,
	)
}
