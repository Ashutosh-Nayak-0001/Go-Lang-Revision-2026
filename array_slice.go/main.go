/*
====================================================================
GO ARRAYS + SLICES
70 CODING QUESTIONS WITH ANSWERS + DEEP EXPLANATIONS
====================================================================

Run:
    go run arrays_slices_70.go

This file covers:
- Go arrays
- Go slices
- Array vs slice comparison
- Indexing
- Traversal
- Input
- Searching
- Sorting
- Duplicates
- Frequency
- Rotation
- Missing numbers
- Two-sum
- Subarrays
- Majority element
- Stock profit

IMPORTANT:
Arrays and slices are NOT the same in Go.

ARRAY:
    var a [5]int

SLICE:
    var s []int

An array has a fixed length that is part of its type.
A slice is a dynamic, flexible view over an underlying array.

====================================================================
ARRAY VS SLICE — QUICK COMPARISON
====================================================================

1. Size
   Array: fixed
       [5]int and [10]int are different types.

   Slice: dynamic
       []int can grow and shrink.

2. Type
   Array:
       [3]int
       [5]int
       [10]int
   These are different types.

   Slice:
       []int
   All []int slices have the same slice type.

3. Assignment
   Arrays are value types.

       a := [3]int{1,2,3}
       b := a
       b[0] = 100

   a is still [1 2 3].

   Slices describe an underlying array. Copying a slice copies
   the slice descriptor, so two slices can refer to the same
   underlying array.

4. append
   Arrays do NOT support append.

   Slices support:
       s = append(s, 10)

5. len
   Both support len().

6. cap
   Slices have capacity:
       cap(s)

   Arrays don't normally use cap() as a slice concept.

7. Function passing
   Arrays are copied when passed by value.

   Slices are small descriptors passed by value, but they refer
   to an underlying array. Therefore element modifications can
   be visible to the caller.

8. Comparison
   Arrays can be compared with == if their element type is
   comparable.

       [3]int{1,2,3} == [3]int{1,2,3}

   Slices cannot be compared using == except against nil.

       s == nil       // valid
       s1 == s2       // invalid

9. Memory
   An array contains its elements directly.

   A slice contains a pointer to an underlying array, length,
   and capacity.

10. When to use
    Array:
        fixed-size data where size is known and part of the
        design.

    Slice:
        almost all normal collection processing in Go,
        especially when size changes or is unknown.

====================================================================
*/

package main

import (
	"errors"
	"fmt"
	"sort"
)

/*
====================================================================
SECTION 1 — BASIC ARRAY QUESTIONS
====================================================================
*/

/*
Q01. Declare and initialize an integer array.

Answer:
Use [N]Type{...}.

Explanation:
[5]int means:
- exactly 5 elements
- each element is int

The length 5 is part of the array type.
*/
func q01() {
	fmt.Println("\nQ01. Declare and initialize array")

	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers)
}

/*
Q02. Print all elements of an array.

Answer:
Use a loop.

Explanation:
Arrays are indexed from 0.

For [5]int:
indexes are 0,1,2,3,4.
*/
func q02() {
	fmt.Println("\nQ02. Print array elements")

	numbers := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

/*
Q03. Take 5 numbers as input and store them in an array.

Answer:
Use fmt.Scan.

Note:
The function demonstrates the input pattern, but for the large
combined file we use sample values so running the entire file
does not pause waiting for terminal input.
*/
func q03() {
	fmt.Println("\nQ03. Store input-like values in array")

	var numbers [5]int

	values := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		numbers[i] = values[i]
	}

	fmt.Println(numbers)

	/*
		Real interactive input would be:

		for i := 0; i < len(numbers); i++ {
			fmt.Scan(&numbers[i])
		}
	*/
}

/*
Q04. Find the length of an array.

Answer:
len(array)

Important:
The result is the number of elements.
*/
func q04() {
	fmt.Println("\nQ04. Array length")

	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Length:", len(numbers))
}

/*
Q05. Find first and last elements.

Answer:
First = array[0]
Last  = array[len(array)-1]
*/
func q05() {
	fmt.Println("\nQ05. First and last element")

	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("First:", numbers[0])
	fmt.Println("Last:", numbers[len(numbers)-1])
}

/*
Q06. Print array using range.

range gives index and value.

    for index, value := range numbers
*/
func q06() {
	fmt.Println("\nQ06. Range over array")

	numbers := [5]int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Println(index, value)
	}
}

