package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "hello"
	// s1 := 'h'
	// fmt.Println(len(s))
	// fmt.Println(s1)

	// for i := len(s) - 1; i >= 0; i-- {
	// fmt.Println(string(s[i]))
	//  }

	//  for i,v := range s{
	// 	fmt.Println(i,string(v))
	//  }

	//  for _,ch := range s{
	// 	fmt.Printf("%c ", ch)
	// 	fmt.Print(ch)
	//  }

	// count := 0

	// for _, ch := range s {
	// 	fmt.Println("characters---", ch)
	// 	switch ch {
	// 	case 'a', 'e', 'i', 'o', 'u',
	// 		'A', 'E', 'I', 'O', 'U':
	// 		continue
	// 	default :
	// 	  count++
	// 	}
	// }

	// fmt.Println(count)

	// count := 0

	// for _, ch := range s {
	// 	if unicode.IsLetter(ch) {
	// 		switch unicode.ToLower(ch) {
	// 		case 'a', 'e', 'i', 'o', 'u':
	// 			continue
	// 		default:
	// 			count++
	// 		}
	// 	}
	// }

	//fmt.Println(unicode.IsLetter('a'))

	runes := []rune(s)

}
