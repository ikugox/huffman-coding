package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func encodeFile(fileName string) {
	inputPath := "./input/" + fileName
	outputPath := "./encoded/" + fileName[:len(fileName)-3] + "huff"
	fmt.Printf("Found .txt file: %s\n", inputPath)

	data, freq := processFile(inputPath)
	root := generateHuffmanTree(data, freq)
	codes := generateHuffmanCodes(root)

	writeEncodedFile(inputPath, outputPath, codes)
	fmt.Printf("Encoding to .huff file: %s\n", outputPath)
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

func encodeFilesFromDir(dirName string) {
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
		encodeFile(file.Name())
	}
}
