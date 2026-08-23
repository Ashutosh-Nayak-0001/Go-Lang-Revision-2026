package main

import "fmt"

/*
===============================================================
GO RECURSION — 50 QUESTIONS & ANSWERS
===============================================================

Run:
    go run recursion_50.go

LEVEL:
    Q01-Q15  Beginner
    Q16-Q35  Intermediate
    Q36-Q50  Interview / Advanced

===============================================================
WHAT IS RECURSION?
===============================================================

Recursion is a technique where a function calls itself to solve
a smaller version of the same problem.

Every useful recursive function normally has two important parts:

1. BASE CASE
   The condition that stops recursion.

2. RECURSIVE CASE
   The function calls itself with a smaller/simpler input.

Basic example:

    func count(n int) {
        if n == 0 {
            return
        }

        fmt.Println(n)
        count(n - 1)
    }

===============================================================
Q01. What is recursion?
===============================================================
ANSWER:

Recursion is when a function calls itself directly or indirectly.

Example:

    func count(n int) {
        if n == 0 {
            return
        }

        fmt.Println(n)
        count(n - 1)
    }

The function keeps calling itself until the base case is reached.

*/

func q01() {
	fmt.Println("\nQ01")
	count01(5)
}

func count01(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	count01(n - 1)
}

/*
===============================================================
Q02. What is a base case?
===============================================================
ANSWER:

The base case is the condition that stops recursive calls.

Without a base case, recursion can continue until the program
runs out of stack space.

Example:

    if n == 0 {
        return
    }

*/

func q02() {
	fmt.Println("\nQ02")

	fmt.Println("Base case: n == 0")

	count02(3)
}

func count02(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	count02(n - 1)
}

/*
===============================================================
Q03. What happens if there is no base case?
===============================================================
ANSWER:

The function keeps calling itself.

Eventually the program can terminate because the call stack
cannot grow indefinitely.

*/

func q03() {
	fmt.Println("\nQ03")
	fmt.Println("Missing base case can cause uncontrolled recursion.")
}

/*
===============================================================
Q04. Print numbers from 1 to N using recursion.
===============================================================
*/

func q04() {
	fmt.Println("\nQ04")
	print1ToN04(1, 5)
}

func print1ToN04(current, n int) {
	if current > n {
		return
	}

	fmt.Println(current)
	print1ToN04(current+1, n)
}

/*
===============================================================
Q05. Print numbers from N to 1 using recursion.
===============================================================
*/

func q05() {
	fmt.Println("\nQ05")
	printNTo104(5)
}

func printNTo104(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	printNTo104(n - 1)
}

/*
===============================================================
Q06. Find sum of numbers from 1 to N.
===============================================================
Formula:

    sum(n) = n + sum(n-1)

Base:

    sum(0) = 0
*/

func q06() {
	fmt.Println("\nQ06")
	fmt.Println("Sum:", sum06(5))
}

func sum06(n int) int {
	if n == 0 {
		return 0
	}

	return n + sum06(n-1)
}

/*
===============================================================
Q07. Find factorial of N.
===============================================================
Formula:

    n! = n * (n-1)!

Base:

    0! = 1
*/

func q07() {
	fmt.Println("\nQ07")
	fmt.Println("5! =", factorial07(5))
}

func factorial07(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial07(n-1)
}

/*
===============================================================
Q08. Find power using recursion.
===============================================================
Calculate:

    base^exponent
*/

func q08() {
	fmt.Println("\nQ08")
	fmt.Println("2^5 =", power08(2, 5))
}

func power08(base, exponent int) int {
	if exponent == 0 {
		return 1
	}

	return base * power08(base, exponent-1)
}

/*
===============================================================
Q09. Print array elements recursively.
===============================================================
*/

func q09() {
	fmt.Println("\nQ09")

	arr := []int{10, 20, 30, 40, 50}

	printArray09(arr, 0)
}

func printArray09(arr []int, index int) {
	if index >= len(arr) {
		return
	}

	fmt.Println(arr[index])

	printArray09(arr, index+1)
}

/*
===============================================================
Q10. Print array in reverse using recursion.
===============================================================
*/

func q10() {
	fmt.Println("\nQ10")

	arr := []int{10, 20, 30, 40, 50}

	printReverse10(arr, len(arr)-1)
}

