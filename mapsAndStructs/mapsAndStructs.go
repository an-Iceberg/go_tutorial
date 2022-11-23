package mapsAndStructs

import (
	"fmt"
	"reflect"
	"strings"
)

// Capitalized names are exported out of the package, making them globally accessible to any user
type doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

// GO does not use OOP, it uses composition instead
type animal struct {
	// Tags
	name   string `required max:"100"`
	origin string
}

type bird struct {
	animal   // Embedding the 'animal' struct into the 'bird' struct
	speedKpH float32
	canFly   bool
}

func Go() {
	fmt.Println(strings.ToUpper("  Maps and Structs"))

	// ! Maps are passed around as reference
	// Creating maps
	countryPopulations := map[string]int{
		"Honduraas":        9_904_607,
		"Seychelles":       9_877,
		"Vietnam":          97_338_579,
		"Tajikistan":       9_660_351,
		"Belarus":          9_449_323,
		"Austria":          9_006_398,
		"Papua New Guinea": 8_947_024,
		"Serbia":           8_737_371,
		"Israel":           8_655_535,
		"Iran":             83_992_949,
		"Germany":          83_783_942,
		"Hong Kong":        7_496_981,
	}

	someMap := make(map[string]int)

	fmt.Println(countryPopulations)
	// Pulling out a specific value
	fmt.Println(countryPopulations["Vietnam"])

	// Adding a new key-value-pair
	countryPopulations["Congo"] = 5_518_087

	fmt.Println(countryPopulations["Congo"])

	// Deleting a specific key-value-pair
	delete(countryPopulations, "Seychelles")

	// ! This throws no error
	fmt.Println(countryPopulations["Seychelles"])

	// Checking if element is present
	population, ok := countryPopulations["Seychelles"]
	fmt.Printf("population:%v ok:%v\n", population, ok)

	if ok {
		fmt.Println(countryPopulations["Seychelles"])
	}

	fmt.Println("Length of countryPopulations:", len(countryPopulations))

	fmt.Println(someMap)
	fmt.Println("")

	// Structs
	docOne := doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(docOne)
	fmt.Println(docOne.actorName)
	fmt.Println(docOne.companions[1])

	// Anonymous struct (structs are independed datasets, copies are passed by copy)
	aDoctor := struct{ name string }{name: "John Pertwee"}
	anotherDoctor := aDoctor
	someDoctor := &aDoctor
	someDoctor.name = "some name"
	anotherDoctor.name = "Tom Baker"
	fmt.Println("aDoctor:", aDoctor)
	fmt.Println("someDoctor:", someDoctor)
	fmt.Println("anotherDoctor:", anotherDoctor)
	fmt.Println(aDoctor.name)
	fmt.Println("")

	// Composition
	emu := bird{}
	emu.name = "Emu"
	emu.origin = "Australia"
	emu.speedKpH = 48.
	emu.canFly = false
	fmt.Println(emu)

	emperorPenguin := bird{
		animal: animal{
			name:   "Emperor Penguin",
			origin: "Antarctica",
		},
		speedKpH: 9.,
		canFly:   false,
	}
	fmt.Println(emperorPenguin)
	fmt.Println("")

	// Tags
	t := reflect.TypeOf(animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field)
}
