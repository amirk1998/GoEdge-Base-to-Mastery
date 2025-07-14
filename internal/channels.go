// channels.go
package internal

import (
	"fmt"
	"time"
)

// RunChannelExamples - main function to run all channel examples
func RunChannelExamples() {
	basicChannelExample()
	bufferedChannelExample()
	channelDirectionExample()
	channelRangeExample()
	channelSelectExample()
	channelCloseExample()
	producerConsumerExample()
	fanOutFanInExample()
}

// Example 1: Basic unbuffered channel
func basicChannelExample() {
	fmt.Println("\n=== Basic Channel Example ===")

	ch := make(chan string)

	// Send data in a goroutine
	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "from"
		ch <- "Channel"
	}()

	// Receive data
	for i := 0; i < 4; i++ {
		msg := <-ch
		fmt.Printf("Received: %s\n", msg)
	}
}

// Example 2: Buffered channel
func bufferedChannelExample() {
	fmt.Println("\n=== Buffered Channel Example ===")

	// Create buffered channel with capacity 3
	ch := make(chan int, 3)

	// Send data (won't block because buffer has space)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("Channel length: %d, capacity: %d\n", len(ch), cap(ch))

	// Receive data
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("Received: %d\n", value)
	}
}

// Example 3: Channel direction (send-only, receive-only)
func sender(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

func channelDirectionExample() {
	fmt.Println("\n=== Channel Direction Example ===")

	ch := make(chan int)

	go sender(ch)
	receiver(ch)
}

// Example 4: Range over channel
func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

func channelRangeExample() {
	fmt.Println("\n=== Channel Range Example ===")

	ch := make(chan int)

	go generateNumbers(ch)

	// Range over channel (continues until channel is closed)
	for num := range ch {
		fmt.Printf("Processing: %d\n", num)
	}

	fmt.Println("All numbers processed")
}

// Example 5: Select with channels
func channelSelectExample() {
	fmt.Println("\n=== Channel Select Example ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()

	// Select waits for first available channel
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg2)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// Non-blocking select
	select {
	case msg := <-ch1:
		fmt.Printf("Got message: %s\n", msg)
	default:
		fmt.Println("No message available")
	}
}

// Example 6: Channel close detection
func channelCloseExample() {
	fmt.Println("\n=== Channel Close Example ===")

	ch := make(chan int, 3)

	// Send some data
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// Method 1: Check if channel is closed
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}

	// Method 2: Using range (automatically detects close)
	ch2 := make(chan string, 2)
	ch2 <- "Hello"
	ch2 <- "World"
	close(ch2)

	fmt.Println("Using range:")
	for msg := range ch2 {
		fmt.Printf("Received: %s\n", msg)
	}
}

// Example 7: Producer-Consumer pattern
func producer(ch chan<- int, id int) {
	for i := 1; i <= 5; i++ {
		product := id*10 + i
		ch <- product
		fmt.Printf("Producer %d produced: %d\n", id, product)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(ch <-chan int, id int) {
	for product := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, product)
		time.Sleep(150 * time.Millisecond)
	}
}

func producerConsumerExample() {
	fmt.Println("\n=== Producer-Consumer Example ===")

	ch := make(chan int, 5) // Buffered channel

	// Start producers
	go producer(ch, 1)
	go producer(ch, 2)

	// Start consumer
	go consumer(ch, 1)

	// Let it run for a while
	time.Sleep(2 * time.Second)
	close(ch)

	// Give consumer time to finish
	time.Sleep(500 * time.Millisecond)
}

// Example 8: Fan-out, Fan-in pattern
func fanOutFanInExample() {
	fmt.Println("\n=== Fan-out, Fan-in Example ===")

	// Input channel
	input := make(chan int)

	// Fan-out: distribute work to multiple workers
	worker1 := make(chan int)
	worker2 := make(chan int)
	worker3 := make(chan int)

	// Start workers
	go func() {
		for n := range worker1 {
			fmt.Printf("Worker 1 processing: %d\n", n)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for n := range worker2 {
			fmt.Printf("Worker 2 processing: %d\n", n)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	go func() {
		for n := range worker3 {
			fmt.Printf("Worker 3 processing: %d\n", n)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Fan-out goroutine
	go func() {
		defer close(worker1)
		defer close(worker2)
		defer close(worker3)

		for n := range input {
			select {
			case worker1 <- n:
			case worker2 <- n:
			case worker3 <- n:
			}
		}
	}()

	// Send work
	go func() {
		defer close(input)
		for i := 1; i <= 9; i++ {
			input <- i
		}
	}()

	// Wait for processing
	time.Sleep(3 * time.Second)
}

// Additional helper functions
func pingPong(ping chan<- string, pong <-chan string) {
	for i := 0; i < 3; i++ {
		ping <- "ping"
		fmt.Printf("Sent: ping\n")
		response := <-pong
		fmt.Printf("Received: %s\n", response)
	}
	close(ping)
}

func pongResponse(ping <-chan string, pong chan<- string) {
	for range ping {
		pong <- "pong"
	}
	close(pong)
}
