package main

import "fmt"

func main() {

	// ============================================================
	// 1. PRINT NUMBERS FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q1 ==========")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// ============================================================
	// 2. PRINT NUMBERS FROM 10 TO 1
	// ============================================================

	fmt.Println("\n========== Q2 ==========")

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// ============================================================
	// 3. PRINT NUMBERS FROM 1 TO 100
	// ============================================================

	fmt.Println("\n========== Q3 ==========")

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// ============================================================
	// 4. PRINT EVEN NUMBERS FROM 1 TO 20
	// ============================================================

	fmt.Println("\n========== Q4 ==========")

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// ============================================================
	// 5. PRINT ODD NUMBERS FROM 1 TO 20
	// ============================================================

	fmt.Println("\n========== Q5 ==========")

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// ============================================================
	// 6. PRINT MULTIPLES OF 5
	// ============================================================

	fmt.Println("\n========== Q6 ==========")

	for i := 5; i <= 100; i += 5 {
		fmt.Println(i)
	}

	// ============================================================
	// 7. PRINT NUMBERS DIVISIBLE BY 3
	// ============================================================

	fmt.Println("\n========== Q7 ==========")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// ============================================================
	// 8. PRINT SQUARES FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q8 ==========")

	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

	// ============================================================
	// 9. PRINT CUBES FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q9 ==========")

	for i := 1; i <= 10; i++ {
		fmt.Println(i * i * i)
	}

	// ============================================================
	// 10. PRINT NUMBER AND SQUARE
	// ============================================================

	fmt.Println("\n========== Q10 ==========")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d -> %d\n", i, i*i)
	}

	// ============================================================
	// 11. SUM FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q11 ==========")

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)

	// ============================================================
	// 12. SUM FROM 1 TO 100
	// ============================================================

	fmt.Println("\n========== Q12 ==========")

	sum = 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)

	// ============================================================
	// 13. SUM OF EVEN NUMBERS
	// ============================================================

	fmt.Println("\n========== Q13 ==========")

	sum = 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum:", sum)

	// ============================================================
	// 14. SUM OF ODD NUMBERS
	// ============================================================

	fmt.Println("\n========== Q14 ==========")

	sum = 0

	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	fmt.Println("Sum:", sum)

	// ============================================================
	// 15. COUNT NUMBERS DIVISIBLE BY 5
	// ============================================================

	fmt.Println("\n========== Q15 ==========")

	count := 0

	for i := 1; i <= 100; i++ {
		if i%5 == 0 {
			count++
		}
	}

	fmt.Println("Count:", count)

	// ============================================================
	// 16. AVERAGE FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q16 ==========")

	sum = 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	average := float64(sum) / 10

	fmt.Println("Average:", average)

	// ============================================================
	// 17. SUM OF SQUARES
	// ============================================================

	fmt.Println("\n========== Q17 ==========")

	sum = 0

	for i := 1; i <= 10; i++ {
		sum += i * i
	}

	fmt.Println("Sum of squares:", sum)

	// ============================================================
	// 18. SUM OF CUBES
	// ============================================================

	fmt.Println("\n========== Q18 ==========")

	sum = 0

	for i := 1; i <= 10; i++ {
		sum += i * i * i
	}

	fmt.Println("Sum of cubes:", sum)

	// ============================================================
	// 19. COUNT EVEN AND ODD NUMBERS
	// ============================================================

	fmt.Println("\n========== Q19 ==========")

	even := 0
	odd := 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)

	// ============================================================
	// 20. SUM OF NUMBERS DIVISIBLE BY BOTH 3 AND 5
	// ============================================================

	fmt.Println("\n========== Q20 ==========")

	sum = 0
      
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum:", sum)

	// ============================================================
	// 21. MULTIPLICATION TABLE OF 5
	// ============================================================

	fmt.Println("\n========== Q21 ==========")

	for i := 1; i <= 10; i++ {
		fmt.Printf("5 x %d = %d\n", i, 5*i)
	}

	// ============================================================
	// 22. MULTIPLICATION TABLE OF ANY NUMBER
	// ============================================================

	fmt.Println("\n========== Q22 ==========")

	tableNumber := 7

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", tableNumber, i, tableNumber*i)
	}

	// ============================================================
	// 23. TABLES FROM 1 TO 5
	// ============================================================

	fmt.Println("\n========== Q23 ==========")

	for num := 1; num <= 5; num++ {

		fmt.Println("Table of", num)

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d\n", num, i, num*i)
		}

		fmt.Println()
	}

	// ============================================================
	// 24. TABLES FROM 1 TO 10
	// ============================================================

	fmt.Println("\n========== Q24 ==========")

	for num := 1; num <= 10; num++ {

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d\n", num, i, num*i)
		}

		fmt.Println()
	}

	// ============================================================
	// 25. COUNT DIGITS
	// ============================================================

	fmt.Println("\n========== Q25 ==========")

	num := 12345
	temp := num
	digitCount := 0

	for temp != 0 {
		temp /= 10
		digitCount++
	}

	fmt.Println("Digits:", digitCount)

	// ============================================================
	// 26. SUM OF DIGITS
	// ============================================================

	fmt.Println("\n========== Q26 ==========")

	num = 12345
	temp = num
	sum = 0

	for temp != 0 {
		digit := temp % 10
		sum += digit
		temp /= 10
	}

	fmt.Println("Sum of digits:", sum)

	// ============================================================
	// 27. REVERSE NUMBER
	// ============================================================

	fmt.Println("\n========== Q27 ==========")

	num = 12345
	temp = num
	reverse := 0

	for temp != 0 {
		digit := temp % 10
		reverse = reverse*10 + digit
		temp /= 10
	}

	fmt.Println("Reverse:", reverse)

	// ============================================================
	// 28. CHECK PALINDROME NUMBER
	// ============================================================

	fmt.Println("\n========== Q28 ==========")

	num = 121
	original := num
	temp = num
	reverse = 0

	for temp != 0 {
		digit := temp % 10
		reverse = reverse*10 + digit
		temp /= 10
	}

	if original == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

	// ============================================================
	// 29. FIND FIRST DIGIT
	// ============================================================

	fmt.Println("\n========== Q29 ==========")

	num = 12345

	for num >= 10 {
		num /= 10
	}

	fmt.Println("First digit:", num)

	// ============================================================
	// 30. FIND LAST DIGIT
	// ============================================================

	fmt.Println("\n========== Q30 ==========")

	num = 12345

	lastDigit := num % 10

	fmt.Println("Last digit:", lastDigit)

	// ============================================================
	// 31. PRODUCT OF DIGITS
	// ============================================================

	fmt.Println("\n========== Q31 ==========")

	num = 1234
	temp = num
	product := 1

	for temp != 0 {
		digit := temp % 10
		product *= digit
		temp /= 10
	}

	fmt.Println("Product:", product)

	// ============================================================
	// 32. COUNT EVEN DIGITS
	// ============================================================

	fmt.Println("\n========== Q32 ==========")

	num = 123456
	temp = num
	count = 0

	for temp != 0 {
		digit := temp % 10

		if digit%2 == 0 {
			count++
		}

		temp /= 10
	}

	fmt.Println("Even digits:", count)

	// ============================================================
	// 33. COUNT ODD DIGITS
	// ============================================================

	fmt.Println("\n========== Q33 ==========")

	num = 123456
	temp = num
	count = 0

	for temp != 0 {
		digit := temp % 10

		if digit%2 != 0 {
			count++
		}

		temp /= 10
	}

	fmt.Println("Odd digits:", count)

	// ============================================================
	// 34. FIND LARGEST DIGIT
	// ============================================================

	fmt.Println("\n========== Q34 ==========")

	num = 58329
	temp = num
	largest := 0

	for temp != 0 {
		digit := temp % 10

		if digit > largest {
			largest = digit
		}

		temp /= 10
	}

	fmt.Println("Largest digit:", largest)

	// ============================================================
	// 35. FIND SMALLEST DIGIT
	// ============================================================

	fmt.Println("\n========== Q35 ==========")

	num = 58329
	temp = num
	smallest := 9

	for temp != 0 {
		digit := temp % 10

		if digit < smallest {
			smallest = digit
		}

		temp /= 10
	}

	fmt.Println("Smallest digit:", smallest)

	// ============================================================
	// 36. FACTORIAL
	// ============================================================

	fmt.Println("\n========== Q36 ==========")

	n := 5
	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Println("Factorial:", factorial)

	// ============================================================
	// 37. POWER OF A NUMBER
	// ============================================================

	fmt.Println("\n========== Q37 ==========")

	base := 2
	exponent := 5
	result := 1

	for i := 1; i <= exponent; i++ {
		result *= base
	}

	fmt.Println("Result:", result)

	// ============================================================
	// 38. CHECK PRIME NUMBER
	// ============================================================

	fmt.Println("\n========== Q38 ==========")

	num = 29
	isPrime := true

	if num < 2 {
		isPrime = false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not prime")
	}

	// ============================================================
	// 39. PRINT PRIME NUMBERS FROM 1 TO 100
	// ============================================================

	fmt.Println("\n========== Q39 ==========")

	for num := 2; num <= 100; num++ {

		isPrime := true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(num)
		}
	}

	// ============================================================
	// 40. COUNT PRIME NUMBERS FROM 1 TO 100
	// ============================================================

	fmt.Println("\n========== Q40 ==========")

	count = 0

	for num := 2; num <= 100; num++ {

		isPrime := true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++
		}
	}

	fmt.Println("Prime count:", count)

	// ============================================================
	// 41. FIND FACTORS OF A NUMBER
	// ============================================================

	fmt.Println("\n========== Q41 ==========")

	num = 24

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}

	// ============================================================
	// 42. COUNT FACTORS
	// ============================================================

	fmt.Println("\n========== Q42 ==========")

	num = 24
	count = 0

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}

	fmt.Println("Factor count:", count)

	// ============================================================
	// 43. CHECK PERFECT NUMBER
	// ============================================================

	fmt.Println("\n========== Q43 ==========")

	num = 6
	sum = 0

	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	if sum == num {
		fmt.Println("Perfect number")
	} else {
		fmt.Println("Not a perfect number")
	}

	// ============================================================
	// 44. PRINT PERFECT NUMBERS FROM 1 TO 1000
	// ============================================================

	fmt.Println("\n========== Q44 ==========")

	for num := 1; num <= 1000; num++ {

		sum = 0

		for i := 1; i < num; i++ {
			if num%i == 0 {
				sum += i
			}
		}

		if sum == num {
			fmt.Println(num)
		}
	}

	// ============================================================
	// 45. FIBONACCI SERIES
	// ============================================================

	fmt.Println("\n========== Q45 ==========")

	n = 10

	first := 0
	second := 1

	for i := 0; i < n; i++ {
		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}

	fmt.Println()

	// ============================================================
	// 46. CHECK ARMSTRONG NUMBER
	// ============================================================

	fmt.Println("\n========== Q46 ==========")

	num = 153
	original = num
	temp = num
	sum = 0

	for temp != 0 {
		digit := temp % 10
		sum += digit * digit * digit
		temp /= 10
	}

	if sum == original {
		fmt.Println("Armstrong number")
	} else {
		fmt.Println("Not Armstrong")
	}

	// ============================================================
	// 47. PRINT ARMSTRONG NUMBERS FROM 1 TO 1000
	// ============================================================

	fmt.Println("\n========== Q47 ==========")

	for num := 1; num <= 1000; num++ {

		original := num
		temp := num
		sum := 0

		for temp != 0 {
			digit := temp % 10
			sum += digit * digit * digit
			temp /= 10
		}

		if sum == original {
			fmt.Println(num)
		}
	}

	// ============================================================
	// 48. FIND GCD
	// ============================================================

	fmt.Println("\n========== Q48 ==========")

	a := 48
	b := 18

	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}

	fmt.Println("GCD:", a)

	// ============================================================
	// 49. FIND LCM
	// ============================================================

	fmt.Println("\n========== Q49 ==========")

	a = 12
	b = 18

	originalA := a
	originalB := b

	for originalB != 0 {
		remainder := originalA % originalB
		originalA = originalB
		originalB = remainder
	}

	gcd := originalA

	lcm := (a * b) / gcd

	fmt.Println("LCM:", lcm)

	// ============================================================
	// 50. CHECK STRONG NUMBER
	// ============================================================

	fmt.Println("\n========== Q50 ==========")

	num = 145
	original = num
	sum = 0

	for num != 0 {

		digit := num % 10

		factorial := 1

		for i := 1; i <= digit; i++ {
			factorial *= i
		}

		sum += factorial
		num /= 10
	}

	if sum == original {
		fmt.Println("Strong number")
	} else {
		fmt.Println("Not strong number")
	}

	// ============================================================
	// 51. 5 x 5 STAR PATTERN
	// ============================================================

	fmt.Println("\n========== Q51 ==========")

	for i := 1; i <= 5; i++ {

		for j := 1; j <= 5; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	// ============================================================
	// 52. RIGHT TRIANGLE
	// ============================================================

	fmt.Println("\n========== Q52 ==========")

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	// ============================================================
	// 53. INVERTED TRIANGLE
	// ============================================================

	fmt.Println("\n========== Q53 ==========")

	for i := 5; i >= 1; i-- {

		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	// ============================================================
	// 54. NUMBER TRIANGLE
	// ============================================================

	fmt.Println("\n========== Q54 ==========")

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// ============================================================
	// 55. SAME NUMBER TRIANGLE
	// ============================================================

	fmt.Println("\n========== Q55 ==========")

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}

		fmt.Println()
	}

	// ============================================================
	// 56. INCREASING NUMBER PATTERN
	// ============================================================

	fmt.Println("\n========== Q56 ==========")

	count = 1

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(count, " ")
			count++
		}

		fmt.Println()
	}

	// ============================================================
	// 57. PYRAMID PATTERN
	// ============================================================

	fmt.Println("\n========== Q57 ==========")

	n = 5

	for i := 1; i <= n; i++ {

		for space := 1; space <= n-i; space++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// ============================================================
	// 58. INVERTED PYRAMID
	// ============================================================

	fmt.Println("\n========== Q58 ==========")

	n = 5

	for i := n; i >= 1; i-- {

		for space := 1; space <= n-i; space++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// ============================================================
	// 59. MULTIPLICATION TABLE GRID
	// ============================================================

	fmt.Println("\n========== Q59 ==========")

	for i := 1; i <= 10; i++ {

		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}

		fmt.Println()
	}

	// ============================================================
	// 60. PRINT COORDINATES
	// ============================================================

	fmt.Println("\n========== Q60 ==========")

	for x := 1; x <= 3; x++ {

		for y := 1; y <= 3; y++ {
			fmt.Printf("(%d,%d) ", x, y)
		}

		fmt.Println()
	}

	// ============================================================
	// 61. FIND LARGEST NUMBER IN SLICE
	// ============================================================

	fmt.Println("\n========== Q61 ==========")

	numbers := []int{10, 25, 7, 50, 18}

	largest = numbers[0]

	for i := 1; i < len(numbers); i++ {

		if numbers[i] > largest {
			largest = numbers[i]
		}
	}

	fmt.Println("Largest:", largest)

	// ============================================================
	// 62. FIND SMALLEST NUMBER IN SLICE
	// ============================================================

	fmt.Println("\n========== Q62 ==========")

	smallest = numbers[0]

	for i := 1; i < len(numbers); i++ {

		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	fmt.Println("Smallest:", smallest)

	// ============================================================
	// 63. FIND SECOND LARGEST NUMBER
	// ============================================================

	fmt.Println("\n========== Q63 ==========")

	numbers = []int{10, 25, 7, 50, 18}

	largest = numbers[0]
	secondLargest := numbers[0]

	for i := 1; i < len(numbers); i++ {

		if numbers[i] > largest {
			secondLargest = largest
			largest = numbers[i]

		} else if numbers[i] > secondLargest &&
			numbers[i] != largest {

			secondLargest = numbers[i]
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Second largest:", secondLargest)

	// ============================================================
	// 64. COUNT POSITIVE, NEGATIVE AND ZERO
	// ============================================================

	fmt.Println("\n========== Q64 ==========")

	numbers = []int{10, -5, 0, 20, -10, 0, 30}

	positive := 0
	negative := 0
	zero := 0

	for _, number := range numbers {

		if number > 0 {
			positive++

		} else if number < 0 {
			negative++

		} else {
			zero++
		}
	}

	fmt.Println("Positive:", positive)
	fmt.Println("Negative:", negative)
	fmt.Println("Zero:", zero)

	// ============================================================
	// 65. SEARCH FOR AN ELEMENT
	// ============================================================

	fmt.Println("\n========== Q65 ==========")

	numbers = []int{10, 20, 30, 40, 50}
	target := 30

	found := false

	for _, number := range numbers {

		if number == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	// ============================================================
	// 66. FIND INDEX OF AN ELEMENT
	// ============================================================

	fmt.Println("\n========== Q66 ==========")

	numbers = []int{10, 20, 30, 40, 50}
	target = 40

	index := -1

	for i := 0; i < len(numbers); i++ {

		if numbers[i] == target {
			index = i
			break
		}
	}

	fmt.Println("Index:", index)

	// ============================================================
	// 67. COUNT OCCURRENCE OF AN ELEMENT
	// ============================================================

	fmt.Println("\n========== Q67 ==========")

	numbers = []int{10, 20, 10, 30, 10, 40}
	target = 10

	count = 0

	for _, number := range numbers {

		if number == target {
			count++
		}
	}

	fmt.Println("Count:", count)

	// ============================================================
	// 68. PRINT DUPLICATE ELEMENTS
	// ============================================================

	fmt.Println("\n========== Q68 ==========")

	numbers = []int{10, 20, 10, 30, 20, 40}

	for i := 0; i < len(numbers); i++ {

		for j := i + 1; j < len(numbers); j++ {

			if numbers[i] == numbers[j] {
				fmt.Println("Duplicate:", numbers[i])
				break
			}
		}
	}

	// ============================================================
	// 69. REMOVE DUPLICATES FROM SLICE
	// ============================================================

	fmt.Println("\n========== Q69 ==========")

	numbers = []int{10, 20, 10, 30, 20, 40}

	resultSlice := []int{}

	for _, number := range numbers {

		found := false

		for _, existing := range resultSlice {

			if existing == number {
				found = true
				break
			}
		}

		if !found {
			resultSlice = append(resultSlice, number)
		}
	}

	fmt.Println("Original:", numbers)
	fmt.Println("Without duplicates:", resultSlice)

	// ============================================================
	// 70. REVERSE A SLICE
	// ============================================================

	fmt.Println("\n========== Q70 ==========")

	numbers = []int{10, 20, 30, 40, 50}

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println("Reversed:", numbers)

	// ============================================================
	// END
	// ============================================================

	fmt.Println("\n========================================")
	fmt.Println("ALL 70 LOOP QUESTIONS COMPLETED")
	fmt.Println("========================================")
}
