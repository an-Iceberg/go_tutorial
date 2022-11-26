package functions

import "fmt"

func anonymousFunctions() {
	fmt.Println(" Anonymous Functions")

	// Lambda/anonymous functions
	func() {
		fmt.Println("Hello GO!")
	}()

	for i := 0; i < 11; i++ {
		func(i int) {
			fmt.Print(i, " ")
		}(i)
	}
	fmt.Println("")

	// Functional patterns
	{
		local_print := func(function func(int, int) int, a, b int)  {
			fmt.Println(function(a, b))
			// someGreeter.greet() // <== This is possible and possibly problematic
		}

		prod := func(a, b int) int { return a * b }
		sum := func(a, b int) int { return a + b }

		local_print(prod, 2, 5)
		local_print(sum, 2, 5)
	}
}
