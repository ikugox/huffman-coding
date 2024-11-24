package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if !(len(os.Args) == 4 || len(os.Args) == 6) {
		running := os.Args[0]
		if strings.Contains(running, "/tmp/") {
			running = "go run *.go"
		}
		fmt.Println("Usage: " + running + " <encode|decode> <-f|-d> <file/directory> [-o <output>]")
		return
	}
	coding, fd, fdName, outputName := os.Args[1], os.Args[2], os.Args[3], os.Args[3]
	if len(os.Args) == 6 {
		outputName = os.Args[5]
	}
	if _, err := os.Stat(outputName); err != nil {
		if err := os.Mkdir(outputName, os.FileMode(0755)); err != nil {
			fmt.Printf("Error with <output> option %s :%s\n", outputName, err)
			return
		}
		fmt.Printf("Created folder named: %s\n", outputName)
	}
	if coding == "encode" {
		if fd == "-f" {
			encodeFile(fdName, outputName)
		} else if fd == "-d" {
			encodeFilesFromDir(fdName, outputName)
		} else {
			fmt.Printf("Invalid option: %s\n", fd)
		}
	} else if coding == "decode" {
		if fd == "-f" {
			decodeFile(fdName, outputName)
		} else if fd == "-d" {
			decodeFilesFromDir(fdName, outputName)
		} else {
			fmt.Printf("Invalid option: %s\n", fd)
		}
	} else {
		fmt.Printf("Invalid option: %s\n", coding)
	}
}
