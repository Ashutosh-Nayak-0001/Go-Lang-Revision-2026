package main

import "fmt"

/*
GO POINTERS — 50 QUESTIONS WITH ANSWERS
=======================================

Run:
    go run pointers_50.go

CORE RULES:

&x  -> address of x
*p  -> value stored at the address held by p

Example:

x := 10
p := &x

fmt.Println(p)  // address
fmt.Println(*p) // 10

*p = 20

Now x is 20.

IMPORTANT:
- A pointer stores a memory address.
- * is used for dereferencing a pointer.
- & gets the address of a variable.
- A pointer can be nil.
- Dereferencing a nil pointer causes panic.
*/

// ============================================================
// Q01. Declare a pointer.
// ============================================================
func q01() {
	x := 10
	var p *int

	p = &x

	fmt.Println("\nQ01 pointer:", p)
}

// ============================================================
// Q02. Get the address of a variable using &.
// ============================================================
func q02() {
	x := 100

	fmt.Println("\nQ02 address:", &x)
}

// ============================================================
// Q03. Dereference a pointer using *.
// ============================================================
func q03() {
	x := 100
	p := &x

	fmt.Println("\nQ03 value:", *p)
}

// ============================================================
// Q04. Change a variable through a pointer.
// ============================================================
func q04() {
	x := 10
	p := &x

	*p = 50

	fmt.Println("\nQ04 x:", x)
}

// ============================================================
// Q05. Print value, address and pointer.
// ============================================================
func q05() {
	x := 25
	p := &x

	fmt.Println("\nQ05 value:", x)
	fmt.Println("Q05 address:", &x)
	fmt.Println("Q05 pointer:", p)
	fmt.Println("Q05 dereferenced:", *p)
}

// ============================================================
// Q06. Pointer to an integer.
// ============================================================
func q06() {
	x := 42
	p := &x

	fmt.Println("\nQ06:", *p)
}

// ============================================================
// Q07. Pointer to a string.
// ============================================================
func q07() {
	name := "Ashutosh"
	p := &name

	fmt.Println("\nQ07:", *p)
}

// ============================================================
// Q08. Pointer to a boolean.
// ============================================================
func q08() {
	flag := true
	p := &flag

	fmt.Println("\nQ08:", *p)
}

// ============================================================
// Q09. Check whether pointer is nil.
// ============================================================
func q09() {
	var p *int

	fmt.Println("\nQ09 is nil:", p == nil)
}

// ============================================================
// Q10. Initialize a nil pointer.
// ============================================================
func q10() {
	var p *int
     
	x := 100
	p = &x

	fmt.Println("\nQ10:", *p)
}

// ============================================================
// Q11. Pointer modification.
// ============================================================
func q11() {
	x := 10
	p := &x

	*p = *p + 20

	fmt.Println("\nQ11:", x)
}

// ============================================================
// Q12. Increment using pointer.
// ============================================================
func q12() {
	x := 10
	p := &x

	*p++

	fmt.Println("\nQ12:", x)
}

// ============================================================
// Q13. Decrement using pointer.
// ============================================================
func q13() {
	x := 10
	p := &x

	*p--

	fmt.Println("\nQ13:", x)
}

