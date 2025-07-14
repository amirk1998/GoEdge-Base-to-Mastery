// interfaces.go
package internal

import "fmt"

// RunInterfaceExamples - main function to run all interface examples
func RunInterfaceExamples() {
	basicInterfaceExample()
	multipleInterfaceExample()
	emptyInterfaceExample()
	typeAssertionExample()
	interfaceCompositionExample()
	polymorphismExample()
}

// Example 1: Basic interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func printShapeInfo(shape Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
}

func basicInterfaceExample() {
	fmt.Println("\n=== Basic Interface Example ===")

	square := Square{Side: 5}

	// Direct call
	fmt.Printf("Square with side 5:\n")
	printShapeInfo(square)

	// Interface variable
	var shape Shape = square
	fmt.Printf("Through interface variable:\n")
	printShapeInfo(shape)
}

// Example 2: Multiple types implementing same interface
type Writer interface {
	Write(data string) error
}

type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data string) error {
	fmt.Printf("Writing to file '%s': %s\n", fw.Filename, data)
	return nil
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) error {
	fmt.Printf("Console output: %s\n", data)
	return nil
}

type NetworkWriter struct {
	URL string
}

func (nw NetworkWriter) Write(data string) error {
	fmt.Printf("Sending to %s: %s\n", nw.URL, data)
	return nil
}

func writeData(w Writer, data string) {
	err := w.Write(data)
	if err != nil {
		fmt.Printf("Error writing data: %v\n", err)
	}
}

func multipleInterfaceExample() {
	fmt.Println("\n=== Multiple Interface Implementation Example ===")

	writers := []Writer{
		FileWriter{Filename: "output.txt"},
		ConsoleWriter{},
		NetworkWriter{URL: "http://api.example.com"},
	}

	for _, writer := range writers {
		writeData(writer, "Hello, World!")
	}
}

// Example 3: Empty interface
func processAnyType(value interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", value, value)
}

func emptyInterfaceExample() {
	fmt.Println("\n=== Empty Interface Example ===")

	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
	}

	for _, value := range values {
		processAnyType(value)
	}
}

// Example 4: Type assertion
func typeAssertionExample() {
	fmt.Println("\n=== Type Assertion Example ===")

	var value interface{} = "hello world"

	// Type assertion with ok check
	if str, ok := value.(string); ok {
		fmt.Printf("String value: %s\n", str)
	} else {
		fmt.Println("Not a string")
	}

	// Type assertion without ok check (can panic)
	str := value.(string)
	fmt.Printf("Direct assertion: %s\n", str)

	// Type switch
	values := []interface{}{42, "hello", 3.14, true}

	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("Integer: %d\n", val)
		case string:
			fmt.Printf("String: %s\n", val)
		case float64:
			fmt.Printf("Float: %.2f\n", val)
		case bool:
			fmt.Printf("Boolean: %t\n", val)
		default:
			fmt.Printf("Unknown type: %T\n", val)
		}
	}
}

// Example 5: Interface composition
type Reader interface {
	Read() (string, error)
}

type Writer2 interface {
	Write(data string) error
}

type ReadWriter interface {
	Reader
	Writer2
}

type FileHandler struct {
	Filename string
	Data     string
}

func (fh *FileHandler) Read() (string, error) {
	fmt.Printf("Reading from file: %s\n", fh.Filename)
	return fh.Data, nil
}

func (fh *FileHandler) Write(data string) error {
	fmt.Printf("Writing to file %s: %s\n", fh.Filename, data)
	fh.Data = data
	return nil
}

func processReadWriter(rw ReadWriter) {
	// Read data
	data, err := rw.Read()
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}

	// Process data
	processedData := "Processed: " + data

	// Write back
	err = rw.Write(processedData)
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	}
}

func interfaceCompositionExample() {
	fmt.Println("\n=== Interface Composition Example ===")

	handler := &FileHandler{
		Filename: "data.txt",
		Data:     "original data",
	}

	processReadWriter(handler)
}

// Example 6: Polymorphism with interfaces
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Sneaking"
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return "Chirp!"
}

func (b Bird) Move() string {
	return "Flying"
}

func interactWithAnimal(animal Animal) {
	fmt.Printf("Animal speaks: %s\n", animal.Speak())
	fmt.Printf("Animal moves: %s\n", animal.Move())
}

func polymorphismExample() {
	fmt.Println("\n=== Polymorphism Example ===")

	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}

	for i, animal := range animals {
		fmt.Printf("Animal %d:\n", i+1)
		interactWithAnimal(animal)
		fmt.Println()
	}
}

// Additional interfaces for demonstration
type Stringer interface {
	String() string
}

type Error interface {
	Error() string
}

type CustomError struct {
	Code    int
	Message string
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}