/*
Q07. Find sum of array elements.
*/
func q07() {
	fmt.Println("\nQ07. Sum of array")

	numbers := [5]int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	fmt.Println("Sum:", sum)
}

/*
Q08. Find average of array elements.

Important:
Use float64 if you want a decimal result.
*/
func q08() {
	fmt.Println("\nQ08. Average of array")

	numbers := [5]int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Println("Average:", average)
}

/*
Q09. Find largest element.
*/
func q09() {
	fmt.Println("\nQ09. Largest element")

	numbers := [5]int{10, 50, 20, 5, 30}

	max := numbers[0]

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	fmt.Println("Largest:", max)
}

/*
Q10. Find smallest element.
*/
func q10() {
	fmt.Println("\nQ10. Smallest element")

	numbers := [5]int{10, 50, 20, 5, 30}

	min := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}
	}

	fmt.Println("Smallest:", min)
}

/*
Q11. Count even numbers.
*/
func q11() {
	fmt.Println("\nQ11. Count even numbers")

	numbers := [6]int{1, 2, 3, 4, 5, 6}

	count := 0

	for _, value := range numbers {
		if value%2 == 0 {
			count++
		}
	}

	fmt.Println("Even count:", count)
}

/*
Q12. Count odd numbers.
*/
func q12() {
	fmt.Println("\nQ12. Count odd numbers")

	numbers := [6]int{1, 2, 3, 4, 5, 6}

	count := 0

	for _, value := range numbers {
		if value%2 != 0 {
			count++
		}
	}

	fmt.Println("Odd count:", count)
}

/*
Q13. Print only even elements.
*/
func q13() {
	fmt.Println("\nQ13. Even elements")

	numbers := [6]int{1, 2, 3, 4, 5, 6}

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}

/*
Q14. Print only odd elements.
*/
func q14() {
	fmt.Println("\nQ14. Odd elements")

	numbers := [6]int{1, 2, 3, 4, 5, 6}

	for _, value := range numbers {
		if value%2 != 0 {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}

/*
Q15. Search for an element.

Return true when found.
*/
func q15() {
	fmt.Println("\nQ15. Search element")

	numbers := [5]int{10, 20, 30, 40, 50}
	target := 30

	found := false

	for _, value := range numbers {
		if value == target {
			found = true
			break
		}
	}

	fmt.Println("Found:", found)
}

/*
Q16. Count occurrences of an element.
*/
func q16() {
	fmt.Println("\nQ16. Count occurrence")

	numbers := [7]int{1, 2, 2, 3, 2, 4, 2}
	target := 2

	count := 0

	for _, value := range numbers {
		if value == target {
			count++
		}
	}

	fmt.Println("Count:", count)
}

/*
Q17. Copy one array into another.

Important:
Array assignment copies all elements.

*/
func q17() {
	fmt.Println("\nQ17. Copy array")

	a := [3]int{10, 20, 30}
	b := a

	b[0] = 100

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	/*
		Output:
		a: [10 20 30]
		b: [100 20 30]

		This proves arrays are value types.
	*/
}

/*
Q18. Compare two arrays.

Arrays can be compared using == if the element type is
comparable.
*/
func q18() {
	fmt.Println("\nQ18. Compare arrays")

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}

	fmt.Println("Equal:", a == b)
}

/*
Q19. Sum elements at even indexes.
*/
func q19() {
	fmt.Println("\nQ19. Sum at even indexes")

	numbers := [6]int{10, 20, 30, 40, 50, 60}

	sum := 0

	for i := 0; i < len(numbers); i += 2 {
		sum += numbers[i]
	}

	fmt.Println("Sum:", sum)
}

/*
Q20. Sum elements at odd indexes.
*/
func q20() {
	fmt.Println("\nQ20. Sum at odd indexes")

	numbers := [6]int{10, 20, 30, 40, 50, 60}

	sum := 0

	for i := 1; i < len(numbers); i += 2 {
		sum += numbers[i]
	}

	fmt.Println("Sum:", sum)
}

/*
====================================================================
SECTION 2 — ARRAY CODING QUESTIONS
====================================================================
*/

/*
Q21. Reverse an array.

Use two pointers:
left  -> beginning
right -> end

Swap and move toward the center.
*/
func q21() {
	fmt.Println("\nQ21. Reverse array")

	numbers := [5]int{1, 2, 3, 4, 5}

	for left, right := 0, len(numbers)-1; left < right; left, right = left+1, right-1 {
		numbers[left], numbers[right] = numbers[right], numbers[left]
	}

	fmt.Println(numbers)
}

/*
Q22. Print array in reverse without modifying it.
*/
func q22() {
	fmt.Println("\nQ22. Print reverse without modifying")

	numbers := [5]int{1, 2, 3, 4, 5}

	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}

	fmt.Println()
	fmt.Println("Original:", numbers)
}

