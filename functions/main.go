/*
============================================================
GO FUNCTIONS — 60 INTERVIEW QUESTIONS WITH DEEP EXPLANATION
============================================================

How to use:
    go run functions_60.go

IMPORTANT:
This is one executable file. Each example is implemented as
a separate function named example01() ... example60().
The main() function calls all of them.

TOPICS COVERED:
1. Basic functions
2. Parameters and return values
3. Multiple return values
4. Named returns
5. error returns
6. Variadic functions
7. Functions as values
8. Functions as arguments
9. Anonymous functions
10. Closures
11. Recursion
12. Functions + slices
13. Higher-order functions
14. defer
15. Practical Go/backend-style functions
============================================================
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
============================================================
Q01. Create a function that prints Hello.
============================================================

Concept:
A function is a reusable block of code.

Syntax:

    func functionName() {
        // code
    }

A function is not executed just because we define it.
We need to CALL it.

    sayHello()
*/
func example01() {
	fmt.Println("\n========== Q01: Basic Function ==========")

	sayHello := func() {
		fmt.Println("Hello")
	}

	sayHello()

	/*
		Output:
		Hello

		Interview point:
		In Go, the function body is executed when the function
		is called.
	*/
}

/*
============================================================
Q02. Function with a parameter.
============================================================

A parameter allows us to send data into a function.

    func greet(name string)

Here:
    name   -> parameter
    string -> parameter type

When calling:
    greet("Ashutosh")

"Ashutosh" is an argument.
*/
func example02() {
	fmt.Println("\n========== Q02: Parameter ==========")

	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet("Ashutosh")
}

/*
============================================================
Q03. Function with two parameters.
============================================================

Go allows multiple parameters.

Instead of:

    func add(a int, b int)

we normally write:

    func add(a, b int)

Both are valid.
*/
func example03() {
	fmt.Println("\n========== Q03: Multiple Parameters ==========")

	add := func(a, b int) {
		fmt.Println("Sum:", a+b)
	}

	add(10, 20)
}

/*
============================================================
Q04. Function returning a value.
============================================================

A function can return data.

    func add(a, b int) int

The final int means:
"The function returns one int value."

return sends the result back to the caller.
*/
func example04() {
	fmt.Println("\n========== Q04: Return Value ==========")

	add := func(a, b int) int {
		return a + b
	}

	result := add(10, 20)

	fmt.Println("Result:", result)
}

/*
============================================================
Q05. Subtract two numbers.
============================================================

This demonstrates that the returned value can be stored
in a variable or used directly.
*/
func example05() {
	fmt.Println("\n========== Q05: Subtract ==========")

	subtract := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtract(20, 8))
}

/*
============================================================
Q06. Multiply two numbers.
============================================================
*/
func example06() {
	fmt.Println("\n========== Q06: Multiply ==========")

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println(multiply(10, 5))
}

/*
============================================================
Q07. Divide two numbers.
============================================================

This simple example assumes b != 0.

Later examples show the correct production-style way
to handle errors such as division by zero.
*/
func example07() {
	fmt.Println("\n========== Q07: Divide ==========")

	divide := func(a, b int) int {
		return a / b
	}

	fmt.Println(divide(20, 5))
}

/*
============================================================
Q08. Check even or odd.
============================================================

A function does not have to return a number.
It can return bool.

    true  -> even
    false -> odd
*/
func example08() {
	fmt.Println("\n========== Q08: Even or Odd ==========")

	isEven := func(num int) bool {
		return num%2 == 0
	}

	fmt.Println("10 is even:", isEven(10))
	fmt.Println("7 is even:", isEven(7))
}

/*
============================================================
Q09. Check positive, negative, or zero.
============================================================
*/
func example09() {
	fmt.Println("\n========== Q09: Positive / Negative / Zero ==========")

	checkNumber := func(num int) string {
		if num > 0 {
			return "Positive"
		}

		if num < 0 {
			return "Negative"
		}

		return "Zero"
	}

	fmt.Println(checkNumber(-10))
	fmt.Println(checkNumber(10))
	fmt.Println(checkNumber(0))
}

