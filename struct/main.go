package main

import (
	"encoding/json"
	"fmt"
)

/*
GO STRUCTS — 60 QUESTIONS WITH ANSWERS
======================================

Run:
    go run structs_60.go

LEVELS:
    Q01-Q20  Beginner
    Q21-Q40  Intermediate
    Q41-Q60  Interview / Advanced

IMPORTANT:
A struct is a collection of fields grouped together into one type.

Example:

    type Employee struct {
        Name   string
        Age    int
        Salary float64
    }

Then:

    e := Employee{
        Name: "Ashutosh",
        Age:  25,
        Salary: 70000,
    }

*/

// ============================================================
// Q01. Create a basic struct.
// ============================================================
func q01() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{
		Name: "Ashutosh",
		Age:  25,
	}

	fmt.Println("\nQ01:", p)
}

// ============================================================
// Q02. Access struct fields.
// ============================================================
func q02() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Ashutosh", Age: 25}

	fmt.Println("\nQ02 name:", p.Name)
	fmt.Println("Q02 age:", p.Age)
}

// ============================================================
// Q03. Modify struct fields.
// ============================================================
func q03() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Ashutosh", Age: 25}

	p.Age = 26

	fmt.Println("\nQ03:", p)
}

// ============================================================
// Q04. Create struct using positional values.
// ============================================================
func q04() {
	type Point struct {
		X int
		Y int
	}

	p := Point{10, 20}

	fmt.Println("\nQ04:", p)
}

// ============================================================
// Q05. Create struct using named fields.
// ============================================================
func q05() {
	type Point struct {
		X int
		Y int
	}

	p := Point{
		Y: 20,
		X: 10,
	}

	fmt.Println("\nQ05:", p)
}

// ============================================================
// Q06. Partial struct initialization.
// ============================================================
func q06() {
	type Employee struct {
		Name   string
		Age    int
		Salary float64
		Active bool
	}

	e := Employee{
		Name: "Ashutosh",
	}

	fmt.Println("\nQ06:", e)
}

// ============================================================
// Q07. Zero value of a struct.
// ============================================================
func q07() {
	type Employee struct {
		Name   string
		Age    int
		Active bool
	}

	var e Employee

	fmt.Println("\nQ07:", e)
}

// ============================================================
// Q08. Anonymous struct.
// ============================================================
func q08() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "Ashutosh",
		Age:  25,
	}

	fmt.Println("\nQ08:", person)
}

// ============================================================
// Q09. Struct as function parameter.
// ============================================================
func q09() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Ashutosh", 25}

	printPerson(p)

	fmt.Println("\nQ09 done")
}

func printPerson(p struct {
	Name string
	Age  int
}) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// ============================================================
// Q10. Named struct type as function parameter.
// ============================================================
func q10() {
	p := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	showPerson(p)

	fmt.Println("\nQ10 done")
}

type PersonBasic struct {
	Name string
	Age  int
}

func showPerson(p PersonBasic) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// ============================================================
// Q11. Pointer to struct.
// ============================================================
func q11() {
	p := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	ptr := &p

	fmt.Println("\nQ11:", ptr.Name, ptr.Age)
}

// ============================================================
// Q12. Modify struct through pointer.
// ============================================================
func q12() {
	p := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	ptr := &p

	ptr.Age = 26

	fmt.Println("\nQ12:", p)
}

// ============================================================
// Q13. Struct pointer parameter.
// ============================================================
func q13() {
	p := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	updateAge(&p, 30)

	fmt.Println("\nQ13:", p)
}

func updateAge(p *PersonBasic, age int) {
	p.Age = age
}

// ============================================================
// Q14. Return struct from function.
// ============================================================
func q14() {
	p := createPerson()

	fmt.Println("\nQ14:", p)
}

func createPerson() PersonBasic {
	return PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}
}

// ============================================================
// Q15. Return pointer to struct.
// ============================================================
func q15() {
	p := createPersonPointer()

	fmt.Println("\nQ15:", p)
}

