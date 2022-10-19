package constants

import "fmt"

const a int16 = 27

func Go() {
	fmt.Println("  Constants")

	const my_constant_value int = 10
	fmt.Printf("%v, %T\n", my_constant_value, my_constant_value)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.141
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	const e = 43
	var f int8 = 27
	fmt.Printf("%v %T\n", e+f, e+f)
	fmt.Println("")

	// Enumerated constants
	const (
		g = iota // 'iota' is a special counter (a new constant block would reset 'iota')
		h
		i
		j
		k
	)
	fmt.Printf("%v %T\n", g, g)
	fmt.Printf("%v %T\n", h, h)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)
	fmt.Println("")

	const (
		_ = iota // _ is a value discarder
		l
		m
		n
	)
	fmt.Printf("%v %T\n", l, l)
	fmt.Printf("%v %T\n", m, m)
	fmt.Printf("%v %T\n", n, n)

}