func printReverse10(arr []int, index int) {
	if index < 0 {
		return
	}

	fmt.Println(arr[index])

	printReverse10(arr, index-1)
}

/*
===============================================================
Q11. Find sum of array elements recursively.
===============================================================
*/

func q11() {
	fmt.Println("\nQ11")

	arr := []int{10, 20, 30, 40}

	fmt.Println("Sum:", arraySum11(arr, 0))
}

func arraySum11(arr []int, index int) int {
	if index == len(arr) {
		return 0
	}

	return arr[index] + arraySum11(arr, index+1)
}

/*
===============================================================
Q12. Find maximum element recursively.
===============================================================
*/

func q12() {
	fmt.Println("\nQ12")

	arr := []int{10, 55, 20, 90, 40}

	fmt.Println("Maximum:", max12(arr, 0))
}

func max12(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[index]
	}

	restMax := max12(arr, index+1)

	if arr[index] > restMax {
		return arr[index]
	}

	return restMax
}

/*
===============================================================
Q13. Count occurrences of a number.
===============================================================
*/

func q13() {
	fmt.Println("\nQ13")

	arr := []int{2, 5, 2, 7, 2, 9}

	fmt.Println("Count:", countOccurrences13(arr, 0, 2))
}

func countOccurrences13(arr []int, index, target int) int {
	if index == len(arr) {
		return 0
	}

	count := countOccurrences13(arr, index+1, target)

	if arr[index] == target {
		count++
	}

	return count
}

/*
===============================================================
Q14. Search an element recursively.
===============================================================
*/

func q14() {
	fmt.Println("\nQ14")

	arr := []int{10, 20, 30, 40}

	fmt.Println("Found:", contains14(arr, 0, 30))
}

func contains14(arr []int, index, target int) bool {
	if index == len(arr) {
		return false
	}

	if arr[index] == target {
		return true
	}

	return contains14(arr, index+1, target)
}

/*
===============================================================
Q15. Reverse a string recursively.
===============================================================
*/

func q15() {
	fmt.Println("\nQ15")

	fmt.Println(reverse15("golang"))
}

func reverse15(s string) string {
	if len(s) <= 1 {
		return s
	}

	return reverse15(s[1:]) + string(s[0])
}

/*
===============================================================
Q16. Check palindrome using recursion.
===============================================================
*/

func q16() {
	fmt.Println("\nQ16")

	fmt.Println("madam:", palindrome16("madam"))
	fmt.Println("hello:", palindrome16("hello"))
}

func palindrome16(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return palindrome16(s[1 : len(s)-1])
}

/*
===============================================================
Q17. Fibonacci using recursion.
===============================================================
Definition:

    fib(0) = 0
    fib(1) = 1

    fib(n) = fib(n-1) + fib(n-2)
*/

func q17() {
	fmt.Println("\nQ17")

	fmt.Println("fib(7):", fibonacci17(7))
}

func fibonacci17(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci17(n-1) + fibonacci17(n-2)
}

/*
IMPORTANT:

This simple Fibonacci implementation is inefficient.

Time complexity is approximately O(2^N).

Later questions show how memoization improves it.
*/

/*
===============================================================
Q18. Find GCD using recursion.
===============================================================
Euclidean algorithm:

    gcd(a,b) = gcd(b, a%b)

Base:

    gcd(a,0) = a
*/

func q18() {
	fmt.Println("\nQ18")

	fmt.Println("GCD:", gcd18(48, 18))
}

func gcd18(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd18(b, a%b)
}

/*
===============================================================
Q19. Calculate multiplication using recursion.
===============================================================
Do not use * for the multiplication itself.

Example:

    5 * 4
    = 5 + 5 + 5 + 5
*/

func q19() {
	fmt.Println("\nQ19")

	fmt.Println("5 x 4:", multiply19(5, 4))
}

func multiply19(a, b int) int {
	if b == 0 {
		return 0
	}

	return a + multiply19(a, b-1)
}

/*
===============================================================
Q20. Calculate sum of digits recursively.
===============================================================
Example:

    1234
    = 1 + 2 + 3 + 4
*/

func q20() {
	fmt.Println("\nQ20")

	fmt.Println("Sum:", digitSum20(1234))
}

