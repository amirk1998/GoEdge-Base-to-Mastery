// file_io.go
package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogEntry represents a log entry structure
type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// CustomWriter implements io.Writer interface
type CustomWriter struct {
	prefix string
}

func (cw CustomWriter) Write(p []byte) (n int, err error) {
	prefixed := fmt.Sprintf("[%s] %s", cw.prefix, string(p))
	return fmt.Print(prefixed)
}

// MultiWriter writes to multiple writers simultaneously
type MultiWriter struct {
	writers []io.Writer
}

func NewMultiWriter(writers ...io.Writer) *MultiWriter {
	return &MultiWriter{writers: writers}
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	for _, writer := range mw.writers {
		n, err = writer.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

// FileProcessor demonstrates various file processing patterns
type FileProcessor struct {
	inputFile  string
	outputFile string
}

func NewFileProcessor(input, output string) *FileProcessor {
	return &FileProcessor{
		inputFile:  input,
		outputFile: output,
	}
}

// ProcessLines processes file line by line
func (fp *FileProcessor) ProcessLines(processor func(string) string) error {
	inputFile, err := os.Open(fp.inputFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(fp.outputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processor(line)
		_, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

// RunFileIOExamples - main function to run all File I/O examples
func RunFileIOExamples() {
	basicFileOperationsExample()
	readerWriterInterfaceExample()
	bufferedIOExample()
	fileProcessingExample()
	csvFileExample()
	binaryFileExample()
	customReaderWriterExample()
	streamingExample()
	fileIOErrorHandlingExample()
	advancedFileOperationsExample()
}

// Basic file operations
func basicFileOperationsExample() {
	fmt.Println(Subtitle("📁 Basic File Operations"))

	// Create a temporary file
	tempFile := "temp_example.txt"
	content := "Hello, World!\nThis is a test file.\nGolang file operations are powerful!"

	// Write to file
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		log.Printf("Error writing file: %v", err)
		return
	}
	fmt.Printf("Created file: %s\n", tempFile)

	// Read entire file
	data, err := os.ReadFile(tempFile)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}
	fmt.Printf("File content:\n%s\n", string(data))

	// Get file info
	info, err := os.Stat(tempFile)
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return
	}
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("File mode: %v\n", info.Mode())
	fmt.Printf("Modified time: %v\n", info.ModTime())

	// Clean up
	os.Remove(tempFile)
	fmt.Println("Temporary file cleaned up")
	fmt.Println()
}

// Reader/Writer interface examples
func readerWriterInterfaceExample() {
	fmt.Println(Subtitle("🔄 Reader/Writer Interface Examples"))

	// String Reader
	stringReader := strings.NewReader("Hello from string reader!")
	buffer := make([]byte, 10)

	fmt.Println(Bold("Reading from string reader:"))
	for {
		n, err := stringReader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading: %v", err)
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}

	// Bytes Buffer (implements both Reader and Writer)
	var bytesBuffer bytes.Buffer

	fmt.Println(Bold("\nUsing bytes.Buffer:"))
	bytesBuffer.WriteString("First line\n")
	bytesBuffer.WriteString("Second line\n")
	bytesBuffer.WriteString("Third line\n")

	// Read from buffer
	fmt.Printf("Buffer contents:\n%s", bytesBuffer.String())

	// Copy from one reader to another writer
	sourceReader := strings.NewReader("Data to copy")
	var destBuffer bytes.Buffer

	n, err := io.Copy(&destBuffer, sourceReader)
	if err != nil {
		log.Printf("Error copying: %v", err)
	} else {
		fmt.Printf("Copied %d bytes: %s\n", n, destBuffer.String())
	}

	// Custom writer example
	customWriter := CustomWriter{prefix: "CUSTOM"}
	customWriter.Write([]byte("Hello from custom writer!\n"))

	fmt.Println()
}

// Buffered I/O examples
func bufferedIOExample() {
	fmt.Println(Subtitle("🚀 Buffered I/O Examples"))

	// Create a test file
	tempFile := "buffered_test.txt"
	content := []string{
		"Line 1: Introduction",
		"Line 2: Buffered I/O is efficient",
		"Line 3: It reduces system calls",
		"Line 4: Perfect for large files",
		"Line 5: Conclusion",
	}

	// Write with buffered writer
	file, err := os.Create(tempFile)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // Important: flush before closing

	fmt.Println(Bold("Writing with buffered writer:"))
	for i, line := range content {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			log.Printf("Error writing line %d: %v", i, err)
			return
		}
		fmt.Printf("Wrote: %s\n", line)
	}

	// Manually flush to ensure data is written
	writer.Flush()

	// Read with buffered reader
	file, err = os.Open(tempFile)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fmt.Println(Bold("\nReading with buffered reader:"))
	lineNum := 1
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				fmt.Printf("Line %d: %s", lineNum, line)
			}
			break
		}
		if err != nil {
			log.Printf("Error reading line: %v", err)
			break
		}
		fmt.Printf("Line %d: %s", lineNum, line)
		lineNum++
	}

	// Scanner example (more convenient for line-by-line reading)
	file, err = os.Open(tempFile)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(Bold("\nUsing Scanner:"))
	lineNum = 1
	for scanner.Scan() {
		fmt.Printf("Scanned line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v", err)
	}

	// Clean up
	os.Remove(tempFile)
	fmt.Println()
}

