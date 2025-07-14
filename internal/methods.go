// methods.go
package internal

import (
	"fmt"
	"math"
)

// RunMethodExamples - main function to run all method examples
func RunMethodExamples() {
	basicMethodExample()
	pointerReceiverExample()
	valueReceiverExample()
	methodSetsExample()
	embeddedMethodExample()
	methodExpressionExample()
}

// Example 1: Basic methods
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

func basicMethodExample() {
	fmt.Println("\n=== Basic Method Example ===")

	rect := Rectangle{Width: 5.0, Height: 3.0}

	fmt.Printf("Rectangle: %s\n", rect.String())
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}

// Example 2: Pointer receiver vs Value receiver
type Counter struct {
	Count int
}

// Pointer receiver - modifies the original
func (c *Counter) Increment() {
	c.Count++
}

// Value receiver - works with a copy
func (c Counter) GetCount() int {
	return c.Count
}

// This won't work as expected (doesn't modify original)
func (c Counter) BrokenIncrement() {
	c.Count++
}

func pointerReceiverExample() {
	fmt.Println("\n=== Pointer Receiver Example ===")

	counter := Counter{Count: 0}

	fmt.Printf("Initial count: %d\n", counter.GetCount())

	counter.Increment()
	fmt.Printf("After increment: %d\n", counter.GetCount())

	counter.BrokenIncrement()
	fmt.Printf("After broken increment: %d\n", counter.GetCount())

	// Working with pointer
	counterPtr := &Counter{Count: 10}
	counterPtr.Increment()
	fmt.Printf("Pointer counter after increment: %d\n", counterPtr.GetCount())
}

// Example 3: Value receiver for performance
type BigStruct struct {
	Data [1000]int
}

// Value receiver - copies the entire struct
func (b BigStruct) ProcessData() int {
	sum := 0
	for _, v := range b.Data {
		sum += v
	}
	return sum
}

// Pointer receiver - more efficient for large structs
func (b *BigStruct) ProcessDataEfficient() int {
	sum := 0
	for _, v := range b.Data {
		sum += v
	}
	return sum
}

func valueReceiverExample() {
	fmt.Println("\n=== Value Receiver Example ===")

	bigStruct := BigStruct{}
	for i := 0; i < 1000; i++ {
		bigStruct.Data[i] = i + 1
	}

	// Both work, but pointer receiver is more efficient
	fmt.Printf("Sum (value receiver): %d\n", bigStruct.ProcessData())
	fmt.Printf("Sum (pointer receiver): %d\n", bigStruct.ProcessDataEfficient())
}

// Example 4: Method sets
type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", m)
}

func (m *MyInt) Double() {
	*m = *m * 2
}

func methodSetsExample() {
	fmt.Println("\n=== Method Sets Example ===")

	var num MyInt = 5
	fmt.Printf("Original: %s\n", num.String())

	num.Double()
	fmt.Printf("After doubling: %s\n", num.String())

	// Pointer to MyInt
	numPtr := &num
	fmt.Printf("Through pointer: %s\n", numPtr.String())
	numPtr.Double()
	fmt.Printf("After doubling through pointer: %s\n", numPtr.String())
}

// Example 5: Embedded methods
type Engine struct {
	Power int
}

func (e Engine) Start() {
	fmt.Printf("Engine with %d HP started\n", e.Power)
}

func (e Engine) Stop() {
	fmt.Println("Engine stopped")
}

type Car struct {
	Brand  string
	Model  string
	Engine // Embedded struct
}

func (c Car) Drive() {
	fmt.Printf("Driving %s %s\n", c.Brand, c.Model)
}

// Method can be overridden
func (c Car) Start() {
	fmt.Printf("Starting %s %s\n", c.Brand, c.Model)
	c.Engine.Start() // Call embedded method explicitly
}

func embeddedMethodExample() {
	fmt.Println("\n=== Embedded Method Example ===")

	car := Car{
		Brand:  "Toyota",
		Model:  "Camry",
		Engine: Engine{Power: 200},
	}

	car.Start() // Uses Car's Start method
	car.Drive() // Car's own method
	car.Stop()  // Engine's method (promoted)
}

// Example 6: Method expressions and method values
type Calculator struct {
	Result float64
}

func (c *Calculator) Add(x float64) {
	c.Result += x
}

func (c *Calculator) Multiply(x float64) {
	c.Result *= x
}

func (c Calculator) GetResult() float64 {
	return c.Result
}

func methodExpressionExample() {
	fmt.Println("\n=== Method Expression Example ===")

	calc := Calculator{Result: 10}

	// Method values (bound to instance)
	addMethod := calc.Add
	multiplyMethod := calc.Multiply

	addMethod(5)
	fmt.Printf("After adding 5: %.2f\n", calc.GetResult())

	multiplyMethod(2)
	fmt.Printf("After multiplying by 2: %.2f\n", calc.GetResult())

	// Method expressions (unbound)
	addExpression := (*Calculator).Add
	multiplyExpression := (*Calculator).Multiply

	calc2 := Calculator{Result: 1}
	addExpression(&calc2, 9)
	fmt.Printf("Second calculator after adding 9: %.2f\n", calc2.GetResult())

	multiplyExpression(&calc2, 3)
	fmt.Printf("Second calculator after multiplying by 3: %.2f\n", calc2.GetResult())
}

// Additional example: Circle with methods
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}
