package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func decodeFile(fileName string) {
	inputPath := "./encoded/" + fileName
	outputPath := "./decoded/" + fileName[:len(fileName)-4] + "txt"
	fmt.Printf("Found .huff file: %s\n", inputPath)

	code, err := readMetadata(inputPath)
	if err != nil {
		fmt.Println("Error reading metadata (huffman code):", err)
		return
	}

	readEncodedFile(inputPath, outputPath, code)
	fmt.Printf("Encoding to .huff file: %s\n", outputPath)
	// PrintHuffmanCodes(codes)
}

func readEncodedFile(inputPath, outputPath string, code map[rune]string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	codeInverted := make(map[string]rune)
	for r, code := range code {
		codeInverted[code] = r
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	readBitByBit(reader, writer, codeInverted)

	return writer.Flush()
}

func readBitByBit(reader *bufio.Reader, writer *bufio.Writer, codeInverted map[string]rune) {
	reader.ReadString('\n')
	var currentBits string
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		for i := 0; i < 8; i++ {
			bit := (b >> (7 - i)) & 1
			if bit == 1 {
				currentBits += "1"
			} else {
				currentBits += "0"
			}
			if char, exists := codeInverted[currentBits]; exists {
				_, err := writer.WriteRune(char)
				if err != nil {
					fmt.Println("Error writting rune:", err)
					return
				}
				currentBits = ""
			}
		}
	}
}

func readMetadata(filePath string) (map[rune]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var code map[rune]string
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&code)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func decodeFilesFromDir(dirName string) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".huff") {
			continue
		}
		decodeFile(file.Name())
	}
}
