package fileSystem

import (
	"bufio"
	"fmt"
	"os"
)

func Go() {
	fmt.Println("  File System")

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
}
