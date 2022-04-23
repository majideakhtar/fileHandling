package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

/*
 In line no. 24, we create a new scanner using the file. The scan() method in line no. 25 reads the next line of the file and the string that is read will be available through the text() method.
 After Scan returns false, the Err() method will return any error that occurred during scanning. If the error is End of File, Err() will return nil.
*/