func createPersonPointer() *PersonBasic {
	return &PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}
}

// ============================================================
// Q16. Struct containing another struct.
// ============================================================
func q16() {
	type Address struct {
		City  string
		State string
	}

	type Employee struct {
		Name    string
		Address Address
	}

	e := Employee{
		Name: "Ashutosh",
		Address: Address{
			City:  "Bhubaneswar",
			State: "Odisha",
		},
	}

	fmt.Println("\nQ16:", e)
	fmt.Println("City:", e.Address.City)
}

// ============================================================
// Q17. Modify nested struct.
// ============================================================
func q17() {
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

	u.Address.City = "Hyderabad"

	fmt.Println("\nQ17:", u)
}

// ============================================================
// Q18. Struct with slice.
// ============================================================
func q18() {
	type Employee struct {
		Name   string
		Skills []string
	}

	e := Employee{
		Name:   "Ashutosh",
		Skills: []string{"Go", "PostgreSQL", "React"},
	}

	e.Skills = append(e.Skills, "Docker")

	fmt.Println("\nQ18:", e)
}

// ============================================================
// Q19. Struct with map.
// ============================================================
func q19() {
	type Employee struct {
		Name   string
		Skills map[string]int
	}

	e := Employee{
		Name: "Ashutosh",
		Skills: map[string]int{
			"Go":         90,
			"PostgreSQL": 85,
		},
	}

	e.Skills["Docker"] = 80

	fmt.Println("\nQ19:", e)
}

// ============================================================
// Q20. Struct with pointer field.
// ============================================================
func q20() {
	type Employee struct {
		Name   string
		Salary *int
	}

	salary := 70000

	e := Employee{
		Name:   "Ashutosh",
		Salary: &salary,
	}

	fmt.Println("\nQ20:", e.Name, *e.Salary)
}

// ============================================================
// Q21. Define a method on a struct.
// ============================================================
func q21() {
	p := PersonMethod{
		Name: "Ashutosh",
	}

	p.Greet()

	fmt.Println("\nQ21 done")
}

type PersonMethod struct {
	Name string
}

func (p PersonMethod) Greet() {
	fmt.Println("Hello", p.Name)
}

