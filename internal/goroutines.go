// goroutines.go
package internal

import (
	"fmt"
	"sync"
	"time"
)

// RunGoroutineExamples - main function to run all goroutine examples
func RunGoroutineExamples() {
	basicGoroutineExample()
	waitGroupExample()
	mutexExample()
	racConditionExample()
	goroutinePoolExample()
	selectStatementExample()
}

// Example 1: Basic goroutine
func printNumbers(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func basicGoroutineExample() {
	fmt.Println("\n=== Basic Goroutine Example ===")

	// Start goroutines
	go printNumbers("Goroutine 1")
	go printNumbers("Goroutine 2")

	// Main goroutine continues
	printNumbers("Main")

	// Wait a bit to see goroutines complete
	time.Sleep(1 * time.Second)
	fmt.Println("Main function ends")
}

// Example 2: WaitGroup for synchronization
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroupExample() {
	fmt.Println("\n=== WaitGroup Example ===")

	var wg sync.WaitGroup

	// Start multiple workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed")
}

// Example 3: Mutex for protecting shared resources
type Counter2 struct {
	mu    sync.Mutex
	value int
}

func (c *Counter2) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter2) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func mutexExample() {
	fmt.Println("\n=== Mutex Example ===")

	counter := &Counter2{}
	var wg sync.WaitGroup

	// Start multiple goroutines that increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}

// Example 4: Race condition demonstration
var unsafeCounter int

func incrementUnsafe(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		unsafeCounter++
	}
}

func racConditionExample() {
	fmt.Println("\n=== Race Condition Example ===")

	unsafeCounter = 0
	var wg sync.WaitGroup

	// Start multiple goroutines without proper synchronization
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementUnsafe(&wg)
	}

	wg.Wait()
	fmt.Printf("Unsafe counter final value: %d (should be 5000)\n", unsafeCounter)
	fmt.Println("Note: This demonstrates race condition - run multiple times to see different results")
}

// Example 5: Goroutine pool pattern
func processTask(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(100 * time.Millisecond)
		results <- task * task
	}
}

func goroutinePoolExample() {
	fmt.Println("\n=== Goroutine Pool Example ===")

	const numWorkers = 3
	const numTasks = 10

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go processTask(i, tasks, results, &wg)
	}

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Example 6: Select statement with channels
func selectStatementExample() {
	fmt.Println("\n=== Select Statement Example ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Select with timeout
	timeout := time.After(500 * time.Millisecond)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Received: %s\n", msg)
		case <-timeout:
			fmt.Println("Timeout reached")
			return
		}
	}
}

// Additional helper functions for demonstration
func longRunningTask(id int, duration time.Duration) {
	fmt.Printf("Task %d starting (duration: %v)\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Task %d completed\n", id)
}

func fibonacci2(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}
