package main

import (
	"fmt"
	"sort"
)

/*
GO MAPS — 70 QUESTIONS WITH ANSWERS
===================================

Run:
    go run maps_70.go

IMPORTANT:
- A map stores key/value pairs.
- Map keys must be comparable.
- A nil map can be read from, but cannot be written to.
- Use make(map[K]V) before inserting.
- Map iteration order is NOT guaranteed.
- Use value, ok := m[key] to check existence.
*/

// Q01. Declare a map.
func q01() {
	m := map[string]int{"Alice": 90, "Bob": 80}
	fmt.Println("\nQ01:", m)
}

// Q02. Create an empty map using make.
func q02() {
	m := make(map[string]int)
	fmt.Println("\nQ02:", m)
}

// Q03. Insert a key/value.
func q03() {
	m := make(map[string]int)
	m["Go"] = 10
	fmt.Println("\nQ03:", m)
}

// Q04. Update a value.
func q04() {
	m := map[string]int{"Go": 10}
	m["Go"] = 20
	fmt.Println("\nQ04:", m)
}

// Q05. Read a value.
func q05() {
	m := map[string]int{"Go": 10}
	fmt.Println("\nQ05:", m["Go"])
}

// Q06. Check whether a key exists.
func q06() {
	m := map[string]int{"Go": 10}
	value, ok := m["Go"]
	fmt.Println("\nQ06 value:", value, "exists:", ok)
}

// Q07. Check a missing key.
func q07() {
	m := map[string]int{"Go": 10}
	value, ok := m["Java"]
	fmt.Println("\nQ07 value:", value, "exists:", ok)
}

// Q08. Delete a key.
func q08() {
	m := map[string]int{"Go": 10, "Java": 20}
	delete(m, "Java")
	fmt.Println("\nQ08:", m)
}

// Q09. Find map length.
func q09() {
	m := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println("\nQ09 length:", len(m))
}

// Q10. Iterate over a map.
func q10() {
	m := map[string]int{"A": 1, "B": 2}
	fmt.Println("\nQ10:")
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// Q11. Create a map with int keys.
func q11() {
	m := map[int]string{1: "One", 2: "Two"}
	fmt.Println("\nQ11:", m)
}

// Q12. Create a map with bool keys.
func q12() {
	m := map[bool]string{true: "Yes", false: "No"}
	fmt.Println("\nQ12:", m)
}

// Q13. Map with struct values.
func q13() {
	type Employee struct {
		Name string
		Age  int
	}

	m := map[int]Employee{
		1: {"Ashutosh", 25},
	}

	fmt.Println("\nQ13:", m)
}

// Q14. Map with slice values.
func q14() {
	m := map[string][]string{
		"Backend": {"Go", "PostgreSQL"},
		"Frontend": {"React", "TypeScript"},
	}
	fmt.Println("\nQ14:", m)
}

// Q15. Map with slice value: append.
func q15() {
	m := make(map[string][]string)
	m["skills"] = append(m["skills"], "Go")
	m["skills"] = append(m["skills"], "PostgreSQL")
	fmt.Println("\nQ15:", m)
}

// Q16. Nil map read.
func q16() {
	var m map[string]int
	fmt.Println("\nQ16 read nil map:", m["Go"])
}

// Q17. Demonstrate nil map safely.
func q17() {
	var m map[string]int
	_, ok := m["Go"]
	fmt.Println("\nQ17 exists:", ok)
}

// Q18. Initialize a nil map before writing.
func q18() {
	var m map[string]int
	m = make(map[string]int)
	m["Go"] = 100
	fmt.Println("\nQ18:", m)
}

// Q19. Count character frequency.
func q19() {
	s := "banana"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	fmt.Println("\nQ19:", freq)
}

// Q20. Count word frequency.
func q20() {
	words := []string{"go", "java", "go", "python", "go"}
	freq := map[string]int{}
	for _, word := range words {
		freq[word]++
	}
	fmt.Println("\nQ20:", freq)
}

// Q21. Find duplicate numbers.
func q21() {
	nums := []int{1, 2, 3, 2, 4, 1}
	seen := map[int]bool{}
	duplicates := map[int]bool{}

	for _, n := range nums {
		if seen[n] {
			duplicates[n] = true
		}
		seen[n] = true
	}

	fmt.Println("\nQ21:", duplicates)
}

// Q22. Find first repeated number.
func q22() {
	nums := []int{4, 1, 3, 2, 3, 1}
	seen := map[int]bool{}
	result := -1

	for _, n := range nums {
		if seen[n] {
			result = n
			break
		}
		seen[n] = true
	}

	fmt.Println("\nQ22:", result)
}

// Q23. Find first non-repeated number.
func q23() {
	nums := []int{4, 5, 1, 5, 4, 2}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result := -1
	for _, n := range nums {
		if freq[n] == 1 {
			result = n
			break
		}
	}

	fmt.Println("\nQ23:", result)
}

// Q24. Find most frequent number.
func q24() {
	nums := []int{1, 2, 2, 3, 2, 4, 3}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result, max := 0, 0
	for n, count := range freq {
		if count > max {
			result, max = n, count
		}
	}

	fmt.Println("\nQ24 number:", result, "count:", max)
}

// Q25. Find least frequent number.
func q25() {
	nums := []int{1, 2, 2, 3, 3, 3}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result, min := 0, len(nums)+1
	for n, count := range freq {
		if count < min {
			result, min = n, count
		}
	}

	fmt.Println("\nQ25 number:", result, "count:", min)
}

// Q26. Find missing number from 1..n.
func q26() {
	nums := []int{1, 2, 4, 5}
	n := 5
	seen := map[int]bool{}

	for _, x := range nums {
		seen[x] = true
	}

	result := -1
	for i := 1; i <= n; i++ {
		if !seen[i] {
			result = i
			break
		}
	}

	fmt.Println("\nQ26:", result)
}

// Q27. Two Sum using a map.
func q27() {
	nums := []int{2, 7, 11, 15}
	target := 9

	seen := map[int]int{}
	result := []int{}

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			result = []int{j, i}
			break
		}
		seen[n] = i
	}

	fmt.Println("\nQ27 indices:", result)
}