/*
Q23. Find second largest element.

Approach:
Track largest and second largest.

This version assumes at least two distinct values.
*/
func q23() {
	fmt.Println("\nQ23. Second largest")

	numbers := [6]int{10, 50, 20, 40, 30, 5}

	max1 := -int(^uint(0)>>1) - 1
	max2 := max1

	for _, value := range numbers {
		if value > max1 {
			max2 = max1
			max1 = value
		} else if value > max2 && value != max1 {
			max2 = value
		}
	}

	fmt.Println("Second largest:", max2)
}

/*
Q24. Find second smallest element.
*/
func q24() {
	fmt.Println("\nQ24. Second smallest")

	numbers := [6]int{10, 50, 20, 40, 30, 5}

	min1 := int(^uint(0)>>1)
	min2 := min1

	for _, value := range numbers {
		if value < min1 {
			min2 = min1
			min1 = value
		} else if value < min2 && value != min1 {
			min2 = value
		}
	}

	fmt.Println("Second smallest:", min2)
}

/*
Q25. Find maximum and minimum together.
*/
func q25() {
	fmt.Println("\nQ25. Min and max together")

	numbers := [5]int{10, 50, 20, 5, 30}

	min := numbers[0]
	max := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

/*
Q26. Find index of largest element.
*/
func q26() {
	fmt.Println("\nQ26. Index of largest")

	numbers := [5]int{10, 50, 20, 5, 30}

	maxIndex := 0

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[maxIndex] {
			maxIndex = i
		}
	}

	fmt.Println("Index:", maxIndex)
	fmt.Println("Value:", numbers[maxIndex])
}

/*
Q27. Find index of smallest element.
*/
func q27() {
	fmt.Println("\nQ27. Index of smallest")

	numbers := [5]int{10, 50, 20, 5, 30}

	minIndex := 0

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[minIndex] {
			minIndex = i
		}
	}

	fmt.Println("Index:", minIndex)
	fmt.Println("Value:", numbers[minIndex])
}

/*
Q28. Replace negative numbers with zero.
*/
func q28() {
	fmt.Println("\nQ28. Replace negatives")

	numbers := [6]int{10, -5, 20, -2, 30, -10}

	for i := range numbers {
		if numbers[i] < 0 {
			numbers[i] = 0
		}
	}

	fmt.Println(numbers)
}

/*
Q29. Replace even numbers with their square.
*/
func q29() {
	fmt.Println("\nQ29. Square even numbers")

	numbers := [5]int{1, 2, 3, 4, 5}

	for i := range numbers {
		if numbers[i]%2 == 0 {
			numbers[i] *= numbers[i]
		}
	}

	fmt.Println(numbers)
}

/*
Q30. Count positive, negative, and zero.
*/
func q30() {
	fmt.Println("\nQ30. Positive negative zero")

	numbers := [7]int{-2, 0, 5, -1, 0, 10, 3}

	positive := 0
	negative := 0
	zero := 0

	for _, value := range numbers {
		switch {
		case value > 0:
			positive++
		case value < 0:
			negative++
		default:
			zero++
		}
	}

	fmt.Println("Positive:", positive)
	fmt.Println("Negative:", negative)
	fmt.Println("Zero:", zero)
}

/*
Q31. Separate positive and negative numbers.
*/
func q31() {
	fmt.Println("\nQ31. Separate positive and negative")

	numbers := [7]int{-2, 0, 5, -1, 0, 10, 3}

	positive := []int{}
	negative := []int{}

	for _, value := range numbers {
		if value > 0 {
			positive = append(positive, value)
		} else if value < 0 {
			negative = append(negative, value)
		}
	}

	fmt.Println("Positive:", positive)
	fmt.Println("Negative:", negative)
}

