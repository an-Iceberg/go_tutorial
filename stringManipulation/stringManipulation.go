package stringManipulation

import (
	"fmt"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  String Manipulation"))

	message := fmt.Sprintf("%d + %d = %d", 2, 3, 5)

	fmt.Println(message)

	fmt.Println(strings.Split(message, " "))

	message = strings.ReplaceAll(message, "=", ":")

	fmt.Println(message)

	number := 74

  // Formatting with positioned arguments
	fmt.Printf("%%v:%[1]v %%+v:%+[1]v %%#v:%#[1]v %%T:%[1]T\n", number)
	fmt.Printf("%#[1]b 0d%[1]d %#[1]o %[1]O %#[1]x %#[1]X %[1]U\n", number)
}