// ============================================================
// Q14. Swap two numbers using pointers.
// ============================================================
func q14() {
	a := 10
	b := 20

	swap(&a, &b)

	fmt.Println("\nQ14 a:", a)
	fmt.Println("Q14 b:", b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// ============================================================
// Q15. Function modifies value using pointer.
// ============================================================
func q15() {
	x := 10

	changeValue(&x)

	fmt.Println("\nQ15:", x)
}

func changeValue(p *int) {
	*p = 100
}

// ============================================================
// Q16. Pass pointer to function.
// ============================================================
func q16() {
	x := 20

	double(&x)

	fmt.Println("\nQ16:", x)
}

func double(p *int) {
	*p = *p * 2
}

// ============================================================
// Q17. Return a pointer from a function.
// ============================================================
func q17() {
	p := createNumber()

	fmt.Println("\nQ17:", *p)
}

func createNumber() *int {
	x := 100
	return &x
}

// ============================================================
// Q18. Pointer to pointer.
// ============================================================
func q18() {
	x := 10

	p := &x
	pp := &p

	fmt.Println("\nQ18 x:", x)
	fmt.Println("Q18 *p:", *p)
	fmt.Println("Q18 **pp:", **pp)
}

// ============================================================
// Q19. Modify value through pointer to pointer.
// ============================================================
func q19() {
	x := 10

	p := &x
	pp := &p

	**pp = 500

	fmt.Println("\nQ19:", x)
}

// ============================================================
// Q20. Pointer to struct.
// ============================================================
func q20() {
	type Employee struct {
		Name string
		Age  int
	}

	e := Employee{
		Name: "Ashutosh",
		Age:  25,
	}

	p := &e

	fmt.Println("\nQ20 name:", p.Name)
	fmt.Println("Q20 age:", p.Age)
}

// ============================================================
// Q21. Modify struct using pointer.
// ============================================================
func q21() {
	type Employee struct {
		Name string
		Age  int
	}

	e := Employee{
		Name: "Ashutosh",
		Age: 25,
	}

	p := &e
	p.Age = 26

	fmt.Println("\nQ21:", e)
}

// ============================================================
// Q22. Pointer receiver.
// ============================================================
func q22() {
	p := Person{Name: "Ashutosh"}

	p.ChangeName("Rahul")

	fmt.Println("\nQ22:", p.Name)
}

type Person struct {
	Name string
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}

// ============================================================
// Q23. Value receiver vs pointer receiver.
// ============================================================
func q23() {
	p := Person2{Name: "Ashutosh"}

	p.SetNameValue("Rahul")

	fmt.Println("\nQ23 after value receiver:", p.Name)

	p.SetNamePointer("Rahul")

	fmt.Println("Q23 after pointer receiver:", p.Name)
}

type Person2 struct {
	Name string
}

func (p Person2) SetNameValue(name string) {
	p.Name = name
}

func (p *Person2) SetNamePointer(name string) {
	p.Name = name
}

// ============================================================
// Q24. Pointer to array.
// ============================================================
func q24() {
	arr := [3]int{10, 20, 30}

	p := &arr

	(*p)[0] = 100

	fmt.Println("\nQ24:", arr)
}

// ============================================================
// Q25. Pointer to slice.
// ============================================================
func q25() {
	s := []int{1, 2, 3}

	p := &s

	*p = append(*p, 4)

	fmt.Println("\nQ25:", s)
}

// ============================================================
// Q26. Modify slice elements using pointer.
// ============================================================
func q26() {
	s := []int{10, 20, 30}

	p := &s

	(*p)[0] = 100

	fmt.Println("\nQ26:", s)
}

// ============================================================
// Q27. Pointer to map.
// ============================================================
func q27() {
	m := map[string]int{
		"Go": 100,
	}

	p := &m

	(*p)["Go"] = 200

	fmt.Println("\nQ27:", m)
}

// ============================================================
// Q28. Pointer to function.
// ============================================================
func q28() {
	add := func(a, b int) int {
		return a + b
	}

	var operation func(int, int) int
	operation = add

	fmt.Println("\nQ28:", operation(10, 20))
}

// ============================================================
// Q29. Pointer to interface variable.
// ============================================================
func q29() {
	var value interface{} = 100

	p := &value

	*p = "Go"

	fmt.Println("\nQ29:", value)
}

// ============================================================
// Q30. Check nil before dereferencing.
// ============================================================
func q30() {
	var p *int

	if p == nil {
		fmt.Println("\nQ30: pointer is nil")
		return
	}

	fmt.Println(*p)
}

// ============================================================
// Q31. Demonstrate nil pointer panic safely.
// ============================================================
func q31() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nQ31 recovered:", r)
		}
	}()

	var p *int

	fmt.Println(*p)
}

// ============================================================
// Q32. Pointer comparison.
// ============================================================
func q32() {
	a := 10
	b := 20

	p1 := &a
	p2 := &a
	p3 := &b

	fmt.Println("\nQ32 p1 == p2:", p1 == p2)
	fmt.Println("Q32 p1 == p3:", p1 == p3)
}

// ============================================================
// Q33. Pointer to same variable.
// ============================================================
func q33() {
	x := 10

	p1 := &x
	p2 := &x

	*p1 = 100

	fmt.Println("\nQ33 x:", x)
	fmt.Println("Q33 *p2:", *p2)
}

// ============================================================
// Q34. Multiple pointers to same variable.
// ============================================================
func q34() {
	x := 10

	p1 := &x
	p2 := &x
	p3 := &x

	*p2 = 500

	fmt.Println("\nQ34:", x, *p1, *p2, *p3)
}