// ============================================================
// Q22. Method returning a value.
// ============================================================
func q22() {
	r := Rectangle{
		Width:  10,
		Height: 20,
	}

	fmt.Println("\nQ22 area:", r.Area())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ============================================================
// Q23. Pointer receiver.
// ============================================================
func q23() {
	p := PersonMethod{
		Name: "Ashutosh",
	}

	p.ChangeName("Rahul")

	fmt.Println("\nQ23:", p.Name)
}

func (p *PersonMethod) ChangeName(name string) {
	p.Name = name
}

// ============================================================
// Q24. Value receiver vs pointer receiver.
// ============================================================
func q24() {
	p := PersonValue{
		Name: "Ashutosh",
	}

	p.SetValue("Rahul")

	fmt.Println("\nQ24 after value receiver:", p.Name)

	p.SetPointer("Rahul")

	fmt.Println("Q24 after pointer receiver:", p.Name)
}

type PersonValue struct {
	Name string
}

func (p PersonValue) SetValue(name string) {
	p.Name = name
}

func (p *PersonValue) SetPointer(name string) {
	p.Name = name
}

// ============================================================
// Q25. Struct method with multiple fields.
// ============================================================
func q25() {
	e := EmployeeAccount{
		Name:   "Ashutosh",
		Salary: 50000,
		Bonus:  10000,
	}

	fmt.Println("\nQ25 total:", e.TotalIncome())
}

type EmployeeAccount struct {
	Name   string
	Salary float64
	Bonus  float64
}

func (e EmployeeAccount) TotalIncome() float64 {
	return e.Salary + e.Bonus
}

// ============================================================
// Q26. Struct constructor pattern.
// ============================================================
func q26() {
	e := NewEmployee("Ashutosh", 25)

	fmt.Println("\nQ26:", e)
}

type EmployeeInfo struct {
	Name string
	Age  int
}

func NewEmployee(name string, age int) EmployeeInfo {
	return EmployeeInfo{
		Name: name,
		Age:  age,
	}
}

// ============================================================
// Q27. Constructor returning pointer.
// ============================================================
func q27() {
	e := NewEmployeePointer("Ashutosh", 25)

	fmt.Println("\nQ27:", e)
}

func NewEmployeePointer(name string, age int) *EmployeeInfo {
	return &EmployeeInfo{
		Name: name,
		Age:  age,
	}
}

// ============================================================
// Q28. Embedded struct.
// ============================================================
func q28() {
	type Address struct {
		City string
	}

	type Employee struct {
		Name string
		Address
	}

	e := Employee{
		Name: "Ashutosh",
		Address: Address{
			City: "Bhubaneswar",
		},
	}

	fmt.Println("\nQ28:", e.Name, e.City)
}

// ============================================================
// Q29. Struct embedding and promoted fields.
// ============================================================
func q29() {
	type Contact struct {
		Email string
	}

	type User struct {
		Name string
		Contact
	}

	u := User{
		Name: "Ashutosh",
		Contact: Contact{
			Email: "ashutosh@example.com",
		},
	}

	fmt.Println("\nQ29:", u.Name, u.Email)
}

// ============================================================
// Q30. Embedded method.
// ============================================================
func q30() {
	type Address struct {
		City string
	}

	type User struct {
		Address
	}

	a := User{
		Address: Address{
			City: "Bhubaneswar",
		},
	}

	a.PrintCity()
	fmt.Println("\nQ30 done")
}

func (a Address) PrintCity() {
	fmt.Println("City:", a.City)
}

// ============================================================
// Q31. Struct implements interface.
// ============================================================
func q31() {
	var s Speaker = EmployeeSpeaker{
		Name: "Ashutosh",
	}

	fmt.Println("\nQ31:")
	s.Speak()
}

type Speaker interface {
	Speak()
}

type EmployeeSpeaker struct {
	Name string
}

func (e EmployeeSpeaker) Speak() {
	fmt.Println("I am", e.Name)
}

// ============================================================
// Q32. Pointer receiver and interface.
// ============================================================
func q32() {
	e := EmployeePointerSpeaker{
		Name: "Ashutosh",
	}

	var s Speaker = &e

	fmt.Println("\nQ32:")
	s.Speak()
}

type EmployeePointerSpeaker struct {
	Name string
}

func (e *EmployeePointerSpeaker) Speak() {
	fmt.Println("I am", e.Name)
}

// ============================================================
// Q33. Compare two structs.
// ============================================================
func q33() {
	a := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	b := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	fmt.Println("\nQ33 equal:", a == b)
}

// ============================================================
// Q34. Struct comparison with unequal values.
// ============================================================
func q34() {
	a := PersonBasic{
		Name: "Ashutosh",
		Age: 25,
	}

	b := PersonBasic{
		Name: "Ashutosh",
		Age: 30,
	}

	fmt.Println("\nQ34 equal:", a == b)
}

// ============================================================
// Q35. Struct containing non-comparable fields.
// ============================================================
func q35() {
	type User struct {
		Name   string
		Skills []string
	}

	a := User{
		Name:   "Ashutosh",
		Skills: []string{"Go"},
	}

	b := User{
		Name:   "Ashutosh",
		Skills: []string{"Go"},
	}

	/*
		a == b is INVALID because []string is not comparable.

		We compare manually.
	*/

	equal := a.Name == b.Name && len(a.Skills) == len(b.Skills)

	if equal {
		for i := range a.Skills {
			if a.Skills[i] != b.Skills[i] {
				equal = false
				break
			}
		}
	}

	fmt.Println("\nQ35 equal:", equal)
}

// ============================================================
// Q36. Copy a struct.
// ============================================================
func q36() {
	a := PersonBasic{
		Name: "Ashutosh",
		Age:  25,
	}

	b := a

	b.Age = 30

	fmt.Println("\nQ36 original:", a)
	fmt.Println("Q36 copy:", b)
}

// ============================================================
// Q37. Struct copy with slice field.
// ============================================================
func q37() {
	type User struct {
		Name   string
		Skills []string
	}

	a := User{
		Name:   "Ashutosh",
		Skills: []string{"Go", "SQL"},
	}

	b := a

	b.Skills[0] = "Rust"

	fmt.Println("\nQ37 original:", a)
	fmt.Println("Q37 copy:", b)

	/*
		The struct itself is copied, but the slice header points
		to the same underlying array.

		Therefore both structs can observe the changed slice element.
	*/
}

// ============================================================
// Q38. Deep-copy a struct containing a slice.
// ============================================================
func q38() {
	type User struct {
		Name   string
		Skills []string
	}

	a := User{
		Name:   "Ashutosh",
		Skills: []string{"Go", "SQL"},
	}

	b := User{
		Name:   a.Name,
		Skills: append([]string(nil), a.Skills...),
	}

	b.Skills[0] = "Rust"

	fmt.Println("\nQ38 original:", a)
	fmt.Println("Q38 deep copy:", b)
}

// ============================================================
// Q39. Struct tags.
// ============================================================
func q39() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	u := User{
		Name:  "Ashutosh",
		Email: "ashutosh@example.com",
	}

	data, _ := json.Marshal(u)

	fmt.Println("\nQ39:", string(data))
}

