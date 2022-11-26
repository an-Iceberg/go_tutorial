package functions

import "fmt"

func regularFunctions() {
	fmt.Println(" Regular Functions")

  fmt.Printf("Hello %s\n", "John")

	fmt.Println(add(3, 6, 2))

	{ // Scopes :D
		a := 4
		four := &a
		increment(four)
		fmt.Println(*four)
	}

	fmt.Println("The sum is", sum(1, 2, 3, 4, 5))

	fmt.Println(*mult(4, 5))

	fmt.Println(sub(10, 4))

	{
		strings := [3]string{"Hello", "World", "!"}
		a, b, c := someFunction(strings)
		fmt.Println(a, b, c)
	}

	{
		number, error := divide(5.0, 0.0)
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println(number)
		}
	}

	// Functions as variables
	{
		f := func() {
			fmt.Println("Hello GO!")
		}

		f()

		var divide func(float64, float64) (float64, error)
		divide = func(f1, f2 float64) (float64, error) {
			if f2 == 0.0 {
				return 0.0, fmt.Errorf("cannot divide by zero")
			}
			return f1 / f2, nil
		}

		d, err := divide(5.0, 3.0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(d)
	}

	// Methods
	someGreeter := greeter{greeting: "Hello", name: "GO"}
	someGreeter.greet()
	someGreeter.print()
	someGreeter.greetAndChangeName()

}

// Passing in multiple variables of the same type
func add(a, b, c int) int {
	return a + b + c
}

// Passing pointers to functions
func increment(number *int) {
	*number++
}

// Variadic parameters
func sum(values ...int) int {
	fmt.Println(values)

	result := 0
	for _, value := range values {
		result += value
	}

	return result
}

// Returning a local variable as a pointer
func mult(a, b int) *int {
	result := a * b
	return &result
}

// Named return values
func sub(a, b int) (result int) {
	result = a - b
	return
}

// Multiple return values
func someFunction(strings [3]string) (string, string, string) {
	return strings[0], strings[1], strings[2]
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Methods
type greeter struct {
	greeting string
	name     string
}

// The struct is passed by copy by default, making field reassignments ineffective
func (greeter greeter) greet() {
	fmt.Println(greeter.greeting, greeter.name)
	greeter.name = "some other name"
}

// Using the struct as a pointer on the left side allows us to modify the struct fields
func (greeter *greeter) greetAndChangeName() {
	greeter.name = "Ghost"
	fmt.Println(greeter.greeting, greeter.name)
}

func (greeter greeter) print() {
	fmt.Printf("Greeting: %s, Name: %s\n", greeter.greeting, greeter.name)
}
