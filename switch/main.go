package main

import "fmt"

func main() {

	// ============================================================
	// 1. DAY BASED ON NUMBER
	// ============================================================

	fmt.Println("\n========== Q1: DAY ==========")

	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}


	// ============================================================
	// 2. MONTH BASED ON NUMBER
	// ============================================================

	fmt.Println("\n========== Q2: MONTH ==========")

	month := 5

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid month")
	}


	// ============================================================
	// 3. CHECK VOWEL
	// ============================================================

	fmt.Println("\n========== Q3: VOWEL ==========")

	ch := 'a'

	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}


	// ============================================================
	// 4. CALCULATOR USING SWITCH
	// ============================================================

	fmt.Println("\n========== Q4: CALCULATOR ==========")

	a := 20
	b := 5
	operator := '*'

	switch operator {
	case '+':
		fmt.Println("Result:", a+b)

	case '-':
		fmt.Println("Result:", a-b)

	case '*':
		fmt.Println("Result:", a*b)

	case '/':
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Cannot divide by zero")
		}

	default:
		fmt.Println("Invalid operator")
	}


	// ============================================================
	// 5. TRAFFIC LIGHT
	// ============================================================

	fmt.Println("\n========== Q5: TRAFFIC LIGHT ==========")

	light := "red"

	switch light {
	case "red":
		fmt.Println("Stop")

	case "yellow":
		fmt.Println("Wait")

	case "green":
		fmt.Println("Go")

	default:
		fmt.Println("Invalid light")
	}


	// ============================================================
	// 6. WEEKDAY OR WEEKEND
	// ============================================================

	fmt.Println("\n========== Q6: WEEKDAY / WEEKEND ==========")

	dayName := "Sunday"

	switch dayName {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}


	// ============================================================
	// 7. GRADE USING SWITCH
	// ============================================================

	fmt.Println("\n========== Q7: GRADE ==========")

	grade := 'A'

	switch grade {
	case 'A':
		fmt.Println("Excellent")

	case 'B':
		fmt.Println("Very Good")

	case 'C':
		fmt.Println("Good")

	case 'D':
		fmt.Println("Pass")

	case 'F':
		fmt.Println("Fail")

	default:
		fmt.Println("Invalid grade")
	}


	// ============================================================
	// 8. HTTP STATUS CODE
	// ============================================================

	fmt.Println("\n========== Q8: HTTP STATUS ==========")

	status := 404

	switch status {
	case 200:
		fmt.Println("OK")

	case 201:
		fmt.Println("Created")

	case 400:
		fmt.Println("Bad Request")

	case 401:
		fmt.Println("Unauthorized")

	case 403:
		fmt.Println("Forbidden")

	case 404:
		fmt.Println("Not Found")

	case 500:
		fmt.Println("Internal Server Error")

	default:
		fmt.Println("Unknown status")
	}


	// ============================================================
	// 9. USER ROLE
	// ============================================================

	fmt.Println("\n========== Q9: USER ROLE ==========")

	role := "admin"

	switch role {
	case "admin":
		fmt.Println("Full access")

	case "manager":
		fmt.Println("Manager access")

	case "employee":
		fmt.Println("Employee access")

	case "guest":
		fmt.Println("Read-only access")

	default:
		fmt.Println("Unknown role")
	}


	// ============================================================
	// 10. HTTP METHOD
	// ============================================================

	fmt.Println("\n========== Q10: HTTP METHOD ==========")

	method := "POST"

	switch method {
	case "GET":
		fmt.Println("Fetch data")

	case "POST":
		fmt.Println("Create data")

	case "PUT":
		fmt.Println("Update data")

	case "DELETE":
		fmt.Println("Delete data")

	case "PATCH":
		fmt.Println("Partially update data")

	default:
		fmt.Println("Unsupported method")
	}


	// ============================================================
	// 11. SWITCH WITHOUT EXPRESSION
	// ============================================================

	fmt.Println("\n========== Q11: AGE CATEGORY ==========")

	age := 25

	switch {
	case age < 0:
		fmt.Println("Invalid age")

	case age < 13:
		fmt.Println("Child")

	case age < 18:
		fmt.Println("Teenager")

	case age < 60:
		fmt.Println("Adult")

	default:
		fmt.Println("Senior citizen")
	}


	// ============================================================
	// 12. POSITIVE / NEGATIVE / ZERO USING SWITCH
	// ============================================================

	fmt.Println("\n========== Q12: NUMBER TYPE ==========")

	num := -10

	switch {
	case num > 0:
		fmt.Println("Positive")

	case num < 0:
		fmt.Println("Negative")

	default:
		fmt.Println("Zero")
	}


	// ============================================================
	// 13. MARKS RANGE USING SWITCH
	// ============================================================

	fmt.Println("\n========== Q13: MARKS ==========")

	marks := 85

	switch {
	case marks < 0 || marks > 100:
		fmt.Println("Invalid marks")

	case marks >= 90:
		fmt.Println("A+")

	case marks >= 80:
		fmt.Println("A")

	case marks >= 70:
		fmt.Println("B")

	case marks >= 60:
		fmt.Println("C")

	case marks >= 40:
		fmt.Println("D")

	default:
		fmt.Println("Fail")
	}


	// ============================================================
	// 14. EVEN / ODD USING SWITCH
	// ============================================================

	fmt.Println("\n========== Q14: EVEN / ODD ==========")

	evenOddNumber := 17

	switch evenOddNumber % 2 {
	case 0:
		fmt.Println("Even")

	case 1:
		fmt.Println("Odd")
	}


	// ============================================================
	// 15. FALLTHROUGH
	// ============================================================

	fmt.Println("\n========== Q15: FALLTHROUGH ==========")

	fallNumber := 1

	switch fallNumber {
	case 1:
		fmt.Println("One")
		fallthrough

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
	}


	// ============================================================
	// 16. NESTED SWITCH
	// ============================================================

	fmt.Println("\n========== Q16: NESTED SWITCH ==========")

	country := "India"
	language := "English"

	switch country {

	case "India":

		switch language {
		case "English":
			fmt.Println("English - India")

		case "Hindi":
			fmt.Println("Hindi - India")

		default:
			fmt.Println("Other language")
		}

	case "USA":
		fmt.Println("USA")

	default:
		fmt.Println("Unknown country")
	}


	// ============================================================
	// 17. MENU-DRIVEN PROGRAM
	// ============================================================

	fmt.Println("\n========== Q17: MENU ==========")

	choice := 2

	switch choice {
	case 1:
		fmt.Println("Create User")

	case 2:
		fmt.Println("View User")

	case 3:
		fmt.Println("Update User")

	case 4:
		fmt.Println("Delete User")

	case 5:
		fmt.Println("Exit")

	default:
		fmt.Println("Invalid choice")
	}


	// ============================================================
	// 18. ROLE + ACCOUNT STATUS
	// ============================================================

	fmt.Println("\n========== Q18: ROLE + STATUS ==========")

	userRole := "admin"
	active := true

	switch userRole {

	case "admin":
		if active {
			fmt.Println("Admin has full access")
		} else {
			fmt.Println("Admin account disabled")
		}

	case "manager":
		if active {
			fmt.Println("Manager access")
		} else {
			fmt.Println("Manager account disabled")
		}

	case "employee":
		if active {
			fmt.Println("Employee access")
		} else {
			fmt.Println("Employee account disabled")
		}

	default:
		fmt.Println("Unknown role")
	}


	// ============================================================
	// 19. TYPE SWITCH
	// ============================================================

	fmt.Println("\n========== Q19: TYPE SWITCH ==========")

	var value interface{} = 100

	switch v := value.(type) {

	case int:
		fmt.Println("Integer:", v)

	case string:
		fmt.Println("String:", v)

	case float64:
		fmt.Println("Float:", v)

	case bool:
		fmt.Println("Boolean:", v)

	default:
		fmt.Println("Unknown type")
	}


	// ============================================================
	// 20. COMPLETE CALCULATOR WITH INPUT
	// ============================================================

	fmt.Println("\n========== Q20: COMPLETE CALCULATOR ==========")

	var firstNumber float64
	var secondNumber float64
	var operation string

	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter operator (+ - * /): ")
	fmt.Scan(&operation)

	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	switch operation {

	case "+":
		fmt.Printf("Result: %.2f\n", firstNumber+secondNumber)

	case "-":
		fmt.Printf("Result: %.2f\n", firstNumber-secondNumber)

	case "*":
		fmt.Printf("Result: %.2f\n", firstNumber*secondNumber)

	case "/":
		if secondNumber == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Printf("Result: %.2f\n", firstNumber/secondNumber)
		}

	default:
		fmt.Println("Invalid operator")
	}


	// ============================================================
	// END
	// ============================================================

	fmt.Println("\n========== ALL 20 QUESTIONS COMPLETED ==========")
}