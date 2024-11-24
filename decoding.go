package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func decodeFile(fileName, outputName string) {
	if strings.HasSuffix(outputName, "/") {
		lastSlash := strings.LastIndex(fileName, "/")
		outputName += fileName[lastSlash+1:]
	}
	if !strings.HasSuffix(fileName, ".huff") {
		fileName += ".huff"
	}
	if strings.HasSuffix(outputName, ".huff") {
		outputName = outputName[:len(outputName)-5]
	}
	if !strings.HasSuffix(outputName, ".txt") {
		outputName += ".txt"
	}

	fmt.Printf("Found .huff file: %s\n", fileName)

	code, err := readMetadata(fileName)
	if err != nil {
		fmt.Printf("Error reading metadata (huffman code): %s\n", err)
		return
	}

	err = readEncodedFile(fileName, outputName, code)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
	fmt.Printf("Decoding to .txt file: %s\n", outputName)
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
					fmt.Printf("Error writting rune %c: %s\n", char, err)
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

func decodeFilesFromDir(dirName, outputDirName string) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".huff") {
			continue
		}
		decodeFile(dirName+"/"+file.Name(), outputDirName+"/"+file.Name())
	}
}
