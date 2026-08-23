package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
GO STRINGS — 80 CODING QUESTIONS WITH ANSWERS

Run:
    go run strings_80.go

IMPORTANT:
- Go strings are immutable.
- len(string) returns bytes, not Unicode characters.
- string[index] returns a byte.
- range over a string gives Unicode code points (runes).
- []byte is useful for byte-level modification.
- []rune is useful for Unicode character-level modification.
*/

// 1. Declare and print a string.
func q01() {
	s := "Hello Go"
	fmt.Println("\nQ01:", s)
}

// 2. Find string length.
func q02() {
	s := "Hello"
	fmt.Println("\nQ02 length:", len(s))
}

// 3. Access first character.
func q03() {
	s := "Hello"
	fmt.Println("\nQ03 first byte:", s[0])
	fmt.Println("Q03 first character:", string(s[0]))
}

// 4. Access last character.
func q04() {
	s := "Hello"
	fmt.Println("\nQ04 last:", string(s[len(s)-1]))
}

// 5. Print characters using index.
func q05() {
	s := "Hello"
	fmt.Print("\nQ05: ")
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), " ")
	}
	fmt.Println()
}

// 6. Print Unicode characters using range.
func q06() {
	s := "Hello नमस्ते"
	fmt.Println("\nQ06:")
	for index, r := range s {
		fmt.Println(index, string(r))
	}
}

// 7. Convert to uppercase.
func q07() {
	fmt.Println("\nQ07:", strings.ToUpper("hello world"))
}

// 8. Convert to lowercase.
func q08() {
	fmt.Println("\nQ08:", strings.ToLower("HELLO WORLD"))
}

// 9. Check substring.
//always return true 
func q09() {
	fmt.Println("\nQ09:", strings.Contains("golang programming", "program"))
}

// 10. Check prefix.
func q10() {
	fmt.Println("\nQ10:", strings.HasPrefix("golang", "go"))
}

// 11. Check suffix.
func q11() {
	fmt.Println("\nQ11:", strings.HasSuffix("main.go", ".go"))
}

// 12. Compare strings.
func q12() {
	a, b := "hello", "hello"
	fmt.Println("\nQ12:", a == b)
}

// 13. Case-insensitive comparison.
func q13() {
	fmt.Println("\nQ13:", strings.EqualFold("GoLang", "golang"))
}

// 14. Concatenate strings.
func q14() {
	fmt.Println("\nQ14:", "Hello"+" "+"World")
}

// 15. Repeat a string.
func q15() {
	fmt.Println("\nQ15:", strings.Repeat("Go ", 3))
}

// 16. Trim spaces.
func q16() {
	fmt.Println("\nQ16:", strings.TrimSpace("   hello   "))
}

// 17. Trim selected characters.
func q17() {
	fmt.Println("\nQ17:", strings.Trim("###hello###", "#"))
}

// 18. Replace text.
func q18() {
	fmt.Println("\nQ18:", strings.ReplaceAll("I like Java", "Java", "Go"))
}

// 19. Split string.
func q19() {
	fmt.Println("\nQ19:", strings.Split("Go,Java,Python", ","))
}

// 20. Join strings.
func q20() {
	fmt.Println("\nQ20:", strings.Join([]string{"Go", "is", "fast"}, " "))
}

// 21. Count vowels.
func q21() {
	s := "hello world"
	count := 0
	for _, r := range strings.ToLower(s) {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	fmt.Println("\nQ21 vowels:", count)
}

// 22. Count consonants.
func q22() {
	s := "hello world"
	count := 0
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
			default:
				count++
			}
		}
	}
	fmt.Println("\nQ22 consonants:", count)
}

// 23. Count digits.
func q23() {
	s := "go12345"
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	fmt.Println("\nQ23 digits:", count)
}

// 24. Count spaces.
func q24() {
	s := "Go is very fast"
	count := 0
	for _, r := range s {
		if r == ' ' {
			count++
		}
	}
	fmt.Println("\nQ24 spaces:", count)
}

// 25. Count words.
func q25() {
	fmt.Println("\nQ25 words:", len(strings.Fields("Go is a powerful language")))
}

