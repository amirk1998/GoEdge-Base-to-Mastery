package internal

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

// RunArraySliceProfessionalExamples - main function to run all professional array and slice examples
func RunArraySliceProfessionalExamples() {
	fmt.Println(Subtitle("📊 Professional Arrays Examples:"))
	demonstrateArrays()

	fmt.Println(Subtitle("🔀 Professional Slices Examples:"))
	demonstrateSlices()
	demonstrateSliceGrowth()
	demonstrateSliceOperations()

	fmt.Println(Subtitle("⚠️  Memory Management & Gotchas:"))
	demonstrateMemoryLeaks()

	fmt.Println(Subtitle("🚀 Advanced Techniques:"))
	demonstrateAdvancedTechniques()

	fmt.Println(Subtitle("⚡ Performance Analysis:"))
	compareSlicePerformance()

	fmt.Println(Subtitle("🌍 Real-World Examples:"))
	demonstrateRealWorldExamples()
}

// ==============================================================================
// 1. ARRAYS - Fixed Size, Value Type
// ==============================================================================

func demonstrateArrays() {
	fmt.Println(InfoText("=== ARRAYS DEMONSTRATION ==="))

	// Arrays are value types with fixed size
	var numbers [5]int
	fmt.Printf("Zero-valued array: %v\n", numbers)

	// Array initialization methods
	primes := [5]int{2, 3, 5, 7, 11}
	auto := [...]int{1, 2, 3, 4, 5}  // Compiler determines size
	sparse := [10]int{1: 42, 9: 100} // Sparse initialization

	fmt.Printf("Primes: %v\n", primes)
	fmt.Printf("Auto-sized: %v (len=%d)\n", auto, len(auto))
	fmt.Printf("Sparse: %v\n", sparse)

	// CRITICAL: Arrays are passed by value - expensive copy operation
	demonstrateArrayCopy(primes)
	fmt.Printf("Original after function call: %v\n", primes)

	// Memory layout comparison
	fmt.Printf("Array size in memory: %d bytes\n", unsafe.Sizeof(primes))
	fmt.Printf("Array address: %p\n", &primes)
	fmt.Printf("First element address: %p\n", &primes[0])

	fmt.Println()
}

func demonstrateArrayCopy(arr [5]int) {
	fmt.Printf("Function received copy at: %p\n", &arr)
	arr[0] = 999 // This won't affect original
	fmt.Printf("Modified copy: %v\n", arr)
}

// ==============================================================================
// 2. SLICES - Dynamic Arrays, Reference Type
// ==============================================================================

