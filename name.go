package main

import "fmt"

type name struct {
  first string
  last  string
}

func (name name) print() {
  fmt.Println(name.first + " " + name.last)
}