/*
============================================================
Q10. Find maximum of two numbers.
============================================================

A function can contain normal if/else logic.
*/
func example10() {
	fmt.Println("\n========== Q10: Maximum of Two ==========")

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	fmt.Println("Maximum:", max(10, 20))
}

/*
============================================================
Q11. Find square of a number.
============================================================
*/
func example11() {
	fmt.Println("\n========== Q11: Square ==========")

	square := func(num int) int {
		return num * num
	}

	fmt.Println("Square:", square(5))
}

/*
============================================================
Q12. Find cube of a number.
============================================================
*/
func example12() {
	fmt.Println("\n========== Q12: Cube ==========")

	cube := func(num int) int {
		return num * num * num
	}

	fmt.Println("Cube:", cube(3))
}

/*
============================================================
Q13. Find maximum of three numbers.
============================================================

This is a common beginner interview question.

Algorithm:
1. Assume a is largest.
2. Compare b with largest.
3. Compare c with largest.
4. Return largest.
*/
func example13() {
	fmt.Println("\n========== Q13: Maximum of Three ==========")

	maxOfThree := func(a, b, c int) int {
		max := a

		if b > max {
			max = b
		}

		if c > max {
			max = c
		}

		return max
	}

	fmt.Println("Maximum:", maxOfThree(10, 50, 30))
}

/*
============================================================
Q14. Factorial using a function.
============================================================

5! = 5 * 4 * 3 * 2 * 1 = 120

Algorithm:
result = 1
for i = 1 to n:
    result = result * i
*/
func example14() {
	fmt.Println("\n========== Q14: Factorial ==========")

	factorial := func(n int) int {
		result := 1

		for i := 1; i <= n; i++ {
			result *= i
		}

		return result
	}

	fmt.Println("5! =", factorial(5))
}

/*
============================================================
Q15. Check whether a number is prime.
============================================================

Prime:
A number greater than 1 that has only two factors:
1 and itself.

Example:
29 -> prime

We try dividing by numbers from 2 up to num-1.
If any number divides it exactly, it is not prime.

Interview improvement:
You can optimize this by checking only i*i <= num.
*/
func example15() {
	fmt.Println("\n========== Q15: Prime Check ==========")

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}

		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}

		return true
	}

	fmt.Println("29:", isPrime(29))
	fmt.Println("30:", isPrime(30))
}

/*
============================================================
Q16. Reverse a number.
============================================================

Example:
12345 -> 54321

Important digit technique:

digit := num % 10
reverse = reverse*10 + digit
num = num / 10
*/
func example16() {
	fmt.Println("\n========== Q16: Reverse Number ==========")

	reverseNumber := func(num int) int {
		reverse := 0

		for num != 0 {
			digit := num % 10
			reverse = reverse*10 + digit
			num /= 10
		}

		return reverse
	}

	fmt.Println("Reverse:", reverseNumber(12345))
}

/*
============================================================
Q17. Check palindrome number.
============================================================

A palindrome reads the same forward and backward.

121 -> palindrome
123 -> not palindrome

We reverse the number and compare it with original.
*/
func example17() {
	fmt.Println("\n========== Q17: Palindrome ==========")

	isPalindrome := func(num int) bool {
		original := num
		reverse := 0

		for num != 0 {
			digit := num % 10
			reverse = reverse*10 + digit
			num /= 10
		}

		return original == reverse
	}

	fmt.Println("121:", isPalindrome(121))
	fmt.Println("123:", isPalindrome(123))
}

/*
============================================================
Q18. Sum of digits.
============================================================

12345 -> 1+2+3+4+5 = 15
*/
func example18() {
	fmt.Println("\n========== Q18: Sum of Digits ==========")

	sumOfDigits := func(num int) int {
		sum := 0

		for num != 0 {
			digit := num % 10
			sum += digit
			num /= 10
		}

		return sum
	}

	fmt.Println("Sum:", sumOfDigits(12345))
}

/*
============================================================
Q19. Count digits.
============================================================

12345 has 5 digits.

Each division by 10 removes one digit.
*/
func example19() {
	fmt.Println("\n========== Q19: Count Digits ==========")

	countDigits := func(num int) int {
		if num == 0 {
			return 1
		}

		count := 0

		for num != 0 {
			count++
			num /= 10
		}

		return count
	}

	fmt.Println("Digits:", countDigits(12345))
}

