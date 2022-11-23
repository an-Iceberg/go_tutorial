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
}
