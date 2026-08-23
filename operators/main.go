package main

import "fmt"

func main() {

	a := 20
	b := 6

	// Arithmetic
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)

	// Comparison
	fmt.Println("Equal:", a == b)
	fmt.Println("Not Equal:", a != b)
	fmt.Println("Greater:", a > b)
	fmt.Println("Less:", a < b)

	// Logical
	x := true
	y := false

	fmt.Println("AND:", x && y)
	fmt.Println("OR:", x || y)
	fmt.Println("NOT:", !x)

	// Assignment
	n := 10

	n += 5
	fmt.Println("After +=:", n)

	n -= 2
	fmt.Println("After -=:", n)

	n *= 2
	fmt.Println("After *=:", n)

	// Increment
	n++
	fmt.Println("After ++:", n)

	// Bitwise
	fmt.Println("AND:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("AND NOT:", a&^b)

    // in go , bitwise operator work directly on the indivisual bits 
	// of integer values , while shift operators move those bits left or right 



	// Shift
	fmt.Println("Left Shift:", a<<1)
	fmt.Println("Right Shift:", a>>1)
}