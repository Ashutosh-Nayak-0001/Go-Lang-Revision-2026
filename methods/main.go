package main

import "fmt"

/*
GO METHODS — 40 QUESTIONS & ANSWERS
====================================

Run:
    go run methods_40.go

LEVEL:
    Q01-Q15  Beginner
    Q16-Q30  Intermediate
    Q31-Q40  Interview / Advanced

IMPORTANT:
A method is a function associated with a receiver.

Basic syntax:

    func (receiver ReceiverType) MethodName(parameters) returnType {
        ...
    }

Example:

    type User struct {
        Name string
    }

    func (u User) Greet() {
        fmt.Println("Hello", u.Name)
    }

The part:

    (u User)

is called the receiver.
*/

func separator(n int) {
	fmt.Printf("\n================ Q%02d ================\n", n)
}

// ============================================================
// Q01. What is a method?
// ============================================================
func q01() {
	separator(1)

	type User struct {
		Name string
	}

	func1 := func(u User) {
		fmt.Println("Hello", u.Name)
	}

	func1(User{Name: "Ashutosh"})

	/*
		A method is a function associated with a receiver.

		Normal function:
		    func greet(u User) {}

		Method:
		    func (u User) Greet() {}
	*/
}

// ============================================================
// Q02. Basic method with value receiver.
// ============================================================
func q02() {
	separator(2)

	u := User02{Name: "Ashutosh"}

	u.Greet()
}

type User02 struct {
	Name string
}

func (u User02) Greet() {
	fmt.Println("Hello", u.Name)
}

// ============================================================
// Q03. Method with parameters.
// ============================================================
func q03() {
	separator(3)

	u := User03{Name: "Ashutosh"}

	u.Greet("Good morning")
}

type User03 struct {
	Name string
}

func (u User03) Greet(message string) {
	fmt.Println(message, u.Name)
}

// ============================================================
// Q04. Method returning a value.
// ============================================================
func q04() {
	separator(4)

	r := Rectangle04{
		Width:  10,
		Height: 5,
	}

	fmt.Println("Area:", r.Area())
}

type Rectangle04 struct {
	Width  float64
	Height float64
}

func (r Rectangle04) Area() float64 {
	return r.Width * r.Height
}

