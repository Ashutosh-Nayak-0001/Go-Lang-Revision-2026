package main

import (
	"fmt"
)

/*
GO INTERFACES — 60 QUESTIONS WITH ANSWERS
==========================================

Run:
    go run interfaces_60.go

LEVELS:
    Q01-Q20  Beginner
    Q21-Q40  Intermediate
    Q41-Q60  Interview / Advanced

CORE IDEA:

An interface defines a set of methods.

Example:

    type Speaker interface {
        Speak()
    }

A type satisfies the interface implicitly by implementing
all required methods.

There is NO "implements" keyword in Go.

*/

func separator(n int) {
	fmt.Printf("\n================ Q%02d ================\n", n)
}

// ============================================================
// Q01. What is an interface?
// ============================================================
func q01() {
	separator(1)

	type Speaker interface {
		Speak()
	}

	fmt.Println("Speaker is an interface containing Speak().")
}

// ============================================================
// Q02. Basic interface implementation.
// ============================================================
func q02() {
	separator(2)

	var s Speaker01 = Person01{Name: "Ashutosh"}

	s.Speak()
}

type Speaker01 interface {
	Speak()
}

type Person01 struct {
	Name string
}

func (p Person01) Speak() {
	fmt.Println("Hello, I am", p.Name)
}

// ============================================================
// Q03. Interfaces are implemented implicitly.
// ============================================================
func q03() {
	separator(3)

	var s Speaker02 = Dog02{Name: "Tommy"}

	s.Speak()
}

type Speaker02 interface {
	Speak()
}

type Dog02 struct {
	Name string
}

func (d Dog02) Speak() {
	fmt.Println(d.Name, "says Woof")
}

// ============================================================
// Q04. Multiple types can implement one interface.
// ============================================================
func q04() {
	separator(4)

	var a Speaker03 = Person03{Name: "Ashutosh"}
	var b Speaker03 = Dog03{Name: "Tommy"}

	a.Speak()
	b.Speak()
}

type Speaker03 interface {
	Speak()
}

type Person03 struct {
	Name string
}

func (p Person03) Speak() {
	fmt.Println("Person:", p.Name)
}

type Dog03 struct {
	Name string
}

func (d Dog03) Speak() {
	fmt.Println("Dog:", d.Name)
}

// ============================================================
// Q05. Interface as function parameter.
// ============================================================
func q05() {
	separator(5)

	printSpeaker(Person04{Name: "Ashutosh"})
	printSpeaker(Dog04{Name: "Tommy"})
}

type Speaker04 interface {
	Speak()
}

type Person04 struct {
	Name string
}

func (p Person04) Speak() {
	fmt.Println("Person:", p.Name)
}

type Dog04 struct {
	Name string
}

func (d Dog04) Speak() {
	fmt.Println("Dog:", d.Name)
}

func printSpeaker(s Speaker04) {
	s.Speak()
}

// ============================================================
// Q06. Interface with multiple methods.
// ============================================================
func q06() {
	separator(6)

	var p Employee06 = Employee06{Name: "Ashutosh"}

	printEmployee06(p)
}

type EmployeeOperations06 interface {
	Work()
	Report()
}

type Employee06 struct {
	Name string
}

func (e Employee06) Work() {
	fmt.Println(e.Name, "is working")
}

func (e Employee06) Report() {
	fmt.Println(e.Name, "submitted report")
}

func printEmployee06(e EmployeeOperations06) {
	e.Work()
	e.Report()
}

// ============================================================
// Q07. Interface can hold different concrete values.
// ============================================================
func q07() {
	separator(7)

	var s Speaker07

	s = Person07{Name: "Ashutosh"}
	s.Speak()

	s = Dog07{Name: "Tommy"}
	s.Speak()
}

type Speaker07 interface {
	Speak()
}

type Person07 struct {
	Name string
}

func (p Person07) Speak() {
	fmt.Println("Person:", p.Name)
}

type Dog07 struct {
	Name string
}

func (d Dog07) Speak() {
	fmt.Println("Dog:", d.Name)
}

// ============================================================
// Q08. Interface value stores dynamic concrete value.
// ============================================================
func q08() {
	separator(8)

	var s Speaker08 = Person08{Name: "Ashutosh"}

	fmt.Printf("Dynamic type: %T\n", s)
	fmt.Printf("Dynamic value: %v\n", s)
}

type Speaker08 interface {
	Speak()
}

type Person08 struct {
	Name string
}

func (p Person08) Speak() {
	fmt.Println(p.Name)
}

// ============================================================
// Q09. Use %T with interface.
// ============================================================
func q09() {
	separator(9)

	var value interface{} = 100

	fmt.Printf("Type: %T\n", value)
}

// ============================================================
// Q10. Empty interface.
// ============================================================
func q10() {
	separator(10)

	var value interface{}

	value = 100
	fmt.Println(value)

	value = "Go"
	fmt.Println(value)

	value = true
	fmt.Println(value)
}

// ============================================================
// Q11. any is alias for interface{}.
// ============================================================
func q11() {
	separator(11)

	var a any = 100
	var b interface{} = 100

	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)

	/*
		any is an alias for interface{}.

		These are equivalent:

		    any
		    interface{}
	*/
}

// ============================================================
// Q12. Interface containing an integer.
// ============================================================
func q12() {
	separator(12)

	var value any = 42

	fmt.Println(value)
}

