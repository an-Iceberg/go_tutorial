package primitives

import (
	"fmt"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  Primitives"))

	// Booleans
	boolean := true // Uninitialized booleans equate to 'false'
	var uninitialized bool
	test1 := 1 == 1
	test2 := 1 == 2

	fmt.Printf("%v, %T\n", boolean, boolean)
	fmt.Printf("%v, %T\n", uninitialized, uninitialized)
	fmt.Printf("%v, %T\n", test1, test1)
	fmt.Printf("%v, %T\n", test2, test2)
	fmt.Println("")

	// Integer types
	n := 42 // If nothing else is specified, a 32-bit int is used
	// Other int types are: int8, int16, int32, int64
	// Unsigned integers: uint8, uint16, uint32
	fmt.Printf("%v, %T\n", n, n)

	a := 10
	b := 3
	// Arithmetics are only allowed same types
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v / %v = %v\n", a, b, a/b)
	fmt.Printf("%v %% %v = %v\n", a, b, a%b)
	fmt.Println("")

	// FLoating point types (float32, float64)
	someFloat := 459.298476
	anotherFloat := 2.34e23 // Exponential notation (1.4E7 also works)
	fmt.Printf("%v, %T\n", someFloat, someFloat)
	fmt.Printf("%v, %T\n", anotherFloat, anotherFloat)

	c := 10.2
	d := 3.7
	fmt.Printf("%v + %v = %v\n", c, d, c+d)
	fmt.Printf("%v - %v = %v\n", c, d, c-d)
	fmt.Printf("%v * %v = %v\n", c, d, c*d)
	fmt.Printf("%v / %v = %v\n", c, d, c/d)
	fmt.Println("")

	// Complex number type
	var one complex64 = 1 + 2i
	two := 3 - 5i
	fmt.Printf("%v, %T\n", one, one)
	fmt.Printf("%v, %T\n", two, two)
	fmt.Printf("%v + %v = %v\n", one, two, complex128(one)+two)
	fmt.Printf("%v - %v = %v\n", one, two, complex128(one)-two)
	fmt.Printf("%v * %v = %v\n", one, two, complex128(one)*two)
	fmt.Printf("%v / %v = %v\n", one, two, complex128(one)/two)

	fmt.Println("real part:", real(one), "imaginary part:", imag(two))

	three := complex(c, d)
	fmt.Println(three)
	fmt.Println("")

	// String type (as usual, strings can be treated as arrays of characters) (strings in GO represent UTF8 characters)
	someString := "this is a string" // Strings are immutable
	fmt.Printf("%v, %T\n", someString, someString)
	fmt.Println(string(someString[13]))

	anotherString := ". This too is a string"
	fmt.Println(someString + anotherString)

	someBytes := []byte(someString)
	fmt.Println(someBytes)
	fmt.Println("")

	// Rune type (runes in GO represent UTF32 characters)
	var someRune rune = 'a'
	fmt.Printf("%v, %T\n", someRune, someRune)

	// Shorthand operators
	a++
	a--
	a += 1
	a -= 1
	a *= 1
	a /= 1
	a %= 1

	// Large numbers can be written like so
	aLargeNumber := 12_345
	fmt.Println(aLargeNumber)
}
