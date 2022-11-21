package main

import (
	"fmt"
	"goTutorialBasics/arraysAndSlices"
	"goTutorialBasics/channels"
	"goTutorialBasics/constants"
	"goTutorialBasics/deferPanicAndRecover"
	"goTutorialBasics/environmentVariables"
	"goTutorialBasics/errorHandling"
	"goTutorialBasics/fileSystem"
	"goTutorialBasics/functions"
	"goTutorialBasics/goroutines"
	"goTutorialBasics/ifAndSwitchStatements"
	"goTutorialBasics/interfaces"
	"goTutorialBasics/looping"
	"goTutorialBasics/mapsAndStructs"
	"goTutorialBasics/pointersAndScoping"
	"goTutorialBasics/primitives"
	"goTutorialBasics/stringManipulation"
	"goTutorialBasics/variables"
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

	errorHandling.Go()
	fmt.Println()

	stringManipulation.Go()
	fmt.Println()

	fileSystem.Go()
	fmt.Println()
}