// ============================================================
// Q05. Method with multiple return values.
// ============================================================
func q05() {
	separator(5)

	u := User05{Name: "Ashutosh", Age: 25}

	name, age := u.Info()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

type User05 struct {
	Name string
	Age  int
}

func (u User05) Info() (string, int) {
	return u.Name, u.Age
}

// ============================================================
// Q06. Method can access receiver fields.
// ============================================================
func q06() {
	separator(6)

	e := Employee06{
		Name:   "Ashutosh",
		Salary: 50000,
	}

	e.PrintSalary()
}

type Employee06 struct {
	Name   string
	Salary int
}

func (e Employee06) PrintSalary() {
	fmt.Println(e.Name, "salary:", e.Salary)
}

// ============================================================
// Q07. Value receiver does not modify original struct.
// ============================================================
func q07() {
	separator(7)

	u := User07{Name: "Ashutosh"}

	u.ChangeName("Rahul")

	fmt.Println("Original:", u.Name)
}

type User07 struct {
	Name string
}

func (u User07) ChangeName(name string) {
	u.Name = name

	fmt.Println("Inside method:", u.Name)
}

/*
ANSWER:
A value receiver receives a copy of the struct.

Therefore changing u.Name inside the method does not change
the original struct.
*/

// ============================================================
// Q08. Pointer receiver modifies original struct.
// ============================================================
func q08() {
	separator(8)

	u := User08{Name: "Ashutosh"}

	u.ChangeName("Rahul")

	fmt.Println("Original:", u.Name)
}

type User08 struct {
	Name string
}

func (u *User08) ChangeName(name string) {
	u.Name = name
}

// ============================================================
// Q09. Value receiver vs pointer receiver.
// ============================================================
func q09() {
	separator(9)

	u := User09{Name: "Ashutosh"}

	u.Print()
	u.ChangeName("Rahul")

	fmt.Println(u.Name)
}

type User09 struct {
	Name string
}

func (u User09) Print() {
	fmt.Println("Name:", u.Name)
}

func (u *User09) ChangeName(name string) {
	u.Name = name
}

/*
VALUE RECEIVER:

    func (u User) Method()

Receives a copy.

POINTER RECEIVER:

    func (u *User) Method()

Receives pointer to original value.
*/

// ============================================================
// Q10. Can a value call a pointer receiver method?
// ============================================================
func q10() {
	separator(10)

	u := User10{Name: "Ashutosh"}

	u.ChangeName("Rahul")

	fmt.Println(u.Name)
}

type User10 struct {
	Name string
}

func (u *User10) ChangeName(name string) {
	u.Name = name
}

/*
Go automatically takes the address when the value is addressable.

Conceptually:

    u.ChangeName("Rahul")

can be treated like:

    (&u).ChangeName("Rahul")
*/

// ============================================================
// Q11. Can a pointer call a value receiver method?
// ============================================================
func q11() {
	separator(11)

	u := &User11{Name: "Ashutosh"}

	u.Print()
}

type User11 struct {
	Name string
}

func (u User11) Print() {
	fmt.Println("Name:", u.Name)
}

// ============================================================
// Q12. Method with pointer receiver.
// ============================================================
func q12() {
	separator(12)

	counter := Counter12{}

	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println("Count:", counter.Value)
}

type Counter12 struct {
	Value int
}

func (c *Counter12) Increment() {
	c.Value++
}

// ============================================================
// Q13. Method can return receiver data.
// ============================================================
func q13() {
	separator(13)

	p := Product13{
		Name:  "Laptop",
		Price: 50000,
	}

	fmt.Println(p.PriceWithTax(18))
}

type Product13 struct {
	Name  string
	Price float64
}

func (p Product13) PriceWithTax(tax float64) float64 {
	return p.Price + (p.Price * tax / 100)
}

// ============================================================
// Q14. Method can call another method.
// ============================================================
func q14() {
	separator(14)

	r := Rectangle14{
		Width:  10,
		Height: 5,
	}

	r.PrintArea()
}

type Rectangle14 struct {
	Width  float64
	Height float64
}

func (r Rectangle14) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle14) PrintArea() {
	fmt.Println("Area:", r.Area())
}

// ============================================================
// Q15. Method on custom type.
// ============================================================
func q15() {
	separator(15)

	age := Age15(25)

	fmt.Println(age.IsAdult())
}

type Age15 int

func (a Age15) IsAdult() bool {
	return a >= 18
}

// ============================================================
// Q16. Methods can be defined on named non-struct types.
// ============================================================
func q16() {
	separator(16)

	name := Name16("Ashutosh")

	fmt.Println(name.Upper())
}

type Name16 string

func (n Name16) Upper() string {
	result := ""

	for _, c := range string(n) {
		if c >= 'a' && c <= 'z' {
			result += string(c - 32)
		} else {
			result += string(c)
		}
	}

	return result
}

/*
Methods can be defined on a named type.

They cannot be defined on arbitrary built-in types directly.

Example that is NOT allowed:

    func (s string) Something() {}

Instead create:

    type MyString string
*/

// ============================================================
// Q17. Method on slice type.
// ============================================================
func q17() {
	separator(17)

	numbers := Numbers17{10, 20, 30}

	fmt.Println("Sum:", numbers.Sum())
}

type Numbers17 []int

func (n Numbers17) Sum() int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}

// ============================================================
// Q18. Method on map type.
// ============================================================
func q18() {
	separator(18)

	scores := Scores18{
		"Go":     90,
		"Postgres": 85,
	}

	fmt.Println("Average:", scores.Average())
}

type Scores18 map[string]int