/*
Q32. Print elements greater than target.
*/
func q32() {
	fmt.Println("\nQ32. Elements greater than target")

	numbers := [6]int{10, 20, 5, 40, 15, 30}
	target := 20

	for _, value := range numbers {
		if value > target {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}

/*
Q33. Print elements smaller than target.
*/
func q33() {
	fmt.Println("\nQ33. Elements smaller than target")

	numbers := [6]int{10, 20, 5, 40, 15, 30}
	target := 20

	for _, value := range numbers {
		if value < target {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}

/*
Q34. Find all duplicate elements.

Basic O(n^2) solution.
*/
func q34() {
	fmt.Println("\nQ34. Duplicate elements")

	numbers := [7]int{1, 2, 3, 2, 4, 1, 5}

	duplicates := []int{}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				alreadyAdded := false

				for _, d := range duplicates {
					if d == numbers[i] {
						alreadyAdded = true
						break
					}
				}

				if !alreadyAdded {
					duplicates = append(duplicates, numbers[i])
				}
			}
		}
	}

	fmt.Println("Duplicates:", duplicates)
}

/*
Q35. Find unique elements.

An element is unique if its frequency is exactly 1.
*/
func q35() {
	fmt.Println("\nQ35. Unique elements")

	numbers := [7]int{1, 2, 2, 3, 4, 4, 5}

	for i := 0; i < len(numbers); i++ {
		count := 0

		for j := 0; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				count++
			}
		}

		if count == 1 {
			fmt.Print(numbers[i], " ")
		}
	}

	fmt.Println()
}

/*
Q36. Remove duplicates.

Use a result slice and check whether the value already exists.
*/
func q36() {
	fmt.Println("\nQ36. Remove duplicates")

	numbers := [8]int{1, 2, 2, 3, 4, 3, 5, 1}

	result := []int{}

	for _, value := range numbers {
		exists := false

		for _, existing := range result {
			if existing == value {
				exists = true
				break
			}
		}

		if !exists {
			result = append(result, value)
		}
	}

	fmt.Println("Without duplicates:", result)
}

/*
Q37. Check whether array contains duplicates.
*/
func q37() {
	fmt.Println("\nQ37. Contains duplicate")

	numbers := [6]int{1, 2, 3, 4, 2, 5}

	found := false

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Println("Contains duplicate:", found)
}

/*
Q38. Check whether array is sorted ascending.
*/
func q38() {
	fmt.Println("\nQ38. Check ascending sorted")

	numbers := [5]int{1, 2, 3, 4, 5}

	sorted := true

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println("Sorted:", sorted)
}

/*
Q39. Check whether array is sorted descending.
*/
func q39() {
	fmt.Println("\nQ39. Check descending sorted")

	numbers := [5]int{5, 4, 3, 2, 1}

	sorted := true

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println("Sorted descending:", sorted)
}

/*
Q40. Sort an array without sort.Ints.

We use bubble sort.

Important:
Array length is fixed, but its elements can be modified.
*/
func q40() {
	fmt.Println("\nQ40. Bubble sort array")

	numbers := [5]int{5, 2, 8, 1, 3}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println(numbers)
}

/*
====================================================================
SECTION 3 — FREQUENCY, SEARCHING, SET-STYLE QUESTIONS
====================================================================
*/

/*
Q41. Sort array descending without sort.Ints.
*/
func q41() {
	fmt.Println("\nQ41. Descending sort")

	numbers := [5]int{5, 2, 8, 1, 3}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println(numbers)
}

/*
Q42. Frequency of every element.

Use a map.

This is much better than nested loops for frequency counting.

Time complexity:
O(n) average.
*/
func q42() {
	fmt.Println("\nQ42. Frequency map")

	numbers := [8]int{1, 2, 2, 3, 1, 4, 2, 3}

	frequency := make(map[int]int)

	for _, value := range numbers {
		frequency[value]++
	}

	fmt.Println(frequency)
}

/*
Q43. Most frequent element.
*/
func q43() {
	fmt.Println("\nQ43. Most frequent element")

	numbers := [8]int{1, 2, 2, 3, 1, 4, 2, 3}

	frequency := make(map[int]int)

	for _, value := range numbers {
		frequency[value]++
	}

	maxValue := numbers[0]
	maxCount := 0

	for value, count := range frequency {
		if count > maxCount {
			maxCount = count
			maxValue = value
		}
	}

	fmt.Println("Value:", maxValue)
	fmt.Println("Frequency:", maxCount)
}

/*
Q44. Least frequent element.
*/
func q44() {
	fmt.Println("\nQ44. Least frequent element")

	numbers := [8]int{1, 2, 2, 3, 1, 4, 2, 3}

	frequency := make(map[int]int)

	for _, value := range numbers {
		frequency[value]++
	}

	minValue := numbers[0]
	minCount := len(numbers) + 1

	for value, count := range frequency {
		if count < minCount {
			minCount = count
			minValue = value
		}
	}

	fmt.Println("Value:", minValue)
	fmt.Println("Frequency:", minCount)
}