func digitSum20(n int) int {
	if n == 0 {
		return 0
	}

	return n%10 + digitSum20(n/10)
}

/*
===============================================================
Q21. Count digits recursively.
===============================================================
*/

func q21() {
	fmt.Println("\nQ21")

	fmt.Println("Digits:", countDigits21(123456))
}

func countDigits21(n int) int {
	if n == 0 {
		return 0
	}

	return 1 + countDigits21(n/10)
}

/*
===============================================================
Q22. Reverse an integer recursively.
===============================================================
*/

func q22() {
	fmt.Println("\nQ22")

	fmt.Println("Reverse:", reverseNumber22(12345, 0))
}

func reverseNumber22(n, result int) int {
	if n == 0 {
		return result
	}

	return reverseNumber22(n/10, result*10+n%10)
}

/*
===============================================================
Q23. Check if number is power of 2 recursively.
===============================================================
*/

func q23() {
	fmt.Println("\nQ23")

	fmt.Println("16:", isPowerOfTwo23(16))
	fmt.Println("18:", isPowerOfTwo23(18))
}

func isPowerOfTwo23(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 || n%2 != 0 {
		return false
	}

	return isPowerOfTwo23(n / 2)
}

/*
===============================================================
Q24. Binary search recursively.
===============================================================
Array must be sorted.
*/

func q24() {
	fmt.Println("\nQ24")

	arr := []int{10, 20, 30, 40, 50, 60}

	index := binarySearch24(arr, 0, len(arr)-1, 40)

	fmt.Println("Index:", index)
}

func binarySearch24(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if target < arr[mid] {
		return binarySearch24(arr, left, mid-1, target)
	}

	return binarySearch24(arr, mid+1, right, target)
}

/*
===============================================================
Q25. Recursive linear search.
===============================================================
*/

func q25() {
	fmt.Println("\nQ25")

	arr := []int{5, 10, 15, 20}

	fmt.Println("Index:", linearSearch25(arr, 0, 15))
}

func linearSearch25(arr []int, index, target int) int {
	if index == len(arr) {
		return -1
	}

	if arr[index] == target {
		return index
	}

	return linearSearch25(arr, index+1, target)
}

/*
===============================================================
Q26. Find minimum element recursively.
===============================================================
*/

func q26() {
	fmt.Println("\nQ26")

	arr := []int{30, 10, 50, 5, 20}

	fmt.Println("Minimum:", min26(arr, 0))
}

func min26(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[index]
	}

	restMin := min26(arr, index+1)

	if arr[index] < restMin {
		return arr[index]
	}

	return restMin
}

/*
===============================================================
Q27. Count vowels recursively.
===============================================================
*/

func q27() {
	fmt.Println("\nQ27")

	fmt.Println("Vowels:", countVowels27("recursion"))
}

func countVowels27(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0

	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u':
		count = 1
	}

	return count + countVowels27(s[1:])
}

/*
===============================================================
Q28. Remove a character recursively.
===============================================================
*/

func q28() {
	fmt.Println("\nQ28")

	fmt.Println(removeChar28("banana", 'a'))
}

func removeChar28(s string, target byte) string {
	if len(s) == 0 {
		return ""
	}

	if s[0] == target {
		return removeChar28(s[1:], target)
	}

	return string(s[0]) + removeChar28(s[1:], target)
}

/*
===============================================================
Q29. Count a character recursively.
===============================================================
*/

func q29() {
	fmt.Println("\nQ29")

	fmt.Println("Count:", countChar29("banana", 'a'))
}

func countChar29(s string, target byte) int {
	if len(s) == 0 {
		return 0
	}

	count := 0

	if s[0] == target {
		count = 1
	}

	return count + countChar29(s[1:], target)
}

/*
===============================================================
Q30. Calculate x^n efficiently.
===============================================================
Use exponentiation by squaring.

For even n:

    x^n = (x^(n/2))^2

For odd n:

    x^n = x * x^(n-1)
*/

func q30() {
	fmt.Println("\nQ30")

	fmt.Println("2^10:", fastPower30(2, 10))
}

func fastPower30(x, n int) int {
	if n == 0 {
		return 1
	}

	half := fastPower30(x, n/2)

	if n%2 == 0 {
		return half * half
	}

	return x * half * half
}