// ============================================================
// Q40. Unmarshal JSON into struct.
// ============================================================
func q40() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	jsonData := `{
		"name": "Ashutosh",
		"email": "ashutosh@example.com",
		"age": 25
	}`

	var u User

	err := json.Unmarshal([]byte(jsonData), &u)

	fmt.Println("\nQ40 error:", err)
	fmt.Println("Q40 user:", u)
}

// ============================================================
// Q41. Exported vs unexported struct fields.
// ============================================================
func q41() {
	type User struct {
		Name  string
		email string
	}

	u := User{
		Name:  "Ashutosh",
		email: "private@example.com",
	}

	fmt.Println("\nQ41 exported:", u.Name)
	fmt.Println("Q41 unexported inside same package:", u.email)

	/*
		Fields beginning with uppercase letters are exported.

		Fields beginning with lowercase letters are unexported.

		Unexported fields cannot be accessed directly from another
		package.
	*/
}

// ============================================================
// Q42. JSON behavior of unexported fields.
// ============================================================
func q42() {
	type User struct {
		Name  string `json:"name"`
		email string `json:"email"`
	}

	u := User{
		Name:  "Ashutosh",
		email: "private@example.com",
	}

	data, _ := json.Marshal(u)

	fmt.Println("\nQ42:", string(data))

	/*
		encoding/json ignores unexported fields.
	*/
}

// ============================================================
// Q43. Struct with JSON omitempty.
// ============================================================
func q43() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email,omitempty"`
		Age   int    `json:"age,omitempty"`
	}

	u := User{
		Name: "Ashutosh",
	}

	data, _ := json.Marshal(u)

	fmt.Println("\nQ43:", string(data))
}

// ============================================================
// Q44. Struct with time-like data represented as string.
// ============================================================
func q44() {
	type Employee struct {
		Name      string
		CreatedAt string
	}

	e := Employee{
		Name:      "Ashutosh",
		CreatedAt: "2026-08-10",
	}

	fmt.Println("\nQ44:", e)
}

// ============================================================
// Q45. Anonymous nested struct.
// ============================================================
func q45() {
	type Employee struct {
		Name    string
		Address struct {
			City  string
			State string
		}
	}

	e := Employee{
		Name: "Ashutosh",
	}

	e.Address.City = "Bhubaneswar"
	e.Address.State = "Odisha"

	fmt.Println("\nQ45:", e)
}

// ============================================================
// Q46. Struct containing pointer to another struct.
// ============================================================
func q46() {
	type Address struct {
		City string
	}

	type Employee struct {
		Name    string
		Address *Address
	}

	e := Employee{
		Name: "Ashutosh",
		Address: &Address{
			City: "Bhubaneswar",
		},
	}

	fmt.Println("\nQ46:", e.Name, e.Address.City)
}