/*
============================================================
Q20. Find first digit.
============================================================

Example:
12345

Repeatedly divide by 10 until only one digit remains.
*/
func example20() {
	fmt.Println("\n========== Q20: First Digit ==========")

	firstDigit := func(num int) int {
		if num < 0 {
			num = -num
		}

		for num >= 10 {
			num /= 10
		}

		return num
	}

	fmt.Println("First digit:", firstDigit(12345))
}

/*
============================================================
Q21. Multiple return values.
============================================================

Go supports returning multiple values directly.

func calculate(a,b int) (int,int)

This is extremely common in Go.
*/
func example21() {
	fmt.Println("\n========== Q21: Multiple Return Values ==========")

	calculate := func(a, b int) (int, int) {
		return a + b, a - b
	}

	sum, difference := calculate(20, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}

/*
============================================================
Q22. Quotient and remainder.
============================================================

17 / 5:
quotient  = 3
remainder = 2

This is a common interview example for multiple returns.
*/
func example22() {
	fmt.Println("\n========== Q22: Quotient + Remainder ==========")

	divide := func(a, b int) (int, int) {
		return a / b, a % b
	}

	quotient, remainder := divide(17, 5)

	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}

/*
============================================================
Q23. Return minimum and maximum of a slice.
============================================================
*/
func example23() {
	fmt.Println("\n========== Q23: Min + Max ==========")

	minMax := func(numbers []int) (int, int) {
		if len(numbers) == 0 {
			return 0, 0
		}

		min := numbers[0]
		max := numbers[0]

		for _, num := range numbers {
			if num < min {
				min = num
			}

			if num > max {
				max = num
			}
		}

		return min, max
	}

	numbers := []int{10, 5, 30, 2, 20}

	min, max := minMax(numbers)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

/*
============================================================
Q24. Return even and odd counts.
============================================================
*/
func example24() {
	fmt.Println("\n========== Q24: Even + Odd Count ==========")

	countEvenOdd := func(numbers []int) (int, int) {
		even := 0
		odd := 0

		for _, num := range numbers {
			if num%2 == 0 {
				even++
			} else {
				odd++
			}
		}

		return even, odd
	}

	even, odd := countEvenOdd([]int{1, 2, 3, 4, 5, 6})

	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)
}

/*
============================================================
Q25. Function returning value + error.
============================================================

This is one of the MOST IMPORTANT Go patterns.

Instead of throwing exceptions for normal application errors,
Go commonly returns:

    result, error

nil means no error.

*/
func example25() {
	fmt.Println("\n========== Q25: Value + Error ==========")

	divide := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}

		return a / b, nil
	}

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

