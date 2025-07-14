// pointers.go
package internal

import "fmt"

// RunPointerExamples - main function to run all pointer examples
func RunPointerExamples() {
	basicPointerExample()
	pointerFunctionExample()
	structPointerExample()
	slicePointerExample()
	nilPointerExample()
	pointerComparisonExample()
	returningPointerExample()
	pointerToPointerExample()
	arraySlicePointerExample()
	performanceExample()
}

// Example 1: Basic pointer usage
func basicPointerExample() {
	fmt.Println("\n=== Basic Pointer Example ===")

	// Create a variable
	x := 42

	// Create a pointer to x
	var ptr *int = &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of ptr (address it points to): %p\n", ptr)
	fmt.Printf("Value at the address ptr points to: %d\n", *ptr)

	// Modify value through pointer
	*ptr = 100
	fmt.Printf("After modifying through pointer, x = %d\n", x)
}

// Example 2: Pointer as function parameter
func modifyValue(ptr *int) {
	*ptr = *ptr * 2
}

func modifyValueByValue(val int) {
	val = val * 2
}

func pointerFunctionExample() {
	fmt.Println("\n=== Pointer Function Example ===")

	num := 10

	// Pass by value (doesn't modify original)
	modifyValueByValue(num)
	fmt.Printf("After pass by value: %d\n", num)

	// Pass by pointer (modifies original)
	modifyValue(&num)
	fmt.Printf("After pass by pointer: %d\n", num)
}

// Example 3: Pointer to struct
type Person struct {
	Name string
	Age  int
}

func (p *Person) celebrateBirthday() {
	p.Age++
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func structPointerExample() {
	fmt.Println("\n=== Struct Pointer Example ===")

	person := Person{Name: "Ali", Age: 25}

	fmt.Printf("Before birthday: %s\n", person.getInfo())

	// Method with pointer receiver
	person.celebrateBirthday()

	fmt.Printf("After birthday: %s\n", person.getInfo())

	// Working with pointer to struct
	personPtr := &person
	personPtr.Name = "Ali Ahmadi"
	fmt.Printf("After name change: %s\n", person.getInfo())
}

// Example 4: Pointer to slice
func appendToSlice(slice *[]int, value int) {
	*slice = append(*slice, value)
}

func slicePointerExample() {
	fmt.Println("\n=== Slice Pointer Example ===")

	numbers := []int{1, 2, 3}
	fmt.Printf("Original slice: %v\n", numbers)

	appendToSlice(&numbers, 4)
	appendToSlice(&numbers, 5)

	fmt.Printf("After appending: %v\n", numbers)
}

// Example 5: Nil pointer handling
func nilPointerExample() {
	fmt.Println("\n=== Nil Pointer Example ===")

	var ptr *int
	fmt.Printf("Nil pointer: %v\n", ptr)

	// Check before dereferencing
	if ptr != nil {
		fmt.Printf("Value: %d\n", *ptr)
	} else {
		fmt.Println("Pointer is nil, cannot dereference")
	}

	// Initialize pointer
	value := 42
	ptr = &value

	if ptr != nil {
		fmt.Printf("Now pointer has value: %d\n", *ptr)
	}
}

// Example 6: Pointer comparison
func pointerComparisonExample() {
	fmt.Println("\n=== Pointer Comparison Example ===")

	x := 10
	y := 20

	ptrX := &x
	ptrY := &y
	ptrX2 := &x

	fmt.Printf("ptrX == ptrY: %v\n", ptrX == ptrY)
	fmt.Printf("ptrX == ptrX2: %v\n", ptrX == ptrX2)
}

// Example 7: Returning pointer from function
func createPerson(name string, age int) *Person {
	// This is safe in Go - the variable will be allocated on heap
	return &Person{Name: name, Age: age}
}

func returningPointerExample() {
	fmt.Println("\n=== Returning Pointer Example ===")

	person := createPerson("Sara", 30)
	fmt.Printf("Created person: %s\n", person.getInfo())
}

// Example 8: Pointer to pointer
func pointerToPointerExample() {
	fmt.Println("\n=== Pointer to Pointer Example ===")

	value := 42
	ptr := &value
	ptrToPtr := &ptr

	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Pointer to value: %p\n", ptr)
	fmt.Printf("Pointer to pointer: %p\n", ptrToPtr)
	fmt.Printf("Value through pointer: %d\n", *ptr)
	fmt.Printf("Value through pointer to pointer: %d\n", **ptrToPtr)

	// Modify through pointer to pointer
	**ptrToPtr = 100
	fmt.Printf("After modification: %d\n", value)
}

// Example 9: Array vs Slice with pointers
func arraySlicePointerExample() {
	fmt.Println("\n=== Array vs Slice Pointer Example ===")

	// Array
	arr := [3]int{1, 2, 3}
	arrPtr := &arr

	fmt.Printf("Original array: %v\n", arr)
	(*arrPtr)[0] = 100
	fmt.Printf("After modification through pointer: %v\n", arr)

	// Slice
	slice := []int{1, 2, 3}
	slicePtr := &slice

	fmt.Printf("Original slice: %v\n", slice)
	(*slicePtr)[0] = 200
	fmt.Printf("After modification through pointer: %v\n", slice)
}

// Example 10: Performance consideration
func expensiveOperation(data []int) {
	// Simulate expensive operation
	for i := 0; i < len(data); i++ {
		data[i] = data[i] * 2
	}
}

func expensiveOperationWithPointer(data *[]int) {
	// Working with pointer to avoid copying
	for i := 0; i < len(*data); i++ {
		(*data)[i] = (*data)[i] * 2
	}
}

func performanceExample() {
	fmt.Println("\n=== Performance Example ===")

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original data: %v\n", data)

	// Using pointer to avoid copying (though slices are reference types anyway)
	expensiveOperationWithPointer(&data)
	fmt.Printf("After operation with pointer: %v\n", data)
}