// ============================================================
// Q47. Recursive struct.
// ============================================================
func q47() {
	type Node struct {
		Value int
		Next  *Node
	}

	first := &Node{Value: 10}
	second := &Node{Value: 20}
	third := &Node{Value: 30}

	first.Next = second
	second.Next = third

	fmt.Println("\nQ47:", first.Value, first.Next.Value, first.Next.Next.Value)
}

// ============================================================
// Q48. Binary tree node.
// ============================================================
func q48() {
	type TreeNode struct {
		Value int
		Left  *TreeNode
		Right *TreeNode
	}

	root := &TreeNode{
		Value: 10,
		Left: &TreeNode{
			Value: 5,
		},
		Right: &TreeNode{
			Value: 15,
		},
	}

	fmt.Println("\nQ48 root:", root.Value)
	fmt.Println("Q48 left:", root.Left.Value)
	fmt.Println("Q48 right:", root.Right.Value)
}

// ============================================================
// Q49. Struct with private data and methods.
// ============================================================
func q49() {
	account := NewBankAccount("Ashutosh", 1000)

	account.Deposit(500)

	fmt.Println("\nQ49:", account.Balance())
}

type BankAccount struct {
	name    string
	balance float64
}

func NewBankAccount(name string, balance float64) *BankAccount {
	return &BankAccount{
		name:    name,
		balance: balance,
	}
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

func (a BankAccount) Balance() float64 {
	return a.balance
}

// ============================================================
// Q50. Struct method validation.
// ============================================================
func q50() {
	u := UserValidator{
		Name:  "Ashutosh",
		Email: "ashutosh@example.com",
		Age:   25,
	}

	fmt.Println("\nQ50 valid:", u.IsValid())
}

type UserValidator struct {
	Name  string
	Email string
	Age   int
}

func (u UserValidator) IsValid() bool {
	return u.Name != "" &&
		u.Email != "" &&
		u.Age >= 18
}

// ============================================================
// Q51. Struct embedding with field name collision.
// ============================================================
func q51() {
	type Contact struct {
		Name string
	}

	type Employee struct {
		Name string
		Contact
	}

	e := Employee{
		Name: "Employee Name",
		Contact: Contact{
			Name: "Contact Name",
		},
	}

	fmt.Println("\nQ51 employee name:", e.Name)
	fmt.Println("Q51 contact name:", e.Contact.Name)
}

// ============================================================
// Q52. Method promotion with embedding.
// ============================================================
func q52() {
	type Logger struct{}

	type Service struct {
		Logger
	}

	logger := Service{}
	logger.Log()

	fmt.Println("\nQ52 done")
}

type Logger struct{}

func (Logger) Log() {
	fmt.Println("Logging...")
}

// ============================================================
// Q53. Struct used as a map key.
// ============================================================
func q53() {
	type Point struct {
		X int
		Y int
	}

	points := map[Point]string{
		{X: 1, Y: 2}: "A",
		{X: 3, Y: 4}: "B",
	}

	fmt.Println("\nQ53:", points[Point{X: 1, Y: 2}])
}

// ============================================================
// Q54. Struct as a map key: non-comparable field warning.
// ============================================================
func q54() {
	/*
		This is INVALID:

		    type Person struct {
		        Name   string
		        Skills []string
		    }

		    map[Person]string

		Reason:
		    []string is not comparable.

		Use only comparable fields in a struct used as a map key.
	*/

	type Point struct {
		X int
		Y int
	}

	m := map[Point]string{
		{1, 2}: "Point A",
	}

	fmt.Println("\nQ54:", m[Point{1, 2}])
}

// ============================================================
// Q55. Struct method chaining by returning pointer.
// ============================================================
func q55() {
	u := NewBuilder().
		SetName("Ashutosh").
		SetAge(25)

	fmt.Println("\nQ55:", u)
}

type UserBuilder struct {
	Name string
	Age  int
}

func NewBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) SetName(name string) *UserBuilder {
	u.Name = name
	return u
}

