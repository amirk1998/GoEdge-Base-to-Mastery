// ioutil_examples.go
package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// RunIOUtilExamples - main function to run all IO/ioutil package examples
func RunIOUtilExamples() {
	fmt.Println(Subtitle("📁 IO/ioutil Package Examples"))
	fmt.Println(Yellow("⚠️  Note: io/ioutil is deprecated since Go 1.16, but still widely used"))
	fmt.Println()

	readFileExample()
	writeFileExample()
	readDirExample()
	tempFileExample()
	tempDirExample()
	readAllExample()
	nopCloserExample()
	discardExample()
}

// ReadFile Example
func readFileExample() {
	fmt.Println(Yellow("📌 ReadFile Operations:"))

	// Create a test file first
	testContent := `This is a test file for demonstrating ioutil.ReadFile.
It contains multiple lines of text.
Line 3: Numbers 123456789
Line 4: Special characters !@#$%^&*()
Final line with some content.`

	fileName := "test_read_file.txt"
	err := ioutil.WriteFile(fileName, []byte(testContent), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %v\n", err)
		return
	}
	defer os.Remove(fileName) // Cleanup

	// Read the entire file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File size: %d bytes\n", len(data))
	fmt.Printf("File content:\n%s\n", Green(string(data)))

	// Read non-existent file
	_, err = ioutil.ReadFile("non_existent_file.txt")
	if err != nil {
		fmt.Printf("Expected error for non-existent file: %s\n", Red(err.Error()))
	}
	fmt.Println()
}

// WriteFile Example
func writeFileExample() {
	fmt.Println(Yellow("📌 WriteFile Operations:"))

	// Write string content to file
	content1 := "Hello, this is content written using ioutil.WriteFile!"
	fileName1 := "write_test1.txt"

	err := ioutil.WriteFile(fileName1, []byte(content1), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	defer os.Remove(fileName1) // Cleanup

	fmt.Printf("Successfully wrote to: %s\n", Green(fileName1))

	// Verify by reading back
	readBack, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Printf("Error reading back: %v\n", err)
	} else {
		fmt.Printf("Read back: %s\n", Cyan(string(readBack)))
	}

	// Write binary data
	binaryData := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x42, 0x69, 0x6E, 0x61, 0x72, 0x79}
	fileName2 := "binary_test.bin"

	err = ioutil.WriteFile(fileName2, binaryData, 0644)
	if err != nil {
		fmt.Printf("Error writing binary file: %v\n", err)
	} else {
		fmt.Printf("Binary file written: %s\n", Yellow(fileName2))

		// Read and display binary data
		binRead, _ := ioutil.ReadFile(fileName2)
		fmt.Printf("Binary content: %x\n", binRead)
		fmt.Printf("As string: %s\n", Green(string(binRead)))
	}
	defer os.Remove(fileName2) // Cleanup

	// Write with different permissions
	restrictedContent := "This file has restricted permissions"
	restrictedFile := "restricted.txt"

	err = ioutil.WriteFile(restrictedFile, []byte(restrictedContent), 0400) // Read-only
	if err != nil {
		fmt.Printf("Error writing restricted file: %v\n", err)
	} else {
		fmt.Printf("Restricted file written: %s\n", Cyan(restrictedFile))

		// Check file permissions
		info, _ := os.Stat(restrictedFile)
		fmt.Printf("File permissions: %s\n", info.Mode().String())
	}
	defer os.Remove(restrictedFile) // Cleanup
	fmt.Println()
}