func (s Scores18) Average() float64 {
	if len(s) == 0 {
		return 0
	}

	total := 0

	for _, score := range s {
		total += score
	}

	return float64(total) / float64(len(s))
}

// ============================================================
// Q19. Method on pointer receiver with nil check.
// ============================================================
func q19() {
	separator(19)

	var u *User19

	u.Print()
}

type User19 struct {
	Name string
}

func (u *User19) Print() {
	if u == nil {
		fmt.Println("User is nil")
		return
	}

	fmt.Println(u.Name)
}

// ============================================================
// Q20. Method returning pointer to receiver.
// ============================================================
func q20() {
	separator(20)

	u := User20{Name: "Ashutosh"}

	result := u.Self()

	fmt.Println(result.Name)
}

type User20 struct {
	Name string
}

func (u *User20) Self() *User20 {
	return u
}

// ============================================================
// Q21. Constructor-like function vs method.
// ============================================================
func q21() {
	separator(21)

	u := NewUser21("Ashutosh")

	u.Greet()
}

type User21 struct {
	Name string
}

func NewUser21(name string) *User21 {
	return &User21{Name: name}
}

func (u User21) Greet() {
	fmt.Println("Hello", u.Name)
}

/*
Important:

NewUser21 is NOT a constructor in the language sense.

It is simply a convention for a function that creates a value.

Go has no constructor keyword.
*/

// ============================================================
// Q22. Method chaining.
// ============================================================
func q22() {
	separator(22)

	u := User22{}

	u.SetName("Ashutosh").
		SetAge(25).
		Print()
}

type User22 struct {
	Name string
	Age  int
}

func (u *User22) SetName(name string) *User22 {
	u.Name = name
	return u
}

func (u *User22) SetAge(age int) *User22 {
	u.Age = age
	return u
}

func (u *User22) Print() *User22 {
	fmt.Println(u.Name, u.Age)
	return u
}

// ============================================================
// Q23. Methods with same name on different types.
// ============================================================
func q23() {
	separator(23)

	car := Car23{}
	bike := Bike23{}

	car.Move()
	bike.Move()
}

type Car23 struct{}

func (Car23) Move() {
	fmt.Println("Car moves")
}

type Bike23 struct{}

func (Bike23) Move() {
	fmt.Println("Bike moves")
}

/*
Methods belong to a type's method set.

Two unrelated types can have methods with the same name.
*/

// ============================================================
// Q24. Methods and interfaces.
// ============================================================
func q24() {
	separator(24)

	var s Speaker24 = Person24{Name: "Ashutosh"}

	s.Speak()
}

type Speaker24 interface {
	Speak()
}

type Person24 struct {
	Name string
}

func (p Person24) Speak() {
	fmt.Println("Hello", p.Name)
}

// ============================================================
// Q25. Pointer receiver and interface implementation.
// ============================================================
func q25() {
	separator(25)

	p := &Person25{Name: "Ashutosh"}

	var s Speaker25 = p

	s.Speak()
}

type Speaker25 interface {
	Speak()
}

type Person25 struct {
	Name string
}

func (p *Person25) Speak() {
	fmt.Println("Hello", p.Name)
}

/*
Because Speak has a pointer receiver:

    *Person25 implements Speaker25

Person25 itself does not implement Speaker25.
*/

// ============================================================
// Q26. Value receiver and interface implementation.
// ============================================================
func q26() {
	separator(26)

	p := Person26{Name: "Ashutosh"}

	var a Speaker26 = p
	var b Speaker26 = &p

	a.Speak()
	b.Speak()
}

type Speaker26 interface {
	Speak()
}

type Person26 struct {
	Name string
}

func (p Person26) Speak() {
	fmt.Println("Hello", p.Name)
}

/*
With a value receiver, both:

    Person26
    *Person26

can satisfy Speaker26.
*/

