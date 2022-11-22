package main

import "fmt"

type address struct {
  street      string
  houseNumber int
  zipCode     int
  city        string
}

func (address address) print() {
  fmt.Println(address.street + " " + fmt.Sprint(address.houseNumber))
  fmt.Println(fmt.Sprint(address.zipCode) + " " + address.city)
}