// ============================================================
// Q13. Interface containing a slice.
// ============================================================
func q13() {
	separator(13)

	var value any = []int{10, 20, 30}

	fmt.Println(value)
}

// ============================================================
// Q14. Interface containing a map.
// ============================================================
func q14() {
	separator(14)

	var value any = map[string]int{
		"Go": 100,
	}

	fmt.Println(value)
}

// ============================================================
// Q15. Interface containing a struct.
// ============================================================
func q15() {
	separator(15)

	type User struct {
		Name string
		Age  int
	}

	var value any = User{
		Name: "Ashutosh",
		Age:  25,
	}

	fmt.Println(value)
}

// ============================================================
// Q16. Type assertion: successful.
// ============================================================
func q16() {
	separator(16)

	var value any = 100

	number := value.(int)

	fmt.Println("Number:", number)
}

// ============================================================
// Q17. Type assertion with comma-ok.
// ============================================================
func q17() {
	separator(17)

	var value any = 100

	number, ok := value.(int)

	fmt.Println("Number:", number)
	fmt.Println("Success:", ok)
}

// ============================================================
// Q18. Type assertion with wrong type.
// ============================================================
func q18() {
	separator(18)

	var value any = 100

	text, ok := value.(string)

	fmt.Println("Text:", text)
	fmt.Println("Success:", ok)

	/*
		Using:

		    value.(string)

		directly would panic because the actual type is int.

		Using comma-ok avoids the panic.
	*/
}

// ============================================================
// Q19. Type switch.
// ============================================================
func q19() {
	separator(19)

	printType19(100)
	printType19("Go")
	printType19(true)
	printType19([]int{1, 2, 3})
}

func printType19(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case []int:
		fmt.Println("[]int:", v)
	default:
		fmt.Println("unknown:", v)
	}
}

// ============================================================
// Q20. Interface with Stringer-like method.
// ============================================================
func q20() {
	separator(20)

	u := User20{
		Name: "Ashutosh",
	}

	fmt.Println(u)
}

type User20 struct {
	Name string
}

func (u User20) String() string {
	return "User: " + u.Name
}

// ============================================================
// Q21. Standard fmt.Stringer interface.
// ============================================================
func q21() {
	separator(21)

	var s fmt.Stringer = User21{
		Name: "Ashutosh",
	}

	fmt.Println(s.String())
}

type User21 struct {
	Name string
}

func (u User21) String() string {
	return "User: " + u.Name
}

// ============================================================
// Q22. Pointer receiver and interface.
// ============================================================
func q22() {
	separator(22)

	p := Person22{Name: "Ashutosh"}

	var s Speaker22 = &p

	s.Speak()
}

type Speaker22 interface {
	Speak()
}

type Person22 struct {
	Name string
}

func (p *Person22) Speak() {
	fmt.Println("Hello", p.Name)
}

// ============================================================
// Q23. Why Person may not implement interface when method
// has pointer receiver.
// ============================================================
func q23() {
	separator(23)

	/*
		Here Speak() has a pointer receiver:

		    func (p *Person23) Speak()

		So:

		    *Person23 implements Speaker23

		but:

		    Person23 does not implement Speaker23
	*/

	p := Person23{Name: "Ashutosh"}

	var s Speaker23 = &p

	s.Speak()
}

type Speaker23 interface {
	Speak()
}

type Person23 struct {
	Name string
}

func (p *Person23) Speak() {
	fmt.Println("Speaking:", p.Name)
}

// ============================================================
// Q24. Value receiver means both value and pointer can satisfy
// the interface.
// ============================================================
func q24() {
	separator(24)

	p := Person24{Name: "Ashutosh"}

	var a Speaker24 = p
	var b Speaker24 = &p

	a.Speak()
	b.Speak()
}

type Speaker24 interface {
	Speak()
}

type Person24 struct {
	Name string
}

func (p Person24) Speak() {
	fmt.Println("Speaking:", p.Name)
}

// ============================================================
// Q25. Interface embedding.
// ============================================================
func q25() {
	separator(25)

	var e Employee25 = Employee25{Name: "Ashutosh"}

	var company CompanyEmployee25 = e

	company.Work()
	company.Report()
}

type Worker25 interface {
	Work()
}

type Reporter25 interface {
	Report()
}

type CompanyEmployee25 interface {
	Worker25
	Reporter25
}

type Employee25 struct {
	Name string
}

func (e Employee25) Work() {
	fmt.Println(e.Name, "working")
}

func (e Employee25) Report() {
	fmt.Println(e.Name, "reporting")
}

// ============================================================
// Q26. Multiple interfaces implemented by one type.
// ============================================================
func q26() {
	separator(26)

	e := Employee26{Name: "Ashutosh"}

	var w Worker26 = e
	var r Reporter26 = e

	w.Work()
	r.Report()
}

type Worker26 interface {
	Work()
}

type Reporter26 interface {
	Report()
}

type Employee26 struct {
	Name string
}

func (e Employee26) Work() {
	fmt.Println("Work:", e.Name)
}

func (e Employee26) Report() {
	fmt.Println("Report:", e.Name)
}

// ============================================================
// Q27. Interface composition.
// ============================================================
func q27() {
	separator(27)

	type Reader27 interface {
		Read()
	}

	type Writer27 interface {
		Write()
	}

	type ReadWriter27 interface {
		Reader27
		Writer27
	}

	var rw ReadWriter27 = File27{}

	rw.Read()
	rw.Write()
}

