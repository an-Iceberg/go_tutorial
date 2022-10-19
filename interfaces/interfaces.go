package interfaces

import "fmt"

// Interfaces can be composed of multiple interfaces

// Interfaces describe behaviour (as opposed to data (struct))
type writer interface {
	write([]byte) (int, error)
}

type console_writer struct{}

// Implicit implementation of a function
func (console_writer console_writer) write(data []byte) (int, error) {
	number, error := fmt.Println(string(data))
	return number, error
}

type incrementer interface {
	increment() int
}

type intCounter int

// Any custom type can implement interfaces
func (ic *intCounter) increment() int {
	*ic++
	return int(*ic)
}

func Go() {
	fmt.Println("  Interfaces")

	// Interfaces
	some_writer := console_writer{}
	some_writer.write([]byte("Hello GO!"))

	myInt := intCounter(0)
	var inc incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.increment())
	}

	// Empty interface
	var myObj interface{} = intCounter(0)

	fmt.Println("myObj:", myObj)
}
