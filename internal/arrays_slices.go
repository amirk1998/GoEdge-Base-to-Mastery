// arrays_slices.go
package internal

import (
	"fmt"
	"sort"
)

// RunArraySliceExamples - main function to run all array and slice examples
func RunArraySliceExamples() {
	fmt.Println(Subtitle("📊 Arrays Examples:"))
	basicArrayExample()
	arrayOperationsExample()
	arrayPassingExample()

	fmt.Println(Subtitle("🔀 Slices Examples:"))
	basicSliceExample()
	sliceOperationsExample()
	sliceMemoryExample()
	sliceAdvancedExample()
	slicePerformanceExample()
}

// basicArrayExample - demonstrates basic array operations
func basicArrayExample() {
	fmt.Println(Bold("1. Basic Array Operations:"))

	// Array declaration and initialization
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20

	// Array literal initialization
	fruits := [3]string{"apple", "banana", "orange"}

	// Array with ... (compiler determines size)
	scores := [...]int{85, 92, 78, 96, 88}

	fmt.Printf("Numbers array: %v\n", numbers)
	fmt.Printf("Fruits array: %v\n", fruits)
	fmt.Printf("Scores array: %v (length: %d)\n", scores, len(scores))

	// Array iteration
	fmt.Println("Iterating through scores:")
	for i, score := range scores {
		fmt.Printf("  Index %d: %d\n", i, score)
	}

	fmt.Println()
}

// arrayOperationsExample - demonstrates array operations and comparisons
func arrayOperationsExample() {
	fmt.Println(Bold("2. Array Operations and Comparisons:"))

	// Array comparison (same type and size)
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1 == arr2: %v\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %v\n", arr1 == arr3)

	// Multidimensional arrays
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:")
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d  ", i, j, val)
		}
		fmt.Println()
	}

	fmt.Println()
}

// arrayPassingExample - demonstrates how arrays are passed to functions
func arrayPassingExample() {
	fmt.Println(Bold("3. Array Passing (By Value):"))

	original := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Original before function call: %v\n", original)

	// Arrays are passed by value (copy)
	modifyArrayValue(original)
	fmt.Printf("Original after function call: %v\n", original)

	// To modify original, pass pointer to array
	modifyArrayPointer(&original)
	fmt.Printf("Original after pointer modification: %v\n", original)

	fmt.Println()
}

// modifyArrayValue - demonstrates array passed by value
func modifyArrayValue(arr [5]int) {
	arr[0] = 999
	fmt.Printf("Inside function (by value): %v\n", arr)
}

// modifyArrayPointer - demonstrates array passed by pointer
func modifyArrayPointer(arr *[5]int) {
	arr[0] = 777
	fmt.Printf("Inside function (by pointer): %v\n", *arr)
}

// basicSliceExample - demonstrates basic slice operations
func basicSliceExample() {
	fmt.Println(Bold("4. Basic Slice Operations:"))

	// Slice declaration and initialization
	var numbers []int
	fmt.Printf("Empty slice: %v (len: %d, cap: %d)\n", numbers, len(numbers), cap(numbers))

	// Slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits slice: %v (len: %d, cap: %d)\n", fruits, len(fruits), cap(fruits))

	// Using make
	scores := make([]int, 3, 5) // length 3, capacity 5
	fmt.Printf("Scores slice: %v (len: %d, cap: %d)\n", scores, len(scores), cap(scores))

	// Slicing arrays and slices
	array := [6]int{10, 20, 30, 40, 50, 60}
	slice1 := array[1:4] // [20, 30, 40]
	slice2 := array[:3]  // [10, 20, 30]
	slice3 := array[2:]  // [30, 40, 50, 60]

	fmt.Printf("Array: %v\n", array)
	fmt.Printf("slice1 [1:4]: %v\n", slice1)
	fmt.Printf("slice2 [:3]: %v\n", slice2)
	fmt.Printf("slice3 [2:]: %v\n", slice3)

	fmt.Println()
}

