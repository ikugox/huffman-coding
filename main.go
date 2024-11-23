package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	encodeFilesFromDir("./input")
	decodeFilesFromDir("./encoded")
}

func processFile(filePath string) (data []rune, freq []int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return nil, nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	hashmap := make(map[rune]int)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file:", err)
			return nil, nil
		}
		hashmap[char]++
	}
	var arr []rune
	var occurances []int
	for char, occurance := range hashmap {
		arr = append(arr, char)
		occurances = append(occurances, occurance)
	}
	return arr, occurances
}
