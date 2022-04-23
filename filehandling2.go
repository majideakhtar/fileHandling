package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

/*
The program above reads the content of the file path passed from the command line. Run this program using the command

filehandling -fpath=/path-of-file/test.txt
Please replace /path-of-file/ with the absolute path of test.txt. For example, in my case, I ran the command

filehandling --fpath=/home/naveen/Documents/filehandling/test.txt
and the program printed.

Contents of file: Hello World. Welcome to file handling in Go.
*/
