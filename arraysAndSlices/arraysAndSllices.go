package arraysAndSlices

import (
	"fmt"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  Arrays and Slices"))

	// Arrays
	var grades [3]float32 = [3]float32{4.5, 5, 4}
	fmt.Printf("%v, %T\n", grades, grades)

	numbers := [...]int{21, 12, 84, 17, 49, 82, 73}
	fmt.Printf("%v, %T\n", numbers, numbers)

	var students [3]string
	fmt.Printf("%v, %T\n", students, students)
	students[0] = "Lisa"
	students[2] = "Mark"
	fmt.Printf("%v, %T\n", students, students)
	fmt.Println("Number of students:", len(students))
	fmt.Println("")

	identityMatrix := [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(identityMatrix)

	someMatrix := [3][3]int{
		{1, 2, 3},
		{3, 1, 2},
		{3, 2, 1},
	}
	fmt.Println(someMatrix)

	// Arrays are copied by value
	a := [...]int{1, 2, 3} // Length inferracne operator
	b := a
	c := &a
	b[1] = 5
	c[2] = 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("")

	// Slices
	someSlice := []int{1, 2, 3, 4}
	fmt.Println(someSlice)
	fmt.Println("Length of slice:", len(someSlice))
	fmt.Println("Capacity of slice:", cap(someSlice))

	// Slices are copied by reference
	anotherSlice := someSlice
	anotherSlice[1] = 55
	fmt.Println(someSlice)

	data := []int{23, 89, 10, 42, 95, 82, 47, 28, 57, 49}

	// These operators can also work with arrays
	// These numbers refer to array indices
	z := data[:]   // Slice of all elements
	y := data[3:]  // Slice from 4th element to end
	x := data[:6]  // Slice of first 6 elements
	w := data[4:8] // Slice from 5th till 8th element
	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(w)

	slice1 := make([]int, 3, 10)
	fmt.Println(slice1)
	fmt.Printf("Length: %v\n", len(slice1))
	fmt.Printf("Capacity: %v\n", cap(slice1))

	slice1 = append(slice1, 3, 102, 512, 2389)
	fmt.Println(slice1)
	fmt.Printf("Length: %v\n", len(slice1))
	fmt.Printf("Capacity: %v\n", cap(slice1))

	slice2 := []int{23, 24, 25, 938, 2948, 2938}
	slice1 = append(slice1, slice2...) // Spread operator
	fmt.Println(slice1)
	fmt.Printf("Length: %v\n", len(slice1))
	fmt.Printf("Capacity: %v\n", cap(slice1))

	stack := []int{4, 8, 2}
	stack = stack[1:] // A way to 'pop' the top element of the stack
	fmt.Println(stack)

	queue := []int{4, 8, 2}
	queue = queue[:len(queue)-1] // A way to 'dequeue' the front of the queue
	fmt.Println(queue)

	compositeSlice := append(queue, stack...)
	fmt.Println(compositeSlice)
}