type File27 struct{}

func (File27) Read() {
	fmt.Println("Reading")
}

func (File27) Write() {
	fmt.Println("Writing")
}

// ============================================================
// Q28. Interface used for polymorphism.
// ============================================================
func q28() {
	separator(28)

	vehicles := []Vehicle28{
		Car28{},
		Bike28{},
	}

	for _, v := range vehicles {
		v.Move()
	}
}

type Vehicle28 interface {
	Move()
}

type Car28 struct{}

func (Car28) Move() {
	fmt.Println("Car moving")
}

type Bike28 struct{}

func (Bike28) Move() {
	fmt.Println("Bike moving")
}

// ============================================================
// Q29. Interface slice.
// ============================================================
func q29() {
	separator(29)

	values := []Speaker29{
		Person29{Name: "Ashutosh"},
		Dog29{Name: "Tommy"},
	}

	for _, value := range values {
		value.Speak()
	}
}

type Speaker29 interface {
	Speak()
}

type Person29 struct {
	Name string
}

func (p Person29) Speak() {
	fmt.Println("Person:", p.Name)
}

type Dog29 struct {
	Name string
}

func (d Dog29) Speak() {
	fmt.Println("Dog:", d.Name)
}

// ============================================================
// Q30. Interface as return type.
// ============================================================
func q30() {
	separator(30)

	s := createSpeaker30("Ashutosh")

	s.Speak()
}

type Speaker30 interface {
	Speak()
}

type Person30 struct {
	Name string
}

func (p Person30) Speak() {
	fmt.Println("Hello", p.Name)
}

func createSpeaker30(name string) Speaker30 {
	return Person30{Name: name}
}

// ============================================================
// Q31. Interface nil.
// ============================================================
func q31() {
	separator(31)

	var s Speaker31

	fmt.Println("Is nil:", s == nil)
}

type Speaker31 interface {
	Speak()
}

// ============================================================
// Q32. Nil interface vs interface containing nil pointer.
// ============================================================
func q32() {
	separator(32)

	var p *Person32 = nil

	var s Speaker32 = p

	fmt.Println("Interface == nil:", s == nil)

	/*
		This is false.

		Why?

		An interface internally has:

		    dynamic type
		    dynamic value

		Here:

		    dynamic type  = *Person32
		    dynamic value = nil

		So the interface itself is not nil.
	*/
}

type Speaker32 interface {
	Speak()
}

type Person32 struct{}

func (p *Person32) Speak() {
	if p == nil {
		fmt.Println("nil receiver")
		return
	}

	fmt.Println("Person speaking")
}

// ============================================================
// Q33. Calling method on typed nil pointer in interface.
// ============================================================
func q33() {
	separator(33)

	var p *Person33 = nil
	var s Speaker33 = p

	s.Speak()
}

type Speaker33 interface {
	Speak()
}

type Person33 struct{}

func (p *Person33) Speak() {
	if p == nil {
		fmt.Println("Safe nil receiver")
		return
	}

	fmt.Println("Person speaking")
}

// ============================================================
// Q34. Interface type assertion to concrete struct.
// ============================================================
func q34() {
	separator(34)

	var s Speaker34 = Person34{Name: "Ashutosh"}

	p, ok := s.(Person34)

	fmt.Println("Value:", p)
	fmt.Println("Success:", ok)
}

type Speaker34 interface {
	Speak()
}

type Person34 struct {
	Name string
}

func (p Person34) Speak() {
	fmt.Println(p.Name)
}

// ============================================================
// Q35. Interface type assertion to pointer.
// ============================================================
func q35() {
	separator(35)

	person := &Person35{Name: "Ashutosh"}

	var s Speaker35 = person

	p, ok := s.(*Person35)

	fmt.Println("Pointer:", p)
	fmt.Println("Success:", ok)
}

type Speaker35 interface {
	Speak()
}

type Person35 struct {
	Name string
}

func (p *Person35) Speak() {
	fmt.Println(p.Name)
}

// ============================================================
// Q36. Type switch with pointer and value.
// ============================================================
func q36() {
	separator(36)

	printSpeakerType36(Person36{Name: "Ashutosh"})
	printSpeakerType36(&Person36{Name: "Ashutosh"})
}

func printSpeakerType36(value any) {
	switch v := value.(type) {
	case Person36:
		fmt.Println("Person value:", v.Name)
	case *Person36:
		fmt.Println("Person pointer:", v.Name)
	default:
		fmt.Println("Unknown")
	}
}

type Person36 struct {
	Name string
}

// ============================================================
// Q37. Interface for dependency injection.
// ============================================================
func q37() {
	separator(37)

	service := NewUserService37(MockRepository37{})

	service.GetUser()
}

type Repository37 interface {
	GetUser()
}

type MockRepository37 struct{}

func (MockRepository37) GetUser() {
	fmt.Println("Mock user returned")
}

type UserService37 struct {
	repo Repository37
}

func NewUserService37(repo Repository37) *UserService37 {
	return &UserService37{
		repo: repo,
	}
}

func (s *UserService37) GetUser() {
	s.repo.GetUser()
}

// ============================================================
// Q38. Real and mock implementations.
// ============================================================
func q38() {
	separator(38)

	var repo Repository38

	repo = RealRepository38{}
	repo.Get()

	repo = MockRepository38{}
	repo.Get()
}

type Repository38 interface {
	Get()
}

type RealRepository38 struct{}