// ============================================================
// Q35. Pointer arithmetic.
// ============================================================
func q35() {
	/*
		Go does NOT support normal pointer arithmetic.

		This is invalid:

		    p++

		For array/slice traversal, use indexes or range.
	*/

	arr := []int{10, 20, 30}

	fmt.Println("\nQ35:", arr[1])
}

// ============================================================
// Q36. Pointer and range variable.
// ============================================================
func q36() {
	values := []int{10, 20, 30}

	/*
		For modern Go versions, taking &v in a range loop
		creates an iteration variable for each iteration.

		For interview discussions, the safest general pattern
		when you need the address of an actual slice element is:

		    &values[i]
	*/

	pointers := []*int{}

	for i := range values {
		pointers = append(pointers, &values[i])
	}

	*pointers[0] = 100

	fmt.Println("\nQ36:", values)
}

// ============================================================
// Q37. Modify array using pointer.
// ============================================================
func q37() {
	arr := [3]int{1, 2, 3}

	modifyArray(&arr)

	fmt.Println("\nQ37:", arr)
}

func modifyArray(arr *[3]int) {
	arr[0] = 100
	arr[1] = 200
}

// ============================================================
// Q38. Return pointer to struct.
// ============================================================
func q38() {
	e := newEmployee()

	fmt.Println("\nQ38:", e.Name, e.Salary)
}

type Employee struct {
	Name   string
	Salary int
}

func newEmployee() *Employee {
	return &Employee{
		Name:   "Ashutosh",
		Salary: 70000,
	}
}

// ============================================================
// Q39. Use new() to allocate a pointer.
// ============================================================
func q39() {
	p := new(int)

	*p = 100

	fmt.Println("\nQ39:", *p)
}

// ============================================================
// Q40. Difference between new and make.
// ============================================================
func q40() {
	/*
		new(T):
		    allocates zero value of T
		    returns *T

		make(T):
		    initializes slices, maps and channels
		    returns T, not *T

		Example:
		    p := new(int)
		    m := make(map[string]int)
	*/

	p := new(int)
	m := make(map[string]int)

	*p = 10
	m["Go"] = 100

	fmt.Println("\nQ40 pointer:", *p)
	fmt.Println("Q40 map:", m)
}

// ============================================================
// Q41. Pointer parameter with multiple values.
// ============================================================
func q41() {
	a, b := 10, 20

	addTen(&a)
	addTen(&b)

	fmt.Println("\nQ41:", a, b)
}

func addTen(p *int) {
	*p += 10
}

// ============================================================
// Q42. Return nil pointer.
// ============================================================
func q42() {
	p := findEmployee(false)

	if p == nil {
		fmt.Println("\nQ42: employee not found")
		return
	}

	fmt.Println(p.Name)
}

func findEmployee(found bool) *Employee {
	if !found {
		return nil
	}

	return &Employee{Name: "Ashutosh"}
}

// ============================================================
// Q43. Pointer in a struct.
// ============================================================
func q43() {
	type EmployeeInfo struct {
		Name   string
		Salary *int
	}

	salary := 80000

	e := EmployeeInfo{
		Name:   "Ashutosh",
		Salary: &salary,
	}

	fmt.Println("\nQ43:", e.Name, *e.Salary)
}

// ============================================================
// Q44. Optional value using pointer.
// ============================================================
func q44() {
	type User struct {
		Name string
		Age  *int
	}

	u1 := User{Name: "A"}
	u2Age := 25
	u2 := User{Name: "B", Age: &u2Age}

	if u1.Age == nil {
		fmt.Println("\nQ44 User A: age not provided")
	}

	if u2.Age != nil {
		fmt.Println("Q44 User B age:", *u2.Age)
	}
}

// ============================================================
// Q45. Pointer and interface method set.
// ============================================================
func q45() {
	type Speaker interface {
		Speak()
	}

	person := Person3{Name: "Ashutosh"}

	/*
		Speak has a pointer receiver, so *Person3
		implements Speaker, while Person3 does not.
	*/

	var s Speaker = &person

	fmt.Println("\nQ45:")
	s.Speak()
}

type Person3 struct {
	Name string
}

func (p *Person3) Speak() {
	fmt.Println("Hello, I am", p.Name)
}

// ============================================================
// Q46. Modify nested struct using pointer.
// ============================================================
func q46() {
	type Address struct {
		City string
	}

	type User struct {
		Name    string
		Address Address
	}

	u := User{
		Name: "Ashutosh",
		Address: Address{
			City: "Bhubaneswar",
		},
	}

	p := &u

	p.Address.City = "Hyderabad"

	fmt.Println("\nQ46:", u)
}