// Q28. Check if two arrays have common elements.
func q28() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 3}

	set := map[int]bool{}
	for _, n := range a {
		set[n] = true
	}

	found := false
	for _, n := range b {
		if set[n] {
			found = true
			break
		}
	}

	fmt.Println("\nQ28:", found)
}

// Q29. Find intersection of two arrays.
func q29() {
	a := []int{1, 2, 2, 3}
	b := []int{2, 3, 4}

	set := map[int]bool{}
	for _, n := range a {
		set[n] = true
	}

	result := []int{}
	added := map[int]bool{}

	for _, n := range b {
		if set[n] && !added[n] {
			result = append(result, n)
			added[n] = true
		}
	}

	fmt.Println("\nQ29:", result)
}

// Q30. Find union of two arrays.
func q30() {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}

	set := map[int]bool{}
	for _, n := range a {
		set[n] = true
	}
	for _, n := range b {
		set[n] = true
	}

	result := []int{}
	for n := range set {
		result = append(result, n)
	}

	sort.Ints(result)
	fmt.Println("\nQ30:", result)
}

// Q31. Count positive and negative numbers.
func q31() {
	nums := []int{-2, 4, -1, 6, 0}
	count := map[string]int{}

	for _, n := range nums {
		if n > 0 {
			count["positive"]++
		} else if n < 0 {
			count["negative"]++
		} else {
			count["zero"]++
		}
	}

	fmt.Println("\nQ31:", count)
}

// Q32. Group numbers by parity.
func q32() {
	nums := []int{1, 2, 3, 4, 5, 6}
	groups := map[string][]int{}

	for _, n := range nums {
		if n%2 == 0 {
			groups["even"] = append(groups["even"], n)
		} else {
			groups["odd"] = append(groups["odd"], n)
		}
	}

	fmt.Println("\nQ32:", groups)
}

