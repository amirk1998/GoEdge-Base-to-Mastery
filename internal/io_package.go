// io_package_examples.go
package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Color and formatting functions
// func Red(text string) string    { return fmt.Sprintf("\033[31m%s\033[0m", text) }
// func Green(text string) string  { return fmt.Sprintf("\033[32m%s\033[0m", text) }
// func Yellow(text string) string { return fmt.Sprintf("\033[33m%s\033[0m", text) }
// func Blue(text string) string   { return fmt.Sprintf("\033[34m%s\033[0m", text) }
// func Cyan(text string) string   { return fmt.Sprintf("\033[36m%s\033[0m", text) }
// func Bold(text string) string   { return fmt.Sprintf("\033[1m%s\033[0m", text) }

func SectionHeader(title string) string {
	return fmt.Sprintf("\n%s\n%s", Bold(Blue("=== "+title+" ===")), strings.Repeat("-", len(title)+8))
}

// func ErrorText(text string) string   { return Red("ERROR: " + text) }
// func WarningText(text string) string { return Yellow("WARNING: " + text) }
// func InfoText(text string) string    { return Cyan("INFO: " + text) }

// RunIOPackageExamples - main function to run all IO package examples
func RunIOPackageExamples() {
	basicReaderWriterExample()
	stringReaderWriterExample()
	copyOperationsExample()
	bufferOperationsExample()
	pipeOperationsExample()
	multiReaderWriterExample()
	limitedReaderExample()
	sectionReaderExample()
	teeReaderExample()
	readerWriterInterfaces()
}

// basicReaderWriterExample demonstrates basic Reader and Writer interfaces
func basicReaderWriterExample() {
	fmt.Println(SectionHeader("Basic Reader and Writer"))

	// Create a string reader
	content := "Hello, World! This is a test string for IO operations."
	reader := strings.NewReader(content)

	fmt.Printf("Original content: %s\n", Yellow(content))
	fmt.Printf("Reader size: %d bytes\n", reader.Size())

	// Read data in chunks
	buffer := make([]byte, 10)
	fmt.Println("Reading in 10-byte chunks:")

	for i := 0; i < 3; i++ {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println(InfoText("Reached end of file"))
				break
			}
			fmt.Printf("Error reading: %s\n", ErrorText(err.Error()))
			break
		}
		fmt.Printf("Chunk %d: %s (read %d bytes)\n", i+1, Green(string(buffer[:n])), n)
	}

	// Reset reader position
	reader.Seek(0, io.SeekStart)

	// Read all at once
	allData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading all: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Read all data: %s\n", Cyan(string(allData)))

	// Write to buffer
	var buffer2 bytes.Buffer
	writer := &buffer2

	data := "This is data being written to a buffer."
	n, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Printf("Error writing: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Written %d bytes: %s\n", n, Green(buffer2.String()))
	fmt.Println()
}