// ============================================================
// Q47. Pointer to pointer function.
// ============================================================
func q47() {
	x := 10
	p := &x

	changePointer(&p)

	fmt.Println("\nQ47:", *p)
}

func changePointer(p **int) {
	x := 500
	*p = &x
}

// ============================================================
// Q48. Pointer with linked-list style node.
// ============================================================
func q48() {
	type Node struct {
		Value int
		Next  *Node
	}

	n1 := &Node{Value: 10}
	n2 := &Node{Value: 20}

	n1.Next = n2

	fmt.Println("\nQ48 first:", n1.Value)
	fmt.Println("Q48 second:", n1.Next.Value)
}

// ============================================================
// Q49. Deep pointer chain.
// ============================================================
func q49() {
	x := 10

	p1 := &x
	p2 := &p1
	p3 := &p2

	***p3 = 999

	fmt.Println("\nQ49:", x)
}

// ============================================================
// Q50. Practical interview problem:
//     modify multiple values through pointers.
// ============================================================
func q50() {
	type Account struct {
		Balance int
		Active  bool
	}

	account := Account{
		Balance: 1000,
		Active:  true,
	}

	updateAccount(&account, 500)

	fmt.Println("\nQ50:", account)
}

func updateAccount(account *Account, amount int) {
	account.Balance += amount
	account.Active = true
}

