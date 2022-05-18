package main

import "fmt"

const PI = 3.14

func main() {

	// the & (ampersand) operator also known as address of operator returns the memory address of a variable.
	name := "Miguel"
	fmt.Println(&name) // -> 0xc000010250

	//** DECLARING AND INITIALIZING POINTERS
	var x int = 2
	// the expression &x means the address of x and creates a pointer to an integer variable,
	// ptr is of type *int, which is pronounced "pointer to int".
	ptr := &x
    _ = ptr
	// prt -> x [2]

	// fmt.Printf("The type of ptr is %T. The address is %v and the value is %v\n", ptr, ptr, *ptr)
	// declaring a pointer without initializing it
	// its zero value is nil

	var ptr1 *float64
	_ = ptr1

	// creating a pointer using new() built-in function.
	p := new(int) // it creates a pointer called p that is a pointer to an int type

	x = 100
	p = &x // initializing p

	fmt.Printf("p is of type %T with value %v\n", p, p) // => p is of type *int with value 0xc00012a008
	//fmt.Printf("address of x is %p\n", &x)              // => address of x is 0xc00012a008

	//** THE DEREFERENCING OPERATOR **//

	// * in front of a pointer is called the dereferencing operator
	*p = 90 //equivalent to x = 90 because *p is x
	// x and *p is the same thing.

	//fmt.Println(*p)                  // => 90
	//fmt.Println("*p == x:", *p == x) // => *p == x: true

	//fmt.Println("Value of x:", *p) // => Value of x: 90 , equivalent to fmt.Println(x)

	*p = 10     // If I write *p = 10, this is equivalent to x = 10
	*p = *p / 2 //dividing x through the pointer
	//fmt.Println(x) // -> 5

	// In a nutshell:
	// &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
	// *pointer => value  -> and if you have pointer you turn it into value value by using the star operator
    xx := 10
    px := &xx
    *px = 2

    fmt.Println(xx) // => 2
}
