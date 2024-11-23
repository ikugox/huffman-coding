package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 6 {
		running := os.Args[0]
		if strings.Contains(running, "/tmp/") {
			running = "go run *.go"
		}
		fmt.Println("Usage: " + running + " <encode|decode> <-f|-d> <file/directory> -o <output>")
		return
	}
	coding, fd, fdName, outputName := os.Args[1], os.Args[2], os.Args[3], os.Args[5]
	if coding == "encode" {
		if fd == "-f" {
			encodeFile(fdName, outputName)
		} else if fd == "-d" {
			encodeFilesFromDir(fdName, outputName)
		} else {
			fmt.Printf("Invalid option: %s", fd)
		}
	} else if coding == "decode" {
		if fd == "-f" {
			decodeFile(fdName, outputName)
		} else if fd == "-d" {
			decodeFilesFromDir(fdName, outputName)
		} else {
			fmt.Printf("Invalid option: %s", fd)
		}
	} else {
		fmt.Printf("Invalid option: %s", coding)
	}
}
