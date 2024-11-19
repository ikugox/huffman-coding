package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	digits := [6]int{1e2, 1e3, 1e4, 1e5, 1e6, 1e7}

	for i, size := range digits {
		text := generateText(size)
		fileName := "ascii" + fmt.Sprintf("%d", i+1) + ".txt"
		writeToFile(fileName, text)
	}
}

func generateText(size int) []byte {
	char := make([]byte, size)
	for i := 0; i < size; i++ {
		num := rand.Intn(26)
		char[i] = 'a' + byte(num)
	}
	return char
}

func writeToFile(fileName string, text []byte) {
	err := os.WriteFile(fileName, text, 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Generated text to file %s\n", fileName)
}