// File processing example
func fileProcessingExample() {
	fmt.Println(Subtitle("⚙️ File Processing Example"))

	// Create input file
	inputFile := "input.txt"
	inputContent := `apple,10,2.50
banana,15,1.20
cherry,8,3.00
date,12,2.80
elderberry,5,4.50`

	err := os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		log.Printf("Error creating input file: %v", err)
		return
	}

	// Process file
	outputFile := "output.txt"
	processor := NewFileProcessor(inputFile, outputFile)

	// Define processing function
	processLine := func(line string) string {
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			return fmt.Sprintf("Product: %s, Stock: %s units, Price: $%s",
				strings.Title(parts[0]), parts[1], parts[2])
		}
		return line
	}

	err = processor.ProcessLines(processLine)
	if err != nil {
		log.Printf("Error processing file: %v", err)
		return
	}

	// Read and display result
	result, err := os.ReadFile(outputFile)
	if err != nil {
		log.Printf("Error reading output file: %v", err)
		return
	}

	fmt.Printf("Input file content:\n%s\n\n", inputContent)
	fmt.Printf("Processed output:\n%s\n", string(result))

	// Clean up
	os.Remove(inputFile)
	os.Remove(outputFile)
	fmt.Println()
}

// CSV file example
func csvFileExample() {
	fmt.Println(Subtitle("📊 CSV File Handling"))

	// Create CSV content
	csvContent := `Name,Age,City,Salary
John Doe,30,New York,75000
Jane Smith,25,Los Angeles,65000
Bob Johnson,35,Chicago,80000
Alice Brown,28,Boston,70000`

	csvFile := "employees.csv"
	err := os.WriteFile(csvFile, []byte(csvContent), 0644)
	if err != nil {
		log.Printf("Error creating CSV file: %v", err)
		return
	}

	// Read and parse CSV
	file, err := os.Open(csvFile)
	if err != nil {
		log.Printf("Error opening CSV file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read header
	if scanner.Scan() {
		header := scanner.Text()
		fmt.Printf("CSV Header: %s\n", header)

		// Process each data row
		fmt.Println(Bold("CSV Data:"))
		rowNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, ",")
			if len(fields) >= 4 {
				fmt.Printf("Row %d: Name=%s, Age=%s, City=%s, Salary=$%s\n",
					rowNum, fields[0], fields[1], fields[2], fields[3])
				rowNum++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading CSV: %v", err)
	}

	// Clean up
	os.Remove(csvFile)
	fmt.Println()
}

// Binary file example
func binaryFileExample() {
	fmt.Println(Subtitle("🔢 Binary File Operations"))

	binaryFile := "binary_data.bin"

	// Create binary data
	var buffer bytes.Buffer

	// Write different data types
	data := []interface{}{
		int32(42),
		float64(3.14159),
		[]byte("Hello Binary"),
		int64(time.Now().Unix()),
	}

	// Write to buffer (in real app, use encoding/binary for proper binary format)
	for _, item := range data {
		switch v := item.(type) {
		case int32:
			buffer.Write([]byte(fmt.Sprintf("%d|", v)))
		case float64:
			buffer.Write([]byte(fmt.Sprintf("%.5f|", v)))
		case []byte:
			buffer.Write(v)
			buffer.Write([]byte("|"))
		case int64:
			buffer.Write([]byte(fmt.Sprintf("%d|", v)))
		}
	}

	// Write binary data to file
	err := os.WriteFile(binaryFile, buffer.Bytes(), 0644)
	if err != nil {
		log.Printf("Error writing binary file: %v", err)
		return
	}

	// Read binary data
	binaryData, err := os.ReadFile(binaryFile)
	if err != nil {
		log.Printf("Error reading binary file: %v", err)
		return
	}

	fmt.Printf("Binary file size: %d bytes\n", len(binaryData))
	fmt.Printf("Binary content (as string): %s\n", string(binaryData))
	fmt.Printf("Binary content (as hex): %x\n", binaryData)

	// Process binary data with io.Reader
	reader := bytes.NewReader(binaryData)
	chunk := make([]byte, 8)

	fmt.Println(Bold("Reading binary data in chunks:"))
	chunkNum := 1
	for {
		n, err := reader.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading chunk: %v", err)
			break
		}
		fmt.Printf("Chunk %d (%d bytes): %s\n", chunkNum, n, string(chunk[:n]))
		chunkNum++
	}

	// Clean up
	os.Remove(binaryFile)
	fmt.Println()
}

// Custom reader/writer example
// Custom reader that converts text to uppercase
// Move UppercaseReader type and its method outside the function

type UppercaseReader struct {
	reader io.Reader
}

func (ur *UppercaseReader) Read(p []byte) (n int, err error) {
	n, err = ur.reader.Read(p)
	if err != nil {
		return n, err
	}

	// Convert to uppercase
	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = p[i] - 'a' + 'A'
		}
	}

	return n, nil
}

