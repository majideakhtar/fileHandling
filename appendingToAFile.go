package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644) // In line no. 9 of the program above, we open the file in append and write only mode.
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine) //After the file is opened successfully, we add a new line to the file in line no. 15. This program will print file appended successfully.
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
