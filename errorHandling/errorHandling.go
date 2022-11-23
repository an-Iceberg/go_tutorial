package errorHandling

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  Error handling"))

	err := errors.New("some new custom error")

	fmt.Println(err.Error())

	someFile, error := os.Open("some_non_existent_file.format")

	if error != nil {
		fmt.Println("File path could not be found")
		return
	} else {
		fmt.Println(someFile)
	}
}