// Q33. Group words by first character.
func q33() {
	words := []string{"apple", "ant", "ball", "banana", "cat"}
	groups := map[rune][]string{}

	for _, word := range words {
		r := []rune(word)
		if len(r) > 0 {
			groups[r[0]] = append(groups[r[0]], word)
		}
	}

	fmt.Println("\nQ33:", groups)
}

// Q34. Group words by length.
func q34() {
	words := []string{"go", "java", "rust", "c", "python"}
	groups := map[int][]string{}

	for _, word := range words {
		groups[len(word)] = append(groups[len(word)], word)
	}

	fmt.Println("\nQ34:", groups)
}

// Q35. Create map of employee salaries.
func q35() {
	salaries := map[string]int{
		"Alice": 50000,
		"Bob":   60000,
	}

	fmt.Println("\nQ35:", salaries)
}

// Q36. Find highest salary.
func q36() {
	salaries := map[string]int{
		"Alice": 50000,
		"Bob":   70000,
		"John":  60000,
	}

	name := ""
	max := 0

	for employee, salary := range salaries {
		if salary > max {
			name, max = employee, salary
		}
	}

	fmt.Println("\nQ36:", name, max)
}

// Q37. Find employees earning above threshold.
func q37() {
	salaries := map[string]int{
		"Alice": 50000,
		"Bob":   70000,
		"John":  60000,
	}

	fmt.Println("\nQ37:")
	for name, salary := range salaries {
		if salary > 55000 {
			fmt.Println(name, salary)
		}
	}
}

// Q38. Sort map keys.
func q38() {
	m := map[string]int{"banana": 2, "apple": 5, "orange": 3}
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("\nQ38:")
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}

// Q39. Print map in deterministic order.
func q39() {
	m := map[int]string{3: "C", 1: "A", 2: "B"}
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	fmt.Print("\nQ39: ")
	for _, key := range keys {
		fmt.Printf("%d=%s ", key, m[key])
	}
	fmt.Println()
}

// Q40. Compare two maps for equality.
func q40() {
	a := map[string]int{"A": 1, "B": 2}
	b := map[string]int{"B": 2, "A": 1}

	equal := mapsEqual(a, b)
	fmt.Println("\nQ40:", equal)
}

func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {
		if other, ok := b[key]; !ok || other != value {
			return false
		}
	}

	return true
}

// Q41. Copy a map.
func q41() {
	original := map[string]int{"A": 1, "B": 2}
	copyMap := make(map[string]int, len(original))

	for key, value := range original {
		copyMap[key] = value
	}

	copyMap["A"] = 100

	fmt.Println("\nQ41 original:", original)
	fmt.Println("Q41 copy:", copyMap)
}

// Q42. Merge two maps.
func q42() {
	a := map[string]int{"A": 1, "B": 2}
	b := map[string]int{"C": 3, "D": 4}

	result := make(map[string]int)

	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}

	fmt.Println("\nQ42:", result)
}

// Q43. Merge maps where second map overwrites duplicate keys.
func q43() {
	a := map[string]int{"A": 1, "B": 2}
	b := map[string]int{"B": 100, "C": 3}

	for k, v := range b {
		a[k] = v
	}

	fmt.Println("\nQ43:", a)
}

// Q44. Find common key/value pairs.
func q44() {
	a := map[string]int{"A": 1, "B": 2, "C": 3}
	b := map[string]int{"B": 2, "C": 9, "D": 4}

	result := map[string]int{}

	for k, v := range a {
		if other, ok := b[k]; ok && other == v {
			result[k] = v
		}
	}

	fmt.Println("\nQ44:", result)
}

// Q45. Remove all keys whose value is below a threshold.
func q45() {
	m := map[string]int{"A": 10, "B": 30, "C": 20}

	for k, v := range m {
		if v < 20 {
			delete(m, k)
		}
	}

	fmt.Println("\nQ45:", m)
}

// Q46. Invert a map.
func q46() {
	m := map[string]int{"Alice": 1, "Bob": 2}
	inverted := map[int]string{}

	for k, v := range m {
		inverted[v] = k
	}

	fmt.Println("\nQ46:", inverted)
}

