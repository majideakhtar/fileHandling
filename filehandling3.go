// bundle the text file along with our binary

// install->  go get -u github.com/gobuffalo/packr/v2/...

/*
packr converts static files such as .txt to .go files which are then embedded directly into the binary. Packer is intelligent enough to fetch the static files from disk rather than from the binary during development. This prevents the need for re-compilation during development when only static files change.
*/

package main

import (
	"fmt"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	// we are creating a New Box named box. A box represents a folder whose contents will be embedded in the binary
	box := packr.New("fileBox", "../filehandling")
	// we read the contents of the file using the FindString method and print it
	data, err := box.FindString("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", data)
}
