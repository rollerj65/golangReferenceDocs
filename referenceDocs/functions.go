package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	fmt.Printf("Max value is : %d\n", ret)

	c, d := swap("Mahesh", "Kumar")
	fmt.Println(c, d)

	fmt.Printf("Functions using pointers\n\n")
	fmt.Printf("Before pointer swap, value of a : %d\n", a)
	fmt.Printf("Before pointer swap, value of b : %d\n", b)

	pointerSwap(&a, &b)

	fmt.Printf("After pointer swap, value of a : %d\n", a)
	fmt.Printf("After pointer swap, value of b : %d\n", b)

	fmt.Printf("\nFunctions as values\n")
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	fmt.Printf("\nfunction closures\n")
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	fmt.Printf("\nmethods\n")
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())

}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

//use to change the values passed in
func pointerSwap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