/*
Time complexity:

Simple power recursion:
    O(N)

Fast power:
    O(log N)
*/

/*
===============================================================
Q31. Generate all subsequences of a string.
===============================================================
*/

func q31() {
	fmt.Println("\nQ31")

	subsequences31("abc", 0, "")
}

func subsequences31(s string, index int, current string) {
	if index == len(s) {
		fmt.Println(current)
		return
	}

	// Exclude current character.
	subsequences31(s, index+1, current)

	// Include current character.
	subsequences31(s, index+1, current+string(s[index]))
}

/*
For N characters:

Number of subsequences = 2^N

Time complexity is approximately O(2^N).
*/

/*
===============================================================
Q32. Generate all permutations of a string.
===============================================================
*/

func q32() {
	fmt.Println("\nQ32")

	permutations32([]byte("abc"), 0)
}

func permutations32(chars []byte, index int) {
	if index == len(chars) {
		fmt.Println(string(chars))
		return
	}

	for i := index; i < len(chars); i++ {
		chars[index], chars[i] = chars[i], chars[index]

		permutations32(chars, index+1)

		// Backtracking.
		chars[index], chars[i] = chars[i], chars[index]
	}
}

/*
For N unique characters:

Number of permutations = N!

This is a classic recursion + backtracking problem.
*/

/*
===============================================================
Q33. Generate binary strings of length N.
===============================================================
*/

func q33() {
	fmt.Println("\nQ33")

	binaryStrings33(3, "")
}

func binaryStrings33(n int, current string) {
	if n == 0 {
		fmt.Println(current)
		return
	}

	binaryStrings33(n-1, current+"0")
	binaryStrings33(n-1, current+"1")
}

/*
For length N:

Number of strings = 2^N.
*/

/*
===============================================================
Q34. Generate balanced parentheses.
===============================================================
*/

func q34() {
	fmt.Println("\nQ34")

	generateParentheses34(3, 3, "")
}

func generateParentheses34(open, close int, current string) {
	if open == 0 && close == 0 {
		fmt.Println(current)
		return
	}

	if open > 0 {
		generateParentheses34(open-1, close, current+"(")
	}

	if close > open {
		generateParentheses34(open, close-1, current+")")
	}
}

/*
Important rule:

A closing parenthesis can be added only when:

    close > open

This prevents invalid sequences.
*/

/*
===============================================================
Q35. Tower of Hanoi.
===============================================================
*/

func q35() {
	fmt.Println("\nQ35")

	towerOfHanoi35(3, 'A', 'C', 'B')
}

func towerOfHanoi35(n int, from, to, auxiliary byte) {
	if n == 0 {
		return
	}

	towerOfHanoi35(n-1, from, auxiliary, to)

	fmt.Printf("Move disk %d from %c to %c\n", n, from, to)

	towerOfHanoi35(n-1, auxiliary, to, from)
}

/*
Minimum moves:

    2^N - 1

For N = 3:

    7 moves
*/

/*
===============================================================
Q36. Fibonacci with memoization.
===============================================================
ANSWER:

Memoization stores already calculated results.

This changes Fibonacci from exponential time to O(N).
*/

func q36() {
	fmt.Println("\nQ36")

	memo := make(map[int]int)

	fmt.Println("fib(40):", fibonacciMemo36(40, memo))
}

func fibonacciMemo36(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if value, exists := memo[n]; exists {
		return value
	}

	result := fibonacciMemo36(n-1, memo) +
		fibonacciMemo36(n-2, memo)

	memo[n] = result

	return result
}

/*
===============================================================
Q37. What is recursion depth?
===============================================================
ANSWER:

Recursion depth is the number of active recursive calls before
the base case is reached.

Example:

    count(5)
       |
       count(4)
          |
          count(3)
             |
             count(2)
                |
                count(1)
                   |
                   count(0)

The maximum depth is approximately N.
*/

func q37() {
	fmt.Println("\nQ37")
	fmt.Println("Recursion depth depends on the number of active calls.")
}

/*
===============================================================
Q38. What is the call stack?
===============================================================
ANSWER:

Each function call gets a stack frame containing information
needed to resume execution.

With recursion:

    function A
       calls A
          calls A
             calls A

Each call creates another stack frame.

When the base case is reached, calls return in reverse order.

*/

