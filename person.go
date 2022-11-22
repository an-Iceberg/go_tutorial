package main

import "fmt"

type person struct {
  name
  height  int
  weight  int
  gender  string
  address
}

func (person person) print() {
  person.name.print()
  fmt.Println(fmt.Sprint(person.height) + " cm")
  fmt.Println(fmt.Sprint(person.weight) + " kg")
  fmt.Println(person.gender)
  person.address.print()
}