// Q47. Find duplicate values in a map.
func q47() {
	m := map[string]int{
		"A": 10,
		"B": 20,
		"C": 10,
		"D": 30,
		"E": 20,
	}

	seen := map[int]string{}
	duplicates := []int{}

	for key, value := range m {
		if _, ok := seen[value]; ok {
			duplicates = append(duplicates, value)
		} else {
			seen[value] = key
		}
	}

	fmt.Println("\nQ47 duplicate values:", duplicates)
}

// Q48. Map of slices: find total number of skills.
func q48() {
	m := map[string][]string{
		"Alice": {"Go", "SQL"},
		"Bob":   {"React", "Go", "Docker"},
	}

	total := 0
	for _, skills := range m {
		total += len(skills)
	}

	fmt.Println("\nQ48:", total)
}

// Q49. Find employees who have a specific skill.
func q49() {
	employees := map[string][]string{
		"Alice": {"Go", "SQL"},
		"Bob":   {"React", "Java"},
		"John":  {"Go", "Docker"},
	}

	target := "Go"
	result := []string{}

	for name, skills := range employees {
		for _, skill := range skills {
			if skill == target {
				result = append(result, name)
				break
			}
		}
	}

	sort.Strings(result)
	fmt.Println("\nQ49:", result)
}

// Q50. Nested map.
func q50() {
	data := map[string]map[string]int{
		"HR": {
			"Alice": 50000,
			"Bob":   60000,
		},
		"IT": {
			"John": 80000,
		},
	}

	fmt.Println("\nQ50:", data["IT"]["John"])
}

// Q51. Safely create a nested map.
func q51() {
	data := make(map[string]map[string]int)

	if data["IT"] == nil {
		data["IT"] = make(map[string]int)
	}

	data["IT"]["John"] = 80000

	fmt.Println("\nQ51:", data)
}

// Q52. Map key as struct.
func q52() {
	type Point struct {
		X int
		Y int
	}

	m := map[Point]string{
		{1, 2}: "A",
		{3, 4}: "B",
	}

	fmt.Println("\nQ52:", m[Point{1, 2}])
}

// Q53. Find pair with target sum using frequency map.
func q53() {
	nums := []int{2, 3, 2, 4}
	target := 6

	freq := map[int]int{}
	pairs := 0

	for _, n := range nums {
		need := target - n
		pairs += freq[need]
		freq[n]++
	}

	fmt.Println("\nQ53 pairs:", pairs)
}

// Q54. Find numbers appearing in both arrays with frequency.
func q54() {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 2, 2, 4, 5}

	freq := map[int]int{}
	for _, n := range a {
		freq[n]++
	}

	result := []int{}
	for _, n := range b {
		if freq[n] > 0 {
			result = append(result, n)
			freq[n]--
		}
	}

	fmt.Println("\nQ54:", result)
}

// Q55. Find majority element using map.
func q55() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result := -1
	for n, count := range freq {
		if count > len(nums)/2 {
			result = n
			break
		}
	}

	fmt.Println("\nQ55:", result)
}

// Q56. Find all numbers appearing exactly twice.
func q56() {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result := []int{}
	for n, count := range freq {
		if count == 2 {
			result = append(result, n)
		}
	}

	sort.Ints(result)
	fmt.Println("\nQ56:", result)
}

// Q57. Find elements that occur only once.
func q57() {
	nums := []int{1, 2, 2, 3, 4, 4}
	freq := map[int]int{}

	for _, n := range nums {
		freq[n]++
	}

	result := []int{}
	for _, n := range nums {
		if freq[n] == 1 {
			result = append(result, n)
		}
	}

	fmt.Println("\nQ57:", result)
}

// Q58. Find longest consecutive sequence using a map/set.
func q58() {
	nums := []int{100, 4, 200, 1, 3, 2}
	set := map[int]bool{}

	for _, n := range nums {
		set[n] = true
	}

	best := 0

	for n := range set {
		if !set[n-1] {
			length := 1

			for set[n+length] {
				length++
			}

			if length > best {
				best = length
			}
		}
	}

	fmt.Println("\nQ58:", best)
}

