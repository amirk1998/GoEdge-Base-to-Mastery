// io_examples.go
package internal

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"
)

// Color and formatting functions (assuming these are defined elsewhere)
// Removed local definitions to avoid redeclaration errors. Use those from colors.go

// RunIOExamples - main function to run all IO package examples
func RunIOExamples() {
	fmt.Println(Subtitle("📄 IO Package Examples"))
	fmt.Println()

	readerInterfaceDemo()
	writerInterfaceDemo()
	readerWriterDemo()
	copyOperationsDemo()
	multiReaderWriterDemo()
	limitedReaderDemo()
	pipeDemo()
	sectionReaderDemo()
	teeReaderDemo()
	ioUtilityFunctionsDemo()
}

// Reader Interface Examples
func readerInterfaceDemo() {
	fmt.Println(Yellow("📌 Reader Interface:"))

	// String reader
	content := "Hello, Go IO package! This is a demonstration of Reader interface."
	stringReader := strings.NewReader(content)

	// Read in chunks
	buffer := make([]byte, 10)
	fmt.Println(Bold("Reading in 10-byte chunks:"))

	for i := 0; i < 3; i++ {
		n, err := stringReader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading: %v\n", err)
			break
		}

		fmt.Printf("Chunk %d (%d bytes): %s\n",
			i+1, n, Green(string(buffer[:n])))

		if err == io.EOF {
			fmt.Println(Dim("Reached end of file"))
			break
		}
	}

	// Reset reader and read all
	stringReader.Reset(content)
	allData, err := io.ReadAll(stringReader)
	if err != nil {
		fmt.Printf("Error reading all: %v\n", err)
	} else {
		fmt.Printf("Read all data: %s\n", Cyan(string(allData)))
	}
	fmt.Println()
}

// Writer Interface Examples
func writerInterfaceDemo() {
	fmt.Println(Yellow("📌 Writer Interface:"))

	// Bytes buffer writer
	var buffer bytes.Buffer

	// Write different types of data
	data := []string{
		"First line of text\n",
		"Second line with numbers: 12345\n",
		"Third line with special chars: !@#$%\n",
	}

	fmt.Println(Bold("Writing to buffer:"))
	totalBytes := 0

	for i, line := range data {
		n, err := buffer.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing: %v\n", err)
			continue
		}
		totalBytes += n
		fmt.Printf("Line %d: wrote %d bytes\n", i+1, n)
	}

	fmt.Printf("Total bytes written: %s\n", Green(fmt.Sprintf("%d", totalBytes)))
	fmt.Printf("Buffer content:\n%s", Cyan(buffer.String()))

	// Write to multiple writers
	var buffer1, buffer2 bytes.Buffer
	multiWriter := io.MultiWriter(&buffer1, &buffer2)

	multiWriter.Write([]byte("This text goes to both buffers!"))
	fmt.Printf("Buffer1: %s\n", Yellow(buffer1.String()))
	fmt.Printf("Buffer2: %s\n", Yellow(buffer2.String()))
	fmt.Println()
}

// ReaderWriter Example
func readerWriterDemo() {
	fmt.Println(Yellow("📌 ReaderWriter Interface:"))

	// Create a buffer that implements both Reader and Writer
	var buffer bytes.Buffer

	// Write some initial data
	initialData := "Initial content in buffer"
	buffer.WriteString(initialData)
	fmt.Printf("Initial buffer: %s\n", Green(buffer.String()))

	// Read from it
	readBuffer := make([]byte, 10)
	n, err := buffer.Read(readBuffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read from buffer: %s\n", Cyan(string(readBuffer[:n])))
	}

	// Write more data
	buffer.WriteString(" + Additional content")
	fmt.Printf("After writing more: %s\n", Yellow(buffer.String()))

	// Demonstrate ReadWriter with file
	tempFile, err := os.CreateTemp("", "readwriter_test_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write to file
	content := "File content for ReadWriter example"
	tempFile.WriteString(content)

	// Seek to beginning to read
	tempFile.Seek(0, 0)

	// Read from file
	fileBuffer := make([]byte, len(content))
	n, err = tempFile.Read(fileBuffer)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
	} else {
		fmt.Printf("Read from file: %s\n", Green(string(fileBuffer[:n])))
	}
	fmt.Println()
}