func q38() {
	fmt.Println("\nQ38")
	stackDemo38(3)
}

func stackDemo38(n int) {
	if n == 0 {
		fmt.Println("Base case")
		return
	}

	fmt.Println("Entering:", n)

	stackDemo38(n - 1)

	fmt.Println("Returning:", n)
}

/*
Output demonstrates:

Entering: 3
Entering: 2
Entering: 1
Base case
Returning: 1
Returning: 2
Returning: 3

This is extremely important for recursion interviews.
*/

/*
===============================================================
Q39. Explain recursion using a dry run.
===============================================================
Question:

What does:

    factorial(4)

do?

Answer:

factorial(4)
= 4 * factorial(3)
= 4 * 3 * factorial(2)
= 4 * 3 * 2 * factorial(1)
= 4 * 3 * 2 * 1
= 24

*/

func q39() {
	fmt.Println("\nQ39")
	fmt.Println("factorial(4) =", factorial39(4))
}

func factorial39(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial39(n-1)
}

/*
===============================================================
Q40. What is tail recursion?
===============================================================
ANSWER:

A recursive call is tail recursive when it is the final
operation performed by the function.

Example:

    func count(n int) {
        if n == 0 {
            return
        }

        count(n - 1)
    }

The recursive call is the last operation.

Compare:

    return n * factorial(n-1)

This is NOT tail recursive because multiplication must happen
after the recursive call returns.

*/

func q40() {
	fmt.Println("\nQ40")

	fmt.Println("Tail recursion example:")
	tailCount40(3)
}

func tailCount40(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	tailCount40(n - 1)
}

/*
===============================================================
Q41. Convert recursive factorial to iterative.
===============================================================
ANSWER:

Many recursive problems can be written using loops.

Recursive:

    factorial(n)
        n * factorial(n-1)

Iterative:

    result := 1

    for i := 2; i <= n; i++ {
        result *= i
    }

*/

func q41() {
	fmt.Println("\nQ41")

	fmt.Println("Iterative factorial:", factorialIterative41(5))
}

func factorialIterative41(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

/*
===============================================================
Q42. Recursion vs iteration.
===============================================================
ANSWER:

Recursion:
    - Often cleaner for tree/graph/backtracking problems.
    - Uses call stack.
    - Can be easier to understand for divide-and-conquer.
    - Can cause deep stack usage.

Iteration:
    - Uses loops.
    - Usually avoids recursive call overhead.
    - Often better for simple counting/repetition.
    - Can be harder to express for trees/backtracking.

*/

func q42() {
	fmt.Println("\nQ42")
	fmt.Println("Use recursion when it naturally represents the problem.")
}

/*
===============================================================
Q43. What is indirect recursion?
===============================================================
ANSWER:

Indirect recursion occurs when function A calls B and B calls A.

    A -> B -> A -> B ...

*/

func q43() {
	fmt.Println("\nQ43")

	indirectA43(3)
}

func indirectA43(n int) {
	if n <= 0 {
		return
	}

	fmt.Println("A:", n)
	indirectB43(n - 1)
}

func indirectB43(n int) {
	if n <= 0 {
		return
	}

	fmt.Println("B:", n)
	indirectA43(n - 1)
}

/*
===============================================================
Q44. Count ways to climb stairs.
===============================================================
Question:

You can climb either 1 or 2 steps.

How many ways can you reach step N?

Recurrence:

    ways(n) = ways(n-1) + ways(n-2)

*/

func q44() {
	fmt.Println("\nQ44")

	fmt.Println("Ways for 5 steps:", climbStairs44(5))
}

func climbStairs44(n int) int {
	if n <= 1 {
		return 1
	}

	return climbStairs44(n-1) + climbStairs44(n-2)
}

/*
This is similar to Fibonacci.

The naive version has exponential complexity.

Memoization can improve it to O(N).
*/

/*
===============================================================
Q45. Find sum of array using divide and conquer recursion.
===============================================================
*/

func q45() {
	fmt.Println("\nQ45")

	arr := []int{10, 20, 30, 40}

	fmt.Println("Sum:", divideSum45(arr, 0, len(arr)-1))
}

func divideSum45(arr []int, left, right int) int {
	if left > right {
		return 0
	}

	if left == right {
		return arr[left]
	}

	mid := left + (right-left)/2

	leftSum := divideSum45(arr, left, mid)
	rightSum := divideSum45(arr, mid+1, right)

	return leftSum + rightSum
}

/*
This demonstrates divide-and-conquer recursion.

The array is repeatedly divided into smaller ranges.
*/

/*
===============================================================
Q46. Recursive merge sort.
===============================================================
*/

func q46() {
	fmt.Println("\nQ46")

	arr := []int{5, 2, 8, 1, 3}

	mergeSort46(arr)

	fmt.Println(arr)
}

func mergeSort46(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	left := append([]int{}, arr[:mid]...)
	right := append([]int{}, arr[mid:]...)

	mergeSort46(left)
	mergeSort46(right)

	merge46(arr, left, right)
}

func merge46(result, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}

		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
}

