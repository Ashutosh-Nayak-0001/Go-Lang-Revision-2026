// package main

// import "fmt"

// func main() {
// 	// var num int

// 	// fmt.Scan(&num)
// 	// fmt.Println(&num)
// 	// fmt.Println("You entered :", num)

// 	// var a, b int

// 	// fmt.Scan(&a, &b)

// 	// fmt.Println("First:", a)
// 	// fmt.Println("Second:", b)

// 	// var a, b int

// 	// fmt.Scan(&a, &b)

// 	// sum := a + b

// 	// fmt.Println("Sum:", sum)/

// 	// var name string
// 	// var age int

// 	// fmt.Scan(&name, &age)

// 	// fmt.Printf("My name is %s and I am %d years old.\n", name, age)

// // 	| Verb   | Meaning                       |
// // | ------ | ----------------------------- |
// // | `%v`   | Default value                 |
// // | `%+v`  | Value with struct field names |
// // | `%#v`  | Go-syntax representation      |
// // | `%T`   | Type                          |
// // | `%d`   | Decimal integer               |
// // | `%f`   | Floating point                |
// // | `%.2f` | Float with 2 decimal places   |
// // | `%s`   | String                        |
// // | `%c`   | Character                     |
// // | `%q`   | Quoted string/character       |
// // | `%t`   | Boolean                       |
// // | `%b`   | Binary                        |
// // | `%o`   | Octal                         |
// // | `%x`   | Hexadecimal                   |
// // | `%X`   | Uppercase hexadecimal         |
// // | `%U`   | Unicode format                |
// // | `%p`   | Pointer address               |
// // | `%%`   | Literal `%`                   |

// }

package main

import "fmt"