/*
============================================================
Q26. Named return values.
============================================================

Go allows return variables to be named.

    (sum int, difference int)

Then we can use a naked return.

Named returns are valid Go, but don't overuse naked returns
because they can make large functions harder to understand.
*/
func example26() {
	fmt.Println("\n========== Q26: Named Returns ==========")

	calculate := func(a, b int) (sum int, difference int) {
		sum = a + b
		difference = a - b

		return
	}

	sum, difference := calculate(20, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}

/*
============================================================
Q27. Named return with error.
============================================================
*/
func example27() {
	fmt.Println("\n========== Q27: Named Return + Error ==========")

	divide := func(a, b float64) (result float64, err error) {
		if b == 0 {
			err = errors.New("division by zero")
			return
		}

		result = a / b
		return
	}

	result, err := divide(20, 5)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

/*
============================================================
Q28. Variadic function.
============================================================

A variadic function accepts zero or more arguments.

    func sum(numbers ...int)

Inside the function, numbers behaves like []int.

Examples:
    sum()
    sum(1)
    sum(1,2,3)
*/
func example28() {
	fmt.Println("\n========== Q28: Variadic Function ==========")

	sum := func(numbers ...int) int {
		total := 0

		for _, num := range numbers {
			total += num
		}

		return total
	}

	fmt.Println("Sum 1:", sum(1, 2, 3))
	fmt.Println("Sum 2:", sum(10, 20, 30, 40))
	fmt.Println("Sum 3:", sum())
}

/*
============================================================
Q29. Variadic maximum.
============================================================
*/
func example29() {
	fmt.Println("\n========== Q29: Variadic Maximum ==========")

	max := func(numbers ...int) int {
		if len(numbers) == 0 {
			return 0
		}

		result := numbers[0]

		for _, num := range numbers {
			if num > result {
				result = num
			}
		}

		return result
	}

	fmt.Println("Maximum:", max(10, 20, 5, 40, 15))
}

/*
============================================================
Q30. Count variadic arguments.
============================================================
*/
func example30() {
	fmt.Println("\n========== Q30: Count Arguments ==========")

	countArguments := func(numbers ...int) int {
		return len(numbers)
	}

	fmt.Println("Count:", countArguments(1, 2, 3, 4, 5))
}

/*
============================================================
Q31. Pass a slice to a variadic function.
============================================================

If numbers is already []int:

    numbers := []int{1,2,3}

Use:

    sum(numbers...)

The ... expands the slice into individual arguments.
*/
func example31() {
	fmt.Println("\n========== Q31: Slice to Variadic ==========")

	sum := func(numbers ...int) int {
		total := 0

		for _, num := range numbers {
			total += num
		}

		return total
	}

	numbers := []int{10, 20, 30, 40}

	fmt.Println("Sum:", sum(numbers...))
}

/*
============================================================
Q32. Store a function in a variable.
============================================================

Functions are first-class values in Go.

That means a function can be:
- stored in a variable
- passed to another function
- returned from another function
*/
func example32() {
	fmt.Println("\n========== Q32: Function as Value ==========")

	add := func(a, b int) int {
		return a + b
	}

	operation := add

	fmt.Println("Result:", operation(10, 20))
}

/*
============================================================
Q33. Pass a function as an argument.
============================================================

A function can accept another function as a parameter.

This is called a higher-order function.
*/
func example33() {
	fmt.Println("\n========== Q33: Function as Argument ==========")

	add := func(a, b int) int {
		return a + b
	}

	calculate := func(
		a, b int,
		operation func(int, int) int,
	) int {
		return operation(a, b)
	}

	fmt.Println("Result:", calculate(10, 20, add))
}

/*
============================================================
Q34. Calculator using function as parameter.
============================================================

This demonstrates the same operation function being reused
with different behaviors.
*/
func example34() {
	fmt.Println("\n========== Q34: Function-Based Calculator ==========")

	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	calculate := func(
		a, b int,
		operation func(int, int) int,
	) int {
		return operation(a, b)
	}

	fmt.Println("Add:", calculate(10, 5, add))
	fmt.Println("Subtract:", calculate(10, 5, subtract))
	fmt.Println("Multiply:", calculate(10, 5, multiply))
}

/*
============================================================
Q35. Return a function.
============================================================

A function can return another function.

This is another important concept behind closures.
*/
func example35() {
	fmt.Println("\n========== Q35: Return a Function ==========")

	createMultiplier := func(factor int) func(int) int {
		return func(num int) int {
			return num * factor
		}
	}

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Println("Double:", double(10))
	fmt.Println("Triple:", triple(10))
}

/*
============================================================
Q36. Anonymous function.
============================================================

An anonymous function has no name.

It can be assigned to a variable:

    add := func(a,b int) int {...}

Useful for small one-time operations.
*/
func example36() {
	fmt.Println("\n========== Q36: Anonymous Function ==========")

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(10, 20))
}

/*
============================================================
Q37. Immediately invoked function.
============================================================

The function is created and called immediately.

    func() {
        ...
    }()

Useful for small scoped initialization or one-time logic.
*/
func example37() {
	fmt.Println("\n========== Q37: Immediately Invoked Function ==========")

	func() {
		fmt.Println("Executed immediately")
	}()
}

/*
============================================================
Q38. Anonymous function with parameters.
============================================================
*/
func example38() {
	fmt.Println("\n========== Q38: Anonymous Function + Parameters ==========")

	result := func(a, b int) int {
		return a * b
	}(10, 5)

	fmt.Println("Result:", result)
}

/*
============================================================
Q39. Anonymous function with a slice.
============================================================

The function is stored and then reused for every element.
*/
func example39() {
	fmt.Println("\n========== Q39: Anonymous Function + Slice ==========")

	numbers := []int{1, 2, 3, 4, 5}

	process := func(num int) {
		fmt.Println("Double:", num*2)
	}

	for _, num := range numbers {
		process(num)
	}
}

/*
============================================================
Q40. Basic closure.
============================================================

A closure is a function that captures variables from its
surrounding scope.

The variable count continues to exist because the returned
function still references it.

This is why:

    c()
    c()
    c()

returns:

    1
    2
    3
*/
func example40() {
	fmt.Println("\n========== Q40: Closure Counter ==========")

	counter := func() func() int {
		count := 0

		return func() int {
			count++
			return count
		}
	}

	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

/*
============================================================
Q41. Closure with multiplier.
============================================================

The returned function remembers factor.

double remembers factor=2.
triple remembers factor=3.
*/
func example41() {
	fmt.Println("\n========== Q41: Closure Multiplier ==========")

	multiplier := func(factor int) func(int) int {
		return func(num int) int {
			return num * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Double:", double(10))
	fmt.Println("Triple:", triple(10))
}

/*
============================================================
Q42. Closure with custom starting value.
============================================================
*/
func example42() {
	fmt.Println("\n========== Q42: Closure Custom Counter ==========")

	counter := func(start int) func() int {
		count := start

		return func() int {
			count++
			return count
		}
	}

	c := counter(100)

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

/*
============================================================
Q43. Factorial using recursion.
============================================================

Recursion means a function calls itself.

Important:
Every recursive function needs a BASE CASE.

Without the base case, recursion continues indefinitely.

factorial(5)
= 5 * factorial(4)
= 5 * 4 * factorial(3)
= ...
= 120
*/
func example43() {
	fmt.Println("\n========== Q43: Recursive Factorial ==========")

	var factorial func(int) int

	factorial = func(n int) int {
		if n == 0 {
			return 1
		}

		return n * factorial(n-1)
	}

	fmt.Println("5! =", factorial(5))
}

/*
============================================================
Q44. Fibonacci using recursion.
============================================================

Fibonacci:
0 1 1 2 3 5 8 13 ...

Formula:
F(n) = F(n-1) + F(n-2)

Base cases:
F(0) = 0
F(1) = 1
*/
func example44() {
	fmt.Println("\n========== Q44: Recursive Fibonacci ==========")

	var fibonacci func(int) int

	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println("Fibonacci(7):", fibonacci(7))
}

/*
============================================================
Q45. Sum from 1 to N using recursion.
============================================================
*/
func example45() {
	fmt.Println("\n========== Q45: Recursive Sum ==========")

	var sum func(int) int

	sum = func(n int) int {
		if n == 0 {
			return 0
		}

		return n + sum(n-1)
	}

	fmt.Println("Sum 1..10:", sum(10))
}

/*
============================================================
Q46. Power using recursion.
============================================================

2^5 = 2 * 2 * 2 * 2 * 2 = 32
*/
func example46() {
	fmt.Println("\n========== Q46: Recursive Power ==========")

	var power func(int, int) int

	power = func(base, exponent int) int {
		if exponent == 0 {
			return 1
		}

		return base * power(base, exponent-1)
	}

	fmt.Println("2^5:", power(2, 5))
}

/*
============================================================
Q47. Reverse a string using recursion.
============================================================

Strings in Go contain UTF-8 bytes.

For simple ASCII interview examples, slicing works directly.

For real Unicode text, prefer []rune when reversing characters.
*/
func example47() {
	fmt.Println("\n========== Q47: Recursive String Reverse ==========")

	var reverse func(string) string

	reverse = func(s string) string {
		if len(s) <= 1 {
			return s
		}

		return reverse(s[1:]) + string(s[0])
	}

	fmt.Println("Reverse:", reverse("hello"))
}

/*
============================================================
Q48. Count digits using recursion.
============================================================
*/
func example48() {
	fmt.Println("\n========== Q48: Recursive Digit Count ==========")

	var countDigits func(int) int

	countDigits = func(num int) int {
		if num == 0 {
			return 0
		}

		return 1 + countDigits(num/10)
	}

	fmt.Println("Digits:", countDigits(12345))
}

/*
============================================================
Q49. Sum slice elements.
============================================================

A slice is commonly passed to a function.

Important:
A slice is a descriptor containing pointer, length and capacity.
Passing a slice to a function does not copy all underlying elements.
*/
func example49() {
	fmt.Println("\n========== Q49: Slice Sum ==========")

	sum := func(numbers []int) int {
		total := 0

		for _, num := range numbers {
			total += num
		}

		return total
	}

	numbers := []int{10, 20, 30, 40}

	fmt.Println("Sum:", sum(numbers))
}

/*
============================================================
Q50. Find largest element in a slice.
============================================================
*/
func example50() {
	fmt.Println("\n========== Q50: Largest Slice Element ==========")

	largest := func(numbers []int) int {
		if len(numbers) == 0 {
			return 0
		}

		max := numbers[0]

		for _, num := range numbers {
			if num > max {
				max = num
			}
		}

		return max
	}

	fmt.Println(
		"Largest:",
		largest([]int{10, 50, 20, 5}),
	)
}

/*
============================================================
Q51. Reverse a slice inside a function.
============================================================

This function does not return the slice.

Why?

Because slices refer to an underlying array.

Swapping elements changes that underlying array.

Therefore the caller sees the changes.
*/
func example51() {
	fmt.Println("\n========== Q51: Reverse Slice ==========")

	reverse := func(numbers []int) {
		for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers := []int{1, 2, 3, 4, 5}

	reverse(numbers)

	fmt.Println("Reversed:", numbers)
}

/*
============================================================
Q52. Filter even numbers.
============================================================

A function can create and return a new slice.

Original slice remains unchanged.
*/
func example52() {
	fmt.Println("\n========== Q52: Filter Even Numbers ==========")

	evenNumbers := func(numbers []int) []int {
		result := []int{}

		for _, num := range numbers {
			if num%2 == 0 {
				result = append(result, num)
			}
		}

		return result
	}

	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Even:", evenNumbers(numbers))
}

/*
============================================================
Q53. Find duplicate numbers.
============================================================

This is a basic interview implementation.

Time complexity:
O(n^2)

For large data, a map can reduce this to approximately O(n).
*/
func example53() {
	fmt.Println("\n========== Q53: Find Duplicates ==========")

	findDuplicates := func(numbers []int) []int {
		result := []int{}

		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {

				if numbers[i] != numbers[j] {
					continue
				}

				exists := false

				for _, value := range result {
					if value == numbers[i] {
						exists = true
						break
					}
				}

				if !exists {
					result = append(result, numbers[i])
				}

				break
			}
		}

		return result
	}

	numbers := []int{1, 2, 3, 2, 4, 1}

	fmt.Println("Duplicates:", findDuplicates(numbers))
}

/*
============================================================
Q54. Higher-order function: apply operation.
============================================================

A higher-order function either:
1. accepts a function, or
2. returns a function.

Here apply accepts:

    func(int) int

and applies it to every element.
*/
func example54() {
	fmt.Println("\n========== Q54: Higher-Order Apply ==========")

	apply := func(
		numbers []int,
		operation func(int) int,
	) []int {

		result := []int{}

		for _, num := range numbers {
			result = append(result, operation(num))
		}

		return result
	}

	numbers := []int{1, 2, 3, 4, 5}

	squared := apply(numbers, func(num int) int {
		return num * num
	})

	fmt.Println("Squared:", squared)
}

/*
============================================================
Q55. Higher-order function: filter.
============================================================

condition decides whether an element should be included.

This is similar to filter/map concepts found in many languages.
*/
func example55() {
	fmt.Println("\n========== Q55: Higher-Order Filter ==========")

	filter := func(
		numbers []int,
		condition func(int) bool,
	) []int {

		result := []int{}

		for _, num := range numbers {
			if condition(num) {
				result = append(result, num)
			}
		}

		return result
	}

	numbers := []int{1, 2, 3, 4, 5, 6}

	even := filter(numbers, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println("Even:", even)
}

/*
============================================================
Q56. Basic defer.
============================================================

defer schedules a function call to execute when the
surrounding function returns.

Output:
Normal
Deferred

Very common uses:
- closing files
- closing HTTP response bodies
- unlocking mutexes
- database cleanup
*/
func example56() {
	fmt.Println("\n========== Q56: Defer ==========")

	test := func() {
		defer fmt.Println("Deferred")
		fmt.Println("Normal")
	}

	test()
}

/*
============================================================
Q57. Multiple defer statements.
============================================================

defer uses LIFO order:

Last defer registered
        ↓
executes first

Therefore:
First
Second
Third

registration becomes:
Third
Second
First

at function return.
*/
func example57() {
	fmt.Println("\n========== Q57: Multiple Defer ==========")

	test := func() {
		defer fmt.Println("First")
		defer fmt.Println("Second")
		defer fmt.Println("Third")

		fmt.Println("Normal")
	}

	test()
}

/*
============================================================
Q58. defer + return.
============================================================

defer executes before the function actually returns to
its caller.

So "Deferred" prints before main receives the result.
*/
func example58() {
	fmt.Println("\n========== Q58: Defer + Return ==========")

	test := func() int {
		defer fmt.Println("Deferred")
		return 10
	}

	result := test()

	fmt.Println("Result:", result)
}

/*
============================================================
Q59. Realistic calculator with error handling.
============================================================

This combines:
- multiple functions
- parameters
- return values
- error
- caller-side error handling

This style is much closer to real Go backend code.
*/
func example59() {
	fmt.Println("\n========== Q59: Calculator + Error ==========")

	add := func(a, b float64) float64 {
		return a + b
	}

	subtract := func(a, b float64) float64 {
		return a - b
	}

	multiply := func(a, b float64) float64 {
		return a * b
	}

	divide := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}

		return a / b, nil
	}

	a := 20.0
	b := 5.0

	fmt.Println("Add:", add(a, b))
	fmt.Println("Subtract:", subtract(a, b))
	fmt.Println("Multiply:", multiply(a, b))

	result, err := divide(a, b)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Divide:", result)
}

/*
============================================================
Q60. Realistic backend-style validation function.
============================================================

This is an important pattern for Go backend interviews.

A common service function looks like:

    func validateUser(user User) error

The function returns nil if everything is valid.

Otherwise it returns an error.

This pattern appears throughout Go APIs, services and
business logic.

Notice that errors are values in Go.

We don't normally use try/catch for ordinary validation.
*/
func example60() {
	fmt.Println("\n========== Q60: Backend-Style Validation ==========")

	type User struct {
		Name  string
		Email string
		Age   int
	}

	validateUser := func(user User) error {

		if strings.TrimSpace(user.Name) == "" {
			return errors.New("name is required")
		}

		if strings.TrimSpace(user.Email) == "" {
			return errors.New("email is required")
		}

		if user.Age < 18 {
			return errors.New("user must be 18 or older")
		}

		return nil
	}

	user := User{
		Name:  "Ashutosh",
		Email: "ashutosh@example.com",
		Age:   25,
	}

	err := validateUser(user)

	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("User is valid")
}

/*
============================================================
MAIN
============================================================
*/
func main() {

	example01()
	example02()
	example03()
	example04()
	example05()
	example06()
	example07()
	example08()
	example09()
	example10()

	example11()
	example12()
	example13()
	example14()
	example15()
	example16()
	example17()
	example18()
	example19()
	example20()

	example21()
	example22()
	example23()
	example24()
	example25()
	example26()
	example27()
	example28()
	example29()
	example30()

	example31()
	example32()
	example33()
	example34()
	example35()
	example36()
	example37()
	example38()
	example39()
	example40()

	example41()
	example42()
	example43()
	example44()
	example45()
	example46()
	example47()
	example48()
	example49()
	example50()

	example51()
	example52()
	example53()
	example54()
	example55()
	example56()
	example57()
	example58()
	example59()
	example60()

	fmt.Println("\n============================================================")
	fmt.Println("ALL 60 FUNCTION EXAMPLES COMPLETED")
	fmt.Println("============================================================")
}