// 26. Reverse string using runes.
func q26() {
	s := "hello"
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println("\nQ26 reverse:", string(r))
}

// 27. Check palindrome.
func q27() {
	fmt.Println("\nQ27 palindrome:", isPalindrome("madam"))
}

// 28. Remove spaces.
func q28() {
	fmt.Println("\nQ28:", strings.ReplaceAll("Go is very fast", " ", ""))
}

// 29. Remove vowels.
func q29() {
	s := "hello world"
	result := strings.Map(func(r rune) rune {
		switch unicode.ToLower(r) {
		case 'a', 'e', 'i', 'o', 'u':
			return -1
		}
		return r
	}, s)
	fmt.Println("\nQ29:", result)
}

// 30. Keep only digits.
func q30() {
	s := "abc123xyz456"
	result := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, s)
	fmt.Println("\nQ30:", result)
}

// 31. Keep only letters.
func q31() {
	s := "Go123-Lang!"
	result := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)
	fmt.Println("\nQ31:", result)
}

// 32. Count a character.
func q32() {
	s := "programming"
	count := 0
	for _, r := range s {
		if r == 'g' {
			count++
		}
	}
	fmt.Println("\nQ32 g count:", count)
}

// 33. Find first occurrence.
func q33() {
	fmt.Println("\nQ33 index:", strings.IndexRune("programming", 'g'))
}

// 34. Find last occurrence.
func q34() {
	fmt.Println("\nQ34 index:", strings.LastIndex("programming", "g"))
}

// 35. Count substring occurrences.
func q35() {
	fmt.Println("\nQ35:", strings.Count("go go go", "go"))
}

// 36. Replace only first occurrence.
func q36() {
	fmt.Println("\nQ36:", strings.Replace("cat cat cat", "cat", "dog", 1))
}

// 37. Check empty string.
func q37() {
	s := ""
	fmt.Println("\nQ37 empty:", s == "")
}

// 38. Remove duplicate characters.
func q38() {
	s := "programming"
	seen := map[rune]bool{}
	var b strings.Builder
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			b.WriteRune(r)
		}
	}
	fmt.Println("\nQ38:", b.String())
}

// 39. Find duplicate characters.
func q39() {
	s := "programming"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	fmt.Print("\nQ39 duplicates: ")
	for r, count := range freq {
		if count > 1 {
			fmt.Print(string(r), " ")
		}
	}
	fmt.Println()
}

// 40. First non-repeating character.
func q40() {
	s := "swiss"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	result := ""
	for _, r := range s {
		if freq[r] == 1 {
			result = string(r)
			break
		}
	}
	fmt.Println("\nQ40:", result)
}

// 41. First repeated character.
func q41() {
	s := "programming"
	seen := map[rune]bool{}
	result := ""
	for _, r := range s {
		if seen[r] {
			result = string(r)
			break
		}
		seen[r] = true
	}
	fmt.Println("\nQ41:", result)
}

// 42. Frequency of every character.
func q42() {
	s := "banana"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	fmt.Println("\nQ42:", freq)
}

// 43. Most frequent character.
func q43() {
	s := "banana"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	var result rune
	max := 0
	for r, count := range freq {
		if count > max {
			max = count
			result = r
		}
	}
	fmt.Println("\nQ43:", string(result), max)
}

// 44. Least frequent character.
func q44() {
	s := "banana"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	var result rune
	minCount := len([]rune(s)) + 1
	for r, count := range freq {
		if count < minCount {
			minCount = count
			result = r
		}
	}
	fmt.Println("\nQ44:", string(result), minCount)
}

// 45. Check anagram.
func q45() {
	fmt.Println("\nQ45 anagram:", areAnagrams("listen", "silent"))
}

// 46. Check anagram ignoring spaces/case.
func q46() {
	a := strings.ReplaceAll(strings.ToLower("conversation"), " ", "")
	b := strings.ReplaceAll(strings.ToLower("voices rant on"), " ", "")
	fmt.Println("\nQ46 anagram:", areAnagrams(a, b))
}