func (RealRepository38) Get() {
	fmt.Println("Database data")
}

type MockRepository38 struct{}

func (MockRepository38) Get() {
	fmt.Println("Mock data")
}

// ============================================================
// Q39. Interface reduces coupling.
// ============================================================
func q39() {
	separator(39)

	notification := EmailNotification39{}

	sendNotification39(notification)
}

type Notifier39 interface {
	Send(message string)
}

type EmailNotification39 struct{}

func (EmailNotification39) Send(message string) {
	fmt.Println("Email:", message)
}

func sendNotification39(n Notifier39) {
	n.Send("Interview preparation")
}

// ============================================================
// Q40. Multiple implementations for same behavior.
// ============================================================
func q40() {
	separator(40)

	sendMessage40(Email40{})
	sendMessage40(SMS40{})
}

type Messenger40 interface {
	Send(string)
}

type Email40 struct{}

func (Email40) Send(message string) {
	fmt.Println("Email:", message)
}

type SMS40 struct{}

func (SMS40) Send(message string) {
	fmt.Println("SMS:", message)
}

func sendMessage40(m Messenger40) {
	m.Send("Hello")
}

// ============================================================
// Q41. Empty interface and type switch practical example.
// ============================================================
func q41() {
	separator(41)

	values := []any{
		100,
		"Go",
		3.14,
		true,
	}

	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case float64:
			fmt.Println("Float:", v)
		case bool:
			fmt.Println("Boolean:", v)
		}
	}
}

// ============================================================
// Q42. Interface can contain function? Yes, but functions do
// not implement arbitrary interfaces unless methods exist.
// ============================================================
func q42() {
	separator(42)

	/*
		A function value itself can be stored in an interface{} / any:

		    var x any = func() {}

		But a function type does not automatically satisfy an
		interface requiring methods.
	*/

	var value any = func() {
		fmt.Println("Function")
	}

	fmt.Printf("Dynamic type: %T\n", value)
}

