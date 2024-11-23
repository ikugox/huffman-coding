package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func encodeFile(fileName, outputName string) {
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}
	if strings.HasSuffix(outputName, ".txt") {
		outputName = outputName[:len(outputName)-4]
	}
	if !strings.HasSuffix(outputName, ".huff") {
		outputName += ".huff"
	}

	fmt.Printf("Found .txt file: %s\n", fileName)

	data, freq := processFile(fileName)
	if data == nil || freq == nil {
		return
	}
	root := generateHuffmanTree(data, freq)
	codes := generateHuffmanCodes(root)

	writeEncodedFile(fileName, outputName, codes)
	fmt.Printf("Encoding to .huff file: %s\n", outputName)
	// PrintHuffmanCodes(codes)
}

func writeEncodedFile(inputPath, outputPath string, code map[rune]string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	reader := bufio.NewReader(inputFile)

	metadata, err := json.Marshal(code)
	if err != nil {
		return err
	}
	_, err = writer.Write(metadata)
	if err != nil {
		return err
	}
	writer.WriteByte('\n')

	writeBitByBit(reader, writer, code)

	return writer.Flush()
}

func writeBitByBit(reader *bufio.Reader, writer *bufio.Writer, code map[rune]string) {
	var bitBuffer byte
	var bitCount int
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file:", err)
			return
		}
		code := code[char]
		for _, bit := range code {
			if bit == '1' {
				bitBuffer |= (1 << (7 - bitCount))
			}
			bitCount++

			if bitCount == 8 {
				writer.WriteByte(bitBuffer)
				bitBuffer = 0
				bitCount = 0
			}
		}
	}
	if bitCount > 0 {
		writer.WriteByte(bitBuffer)
	}
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

func encodeFilesFromDir(dirName, outputDirName string) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".txt") {
			continue
		}
		encodeFile(dirName+"/"+file.Name(), outputDirName+"/"+file.Name())
	}
}