/*
Q45. First repeated element.

Example:
1 2 3 2 4
The first element that repeats is 2.
*/
func q45() {
	fmt.Println("\nQ45. First repeated element")

	numbers := [6]int{1, 2, 3, 2, 4, 5}

	seen := make(map[int]bool)
	result := -1

	for _, value := range numbers {
		if seen[value] {
			result = value
			break
		}

		seen[value] = true
	}

	fmt.Println("First repeated:", result)
}

/*
Q46. First non-repeated element.
*/
func q46() {
	fmt.Println("\nQ46. First non-repeated element")

	numbers := [7]int{2, 3, 2, 4, 3, 5, 4}

	frequency := make(map[int]int)

	for _, value := range numbers {
		frequency[value]++
	}

	result := -1

	for _, value := range numbers {
		if frequency[value] == 1 {
			result = value
			break
		}
	}

	fmt.Println("First non-repeated:", result)
}

/*
Q47. Find missing number from 1..N.

Example:
[1,2,3,5]
N = 5
Missing = 4

Formula:
1+2+...+N = N*(N+1)/2

Missing = expected - actual
*/
func q47() {
	fmt.Println("\nQ47. Missing number")

	numbers := []int{1, 2, 3, 5}
	n := 5

	expected := n * (n + 1) / 2

	actual := 0

	for _, value := range numbers {
		actual += value
	}

	fmt.Println("Missing:", expected-actual)
}

/*
Q48. Two numbers with target sum.

Use a map for O(n) average time.

For every x, check whether target-x was already seen.
*/
func q48() {
	fmt.Println("\nQ48. Two sum")

	numbers := []int{2, 7, 11, 15}
	target := 9

	seen := make(map[int]bool)

	for _, value := range numbers {
		needed := target - value

		if seen[needed] {
			fmt.Println("Pair:", needed, value)
			return
		}

		seen[value] = true
	}

	fmt.Println("No pair found")
}

/*
Q49. Three numbers with target sum.

Simple O(n^3) interview solution.
*/
func q49() {
	fmt.Println("\nQ49. Three sum")

	numbers := []int{1, 2, 3, 4, 5, 6}
	target := 12

	found := false

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == target {
					fmt.Println("Triplet:", numbers[i], numbers[j], numbers[k])
					found = true
					return
				}
			}
		}
	}

	if !found {
		fmt.Println("No triplet")
	}
}

/*
Q50. Intersection of two arrays.

Intersection means values common to both.
*/
func q50() {
	fmt.Println("\nQ50. Intersection")

	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	set := make(map[int]bool)

	for _, value := range a {
		set[value] = true
	}

	result := []int{}

	for _, value := range b {
		if set[value] {
			result = append(result, value)
			delete(set, value)
		}
	}

	fmt.Println("Intersection:", result)
}

/*
Q51. Union of two arrays.

Union contains every distinct value from both.
*/
func q51() {
	fmt.Println("\nQ51. Union")

	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}

	set := make(map[int]bool)

	for _, value := range a {
		set[value] = true
	}

	for _, value := range b {
		set[value] = true
	}

	result := []int{}

	for value := range set {
		result = append(result, value)
	}

	sort.Ints(result)

	fmt.Println("Union:", result)
}

/*
Q52. Common elements among three arrays.
*/
func q52() {
	fmt.Println("\nQ52. Common in three arrays")

	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 4, 6}
	c := []int{0, 2, 3, 4}

	setB := make(map[int]bool)
	setC := make(map[int]bool)

	for _, value := range b {
		setB[value] = true
	}

	for _, value := range c {
		setC[value] = true
	}

	result := []int{}

	for _, value := range a {
		if setB[value] && setC[value] {
			result = append(result, value)
		}
	}

	fmt.Println("Common:", result)
}

/*
Q53. Find index of target.

Return -1 if not found.
*/
func q53() {
	fmt.Println("\nQ53. Find target index")

	numbers := []int{10, 20, 30, 40}
	target := 30

	index := -1

	for i, value := range numbers {
		if value == target {
			index = i
			break
		}
	}

	fmt.Println("Index:", index)
}

/*
Q54. Count elements greater than average.
*/
func q54() {
	fmt.Println("\nQ54. Count above average")

	numbers := []int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	average := float64(sum) / float64(len(numbers))

	count := 0

	for _, value := range numbers {
		if float64(value) > average {
			count++
		}
	}

	fmt.Println("Average:", average)
	fmt.Println("Above average:", count)
}

