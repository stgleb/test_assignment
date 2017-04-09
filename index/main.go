package main

import (
	"flag"
)

func ReadIndex(fileName string) (map[string][]string) {
	return nil
}

func UpdateIndex(fileName string, index map[string][]string) map[string][]string {
	return nil
}

func WriteIndex(indexFile string, index map[string][]string) {

}

func main() {
	flag.Parse()
	fileName := *flag.String("file", "file.txt", "File name")
	indexFile := *flag.String("index", "index.dat", "Index file name")
	index := ReadIndex(indexFile)
	index = UpdateIndex(fileName, index)
	WriteIndex(indexFile, index)
}
