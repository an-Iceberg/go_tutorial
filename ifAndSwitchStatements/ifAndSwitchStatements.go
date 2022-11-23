package ifAndSwitchStatements

import (
	"fmt"
	"math"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  If and Switch Statements"))

	// Control flow statements
	if true { fmt.Println("This is true") }

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	// Initializer syntax
	if n, ok := numbers["two"]; ok { fmt.Println(n) }

	number := 50
	guess := 30
	if guess < 1 || guess > 100 { fmt.Println("The guess must be between 1 and 100") }

	if guess >= 1 && guess <= 100 {
		if guess < number { fmt.Println("Too low") }
		if guess > number { fmt.Println("Too high") }
		if guess == number { fmt.Println("Correct guess") }
	}
	fmt.Println("number <= guess:", number <= guess, " number >= guess:", number >= guess, " number != guess:", number != guess)

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	} else {
		if guess < number { fmt.Println("Too low")
		} else if guess > number { fmt.Println("Too high")
		} else { fmt.Println("Correct guess") }
	}

	// Comparing two floats to a defined accuracy
	myNum := 0.123456789
	errorParameter := 0.001
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < errorParameter {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// Switch statements
	// By default in GO case statements prevent fallthrough
	switch i := 2 * 3; i {
		case 1, 5, 10:
			fmt.Println("one, five or ten")

			// One can also break out of a switch statement like so
			if i == 10 { break }

			fmt.Println("This will print too")
		case 2, 4, 6: fmt.Println("two, four or six")
		default: fmt.Println("another number")
	}

	// Tag list syntax
	i := 10
	switch {
		case i <= 10: fmt.Println("Less than or equal to 10")
			fallthrough // Logicless, will execute the next case regardless if the logic test fails or not

		case i <= 20: fmt.Println("Less than or equal to 20")

		default: fmt.Println("Greater than 20")
	}

	// Type switch
	var j interface{} = 1
	switch j.(type) {
		case int: fmt.Println("integer")

		case float32: fmt.Println("float32")

		case string: fmt.Println("string")

		default: fmt.Println("different type")
	}
}
