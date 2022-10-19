package main

import (
	"fmt"
	"go_tutorial_basics/arraysAndSlices"
	"go_tutorial_basics/channels"
	"go_tutorial_basics/constants"
	"go_tutorial_basics/deferPanicAndRecover"
	"go_tutorial_basics/environmentVariables"
	"go_tutorial_basics/functions"
	"go_tutorial_basics/goroutines"
	"go_tutorial_basics/ifAndSwitchStatements"
	"go_tutorial_basics/interfaces"
	"go_tutorial_basics/looping"
	"go_tutorial_basics/mapsAndStructs"
	"go_tutorial_basics/pointersAndScoping"
	"go_tutorial_basics/primitives"
	math "go_tutorial_basics/testing"
	"go_tutorial_basics/variables"
)

func main() {
	variables.Go()
	fmt.Println()

	primitives.Go()
	fmt.Println()

	constants.Go()
	fmt.Println()

	arraysAndSlices.Go()
	fmt.Println()

	mapsAndStructs.Go()
	fmt.Println()

	ifAndSwitchStatements.Go()
	fmt.Println()

	looping.Go()
	fmt.Println()

	deferPanicAndRecover.Go()
	fmt.Println()

	pointersAndScoping.Go()
	fmt.Println()

	functions.Go()
	fmt.Println()

	interfaces.Go()
	fmt.Println()

	goroutines.Go()
	fmt.Println()

	channels.Go()
	fmt.Println()

	environmentVariables.Go()
	fmt.Println()

	fmt.Println(math.Add(1, 2))
}
