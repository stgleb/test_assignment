package main

import (
	"flag"
	"fmt"
	full_text "test_assignment"
)

func main() {
	flag.Parse()
	fileName := *flag.String("file", "file.txt", "File name")
	indexFile := *flag.String("index", "index.txt", "Index file name")
	index, err := full_text.ReadIndex(indexFile)

	if err != nil {
		fmt.Print(err)
	}

	index, err = full_text.UpdateIndex(fileName, index)

	if err != nil {
		fmt.Print(err)
	}

	full_text.WriteIndex(indexFile, index)
}
