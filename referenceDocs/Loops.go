package main

import "fmt"

func main() {
	//for loops
	fmt.Printf("for loops\n\n")

	var b int = 5
	var a int
	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	//nested for loop
	//prime number generator
	fmt.Printf("nested for loop\n\n")

	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // if factor found, not prime
			}
		}
	}
	if j > (i / j) {
		fmt.Printf("%d is prime\n", i)
	}

	//break statements
	fmt.Printf("break statements\n\n")

	a = 10
	for a < 20 {
		fmt.Printf("value of a: %d\n", a)
		a++
		if a > 15 {
			//terminate loop
			break
		}
	}

	//continue statements
	fmt.Printf("continue statements\n\n")

	a = 10
	for a < 20 {
		if a == 15 {
			//skip iteration
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
