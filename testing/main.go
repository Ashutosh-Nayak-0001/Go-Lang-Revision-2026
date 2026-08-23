// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	num := 12321
	temp := num
	reverse := 0
	
	

	for temp != 0 {
		rnum := temp % 10
		reverse = reverse*10 + rnum
		temp /= 10
	}

	if num == reverse {
		fmt.Println("num is palindrome")
	}

}