// ============================================================
// Q43. Interface and error.
// ============================================================
func q43() {
	separator(43)

	err := divide43(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

type Error43 struct {
	Message string
}

func (e Error43) Error() string {
	return e.Message
}

func divide43(a, b int) error {
	if b == 0 {
		return Error43{
			Message: "cannot divide by zero",
		}
	}

	return nil
}

// ============================================================
// Q44. Custom error implements error interface.
// ============================================================
func q44() {
	separator(44)

	err := Error44{
		Code:    404,
		Message: "User not found",
	}

	var e error = err

	fmt.Println("Error:", e.Error())
}

type Error44 struct {
	Code    int
	Message string
}

func (e Error44) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// ============================================================
// Q45. Interface method set concept.
// ============================================================
func q45() {
	separator(45)

	/*
		For:

		    type Person45 struct{}

		    func (Person45) Speak() {}

		both:

		    Person45
		    *Person45

		can satisfy:

		    Speaker45

		because a value receiver method belongs to the method set
		of both the value and pointer.

		For a pointer receiver:

		    func (*Person45) Speak() {}

		only *Person45 satisfies Speaker45.
	*/

	var a Speaker45 = Person45{}
	var b Speaker45 = &Person45{}

	a.Speak()
	b.Speak()
}

type Speaker45 interface {
	Speak()
}

type Person45 struct{}

func (Person45) Speak() {
	fmt.Println("Speaking")
}

// ============================================================
// Q46. Compile-time interface assertion.
// ============================================================
func q46() {
	separator(46)

	/*
		This pattern verifies at compile time that a type implements
		an interface:

		    var _ Speaker46 = (*Person46)(nil)

		It creates no runtime work.
	*/

	fmt.Println("Person46 implements Speaker46")
}

type Speaker46 interface {
	Speak()
}

type Person46 struct{}

func (*Person46) Speak() {}

var _ Speaker46 = (*Person46)(nil)

// ============================================================
// Q47. Interface embedding with io-like behavior.
// ============================================================
func q47() {
	separator(47)

	type Reader47 interface {
		Read()
	}

	type Closer47 interface {
		Close()
	}

	type ReadCloser47 interface {
		Reader47
		Closer47
	}

	var rc ReadCloser47 = File47{}

	rc.Read()
	rc.Close()
}

type File47 struct{}

func (File47) Read() {
	fmt.Println("Read")
}

func (File47) Close() {
	fmt.Println("Close")
}

// ============================================================
// Q48. Interface segregation.
// ============================================================
func q48() {
	separator(48)

	/*
		Prefer small interfaces.

		Instead of:

		    type Huge interface {
		        Read()
		        Write()
		        Delete()
		        Update()
		    }

		split responsibilities when appropriate:

		    Reader
		    Writer
		    Deleter
		    Updater
	*/

	var r Reader48 = File48{}

	r.Read()
}

type Reader48 interface {
	Read()
}

type File48 struct{}

func (File48) Read() {
	fmt.Println("Reading")
}

func (File48) Write() {
	fmt.Println("Writing")
}

// ============================================================
// Q49. Accept interface, return concrete type principle.
// ============================================================
func q49() {
	separator(49)

	/*
		A common Go API design principle is:

		    Accept interfaces when abstraction is useful.
		    Return concrete types when callers benefit from the
		    concrete API.

		This is not an absolute rule, but it is a useful guideline.
	*/

	result := createUser49()

	fmt.Println(result.Name)
}

type User49 struct {
	Name string
}

func createUser49() User49 {
	return User49{Name: "Ashutosh"}
}

// ============================================================
// Q50. Interface and nil error trap.
// ============================================================
func q50() {
	separator(50)

	err := getError50()

	if err == nil {
		fmt.Println("No error")
	} else {
		fmt.Println("Error interface is non-nil")
	}
}

type CustomError50 struct{}

func (CustomError50) Error() string {
	return "custom error"
}

func getError50() error {
	var err *CustomError50 = nil

	/*
		IMPORTANT:

		    return err

		returns an error interface containing:

		    dynamic type  = *CustomError50
		    dynamic value = nil

		Therefore the returned interface is NOT nil.
	*/

	return err
}

// ============================================================
// Q51. Correct way to return nil error.
// ============================================================
func q51() {
	separator(51)

	err := getError51()

	if err == nil {
		fmt.Println("Correctly nil")
	}
}

func getError51() error {
	return nil
}

// ============================================================
// Q52. Type assertion on interface containing pointer.
// ============================================================
func q52() {
	separator(52)

	var value any = &Person52{Name: "Ashutosh"}

	p, ok := value.(*Person52)

	if ok {
		fmt.Println("Name:", p.Name)
	}
}

type Person52 struct {
	Name string
}

// ============================================================
// Q53. Type assertion helper function.
// ============================================================
func q53() {
	separator(53)

	printInt53(100)
	printInt53("Go")
}

func printInt53(value any) {
	number, ok := value.(int)

	if ok {
		fmt.Println("Integer:", number)
		return
	}

	fmt.Println("Not an integer")
}

// ============================================================
// Q54. Type switch helper for API-like data.
// ============================================================
func q54() {
	separator(54)

	printValue54(100)
	printValue54("Ashutosh")
	printValue54(nil)
}

func printValue54(value any) {
	switch v := value.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("other type: %T\n", v)
	}
}

// ============================================================
// Q55. Practical repository interface.
// ============================================================
func q55() {
	separator(55)

	repo := NewRepository55()
	service := NewService55(repo)

	user := service.GetUser(101)

	fmt.Println("User:", user)
}

type User55 struct {
	ID   int
	Name string
}

type UserRepository55 interface {
	FindByID(id int) User55
}

type RepositoryImpl55 struct{}

func NewRepository55() UserRepository55 {
	return RepositoryImpl55{}
}

func (RepositoryImpl55) FindByID(id int) User55 {
	return User55{
		ID:   id,
		Name: "Ashutosh",
	}
}

type UserService55 struct {
	repo UserRepository55
}

func NewService55(repo UserRepository55) *UserService55 {
	return &UserService55{
		repo: repo,
	}
}

func (s *UserService55) GetUser(id int) User55 {
	return s.repo.FindByID(id)
}

// ============================================================
// Q56. Practical notification interface.
// ============================================================
func q56() {
	separator(56)

	processNotification56(Email56{})
	processNotification56(SMS56{})
	processNotification56(Push56{})
}

type Notification56 interface {
	Send(message string) error
}

type Email56 struct{}

func (Email56) Send(message string) error {
	fmt.Println("Email:", message)
	return nil
}

type SMS56 struct{}

func (SMS56) Send(message string) error {
	fmt.Println("SMS:", message)
	return nil
}

type Push56 struct{}

func (Push56) Send(message string) error {
	fmt.Println("Push:", message)
	return nil
}

func processNotification56(n Notification56) {
	if err := n.Send("Invoice approved"); err != nil {
		fmt.Println("Error:", err)
	}
}

// ============================================================
// Q57. Practical payment interface.
// ============================================================
func q57() {
	separator(57)

	processPayment57(CreditCard57{})
	processPayment57(UPI57{})
}

type Payment57 interface {
	Pay(amount float64) bool
}

type CreditCard57 struct{}

func (CreditCard57) Pay(amount float64) bool {
	fmt.Println("Credit card payment:", amount)
	return true
}

type UPI57 struct{}

func (UPI57) Pay(amount float64) bool {
	fmt.Println("UPI payment:", amount)
	return true
}

func processPayment57(p Payment57) {
	success := p.Pay(1500)

	fmt.Println("Success:", success)
}

// ============================================================
// Q58. Practical logger interface.
// ============================================================
func q58() {
	separator(58)

	service := Service58{
		Logger: ConsoleLogger58{},
	}

	service.Run()
}

type Logger58 interface {
	Info(message string)
	Error(message string)
}

type ConsoleLogger58 struct{}

func (ConsoleLogger58) Info(message string) {
	fmt.Println("INFO:", message)
}

func (ConsoleLogger58) Error(message string) {
	fmt.Println("ERROR:", message)
}

type Service58 struct {
	Logger Logger58
}

func (s Service58) Run() {
	s.Logger.Info("service started")
	s.Logger.Error("example error")
}

// ============================================================
// Q59. Interface-based sorting-style abstraction.
// ============================================================
func q59() {
	separator(59)

	/*
		Go's sort package uses interfaces such as sort.Interface
		to define behavior.

		The important design idea:

		    Len()
		    Less()
		    Swap()

		A type implements the required behavior and can be passed
		to a generic algorithm.

		This example demonstrates the same concept manually.
	*/

	values := IntCollection59{3, 1, 2}

	sortAscending59(&values)

	fmt.Println("Sorted:", values.Values)
}

type IntCollection59 struct {
	Values []int
}

func (c *IntCollection59) Len() int {
	return len(c.Values)
}

func (c *IntCollection59) Less(i, j int) bool {
	return c.Values[i] < c.Values[j]
}

func (c *IntCollection59) Swap(i, j int) {
	c.Values[i], c.Values[j] = c.Values[j], c.Values[i]
}

type Sorter59 interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func sortAscending59(s Sorter59) {
	for i := 0; i < s.Len(); i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, i) {
				s.Swap(i, j)
			}
		}
	}
}