func (u *UserBuilder) SetAge(age int) *UserBuilder {
	u.Age = age
	return u
}

// ============================================================
// Q56. Struct as receiver: calculate total.
// ============================================================
func q56() {
	order := Order{
		Items: []Item{
			{Name: "Keyboard", Price: 1000, Quantity: 2},
			{Name: "Mouse", Price: 500, Quantity: 1},
		},
	}

	fmt.Println("\nQ56 total:", order.Total())
}

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	Items []Item
}

func (o Order) Total() float64 {
	total := 0.0

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

// ============================================================
// Q57. Update struct inside a slice.
// ============================================================
func q57() {
	type Employee struct {
		Name   string
		Salary int
	}

	employees := []Employee{
		{Name: "A", Salary: 50000},
		{Name: "B", Salary: 60000},
	}

	for i := range employees {
		if employees[i].Name == "B" {
			employees[i].Salary = 70000
		}
	}

	fmt.Println("\nQ57:", employees)
}

// ============================================================
// Q58. Find highest salary from struct slice.
// ============================================================
func q58() {
	type Employee struct {
		Name   string
		Salary int
	}

	employees := []Employee{
		{Name: "A", Salary: 50000},
		{Name: "B", Salary: 80000},
		{Name: "C", Salary: 70000},
	}

	highest := employees[0]

	for _, employee := range employees[1:] {
		if employee.Salary > highest.Salary {
			highest = employee
		}
	}

	fmt.Println("\nQ58:", highest)
}

// ============================================================
// Q59. Group structs by department.
// ============================================================
func q59() {
	type Employee struct {
		Name       string
		Department string
	}

	employees := []Employee{
		{Name: "A", Department: "IT"},
		{Name: "B", Department: "HR"},
		{Name: "C", Department: "IT"},
		{Name: "D", Department: "Finance"},
	}

	groups := map[string][]Employee{}

	for _, employee := range employees {
		groups[employee.Department] =
			append(groups[employee.Department], employee)
	}

	fmt.Println("\nQ59:", groups)
}

// ============================================================
// Q60. Practical interview problem:
//      model an API response with nested structs and JSON.
// ============================================================
func q60() {
	type Address struct {
		City  string `json:"city"`
		State string `json:"state"`
	}

	type User struct {
		ID      int     `json:"id"`
		Name    string  `json:"name"`
		Address Address `json:"address"`
	}

	type APIResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    User   `json:"data"`
	}

	response := APIResponse{
		Success: true,
		Message: "User fetched successfully",
		Data: User{
			ID:   101,
			Name: "Ashutosh",
			Address: Address{
				City:  "Bhubaneswar",
				State: "Odisha",
			},
		},
	}

	data, err := json.MarshalIndent(response, "", "  ")

	fmt.Println("\nQ60 error:", err)
	fmt.Println("Q60 response:")
	fmt.Println(string(data))
}