/*
Q55. Find pair with maximum product.

For positive/negative mixed arrays, a robust solution can compare
all pairs. This is O(n^2).

A more optimized solution depends on the input constraints.
*/
func q55() {
	fmt.Println("\nQ55. Maximum product pair")

	numbers := []int{-10, -20, 1, 2, 3}

	maxProduct := numbers[0] * numbers[1]
	var pairA, pairB = numbers[0], numbers[1]

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			product := numbers[i] * numbers[j]

			if product > maxProduct {
				maxProduct = product
				pairA = numbers[i]
				pairB = numbers[j]
			}
		}
	}

	fmt.Println("Pair:", pairA, pairB)
	fmt.Println("Product:", maxProduct)
}

/*
====================================================================
SECTION 4 — INTERVIEW LEVEL
====================================================================
*/

/*
Q56. Move all zeros to the end.

Input:
[0,1,0,3,12]

Output:
[1,3,12,0,0]

Approach:
Keep a position for the next non-zero element.
*/
func q56() {
	fmt.Println("\nQ56. Move zeros to end")

	numbers := []int{0, 1, 0, 3, 12}

	position := 0

	for _, value := range numbers {
		if value != 0 {
			numbers[position] = value
			position++
		}
	}

	for position < len(numbers) {
		numbers[position] = 0
		position++
	}

	fmt.Println(numbers)
}

/*
Q57. Move negative numbers to the beginning.

This version preserves relative order.

Example:
[1,-2,3,-4,5]
-> [-2,-4,1,3,5]
*/
func q57() {
	fmt.Println("\nQ57. Move negatives to beginning")

	numbers := []int{1, -2, 3, -4, 5}

	result := []int{}

	for _, value := range numbers {
		if value < 0 {
			result = append(result, value)
		}
	}

	for _, value := range numbers {
		if value >= 0 {
			result = append(result, value)
		}
	}

	fmt.Println(result)
}

/*
Q58. Move zeros to beginning.
*/
func q58() {
	fmt.Println("\nQ58. Move zeros to beginning")

	numbers := []int{1, 0, 2, 0, 3, 4}

	result := []int{}

	for _, value := range numbers {
		if value == 0 {
			result = append(result, value)
		}
	}

	for _, value := range numbers {
		if value != 0 {
			result = append(result, value)
		}
	}

	fmt.Println(result)
}

/*
Q59. Rotate array left by one.

[1,2,3,4,5]
-> [2,3,4,5,1]
*/
func q59() {
	fmt.Println("\nQ59. Left rotate by one")

	numbers := []int{1, 2, 3, 4, 5}

	first := numbers[0]

	for i := 0; i < len(numbers)-1; i++ {
		numbers[i] = numbers[i+1]
	}

	numbers[len(numbers)-1] = first

	fmt.Println(numbers)
}

/*
Q60. Rotate array right by one.

[1,2,3,4,5]
-> [5,1,2,3,4]
*/
func q60() {
	fmt.Println("\nQ60. Right rotate by one")

	numbers := []int{1, 2, 3, 4, 5}

	last := numbers[len(numbers)-1]

	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i] = numbers[i-1]
	}

	numbers[0] = last

	fmt.Println(numbers)
}

/*
Q61. Left rotate by K positions.

Example:
[1,2,3,4,5], k=2
-> [3,4,5,1,2]

We use:
k = k % n

so rotating by 7 positions in a 5-element array is the same
as rotating by 2.
*/
func q61() {
	fmt.Println("\nQ61. Left rotate by K")

	numbers := []int{1, 2, 3, 4, 5}
	k := 2

	k %= len(numbers)

	result := append([]int{}, numbers[k:]...)
	result = append(result, numbers[:k]...)

	fmt.Println(result)
}

/*
Q62. Right rotate by K positions.

Example:
[1,2,3,4,5], k=2
-> [4,5,1,2,3]
*/
func q62() {
	fmt.Println("\nQ62. Right rotate by K")

	numbers := []int{1, 2, 3, 4, 5}
	k := 2

	k %= len(numbers)

	result := append([]int{}, numbers[len(numbers)-k:]...)
	result = append(result, numbers[:len(numbers)-k]...)

	fmt.Println(result)
}