/*
Merge sort:

Time:
    O(N log N)

Recursion depth:
    O(log N)

This is a classic recursion interview problem.
*/

/*
===============================================================
Q47. Recursive quicksort.
===============================================================
*/

func q47() {
	fmt.Println("\nQ47")

	arr := []int{5, 1, 8, 3, 2, 7}

	quickSort47(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSort47(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition47(arr, low, high)

	quickSort47(arr, low, pivotIndex-1)
	quickSort47(arr, pivotIndex+1, high)
}

func partition47(arr []int, low, high int) int {
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i
}

/*
Average time:
    O(N log N)

Worst case:
    O(N^2)

Recursion is used to process the two partitions.
*/

/*
===============================================================
Q48. Recursive tree traversal.
===============================================================
*/

type TreeNode48 struct {
	Value int
	Left  *TreeNode48
	Right *TreeNode48
}

func q48() {
	fmt.Println("\nQ48")

	root := &TreeNode48{
		Value: 1,
		Left: &TreeNode48{
			Value: 2,
			Left: &TreeNode48{
				Value: 4,
			},
			Right: &TreeNode48{
				Value: 5,
			},
		},
		Right: &TreeNode48{
			Value: 3,
		},
	}

	fmt.Print("Inorder: ")
	inorder48(root)
	fmt.Println()
}

func inorder48(root *TreeNode48) {
	if root == nil {
		return
	}

	inorder48(root.Left)

	fmt.Print(root.Value, " ")

	inorder48(root.Right)
}

/*
Trees are one of the most important practical uses of recursion.

Inorder:

    Left
    Root
    Right

Preorder:

    Root
    Left
    Right

Postorder:

    Left
    Right
    Root
*/

/*
===============================================================
Q49. Recursive DFS on a graph.
===============================================================
*/

func q49() {
	fmt.Println("\nQ49")

	graph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}

	visited := make(map[int]bool)

	fmt.Print("DFS: ")

	dfs49(graph, 1, visited)

	fmt.Println()
}

func dfs49(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true

	fmt.Print(node, " ")

	for _, neighbor := range graph[node] {
		dfs49(graph, neighbor, visited)
	}
}

/*
DFS = Depth First Search.

Recursion naturally represents:

    Visit node
       |
       +-- Visit neighbor
              |
              +-- Visit next neighbor

IMPORTANT:
Always track visited nodes when cycles are possible.

Otherwise:

    A -> B -> A

can cause infinite recursion.
*/

/*
===============================================================
Q50. INTERVIEW PROBLEM — Generate all combinations.
===============================================================

Question:

Given:

    [1, 2, 3]

generate all combinations of size 2.

Expected:

    [1 2]
    [1 3]
    [2 3]

This combines recursion with backtracking.
*/

func q50() {
	fmt.Println("\nQ50")

	arr := []int{1, 2, 3}
	current := []int{}

	combinations50(arr, 0, 2, current)
}

func combinations50(arr []int, start, k int, current []int) {
	if len(current) == k {
		fmt.Println(current)
		return
	}

	for i := start; i < len(arr); i++ {
		current = append(current, arr[i])

		combinations50(
			arr,
			i+1,
			k,
			current,
		)

		// Backtrack.
		current = current[:len(current)-1]
	}
}

/*
===============================================================
DEEP INTERVIEW NOTES
===============================================================

1. EVERY RECURSIVE FUNCTION SHOULD HAVE A STOPPING CONDITION

Example:

    func f(n int) {
        if n == 0 {
            return
        }

        f(n-1)
    }

Without a correct base case, the recursion may never terminate.

---------------------------------------------------------------

2. RECURSIVE CALL MUST MOVE TOWARD THE BASE CASE

Correct:

    f(n-1)

if base case is n == 0.

Potentially incorrect:

    f(n+1)

because it moves away from n == 0.

---------------------------------------------------------------

3. RECURSION USES THE CALL STACK

For:

    f(3)

you can imagine:

    f(3)
      f(2)
        f(1)
          f(0)

Then returning:

    f(0)
    f(1)
    f(2)
    f(3)

This explains why recursive functions can use significant
memory for deep recursion.

---------------------------------------------------------------

4. TIME COMPLEXITY MATTERS

Factorial:

    O(N)

Linear recursive search:

    O(N)

Binary search:

    O(log N)

Simple Fibonacci:

    O(2^N)

Merge sort:

    O(N log N)

Generating all subsets:

    O(2^N)

Generating permutations:

    O(N!)

---------------------------------------------------------------

5. SPACE COMPLEXITY

For simple recursion:

    f(n)
      f(n-1)
        f(n-2)

maximum active calls can be O(N).

So recursion may use:

    O(N)

stack space.

---------------------------------------------------------------

6. RECURSION TREE

For Fibonacci:

                fib(5)
              /       \
          fib(4)      fib(3)
          /   \        /   \
      fib(3) fib(2) fib(2) fib(1)

The same subproblems are calculated repeatedly.

Memoization removes this duplication.

---------------------------------------------------------------

7. MEMOIZATION

Memoization means:

    Calculate once
    Store result
    Reuse result

Example:

    memo[n] = result

This is very important in dynamic programming.

---------------------------------------------------------------

8. BACKTRACKING

Backtracking follows:

    choose
    explore
    undo

Example:

    current = append(current, choice)

    recurse(...)

    current = current[:len(current)-1]

Used for:

    permutations
    combinations
    subsets
    N-Queens
    Sudoku
    maze problems
    balanced parentheses

---------------------------------------------------------------

9. DIVIDE AND CONQUER

Divide problem into smaller problems.

Examples:

    Merge sort
    Quick sort
    Binary search

Pattern:

    divide
    solve recursively
    combine

---------------------------------------------------------------

10. RECURSION VS LOOP

Use a loop when the problem is simple repetition:

    for i := 0; i < n; i++ {}

Use recursion when the problem naturally contains nested
structure:

    trees
    DFS
    backtracking
    divide and conquer

---------------------------------------------------------------

11. TAIL RECURSION

Example:

    func f(n int) {
        if n == 0 {
            return
        }

        f(n-1)
    }

The recursive call is the final operation.

However, do not assume that Go automatically performs general
tail-call optimization.

Deep recursion should therefore be considered carefully.

---------------------------------------------------------------

12. DIRECT RECURSION

A function calls itself:

    func f() {
        f()
    }

---------------------------------------------------------------

13. INDIRECT RECURSION

A calls B.

B calls A.

    A -> B -> A

---------------------------------------------------------------

14. MUTUAL RECURSION

Mutual recursion is another name often used when multiple
functions recursively call one another.

Example:

    isEven()
       |
       v
    isOdd()
       |
       v
    isEven()

---------------------------------------------------------------

15. RECURSION WITH ARRAYS

Typical pattern:

    func solve(arr []int, index int) {
        if index == len(arr) {
            return
        }

        // process arr[index]

        solve(arr, index+1)
    }

The index acts as the state of recursion.

---------------------------------------------------------------

16. RECURSION WITH STRINGS

Typical pattern:

    func solve(s string) {
        if len(s) == 0 {
            return
        }

        solve(s[1:])
    }

Be aware that repeated string slicing and concatenation can
have performance implications.

---------------------------------------------------------------

17. RECURSION WITH TREES

Typical pattern:

    func traverse(root *Node) {
        if root == nil {
            return
        }

        traverse(root.Left)
        traverse(root.Right)
    }

The nil node is usually the base case.

---------------------------------------------------------------

18. RECURSION WITH GRAPHS

Always think about cycles.

Example:

    A -> B
    B -> C
    C -> A

Without visited tracking:

    dfs(A)
      dfs(B)
        dfs(C)
          dfs(A)
            ...

This never terminates.

---------------------------------------------------------------

19. RECURSION AND GOROUTINES

Do not confuse recursion with concurrency.

A recursive function can run synchronously.

A recursive function can also start goroutines, but that is a
separate concern.

---------------------------------------------------------------

20. COMMON INTERVIEW TRAPS

TRAP 1:
Every recursive problem is faster than iteration.

FALSE.

TRAP 2:
Recursion always causes stack overflow immediately.

FALSE.

It depends on recursion depth and available stack.

TRAP 3:
Every recursion problem requires memoization.

FALSE.

Memoization is useful when subproblems repeat.

TRAP 4:
Binary search requires recursion.

FALSE.

It can be implemented iteratively.

TRAP 5:
Trees always require recursion.

FALSE.

Trees can also be traversed iteratively using stacks/queues.

---------------------------------------------------------------

21. BEST WAY TO SOLVE A RECURSION QUESTION IN AN INTERVIEW

Use this sequence:

STEP 1:
Identify the smallest possible input.

STEP 2:
Define the base case.

STEP 3:
Ask:

    "How can I solve a smaller version of this problem?"

STEP 4:
Write the recursive call.

STEP 5:
Combine the result.

STEP 6:
Dry-run with a small input.

STEP 7:
State time and space complexity.

---------------------------------------------------------------

22. RECURSION TEMPLATE

Most beginner problems:

    func solve(n int) result {
        if baseCase {
            return baseResult
        }

        smaller := solve(smallerInput)

        return combine(n, smaller)
    }

Backtracking:

    func solve(state) {
        if solution {
            process()
            return
        }

        for _, choice := range choices {
            choose(choice)
            solve(newState)
            undo(choice)
        }
    }

===============================================================
50-QUESTION REVISION CHECKLIST
===============================================================

[01] What is recursion?
[02] What is a base case?
[03] What happens without a base case?
[04] Print 1 to N.
[05] Print N to 1.
[06] Sum 1 to N.
[07] Factorial.
[08] Power.
[09] Print array.
[10] Reverse array.
[11] Array sum.
[12] Maximum.
[13] Count occurrences.
[14] Search element.
[15] Reverse string.
[16] Palindrome.
[17] Fibonacci.
[18] GCD.
[19] Multiplication.
[20] Sum of digits.
[21] Count digits.
[22] Reverse integer.
[23] Power of 2.
[24] Binary search.
[25] Linear search.
[26] Minimum.
[27] Count vowels.
[28] Remove character.
[29] Count character.
[30] Fast power.
[31] Subsequences.
[32] Permutations.
[33] Binary strings.
[34] Balanced parentheses.
[35] Tower of Hanoi.
[36] Fibonacci memoization.
[37] Recursion depth.
[38] Call stack.
[39] Dry run.
[40] Tail recursion.
[41] Recursion to iteration.
[42] Recursion vs iteration.
[43] Indirect recursion.
[44] Climbing stairs.
[45] Divide and conquer.
[46] Merge sort.
[47] Quick sort.
[48] Tree traversal.
[49] Graph DFS.
[50] Combinations/backtracking.

===============================================================
END OF RECURSION 50
===============================================================
*/

func main() {
	q01()
	q02()
	q03()
	q04()
	q05()
	q06()
	q07()
	q08()
	q09()
	q10()
	q11()
	q12()
	q13()
	q14()
	q15()
	q16()
	q17()
	q18()
	q19()
	q20()
	q21()
	q22()
	q23()
	q24()
	q25()
	q26()
	q27()
	q28()
	q29()
	q30()
	q31()
	q32()
	q33()
	q34()
	q35()
	q36()
	q37()
	q38()
	q39()
	q40()
	q41()
	q42()
	q43()
	q44()
	q45()
	q46()
	q47()
	q48()
	q49()
	q50()

	fmt.Println("\n===============================================================")
	fmt.Println("ALL 50 RECURSION QUESTIONS COMPLETED")
	fmt.Println("===============================================================")
}
