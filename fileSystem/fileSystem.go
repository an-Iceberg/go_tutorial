package fileSystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Go() {
	fmt.Println(strings.ToUpper("  File System"))

	// Reading a file directly
	data, error := os.ReadFile("example.txt")
	if error != nil { fmt.Println(error.Error()) }
	fmt.Println(data)
	fmt.Println(string(data))

	// Obtaining a file object
	file, error := os.Open("example.txt")
	if error != nil { fmt.Println(error.Error()) }

	// Read some bytes from the beginning of the file
	bytes := make([]byte, 5)
	amount, error := file.Read(bytes)
	if error != nil { fmt.Println(error.Error()) }
	fmt.Printf("%d bytes: %s\n", amount, string(bytes))

	file.Close()

	// Using a buffered reader
	file, error = os.Open("example.txt")
	if error != nil { fmt.Println(error.Error()) }
	bufferedReader := bufio.NewReader(file)

	bytes = make([]byte, 11)

	amount, error = bufferedReader.Read(bytes)
	if error != nil { fmt.Println(error.Error()) }

	fmt.Printf("%d bytes read: %s\n", amount, bytes)

	file.Close()

	// Splitting the buffered reader
	file, error = os.Open("example.txt")
	if error != nil { fmt.Println(error.Error()) }
	bufferedScanner := bufio.NewScanner(file)
	bufferedScanner.Split(bufio.ScanBytes)

	// Prints each byte to the console
	for bufferedScanner.Scan() {
		fmt.Print(bufferedScanner.Text())
	}

	file.Close()

  // Writing to a file
  {
    // Creating the file (if it already exists, it will be reset)
    file, error := os.Create("example_text_file.txt")
    if error != nil {
      fmt.Println(error.Error())
      os.Exit(-1)
    }
    fmt.Println("Opened file \"example_text_file.txt\"")

    number, error := file.WriteString(fmt.Sprintf("Hello World, this is a decimal number: %[1]d, and here is the hexadecimal representation of it: %[1]x\n", 75))
    if error != nil {
      fmt.Println(error.Error())
      os.Exit(-1)
    }
    fmt.Printf("Wrote %d bytes to file\n", number)

    file.Close()
    fmt.Println("Closed file")
  }
}
