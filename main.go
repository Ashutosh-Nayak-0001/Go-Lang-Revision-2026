package main

import (
	"fmt"
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

	// var a = []int{1,2,3}
	// fmt.Println(reflect.TypeOf(a))

	//unary operator
	// a := 10
	// a++
	// fmt.Println(a)

	// a := 10
	// b := 20
	// sum := a + b
	// fmt.Println("The sum of", a, "and", b, "is", sum)

	// firstName := "Alice"
	// fmt.Println("Hello,", firstName)
	// firstName = "Bob"
	// fmt.Println("Hello,", firstName)

	// lastName := "Smith"
	// fmt.Printf("Full name: %s\n %s\n", firstName, lastName)

	//Different operators
	// x := 10
	// y := 5
	// fmt.Println("x + y =", x+y)
	// fmt.Println("x - y =", x-y)
	// fmt.Println("x * y =", x*y)
	// fmt.Println("x / y =", x/y)
	// fmt.Println("x % y =", x%y)

	//complex numbers
	// c1 := complex(2, 3)
	// c2 := complex(4, 5)
	// fmt.Println("c1 + c2 =", c1+c2)
	// fmt.Println("c1 - c2 =", c1-c2)
	// fmt.Println("c1 * c2 =", c1*c2)
	// fmt.Println("c1 / c2 =", c1/c2)

	//bits and bitwise operators
	// x := 202  // 0101 in binary
	// y := 141  // 0011 in binary
	// fmt.Println("x & y =", x&y)   // 0001 (bitwise AND)
	// fmt.Println("x | y =", x|y)   // 0111 (bitwise OR)
	// fmt.Println("x ^ y =", x^y)   // 0110 (bitwise XOR)
	// fmt.Println("x &^ y =", x&^y) // 0100 (bit clear)

	//ARRay
	//r a [5]int //Zero value of int is 0
	//b:= [3]int{1,2,3} //literal
	//c := [...]int{4,5,6} // compiler counts length

	//Accessing and modifiying

	//    a[0] = 10
	// fmt.Println(a[0],len(a))

	//iterating
	//  for i, v := range b {
	//     fmt.Println(i, v)
	//  }

	//Slicing an Array/Slices

	//s := []int{1, 2, 3}      // slice literal
	// s2 := make([]int, 5)     // len=5, cap=5
	// s3 := make([]int, 3, 10) // len=3, cap=10

	//    arr := [5]int{10,20,30,40,50}
	//    s1 := arr[1:4]
	//    fmt.Println(s1)

	// s = append(s, 4, 5)            // add elements
	// s = append(s, []int{6,7}...)   // append another slice

	//  dst := make([]int, len(s))
	//  copy(dst, s)                   // copy into dst

	// fmt.Println(s);
	// fmt.Println(dst)

	//2D Slices

	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6,7},
	// }

	// fmt.Println(matrix)
	// fmt.Println(matrix[1][2]) // 6

	//Reverse an array/slice =================================================

	// input := [5]int{1, 2, 3, 4, 5}
	// left := 0
	// right := len(input) - 1

	// for left < right {
	// 	input[left], input[right] = input[right], input[left]
	// 	left++
	// 	right--
	// }
	// fmt.Println(input)

	//find mi/ max =================================================
	// input := []int{1, 12, 53, 84, 85}

	// min := input[0]
	// max := input[0]

	// for i := 1; i < len(input); i++ {
	//     if input[i] > max {
	//         max = input[i]   // ✅ found bigger value
	//     }
	//     if input[i] < min {
	//         min = input[i]   // ✅ found smaller value
	//     }
	// }

	// fmt.Println("Max =", max) // Max = 85
	// fmt.Println("Min =", min) // Min = 1

	//Remove duplicates from a slice =========================================================================
	// input := []int{1, 12, 53, 53, 84, 84, 85}

	// seen := make(map[int]bool)
	// result := []int{}

	// for i := 0; i < len(input); i++ {
	// 	if !seen[input[i]] {
	// 		result = append(result, input[i])
	// 		seen[input[i]] = true
	// 		fmt.Println(seen)
	// 	}
	// }

	// fmt.Println(result)

	// 	input := []int{1,2,3,4,5}

	// 	k := 2

	// 	input = append(input[k:],input[:k]... )
	//   fmt.Println(input)

	// input := [][]int{
	//     {1, 2, 3},
	//     {4, 5},
	//     {6, 7, 8, 9},
	// }

	// result := []int{}

	// for i := 0; i < len(input); i++ {
	//     result = append(result, input[i]...)  // spreads entire row at once!
	// }

	// fmt.Println(result)

	//input := []int{1,2,3,4,5,6,7,8,9}

	// result := contains([]int{1, 2, 3, 4, 5}, 30)
	// fmt.Println(result) // true

	

}

// func contains(array []int, target int) bool {
// 	for i := 0; i < len(array); i++ {
// 		if array[i] == target {
// 			return true
// 		}
// 	}
// 	return false

// }