/*
Q63. Maximum sum subarray.

Kadane's algorithm.

Example:
[-2,1,-3,4,-1,2,1,-5,4]

Maximum sum = 6
Subarray = [4,-1,2,1]

Core idea:
currentSum = max(value, currentSum+value)
maxSum = max(maxSum, currentSum)

Time:
O(n)
*/
func q63() {
	fmt.Println("\nQ63. Maximum subarray sum")

	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	currentSum := numbers[0]
	maxSum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		currentSum = max(numbers[i], currentSum+numbers[i])

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	fmt.Println("Maximum sum:", maxSum)
}

/*
Q64. Minimum sum subarray.

This is the opposite version of Kadane's algorithm.
*/
func q64() {
	fmt.Println("\nQ64. Minimum subarray sum")

	numbers := []int{3, -4, 2, -3, -1, 7}

	currentSum := numbers[0]
	minSum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		currentSum = min(numbers[i], currentSum+numbers[i])

		if currentSum < minSum {
			minSum = currentSum
		}
	}

	fmt.Println("Minimum sum:", minSum)
}

/*
Q65. Longest increasing contiguous subarray.

Contiguous means the elements must be next to each other.

Example:
[1,2,3,2,4,5,6]

Longest increasing contiguous sequence:
[2,4,5,6] -> length 4
*/
func q65() {
	fmt.Println("\nQ65. Longest increasing contiguous subarray")

	numbers := []int{1, 2, 3, 2, 4, 5, 6}

	if len(numbers) == 0 {
		fmt.Println(0)
		return
	}

	current := 1
	longest := 1

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			current++
		} else {
			current = 1
		}

		if current > longest {
			longest = current
		}
	}

	fmt.Println("Longest length:", longest)
}

/*
Q66. Longest decreasing contiguous subarray.
*/
func q66() {
	fmt.Println("\nQ66. Longest decreasing contiguous subarray")

	numbers := []int{9, 7, 5, 6, 4, 3, 2}

	if len(numbers) == 0 {
		fmt.Println(0)
		return
	}

	current := 1
	longest := 1

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			current++
		} else {
			current = 1
		}

		if current > longest {
			longest = current
		}
	}

	fmt.Println("Longest length:", longest)
}

/*
Q67. Pair with minimum difference.

Sort first, then compare adjacent elements.

Why only adjacent after sorting?
The closest pair must be adjacent in sorted order.
*/
func q67() {
	fmt.Println("\nQ67. Minimum difference pair")

	numbers := []int{10, 3, 20, 8, 15}

	sort.Ints(numbers)

	minDiff := numbers[1] - numbers[0]
	a := numbers[0]
	b := numbers[1]

	for i := 1; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]

		if diff < minDiff {
			minDiff = diff
			a = numbers[i]
			b = numbers[i+1]
		}
	}

	fmt.Println("Pair:", a, b)
	fmt.Println("Difference:", minDiff)
}

/*
Q68. Majority element.

An element is a majority element if it occurs more than n/2 times.

Boyer-Moore Voting Algorithm:
- keep candidate
- keep count
- same candidate -> count++
- different -> count--
- count reaches zero -> choose new candidate

After finding candidate, verify it.
*/
func q68() {
	fmt.Println("\nQ68. Majority element")

	numbers := []int{2, 2, 1, 1, 1, 2, 2}

	candidate := 0
	count := 0

	for _, value := range numbers {
		if count == 0 {
			candidate = value
			count = 1
		} else if value == candidate {
			count++
		} else {
			count--
		}
	}

	occurrences := 0

	for _, value := range numbers {
		if value == candidate {
			occurrences++
		}
	}

	if occurrences > len(numbers)/2 {
		fmt.Println("Majority:", candidate)
	} else {
		fmt.Println("No majority element")
	}
}

/*
Q69. Maximum profit from buying and selling stock once.

Example:
[7,1,5,3,6,4]

Buy at 1
Sell at 6
Profit = 5

Algorithm:
Track minimum price seen so far.
For every price:
profit = price - minimumPrice
keep maximum profit.

Time:
O(n)
*/
func q69() {
	fmt.Println("\nQ69. Maximum stock profit")

	prices := []int{7, 1, 5, 3, 6, 4}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	fmt.Println("Maximum profit:", maxProfit)
}