// Custom writer that adds line numbers
// Move LineNumberWriter type and its method outside the function

type LineNumberWriter struct {
	writer    io.Writer
	lineCount int
}

func (lnw *LineNumberWriter) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	var output strings.Builder

	for i, line := range lines {
		if i == len(lines)-1 && line == "" {
			break // Don't add number to final empty line
		}
		lnw.lineCount++
		output.WriteString(fmt.Sprintf("%3d: %s\n", lnw.lineCount, line))
	}

	return lnw.writer.Write([]byte(output.String()))
}

func customReaderWriterExample() {
	fmt.Println(Subtitle("🎨 Custom Reader/Writer Implementation"))

	// Test custom reader
	originalText := "Hello, World! This text will be converted to uppercase."
	uppercaseReader := &UppercaseReader{reader: strings.NewReader(originalText)}

	var result bytes.Buffer
	_, err := io.Copy(&result, uppercaseReader)
	if err != nil {
		log.Printf("Error copying: %v", err)
		return
	}

	fmt.Printf("Original text: %s\n", originalText)
	fmt.Printf("Uppercase text: %s\n", result.String())

	// Test custom writer
	var lineNumberedOutput bytes.Buffer
	lineWriter := &LineNumberWriter{writer: &lineNumberedOutput}

	sampleText := "First line\nSecond line\nThird line\nFourth line"
	lineWriter.Write([]byte(sampleText))

	fmt.Printf("\nOriginal text:\n%s\n", sampleText)
	fmt.Printf("Line numbered text:\n%s", lineNumberedOutput.String())

	// Multi-writer example
	var buffer1, buffer2 bytes.Buffer
	multiWriter := NewMultiWriter(&buffer1, &buffer2)

	multiWriter.Write([]byte("This text goes to multiple writers!"))

	fmt.Printf("Buffer 1: %s\n", buffer1.String())
	fmt.Printf("Buffer 2: %s\n", buffer2.String())
	fmt.Println()
}

