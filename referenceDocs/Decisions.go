package main

import "fmt"

func main() {
	//if statement
	var a int = 10

	fmt.Printf("if statement\n\n")
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is: %d\n", a)

	//if-else statement
	a = 100

	fmt.Printf("if-else statement\n\n")
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	} else {
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is: %d\n", a)

	//nested if statements
	fmt.Printf("nested if statements\n\n")
	a = 100
	var b int = 100

	/* check boolean condition */
	if a == 100 {
		if b == 200 {
			fmt.Printf("Value of a is 100 and b is 200\n")
		}
	}
	fmt.Printf("Exact value of a is :%d\n", a)
	fmt.Printf("Exact value of b is :%d\n", b)

	//switch statements
	fmt.Printf("Switch statement\n\n")
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is %s\n", grade)

	//select statements
	fmt.Printf("select statements\n\n")
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("received ", i2, "from c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication")
	}
}
