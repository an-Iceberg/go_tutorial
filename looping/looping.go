package looping

import (
	"fmt"
	"strconv"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  Loops"))

	for i := 0; i < 20; i++ {
		fmt.Print(strconv.Itoa(i) + " ")

		// If i is even divide by 2
		if i%2 == 0 {
			i /= 2
			// Else double it and add 1
		} else {
			i = 2*i + 1
		}
	}

	fmt.Println("")

	for i, j := 0, 0; i <= 10; i, j = i+1, j+2 {
    fmt.Printf("i:%d j:%d\n", i, j)
	}

	fmt.Println("")

	index := 0

	for ; index < 11; index++ {
    fmt.Printf("index:%d\n", index)
	}

	fmt.Println("")

	// While loops in GO
	for i := 0; i < 11; {
		fmt.Println(i)
		i++
	}

	fmt.Println("")

	// Infinite loop
	for {
		index++

		if index > 100 {
			break
		} else {
			continue
		}
	}

	// Labels
Loop:
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			fmt.Print(strconv.Itoa(i*j) + " ")

			if i*j > 50 {
				break Loop
			}
		}
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("")

	// Iteration over arrays/slices/maps
	someSlice := []int{34, 29, 83, 72, 27, 58, 27, 84}

	for _, value := range someSlice {
		fmt.Println(value)
	}
}
