package main

import (
	"fmt"
	"reflect"
)

func main() {
	//println("Hello, World!")
	// 	var a int = 10
	// 	var b int = 20
	// 	var sum int = a + b
	// 	println("The sum of", a, "and", b, "is", sum)

	// 	var name string = "Alice"
	// 	println("Hello,", name)
	// 	name = "Bob"
	// 	println("Hello,", name)
	// 	println(strings.Split(name, ""))

	// 	title := "david"
	//    fmt.Printf("The title is: %s\n", title)

	// var a [10]int
	// for i := 0; i < len(a); i++ {
	// 	a[i] = i + 1
	// }
	// fmt.Println(a)

	// var a = []int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	// var a = [...]int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	//sparse array
	// var a = [10]int{0: 1, 2: 3, 4: 5}
	// fmt.Println(a)

	//implicit array length
	// var a = [...]int{1, 2, 3, 4, 5}
	// fmt.Println(a)
	//	fmt.Println(len(a))

	//Accessing array elements
	//	fmt.Println(a[0])

	//2 D array
	// var a = [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// fmt.Println(a)

	//=================================slices======================

	//slices := []int{1, 2, 3, 4, 5}
	//  a := [5]int{1, 2, 3, 4, 5}
	//  slices := a[1:4]
	// fmt.Println(slices)

	// a := []int{1, 2:7, 3, 4:10, 5}
	// fmt.Println(a)

	// a := make([]int, 5,100)
	// a = append(a, 1, 2, 3, 4, 5)
	// fmt.Println(a)

	var a = []int{1,2,3}
	fmt.Println(reflect.TypeOf(a))
}
