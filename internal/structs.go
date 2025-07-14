package internal

import (
	"fmt"
	"math"
	"strings"
)

// Basic struct definition
type PersonStruct struct {
	Name    string
	Age     int
	Email   string
	Address AddressStruct // embedded struct
}

// Nested struct
type AddressStruct struct {
	Street  string
	City    string
	Country string
}

// Struct with methods
type RectangleStruct struct {
	Width  float64
	Height float64
}

// Value receiver method
func (r RectangleStruct) Area() float64 {
	return r.Width * r.Height
}

// Pointer receiver method
func (r *RectangleStruct) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Method that returns multiple values
func (r RectangleStruct) Dimensions() (float64, float64) {
	return r.Width, r.Height
}

// Struct for demonstrating different method types
type EmployeeStruct struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// Method with pointer receiver for modifying data
func (e *EmployeeStruct) GiveRaise(percentage float64) {
	e.Salary *= (1 + percentage/100)
}

// Method with value receiver for reading data
func (e EmployeeStruct) GetInfo() string {
	return fmt.Sprintf("ID: %d, Name: %s, Position: %s, Salary: %.2f",
		e.ID, e.Name, e.Position, e.Salary)
}

// Method that validates data
func (e EmployeeStruct) IsValid() bool {
	return e.ID > 0 && e.Name != "" && e.Salary > 0
}

// Struct with embedded type
type ManagerStruct struct {
	EmployeeStruct // embedded struct
	Department     string
	Subordinates   []string
}

// Method for embedded struct
func (m ManagerStruct) GetFullInfo() string {
	return fmt.Sprintf("%s, Department: %s, Manages: %d people",
		m.GetInfo(), m.Department, len(m.Subordinates))
}

// Interface demonstration
type ShapeInterface interface {
	Area() float64
	Perimeter() float64
}

// Circle struct implementing Shape interface
type CircleStruct struct {
	Radius float64
}

func (c CircleStruct) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c CircleStruct) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implementing Shape interface
func (r RectangleStruct) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function that works with interface
func PrintShapeInfo(s ShapeInterface) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Advanced struct with constructor pattern
type BankAccount struct {
	accountNumber string
	balance       float64
	owner         string
}

// Constructor function
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		accountNumber: generateAccountNumber(),
		balance:       initialBalance,
		owner:         owner,
	}
}

func generateAccountNumber() string {
	return "ACC-" + fmt.Sprintf("%06d", 123456) // simplified
}