// ============================================================
// Q27. Method set interview question.
// ============================================================
func q27() {
	separator(27)

	fmt.Println("See explanation in source code.")

	/*
		Question:

		    type Person struct{}

		    func (p Person) Speak() {}

		Who implements Speaker?

		    Person
		    *Person

		Answer:

		Both.

		Question:

		    func (p *Person) Speak() {}

		Who implements Speaker?

		    *Person only.
	*/
}

// ============================================================
// Q28. Method expression.
// ============================================================
func q28() {
	separator(28)

	u := User28{Name: "Ashutosh"}

	/*
		Method expression converts a method into a function.

		Normal:

		    u.Greet()

		Method expression:

		    User28.Greet(u)
	*/

	greet := User28.Greet

	greet(u)
}

type User28 struct {
	Name string
}

func (u User28) Greet() {
	fmt.Println("Hello", u.Name)
}

// ============================================================
// Q29. Method value.
// ============================================================
func q29() {
	separator(29)

	u := User29{Name: "Ashutosh"}

	/*
		Method value binds the receiver to the function.

		    greet := u.Greet

		Now greet can be called directly.
	*/

	greet := u.Greet

	greet()
}

type User29 struct {
	Name string
}

func (u User29) Greet() {
	fmt.Println("Hello", u.Name)
}

// ============================================================
// Q30. Method expression vs method value.
// ============================================================
func q30() {
	separator(30)

	u := User30{Name: "Ashutosh"}

	// Method expression.
	expression := User30.Greet
	expression(u)

	// Method value.
	value := u.Greet
	value()

	/*
	METHOD EXPRESSION:

	    User30.Greet(u)

	The receiver is supplied explicitly.

	METHOD VALUE:

	    u.Greet()

	The receiver is already bound.
	*/
}

type User30 struct {
	Name string
}

func (u User30) Greet() {
	fmt.Println("Hello", u.Name)
}

// ============================================================
// Q31. Method on embedded struct.
// ============================================================
func q31() {
	separator(31)

	e := Employee31{
		Person31: Person31{Name: "Ashutosh"},
	}

	e.Speak()
}

type Person31 struct {
	Name string
}

func (p Person31) Speak() {
	fmt.Println("Hello", p.Name)
}

type Employee31 struct {
	Person31
}

/*
Employee31 gets promoted access to Speak through embedding.

So:

    e.Speak()

works.
*/

// ============================================================
// Q32. Method overriding-like behavior with embedding.
// ============================================================
func q32() {
	separator(32)

	e := Employee32{
		Person32: Person32{Name: "Ashutosh"},
	}

	e.Speak()
	e.Person32.Speak()
}

type Person32 struct {
	Name string
}

func (p Person32) Speak() {
	fmt.Println("Person:", p.Name)
}

type Employee32 struct {
	Person32
}

func (e Employee32) Speak() {
	fmt.Println("Employee:", e.Name)
}

/*
Go does not have traditional class inheritance.

A promoted method can be shadowed by a method declared on the
outer type.
*/

// ============================================================
// Q33. Method can modify slice/map through value receiver.
// ============================================================
func q33() {
	separator(33)

	b := Basket33{
		Items: []string{},
	}

	b.Add("Laptop")
	b.Add("Mouse")

	fmt.Println(b.Items)
}

type Basket33 struct {
	Items []string
}

func (b Basket33) Add(item string) {
	b.Items = append(b.Items, item)

	/*
		This example demonstrates an important subtlety.

		A slice field contains a slice header.

		Whether the original slice length changes depends on how
		the slice is modified and whether append reallocates.

		For predictable mutation of the struct's slice field,
		use a pointer receiver.
	*/
}

// ============================================================
// Q34. Pointer receiver for reliable struct mutation.
// ============================================================
func q34() {
	separator(34)

	b := Basket34{}

	b.Add("Laptop")
	b.Add("Mouse")

	fmt.Println(b.Items)
}

type Basket34 struct {
	Items []string
}