// stringReaderWriterExample demonstrates string-based readers and writers
func stringReaderWriterExample() {
	fmt.Println(SectionHeader("String Reader and Writer"))

	// String Reader example
	text := "The quick brown fox jumps over the lazy dog."
	reader := strings.NewReader(text)

	fmt.Printf("Original text: %s\n", Bold(text))
	fmt.Printf("Reader length: %d\n", reader.Len())

	// Read specific number of bytes
	buffer := make([]byte, 15)
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, Yellow(string(buffer[:n])))

	// Read a single byte
	b, err := reader.ReadByte()
	if err != nil {
		fmt.Printf("Error reading byte: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Next byte: %s\n", Green(string(b)))

	// Read until delimiter
	reader.Seek(0, io.SeekStart) // Reset position
	word, err := reader.ReadByte()
	if err != nil {
		fmt.Printf("Error reading: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("First character: %s\n", Cyan(string(word)))

	// String Builder (Writer) example
	var builder strings.Builder
	builder.WriteString("Building a string ")
	builder.WriteString("piece by piece ")
	builder.WriteByte('!')

	result := builder.String()
	fmt.Printf("Built string: %s\n", Green(result))
	fmt.Printf("Builder length: %d\n", builder.Len())

	// Reset and build again
	builder.Reset()
	builder.WriteString("Fresh start after reset")
	fmt.Printf("After reset: %s\n", Yellow(builder.String()))
	fmt.Println()
}

// copyOperationsExample demonstrates copy operations
func copyOperationsExample() {
	fmt.Println(SectionHeader("Copy Operations"))

	// Basic copy
	source := "This is the source data that will be copied."
	reader := strings.NewReader(source)
	var destination bytes.Buffer

	n, err := io.Copy(&destination, reader)
	if err != nil {
		fmt.Printf("Error copying: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Source: %s\n", Yellow(source))
	fmt.Printf("Copied %d bytes\n", n)
	fmt.Printf("Destination: %s\n", Green(destination.String()))

	// Copy with limit
	source2 := "This is a longer source text that will be partially copied."
	reader2 := strings.NewReader(source2)
	var destination2 bytes.Buffer

	n, err = io.CopyN(&destination2, reader2, 20)
	if err != nil {
		fmt.Printf("Error copying with limit: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Limited copy source: %s\n", Yellow(source2))
	fmt.Printf("Copied %d bytes (limited)\n", n)
	fmt.Printf("Limited destination: %s\n", Green(destination2.String()))

	// Copy to file
	tempFile, err := os.CreateTemp("", "copy_example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %s\n", ErrorText(err.Error()))
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	fileContent := "This content will be written to a file."
	reader3 := strings.NewReader(fileContent)

	n, err = io.Copy(tempFile, reader3)
	if err != nil {
		fmt.Printf("Error copying to file: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Copied %d bytes to file: %s\n", n, Cyan(tempFile.Name()))

	// Read back from file
	tempFile.Seek(0, io.SeekStart)
	fileData, err := io.ReadAll(tempFile)
	if err != nil {
		fmt.Printf("Error reading from file: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Read from file: %s\n", Green(string(fileData)))
	fmt.Println()
}

// bufferOperationsExample demonstrates buffer operations
func bufferOperationsExample() {
	fmt.Println(SectionHeader("Buffer Operations"))

	// Create a buffer and write different types of data
	var buffer bytes.Buffer

	// Write string
	buffer.WriteString("Hello, ")

	// Write bytes
	buffer.Write([]byte("World! "))

	// Write single byte
	buffer.WriteByte('!')

	// Write formatted string
	fmt.Fprintf(&buffer, " Current time: %s", time.Now().Format("15:04:05"))

	fmt.Printf("Buffer content: %s\n", Yellow(buffer.String()))
	fmt.Printf("Buffer length: %d bytes\n", buffer.Len())

	// Read from buffer
	readBuffer := make([]byte, 10)
	n, err := buffer.Read(readBuffer)
	if err != nil {
		fmt.Printf("Error reading from buffer: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, Green(string(readBuffer[:n])))
	fmt.Printf("Remaining buffer: %s\n", Cyan(buffer.String()))

	// Reset buffer
	buffer.Reset()
	fmt.Printf("After reset, buffer length: %d\n", buffer.Len())

	// Buffer with initial content
	initialContent := "Initial buffer content"
	buffer2 := bytes.NewBufferString(initialContent)
	fmt.Printf("Buffer with initial content: %s\n", Green(buffer2.String()))

	// Truncate buffer
	buffer2.WriteString(" - Additional content")
	fmt.Printf("After adding content: %s\n", Yellow(buffer2.String()))

	buffer2.Truncate(len(initialContent))
	fmt.Printf("After truncation: %s\n", Cyan(buffer2.String()))
	fmt.Println()
}

// pipeOperationsExample demonstrates pipe operations
func pipeOperationsExample() {
	fmt.Println(SectionHeader("Pipe Operations"))

	// Create a pipe
	reader, writer := io.Pipe()

	// Write to pipe in a goroutine
	go func() {
		defer writer.Close()

		messages := []string{
			"First message through pipe",
			"Second message through pipe",
			"Third message through pipe",
		}

		for _, message := range messages {
			fmt.Fprintf(writer, "%s\n", message)
			time.Sleep(100 * time.Millisecond) // Simulate processing time
		}
	}()

	// Read from pipe
	fmt.Println("Reading from pipe:")
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading from pipe: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Pipe data:\n%s", Green(string(data)))

	// Pipe with error handling
	reader2, writer2 := io.Pipe()

	go func() {
		defer writer2.Close()

		// Simulate an error condition
		_, err := writer2.Write([]byte("Some data before error"))
		if err != nil {
			fmt.Printf("Error writing to pipe: %s\n", ErrorText(err.Error()))
			return
		}

		// Close with error
		writer2.CloseWithError(fmt.Errorf("simulated pipe error"))
	}()

	// Read from pipe that will have an error
	data2, err := io.ReadAll(reader2)
	if err != nil {
		fmt.Printf("Expected error from pipe: %s\n", WarningText(err.Error()))
	} else {
		fmt.Printf("Data before error: %s\n", Yellow(string(data2)))
	}
	fmt.Println()
}

// multiReaderWriterExample demonstrates multi-reader and multi-writer
func multiReaderWriterExample() {
	fmt.Println(SectionHeader("Multi-Reader and Multi-Writer"))

	// Multi-Reader example
	reader1 := strings.NewReader("First part. ")
	reader2 := strings.NewReader("Second part. ")
	reader3 := strings.NewReader("Third part.")

	multiReader := io.MultiReader(reader1, reader2, reader3)

	data, err := io.ReadAll(multiReader)
	if err != nil {
		fmt.Printf("Error reading from multi-reader: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Multi-reader result: %s\n", Green(string(data)))

	// Multi-Writer example
	var buffer1 bytes.Buffer
	var buffer2 bytes.Buffer
	var buffer3 bytes.Buffer

	multiWriter := io.MultiWriter(&buffer1, &buffer2, &buffer3)

	message := "This message will be written to multiple destinations."
	n, err := multiWriter.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to multi-writer: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Written %d bytes to multi-writer\n", n)
	fmt.Printf("Buffer 1: %s\n", Green(buffer1.String()))
	fmt.Printf("Buffer 2: %s\n", Yellow(buffer2.String()))
	fmt.Printf("Buffer 3: %s\n", Cyan(buffer3.String()))
	fmt.Println()
}

// limitedReaderExample demonstrates LimitedReader
func limitedReaderExample() {
	fmt.Println(SectionHeader("Limited Reader"))

	// Create a long string to read from
	longText := "This is a very long string that we will read with a limit. " +
		"The LimitedReader will stop reading after a certain number of bytes. " +
		"This is useful for preventing reading too much data from a source."

	reader := strings.NewReader(longText)

	// Create a limited reader that will only read 50 bytes
	limitedReader := &io.LimitedReader{
		R: reader,
		N: 50, // Limit to 50 bytes
	}

	fmt.Printf("Original text (%d bytes): %s\n", len(longText), Yellow(longText))
	fmt.Printf("Reading with limit of %d bytes\n", 50)

	// Read all data (will be limited to 50 bytes)
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		fmt.Printf("Error reading limited data: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Limited read result (%d bytes): %s\n", len(data), Green(string(data)))
	fmt.Printf("Remaining limit: %d bytes\n", limitedReader.N)

	// Try to read more (should return EOF)
	moreData := make([]byte, 10)
	n, err := limitedReader.Read(moreData)
	if err != nil {
		if err == io.EOF {
			fmt.Println(InfoText("Reached limit - EOF returned"))
		} else {
			fmt.Printf("Error reading more: %s\n", ErrorText(err.Error()))
		}
	} else {
		fmt.Printf("Read %d more bytes: %s\n", n, string(moreData[:n]))
	}
	fmt.Println()
}

// sectionReaderExample demonstrates SectionReader
func sectionReaderExample() {
	fmt.Println(SectionHeader("Section Reader"))

	// Create a string reader
	fullText := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	reader := strings.NewReader(fullText)

	fmt.Printf("Full text: %s\n", Yellow(fullText))
	fmt.Printf("Full text length: %d\n", len(fullText))

	// Create a section reader that reads from position 10 to 20
	sectionReader := io.NewSectionReader(reader, 10, 10)

	fmt.Printf("Creating section reader: offset=10, length=10\n")
	fmt.Printf("Section size: %d\n", sectionReader.Size())

	// Read the section
	sectionData, err := io.ReadAll(sectionReader)
	if err != nil {
		fmt.Printf("Error reading section: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Section content: %s\n", Green(string(sectionData)))

	// Seek within the section
	sectionReader.Seek(5, io.SeekStart)
	remainingData := make([]byte, 3)
	n, err := sectionReader.Read(remainingData)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading after seek: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("After seeking to position 5 in section, read %d bytes: %s\n",
		n, Cyan(string(remainingData[:n])))

	// Try to read beyond section
	sectionReader.Seek(0, io.SeekStart)
	largeBuffer := make([]byte, 20)
	n, err = sectionReader.Read(largeBuffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading large buffer: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Attempted to read 20 bytes, actually read %d: %s\n",
		n, Green(string(largeBuffer[:n])))
	fmt.Println()
}

// teeReaderExample demonstrates TeeReader
func teeReaderExample() {
	fmt.Println(SectionHeader("Tee Reader"))

	// Create source data
	sourceText := "This data will be read and simultaneously written to another destination."
	sourceReader := strings.NewReader(sourceText)

	// Create a buffer to capture the tee output
	var teeBuffer bytes.Buffer

	// Create a TeeReader that will write to teeBuffer as we read
	teeReader := io.TeeReader(sourceReader, &teeBuffer)

	fmt.Printf("Source text: %s\n", Yellow(sourceText))
	fmt.Println("Reading through TeeReader (data will be copied to tee buffer)...")

	// Read from the tee reader in chunks
	buffer := make([]byte, 20)
	var readData bytes.Buffer

	for {
		n, err := teeReader.Read(buffer)
		if n > 0 {
			readData.Write(buffer[:n])
			fmt.Printf("Read chunk: %s\n", Green(string(buffer[:n])))
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println(InfoText("Finished reading"))
				break
			}
			fmt.Printf("Error reading: %s\n", ErrorText(err.Error()))
			break
		}
	}

	fmt.Printf("Total data read: %s\n", Cyan(readData.String()))
	fmt.Printf("Tee buffer content: %s\n", Green(teeBuffer.String()))
	fmt.Printf("Tee buffer length: %d bytes\n", teeBuffer.Len())

	// Demonstrate that both contain the same data
	if readData.String() == teeBuffer.String() {
		fmt.Println(InfoText("✓ Read data and tee buffer contain identical content"))
	} else {
		fmt.Println(ErrorText("✗ Read data and tee buffer content differ"))
	}
	fmt.Println()
}

// Custom Writer that counts characters
type CountingWriter struct {
	CharCount  map[rune]int
	TotalBytes int
}

func (w *CountingWriter) Write(p []byte) (n int, err error) {
	if w.CharCount == nil {
		w.CharCount = make(map[rune]int)
	}

	for _, b := range p {
		w.CharCount[rune(b)]++
		w.TotalBytes++
	}

	return len(p), nil
}

// Custom Reader that repeats a string
type RepeatingReader struct {
	data  string
	count int
	pos   int
}

func (r *RepeatingReader) Read(p []byte) (n int, err error) {
	if r.count <= 0 || len(r.data) == 0 {
		return 0, io.EOF
	}

	for n < len(p) && r.count > 0 {
		p[n] = byte(r.data[r.pos])
		r.pos++
		n++
		if r.pos >= len(r.data) {
			r.pos = 0
			r.count--
			if r.count <= 0 {
				break
			}
		}
	}

	if n == 0 {
		return 0, io.EOF
	}
	return n, nil
}

// readerWriterInterfaces demonstrates custom Reader and Writer implementations
func readerWriterInterfaces() {
	fmt.Println(SectionHeader("Custom Reader and Writer Interfaces"))

	// Test custom reader
	fmt.Println("Testing custom RepeatingReader:")
	repeatingReader := &RepeatingReader{
		data:  "Go! ",
		count: 5,
	}

	data, err := io.ReadAll(repeatingReader)
	if err != nil {
		fmt.Printf("Error reading from custom reader: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Custom reader output: %s\n", Green(string(data)))

	// Test custom writer
	fmt.Println("Testing custom CountingWriter:")
	countingWriter := &CountingWriter{}

	text := "Hello, World! This is a test of the counting writer."
	n, err := countingWriter.Write([]byte(text))
	if err != nil {
		fmt.Printf("Error writing to custom writer: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Written %d bytes: %s\n", n, Yellow(text))
	fmt.Printf("Total bytes processed: %d\n", countingWriter.TotalBytes)
	fmt.Println("Character counts:")

	for char, count := range countingWriter.CharCount {
		if char == ' ' {
			fmt.Printf("  Space: %d\n", count)
		} else if char == '\n' {
			fmt.Printf("  Newline: %d\n", count)
		} else {
			fmt.Printf("  '%c': %d\n", char, count)
		}
	}

	// Demonstrate io.WriteString
	fmt.Println("\nTesting io.WriteString with custom writer:")
	countingWriter2 := &CountingWriter{}

	n, err = io.WriteString(countingWriter2, "Testing WriteString function")
	if err != nil {
		fmt.Printf("Error with WriteString: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("WriteString wrote %d bytes\n", n)
	fmt.Printf("Total bytes in writer: %d\n", countingWriter2.TotalBytes)
	fmt.Println()
}

// main function to run the examples
// func main() {
// 	fmt.Println(Bold(Blue("Go IO Package Examples")))
// 	fmt.Println(Bold(Blue("======================")))

// 	RunIOPackageExamples()

// 	fmt.Println(SectionHeader("Examples Complete"))
// 	fmt.Println(InfoText("All IO package examples have been demonstrated."))
// 	fmt.Println(InfoText("Key concepts covered:"))
// 	fmt.Println("• Basic Reader and Writer interfaces")
// 	fmt.Println("• String readers and builders")
// 	fmt.Println("• Copy operations (Copy, CopyN)")
// 	fmt.Println("• Buffer operations")
// 	fmt.Println("• Pipe operations with goroutines")
// 	fmt.Println("• Multi-reader and multi-writer")
// 	fmt.Println("• Limited reader for controlled reading")
// 	fmt.Println("• Section reader for reading portions")
// 	fmt.Println("• Tee reader for simultaneous read/write")
// 	fmt.Println("• Custom reader and writer implementations")
// }