// sliceOperationsExample - demonstrates slice operations
func sliceOperationsExample() {
	fmt.Println(Bold("5. Slice Operations:"))

	// Append operation
	slice := []int{1, 2, 3}
	fmt.Printf("Original slice: %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))

	slice = append(slice, 4, 5, 6)
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))

	// Append another slice
	other := []int{7, 8, 9}
	slice = append(slice, other...)
	fmt.Printf("After append slice: %v (len: %d, cap: %d)\n", slice, len(slice), cap(slice))

	// Copy operation
	source := []int{10, 20, 30, 40, 50}
	dest := make([]int, 3)
	copied := copy(dest, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", dest)
	fmt.Printf("Elements copied: %d\n", copied)

	// Slice deletion (remove element at index 2)
	numbers := []int{1, 2, 3, 4, 5}
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Printf("After deletion at index 2: %v\n", numbers)

	fmt.Println()
}

// sliceMemoryExample - demonstrates slice memory behavior
func sliceMemoryExample() {
	fmt.Println(Bold("6. Slice Memory and References:"))

	// Slices share underlying array
	array := [5]int{1, 2, 3, 4, 5}
	slice1 := array[1:4]
	slice2 := array[2:5]

	fmt.Printf("Original array: %v\n", array)
	fmt.Printf("slice1 [1:4]: %v\n", slice1)
	fmt.Printf("slice2 [2:5]: %v\n", slice2)

	// Modifying slice affects shared array
	slice1[1] = 999
	fmt.Printf("After slice1[1] = 999:\n")
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	// Capacity and growth
	fmt.Println("\nCapacity growth example:")
	growth := make([]int, 0, 1)
	for i := 0; i < 8; i++ {
		growth = append(growth, i)
		fmt.Printf("len: %d, cap: %d, slice: %v\n", len(growth), cap(growth), growth)
	}

	fmt.Println()
}

// sliceAdvancedExample - demonstrates advanced slice techniques
func sliceAdvancedExample() {
	fmt.Println(Bold("7. Advanced Slice Techniques:"))

	// Filter slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterEvens(numbers)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Evens: %v\n", evens)

	// Map slice
	doubled := mapSlice(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Doubled: %v\n", doubled)

	// Reduce slice
	sum := reduceSlice(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", sum)

	// Sort slice
	names := []string{"Charlie", "Alice", "Bob", "David"}
	fmt.Printf("Before sort: %v\n", names)
	sort.Strings(names)
	fmt.Printf("After sort: %v\n", names)

	// Custom sort
	people := []PersonStr{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	fmt.Printf("Before custom sort: %v\n", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("After age sort: %v\n", people)

	fmt.Println()
}

// Person - struct for sorting example
type PersonStr struct {
	Name string
	Age  int
}

// filterEvens - filters even numbers from slice
func filterEvens(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

// mapSlice - applies function to each element
func mapSlice(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// reduceSlice - reduces slice to single value
func reduceSlice(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// slicePerformanceExample - demonstrates slice performance considerations
func slicePerformanceExample() {
	fmt.Println(Bold("8. Slice Performance Considerations:"))

	// Pre-allocate with known capacity
	fmt.Println("Performance comparison:")

	// Without pre-allocation
	slice1 := []int{}
	for i := 0; i < 1000; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Printf("Without pre-allocation - len: %d, cap: %d\n", len(slice1), cap(slice1))

	// With pre-allocation
	slice2 := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Printf("With pre-allocation - len: %d, cap: %d\n", len(slice2), cap(slice2))

	// Memory leak prevention
	fmt.Println("\nMemory leak prevention:")
	large := make([]int, 1000000)
	// Bad: keeps reference to large array
	// small := large[:5]

	// Good: copy to new slice
	small := make([]int, 5)
	copy(small, large[:5])
	large = nil // can be garbage collected

	fmt.Printf("Small slice from large array: %v (len: %d, cap: %d)\n", small, len(small), cap(small))

	fmt.Println()
}
