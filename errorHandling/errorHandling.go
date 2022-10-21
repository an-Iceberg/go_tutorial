package errorHandling

import (
	"errors"
	"fmt"
	"os"
)

func Go() {
	fmt.Println("  Error handling")

	err := errors.New("some new custom error")

	fmt.Println(err.Error())

	someFile, error := os.Open("some_non_existent_file.format")

	if error != nil {
		fmt.Println("File could not be found")
		return
	} else {
		fmt.Println(someFile)
	}
}