// Copy Operations
func copyOperationsDemo() {
	fmt.Println(Yellow("📌 Copy Operations:"))

	// Basic copy
	source := strings.NewReader("This is the source content that will be copied.")
	var destination bytes.Buffer

	n, err := io.Copy(&destination, source)
	if err != nil {
		fmt.Printf("Error copying: %v\n", err)
		return
	}

	fmt.Printf("Copied %d bytes\n", n)
	fmt.Printf("Destination content: %s\n", Green(destination.String()))

	// Copy with limit
	source2 := strings.NewReader("This content will be partially copied due to limit.")
	var limitedDestination bytes.Buffer

	n, err = io.CopyN(&limitedDestination, source2, 20)
	if err != nil && err != io.EOF {
		fmt.Printf("Error in limited copy: %v\n", err)
	} else {
		fmt.Printf("Limited copy - copied %d bytes: %s\n",
			n, Cyan(limitedDestination.String()))
	}

	// Copy buffer (with custom buffer size)
	source3 := strings.NewReader("This demonstrates copy with custom buffer size for efficiency.")
	var bufferedDestination bytes.Buffer

	buffer := make([]byte, 8) // Small buffer for demonstration
	n, err = io.CopyBuffer(&bufferedDestination, source3, buffer)
	if err != nil {
		fmt.Printf("Error in buffered copy: %v\n", err)
	} else {
		fmt.Printf("Buffered copy - copied %d bytes: %s\n",
			n, Yellow(bufferedDestination.String()))
	}
	fmt.Println()
}

// MultiReader and MultiWriter Examples
func multiReaderWriterDemo() {
	fmt.Println(Yellow("📌 MultiReader and MultiWriter:"))

	// MultiReader - combines multiple readers
	reader1 := strings.NewReader("First part. ")
	reader2 := strings.NewReader("Second part. ")
	reader3 := strings.NewReader("Third part.")

	multiReader := io.MultiReader(reader1, reader2, reader3)

	combined, err := io.ReadAll(multiReader)
	if err != nil {
		fmt.Printf("Error reading from MultiReader: %v\n", err)
	} else {
		fmt.Printf("MultiReader result: %s\n", Green(string(combined)))
	}

	// MultiWriter - writes to multiple destinations
	var buffer1, buffer2, buffer3 bytes.Buffer
	multiWriter := io.MultiWriter(&buffer1, &buffer2, &buffer3)

	content := "This content goes to all three buffers simultaneously!"
	n, err := multiWriter.Write([]byte(content))
	if err != nil {
		fmt.Printf("Error writing to MultiWriter: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes to multiple destinations\n", n)
		fmt.Printf("Buffer 1: %s\n", Cyan(buffer1.String()))
		fmt.Printf("Buffer 2: %s\n", Yellow(buffer2.String()))
		fmt.Printf("Buffer 3: %s\n", Green(buffer3.String()))
	}
	fmt.Println()
}

// LimitedReader Example
func limitedReaderDemo() {
	fmt.Println(Yellow("📌 LimitedReader:"))

	source := strings.NewReader("This is a long string that will be limited by LimitedReader.")

	// Create limited reader with 25 byte limit
	limitedReader := &io.LimitedReader{R: source, N: 25}

	result, err := io.ReadAll(limitedReader)
	if err != nil {
		fmt.Printf("Error reading from LimitedReader: %v\n", err)
	} else {
		fmt.Printf("Limited read (25 bytes): %s\n", Green(string(result)))
		fmt.Printf("Bytes remaining in limit: %d\n", limitedReader.N)
	}

	// Reset and try reading more than limit
	source.Reset("Another example with different content for limited reading.")
	limitedReader2 := &io.LimitedReader{R: source, N: 15}

	buffer := make([]byte, 30) // Buffer larger than limit
	n, err := limitedReader2.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes (limit was 15): %s\n",
			n, Cyan(string(buffer[:n])))
	}
	fmt.Println()
}

// Pipe Example
func pipeDemo() {
	fmt.Println(Yellow("📌 Pipe Operations:"))

	// Create a pipe
	pipeReader, pipeWriter := io.Pipe()

	// Goroutine to write to pipe
	go func() {
		defer pipeWriter.Close()

		data := []string{
			"First chunk of data\n",
			"Second chunk of data\n",
			"Final chunk of data\n",
		}

		for i, chunk := range data {
			n, err := pipeWriter.Write([]byte(chunk))
			if err != nil {
				fmt.Printf("Error writing to pipe: %v\n", err)
				return
			}
			fmt.Printf("Wrote chunk %d: %d bytes\n", i+1, n)
		}
		fmt.Println(Dim("Pipe writer closed"))
	}()

	// Read from pipe
	fmt.Println(Bold("Reading from pipe:"))
	buffer := make([]byte, 1024)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println(Dim("Pipe reader reached EOF"))
				break
			}
			fmt.Printf("Error reading from pipe: %v\n", err)
			break
		}

		if n > 0 {
			fmt.Printf("Read from pipe: %s", Green(string(buffer[:n])))
		}
	}
	pipeReader.Close()
	fmt.Println()
}

