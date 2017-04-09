package index

import (
	"os"
	"bufio"
	"fmt"
)

func ReadIndex(fileName string) (map[string][]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
	}

	return nil, nil
}

func UpdateIndex(fileName string, index map[string][]string) (map[string][]string, error) {
	return nil, nil
}

func WriteIndex(indexFile string, index map[string][]string) {

}

func Search(index map[string][]string, word string) []string {
	return nil
}