// ============================================================
// Q60. Final practical interview problem:
//      Build a service using interface + mock implementation.
// ============================================================
func q60() {
	separator(60)

	/*
		This pattern is extremely useful in real Go backend projects.

		Architecture:

		    Handler/Service
		          |
		          v
		    UserRepository interface
		          |
		          +---- PostgreSQL implementation
		          |
		          +---- Mock implementation for tests

		The service does not need to know the concrete database type.
	*/

	repo := MockUserRepository60{
		User: User60{
			ID:   101,
			Name: "Ashutosh",
		},
	}

	service := UserService60{
		Repo: repo,
	}

	user, err := service.GetUser(101)

	fmt.Println("User:", user)
	fmt.Println("Error:", err)
}

type User60 struct {
	ID   int
	Name string
}

type UserRepository60 interface {
	FindByID(id int) (User60, error)
}

type MockUserRepository60 struct {
	User User60
}

func (m MockUserRepository60) FindByID(id int) (User60, error) {
	if m.User.ID != id {
		return User60{}, fmt.Errorf("user not found")
	}

	return m.User, nil
}

type UserService60 struct {
	Repo UserRepository60
}

func (s UserService60) GetUser(id int) (User60, error) {
	return s.Repo.FindByID(id)
}

/*
=====================================================================
DEEP INTERFACE INTERVIEW NOTES
=====================================================================

1. WHAT IS AN INTERFACE?

An interface defines behavior through methods.

Example:

    type Speaker interface {
        Speak()
    }

Any type that has:

    Speak()

satisfies the interface.

---------------------------------------------------------------------

2. DOES GO HAVE EXPLICIT IMPLEMENTATION?

No.

There is no:

    implements Speaker

keyword.

Implementation is implicit.

---------------------------------------------------------------------

3. WHY ARE INTERFACES USED?

Main reasons:

- Abstraction
- Polymorphism
- Loose coupling
- Dependency injection
- Testing / mocking
- Replacing implementations
- Designing small APIs

---------------------------------------------------------------------

4. INTERFACE VALUE

An interface value conceptually contains:

    dynamic type
    dynamic value

Example:

    var x any = 100

Conceptually:

    dynamic type  = int
    dynamic value = 100

---------------------------------------------------------------------

5. %T

Very important for interviews.

    fmt.Printf("%T\n", value)

prints the dynamic type of the value.

Example:

    var x any = 100

    fmt.Printf("%T\n", x)

Output:

    int

For:

    var x any = "Go"

Output:

    string

---------------------------------------------------------------------

6. %v

    %v

prints the value in its default format.

Example:

    fmt.Printf("%v\n", 100)

Output:

    100

---------------------------------------------------------------------

7. %T VS %v

    %T -> type
    %v -> value

Example:

    x := 100

    fmt.Printf("%T %v\n", x, x)

Output:

    int 100

---------------------------------------------------------------------

8. EMPTY INTERFACE

Before Go 1.18:

    interface{}

Modern Go:

    any

They are aliases.

Example:

    var value any

It can hold values of different concrete types.

---------------------------------------------------------------------

9. TYPE ASSERTION

Syntax:

    value.(Type)

Example:

    var x any = 100

    n := x.(int)

This extracts the int.

If the actual dynamic type is not int, direct assertion panics.

---------------------------------------------------------------------

10. COMMA-OK TYPE ASSERTION

Safer:

    n, ok := x.(int)

If successful:

    ok == true

If unsuccessful:

    ok == false

No panic.

---------------------------------------------------------------------

11. TYPE SWITCH

Syntax:

    switch v := value.(type) {
    case int:
        ...
    case string:
        ...
    }

Used to handle multiple dynamic types.

---------------------------------------------------------------------

12. NIL INTERFACE

This is nil:

    var x Speaker
    x == nil

Because both interface components are absent.

---------------------------------------------------------------------

13. TYPED NIL INSIDE INTERFACE

This is NOT nil:

    var p *Person = nil
    var x Speaker = p

Because:

    dynamic type  = *Person
    dynamic value = nil

Therefore:

    x != nil

This is one of the most important Go interface interview traps.

---------------------------------------------------------------------

14. INTERFACE METHOD SET

Suppose:

    type Speaker interface {
        Speak()
    }

If:

    func (p Person) Speak() {}

then both:

    Person
    *Person

can satisfy Speaker.

If:

    func (p *Person) Speak() {}

then only:

    *Person

satisfies Speaker.

---------------------------------------------------------------------

15. VALUE RECEIVER

Example:

    func (p Person) Speak() {}

Method belongs to the value method set and is also available
through a pointer.

---------------------------------------------------------------------

16. POINTER RECEIVER

Example:

    func (p *Person) Speak() {}

Only *Person satisfies an interface requiring Speak.

---------------------------------------------------------------------

17. INTERFACE EMBEDDING

You can compose interfaces:

    type Reader interface {
        Read()
    }

    type Writer interface {
        Write()
    }

    type ReadWriter interface {
        Reader
        Writer
    }

A type must implement both methods to satisfy ReadWriter.

---------------------------------------------------------------------

18. INTERFACE SEGREGATION

Prefer small interfaces.

Instead of a giant interface:

    type Database interface {
        Create()
        Read()
        Update()
        Delete()
        Backup()
        Restore()
        ...
    }

Use smaller interfaces where useful:

    Reader
    Writer
    Deleter

Small interfaces are easier to implement, test and reuse.

---------------------------------------------------------------------

19. io.Reader IS A FAMOUS EXAMPLE

Go standard library uses small interfaces extensively.

Conceptually:

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

Many different types can implement this behavior.

---------------------------------------------------------------------

20. ERROR IS AN INTERFACE

The built-in error type is an interface conceptually equivalent to:

    type error interface {
        Error() string
    }

Any type implementing:

    Error() string

can be used as an error.

---------------------------------------------------------------------

21. CUSTOM ERROR

Example:

    type MyError struct {
        Message string
    }

    func (e MyError) Error() string {
        return e.Message
    }

Now MyError implements error.

---------------------------------------------------------------------

22. NIL ERROR TRAP

This is dangerous:

    func getError() error {
        var err *MyError = nil
        return err
    }

The returned error interface is not nil.

Why?

    dynamic type  = *MyError
    dynamic value = nil

Correct when there is no error:

    return nil

---------------------------------------------------------------------

23. DEPENDENCY INJECTION

Instead of hard-coding:

    service -> PostgreSQL

use:

    service -> Repository interface

Then:

    PostgreSQLRepository implements Repository
    MockRepository implements Repository

This makes testing easier.

---------------------------------------------------------------------

24. MOCKING

Suppose:

    type Repository interface {
        FindByID(id int) User
    }

Production:

    PostgreSQLRepository

Testing:

    MockRepository

The service can use either one.

---------------------------------------------------------------------

25. POLYMORPHISM

A function accepting an interface can work with many types.

Example:

    func send(n Notifier) {
        n.Send()
    }

Then:

    Email
    SMS
    Push

can all be passed to send() if they implement Notifier.

---------------------------------------------------------------------

26. INTERFACE AS RETURN TYPE

Possible:

    func create() Speaker {
        return Person{}
    }

This hides the concrete implementation.

Use it when returning the abstraction is beneficial.

---------------------------------------------------------------------

27. ACCEPT INTERFACE, RETURN CONCRETE

A common Go design guideline:

    Accept interfaces where abstraction is useful.

    Return concrete types when callers benefit from the
    concrete API.

It is a guideline, not an absolute language rule.

---------------------------------------------------------------------

28. COMPILE-TIME INTERFACE CHECK

Very useful:

    var _ Speaker = (*Person)(nil)

This ensures the compiler verifies that *Person implements Speaker.

If Speak() is removed or changed incorrectly, compilation fails.

---------------------------------------------------------------------

29. INTERFACE VS STRUCT

Struct:

    describes data.

Interface:

    describes behavior.

Example:

    type Employee struct {
        Name string
    }

    type Worker interface {
        Work()
    }

Struct stores state.

Interface defines required behavior.

---------------------------------------------------------------------

30. INTERFACE VS POINTER

Pointer:

    stores an address/reference to a value.

Interface:

    represents a set of required methods and can hold a
    concrete value.

They solve different problems.

---------------------------------------------------------------------

31. INTERFACE VS ANY

any is an alias for interface{}.

An ordinary interface:

    type Reader interface {
        Read()
    }

restricts values to types implementing Read().

any:

    var x any

can hold any value.

---------------------------------------------------------------------

32. WHEN NOT TO USE INTERFACE

Do not create interfaces automatically for every struct.

If there is only one concrete implementation and no useful
abstraction/testing boundary, an interface may add unnecessary
complexity.

Prefer simple concrete code when abstraction is not needed.

---------------------------------------------------------------------

33. INTERFACE DESIGN RULE

Define interfaces near the code that USES the behavior, rather
than automatically defining a huge interface next to every
implementation.

This encourages small, focused interfaces.

---------------------------------------------------------------------

34. COMMON INTERVIEW TRAPS

TRAP 1:

    Go uses explicit implements.

WRONG.

Go uses implicit interface satisfaction.

TRAP 2:

    interface == nil whenever the underlying pointer is nil.

WRONG.

A typed nil pointer stored in an interface makes the interface
non-nil.

TRAP 3:

    any and interface{} are different.

WRONG.

They are aliases.

TRAP 4:

    Type assertion and type conversion are the same.

WRONG.

Assertion extracts a dynamic value from an interface.

Conversion changes a value from one type to another when
conversion is allowed.

---------------------------------------------------------------------

35. TYPE ASSERTION VS TYPE CONVERSION

Assertion:

    x.(int)

requires x to be an interface value.

Conversion:

    int64(x)

changes one concrete type to another when allowed.

---------------------------------------------------------------------

36. INTERFACE METHOD SET QUESTION

Question:

    type Person struct{}

    func (p *Person) Speak() {}

Does Person implement:

    type Speaker interface {
        Speak()
    }

Answer:

No.

*Person implements Speaker.

---------------------------------------------------------------------

37. WHY USE POINTER RECEIVER WITH INTERFACE?

Because the method may need to modify the receiver:

    func (p *Person) ChangeName(...) {}

It also avoids copying large structs.

---------------------------------------------------------------------

38. INTERFACE + JSON

JSON data is commonly decoded into structs, but dynamic JSON
data can also be decoded into:

    map[string]any

Use structs when the schema is known.

Use any/map structures when data is genuinely dynamic.

---------------------------------------------------------------------

39. INTERFACE + DATABASE

A common backend architecture:

    Handler
       |
       v
    Service
       |
       v
    Repository interface
       |
       +------ PostgreSQL
       |
       +------ Mock

This is useful for testing and loose coupling.

---------------------------------------------------------------------

40. INTERFACE + NOTIFICATION SYSTEM

A real-world design:

    type Notifier interface {
        Send(message string) error
    }

Implement:

    EmailNotifier
    SMSNotifier
    WhatsAppNotifier
    PushNotifier

Then:

    func Notify(n Notifier) error

can work with all implementations.

---------------------------------------------------------------------

41. INTERFACE + PAYMENT SYSTEM

Another practical example:

    type Payment interface {
        Pay(amount float64) error
    }

Implement:

    CreditCard
    UPI
    PayPal
    Stripe

Business logic depends on Payment rather than one specific
payment provider.

---------------------------------------------------------------------

42. INTERFACE + LOW-CODE PLATFORM

For your low-code platform, interfaces can be useful for
pluggable components such as:

    ResourceExecutor
    NotificationProvider
    WorkflowNode
    StorageProvider
    RuleEvaluator

For example:

    type NotificationProvider interface {
        Send(to string, message string) error
    }

Then email/SMS/WhatsApp providers can implement it.

---------------------------------------------------------------------

43. INTERFACE + WORKFLOW NODE

Example design:

    type WorkflowNode interface {
        Execute(ctx Context) error
    }

Different nodes:

    ApprovalNode
    ConditionNode
    EmailNode
    WebhookNode

can implement Execute().

The workflow engine only needs WorkflowNode.

---------------------------------------------------------------------

44. INTERFACE + RESOURCE BUILDER

A resource executor could be:

    type ResourceExecutor interface {
        Execute(query string, params map[string]any) (any, error)
    }

Implementations could be:

    PostgreSQLExecutor
    MySQLExecutor
    RESTExecutor

The platform can work with the abstraction.

---------------------------------------------------------------------

45. INTERFACE + TESTING

Production:

    type PostgresRepository struct{}

Testing:

    type MockRepository struct{}

Both implement:

    UserRepository

Your service can be tested without a real database.

---------------------------------------------------------------------

46. INTERFACE EMBEDDING VS STRUCT EMBEDDING

Struct embedding:

    type Employee struct {
        Person
    }

Interface embedding:

    type ReadWriter interface {
        Reader
        Writer
    }

Both provide composition, but they operate on different concepts.

---------------------------------------------------------------------

47. INTERFACE VALUE COPYING

When an interface value is assigned:

    b := a

the interface value is copied.

The underlying concrete value may itself be a value or may contain
references/pointers.

Do not automatically assume deep copying.

---------------------------------------------------------------------

48. INTERFACE AND CONCURRENCY

Interfaces do not automatically make code thread-safe.

If multiple goroutines access shared concrete state stored behind
an interface, synchronization may still be required.

Use:

    mutex
    channels
    atomic operations

as appropriate.

---------------------------------------------------------------------

49. TOP 20 INTERFACE QUESTIONS TO MASTER

1. What is an interface?
2. Why use interfaces?
3. How does Go implement interfaces?
4. What does implicit implementation mean?
5. What is an empty interface?
6. What is any?
7. What is type assertion?
8. What is comma-ok?
9. What is type switch?
10. What is a nil interface?
11. What is typed nil?
12. Value receiver vs pointer receiver?
13. What is method set?
14. What is interface embedding?
15. What is polymorphism?
16. How are interfaces used for dependency injection?
17. How are interfaces used for mocking?
18. What is the error interface?
19. Why can a typed nil error be non-nil?
20. What is compile-time interface verification?

---------------------------------------------------------------------
50. INTERVIEW CODE YOU SHOULD WRITE WITHOUT LOOKING
---------------------------------------------------------------------

A)

    type Speaker interface {
        Speak()
    }

B)

    func (p Person) Speak() {}

C)

    var s Speaker = Person{}

D)

    value, ok := x.(int)

E)

    switch v := x.(type) {
    case int:
        ...
    case string:
        ...
    }

F)

    var _ Speaker = (*Person)(nil)

G)

    type ReadWriter interface {
        Reader
        Writer
    }

H)

    type Repository interface {
        FindByID(id int) User
    }

I)

    type MockRepository struct{}

    func (MockRepository) FindByID(id int) User {
        ...
    }

J)

    func Service(repo Repository) {
        ...
    }

=====================================================================
IMPORTANT FOR YOUR 2-YEAR GO INTERVIEW
=====================================================================

Master these first:

Q02-Q06    Basic interfaces
Q08-Q11    Interface values, %T, any
Q16-Q19    Type assertion and type switch
Q21-Q25    Pointer receiver + method sets
Q25-Q30    Interface composition and polymorphism
Q31-Q36    NIL and type assertions
Q37-Q40    Dependency injection and loose coupling
Q43-Q46    error, method sets, compile-time checks
Q50-Q55    nil error and practical repositories
Q56-Q60    Real-world backend interface design

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
	fmt.Println("ALL 60 INTERFACE QUESTIONS COMPLETED")
	fmt.Println("============================================================")
}