// ReadDir Example
func readDirExample() {
	fmt.Println(Yellow("📌 ReadDir Operations:"))

	// Create test directory structure
	testDir := "test_directory"
	os.Mkdir(testDir, 0755)
	defer os.RemoveAll(testDir) // Cleanup

	// Create some files and subdirectories
	files := []string{
		"file1.txt",
		"file2.log",
		"document.pdf",
	}

	for i, file := range files {
		filePath := filepath.Join(testDir, file)
		content := fmt.Sprintf("Content of %s (file %d)", file, i+1)
		ioutil.WriteFile(filePath, []byte(content), 0644)
	}

	// Create subdirectories
	subdirs := []string{"subdir1", "subdir2"}
	for _, subdir := range subdirs {
		os.Mkdir(filepath.Join(testDir, subdir), 0755)

		// Add file to subdirectory
		subFile := filepath.Join(testDir, subdir, "nested_file.txt")
		ioutil.WriteFile(subFile, []byte("Nested file content"), 0644)
	}

	// Read directory contents
	entries, err := ioutil.ReadDir(testDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	fmt.Printf("Directory '%s' contains %d entries:\n", Green(testDir), len(entries))

	for _, entry := range entries {
		entryType := "FILE"
		if entry.IsDir() {
			entryType = "DIR "
		}

		fmt.Printf("  [%s] %s (size: %d, mode: %s)\n",
			Yellow(entryType),
			entry.Name(),
			entry.Size(),
			Cyan(entry.Mode().String()))
	}

	// Read subdirectory
	subDirPath := filepath.Join(testDir, "subdir1")
	subEntries, err := ioutil.ReadDir(subDirPath)
	if err != nil {
		fmt.Printf("Error reading subdirectory: %v\n", err)
	} else {
		fmt.Printf("\nSubdirectory '%s' contains:\n", Green(subDirPath))
		for _, entry := range subEntries {
			fmt.Printf("  %s\n", entry.Name())
		}
	}
	fmt.Println()
}

// TempFile Example
func tempFileExample() {
	fmt.Println(Yellow("📌 TempFile Operations:"))

	// Create temporary file with default temp directory
	tempFile1, err := ioutil.TempFile("", "example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile1.Name()) // Cleanup
	defer tempFile1.Close()

	fmt.Printf("Created temp file: %s\n", Green(tempFile1.Name()))

	// Write content to temp file
	content := "This is temporary content for demonstration"
	_, err = tempFile1.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to temp file: %v\n", err)
	} else {
		fmt.Printf("Written to temp file: %s\n", Cyan(content))
	}

	// Create temp file in specific directory
	tempDir := "custom_temp_dir"
	os.Mkdir(tempDir, 0755)
	defer os.RemoveAll(tempDir) // Cleanup

	tempFile2, err := ioutil.TempFile(tempDir, "custom_temp_*.log")
	if err != nil {
		fmt.Printf("Error creating custom temp file: %v\n", err)
	} else {
		defer tempFile2.Close()
		defer os.Remove(tempFile2.Name())

		fmt.Printf("Custom temp file: %s\n", Yellow(tempFile2.Name()))

		// Write JSON-like content
		jsonContent := `{
  "timestamp": "2024-01-15T10:30:00Z",
  "level": "INFO",
  "message": "Temporary log entry",
  "data": {
    "user_id": 12345,
    "action": "login"
  }
}`
		tempFile2.WriteString(jsonContent)
		fmt.Printf("JSON content written to temp file\n")
	}

	// Multiple temp files with different patterns
	patterns := []string{"log_*.txt", "data_*.csv", "config_*.json"}

	fmt.Println(Bold("Creating multiple temp files:"))
	for i, pattern := range patterns {
		tf, err := ioutil.TempFile("", pattern)
		if err != nil {
			fmt.Printf("Error creating temp file %d: %v\n", i+1, err)
			continue
		}

		fmt.Printf("  %d. %s\n", i+1, tf.Name())
		tf.WriteString(fmt.Sprintf("Content for temp file %d", i+1))
		tf.Close()
		os.Remove(tf.Name()) // Immediate cleanup for demo
	}
	fmt.Println()
}

// TempDir Example
func tempDirExample() {
	fmt.Println(Yellow("📌 TempDir Operations:"))

	// Create temporary directory
	tempDir1, err := ioutil.TempDir("", "example_dir_*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir1) // Cleanup

	fmt.Printf("Created temp directory: %s\n", Green(tempDir1))

	// Create files in temp directory
	files := []struct {
		name, content string
	}{
		{"config.json", `{"app": "demo", "version": "1.0"}`},
		{"data.csv", "id,name,value\n1,item1,100\n2,item2,200"},
		{"readme.txt", "This is a temporary directory with various files"},
	}

	for _, file := range files {
		filePath := filepath.Join(tempDir1, file.name)
		err := ioutil.WriteFile(filePath, []byte(file.content), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file.name, err)
		} else {
			fmt.Printf("Created file: %s\n", Cyan(file.name))
		}
	}

	// Create nested directory structure
	nestedDir := filepath.Join(tempDir1, "nested", "deep", "structure")
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("Error creating nested structure: %v\n", err)
	} else {
		fmt.Printf("Created nested structure: %s\n", Yellow("nested/deep/structure"))

		// Add file to nested directory
		nestedFile := filepath.Join(nestedDir, "deep_file.txt")
		ioutil.WriteFile(nestedFile, []byte("File in deep nested directory"), 0644)
	}

	// List all contents recursively
	fmt.Println(Bold("Directory contents:"))
	err = filepath.Walk(tempDir1, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(tempDir1, path)
		if relPath == "." {
			return nil
		}

		indent := strings.Repeat("  ", strings.Count(relPath, string(filepath.Separator)))
		entryType := "📄"
		if info.IsDir() {
			entryType = "📁"
		}

		fmt.Printf("%s%s %s\n", indent, entryType, relPath)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

	// Create temp directory in custom location
	customBase := "custom_base"
	os.Mkdir(customBase, 0755)
	defer os.RemoveAll(customBase) // Cleanup

	tempDir2, err := ioutil.TempDir(customBase, "app_temp_*")
	if err != nil {
		fmt.Printf("Error creating custom temp dir: %v\n", err)
	} else {
		defer os.RemoveAll(tempDir2)
		fmt.Printf("Custom location temp dir: %s\n", Green(tempDir2))
	}
	fmt.Println()
}

// ReadAll Example
func readAllExample() {
	fmt.Println(Yellow("📌 ReadAll Operations:"))

	// ReadAll from string reader
	content := "This is content that will be read all at once using ioutil.ReadAll"
	reader := strings.NewReader(content)

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading all: %v\n", err)
		return
	}

	fmt.Printf("Read all data (%d bytes): %s\n", len(data), Green(string(data)))

	// ReadAll from file
	testFile := "readall_test.txt"
	fileContent := `Line 1: Introduction
Line 2: Body content with important information
Line 3: More detailed explanation
Line 4: Additional notes and references
Line 5: Conclusion and summary`

	err = ioutil.WriteFile(testFile, []byte(fileContent), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %v\n", err)
		return
	}
	defer os.Remove(testFile) // Cleanup

	// Open file and read all
	file, err := os.Open(testFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading all from file: %v\n", err)
	} else {
		lines := strings.Split(string(fileData), "\n")
		fmt.Printf("File read completely (%d lines):\n", len(lines))
		for i, line := range lines {
			fmt.Printf("  %d: %s\n", i+1, Cyan(line))
		}
	}

	// ReadAll with limited content
	limitedContent := "Short content"
	limitedReader := strings.NewReader(limitedContent)

	limitedData, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		fmt.Printf("Error reading limited content: %v\n", err)
	} else {
		fmt.Printf("Limited content: %s\n", Yellow(string(limitedData)))
	}
	fmt.Println()
}