// Streaming example
func streamingExample() {
	fmt.Println(Subtitle("🌊 Streaming Data Processing"))

	// Create a large dataset
	dataFile := "large_dataset.txt"

	// Write large dataset
	file, err := os.Create(dataFile)
	if err != nil {
		log.Printf("Error creating dataset: %v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Generate sample data
	for i := 1; i <= 1000; i++ {
		line := fmt.Sprintf("Record %d: Value=%d, Timestamp=%s\n",
			i, i*10, time.Now().Add(time.Duration(i)*time.Second).Format("2006-01-02 15:04:05"))
		writer.WriteString(line)
	}
	writer.Flush()

	// Stream process the data
	file, err = os.Open(dataFile)
	if err != nil {
		log.Printf("Error opening dataset: %v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fmt.Println(Bold("Processing stream (showing first 10 records):"))

	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading stream: %v", err)
			break
		}

		// Process line (example: extract record number)
		if strings.Contains(line, "Record") {
			count++
			if count <= 10 {
				fmt.Printf("Processed: %s", line)
			}
		}
	}

	fmt.Printf("Total records processed: %d\n", count)

	// Clean up
	os.Remove(dataFile)
	fmt.Println()
}

// Error handling example
func fileIOErrorHandlingExample() {
	fmt.Println(Subtitle("🚨 Error Handling in File Operations"))

	// Test various error scenarios

	// 1. File not found
	fmt.Println(Bold("1. File not found error:"))
	_, err := os.Open("nonexistent_file.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if os.IsNotExist(err) {
			fmt.Println("   This is a 'file not found' error")
		}
	}

	// 2. Permission denied (try to write to read-only file)
	fmt.Println(Bold("2. Permission handling:"))
	readOnlyFile := "readonly.txt"
	os.WriteFile(readOnlyFile, []byte("read only content"), 0444) // Read-only

	err = os.WriteFile(readOnlyFile, []byte("trying to overwrite"), 0644)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if os.IsPermission(err) {
			fmt.Println("   This is a permission error")
		}
	}

	// 3. Proper resource cleanup with defer
	fmt.Println(Bold("3. Resource cleanup example:"))
	func() {
		file, err := os.Create("temp_cleanup.txt")
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("Error closing file: %v\n", err)
			}
			os.Remove("temp_cleanup.txt")
			fmt.Println("   File cleaned up successfully")
		}()

		_, err = file.WriteString("Temporary content")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}

		fmt.Println("   File operations completed successfully")
	}()

	// 4. Handling EOF correctly
	fmt.Println(Bold("4. EOF handling:"))
	reader := strings.NewReader("Short content")
	buffer := make([]byte, 20) // Larger than content

	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Unexpected error: %v\n", err)
	} else {
		fmt.Printf("   Read %d bytes: %s\n", n, string(buffer[:n]))
		if err == io.EOF {
			fmt.Println("   Reached end of file")
		}
	}

	// Clean up
	os.Remove(readOnlyFile)
	fmt.Println()
}

// Advanced file operations
func advancedFileOperationsExample() {
	fmt.Println(Subtitle("🔧 Advanced File Operations"))

	// 1. Working with directories
	fmt.Println(Bold("1. Directory operations:"))

	testDir := "test_directory"
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		log.Printf("Error creating directory: %v", err)
		return
	}

	// Create files in directory
	for i := 1; i <= 3; i++ {
		filename := filepath.Join(testDir, fmt.Sprintf("file_%d.txt", i))
		content := fmt.Sprintf("This is file number %d", i)
		os.WriteFile(filename, []byte(content), 0644)
	}

	// List directory contents
	entries, err := os.ReadDir(testDir)
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}

	fmt.Printf("Directory contents:\n")
	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("  %s (%d bytes)\n", entry.Name(), info.Size())
	}

	// 2. File copying
	fmt.Println(Bold("2. File copying:"))

	sourceFile := filepath.Join(testDir, "file_1.txt")
	destFile := filepath.Join(testDir, "file_1_copy.txt")

	err = copyFile(sourceFile, destFile)
	if err != nil {
		log.Printf("Error copying file: %v", err)
	} else {
		fmt.Printf("Successfully copied %s to %s\n", sourceFile, destFile)
	}

	// 3. File seeking
	fmt.Println(Bold("3. File seeking:"))

	file, err := os.Open(sourceFile)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// Read first 5 bytes
	buffer := make([]byte, 5)
	n, err := file.Read(buffer)
	if err != nil {
		log.Printf("Error reading: %v", err)
		return
	}
	fmt.Printf("First 5 bytes: %s\n", string(buffer[:n]))

	// Seek to position 10
	offset, err := file.Seek(10, io.SeekStart)
	if err != nil {
		log.Printf("Error seeking: %v", err)
		return
	}
	fmt.Printf("Seeked to position: %d\n", offset)

	// Read next 5 bytes
	n, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error reading after seek: %v", err)
		return
	}
	fmt.Printf("Next 5 bytes: %s\n", string(buffer[:n]))

	// 4. Temporary files
	fmt.Println(Bold("4. Temporary files:"))

	tempFile, err := os.CreateTemp("", "example_*.txt")
	if err != nil {
		log.Printf("Error creating temp file: %v", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	tempFile.WriteString("This is temporary content")

	// Clean up test directory
	os.RemoveAll(testDir)
	fmt.Println("Test directory cleaned up")
	fmt.Println()
}

// Helper function to copy files
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