/*
=====================================================================
DEEP STRUCT INTERVIEW NOTES
=====================================================================

1. WHAT IS A STRUCT?

A struct is a user-defined type that groups multiple fields.

Example:

    type Employee struct {
        Name   string
        Age    int
        Salary float64
    }

It is similar to a class's data structure, but Go structs do not
automatically come with inheritance or constructors.

---------------------------------------------------------------------

2. STRUCT VS CLASS

Go does NOT have traditional classes.

A struct can contain:

    fields

and methods can be defined separately:

    func (e Employee) Print() {}

Go uses composition instead of traditional class inheritance.

---------------------------------------------------------------------

3. STRUCT ZERO VALUE

If you write:

    var e Employee

all fields receive their zero values.

Example:

    string -> ""
    int    -> 0
    bool   -> false
    pointer -> nil
    slice -> nil
    map -> nil

---------------------------------------------------------------------

4. NAMED INITIALIZATION

Preferred:

    e := Employee{
        Name: "Ashutosh",
        Age: 25,
    }

It is safer when the struct has many fields.

---------------------------------------------------------------------

5. POSITIONAL INITIALIZATION

You can write:

    e := Employee{"Ashutosh", 25}

But this depends on field order and is less maintainable.

---------------------------------------------------------------------

6. STRUCT COPY

Structs are value types.

Example:

    a := Employee{Name: "A"}
    b := a

b is a copy of the struct.

Changing simple fields in b does not change a.

However, if the struct contains slices, maps, pointers, etc.,
those fields can still refer to shared underlying data.

---------------------------------------------------------------------

7. STRUCT POINTER

Example:

    e := Employee{Name: "A"}
    p := &e

You can access fields directly:

    p.Name

Go automatically dereferences the pointer for field selection.

Equivalent conceptually to:

    (*p).Name

---------------------------------------------------------------------

8. VALUE RECEIVER

Example:

    func (e Employee) ChangeName(name string) {
        e.Name = name
    }

e is a copy.

The original struct is not changed.

---------------------------------------------------------------------

9. POINTER RECEIVER

Example:

    func (e *Employee) ChangeName(name string) {
        e.Name = name
    }

The original struct can be modified.

Use pointer receivers when:

- method modifies the receiver
- copying the struct is expensive
- consistency across method set is desired

---------------------------------------------------------------------

10. STRUCT EMBEDDING

Go supports embedding:

    type Employee struct {
        Person
        Salary int
    }

Employee gets promoted access to Person's fields and methods.

This is composition, not traditional inheritance.

---------------------------------------------------------------------

11. STRUCT TAGS

Example:

    type User struct {
        Name string `json:"name"`
    }

The tag tells packages such as encoding/json how to map fields.

---------------------------------------------------------------------

12. EXPORTED FIELDS

Uppercase:

    Name string

is exported.

Lowercase:

    name string

is unexported.

This matters across packages.

---------------------------------------------------------------------

13. JSON AND STRUCTS

encoding/json works with exported fields.

Example:

    type User struct {
        Name string `json:"name"`
    }

Marshal:

    data, err := json.Marshal(user)

Unmarshal:

    json.Unmarshal(data, &user)

The destination generally needs a pointer so the decoder can
populate the struct.

---------------------------------------------------------------------

14. OMITEMPTY

Example:

    Age int `json:"age,omitempty"`

If Age is its zero value, it can be omitted from JSON output.

---------------------------------------------------------------------

15. STRUCT COMPARISON

Structs can be compared using == only if all their fields are
comparable.

Comparable examples:

    int
    string
    bool
    arrays of comparable values
    structs containing comparable fields
    pointers
    channels
    interfaces (with caveats about dynamic values)

Not comparable:

    slices
    maps
    functions

So this is invalid:

    type User struct {
        Skills []string
    }

    user1 == user2

---------------------------------------------------------------------

16. STRUCT AS MAP KEY

A struct can be a map key only when the entire struct is comparable.

Valid:

    type Point struct {
        X int
        Y int
    }

    map[Point]string

Invalid:

    type Point struct {
        X int
        Values []int
    }

    map[Point]string

because []int is not comparable.

---------------------------------------------------------------------

17. STRUCT WITH SLICE

Example:

    type User struct {
        Skills []string
    }

Copying the struct copies the slice header, not necessarily the
underlying array.

Therefore:

    b := a
    b.Skills[0] = "Rust"

can affect a.Skills too.

---------------------------------------------------------------------

18. DEEP COPY

For a slice:

    copied := append([]string(nil), original...)

For maps, create a new map and copy each entry.

For nested structures, copy each reference-like field as required.

---------------------------------------------------------------------

19. RECURSIVE STRUCT

A struct cannot contain itself directly:

    type Node struct {
        Next Node // invalid recursive value
    }

But it can contain a pointer to itself:

    type Node struct {
        Next *Node
    }

This is how linked lists and trees are built.

---------------------------------------------------------------------

20. CONSTRUCTOR PATTERN

Go has no special constructor keyword.

Common pattern:

    func NewEmployee(...) *Employee {
        return &Employee{...}
    }

This is a convention, not a language requirement.

---------------------------------------------------------------------

21. PRIVATE DATA / ENCAPSULATION

Go uses capitalization for visibility.

Example:

    type Account struct {
        balance float64
    }

External packages cannot directly access balance.

Expose behavior through methods:

    func (a *Account) Deposit(...) {}

This is commonly used for encapsulation.

---------------------------------------------------------------------

22. STRUCT + INTERFACE

A struct implicitly implements an interface when it has all
required methods.

There is no explicit "implements" keyword.

Example:

    type Speaker interface {
        Speak()
    }

    type Person struct{}

    func (Person) Speak() {}

Person implements Speaker.

---------------------------------------------------------------------

23. POINTER RECEIVER + INTERFACE

If:

    func (p *Person) Speak() {}

then *Person implements Speaker.

Person may not implement Speaker.

This is a very common Go interview question.

---------------------------------------------------------------------

24. STRUCT VS MAP

Struct:

    fixed, known fields

Map:

    dynamic keys

Use a struct when the data model is known:

    Employee{
        Name,
        Age,
        Salary,
    }

Use a map when keys are dynamic:

    map[string]interface{}

Although for modern Go code, prefer concrete structs over
map[string]interface{} when the schema is known.

---------------------------------------------------------------------

25. STRUCT VS SLICE

Struct:

    different named fields

Slice:

    ordered collection of same element type

Example:

    Employee{Name: "A", Age: 25}

versus:

    []int{10, 20, 30}

---------------------------------------------------------------------

26. IMPORTANT INTERVIEW QUESTIONS

Practice these without looking:

1. What is a struct?
2. Why are structs used?
3. Struct vs class?
4. Struct vs map?
5. What is struct zero value?
6. Named vs positional initialization?
7. What is an anonymous struct?
8. Can a struct contain another struct?
9. Can a struct contain a pointer?
10. Can a struct contain a slice?
11. Can a struct contain a map?
12. What is a pointer to struct?
13. Value receiver vs pointer receiver?
14. When should you use pointer receiver?
15. What is struct embedding?
16. Does Go support inheritance?
17. What is composition?
18. What are struct tags?
19. How does JSON work with structs?
20. What does omitempty mean?
21. Exported vs unexported fields?
22. Can structs be compared?
23. Why can't structs with slices be compared?
24. Can structs be map keys?
25. Why can't a struct with a slice be a map key?
26. How are structs copied?
27. What happens when a struct contains a slice?
28. What happens when a struct contains a map?
29. What is a recursive struct?
30. How do you create a linked list node?
31. How do you create a tree node?
32. What is a constructor pattern in Go?
33. Can a function return a struct?
34. Can a function return *Struct?
35. How do you modify a struct inside a function?
36. What is method promotion?
37. How does embedding affect methods?
38. How does a struct implement an interface?
39. Pointer receiver and interface method set?
40. How do you encapsulate struct fields?

---------------------------------------------------------------------
MOST IMPORTANT FOR YOUR 2-YEAR GO INTERVIEW
---------------------------------------------------------------------

Master these first:

Q01  Basic struct
Q05  Named initialization
Q07  Zero values
Q11  Struct pointers
Q13  Pointer parameters
Q16  Nested structs
Q18  Struct + slice
Q19  Struct + map
Q21  Methods
Q23  Pointer receiver
Q24  Value vs pointer receiver
Q28  Embedding
Q31  Interface
Q32  Pointer receiver + interface
Q33-Q35 Struct comparison
Q36-Q38 Struct copying
Q39-Q43 JSON + tags
Q47-Q48 Recursive structs
Q49 Encapsulation
Q53-Q54 Struct as map key
Q57-Q59 Struct slices/grouping
Q60 API response modeling

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

	fmt.Println("\n============================================================")
	fmt.Println("ALL 60 STRUCT QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}