// 47. Check string rotation.
func q47() {
	a, b := "abcd", "cdab"
	fmt.Println("\nQ47 rotation:", len(a) == len(b) && strings.Contains(a+a, b))
}

// 48. Count words with Fields.
func q48() {
	fmt.Println("\nQ48:", len(strings.Fields("  Go   is   fast  ")))
}

// 49. Reverse words in sentence.
func q49() {
	words := strings.Fields("Go is very fast")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	fmt.Println("\nQ49:", strings.Join(words, " "))
}

// 50. Reverse every word.
func q50() {
	words := strings.Fields("Go is fast")
	for i, word := range words {
		r := []rune(word)
		for a, b := 0, len(r)-1; a < b; a, b = a+1, b-1 {
			r[a], r[b] = r[b], r[a]
		}
		words[i] = string(r)
	}
	fmt.Println("\nQ50:", strings.Join(words, " "))
}

// 51. Capitalize every word.
func q51() {
	words := strings.Fields("go is powerful")
	for i, word := range words {
		r := []rune(word)
		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}
		words[i] = string(r)
	}
	fmt.Println("\nQ51:", strings.Join(words, " "))
}

// 52. Toggle case.
func q52() {
	s := "GoLang"
	result := strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, s)
	fmt.Println("\nQ52:", result)
}

// 53. Check only digits.
func q53() {
	s := "123456"
	valid := s != ""
	for _, r := range s {
		if !unicode.IsDigit(r) {
			valid = false
			break
		}
	}
	fmt.Println("\nQ53:", valid)
}

// 54. Check only letters.
func q54() {
	s := "Golang"
	valid := s != ""
	for _, r := range s {
		if !unicode.IsLetter(r) {
			valid = false
			break
		}
	}
	fmt.Println("\nQ54:", valid)
}

// 55. Palindrome ignoring spaces and case.
func q55() {
	s := "A man a plan a canal Panama"
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
	fmt.Println("\nQ55:", isPalindrome(clean))
}

// 56. Find longest word.
func q56() {
	words := strings.Fields("Go programming is powerful")
	longest := ""
	for _, word := range words {
		if len([]rune(word)) > len([]rune(longest)) {
			longest = word
		}
	}
	fmt.Println("\nQ56:", longest)
}

// 57. Find shortest word.
func q57() {
	words := strings.Fields("Go programming is powerful")
	shortest := words[0]
	for _, word := range words[1:] {
		if len([]rune(word)) < len([]rune(shortest)) {
			shortest = word
		}
	}
	fmt.Println("\nQ57:", shortest)
}

// 58. Longest substring without repeating characters.
func q58() {
	s := "abcabcbb"
	r := []rune(s)
	last := map[rune]int{}
	start, best := 0, 0
	for i, ch := range r {
		if p, ok := last[ch]; ok && p >= start {
			start = p + 1
		}
		last[ch] = i
		if i-start+1 > best {
			best = i - start + 1
		}
	}
	fmt.Println("\nQ58 length:", best)
}

// 59. Find all substring occurrences.
func q59() {
	s, target := "abababa", "aba"
	positions := []int{}
	for i := 0; i+len(target) <= len(s); i++ {
		if s[i:i+len(target)] == target {
			positions = append(positions, i)
		}
	}
	fmt.Println("\nQ59:", positions)
}

// 60. Valid palindrome ignoring punctuation.
func q60() {
	s := "No 'x' in Nixon"
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
	fmt.Println("\nQ60:", isPalindrome(clean))
}

// 61. Manual string reverse.
func q61() {
	s := "Golang"
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println("\nQ61:", string(r))
}

// 62. Manual palindrome.
func q62() {
	fmt.Println("\nQ62:", isPalindrome("level"))
}

// 63. First unique character.
func q63() {
	s := "leetcode"
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	result := ""
	for _, r := range s {
		if freq[r] == 1 {
			result = string(r)
			break
		}
	}
	fmt.Println("\nQ63:", result)
}