// Methods for BankAccount
func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	ba.balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > ba.balance {
		return fmt.Errorf("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

func (ba BankAccount) GetBalance() float64 {
	return ba.balance
}

func (ba BankAccount) GetOwner() string {
	return ba.owner
}

// Main function that demonstrates all concepts
func RunStructureExamples() {
	basicStructureExample()
	structMethodsExampleDemo()
	pointerReceiverExampleDemo()
	valueReceiverExampleDemo()
	embeddedStructExampleDemo()
	interfaceExampleDemo()
	constructorPatternExampleDemo()
	anonymousStructExampleDemo()
}

func basicStructureExample() {
	fmt.Println("=== Basic Struct Example ===")

	// Creating struct instances
	person1 := PersonStruct{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: AddressStruct{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	// Another way to create struct
	person2 := PersonStruct{}
	person2.Name = "Jane Smith"
	person2.Age = 25
	person2.Email = "jane@example.com"

	// Using struct literals
	person3 := &PersonStruct{
		Name:  "Bob Johnson",
		Age:   35,
		Email: "bob@example.com",
	}

	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n", person2)
	fmt.Printf("Person 3: %+v\n", *person3)

	// Accessing nested struct fields
	fmt.Printf("Person 1 lives in: %s, %s\n", person1.Address.City, person1.Address.Country)
	fmt.Println()
}

func structMethodsExampleDemo() {
	fmt.Println("=== Struct Methods Example ===")

	rect := RectangleStruct{Width: 5, Height: 3}

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	width, height := rect.Dimensions()
	fmt.Printf("Dimensions: %.2f x %.2f\n", width, height)
	fmt.Println()
}

func pointerReceiverExampleDemo() {
	fmt.Println("=== Pointer Receiver Example ===")

	rect := RectangleStruct{Width: 4, Height: 6}
	fmt.Printf("Before scaling: %+v\n", rect)

	// Using pointer receiver method
	rect.Scale(2.0)
	fmt.Printf("After scaling by 2: %+v\n", rect)

	// Working with Employee
	emp := EmployeeStruct{
		ID:       1,
		Name:     "Alice",
		Position: "Developer",
		Salary:   50000,
	}

	fmt.Printf("Before raise: %s\n", emp.GetInfo())
	emp.GiveRaise(10) // 10% raise
	fmt.Printf("After 10%% raise: %s\n", emp.GetInfo())
	fmt.Println()
}

func valueReceiverExampleDemo() {
	fmt.Println("=== Value Receiver Example ===")

	emp := EmployeeStruct{
		ID:       2,
		Name:     "Charlie",
		Position: "Designer",
		Salary:   45000,
	}

	// Value receiver doesn't modify original
	info := emp.GetInfo()
	fmt.Printf("Employee Info: %s\n", info)

	// Validation method
	if emp.IsValid() {
		fmt.Println("Employee data is valid")
	} else {
		fmt.Println("Employee data is invalid")
	}

	// Testing with invalid employee
	invalidEmp := EmployeeStruct{ID: 0, Name: "", Salary: -1000}
	if !invalidEmp.IsValid() {
		fmt.Println("Invalid employee detected")
	}
	fmt.Println()
}

func embeddedStructExampleDemo() {
	fmt.Println("=== Embedded Struct Example ===")

	manager := ManagerStruct{
		EmployeeStruct: EmployeeStruct{
			ID:       3,
			Name:     "David",
			Position: "Manager",
			Salary:   75000,
		},
		Department:   "Engineering",
		Subordinates: []string{"Alice", "Bob", "Charlie"},
	}

	// Accessing embedded struct methods
	fmt.Printf("Manager Info: %s\n", manager.GetFullInfo())

	// Can also access embedded struct fields directly
	fmt.Printf("Manager Name: %s\n", manager.Name)
	fmt.Printf("Manager Salary: %.2f\n", manager.Salary)

	// Give manager a raise
	manager.GiveRaise(15)
	fmt.Printf("After raise: %s\n", manager.GetFullInfo())
	fmt.Println()
}

func interfaceExampleDemo() {
	fmt.Println("=== Interface Example ===")

	circle := CircleStruct{Radius: 5}
	rectangle := RectangleStruct{Width: 4, Height: 6}

	// Both structs implement Shape interface
	fmt.Print("Circle - ")
	PrintShapeInfo(circle)

	fmt.Print("Rectangle - ")
	PrintShapeInfo(rectangle)

	// Using slice of interfaces
	shapes := []ShapeInterface{circle, rectangle}

	fmt.Println("All shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d - ", i+1)
		PrintShapeInfo(shape)
	}
	fmt.Println()
}

func constructorPatternExampleDemo() {
	fmt.Println("=== Constructor Pattern Example ===")

	// Using constructor function
	account1 := NewBankAccount("John Doe", 1000.0)
	account2 := NewBankAccount("Jane Smith", 2000.0)

	fmt.Printf("Account 1 - Owner: %s, Balance: %.2f\n",
		account1.GetOwner(), account1.GetBalance())

	// Deposit money
	err := account1.Deposit(500)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After deposit: %.2f\n", account1.GetBalance())
	}

	// Withdraw money
	err = account1.Withdraw(300)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After withdrawal: %.2f\n", account1.GetBalance())
	}

	// Try to withdraw more than balance
	err = account1.Withdraw(2000)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Account 2 - Owner: %s, Balance: %.2f\n",
		account2.GetOwner(), account2.GetBalance())
	fmt.Println()
}

func anonymousStructExampleDemo() {
	fmt.Println("=== Anonymous Struct Example ===")

	// Anonymous struct - useful for temporary data structures
	config := struct {
		Host     string
		Port     int
		Database string
		SSL      bool
	}{
		Host:     "localhost",
		Port:     5432,
		Database: "myapp",
		SSL:      true,
	}

	fmt.Printf("Config: %+v\n", config)

	// Anonymous struct in slice
	users := []struct {
		Name string
		Role string
	}{
		{"Admin", "administrator"},
		{"User1", "regular"},
		{"User2", "regular"},
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("- %s (%s)\n", user.Name, user.Role)
	}

	// Anonymous struct with methods (rare but possible)
	calculator := struct {
		value    float64
		add      func(float64) float64
		multiply func(float64) float64
	}{
		value: 10,
	}

	calculator.add = func(x float64) float64 {
		calculator.value += x
		return calculator.value
	}

	calculator.multiply = func(x float64) float64 {
		calculator.value *= x
		return calculator.value
	}

	fmt.Printf("Initial value: %.2f\n", calculator.value)
	fmt.Printf("After adding 5: %.2f\n", calculator.add(5))
	fmt.Printf("After multiplying by 2: %.2f\n", calculator.multiply(2))
	fmt.Println()
}

// Bonus: Struct tags example (commonly used with JSON)
type ProductStruct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	IsAvailable bool    `json:"is_available"`
}

func (p ProductStruct) String() string {
	availability := "Available"
	if !p.IsAvailable {
		availability = "Not Available"
	}
	return fmt.Sprintf("%s (ID: %d) - $%.2f [%s] - %s",
		p.Name, p.ID, p.Price, p.Category, availability)
}

// Method to format product name
func (p ProductStruct) FormattedName() string {
	return strings.ToUpper(p.Name)
}

func structTagsExampleDemo() {
	fmt.Println("=== Struct Tags Example ===")

	product := ProductStruct{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Category:    "Electronics",
		IsAvailable: true,
	}

	fmt.Printf("Product: %s\n", product.String())
	fmt.Printf("Formatted Name: %s\n", product.FormattedName())
	fmt.Println()
}
