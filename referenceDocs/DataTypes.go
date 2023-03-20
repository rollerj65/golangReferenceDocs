package main

import "fmt"

func main() {
	//static type declaration
	var x float64
	x = 20.0
	fmt.Printf("x: %g\n", x)
	fmt.Printf("x is of type: %T\n", x)

	//dynamic type declaration
	y := 42
	fmt.Printf("x: %g\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("x is of type: %T\n", x)
	fmt.Printf("y is of type: %T\n", y)

	//mixed type declaration

	var a, b, c = 3, 4, "foo"

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

}