func (b *Basket34) Add(item string) {
	b.Items = append(b.Items, item)
}

/*
Pointer receiver is preferred because the method changes the
struct's Items field itself.
*/

// ============================================================
// Q35. Method on a struct containing a map.
// ============================================================
func q35() {
	separator(35)

	u := UserScores35{
		Scores: map[string]int{},
	}

	u.SetScore("Go", 95)

	fmt.Println(u.Scores)
}

type UserScores35 struct {
	Scores map[string]int
}

func (u UserScores35) SetScore(subject string, score int) {
	u.Scores[subject] = score
}

/*
A map is a reference-like data structure.

The map itself can be modified through a copied struct receiver
because the copied map header refers to the same underlying map.

However, if you replace the map field itself, a pointer receiver
is needed to change the original struct field.
*/

// ============================================================
// Q36. Pointer receiver and large structs.
// ============================================================
func q36() {
	separator(36)

	data := LargeData36{
		A: 10,
		B: 20,
		C: 30,
		D: 40,
	}

	data.Print()
}

type LargeData36 struct {
	A int
	B int
	C int
	D int
}

func (d *LargeData36) Print() {
	fmt.Println(d.A, d.B, d.C, d.D)
}

/*
Pointer receivers can avoid copying large structs.

Performance should be measured rather than assumed, but pointer
receivers are commonly appropriate for large mutable types.
*/

// ============================================================
// Q37. Consistency of receiver choice.
// ============================================================
func q37() {
	separator(37)

	/*
		If a type has methods that mutate its state, pointer receivers
		are usually appropriate.

		For consistency, many methods on that type may also use
		pointer receivers.

		Example:

		    func (u *User) ChangeName(...)
		    func (u *User) ChangeAge(...)
		    func (u *User) Save(...)

		There is no absolute rule requiring every method to use the
		same receiver, but consistency improves readability.
	*/

	fmt.Println("Use receiver choice deliberately.")
}

// ============================================================
// Q38. Method can return an error.
// ============================================================
func q38() {
	separator(38)

	account := Account38{
		Balance: 1000,
	}

	err := account.Withdraw(1500)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Balance:", account.Balance)
}

type Account38 struct {
	Balance float64
}

func (a *Account38) Withdraw(amount float64) error {
	if amount > a.Balance {
		return fmt.Errorf("insufficient balance")
	}

	a.Balance -= amount
	return nil
}

// ============================================================
// Q39. Practical service method.
// ============================================================
func q39() {
	separator(39)

	service := UserService39{
		Repository: UserRepository39{},
	}

	user, err := service.GetUser(101)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user)
}

type User39 struct {
	ID   int
	Name string
}

type UserRepository39 struct{}

func (UserRepository39) FindByID(id int) (User39, error) {
	return User39{
		ID:   id,
		Name: "Ashutosh",
	}, nil
}

type UserService39 struct {
	Repository UserRepository39
}

func (s UserService39) GetUser(id int) (User39, error) {
	return s.Repository.FindByID(id)
}

/*
A method can contain application/business logic.

In a real Go backend:

    Handler
       |
       v
    Service method
       |
       v
    Repository method
       |
       v
    Database
*/

// ============================================================
// Q40. Final interview problem: value vs pointer receiver +
//      interface + method set.
// ============================================================
func q40() {
	separator(40)

	p := Person40{Name: "Ashutosh"}

	var a Speaker40 = p
	var b Speaker40 = &p

	a.Speak()
	b.Speak()

	p.ChangeName("Rahul")

	fmt.Println("Final:", p.Name)
}

type Speaker40 interface {
	Speak()
}

type Person40 struct {
	Name string
}

// Value receiver.
// Both Person40 and *Person40 can satisfy Speaker40.
func (p Person40) Speak() {
	fmt.Println("Speaking:", p.Name)
}

// Pointer receiver.
// This changes the original struct.
func (p *Person40) ChangeName(name string) {
	p.Name = name
}

