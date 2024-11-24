# Huffman coding

A CLI tool for encoding or decoding .txt files using the huffman algorithm.

(The output is in .huff, but it can easily be opened as a .txt file)

## Usage
```bash
{go run *.go|./huffman} <encode|decode> <-f|-d> <file/directory> [-o <output>]
```

* `<encode|decode>` - what to do with file/directory
* `<-f|-d>` - specify if next line is file or directory
* `<file/directory>` - the name of a specific file or folder
    * if `-f` - file extensions (".txt" or ".huff") don't matter
    * if `-d` - "./" in beggining or "/" at end doesn't matter
* `[-o <output>]` - (OPTIONAL) specify the file or directory where you want the output
    * if `-f` - file extensions (".txt" or ".huff") don't matter
    * if `-d` - "./" in beggining or "/" at end doesn't matter  
                if the folder doesnt exist, it will be created
## Prerequisites

- [Go](https://go.dev).

## Installation

To get started with this project, clone the repository and run (or compile, then run) using the `go` command.

```bash
git clone https://github.com/ikugox/huffman-coding.git
cd huffman-coding
```
You can run the code from `huffman-coding/` using
```bash
go run src/*.go
```
or
```bash
go build -o huffman src/*.go
```
and then run with
``` ./huffman ```