/*
=====================================================================
DEEP POINTER INTERVIEW NOTES
=====================================================================

1. WHAT IS A POINTER?

A pointer is a variable that stores the memory address of another
variable.

Example:

    x := 10
    p := &x

p stores the address of x.

    *p

gives the value stored at that address.

Therefore:

    fmt.Println(p)   // address
    fmt.Println(*p)  // value

---------------------------------------------------------------------

2. WHAT DOES & MEAN?

& means "address of".

    x := 10
    p := &x

p now points to x.

---------------------------------------------------------------------

3. WHAT DOES * MEAN?

There are two important uses.

A) Declaration:

    var p *int

This means p is a pointer to int.

B) Dereferencing:

    *p

This means "value stored at the address p points to".

---------------------------------------------------------------------

4. BASIC POINTER FLOW

    x := 10
       |
       | address
       v
    p := &x

p ---------> x
             10

Then:

    *p = 50

becomes:

p ---------> x
             50

---------------------------------------------------------------------

5. WHY DO WE USE POINTERS?

Main reasons:

- Modify a variable inside a function.
- Avoid copying large structs.
- Represent optional values using nil.
- Build linked lists, trees and other data structures.
- Work with mutable data.
- Use pointer receivers on methods.

---------------------------------------------------------------------

6. GO IS PASS-BY-VALUE

This is a very important interview question.

Go passes function arguments by value.

Example:

    func change(x int) {
        x = 100
    }

    a := 10
    change(a)

a is still 10.

But if you pass a pointer:

    func change(x *int) {
        *x = 100
    }

    change(&a)

a becomes 100.

You are still passing a value:
the value being passed is the pointer/address.

---------------------------------------------------------------------

7. POINTER TO POINTER

Example:

    x := 10
    p := &x
    pp := &p

Relationship:

    pp -> p -> x

Therefore:

    *pp  = p
    **pp = x

---------------------------------------------------------------------

8. NIL POINTER

A pointer can have nil as its zero value:

    var p *int

p == nil

But this is dangerous:

    fmt.Println(*p)

because it dereferences nil and causes a panic.

Always check when nil is possible:

    if p != nil {
        fmt.Println(*p)
    }

---------------------------------------------------------------------

9. new()

new(T) allocates zero value storage for T
and returns *T.

Example:

    p := new(int)

    *p = 100

---------------------------------------------------------------------

10. new() vs make()

new:

    p := new(int)

returns:

    *int

make:

    m := make(map[string]int)

returns:

    map[string]int

make is used for:

- maps
- slices
- channels

---------------------------------------------------------------------

11. POINTER TO STRUCT

You can use:

    p.Name

instead of:

    (*p).Name

Go automatically dereferences the pointer for struct field access.

Example:

    employee := Employee{Name: "A"}

    p := &employee

    fmt.Println(p.Name)

---------------------------------------------------------------------

12. POINTER RECEIVER

Example:

    func (p *Person) ChangeName(name string) {
        p.Name = name
    }

Use a pointer receiver when the method needs to modify
the receiver or when copying the receiver is undesirable.

---------------------------------------------------------------------

13. VALUE RECEIVER

Example:

    func (p Person) ChangeName(name string) {
        p.Name = name
    }

This changes only a copy.

Original value remains unchanged.

---------------------------------------------------------------------

14. POINTER RECEIVER AND INTERFACES

Very important interview topic.

If:

    func (p *Person) Speak() {}

then:

    *Person

implements the interface.

Person itself may not implement it.

Example:

    var s Speaker = &person

---------------------------------------------------------------------

15. POINTER ARITHMETIC

Go does NOT support C-style pointer arithmetic.

This is not normal Go:

    p++

For slice/array traversal use:

    for i := range arr {
        ...
    }

or:

    for _, value := range arr {
        ...
    }

---------------------------------------------------------------------

16. POINTER TO SLICE

Possible:

    s := []int{1, 2, 3}
    p := &s

Then:

    *p = append(*p, 4)

But most of the time you don't need a pointer to a slice,
because slices already contain a pointer to an underlying
array plus length and capacity.

---------------------------------------------------------------------

17. POINTER TO MAP

Possible:

    m := map[string]int{}
    p := &m

But maps are already reference-like descriptors, so a pointer
to a map is usually unnecessary.

---------------------------------------------------------------------

18. POINTER TO INTERFACE

Possible:

    var x interface{} = 10
    p := &x

But pointer-to-interface is uncommon and should not be used
unless there is a specific reason.

---------------------------------------------------------------------

19. RETURNING A POINTER FROM A FUNCTION

This is safe in Go:

    func create() *int {
        x := 10
        return &x
    }

Go's compiler/runtime manages memory lifetime using escape
analysis and garbage collection.

You do NOT manually free this memory.

---------------------------------------------------------------------

20. DOES GO HAVE MANUAL free()?

No.

Go uses garbage collection.

You normally do not manually free memory.

---------------------------------------------------------------------

21. MAPS, SLICES AND POINTERS

Interviewers may ask whether maps and slices are "reference types".

A better explanation:

Go has no formal "reference type" category like some languages.

Maps, slices, channels, pointers, interfaces and functions have
reference-like or descriptor behavior in different ways.

Do not simply say "everything is passed by reference".

Go function arguments are passed by value.

---------------------------------------------------------------------

22. POINTER COMPARISON

Pointers can be compared:

    p1 == p2

This checks whether they point to the same variable/address.

Example:

    x := 10

    p1 := &x
    p2 := &x

    p1 == p2 // true

---------------------------------------------------------------------

23. POINTER ZERO VALUE

The zero value of a pointer is:

    nil

Example:

    var p *int

p == nil

---------------------------------------------------------------------

24. POINTER TYPE

Examples:

    *int
    *string
    *bool
    *Employee
    *[]int
    *map[string]int

---------------------------------------------------------------------

25. MOST IMPORTANT INTERVIEW QUESTIONS

Practice these without looking:

1. What is a pointer?
2. What does & mean?
3. What does * mean?
4. Difference between *int and int.
5. What is nil pointer?
6. What happens when nil pointer is dereferenced?
7. How do you modify a variable using pointer?
8. Why pass pointers to functions?
9. Is Go pass-by-reference?
10. Explain pass-by-value with pointers.
11. What is pointer receiver?
12. Pointer receiver vs value receiver.
13. What is pointer to pointer?
14. What is new()?
15. new() vs make().
16. Can a pointer be nil?
17. Does Go support pointer arithmetic?
18. Can a function return a pointer?
19. Can a struct contain a pointer?
20. Can a pointer point to a struct?
21. How do interfaces interact with pointer receivers?
22. Why are pointers useful for large structs?
23. How does garbage collection affect pointers?
24. Can maps/slices use pointers?
25. What is escape analysis?

---------------------------------------------------------------------
INTERVIEW FORMULA TO REMEMBER
---------------------------------------------------------------------

    &x
      ↓
    address of x

    p := &x
      ↓
    p stores address of x

    *p
      ↓
    value at that address

    *p = 100
      ↓
    modify x

So:

    x := 10
    p := &x
    *p = 100

Result:

    x == 100

=====================================================================
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

	fmt.Println("\n============================================================")
	fmt.Println("ALL 50 POINTER QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}