// Q59. Find pair whose difference equals target.
func q59() {
	nums := []int{5, 20, 3, 2, 50, 80}
	target := 78

	seen := map[int]bool{}
	found := false

	for _, n := range nums {
		if seen[n-target] || seen[n+target] {
			found = true
			break
		}
		seen[n] = true
	}

	fmt.Println("\nQ59:", found)
}

// Q60. Find subarray with sum zero using prefix sums.
func q60() {
	nums := []int{4, 2, -3, 1, 6}
	seen := map[int]bool{0: true}
	sum := 0
	found := false

	for _, n := range nums {
		sum += n
		if seen[sum] {
			found = true
			break
		}
		seen[sum] = true
	}

	fmt.Println("\nQ60 zero-sum subarray:", found)
}

// Q61. Count subarrays whose sum equals K.
func q61() {
	nums := []int{1, 1, 1}
	k := 2

	prefixFreq := map[int]int{0: 1}
	sum := 0
	count := 0

	for _, n := range nums {
		sum += n
		count += prefixFreq[sum-k]
		prefixFreq[sum]++
	}

	fmt.Println("\nQ61:", count)
}

// Q62. Find two strings with same character frequency.
func q62() {
	a := "abbc"
	b := "cbba"

	fmt.Println("\nQ62:", areAnagramsMap(a, b))
}

func areAnagramsMap(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	freq := map[rune]int{}

	for _, r := range a {
		freq[r]++
	}

	for _, r := range b {
		freq[r]--
		if freq[r] < 0 {
			return false
		}
	}

	return true
}

// Q63. Find common words between two lists.
func q63() {
	a := []string{"go", "java", "python"}
	b := []string{"rust", "go", "java"}

	set := map[string]bool{}
	for _, word := range a {
		set[word] = true
	}

	result := []string{}
	used := map[string]bool{}

	for _, word := range b {
		if set[word] && !used[word] {
			result = append(result, word)
			used[word] = true
		}
	}

	fmt.Println("\nQ63:", result)
}

// Q64. Find the top frequent numbers.
func q64() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}

	type Pair struct {
		Num   int
		Count int
	}

	pairs := []Pair{}
	for n, count := range freq {
		pairs = append(pairs, Pair{n, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	result := []int{}
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i].Num)
	}

	fmt.Println("\nQ64:", result)
}

// Q65. Demonstrate map reference-like behavior.
func q65() {
	a := map[string]int{"x": 1}
	b := a

	b["x"] = 100

	fmt.Println("\nQ65 a:", a)
	fmt.Println("Q65 b:", b)
}

// Q66. Demonstrate that map is not comparable with == except nil.
func q66() {
	a := map[string]int{"x": 1}
	b := map[string]int{"x": 1}

	// a == b is INVALID in Go.
	// Use a custom comparison function instead.
	fmt.Println("\nQ66 equal:", mapsEqual(a, b))
}

// Q67. Delete while ranging over a map.
func q67() {
	m := map[string]int{
		"A": 10,
		"B": 20,
		"C": 30,
	}

	for key, value := range m {
		if value < 25 {
			delete(m, key)
		}
	}

	fmt.Println("\nQ67:", m)
}

// Q68. Demonstrate map iteration order is not guaranteed.
func q68() {
	m := map[string]int{"A": 1, "B": 2, "C": 3}

	fmt.Println("\nQ68 map iteration:")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Run order may differ. Sort keys when deterministic output is required.")
}

// Q69. Explain and demonstrate concurrent-safe access concept.
func q69() {
	/*
		Normal Go maps are NOT safe for concurrent read/write access.

		This is unsafe:
		    go func() { m["A"] = 1 }()
		    go func() { fmt.Println(m["A"]) }()

		Use synchronization such as sync.RWMutex, or sync.Map
		when appropriate.

		This example only shows the normal map used in one goroutine.
	*/
	m := make(map[string]int)
	m["A"] = 1

	fmt.Println("\nQ69:", m)
}

