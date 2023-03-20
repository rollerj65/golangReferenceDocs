package main

import "fmt"

func main() {

	fmt.Printf("\nlen() and cap() functions\n")
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	fmt.Printf("\n\nnil slice functions\n")

	var emptyNumbers []int
	printSlice(emptyNumbers)

	if emptyNumbers == nil {
		fmt.Printf("slice is nil")
	}

	fmt.Printf("\n\nsub slices\n")
	numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* print the original slice */
	fmt.Println("numbers ==", numbers)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbers[:2]
	printSlice(number2)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbers[2:5]
	printSlice(number3)

	fmt.Printf("\n\nappend() and copy() functions\n")

	var appendAndCopyNumbers []int
	printSlice(appendAndCopyNumbers)

	/* append allows nil slice */
	appendAndCopyNumbers = append(appendAndCopyNumbers, 0)
	printSlice(appendAndCopyNumbers)

	/* add one element to slice*/
	appendAndCopyNumbers = append(appendAndCopyNumbers, 1)

	/* add more than one element at a time*/
	appendAndCopyNumbers = append(appendAndCopyNumbers, 2, 3, 4)
	printSlice(appendAndCopyNumbers)

	/* create a slice numbers1 with double the capacity of earlier slice*/
	doubleCapacitySlice := make([]int, len(appendAndCopyNumbers), (cap(appendAndCopyNumbers))*2)

	/* copy content of numbers to numbers1 */
	copy(doubleCapacitySlice, appendAndCopyNumbers)
	printSlice(appendAndCopyNumbers)

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