/*
Q70. Deep array vs slice practical example.

This is the most important comparison exercise.

Part A:
Array copy

Part B:
Slice copy

Part C:
append

Part D:
Passing to function

This question is extremely useful for Go interviews.
*/
func q70() {
	fmt.Println("\nQ70. Array vs Slice practical comparison")

	// ---------------- ARRAY ----------------

	a := [3]int{1, 2, 3}
	b := a

	b[0] = 100

	fmt.Println("Array a:", a)
	fmt.Println("Array b:", b)

	/*
		a remains [1 2 3]

		Why?
		Because b = a copies the entire array.
	*/

	// ---------------- SLICE ----------------

	s1 := []int{1, 2, 3}
	s2 := s1

	s2[0] = 100

	fmt.Println("Slice s1:", s1)
	fmt.Println("Slice s2:", s2)

	/*
		Both usually show:

		    [100 2 3]

		Why?
		s1 and s2 are slice descriptors referring to the same
		underlying array in this example.

		IMPORTANT:
		A slice itself is passed by value. But its descriptor
		points to underlying storage.

		Therefore modifying an existing element can be visible
		through another slice sharing that storage.
	*/

	// ---------------- APPEND ----------------

	s := []int{1, 2, 3}

	s = append(s, 4, 5)

	fmt.Println("After append:", s)

	/*
		append may reuse existing capacity or allocate a new
		underlying array.

		Therefore never assume append always modifies the same
		underlying array.

		Always use:
		    s = append(s, value)
	*/

	// ---------------- NIL SLICE ----------------

	var nilSlice []int

	fmt.Println("Nil slice:", nilSlice)
	fmt.Println("Nil slice == nil:", nilSlice == nil)

	/*
		A nil slice is valid.

		    var s []int

		Then:
		    len(s) == 0
		    cap(s) == 0
		    s == nil

		You can append to it safely.
	*/

	nilSlice = append(nilSlice, 10)

	fmt.Println("After append:", nilSlice)
}

/*
====================================================================
HELPER FUNCTIONS
====================================================================
*/

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
====================================================================
ARRAY VS SLICE INTERVIEW QUESTIONS
====================================================================

1. What is the difference between [5]int and []int?
2. Why are [5]int and [10]int different types?
3. Can an array grow?
4. Can a slice grow?
5. What does append do?
6. What is len?
7. What is cap?
8. Can arrays be compared using ==?
9. Can slices be compared using ==?
10. Why can a slice be nil?
11. What happens when an array is assigned to another array?
12. What happens when a slice is assigned to another slice?
13. What is the underlying array of a slice?
14. What are pointer, length and capacity in a slice?
15. Why does modifying a slice element sometimes affect another slice?
16. Why can append cause a new underlying array?
17. What happens when a slice is passed to a function?
18. What happens when an array is passed to a function?
19. What is a slice expression: a[1:4]?
20. What is the difference between len and cap?

KEY ANSWERS:

Array:
    [N]T

Slice:
    []T

Slice descriptor concept:
    pointer + length + capacity

Array assignment:
    copies elements

Slice assignment:
    copies descriptor; underlying storage may be shared

Array comparison:
    possible with == when elements are comparable

Slice comparison:
    only == nil is allowed; use slices.Equal or manual comparison
    for comparing two slices.

Slice append:
    s = append(s, value)

Slice expression:
    s[low:high]

The high index is excluded.

Example:
    s := []int{10,20,30,40,50}
    s[1:4]
    -> [20 30 40]

====================================================================
COMMON INTERVIEW TRAPS
====================================================================

TRAP 1:
    a := [3]int{1,2,3}
    b := a
    b[0] = 99

    a is unchanged.

TRAP 2:
    a := []int{1,2,3}
    b := a
    b[0] = 99

    a may now show [99 2 3] because both slices share the
    underlying array.

TRAP 3:
    s := []int{1,2,3}
    s = append(s, 4)

    append returns a slice. Always capture the returned value.

TRAP 4:
    s1 == s2

    Invalid for slices.

Use:
    slices.Equal(s1, s2)
or a manual comparison depending on Go version/project needs.

TRAP 5:
    [3]int and [4]int are different types.

TRAP 6:
    len(s) is current number of elements.

    cap(s) is the capacity available from the slice's starting
    position before another allocation may be required.

====================================================================
MAIN
====================================================================
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

	q51()
	q52()
	q53()
	q54()
	q55()
	q56()
	q57()
	q58()
	q59()
	q60()

	q61()
	q62()
	q63()
	q64()
	q65()
	q66()
	q67()
	q68()
	q69()
	q70()

	fmt.Println("\n============================================================")
	fmt.Println("ALL 70 ARRAY + SLICE QUESTIONS COMPLETED")
	fmt.Println("============================================================")
	fmt.Println("Next practice target: Strings + Maps + Structs")
}

/*
====================================================================
END
====================================================================
*/