func demonstrateSlices() {
	fmt.Println(InfoText("=== SLICES DEMONSTRATION ==="))

	// Slice creation methods
	var nilSlice []int
	emptySlice := []int{}
	makeSlice := make([]int, 5)       // length 5, capacity 5
	makeWithCap := make([]int, 3, 10) // length 3, capacity 10

	fmt.Printf("Nil slice: %v (len=%d, cap=%d, nil=%v)\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("Empty slice: %v (len=%d, cap=%d, nil=%v)\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("Make slice: %v (len=%d, cap=%d)\n",
		makeSlice, len(makeSlice), cap(makeSlice))
	fmt.Printf("Make with capacity: %v (len=%d, cap=%d)\n",
		makeWithCap, len(makeWithCap), cap(makeWithCap))

	// Slice header anatomy
	demonstrateSliceHeader(makeWithCap)

	fmt.Println()
}

func demonstrateSliceHeader(s []int) {
	fmt.Printf("\nSlice header analysis:\n")
	fmt.Printf("Slice value: %v\n", s)
	fmt.Printf("Slice header size: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("Slice pointer: %p\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
	fmt.Printf("Slice length: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len)
	fmt.Printf("Slice capacity: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Cap)
}

// ==============================================================================
// 3. MEMORY MANAGEMENT AND GROWTH STRATEGY
// ==============================================================================

func demonstrateSliceGrowth() {
	fmt.Println(InfoText("=== SLICE GROWTH STRATEGY ==="))

	var numbers []int
	prevCap := 0

	for i := 0; i < 20; i++ {
		numbers = append(numbers, i)
		currentCap := cap(numbers)

		if currentCap != prevCap {
			fmt.Printf("Append %d: len=%d, cap=%d (growth from %d)\n",
				i, len(numbers), currentCap, prevCap)
			prevCap = currentCap
		}
	}

	// Pre-allocate for better performance
	fmt.Println("\nPre-allocation benefits:")
	demonstratePreAllocation()

	fmt.Println()
}

func demonstratePreAllocation() {
	const size = 1000000

	// Without pre-allocation
	start := time.Now()
	var slice1 []int
	for i := 0; i < size; i++ {
		slice1 = append(slice1, i)
	}
	withoutPreAlloc := time.Since(start)

	// With pre-allocation
	start = time.Now()
	slice2 := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice2 = append(slice2, i)
	}
	withPreAlloc := time.Since(start)

	fmt.Printf("Without pre-allocation: %v\n", withoutPreAlloc)
	fmt.Printf("With pre-allocation: %v\n", withPreAlloc)
	fmt.Printf("Performance improvement: %.2fx\n",
		float64(withoutPreAlloc)/float64(withPreAlloc))
}

// ==============================================================================
// 4. SLICE OPERATIONS AND GOTCHAS
// ==============================================================================

func demonstrateSliceOperations() {
	fmt.Println(InfoText("=== SLICE OPERATIONS ==="))

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Slicing operations
	sub1 := original[2:5] // elements 3, 4, 5
	sub2 := original[5:]  // elements from index 5 to end
	sub3 := original[:4]  // elements from start to index 3

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Sub1 [2:5]: %v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("Sub2 [5:]: %v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("Sub3 [:4]: %v (len=%d, cap=%d)\n", sub3, len(sub3), cap(sub3))

	// CRITICAL GOTCHA: Shared underlying array
	fmt.Println("\nShared underlying array demonstration:")
	sub1[0] = 999
	fmt.Printf("After modifying sub1[0]: original=%v, sub1=%v\n", original, sub1)

	// Full slice expression to control capacity
	safeSub := original[2:5:5] // [low:high:max] - capacity = max-low
	fmt.Printf("Safe sub with full slice: %v (len=%d, cap=%d)\n",
		safeSub, len(safeSub), cap(safeSub))

	fmt.Println()
}

// ==============================================================================
// 5. MEMORY LEAKS AND SLICE GOTCHAS
// ==============================================================================

func demonstrateMemoryLeaks() {
	fmt.Println(InfoText("=== MEMORY LEAK SCENARIOS ==="))

	// Scenario 1: Large slice with small sub-slice
	largeSlice := make([]byte, 1000000) // 1MB
	// Fill with some data
	for i := range largeSlice {
		largeSlice[i] = byte(i % 256)
	}

	// WRONG WAY - keeps reference to entire 1MB
	wrongSubSlice := largeSlice[:10]
	fmt.Printf("Wrong way - capacity kept: %d bytes\n", cap(wrongSubSlice))

	// RIGHT WAY - copy to break reference
	rightSubSlice := make([]byte, 10)
	copy(rightSubSlice, largeSlice[:10])
	fmt.Printf("Right way - capacity: %d bytes\n", cap(rightSubSlice))

	// Scenario 2: Slice append gotcha
	demonstrateAppendGotcha()

	fmt.Println()
}

func demonstrateAppendGotcha() {
	fmt.Println("\nAppend gotcha demonstration:")

	slice1 := make([]int, 3, 5)
	slice1[0], slice1[1], slice1[2] = 1, 2, 3

	slice2 := slice1[1:3] // shares underlying array
	fmt.Printf("slice1: %v (len=%d, cap=%d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v (len=%d, cap=%d)\n", slice2, len(slice2), cap(slice2))

	// This will overwrite slice1's data!
	slice2 = append(slice2, 4)
	fmt.Printf("After append to slice2:\n")
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
}

// ==============================================================================
// 6. PROFESSIONAL PATTERNS AND TECHNIQUES
// ==============================================================================

// Generic slice utility functions (Go 1.18+)
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice)) // Pre-allocate with capacity
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice)) // Pre-allocate with exact size
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

func Reduce[T, R any](slice []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// Thread-safe slice operations
type SafeSlice[T any] struct {
	mu    sync.RWMutex
	items []T
}

func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{
		items: make([]T, 0),
	}
}