// 64. Group anagrams.
func q64() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := map[string][]string{}
	for _, word := range words {
		r := []rune(word)
		sortRunes(r)
		groups[string(r)] = append(groups[string(r)], word)
	}
	fmt.Println("\nQ64:", groups)
}

// 65. Longest common prefix.
func q65() {
	words := []string{"flower", "flow", "flight"}
	prefix := words[0]
	for _, word := range words[1:] {
		for !strings.HasPrefix(word, prefix) {
			if prefix == "" {
				break
			}
			r := []rune(prefix)
			prefix = string(r[:len(r)-1])
		}
	}
	fmt.Println("\nQ65:", prefix)
}

// 66. Check isomorphic strings.
func q66() {
	fmt.Println("\nQ66:", isIsomorphic("egg", "add"))
}

// 67. Check subsequence.
func q67() {
	small, large := []rune("ace"), []rune("abcde")
	i := 0
	for _, r := range large {
		if i < len(small) && small[i] == r {
			i++
		}
	}
	fmt.Println("\nQ67:", i == len(small))
}

// 68. Remove consecutive duplicate characters.
func q68() {
	s := "aaabbcddd"
	var b strings.Builder
	var previous rune
	first := true
	for _, r := range s {
		if first || r != previous {
			b.WriteRune(r)
		}
		previous = r
		first = false
	}
	fmt.Println("\nQ68:", b.String())
}

// 69. Run-length encoding.
func q69() {
	s := "aaabbc"
	r := []rune(s)
	var b strings.Builder
	count := 1
	for i := 1; i <= len(r); i++ {
		if i < len(r) && r[i] == r[i-1] {
			count++
		} else {
			b.WriteRune(r[i-1])
			b.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	fmt.Println("\nQ69:", b.String())
}

// 70. Run-length decoding, one digit count per character.
func q70() {
	s := "a3b2c1"
	r := []rune(s)
	var b strings.Builder
	for i := 0; i+1 < len(r); i += 2 {
		count := int(r[i+1] - '0')
		b.WriteString(strings.Repeat(string(r[i]), count))
	}
	fmt.Println("\nQ70:", b.String())
}

// 71. Count character difference between two strings.
func q71() {
	a, b := "abc", "cde"
	freq := map[rune]int{}
	for _, r := range a {
		freq[r]++
	}
	for _, r := range b {
		freq[r]--
	}
	total := 0
	for _, n := range freq {
		if n < 0 {
			n = -n
		}
		total += n
	}
	fmt.Println("\nQ71:", total)
}

// 72. Longest palindromic substring.
func q72() {
	s := []rune("babad")
	bestStart, bestLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > bestLen {
				bestStart = left
				bestLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := range s {
		expand(i, i)
		expand(i, i+1)
	}
	fmt.Println("\nQ72:", string(s[bestStart:bestStart+bestLen]))
}

// 73. Check balanced brackets.
func q73() {
	s := "{[()]}"
	stack := []rune{}
	valid := true

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				valid = false
				break
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !matches(last, r) {
				valid = false
				break
			}
		}
	}
	if len(stack) != 0 {
		valid = false
	}
	fmt.Println("\nQ73:", valid)
}

// 74. Longest word with no repeated character.
func q74() {
	words := strings.Fields("hello world abcdef")
	best := ""
	for _, word := range words {
		seen := map[rune]bool{}
		valid := true
		for _, r := range word {
			if seen[r] {
				valid = false
				break
			}
			seen[r] = true
		}
		if valid && len([]rune(word)) > len([]rune(best)) {
			best = word
		}
	}
	fmt.Println("\nQ74:", best)
}

// 75. Convert string to integer.
func q75() {
	n, err := strconv.Atoi("12345")
	fmt.Println("\nQ75:", n, err)
}

// 76. Convert integer to string.
func q76() {
	fmt.Println("\nQ76:", strconv.Itoa(12345))
}

// 77. Validate integer strings.
func q77() {
	for _, s := range []string{"123", "12a", "-50"} {
		_, err := strconv.Atoi(s)
		fmt.Println("\nQ77", s, "valid:", err == nil)
	}
}

// 78. Most common word ignoring punctuation and case.
func q78() {
	s := "Go, go, GO! Java."
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, s)
	words := strings.Fields(clean)
	freq := map[string]int{}
	for _, word := range words {
		freq[word]++
	}
	best, count := "", 0
	for _, word := range words {
		if freq[word] > count {
			best, count = word, freq[word]
		}
	}
	fmt.Println("\nQ78:", best, count)
}