/*
===============================================================
DEEP INTERVIEW NOTES
===============================================================

1. METHOD VS FUNCTION

Function:

    func Add(a, b int) int {
        return a + b
    }

Method:

    func (u User) Greet() {
        ...
    }

A method has a receiver.

---------------------------------------------------------------

2. RECEIVER

In:

    func (u User) Greet() {}

    u

is the receiver variable.

    User

is the receiver type.

---------------------------------------------------------------

3. VALUE RECEIVER

    func (u User) ChangeName(name string) {
        u.Name = name
    }

The receiver is a copy.

Changing u normally does not change the original struct.

---------------------------------------------------------------

4. POINTER RECEIVER

    func (u *User) ChangeName(name string) {
        u.Name = name
    }

The receiver points to the original value.

Changes are visible outside the method.

---------------------------------------------------------------

5. WHEN TO USE POINTER RECEIVER?

Common reasons:

- Method needs to modify receiver.
- Avoid copying a large struct.
- Keep receiver behavior consistent.
- The type contains state that should be mutated.

---------------------------------------------------------------

6. WHEN TO USE VALUE RECEIVER?

Common when:

- Receiver is small.
- Method does not mutate state.
- Copying the value is acceptable.
- The type behaves like a value.

---------------------------------------------------------------

7. CAN VALUE CALL POINTER METHOD?

Usually yes if the value is addressable.

Example:

    u.ChangeName()

Go can effectively take its address.

---------------------------------------------------------------

8. CAN POINTER CALL VALUE METHOD?

Yes.

If:

    func (u User) Greet() {}

then:

    u.Greet()

and:

    p.Greet()

can work where p is *User.

---------------------------------------------------------------

9. METHOD SET

Very important for interviews.

Suppose:

    type Speaker interface {
        Speak()
    }

And:

    func (u User) Speak() {}

Then both:

    User
    *User

can implement Speaker.

But:

    func (u *User) Speak() {}

means:

    *User

implements Speaker, while User does not.

---------------------------------------------------------------

10. INTERFACE + POINTER RECEIVER

Example:

    type Speaker interface {
        Speak()
    }

    type User struct{}

    func (u *User) Speak() {}

This works:

    var s Speaker = &User{}

This does NOT work:

    var s Speaker = User{}

because User's method set does not contain the pointer-receiver method.

---------------------------------------------------------------

11. METHOD EXPRESSION

Example:

    type User struct{}

    func (u User) Greet() {}

Method expression:

    f := User.Greet

Call:

    f(user)

The receiver becomes an explicit function argument.

---------------------------------------------------------------

12. METHOD VALUE

Example:

    f := user.Greet

The receiver is already bound.

Call:

    f()

---------------------------------------------------------------

13. METHOD EXPRESSION VS METHOD VALUE

Method expression:

    User.Greet(user)

Receiver supplied explicitly.

Method value:

    user.Greet()

Receiver already attached to function value.

---------------------------------------------------------------

14. METHODS ON CUSTOM TYPES

Allowed:

    type Age int

    func (a Age) IsAdult() bool {
        return a >= 18
    }

Not allowed:

    func (i int) Something() {}

You cannot define methods on a type defined in another package,
and the receiver base type must be defined in the same package.

---------------------------------------------------------------

15. METHODS ON SLICE TYPES

Allowed:

    type Numbers []int

    func (n Numbers) Sum() int {
        ...
    }

This is useful for creating domain-specific behavior.

---------------------------------------------------------------

16. METHODS ON MAP TYPES

Allowed:

    type Scores map[string]int

    func (s Scores) Average() float64 {
        ...
    }

---------------------------------------------------------------

17. METHODS AND EMBEDDING

Struct embedding can promote methods.

Example:

    type Person struct{}

    func (Person) Speak() {}

    type Employee struct {
        Person
    }

Then:

    e.Speak()

can access the embedded method.

Go does NOT have traditional class inheritance.

---------------------------------------------------------------

18. METHOD SHADOWING WITH EMBEDDING

If the outer struct defines the same method:

    func (Employee) Speak() {}

then:

    e.Speak()

uses Employee's method.

The embedded method can still be accessed:

    e.Person.Speak()

---------------------------------------------------------------

19. METHODS AND INTERFACES

An interface describes required methods.

Example:

    type Speaker interface {
        Speak()
    }

If:

    func (u User) Speak() {}

then User satisfies Speaker.

---------------------------------------------------------------

20. METHODS AND POLYMORPHISM

Different types can implement the same interface method:

    Car.Move()
    Bike.Move()

Then:

    func MoveVehicle(v Vehicle) {
        v.Move()
    }

can work with both.

---------------------------------------------------------------

21. METHODS ARE NOT CLASS METHODS

Go does not have traditional class methods.

Go has methods associated with receiver types.

---------------------------------------------------------------

22. NO METHOD OVERLOADING

Go does not support traditional method overloading.

You cannot define:

    func (u User) Print()
    func (u User) Print(name string)

with the same name on the same receiver type.

Use different method names or other API designs.

---------------------------------------------------------------

23. METHODS CAN RETURN ERRORS

Very common in backend Go:

    func (a *Account) Withdraw(amount float64) error

Caller:

    if err := account.Withdraw(100); err != nil {
        ...
    }

---------------------------------------------------------------

24. METHODS CAN RETURN MULTIPLE VALUES

Example:

    func (u User) Info() (string, int) {
        return u.Name, u.Age
    }

Go methods follow the same return-value rules as functions.

---------------------------------------------------------------

25. METHODS CAN HAVE PARAMETERS

Example:

    func (u User) Greet(message string) {
        ...
    }

Receiver and ordinary parameters are separate concepts.

---------------------------------------------------------------

26. POINTER RECEIVER DOES NOT MEAN POINTER FIELD

This:

    func (u *User) ChangeName() {}

means the receiver is a pointer.

It does NOT mean every field of User is a pointer.

---------------------------------------------------------------

27. VALUE RECEIVER DOES NOT ALWAYS MEAN NOTHING CAN MUTATE

Important subtlety:

If a struct contains reference-like fields such as:

    slice
    map
    pointer
    channel
    function

a copied struct can still refer to shared underlying data.

Example:

    type User struct {
        Scores map[string]int
    }

A value receiver can mutate:

    u.Scores["Go"] = 100

because the copied map header points to the same underlying map.

But replacing:

    u.Scores = newMap

only changes the receiver copy.

---------------------------------------------------------------

28. SLICE + VALUE RECEIVER SUBTLETY

A slice is a descriptor containing:

    pointer
    length
    capacity

Copying a slice copies this descriptor, not all underlying elements.

Therefore append behavior can be subtle.

For methods that intentionally change a struct's slice field,
a pointer receiver is usually clearer.

---------------------------------------------------------------

29. METHOD CHAINING

Possible:

    func (u *User) SetName(name string) *User {
        u.Name = name
        return u
    }

Then:

    u.SetName("Ashutosh").
      SetAge(25).
      Save()

This is useful for builder-style APIs.

Use it only when it improves readability.

---------------------------------------------------------------

30. CONSTRUCTOR PATTERN

Go has no constructor keyword.

Common convention:

    func NewUser(name string) *User {
        return &User{Name: name}
    }

Then:

    user := NewUser("Ashutosh")

---------------------------------------------------------------

31. METHOD RECEIVER NAMING

Usually use a short receiver name:

    func (u User) Greet() {}

Common style is one or two letters based on the type:

    u for User
    p for Person
    e for Employee
    r for Request

Do not use unnecessarily long names like:

    func (currentUser User) Greet()

unless there is a compelling reason.

---------------------------------------------------------------

32. METHODS AND ENCAPSULATION

Go does not have private/public methods using keywords.

Exporting is controlled by capitalization.

    func (u User) Save() {}

is exported.

    func (u User) save() {}

is unexported outside the package.

---------------------------------------------------------------

33. POINTER RECEIVER + NIL

A pointer receiver can be nil:

    var u *User

Calling a method is possible if the method handles nil safely.

Example:

    func (u *User) Print() {
        if u == nil {
            return
        }
    }

Be careful when dereferencing nil.

---------------------------------------------------------------

34. METHODS AND CONCURRENCY

A method is not automatically thread-safe.

If:

    func (c *Counter) Increment() {
        c.Value++
    }

is called concurrently by many goroutines, it can cause a race.

You may need:

    sync.Mutex
    sync.RWMutex
    atomic operations
    channels

depending on the design.

---------------------------------------------------------------

35. METHOD VS FUNCTION INTERVIEW ANSWER

Question:
What is the difference between a function and a method?

Strong answer:

"A function is a standalone callable unit, while a method is a
function associated with a receiver type. The receiver allows
the method to operate on the state or behavior associated with
that type."

---------------------------------------------------------------

36. VALUE VS POINTER RECEIVER INTERVIEW ANSWER

Question:
What is the difference?

Strong answer:

"A value receiver receives a copy of the receiver, so changes to
its fields normally do not affect the original value. A pointer
receiver receives a pointer to the original value, allowing the
method to mutate it and potentially avoiding copies of large
values."

---------------------------------------------------------------

37. METHOD SET INTERVIEW ANSWER

Question:
What is a method set?

Strong answer:

"A method set is the set of methods associated with a type that
determines which interfaces the type satisfies. For a defined
type T, its method set contains methods with receiver T. The
method set of *T contains methods with receiver T as well as
*T."

---------------------------------------------------------------

38. MOST COMMON INTERVIEW TRAPS

TRAP 1:
Go supports method overloading.

FALSE.

TRAP 2:
Pointer receiver is required for every method.

FALSE.

TRAP 3:
Value receiver can never modify underlying data.

FALSE in all cases.

Reference-like fields can still refer to shared data.

TRAP 4:
User automatically implements an interface when it has similar
fields.

FALSE.

Interfaces care about methods, not fields.

TRAP 5:
Go has class inheritance.

FALSE.

Go uses composition and embedding instead.

---------------------------------------------------------------

39. CODE YOU SHOULD BE ABLE TO WRITE IN AN INTERVIEW

Write this without looking:

    type User struct {
        Name string
        Age  int
    }

    func (u User) Greet() {
        fmt.Println("Hello", u.Name)
    }

    func (u *User) SetName(name string) {
        u.Name = name
    }

    func (u User) IsAdult() bool {
        return u.Age >= 18
    }

Then explain:

    User.Greet()
    user.Greet()
    (&user).SetName(...)
    user.SetName(...)
    user.IsAdult()

---------------------------------------------------------------

40. FINAL INTERVIEW CHECKLIST

You should be able to explain and code:

[ ] What is a method?
[ ] What is a receiver?
[ ] Value receiver
[ ] Pointer receiver
[ ] When to use pointer receiver
[ ] Method set
[ ] Pointer method set
[ ] Value vs pointer interface implementation
[ ] Methods on custom types
[ ] Methods on slices
[ ] Methods on maps
[ ] Methods with parameters
[ ] Methods with return values
[ ] Multiple return values
[ ] Methods returning errors
[ ] Method expressions
[ ] Method values
[ ] Method chaining
[ ] Constructor pattern
[ ] Embedded methods
[ ] Method shadowing
[ ] Methods + interfaces
[ ] Methods + polymorphism
[ ] Nil pointer receiver
[ ] Slice/map receiver behavior
[ ] No method overloading
[ ] No class inheritance
[ ] Exported/unexported methods
[ ] Methods and concurrency
[ ] Practical service methods

===============================================================
END
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

	fmt.Println("\n============================================================")
	fmt.Println("ALL 40 METHOD QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}
