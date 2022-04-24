package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt") // The create function in line no. 9 of the program above creates a file named test.txt. If a file with that name already exists, then the create function truncates the file. This function returns a File descriptor.
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World") // In line no 14, we write the string Hello World to the file using the WriteString method. This method returns the number of bytes written and error if any.
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close() // we close the file in line no. 21.
	if err != nil {
		fmt.Println(err)
		return
	}
}
