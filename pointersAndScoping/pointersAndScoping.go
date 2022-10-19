package pointersAndScoping

import (
	"fmt"
	"strings"
)

type myStruct struct {
	foo int
}

func string_replace(string string, replacements map[string]string) string {
	for key, value := range replacements {
		string = strings.Replace(string, key, value, -1)
	}
	return string
}

func Go() {
	fmt.Println("  Pointers and Scoping")

	{
		// Pointers
		a := 42
		var b *int = &a
		fmt.Println(a, *b)
		a = 27
		fmt.Println(a, *b)
		*b = 17
		fmt.Println(a, *b)
	}

	{
		// Pointer arithmetics is not allowed in GO
		a := [3]int{4, 23, 8}
		b := &a[0]
		c := &a[2]
		fmt.Printf("%v %v %p\n", a, *b, c)
	}

	{
		var ms *myStruct
		fmt.Printf("ms(uninitialised):%v\n", ms)
		ms = &myStruct{foo: 42}
		fmt.Printf("ms:%v\n", ms)
		newStruct := new(myStruct)
		fmt.Printf("newStruct:%v\n", newStruct)
		newStruct.foo = 37
		fmt.Printf("newStruct:%v\n", newStruct.foo)
	}

	fmt.Println(string_replace(
		"{name} is {age} years old",
		map[string]string{
			"{name}": "John",
			"{age}":  "24",
		},
	))

	{
		// Arrays are copied, Slices and Maps work with pointers
		a := [3]int{4, 39, 74}
		b := a
		fmt.Println(a, b)
		a[0] = 69
		fmt.Println(a, b)
	}
}