func main() {

	// ============================================================
	// 1. PRINT TWO LINES
	// ============================================================

	fmt.Println("========== 1. PRINT ==========")

	fmt.Println("Hello, Go!")
	fmt.Println("Welcome to coding practice.")

	fmt.Println()

	// ============================================================
	// 2. READ ONE INTEGER
	// ============================================================

	fmt.Println("========== 2. READ INTEGER ==========")

	var num int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	fmt.Println("You entered:", num)

	fmt.Println()

	// ============================================================
	// 3. READ TWO INTEGERS
	// ============================================================

	fmt.Println("========== 3. TWO INTEGERS ==========")

	var a, b int

	fmt.Print("Enter two integers: ")
	fmt.Scan(&a, &b)

	fmt.Println("First:", a)
	fmt.Println("Second:", b)

	fmt.Println()

	// ============================================================
	// 4. SUM OF TWO INTEGERS
	// ============================================================

	fmt.Println("========== 4. SUM ==========")

	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	sum := x + y

	fmt.Println("Sum:", sum)

	fmt.Println()

	// ============================================================
	// 5. NAME AND AGE
	// ============================================================

	fmt.Println("========== 5. NAME AND AGE ==========")

	var name string
	var age int

	fmt.Print("Enter name and age: ")
	fmt.Scan(&name, &age)

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	fmt.Println()

	// ============================================================
	// 6. INTEGER, FLOAT, STRING, BOOLEAN
	// ============================================================

	fmt.Println("========== 6. DIFFERENT DATA TYPES ==========")

	var userAge int
	var salary float64
	var userName string
	var active bool

	fmt.Print("Enter age, salary, name, active: ")
	fmt.Scan(&userAge, &salary, &userName, &active)

	fmt.Printf("Age: %d\n", userAge)
	fmt.Printf("Salary: %f\n", salary)
	fmt.Printf("Name: %s\n", userName)
	fmt.Printf("Active: %t\n", active)

	fmt.Println()

	// ============================================================
	// 7. FLOAT PRECISION
	// ============================================================

	fmt.Println("========== 7. FLOAT PRECISION ==========")

	var floatNum float64

	fmt.Print("Enter a floating-point number: ")
	fmt.Scan(&floatNum)

	fmt.Printf("Default : %f\n", floatNum)
	fmt.Printf("2 digits: %.2f\n", floatNum)
	fmt.Printf("3 digits: %.3f\n", floatNum)

	fmt.Println()

	// ============================================================
	// 8. DECIMAL, BINARY, OCTAL, HEX
	// ============================================================

	fmt.Println("========== 8. NUMBER FORMATS ==========")

	var number int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)

	fmt.Printf("Decimal     : %d\n", number)
	fmt.Printf("Binary      : %b\n", number)
	fmt.Printf("Octal       : %o\n", number)
	fmt.Printf("Hexadecimal : %x\n", number)

	fmt.Println()

	// ============================================================
	// 9. CHARACTER, VALUE AND TYPE
	// ============================================================

	fmt.Println("========== 9. CHARACTER ==========")

	var ch rune

	fmt.Print("Enter a character: ")
	fmt.Scan(&ch)

	fmt.Printf("Character : %c\n", ch)
	fmt.Printf("Value     : %d\n", ch)
	fmt.Printf("Type      : %T\n", ch)
	fmt.Printf("Unicode   : %U\n", ch)

	fmt.Println()

	// ============================================================
	// 10. %v, %T, %#v
	// ============================================================

	fmt.Println("========== 10. FORMAT VERBS ==========")

	testValue := "Ashutosh"

	fmt.Printf("Value            : %v\n", testValue)
	fmt.Printf("Type             : %T\n", testValue)
	fmt.Printf("Go Representation: %#v\n", testValue)

	fmt.Println()

	// ============================================================
	// 11. EMPLOYEE PROFILE
	// ============================================================

	fmt.Println("========== 11. EMPLOYEE PROFILE ==========")

	var employeeName string
	var employeeAge int
	var city string
	var employeeSalary float64

	fmt.Print("Enter name, age, city, salary: ")
	fmt.Scan(
		&employeeName,
		&employeeAge,
		&city,
		&employeeSalary,
	)

	fmt.Println("----- Employee Profile -----")
	fmt.Printf("Name   : %s\n", employeeName)
	fmt.Printf("Age    : %d\n", employeeAge)
	fmt.Printf("City   : %s\n", city)
	fmt.Printf("Salary : %.2f\n", employeeSalary)

	fmt.Println()

	// ============================================================
	// 12. READ THREE INTEGERS
	// ============================================================

	fmt.Println("========== 12. THREE INTEGERS ==========")

	var n1, n2, n3 int

	fmt.Print("Enter three integers: ")
	fmt.Scan(&n1, &n2, &n3)

	fmt.Println("First :", n1)
	fmt.Println("Second:", n2)
	fmt.Println("Third :", n3)

	fmt.Println()

	// ============================================================
	// 13. FIRST NAME AND LAST NAME
	// ============================================================

	fmt.Println("========== 13. FULL NAME ==========")

	var firstName string
	var lastName string

	fmt.Print("Enter first and last name: ")
	fmt.Scan(&firstName, &lastName)

	fmt.Printf("Full Name: %s %s\n", firstName, lastName)

	fmt.Println()

	// ============================================================
	// 14. PRODUCT CALCULATION
	// ============================================================

	fmt.Println("========== 14. PRODUCT ==========")

	var productName string
	var productPrice float64
	var quantity int

	fmt.Print("Enter product, price and quantity: ")
	fmt.Scan(&productName, &productPrice, &quantity)

	total := productPrice * float64(quantity)

	fmt.Printf("Product  : %s\n", productName)
	fmt.Printf("Price    : %.2f\n", productPrice)
	fmt.Printf("Quantity : %d\n", quantity)
	fmt.Printf("Total    : %.2f\n", total)

	fmt.Println()

	// ============================================================
	// 15. MULTIPLE fmt.Scan()
	// ============================================================

	fmt.Println("========== 15. MULTIPLE SCAN ==========")

	var scanName string
	var scanAge int

	fmt.Print("Enter name: ")
	fmt.Scan(&scanName)

	fmt.Print("Enter age: ")
	fmt.Scan(&scanAge)

	fmt.Println("Name:", scanName)
	fmt.Println("Age :", scanAge)

	fmt.Println()

	// ============================================================
	// 16. fmt.Scanf()
	// ============================================================

	fmt.Println("========== 16. SCANF ==========")

	var scanA, scanB int

	fmt.Print("Enter two integers: ")
	fmt.Scanf("%d %d", &scanA, &scanB)

	fmt.Println("Sum:", scanA+scanB)

	fmt.Println()

	// ============================================================
	// 17. SCANF WITH NAME AND AGE
	// ============================================================

	fmt.Println("========== 17. SCANF NAME AGE ==========")

	var scanfName string
	var scanfAge int

	fmt.Print("Enter name and age: ")
	fmt.Scanf("%s %d", &scanfName, &scanfAge)

	fmt.Printf("Name: %s\n", scanfName)
	fmt.Printf("Age : %d\n", scanfAge)

	fmt.Println()

	// ============================================================
	// 18. fmt.Sprint()
	// ============================================================

	fmt.Println("========== 18. SPRINT ==========")

	sprintName := "Ashutosh"
	sprintAge := 25

	result1 := fmt.Sprint(
		"My name is ",
		sprintName,
		" and I am ",
		sprintAge,
		" years old.",
	)

	fmt.Println(result1)

	fmt.Println()

	// ============================================================
	// 19. fmt.Sprintf()
	// ============================================================

	fmt.Println("========== 19. SPRINTF ==========")

	sprintfName := "Ashutosh"
	sprintfAge := 25

	result2 := fmt.Sprintf(
		"My name is %s and I am %d years old.",
		sprintfName,
		sprintfAge,
	)

	fmt.Println(result2)

	fmt.Println()

	// ============================================================
	// 20. STUDENT REPORT
	// ============================================================

	fmt.Println("========== 20. STUDENT REPORT ==========")

	var studentName string
	var studentAge int
	var rollNumber int

	var math int
	var science int
	var english int

	fmt.Print("Enter student name: ")
	fmt.Scan(&studentName)

	fmt.Print("Enter age: ")
	fmt.Scan(&studentAge)

	fmt.Print("Enter roll number: ")
	fmt.Scan(&rollNumber)

	fmt.Print("Enter Math marks: ")
	fmt.Scan(&math)

	fmt.Print("Enter Science marks: ")
	fmt.Scan(&science)

	fmt.Print("Enter English marks: ")
	fmt.Scan(&english)

	studentTotal := math + science + english
	percentage := float64(studentTotal) / 3

	fmt.Println()
	fmt.Println("========== STUDENT REPORT ==========")

	fmt.Printf("Name        : %s\n", studentName)
	fmt.Printf("Age         : %d\n", studentAge)
	fmt.Printf("Roll Number : %d\n", rollNumber)

	fmt.Println()

	fmt.Printf("Math        : %d\n", math)
	fmt.Printf("Science     : %d\n", science)
	fmt.Printf("English     : %d\n", english)

	fmt.Println()

	fmt.Printf("Total       : %d\n", studentTotal)
	fmt.Printf("Percentage  : %.2f%%\n", percentage)

	fmt.Println()
	fmt.Println("====================================")
}