// 79. Count Unicode characters correctly.
func q79() {
	s := "नमस्ते"
	fmt.Println("\nQ79 bytes:", len(s))
	fmt.Println("Q79 runes:", utf8.RuneCountInString(s))
}

// 80. Build a string efficiently with strings.Builder.
func q80() {
	words := []string{"Go", "is", "fast"}
	var b strings.Builder
	for i, word := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(word)
	}
	fmt.Println("\nQ80:", b.String())
}

func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func areAnagrams(a, b string) bool {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return false
	}
	freq := map[rune]int{}
	for _, r := range ar {
		freq[r]++
	}
	for _, r := range br {
		freq[r]--
		if freq[r] < 0 {
			return false
		}
	}
	return true
}

func sortRunes(r []rune) {
	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < len(r)-i-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}

func isIsomorphic(a, b string) bool {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return false
	}
	ab := map[rune]rune{}
	ba := map[rune]rune{}

	for i := range ar {
		x, y := ar[i], br[i]
		if v, ok := ab[x]; ok && v != y {
			return false
		}
		if v, ok := ba[y]; ok && v != x {
			return false
		}
		ab[x] = y
		ba[y] = x
	}
	return true
}

func matches(open, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '[' && close == ']') ||
		(open == '{' && close == '}')
}

func main() {
	q01(); q02(); q03(); q04(); q05(); q06(); q07(); q08(); q09(); q10()
	q11(); q12(); q13(); q14(); q15(); q16(); q17(); q18(); q19(); q20()
	q21(); q22(); q23(); q24(); q25(); q26(); q27(); q28(); q29(); q30()
	q31(); q32(); q33(); q34(); q35(); q36(); q37(); q38(); q39(); q40()
	q41(); q42(); q43(); q44(); q45(); q46(); q47(); q48(); q49(); q50()
	q51(); q52(); q53(); q54(); q55(); q56(); q57(); q58(); q59(); q60()
	q61(); q62(); q63(); q64(); q65(); q66(); q67(); q68(); q69(); q70()
	q71(); q72(); q73(); q74(); q75(); q76(); q77(); q78(); q79(); q80()

	fmt.Println("\n============================================================")
	fmt.Println("ALL 80 STRING QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}

/*
====================================================================
STRING INTERVIEW CHEAT SHEET
====================================================================

1. string is immutable.
2. len(s) counts bytes.
3. s[i] gives a byte.
4. range gives rune values.
5. []byte(s) gives mutable bytes.
6. []rune(s) gives mutable Unicode code points.
7. strings.Contains checks substring.
8. strings.Split returns []string.
9. strings.Join joins []string.
10. strings.Fields handles whitespace-separated words.
11. strings.TrimSpace removes surrounding whitespace.
12. strings.ReplaceAll replaces all occurrences.
13. strings.EqualFold compares case-insensitively.
14. strings.Builder is useful for repeated construction.
15. strconv.Atoi converts string to int.
16. strconv.Itoa converts int to string.
17. utf8.RuneCountInString counts Unicode code points.
18. For ASCII, byte indexing is straightforward.
19. For Unicode, prefer range or []rune.
20. Never try s[0] = 'A'; strings are immutable.

COMMON INTERVIEW QUESTIONS:
- Reverse a string.
- Check palindrome.
- Check anagram.
- Find first non-repeating character.
- Find first repeated character.
- Count character frequency.
- Remove duplicate characters.
- Longest substring without repeating characters.
- Longest common prefix.
- Group anagrams.
- Check string rotation.
- Check subsequence.
- Longest palindromic substring.
- Balanced brackets.
- Run-length encoding.
- String to integer.
- Unicode length.
*/