// SectionReader Example
func sectionReaderDemo() {
	fmt.Println(Yellow("📌 SectionReader:"))

	// Create content
	fullContent := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	source := strings.NewReader(fullContent)

	// Create section reader (offset: 10, length: 15)
	sectionReader := io.NewSectionReader(source, 10, 15)

	// Read the section
	sectionData, err := io.ReadAll(sectionReader)
	if err != nil {
		fmt.Printf("Error reading section: %v\n", err)
	} else {
		fmt.Printf("Full content: %s\n", Dim(fullContent))
		fmt.Printf("Section (offset 10, length 15): %s\n", Green(string(sectionData)))
	}

	// Demonstrate seeking within section
	sectionReader.Seek(0, 0) // Reset to beginning of section
	buffer := make([]byte, 5)
	n, err := sectionReader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading after seek: %v\n", err)
	} else {
		fmt.Printf("First 5 bytes of section: %s\n", Cyan(string(buffer[:n])))
	}

	// Get size and current position
	fmt.Printf("Section size: %d\n", sectionReader.Size())
	fmt.Println()
}

// TeeReader Example
func teeReaderDemo() {
	fmt.Println(Yellow("📌 TeeReader:"))

	source := strings.NewReader("This content will be read and simultaneously copied to another writer.")
	var teeOutput bytes.Buffer

	// Create TeeReader
	teeReader := io.TeeReader(source, &teeOutput)

	// Read from TeeReader
	data, err := io.ReadAll(teeReader)
	if err != nil {
		fmt.Printf("Error reading from TeeReader: %v\n", err)
	} else {
		fmt.Printf("Read data: %s\n", Green(string(data)))
		fmt.Printf("Tee output: %s\n", Cyan(teeOutput.String()))
		fmt.Printf("Data matches: %s\n",
			Yellow(fmt.Sprintf("%t", string(data) == teeOutput.String())))
	}

	// Practical example: reading file while computing hash
	tempFile, err := os.CreateTemp("", "tee_example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write content to file
	fileContent := "Content for TeeReader hash computation example."
	tempFile.WriteString(fileContent)
	tempFile.Seek(0, 0)

	// Use TeeReader to read file and copy to buffer simultaneously
	var copyBuffer bytes.Buffer
	teeReader2 := io.TeeReader(tempFile, &copyBuffer)

	fileData, err := io.ReadAll(teeReader2)
	if err != nil {
		fmt.Printf("Error reading file with TeeReader: %v\n", err)
	} else {
		fmt.Printf("File content: %s\n", Green(string(fileData)))
		fmt.Printf("Copy buffer: %s\n", Yellow(copyBuffer.String()))
	}
	fmt.Println()
}

// IO Utility Functions
func ioUtilityFunctionsDemo() {
	fmt.Println(Yellow("📌 IO Utility Functions:"))

	// WriteString example
	var buffer bytes.Buffer
	n, err := io.WriteString(&buffer, "Hello from WriteString!")
	if err != nil {
		fmt.Printf("Error with WriteString: %v\n", err)
	} else {
		fmt.Printf("WriteString wrote %d bytes: %s\n", n, Green(buffer.String()))
	}

	// ReadAtLeast example
	source := strings.NewReader("This is content for ReadAtLeast demonstration.")
	readBuffer := make([]byte, 50)

	n, err = io.ReadAtLeast(source, readBuffer, 10)
	if err != nil {
		fmt.Printf("Error with ReadAtLeast: %v\n", err)
	} else {
		fmt.Printf("ReadAtLeast read %d bytes (minimum 10): %s\n",
			n, Cyan(string(readBuffer[:n])))
	}

	// ReadFull example
	source2 := strings.NewReader("Exact content for ReadFull")
	fullBuffer := make([]byte, 25) // Exact size

	n, err = io.ReadFull(source2, fullBuffer)
	if err != nil {
		fmt.Printf("Error with ReadFull: %v\n", err)
	} else {
		fmt.Printf("ReadFull read %d bytes: %s\n", n, Yellow(string(fullBuffer)))
	}

	// Random data example
	randomBuffer := make([]byte, 16)
	n, err = rand.Read(randomBuffer)
	if err != nil {
		fmt.Printf("Error reading random data: %v\n", err)
	} else {
		fmt.Printf("Random data (%d bytes): %s\n", n, Green(fmt.Sprintf("%x", randomBuffer)))
	}
	fmt.Println()
}