// Q70. Practical interview problem: group anagrams.
func q70() {
	words := []string{
		"eat",
		"tea",
		"tan",
		"ate",
		"nat",
		"bat",
	}

	groups := map[string][]string{}

	for _, word := range words {
		r := []rune(word)

		for i := 0; i < len(r)-1; i++ {
			for j := 0; j < len(r)-i-1; j++ {
				if r[j] > r[j+1] {
					r[j], r[j+1] = r[j+1], r[j]
				}
			}
		}

		key := string(r)
		groups[key] = append(groups[key], word)
	}

	fmt.Println("\nQ70 groups:", groups)
}

/*
===============================================================
IMPORTANT GO MAP INTERVIEW NOTES
===============================================================

1. What is a map?

A map stores key/value pairs:

    map[keyType]valueType

Example:

    map[string]int

2. How do you create a map?

    m := make(map[string]int)

or:

    m := map[string]int{
        "Go": 100,
    }

3. Can you write to a nil map?

NO.

    var m map[string]int
    m["Go"] = 10 // panic

Initialize first:

    m = make(map[string]int)

4. Can you read from a nil map?

YES.

    var m map[string]int
    fmt.Println(m["Go"]) // 0

5. How do you check if a key exists?

    value, ok := m["Go"]

If the key exists:
    ok == true

6. How do you delete?

    delete(m, "Go")

Deleting a missing key is safe.

7. Is map iteration ordered?

NO.

Never depend on range order.

8. Can map keys be slices?

NO.

This is invalid:

    map[[]int]string

because slices are not comparable.

9. Can arrays be map keys?

YES, if their element type is comparable.

10. Can structs be map keys?

YES, if all fields are comparable.

11. Can maps be compared with ==?

Maps cannot be compared with == to each other.

You can only compare a map with nil:

    m == nil

For equality, compare length and every key/value.

12. Are maps reference types?

Maps are commonly described as reference-like types.

Assignment does not create an independent copy:

    a := map[string]int{"x": 1}
    b := a
    b["x"] = 100

a["x"] is also 100.

13. How do you copy a map?

Create a new map and copy every entry.

14. Can you append directly to a map?

Not to a map itself, but you can append to a slice stored as a map value:

    m["skills"] = append(m["skills"], "Go")

15. Are maps safe for concurrent use?

A normal Go map is not safe for concurrent read/write.

Use:
    sync.RWMutex
or:
    sync.Map

when appropriate.

16. Time complexity

Average:
    lookup   O(1)
    insert   O(1)
    delete   O(1)

Worst-case behavior can differ because maps are hash-table based.

17. Why are maps heavily used in coding interviews?

They provide fast lookup and are useful for:

- frequency counting
- duplicate detection
- two sum
- grouping
- sets
- prefix sums
- anagrams
- intersections
- caching

18. Map as a set

Go does not have a built-in set type.

Common pattern:

    set := map[int]bool{}

or:

    set := map[int]struct{}{}

Using struct{} avoids storing a meaningful value.

19. Map of slices

Very common pattern:

    groups := map[string][]string{}

Then:

    groups["Go"] = append(groups["Go"], "Gin")

20. Nested maps

Example:

    users := map[string]map[string]int{}

Before writing to a nested map, initialize the inner map.

===============================================================
MOST IMPORTANT QUESTIONS TO PRACTICE WITHOUT LOOKING
===============================================================

Q19  Character frequency
Q20  Word frequency
Q21  Duplicates
Q22  First repeated number
Q23  First non-repeated number
Q24  Most frequent number
Q26  Missing number
Q27  Two Sum
Q29  Intersection
Q30  Union
Q32  Group by category
Q34  Group by length
Q36  Maximum value
Q38  Sort map keys
Q40  Compare maps
Q41  Copy map
Q42  Merge maps
Q46  Invert map
Q49  Map of slices
Q50  Nested map
Q53  Pair count
Q55  Majority element
Q58  Longest consecutive sequence
Q60  Zero-sum subarray
Q61  Subarray sum K
Q64  Top K frequent elements
Q65  Map assignment/reference behavior
Q67  Delete during range
Q69  Concurrent map safety
Q70  Group anagrams

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
	fmt.Println("ALL 70 MAP QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}