func (s *SafeSlice[T]) Append(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

func (s *SafeSlice[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var zero T
	if index < 0 || index >= len(s.items) {
		return zero, false
	}
	return s.items[index], true
}

func (s *SafeSlice[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

func (s *SafeSlice[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// ==============================================================================
// 7. ADVANCED SLICE TECHNIQUES
// ==============================================================================

func demonstrateAdvancedTechniques() {
	fmt.Println(InfoText("=== ADVANCED TECHNIQUES ==="))

	// 1. Efficient removal without preserving order
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	indexToRemove := 4

	// O(1) removal - swap with last element
	numbers[indexToRemove] = numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-1]
	fmt.Printf("After O(1) removal: %v\n", numbers)

	// 2. Efficient insertion at beginning
	numbers = append([]int{0}, numbers...)
	fmt.Printf("After prepend: %v\n", numbers)

	// 3. Slice pooling for better performance
	demonstrateSlicePooling()

	// 4. Generic slice utilities
	evenNumbers := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	squares := Map([]int{1, 2, 3, 4, 5}, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)

	sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("Sum: %d\n", sum)

	fmt.Println()
}

var slicePool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 100)
	},
}

func demonstrateSlicePooling() {
	fmt.Println("\nSlice pooling demonstration:")

	// Get slice from pool
	slice := slicePool.Get().([]int)
	defer slicePool.Put(slice[:0]) // Reset length and return to pool

	// Use the slice
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	fmt.Printf("Pooled slice: %v\n", slice)
}

// ==============================================================================
// 8. PERFORMANCE BENCHMARKING PATTERNS
// ==============================================================================

func compareSlicePerformance() {
	fmt.Println(InfoText("=== PERFORMANCE COMPARISON ==="))

	const size = 1000000

	// Test 1: Append vs Pre-allocation
	start := time.Now()
	var slice1 []int
	for i := 0; i < size; i++ {
		slice1 = append(slice1, i)
	}
	appendTime := time.Since(start)

	start = time.Now()
	slice2 := make([]int, size)
	for i := 0; i < size; i++ {
		slice2[i] = i
	}
	indexTime := time.Since(start)

	fmt.Printf("Append: %v\n", appendTime)
	fmt.Printf("Direct indexing: %v\n", indexTime)
	fmt.Printf("Indexing is %.2fx faster\n",
		float64(appendTime)/float64(indexTime))

	// Test 2: Copy vs manual loop
	src := make([]int, size)

	start = time.Now()
	dst1 := make([]int, size)
	copy(dst1, src)
	copyTime := time.Since(start)

	start = time.Now()
	dst2 := make([]int, size)
	for i := 0; i < size; i++ {
		dst2[i] = src[i]
	}
	loopTime := time.Since(start)

	fmt.Printf("Built-in copy: %v\n", copyTime)
	fmt.Printf("Manual loop: %v\n", loopTime)
	fmt.Printf("Copy is %.2fx faster\n",
		float64(loopTime)/float64(copyTime))

	fmt.Println()
}

// ==============================================================================
// 9. REAL-WORLD SCENARIOS
// ==============================================================================

// Scenario 1: Buffer management for network operations
type CircularBuffer struct {
	buffer []byte
	head   int
	tail   int
	size   int
	full   bool
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer: make([]byte, size),
		size:   size,
	}
}

func (cb *CircularBuffer) Write(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	written := 0
	for _, b := range data {
		if cb.full && cb.head == cb.tail {
			cb.tail = (cb.tail + 1) % cb.size
		}

		cb.buffer[cb.head] = b
		cb.head = (cb.head + 1) % cb.size
		written++

		if cb.head == cb.tail {
			cb.full = true
		}
	}

	return written
}

func (cb *CircularBuffer) Read(data []byte) int {
	if len(data) == 0 || (!cb.full && cb.head == cb.tail) {
		return 0
	}

	read := 0
	for i := 0; i < len(data) && (cb.full || cb.head != cb.tail); i++ {
		data[i] = cb.buffer[cb.tail]
		cb.tail = (cb.tail + 1) % cb.size
		cb.full = false
		read++
	}

	return read
}

// Scenario 2: Event processing with sliding window
type SlidingWindow struct {
	events     []time.Time
	windowSize time.Duration
	maxEvents  int
}

func NewSlidingWindow(windowSize time.Duration, maxEvents int) *SlidingWindow {
	return &SlidingWindow{
		events:     make([]time.Time, 0, maxEvents),
		windowSize: windowSize,
		maxEvents:  maxEvents,
	}
}

func (sw *SlidingWindow) AddEvent() bool {
	now := time.Now()

	// Remove old events outside window
	cutoff := now.Add(-sw.windowSize)
	validStart := 0
	for i, event := range sw.events {
		if event.After(cutoff) {
			validStart = i
			break
		}
	}

	// Efficiently remove old events
	if validStart > 0 {
		copy(sw.events, sw.events[validStart:])
		sw.events = sw.events[:len(sw.events)-validStart]
	}

	// Check if we can add new event
	if len(sw.events) >= sw.maxEvents {
		return false
	}

	sw.events = append(sw.events, now)
	return true
}

func (sw *SlidingWindow) CurrentCount() int {
	return len(sw.events)
}

// ==============================================================================
// REAL-WORLD EXAMPLES DEMONSTRATION
// ==============================================================================

func demonstrateRealWorldExamples() {
	fmt.Println(InfoText("=== REAL-WORLD EXAMPLES ==="))

	// Circular buffer example
	buffer := NewCircularBuffer(5)
	buffer.Write([]byte("Hello"))
	readData := make([]byte, 10)
	n := buffer.Read(readData)
	fmt.Printf("Circular buffer read: %s (%d bytes)\n", string(readData[:n]), n)

	// Sliding window example
	window := NewSlidingWindow(time.Second, 3)
	for i := 0; i < 5; i++ {
		allowed := window.AddEvent()
		fmt.Printf("Event %d: allowed=%v, count=%d\n",
			i+1, allowed, window.CurrentCount())
		time.Sleep(300 * time.Millisecond)
	}

	// Thread-safe slice example
	safeSlice := NewSafeSlice[int]()
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			safeSlice.Append(val)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Thread-safe slice: %v\n", safeSlice.ToSlice())

	// Memory usage info
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\nMemory usage: %d KB\n", m.Alloc/1024)

	fmt.Println()
}