// NopCloser Example
func nopCloserExample() {
	fmt.Println(Yellow("📌 NopCloser Operations:"))

	// Create a reader that doesn't implement io.ReadCloser
	content := "This content comes from a reader that will be wrapped with NopCloser"
	reader := strings.NewReader(content)

	// Wrap with NopCloser to make it a ReadCloser
	readCloser := ioutil.NopCloser(reader)

	fmt.Printf("Original reader type: %T\n", reader)
	fmt.Printf("NopCloser type: %T\n", readCloser)

	// Read from the NopCloser
	data, err := ioutil.ReadAll(readCloser)
	if err != nil {
		fmt.Printf("Error reading from NopCloser: %v\n", err)
	} else {
		fmt.Printf("Read from NopCloser: %s\n", Green(string(data)))
	}

	// Close the NopCloser (this is a no-op)
	err = readCloser.Close()
	if err != nil {
		fmt.Printf("Error closing NopCloser: %v\n", err)
	} else {
		fmt.Printf("NopCloser closed successfully (no-op)\n")
	}

	// Practical example: HTTP response body simulation
	responseBody := "HTTP response body content that needs to be a ReadCloser"
	responseReader := strings.NewReader(responseBody)
	httpBodyCloser := ioutil.NopCloser(responseReader)

	// Simulate reading HTTP response
	fmt.Println(Bold("Simulated HTTP response reading:"))
	bodyContent, err := ioutil.ReadAll(httpBodyCloser)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	} else {
		fmt.Printf("Response body: %s\n", Cyan(string(bodyContent)))
	}

	// Always close the response body (even though it's a no-op)
	httpBodyCloser.Close()
	fmt.Printf("Response body closed\n")
	fmt.Println()
}

// Discard Example
func discardExample() {
	fmt.Println(Yellow("📌 Discard Operations:"))

	// Demonstrate ioutil.Discard
	content := "This content will be discarded - it goes nowhere!"

	// Write to discard
	n, err := ioutil.Discard.Write([]byte(content))
	if err != nil {
		fmt.Printf("Error writing to discard: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes to discard (content: %s)\n",
			n, Dim(content))
	}

	// Copy to discard (useful for measuring read speed)
	largeContent := strings.Repeat("This is repeated content for performance testing. ", 1000)
	reader := strings.NewReader(largeContent)

	fmt.Printf("Original content size: %d bytes\n", len(largeContent))

	// Copy all content to discard
	discardedBytes, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		fmt.Printf("Error copying to discard: %v\n", err)
	} else {
		fmt.Printf("Discarded %s bytes\n", Green(fmt.Sprintf("%d", discardedBytes)))
	}

	// Practical example: draining a reader without storing content
	tempFile, err := ioutil.TempFile("", "discard_test_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write content to file
	fileContent := "Content in file that we want to drain without reading into memory"
	tempFile.WriteString(fileContent)
	tempFile.Seek(0, 0) // Reset to beginning

	// Drain the file content
	drainedBytes, err := io.Copy(ioutil.Discard, tempFile)
	if err != nil {
		fmt.Printf("Error draining file: %v\n", err)
	} else {
		fmt.Printf("Drained %s bytes from file to discard\n",
			Yellow(fmt.Sprintf("%d", drainedBytes)))
	}

	// Multiple writes to discard
	fmt.Println(Bold("Multiple writes to discard:"))
	messages := []string{
		"Message 1: This will be discarded",
		"Message 2: This will also be discarded",
		"Message 3: All content goes to the void",
	}

	totalDiscarded := 0
	for i, msg := range messages {
		n, err := ioutil.Discard.Write([]byte(msg))
		if err != nil {
			fmt.Printf("Error in write %d: %v\n", i+1, err)
		} else {
			totalDiscarded += n
			fmt.Printf("  Write %d: %d bytes\n", i+1, n)
		}
	}

	fmt.Printf("Total bytes discarded: %s\n", Cyan(fmt.Sprintf("%d", totalDiscarded)))
	fmt.Println()
}
