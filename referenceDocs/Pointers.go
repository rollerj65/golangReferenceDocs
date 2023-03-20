package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	const MAX int = 3
	var ptr [MAX]*int
	ip = &a

	fmt.Printf("\nPointers in go\n")
	fmt.Printf("Address of a variable: %x\n", &a)
	fmt.Printf("Address stored in ip variable: %x\n", *ip)
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	fmt.Printf("\nNill pointers in go\n")
	fmt.Printf("The value of ptr is : %x\n", ptr)

	fmt.Printf("\nArray of pointers\n")

	ar := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX; i++ {
		ptr[i] = &ar[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}

	fmt.Printf("\npointer to pointer\n")
	a = 3000
	var ptr2 *int
	var pptr **int

	ptr2 = &a

	pptr = &ptr2

	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr2 = %d\n", *ptr2)
	fmt.Printf("Value available at **ptr = %d\n", **pptr)

	fmt.Printf("\nPassing pointers to functions\n")
	a = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */

	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
